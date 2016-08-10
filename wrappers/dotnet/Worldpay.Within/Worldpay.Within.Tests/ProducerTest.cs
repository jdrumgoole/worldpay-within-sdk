using Common.Logging;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Worldpay.Innovation.WPWithin;

namespace Worldpay.Within.Tests
{
    [TestClass]
    public class ProducerTest : ThriftTest
    {
        private static readonly ILog Log = LogManager.GetLogger<ProducerTest>();


        [TestMethod]
        public void SendSimpleMessage()
        {
            ThriftClient.SetupDevice("DotNet RPC client", "This is coming from C# via Thrift RPC.");
            Log.Info("Initialising Producer");
            ThriftClient.InitProducer("cl_key", "srv_key");
            ThriftClient.StartServiceBroadcast(2000);
            var svcMsgs = ThriftClient.DeviceDiscovery(2000);

            if (svcMsgs != null)
            {
                foreach (ServiceMessage svcMsg in svcMsgs)
                {
                    Log.InfoFormat("{0} - {1} - {2} - {3}", svcMsg.DeviceDescription, svcMsg.Hostname, svcMsg.PortNumber,
                        svcMsg.ServerId);
                }
            }
            else
            {
                Log.Info("Broadcast ok, but no services found");
            }
            Log.Info("All done, closing transport");
        }

        [ClassInitialize]
        public static void StartThriftRpcService(TestContext context)
        {
            StartThriftRpcAgentProcess(context);
        }

        [ClassCleanup]
        public static void StopThriftRpcService()
        {
            KillThriftRpcAgentProcess();
        }
    }
}

