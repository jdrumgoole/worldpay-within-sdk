package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceMessage;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;

import java.util.Set;

public class Main {
    public static void main(String [] args) {

        String host = "127.0.0.1";
        int port = 9091;

        try {
            TTransport transport = new TSocket(host, port);
            transport.open();

            TProtocol protocol = new TBinaryProtocol(transport);
            WPWithin.Client client = new WPWithin.Client(protocol);

            defaultDevice(client);
            discovery(client);
//            initProducer(client);
//            broadcast(client);

            transport.close();

        } catch (TException x) {
            x.printStackTrace();
        }
    }

    private static void discovery(WPWithin.Client client) throws TException
    {
        try {

            Set<ServiceMessage> svcMsgs = client.serviceDiscovery(20000);

            if(svcMsgs != null && svcMsgs.size() > 0) {

                for(ServiceMessage svcMsg : svcMsgs) {

                    System.out.printf("%s - %s - %d - %s", svcMsg.getDeviceDescription(), svcMsg.getHostname(), svcMsg.getPortNumber(), svcMsg.getServerId());
                }
            }

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

            client.setup("Java RPC client", "This is coming from Java via Thrift RPC.");

        } catch (Exception e) {

            e.printStackTrace();
        }
    }

    private static void initProducer(WPWithin.Client client) throws TException {

        try {

            client.initHTE("cl_key", "srv_key");
            client.initProducer();

        } catch (Exception e) {

            e.printStackTrace();
        }
    }
}
