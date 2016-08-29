/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.worldpay.innovation.wpwithin;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.launcher.*;
import com.worldpay.innovation.wpwithin.thriftadapter.*;
import com.worldpay.innovation.wpwithin.types.*;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;
import org.apache.thrift.transport.TTransportException;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 *
 * @author worldpay
 */
public class WPWithinWrapperImpl implements WPWithinWrapper {

    private static final Logger log = Logger.getLogger(WPWithinWrapperImpl.class.getName());

    private String hostConfig;
    private Integer portConfig;
    private WPWithin.Client cachedClient;
    private Launcher launcher;

    private void startRPCAgent(int port) {

        launcher = new Launcher();

        Map<OS, PlatformConfig> launchConfig = new HashMap<>(3);

        PlatformConfig winConfig = new PlatformConfig();
        winConfig.setCommand(Architecture.IA32, String.format("./rpc-agent/rpc-agent-win-32 -port=%d", port));
        winConfig.setCommand(Architecture.X86_64, String.format("./rpc-agent/rpc-agent-win-64 -port=%d", port));
        winConfig.setCommand(Architecture.ARM, String.format("./rpc-agent/rpc-agent-win-armv5 -port=%d", port));
        launchConfig.put(OS.WINDOWS, winConfig);

        PlatformConfig linuxConfig = new PlatformConfig();
        linuxConfig.setCommand(Architecture.IA32, String.format("./rpc-agent/rpc-agent-linux-32 -port=%d", port));
        linuxConfig.setCommand(Architecture.X86_64, String.format("./rpc-agent/rpc-agent-linux-64 -port=%d", port));
        linuxConfig.setCommand(Architecture.ARM, String.format("./rpc-agent/rpc-agent-linux-armv5 -port=%d", port));
        launchConfig.put(OS.LINUX, linuxConfig);


        PlatformConfig macConfig = new PlatformConfig();
        macConfig.setCommand(Architecture.IA32, String.format("./rpc-agent/rpc-agent/rpc-agent-mac-32 -port=%d", port));
        macConfig.setCommand(Architecture.X86_64, String.format("./rpc-agent/rpc-agent-mac-64 -port=%d", port));
        macConfig.setCommand(Architecture.ARM, String.format("./rpc-agent/rpc-agent-mac-armv5 -port=%d", port));
        launchConfig.put(OS.MAC, macConfig);


        Listener listener = new Listener() {

            @Override
            public void onApplicationExit(int exitCode, String stdOutput, String errOutput) {

                System.out.printf("RPC Agent did exit with code: %d\n", exitCode);

                try {

                    String output = launcher.getStdOutput();
                    String error = launcher.getErrorOutput();

                    System.out.println("Output: " + output);
                    System.out.println("Error: " + error);


                } catch (Exception e) {

                    e.printStackTrace();
                }
            }
        };

        try {

            launcher.startProcess(launchConfig, listener);

        } catch (WPWithinGeneralException ioe) {

            ioe.printStackTrace();
        }
    }

    public WPWithinWrapperImpl(String host, Integer port, boolean startRPCAgent) {

        this.hostConfig = host;
        this.portConfig = port;

        if(startRPCAgent) {

            startRPCAgent(port);
        }

        setClientIfNotSet();
    }
    
    private void setClientIfNotSet() {
        if(this.cachedClient == null) {
            this.cachedClient = openRpcListener();
        }        
    }
    
    private WPWithin.Client getClient() {
       setClientIfNotSet();
       return this.cachedClient;
    }

    private WPWithin.Client openRpcListener() {

        TTransport transport = new TSocket(hostConfig, portConfig);

        try {
            transport.open();
        } catch (TTransportException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Could not open transport socket", ex);
        }

        TProtocol protocol = new TBinaryProtocol(transport);
        WPWithin.Client client = new WPWithin.Client(protocol);

        return client;
    }

    @Override
    public void setup(String name, String description) throws WPWithinGeneralException {
        try {
            getClient().setup(name, description);
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Failure to setup in the wrapper", ex);
            throw new WPWithinGeneralException("Failure to setup in the wrapper");
        }
    }

