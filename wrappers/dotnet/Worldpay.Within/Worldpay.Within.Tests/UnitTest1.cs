using System;
using System.Collections.Generic;
using System.Linq;
using Common.Logging;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Worldpay.Innovation.WPWithin;

namespace Worldpay.Within.Tests
{
    [TestClass]
    public class UnitTest1
    {
        private static readonly ILog Log = LogManager.GetLogger<UnitTest1>();

        [TestMethod]
        public void SendSimpleMessage()
        {
            const string host = "127.0.0.1";
            const int port = 9091;
            using (WPWithinService client = new WPWithinService(host, port))
            {
                DefaultDevice(client);
                InitProducer(client);
                Broadcast(client);
                Discovery(client);
                Log.Info("All done, closing transport");
            }
        }


        private static void Discovery(WPWithinService client)
        {
            IEnumerable<ServiceMessage> svcMsgs = client.DeviceDiscovery(2000);

            if (svcMsgs != null)
            {
                foreach (ServiceMessage svcMsg in svcMsgs)
                {
                    Log.InfoFormat("{0} - {1} - {2} - {3}", svcMsg.DeviceDescription, svcMsg.Hostname, svcMsg.PortNumber, svcMsg.ServerId);
                }
            }
            else
            {
                Log.Info("Broadcast ok, but no services found");
            }
        }

        private static void Broadcast(WPWithinService client)
        {
            client.StartServiceBroadcast(2000);
        }

        private static void DefaultDevice(WPWithinService client)
        {
            client.SetupDevice("DotNet RPC client", "This is coming from C# via Thrift RPC.");
        }

        private static void InitProducer(WPWithinService client)
        {
            client.InitProducer("cl_key", "srv_key");
        }

    }
}
