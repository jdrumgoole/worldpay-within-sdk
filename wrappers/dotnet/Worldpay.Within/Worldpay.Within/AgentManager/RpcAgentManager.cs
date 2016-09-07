using Common.Logging;
using System;
using System.Configuration;
using System.Diagnostics;
using System.IO;
using System.Runtime.InteropServices;

namespace Worldpay.Innovation.WPWithin.AgentManager
{
    /// <summary>
    /// Manages the lifecycle of a Thrift RPC Agent (see /applications/rpc-agent).
    /// </summary>
    /// <remarks>
    /// <para>The Thift RPC agent is what this code talks to in order to communicate with the other participant in the conversation.</para>
    /// <para>For example, if we are writing a .NET producer, we commnunicate with the consumer by invoking the Thrift RPC Agent (typically
    /// located locally) which then talks to the consumer service (typically located remotely). </para>
    /// <para>
    /// As the Thrift RPC Agent runs as separate process, we have to invoke it as if it were a separate tool, rather than loading in to the 
    /// same address space.</para>
    /// <list type="bullet">
    /// <item><description>Start a binary and keep reference to process.</description></item>
    /// <item><description>Using the reference during start, be able to stop the process.</description></item>
    /// <item><description>Monitor std and err outputs.  Surface errors to calling application</description></item>
    /// <item><description>Monitor the exit code of the RPC Agent. 0 = success, not 0 = error ( will be updating this app to return the appropriate exit codes)</description></item>
    /// <item><description>Ability to support multiple operating systems. Most of the wrappers will be able to run on multiple OS’s we should aim to support { Linux, Windows, Mac OS }</description></item>
    /// <item><description>We should support 3 architectures { ia32, x86-64, arm }</description></item>
    /// <item><description>The launcher should be able to determine the OS and CPU architecture that the application is running on.This will enable the correct selection of the binary to launch. Please note that the wrapper distribution is to include the binaries for the various platforms / architectures.</description></item>
    /// <item><description>We will agree on a filename convention such as rpc-agent-os-cpu(.ext) e.g.rpc-agent-win-64.exe</description></item>
    /// </list>
    /// </remarks>
    public class RpcAgentManager
    {
        private static readonly ILog Log = LogManager.GetLogger<RpcAgentManager>();
        private static readonly ILog ThriftRpcLog = LogManager.GetLogger("ThriftRpcAgent");

        /// <summary>
        /// Delegate used for <see cref="RpcAgentManager.RpcAgentMessages"/> and <see cref="RpcAgentManager.RpcAgentErrors"/>.
        /// </summary>
        /// <param name="process">The process that the RPC Agent process is running under.</param>
        /// <param name="message">The message received from the RPC Agent process.</param>
        public delegate void ThriftRpcAgentOutput(Process process, string message);

        /// <summary>
        /// Invoked whenever a message is sent to the RPC Agent process's standard output stream.
        /// </summary>
        public event ThriftRpcAgentOutput RpcAgentMessages;

        /// <summary>
        /// Invoked whenever a message is sent ot the RPC Agent process's standard error stream.
        /// </summary>
        public event ThriftRpcAgentOutput RpcAgentErrors;

        /// <summary>
        /// Invoked when the Thrift RPC Agent has successfully been started.
        /// </summary>
        public event EventHandler ThriftRpcAgentStarted;

        /// <summary>
        /// Invoked when the Thrift RPC Agent has successfully been stopped, or has crashed out.
        /// </summary>
        public event EventHandler ThriftRpcAgentExited;

