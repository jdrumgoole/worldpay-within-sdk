/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.worldpay.innovation.wpwithin;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.types.Device;
import com.worldpay.innovation.wpwithin.rpc.types.HCECard;
import com.worldpay.innovation.wpwithin.rpc.types.PaymentResponse;
import com.worldpay.innovation.wpwithin.rpc.types.Price;
import com.worldpay.innovation.wpwithin.rpc.types.PricePerUnit;
import com.worldpay.innovation.wpwithin.rpc.types.Service;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceDeliveryToken;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceDetails;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceMessage;
import com.worldpay.innovation.wpwithin.rpc.types.TotalPriceResponse;
import com.worldpay.innovation.wpwithin.thriftadapter.*;
import com.worldpay.innovation.wpwithin.types.WWDevice;
import com.worldpay.innovation.wpwithin.types.WWHCECard;
import com.worldpay.innovation.wpwithin.types.WWPaymentResponse;
import com.worldpay.innovation.wpwithin.types.WWPrice;
import com.worldpay.innovation.wpwithin.types.WWPricePerUnit;
import com.worldpay.innovation.wpwithin.types.WWService;
import com.worldpay.innovation.wpwithin.types.WWServiceDeliveryToken;
import com.worldpay.innovation.wpwithin.types.WWServiceDetails;
import com.worldpay.innovation.wpwithin.types.WWServiceMessage;
import com.worldpay.innovation.wpwithin.types.WWTotalPriceResponse;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;
import org.apache.thrift.transport.TTransportException;

/**
 *
 * @author worldpay
 */
public class WPWithinWrapperImpl implements WPWithinWrapper {

    private static final Logger log = Logger.getLogger(WPWithinWrapperImpl.class.getName());

    private String hostConfig;
    private Integer portConfig;
    private WPWithin.Client cachedClient;

    public WPWithinWrapperImpl(String host, Integer port) {
        this.hostConfig = host;
        this.portConfig = port;
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
}
