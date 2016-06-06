package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.types.Device;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceMessage;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.transport.TSSLTransportFactory;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;

import java.util.Set;

public class Main {
    public static void main(String [] args) {

//        if (args.length != 1) {
//            System.out.println("Please enter 'simple' or 'secure'");
//            System.exit(0);
//        }

        // Override need for argument to be provided...
        args = new String[]{ "simple" };

        String host = "127.0.0.1";
        int port = 9091;

        try {
            TTransport transport;
            if (args[0].contains("simple")) {
                transport = new TSocket(host, port);
                transport.open();
            }
            else {
        /*
         * Similar to the server, you can use the parameters to setup client parameters or
         * use the default settings. On the client side, you will need a TrustStore which
         * contains the trusted certificate along with the public key.
         * For this example it's a self-signed cert.
         */
                TSSLTransportFactory.TSSLTransportParameters params = new TSSLTransportFactory.TSSLTransportParameters();
                params.setTrustStore("../../lib/java/test/.truststore", "thrift", "SunX509", "JKS");
        /*
         * Get a client transport instead of a server transport. The connection is opened on
         * invocation of the factory method, no need to specifically call open()
         */
                transport = TSSLTransportFactory.getClientSocket(host, port, 0, params);
            }

            TProtocol protocol = new TBinaryProtocol(transport);
            WPWithin.Client client = new WPWithin.Client(protocol);

//            perform(client);
            System.out.println("before deafult device");
            defaultDevice(client);
            initProducer(client);
            System.out.println("before broadcast");
            broadcast(client);
            transport.close();
        } catch (TException x) {
            x.printStackTrace();
        }
    }

    private static void perform(WPWithin.Client client) throws TException
    {
        try {

            Set<ServiceMessage> svcMsgs = client.serviceDiscovery(20000);

            if(svcMsgs != null && svcMsgs.size() > 0) {

                for(ServiceMessage svcMsg : svcMsgs) {

                    System.out.printf("%s - %s - %d - %s", svcMsg.getDeviceDescription(), svcMsg.getHostname(), svcMsg.getPortNumber(), svcMsg.getServerId());
                }
            }

            //System.out.println("Device name: " + device.getName());

        } catch (Exception e) {

            e.printStackTrace();
        }
    }

    private static void broadcast(WPWithin.Client client) throws TException {

        try {

            client.startServiceBroadcast(20000);

        } catch (Exception e) {

            e.printStackTrace();
        }
    }

    private static void defaultDevice(WPWithin.Client client) throws TException {

        try {

            client.setup("name.rpc.test", "descroption.rpc.test");

        } catch (Exception e) {

            e.printStackTrace();
        }
    }

    private static void initProducer(WPWithin.Client client) throws TException {

        try {

            System.out.println("before init hte");
            client.initHTE("cl_key", "srv_key");
            System.out.println("before init producer");
            client.initProducer();

        } catch (Exception e) {

            e.printStackTrace();
        }
    }
}
