using System;
using System.Text;
using System.Collections.Generic;
using Common.Logging;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Worldpay.Innovation.WPWithin;

namespace Worldpay.Within.Tests
{
    /// <summary>
    /// Summary description for ConsumerTest
    /// </summary>
    [TestClass]
    public class ConsumerTest : ThriftTest
    {

        private static readonly ILog Log = LogManager.GetLogger<ConsumerTest>();

        #region Additional test attributes
        //
        // You can use the following additional attributes as you write your tests:
        //
        // Use ClassInitialize to run code before running the first test in the class
        // [ClassInitialize()]
        // public static void MyClassInitialize(TestContext testContext) { }
        //
        // Use ClassCleanup to run code after all tests in a class have run
        // [ClassCleanup()]
        // public static void MyClassCleanup() { }
        //
        // Use TestInitialize to run code before running each test 
        // [TestInitialize()]
        // public void MyTestInitialize() { }
        //
        // Use TestCleanup to run code after each test has run
        // [TestCleanup()]
        // public void MyTestCleanup() { }
        //
        #endregion

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

        [TestMethod]
        public void CarWashConsumer()
        {
            string defaultDeviceName = Environment.MachineName;
            string defaultDeviceDescription =
                $".net wrapper unit test.  Class: {this.GetType().Name}.  Method: {System.Reflection.MethodBase.GetCurrentMethod().Name}";
            ThriftClient.SetupDevice(defaultDeviceName, defaultDeviceDescription);
            foreach (ServiceMessage service in ThriftClient.DeviceDiscovery(20000))
            {
                Log.InfoFormat("Found service: ", service);
            }

        }
    }
}
