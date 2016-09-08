using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Worldpay.Innovation.WPWithin.AgentManager;
using System.Threading;

namespace Worldpay.Within.Tests
{
    [TestClass]
    public class RpcAgentManagerTest
    {
        [TestMethod]
        public void StartAndStopViaEvent()
        {
            RpcAgentManager mgr = new RpcAgentManager();

            bool started = false;
            mgr.OnStarted += (s,e) =>
            {
                mgr.StopThriftRpcAgentProcess();
                started = true;
            };
            mgr.StartThriftRpcAgentProcess();

            int retries = 0;
            while(!started && retries<10)
            {
                Thread.Sleep(500);
                retries++;
            }
            if (!started) Assert.Fail("Thrift RPC Agent didn't start");
        }

        

    }
}
