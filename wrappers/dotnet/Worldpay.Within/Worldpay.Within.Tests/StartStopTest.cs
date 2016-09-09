using System;
using Common.Logging;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Worldpay.Innovation.WPWithin;

namespace Worldpay.Within.Tests
{
    [TestClass]
    public class StartStopTest
    {
        private static readonly ILog Log = LogManager.GetLogger<StartStopTest>();

        [TestMethod]
        public void StartAndStop()
        {
            using (WPWithinService service = new WPWithinService("localhost", 9091, 9092))
            {
                Log.InfoFormat("Successfully created service {0}", service);
            }
        }
    }
}