        /// <summary>
        /// The application config property name for the full file path to the Thrift RPC Agent.
        /// </summary>
        public static readonly string RpcAgentPathProperty = "ThriftRpcAgent.Path";
        /// <summary>
        /// The application config property name for the host name to bind the Thrift RPC Agent to.
        /// </summary>
        public static readonly string RpcAgentHostProperty = "ThriftRpcAgent.Host";
        /// <summary>
        /// The default value for the full file path to the Thrift RPC Agent.
        /// </summary>
        public static readonly string RpcAgentHostPropertyDefault = "127.0.0.1";
        /// <summary>
        /// The application config property name for the port to launch the Thrift RPC Agent on.
        /// </summary>
        public static readonly string RpcAgentPortProperty = "ThriftRpcAgent.Port";
        /// <summary>
        /// The default value for the port to launch the Thrift RPC Agent on.
        /// </summary>
        public static readonly int RpcAgentPortPropertyDefault = 9091;
        /// <summary>
        /// The application config property name for the Thrift protocol to use to connect to the Thrift RPC Agent.
        /// </summary>
        public static readonly string RpcAgentProtocolProperty = "ThriftRpcAgent.Protocol";
        /// <summary>
        /// The default value for the protocol to use to connect to the Thrift RPC Agent.
        /// </summary>
        public static readonly string RpcAgentProtocolPropertyDefault = "binary";

        private Process _thriftRpcProcess;
        private string _rpcAgentPath;

        /// <summary>
        /// Retrieves the RPC Agent host property from application config or provides default value.
        /// </summary>
        public static string RpcAgentServiceHost
            => ConfigurationManager.AppSettings[RpcAgentHostProperty] ?? RpcAgentHostPropertyDefault;

        /// <summary>
        /// Retrieves the RPC Agent protocol property from application config or provides default value.
        /// </summary>
        public static string RpcAgentProtcol => ConfigurationManager.AppSettings[RpcAgentProtocolProperty] ?? RpcAgentProtocolPropertyDefault;

        /// <summary>
        /// Retrieves the RPC Agent port property from application config or provides default value.
        /// </summary>
        public int RpcAgentServicePort
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

