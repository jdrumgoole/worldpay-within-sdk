using System;
using System.Linq;
using System.Collections.Generic;
using System.IO;
using Worldpay.Innovation.WPWithin.AgentManager;

namespace Worldpay.Innovation.WPWithin.Sample.Commands
{
    internal class CommandMenu
    {

        private List<Command> MenuItems;
        private TextWriter _output;
        private TextWriter _error;
        private TextReader _reader;
        private RpcAgentManager _rpcManager;

        public CommandMenu()
        {
            MenuItems = new List<Command>(new Command[] {
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
            if (this._rpcManager==null)
            {
                _error.WriteLine("Thift RPC Agent not active.  Start it before trying to stop it.");
                return CommandResult.Failure;
            }
            this._rpcManager.StopThriftRpcAgentProcess();
            this._rpcManager = null;
            return CommandResult.Success;
        }

        private CommandResult StartRpcClient(string[] arg)
        {
            if (this._rpcManager!=null)
            {
                _error.WriteLine("Thrift RPC Agent already active.  Stop it before trying to start a new one");
                return CommandResult.Failure;
            }
            this._rpcManager = new RpcAgentManager();
            this._rpcManager.StartThriftRpcAgentProcess();
            return CommandResult.Success;
        }

        internal CommandResult ProcessCommand()
        {

            _output.WriteLine("Sample Application.");
            foreach (Command item in MenuItems)
            {
                _output.WriteLine("{0}: {1}", item.Name, item.Description);
            }

            // Read
            _output.Write("\nCommand: ");
            string[] args = _reader.ReadLine().Split();

            // If no arguments, then don't error, just return success;
            if (args.Length == 0 || String.IsNullOrEmpty(args[0]))
            {
                return CommandResult.NoOp;
            }

            Command selectedItem = MenuItems.FirstOrDefault(m => m.Name.Equals(args[0]));
            if (selectedItem == null)
            {
                _output.WriteLine("No such option.");
                return CommandResult.NoSuchCommand;
            }
            else
            {
                return selectedItem.Function(args);
            }

        }
    }
}