    @Override
    public void addService(WWService theService) throws WPWithinGeneralException {

        Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.INFO, "About to add service");
        try {
            getClient().addService(ServiceAdapter.convertWWService(theService));
        } catch(TException ex) {
            throw new WPWithinGeneralException("Add service to producer failed with Rpc call to the SDK lower level");
        }
        Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.INFO, "Should have successfully added service");

    }

    @Override
    public void removeService(WWService svc) throws WPWithinGeneralException {
        try {
            getClient().removeService(ServiceAdapter.convertWWService(svc));
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Removal of service failed in the wrapper", ex);
            throw new WPWithinGeneralException("Removal of service failed in the wrapper");
        }
    }

    @Override
    public void initConsumer(String scheme, String hostname, Integer port, String urlPrefix, String serverId, WWHCECard hceCard) throws WPWithinGeneralException {
        try {
            getClient().initConsumer(scheme, hostname, port, urlPrefix, serverId, HCECardAdapter.convertWWHCECard(hceCard));
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Initiating the consumer failed in the wrapper", ex);
            throw new WPWithinGeneralException("Initiating the consumer failed in the wrapper");
        }
    }

    @Override
    public void initProducer(String merchantClientKey, String merchantServiceKey) throws WPWithinGeneralException {
        try {
            getClient().initProducer(merchantClientKey, merchantServiceKey);
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Initiating the producer failed in the wrapper", ex);
            throw new WPWithinGeneralException("Initiating the producer failed in the wrapper");
        }
    }

    @Override
    public WWDevice getDevice() throws WPWithinGeneralException {
        try {
            return DeviceAdapter.convertDevice(getClient().getDevice());
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Get device in wrapper failed", ex);
            throw new WPWithinGeneralException("Get device in wrapper failed");
        }
    }

    @Override
    public void startServiceBroadcast(Integer timeoutMillis) throws WPWithinGeneralException {
        try {
            getClient().startServiceBroadcast(timeoutMillis);
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Start service broadcast in wrapper failed", ex);
            throw new WPWithinGeneralException("Start service broadcast in wrapper failed");
        }
    }

    @Override
    public void stopServiceBroadcast() throws WPWithinGeneralException {
        try {
            getClient().stopServiceBroadcast();
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Stop service broadcast failed", ex);
            throw new WPWithinGeneralException("Stop service broadcast failed");
        }
    }

    @Override
    public Set<WWServiceMessage> deviceDiscovery(Integer timeoutMillis) throws WPWithinGeneralException {
        try {
            return ServiceMessageAdapter.convertServiceMessages(getClient().deviceDiscovery(timeoutMillis));
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Failed device discovery in wrapper", ex);
            throw new WPWithinGeneralException("Failed device discovery in wrapper");
        }
    }

    @Override
    public Set<WWServiceDetails> requestServices() throws WPWithinGeneralException {
            
        try {
            return ServiceDetailsAdapter.convertServiceDetails(getClient().requestServices());
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Request Services failed in wrapper", ex);
            throw new WPWithinGeneralException("Request Services failed in wrapper");
        }

    }

    @Override
    public Set<WWPrice> getServicePrices(Integer serviceId) throws WPWithinGeneralException {
        try {
            return PriceAdapter.convertServicePrices(getClient().getServicePrices(serviceId));
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Get Service Prices failed in wrapper", ex);
            throw new WPWithinGeneralException("Get Service Prices failed in wrapper");
        }
    }
    
    @Override
    public WWTotalPriceResponse selectService(Integer serviceId, Integer numberOfUnits, Integer priceId) throws WPWithinGeneralException {
        try {
            return TotalPriceResponseAdapter.convertTotalPriceResponse(getClient().selectService(serviceId, numberOfUnits, priceId));
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Select service failed in wrapper", ex);
            throw new WPWithinGeneralException("Select service failed in wrapper");
        }
    }

    @Override
    public WWPaymentResponse makePayment(WWTotalPriceResponse request) throws WPWithinGeneralException {
  
        try {
            return PaymentResponseAdapter.convertPaymentResponse(getClient().makePayment(TotalPriceResponseAdapter.convertWWTotalPriceResponse(request)));
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Failed to make payment in the wrapper", ex);
            throw new WPWithinGeneralException("Failed to make payment in the wrapper");
        }
    }

    @Override
    public void beginServiceDelivery(String clientId, WWServiceDeliveryToken serviceDeliveryToken, Integer unitsToSupply) throws WPWithinGeneralException {
        try {
            getClient().beginServiceDelivery(clientId, ServiceDeliveryTokenAdapter.convertWWServiceDeliveryToken(serviceDeliveryToken), unitsToSupply);
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Failed to begin Service Delivery in the wrapper", ex);
            throw new WPWithinGeneralException("Failed to begin Service Delivery in the wrapper");
        }
    }

    @Override
    public void endServiceDelivery(String clientId, WWServiceDeliveryToken serviceDeliveryToken, Integer unitsReceived) throws WPWithinGeneralException {
        try {
            getClient().endServiceDelivery(clientId, ServiceDeliveryTokenAdapter.convertWWServiceDeliveryToken(serviceDeliveryToken), unitsReceived);
        } catch (TException ex) {
            Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.SEVERE, "Failed to end Service Delivery in the wrapper", ex);
            throw new WPWithinGeneralException("Failed to end Service Delivery in the wrapper");
        }
    }

    @Override
    public void stopRPCAgent() {

        try {

            launcher.stopProcess();

        } catch (Exception e) {

            throw new RuntimeException(e);
        }
    }
}