        /// <summary>
        /// Retrieves the property RPC Agent Path (<see cref="RpcAgentPathProperty"/>) from the application configuration,
        /// or attempts to work it out by searching up from the current directory, looking for <code>applications/rpc-agent/rpc-agent.exe</code>.
        /// </summary>
        public string RpcAgentPath
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
                    Log.InfoFormat("No {0} property found in application configuration, searching for it relative to {1}",
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


        /// <summary>
        /// Starts the Thrift RPC Agent Process.
        /// </summary>
        /// <remarks>
        /// <para>This method launches a sub-process to start the Thrift RPC Agent.  Before calling this, you should add an event handler for
        /// the <see cref="ThriftRpcAgentStarted"/> event.  Until this event is invoked, it is not guaranteed that the process has started.</para>
        /// <para>If the agent fails to start properly then the <see cref="ThriftRpcAgentExited"/>.</para>
        /// </remarks>
        public void StartThriftRpcAgentProcess()
        {
            Log.InfoFormat("Attempting to launch Thrift RPC Agent {0}", RpcAgentPath);

            Process thriftRpcProcess = new Process
            {
                StartInfo = new ProcessStartInfo(RpcAgentPath, string.Join(" ",
                    "-host", RpcAgentServiceHost,
                    "-port", RpcAgentServicePort,
                    "-protocol", RpcAgentProtcol
                    ))
                {
                    RedirectStandardOutput = true,
                    RedirectStandardError = true,
                    UseShellExecute = false,
                    CreateNoWindow = true
                }
            };
            thriftRpcProcess.OutputDataReceived += ThriftRpcProcess_OutputDataReceived;
            thriftRpcProcess.ErrorDataReceived += ThriftRpcProcess_ErrorDataReceived;
            thriftRpcProcess.Exited += ThriftRpcProcess_Exited;
            thriftRpcProcess.Start();
            thriftRpcProcess.BeginOutputReadLine();
            thriftRpcProcess.BeginErrorReadLine();
            _thriftRpcProcess = thriftRpcProcess;
        }

        private void ThriftRpcProcess_Exited(object sender, EventArgs e)
        {
            Process proc = (Process)sender;
            if (proc.ExitCode == 0)
            {
                ThriftRpcLog.Info("Thrift RPC Agent has terminated with exit code 0");
            }
            else
            {
                ThriftRpcLog.Fatal($"Thrift RPC Agent has exited abnormally with exit code {proc.ExitCode}");
            }
            ThriftRpcAgentExited?.Invoke(this, EventArgs.Empty);
        }


        /// <summary>
        /// Terminates the Thrift RPC Agent.
        /// </summary>
        /// <remarks>
        /// Attempts to send a Ctrl-C signal to the agent.  If this cannot be done, or fails, then the process is killed via
        /// <see cref="Process.Kill"/>.</remarks>
        public void StopThriftRpcAgentProcess()
        {
            if (_thriftRpcProcess == null) throw new InvalidOperationException("Cannot stop Thrift RPC Agent process unless it's started");
            Log.InfoFormat("Attempting to stop Thrift RPC Agent {0}", _thriftRpcProcess.Id);

            if (_thriftRpcProcess.HasExited)
            {
                Log.Warn("Ignoring call to stop Thrift RPC agent as the process has already exited");
            }
            else
            {
                if (!SentCtrlCToProcess(_thriftRpcProcess))
                {
                    Log.Info("Unable to gracefully stop Thift RPC Agent, issuing kill instead");
                    _thriftRpcProcess.Kill();
                }
            }
        }

        private void ThriftRpcProcess_ErrorDataReceived(object sender, DataReceivedEventArgs e)
        {
            Process process = (Process)sender;
            ThriftRpcLog.Error($"RpcAgent({process.Id}): {e.Data}");
            RpcAgentErrors?.Invoke(process, e.Data);
        }

        private void ThriftRpcProcess_OutputDataReceived(object sender, DataReceivedEventArgs e)
        {
            Process process = (Process)sender;
            ThriftRpcLog.Info($"RpcAgent({process.Id}): {e.Data}");
            RpcAgentMessages?.Invoke(process, e.Data);
            const string startString = "Starting the rpc server on";
            if (e.Data!=null && e.Data.Contains(startString))
            {
                Log.InfoFormat("Found magic string \"{0}\" indicating that RPC Agent {1} has started successfully", startString, process.Id);
                ThriftRpcAgentStarted?.Invoke(this, EventArgs.Empty);
            }
        }

        #region Managing Process shutdown gracefully (Ctrl+C instead of kill)

        internal const int CtrlCEvent = 0;
        [DllImport("kernel32.dll")]
        internal static extern bool GenerateConsoleCtrlEvent(uint dwCtrlEvent, uint dwProcessGroupId);
        [DllImport("kernel32.dll", SetLastError = true)]
        internal static extern bool AttachConsole(uint dwProcessId);
        [DllImport("kernel32.dll", SetLastError = true, ExactSpelling = true)]
        internal static extern bool FreeConsole();
        [DllImport("kernel32.dll")]
        static extern bool SetConsoleCtrlHandler(ConsoleCtrlDelegate handlerRoutine, bool add);

        // Delegate type to be used as the Handler Routine for SCCH
        delegate Boolean ConsoleCtrlDelegate(uint ctrlType);

        private bool SentCtrlCToProcess(Process process)
        {
            FreeConsole();

            if (AttachConsole((uint)process.Id))
            {
                SetConsoleCtrlHandler(null, true);
                try
                {
                    if (!GenerateConsoleCtrlEvent(CtrlCEvent, 0))
                    {
                        Log.Warn("Unable to generate Ctrl event for process");
                        return false;
                    }
                    Log.Info("Send Ctrl-C to Thrid RPC agent, now waiting for exit");
                    process.WaitForExit();
                    Log.Info("Thrift RPC agent has exited cleanly.");
                }
                finally
                {
                    Log.Debug("Restoring console handler for host process (i.e. us)");
                    FreeConsole();
                    SetConsoleCtrlHandler(null, false);
                }
                return true;
            }
            else
            {
                Log.Warn($"Unable to attach to console on process {process.Id}");
                return false;
            }
        }

        #endregion

    }
}
