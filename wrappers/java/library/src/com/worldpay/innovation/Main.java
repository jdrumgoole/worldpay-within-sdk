package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.transport.TSSLTransportFactory;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;

public class Main {
    public static void main(String [] args) {

//        if (args.length != 1) {
//            System.out.println("Please enter 'simple' or 'secure'");
//            System.exit(0);
//        }

        // Override need for argument to be provided...
        args = new String[]{ "simple" };

        try {
            TTransport transport;
            if (args[0].contains("simple")) {
                transport = new TSocket("localhost", 9090);
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
                transport = TSSLTransportFactory.getClientSocket("localhost", 9090, 0, params);
            }

            TProtocol protocol = new TBinaryProtocol(transport);
            WPWithin.Client client = new WPWithin.Client(protocol);

            perform(client);

            transport.close();
        } catch (TException x) {
            x.printStackTrace();
        }
    }

    private static void perform(WPWithin.Client client) throws TException
    {
        try {
            client.initHTE("foo", "bar");
        } catch (Exception e) {

            e.printStackTrace();
        }
    }
}
