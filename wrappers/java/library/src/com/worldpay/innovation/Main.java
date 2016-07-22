package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.types.*;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;

import java.util.Map;
import java.util.Set;
import java.util.logging.ConsoleHandler;
import java.util.logging.Level;
import java.util.logging.Logger;

public class Main {
    private static final Logger log= Logger.getLogger( Main.class.getName() );
    
    public static void main(String [] args) {

        String host = "127.0.0.1";
        int port = 9091;

        try {
            TTransport transport = new TSocket(host, port);
            transport.open();

            TProtocol protocol = new TBinaryProtocol(transport);
            WPWithin.Client client = new WPWithin.Client(protocol);

// This was unncessary, and slowing down startup - KG
//            defaultDevice(client);
//            discovery(client);
//            initProducer(client);
//            broadcast(client);

            doUI(client);

            transport.close();

        } catch (TException x) {
            x.printStackTrace();
        }
    }

    private static void discovery(WPWithin.Client client) throws TException
    {
        try {

            Set<ServiceMessage> svcMsgs = client.deviceDiscovery(20000);

            if(svcMsgs != null && svcMsgs.size() > 0) {

                ServiceMessage svcMsg = svcMsgs.iterator().next();

                System.out.printf("%s - %s - %d - %s", svcMsg.getDeviceDescription(), svcMsg.getHostname(), svcMsg.getPortNumber(), svcMsg.getServerId());

                HCECard card = new HCECard("Bilbo", "Baggins", 11, 2018, "5555555555554444", "Card", "113");

                client.initConsumer("http://", svcMsg.getHostname(), svcMsg.getPortNumber(), svcMsg.getUrlPrefix(), svcMsg.getServerId(), card);

                Set<ServiceDetails> svcDetails = client.requestServices();

                if(svcDetails != null) {

                    for(ServiceDetails svcDetail : svcDetails) {

                        System.out.printf("%d - %s\n", svcDetail.getServiceId(), svcDetail.getServiceDescription());

                        Set<Price> prices = client.getServicePrices(svcDetail.getServiceId());

                        if(prices != null && prices.size() > 0) {

                            for(Price price : prices) {

                                System.out.printf("%d - %s (Unit:: Descr: %s, PricePerUnit: %d, CurrencyCode: %s, UnitId: %d)", price.getId(), price.getDescription(), price.getUnitDescription(), price.getPricePerUnit().getAmount(), price.getPricePerUnit().getCurrencyCode(), price.getUnitId());
                            }
                        } else {

                            System.out.printf("No prices found for service %d\n", svcDetail.getServiceId());
                        }
                    }
                } else {

                    System.out.printf("No services found @ %s:%d\n", svcMsg.getHostname(), svcMsg.getPortNumber());
                }

            } else {

                System.out.printf("No devices found on network\n");
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

//            client.initHTE();
            client.initProducer("cl_key", "srv_key");

        } catch (Exception e) {

            e.printStackTrace();
        }
    }
    
    private static void doUI(WPWithin.Client client) {
        
        log.setLevel(Level.FINE);
        ConsoleHandler handler = new ConsoleHandler();
        handler.setLevel(Level.FINE);
        log.addHandler(handler);
        
        
        
        log.log( Level.INFO, "Starting UI");
        (new MenuSystem()).doUI(client);
        log.log( Level.INFO, "FINISHING UI");
    }

    public static void testSvcIds() {

        Service svc1 = new Service(0, "1 - name", "1 - description");
        svc1.setId(2);

        Service svc2 = new Service(0, "2 - name", "2 - description");
        svc2.setId(33);


        String host = "127.0.0.1";
        int port = 9091;

        try {

            TTransport transport = new TSocket(host, port);
            transport.open();

            TProtocol protocol = new TBinaryProtocol(transport);
            WPWithin.Client client = new WPWithin.Client(protocol);

            client.setup("ConorH-WP", "Running on Macbook Pro (Java)");

            client.addService(svc1);
            client.addService(svc2);

            Map<Integer, Service> svcs = client.getDevice().getServices();

            for (Map.Entry<Integer, Service> svc : svcs.entrySet()) {

                System.out.printf("%d - %s - %s\n", svc.getValue().getId(), svc.getValue().getName(), svc.getValue().getDescription());
            }

            transport.close();

        } catch (TException x) {
            x.printStackTrace();
        }
    }
}
