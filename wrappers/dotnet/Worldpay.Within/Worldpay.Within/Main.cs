using Common.Logging;
using System;
using Thrift.Collections;
using Thrift.Protocol;
using Thrift.Transport;
using WWRpc = Worldpay.Innovation.WPWithin.Rpc.WPWithin;
using WWRpcTypes = Worldpay.Innovation.WPWithin.Rpc.Types;

namespace Worldpay.Innovation.WPWithin
{

    public class Main
    {
        private static readonly ILog Log = LogManager.GetLogger<Main>();

        public void SendSimpleMessage()
        {
            const string host = "127.0.0.1";
            const int port = 9091;

            Log.InfoFormat("Opening TSocket to {0}:{1}", host, port);
            TTransport transport = new TSocket(host, port);
            transport.Open();

            TProtocol protocol = new TBinaryProtocol(transport);
            WWRpc.Client client = new WWRpc.Client(protocol);

            DefaultDevice(client);
            InitProducer(client);
            Broadcast(client);
            Discovery(client);

            Log.Info("All done, closing transport");
            transport.Close();
        }


        private static void Discovery(WWRpc.Client client)
        {
            THashSet<WWRpcTypes.ServiceMessage> svcMsgs = client.serviceDiscovery(2000);

            if (svcMsgs != null && svcMsgs.Count!=0)
            {
                foreach (WWRpcTypes.ServiceMessage svcMsg in svcMsgs)
                {
                    Log.InfoFormat("{0} - {1} - {2} - {3}", svcMsg.DeviceDescription, svcMsg.Hostname, svcMsg.PortNumber, svcMsg.ServerId);
                }
            } else
            {
                Log.Info("Broadcast ok, but no services found");
            }
        }

        private static void Broadcast(WWRpc.Client client)
        {
            client.startServiceBroadcast(2000);
        }

        private static void DefaultDevice(WWRpc.Client client)
        {
            client.setup("DotNet RPC client", "This is coming from C# via Thrift RPC.");
        }

        private static void InitProducer(WWRpc.Client client)
        {
            client.initHTE("cl_key", "srv_key");
            client.initProducer();
        }
    }
}
