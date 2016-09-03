using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Worldpay.Innovation.WPWithin.AgentManager;

namespace Worldpay.Innovation.WPWithin.Sample.Commands
{
    internal class CommandMenu
    {
        private readonly TextWriter _error;

        private readonly List<Command> _menuItems;
        private readonly TextWriter _output;
        private readonly TextReader _reader;
        private RpcAgentManager _rpcManager;

        public CommandMenu()
        {
            _menuItems = new List<Command>(new[]
            {
                new Command("Exit", "Exits the application.", (a) => CommandResult.Exit),
                new Command("StartRPCClient", "Starts the Thrift RPC Client", StartRpcClient),
                new Command("StopRPCClient", "Stops the Thrift RPC Client", StopRpcClient),
            });

            // TODO Parameterise these so output can be written to a specific file
            _output = Console.Out;
            _error = Console.Error;
            _reader = Console.In;
        }

        private CommandResult StopRpcClient(string[] arg)
        {
            if (_rpcManager == null)
            {
                _error.WriteLine("Thift RPC Agent not active.  Start it before trying to stop it.");
                return CommandResult.Failure;
            }
            _rpcManager.StopThriftRpcAgentProcess();
            _rpcManager = null;
            return CommandResult.Success;
        }

        private CommandResult StartRpcClient(string[] arg)
        {
            if (_rpcManager != null)
            {
                _error.WriteLine("Thrift RPC Agent already active.  Stop it before trying to start a new one");
                return CommandResult.Failure;
            }
            _rpcManager = new RpcAgentManager();
            _rpcManager.StartThriftRpcAgentProcess();
            return CommandResult.Success;
        }

        internal CommandResult ReadEvalPrint()
        {
            _output.WriteLine("Sample Application.");
            foreach (Command item in _menuItems)
            {
                _output.WriteLine("{0}: {1}", item.Name, item.Description);
            }

            // Read
            _output.Write("\nCommand: ");
            string readLine = _reader.ReadLine();
            if (readLine == null)
            {
                return CommandResult.NoOp;
            }

            string[] args = readLine.Split();

            // If no arguments, then don't error, just return success;
            if (args.Length == 0 || string.IsNullOrEmpty(args[0]))
            {
                return CommandResult.NoOp;
            }

            Command selectedItem = _menuItems.FirstOrDefault(m => m.Name.Equals(args[0]));
            if (selectedItem != null)
            {
                return selectedItem.Function(args);
            }

            _output.WriteLine("No such option.");
            return CommandResult.NoSuchCommand;
        }
    }
}