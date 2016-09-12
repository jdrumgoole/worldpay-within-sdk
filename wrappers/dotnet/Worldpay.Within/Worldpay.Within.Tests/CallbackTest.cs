using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using Common.Logging;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Worldpay.Innovation.WPWithin;

namespace Worldpay.Within.Tests
{
    [TestClass]
    public class CallbackTest
    {

        private static readonly ILog Log = LogManager.GetLogger<CallbackTest>();

        [TestMethod]
        public void SimpleCallbackTest()
        {
            using (WPWithinService service = new WPWithinService("localhost", 9091, 9092))
            {
                service.SetupDevice(this.GetType().Name, "Unit test from .NET wrapper");
                bool beginEventReceived = false;
                bool endEventReceived = false;

                service.OnBeginServiceDelivery += (id, token, supply) =>
                {
                    Log.InfoFormat("BeginServiceDelivery event received: id={0}, token={1}, supply={2}", id, token,
                        supply);
                    beginEventReceived = true;
                };
                service.OnEndServiceDelivery += (id, token, received) =>
                {
                    Log.InfoFormat("BeginServiceDelivery event received: id={0}, token={1}, supply={2}", id, token,
                        received);
                    endEventReceived = true;
                };

                Log.Info("Initialising Producer");
                service.InitProducer("cl_key", "srv_key");
                service.StartServiceBroadcast(2000);
                List<ServiceMessage> svcMsgs = service.DeviceDiscovery(2000).ToList();

                // Invoke a service and pay for it so the start and stop events are fired.
                Assert.IsNotNull(svcMsgs, "Discovered no services (null)");
                Assert.IsTrue(svcMsgs.Count>0, "Discovered no services (empty)");

                ServiceMessage serviceDescription = svcMsgs[0];
               
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
                    Assert.Fail("Unable to find service to invoke");
                }

                // Wait 5 seconds for the events to be triggered (could probably do this more elegantly, another day...)
                // If not triggered, then test fails.
                for (int i = 0; i < 100; i++)
                {
                    Thread.Sleep(50);
                    if (beginEventReceived && endEventReceived) break;
                }
                if (!beginEventReceived || !endEventReceived)
                {
                    Assert.Fail("BeginEvent received {0}; EndEvent received {1}", beginEventReceived, endEventReceived);
                }
            }
        }
    }
}
