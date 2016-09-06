package com.worldpay.innovation.wpwithin.eventlistener;

import com.worldpay.innovation.wpwithin.rpc.WPWithinCallback;
import org.apache.thrift.server.TServer;
import org.apache.thrift.server.TThreadPoolServer;
import org.apache.thrift.transport.TServerSocket;
import org.apache.thrift.transport.TServerTransport;

/**
 * A server to listen for events sent from the SDK Core via RPC
 */
public class EventServer {

    TServerTransport serverTransport;
    //TServer server = new TSimpleServer(new TServer.Args(serverTransport).processor(processor));

    // Use this for a multithreaded server
    TServer server;

    public void start(EventListener listener, final int port) {

        try {

            final WPWithinCallback.Iface handler = new EventHandler(listener);
            final WPWithinCallback.Processor processor = new WPWithinCallback.Processor(handler);

            serverTransport = new TServerSocket(port);
//            server = new TSimpleServer(new TServer.Args(serverTransport).processor(processor));

            // Use this for a multithreaded server
            server = new TThreadPoolServer(new TThreadPoolServer.Args(serverTransport).processor(processor));

            Runnable simple = new Runnable() {

                public void run() {

                    try {

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

    public void stop() {

        if(server != null) {

            server.setShouldStop(true);
        }
    }
}
