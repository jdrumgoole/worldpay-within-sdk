using System;
using System.Configuration;
using System.Diagnostics;
using System.IO;
using Common.Logging;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Worldpay.Innovation.WPWithin;

namespace Worldpay.Within.Tests
{
    public class ThriftTest
    {
        private static readonly ILog Log = LogManager.GetLogger<ProducerTest>();

        public static readonly string RpcAgentPathProperty = "ThriftRpcAgent.Path";
        public static readonly string RpcAgentHostProperty = "ThriftRpcAgent.Host";
        public static readonly string RpcAgentHostPropertyDefault = "127.0.0.1";
        public static readonly string RpcAgentPortProperty = "ThriftRpcAgent.Port";
        public static readonly int RpcAgentPortPropertyDefault = 9091;
        public static readonly string RpcAgentProtocolProperty = "ThriftRpcAgent.Protocol";
        public static readonly string RpcAgentProtocolPropertyDefault = "binary";

        private static Process _thriftRpcProcess;
        private static string _rpcAgentPath;

        protected WPWithinService ThriftClient { get; private set; }

        public TestContext TestContext { get; set; }

        public static string RpcAgentServiceHost
            => ConfigurationManager.AppSettings[RpcAgentHostProperty] ?? RpcAgentHostPropertyDefault;

        public static int RpcAgentServicePort
        {
            get
            {
                string portString = ConfigurationManager.AppSettings[RpcAgentPortProperty];
                int port;
                if (portString == null || !int.TryParse(portString, out port))
                {
                    port = RpcAgentPortPropertyDefault;
                }
                return port;
            }
        }

        public static string RpcAgentPath
        {
            get
            {
                if (_rpcAgentPath != null)
                {
                    return _rpcAgentPath;
                }
                string agentPath = ConfigurationManager.AppSettings[RpcAgentPathProperty];
                if (agentPath == null)
                {
                    DirectoryInfo parent = new DirectoryInfo(".");
                    Log.InfoFormat("No {0} property found in App.Confing, searching for it relative to {1}",
                        RpcAgentPathProperty, parent.FullName);
                    const string sdkDir = "worldpay-within-sdk";
                    while (parent != null && !parent.Name.Equals(sdkDir))
                    {
                        parent = parent.Parent;
                    }
                    if (parent == null)
                    {
                        throw new Exception("Unable to locate " + sdkDir +
                                            " override with property " + RpcAgentPathProperty + "property in App.config");
                    }
                    _rpcAgentPath =
                        new FileInfo(string.Join(Path.DirectorySeparatorChar.ToString(), parent.FullName, "applications",
                            "rpc-agent",
                            "rpc-agent.exe")).FullName;
                }
                return _rpcAgentPath;
            }
        }

        [TestInitialize]
        public void CreateClient()
        {
            ThriftClient = new WPWithinService(RpcAgentServiceHost, RpcAgentServicePort);
        }

        [TestCleanup]
        public void DisposeClient()
        {
            ThriftClient.Dispose();
        }

        public static void StartThriftRpcAgentProcess(TestContext context)
        {
            Log.InfoFormat("Attempting to launch Thrift RPC Agent {0}", RpcAgentPath);

            ProcessStartInfo thriftRpcService =
                new ProcessStartInfo(RpcAgentPath, string.Join(" ",
                    "-host", RpcAgentServiceHost,
                    "-port", RpcAgentServicePort,
                    "-protocol",
                    ConfigurationManager.AppSettings[RpcAgentProtocolProperty] ?? RpcAgentProtocolPropertyDefault
                    ));
            _thriftRpcProcess = Process.Start(thriftRpcService);
        }

        public static void KillThriftRpcAgentProcess()
        {
            _thriftRpcProcess.Kill();
        }
    }
}