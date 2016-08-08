package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.WPWithinCallback;
import com.worldpay.innovation.wpwithin.rpc.types.*;
import com.worldpay.innovation.wpwithin.rpc.types.Error;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TSimpleServer;
import org.apache.thrift.transport.TServerSocket;
import org.apache.thrift.transport.TServerTransport;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;

import java.util.HashMap;
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
//kevTest();
            setupEventHandler();
            carwashConsumer();
// This was unncessary, and slowing down startup - KG
//            defaultDevice(client);
//            discovery(client);
//            initProducer(client);
//            broadcast(client);

            //doUI(client);

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

    static void kevTest() {

        try {

            TTransport transport = new TSocket("127.0.0.1", 9091);
            transport.open();

            TProtocol protocol = new TBinaryProtocol(transport);
            WPWithin.Client sdk = new WPWithin.Client(protocol);

            sdk.setup("TD_1", "TestDevice_1");

            Device dev = sdk.getDevice();

            Service roboWash = new Service();
            roboWash.setName("RoboWash");
            roboWash.setDescription("Car washed by robot");
            roboWash.setId(0);

            Price washPriceCar = new Price();
            washPriceCar.setUnitId(1);
            washPriceCar.setId(1);
            washPriceCar.setDescription("Car wash");
            washPriceCar.setUnitDescription("Single wash");
            washPriceCar.setPricePerUnit(new PricePerUnit(500, "GBP"));

            Price washPriceSUV = new Price();
            washPriceSUV.setUnitId(1);
            washPriceSUV.setId(2);
            washPriceSUV.setDescription("SUV wash");
            washPriceSUV.setUnitDescription("Single wash");
            washPriceSUV.setPricePerUnit(new PricePerUnit(650, "GBP"));

            Map<Integer,Price> prices = new HashMap();
            prices.put(0, washPriceCar);
            prices.put(1, washPriceSUV);

            roboWash.setPrices(prices);


            dev.putToServices(0, roboWash);


            try {
                System.out.println("trying to get services");
                //this.sdk.getDevice().getServices()
                Map svcDetails = dev.getServices();

                if(svcDetails != null) {

                    System.out.println("trying to get services - GOT SERVICES");

                    System.out.printf("Service Count: %d\n", dev.getServicesSize());
                    System.out.printf("Device Description: %s\n", dev.getDescription());


                    for(int i=0; i < dev.getServicesSize(); i++) {

                        Service svcDetail = (Service)svcDetails.get(i);

                        System.out.printf("%d - %s\n", svcDetail.getId(), svcDetail.getDescription());

                        Map<Integer, Price> pricesAfter = svcDetail.getPrices();

                        if(pricesAfter != null && pricesAfter.size() > 0) {

                            for(Map.Entry<Integer, Price> price : pricesAfter.entrySet()) {

                                System.out.printf("%d - %s (Unit:: Descr: %s, PricePerUnit: %d, CurrencyCode: %s, UnitId: %d)\n", price.getValue().getId(), price.getValue().getDescription(), price.getValue().getUnitDescription(), price.getValue().getPricePerUnit().getAmount(), price.getValue().getPricePerUnit().getCurrencyCode(), price.getValue().getUnitId());
                            }
                            System.out.println("trying to get services - DID IT OUTPUT PRICES");

                        } else {

                            System.out.printf("No prices found for service %d\n", svcDetail.getId());
                        }
                    }
                    System.out.println("trying to get services - DID IT OUTPUT SERVICES?");

                } else {
                    System.out.println("trying to get services - NO SERVICES FOUND");

                    System.out.println("No services found ");
                }


            } catch (Exception e) {

                e.printStackTrace();
            }
        } catch (Exception e) {

            e.printStackTrace();
        }
    }

    private static void carwashConsumer() {

        try {

            TTransport transport = new TSocket("127.0.0.1", 9090);
            transport.open();

            TProtocol protocol = new TBinaryProtocol(transport);
            WPWithin.Client sdk = new WPWithin.Client(protocol);

            sdk.setup("TD_1", "TestDevice_1");

            Device device = sdk.getDevice();

            Set<ServiceMessage> devices = sdk.deviceDiscovery(15000);

            if(devices != null && devices.size() > 0) {

                ServiceMessage[] arrMsgs = devices.toArray(new ServiceMessage[]{});

                ServiceMessage dev1 = arrMsgs[0];

                HCECard hceCard = new HCECard();
                hceCard.setFirstName("Bilbo");
                hceCard.setLastName("Baggins");
                hceCard.setExpMonth(11);
                hceCard.setExpYear(2018);
                hceCard.setCardNumber("5555555555554444");
                hceCard.setType("Card");
                hceCard.setCvc("113");

                sdk.initConsumer("http://", dev1.getHostname(), dev1.getPortNumber(),dev1.getUrlPrefix(), dev1.getServerId(), hceCard);

                Set<ServiceDetails> svcs = sdk.requestServices();

                if(svcs != null && svcs.size() > 0) {

                    ServiceDetails[] svcDetails = svcs.toArray(new ServiceDetails[]{});

                    int svcId = svcDetails[0].getServiceId();

                    Set<Price> prices = sdk.getServicePrices(svcId);

                    if(prices != null && prices.size() > 0) {

                        Price[] arrPrices = prices.toArray(new Price[]{});

                        TotalPriceResponse tpr = sdk.selectService(svcId, 10, arrPrices[0].getId());

                        PaymentResponse pr = sdk.makePayment(tpr);

                        if(pr != null) {

                            System.out.printf("SDT= %s |||| Paid = %d", pr.getServiceDeliveryToken().getKey(), pr.getTotalPaid());

                            sdk.beginServiceDelivery(device.getUid(), pr.serviceDeliveryToken, 10);
                        } else {

                            System.out.println("Payment response is null");
                        }

                    } else {

                        System.out.println("No prices found");
                    }
                } else {

                    System.out.println("No services found");
                }

            } else {

                System.out.println("No devices found");
            }


        } catch (Exception e) {

            e.printStackTrace();
        }
    }

    public static void setupEventHandler() {

        try {

            WPWithinHandler.Callbacks cb = new WPWithinHandler.Callbacks() {

                @Override
                public void beginServiceDelivery(String clientId, ServiceDeliveryToken serviceDeliveryToken, int unitsToSupply) throws Error, TException {

                }

                @Override
                public void endServiceDelivery(String clientId, ServiceDeliveryToken serviceDeliveryToken, int unitsReceived) throws Error, TException {

                }
            };

            final WPWithinCallback.Iface handler = new WPWithinHandler(cb);
            final WPWithinCallback.Processor processor = new WPWithinCallback.Processor(handler);

            Runnable simple = new Runnable() {
                public void run() {

                    try {
                        TServerTransport serverTransport = new TServerSocket(9091);
                        TServer server = new TSimpleServer(new TServer.Args(serverTransport).processor(processor));

                        // Use this for a multithreaded server
                        // TServer server = new TThreadPoolServer(new TThreadPoolServer.Args(serverTransport).processor(processor));

                        System.out.println("Starting the simple server...");
                        server.serve();

                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }
            };

            new Thread(simple).start();

        } catch (Exception x) {
            x.printStackTrace();
        }
    }
}