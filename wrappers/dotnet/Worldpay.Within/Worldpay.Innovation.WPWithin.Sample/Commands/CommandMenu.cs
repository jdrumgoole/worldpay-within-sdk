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
        private SimpleProducer _simpleProducer;

        public CommandMenu()
        {
            _menuItems = new List<Command>(new[]
            {
                new Command("Exit", "Exits the application.", (a) =>
                {
                    _output.WriteLine("Exiting...");
                    return CommandResult.Exit;
                }),
                new Command("StartRPCClient", "Starts the Thrift RPC Client", StartRpcClient),
                new Command("StopRPCClient", "Stops the Thrift RPC Client", StopRpcClient),
                new Command("StartSimpleProducer", "Starts a simple producer", StartSimpleProducer),
                new Command("StopSimpleProducer", "Starts a simple producer", StopSimpleProducer),
                new Command("ConsumePurchase", "Consumes a service (first price of first service found)", ConsumePurchase),
            });

            // TODO Parameterise these so output can be written to a specific file
            _output = Console.Out;
            _error = Console.Error;
            _reader = Console.In;
        }

        private CommandResult ConsumePurchase(string[] arg)
        {
            SimpleConsumer consumer = new SimpleConsumer(_output, _error);
            consumer.MakePurchase(9091);
            return CommandResult.Success;
        }

        private CommandResult StartSimpleProducer(string[] arg)
        {
            if (_simpleProducer != null)
            {
                _output.WriteLine("Simple producer already started, stop it before trying to start it again.");
                return CommandResult.NonCriticalError;
            }
            _simpleProducer = new SimpleProducer(_output, _error);
            _simpleProducer.Start();
            return CommandResult.Success;
        }

        private CommandResult StopSimpleProducer(string[] arg)
        {
            if (_simpleProducer == null)
            {
                _output.WriteLine("Cannot stop Simple producer as it is not started.");
                return CommandResult.NonCriticalError;
            }
            _simpleProducer.Stop();
            _simpleProducer = null;
            return CommandResult.Success;
        }

        private CommandResult StopRpcClient(string[] arg)
        {
            if (_rpcManager == null)
            {
                _error.WriteLine("Thift RPC Agent not active.  Start it before trying to stop it.");
                return CommandResult.NonCriticalError;
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
                return CommandResult.NonCriticalError;
            }
            _rpcManager = new RpcAgentManager();
            _rpcManager.StartThriftRpcAgentProcess();
            return CommandResult.Success;
        }

        internal CommandResult ReadEvalPrint(string[] args)
        {
            // Only show the menu if there isn't already a command line to deal with
            if (args != null)
            {
                _output.WriteLine("\nSample Application.");
                int count = 0;
                foreach (Command item in _menuItems)
                {
                    _output.WriteLine("{0}. {1}: {2}", count, item.Name, item.Description);
                    count++;
                }

                // Read
                _output.Write("\nCommand: ");
                string readLine = _reader.ReadLine();
                if (readLine == null)
                {
                    return CommandResult.NoOp;
                }

                args = readLine.Split();
            }

            // If no arguments, then don't error, just return success;
            if (args == null || args.Length == 0 || string.IsNullOrEmpty(args[0]))
            {
                return CommandResult.NoOp;
            }

            int optionNumber;
            Command selectedItem = int.TryParse(args[0], out optionNumber) ? _menuItems[optionNumber] : _menuItems.FirstOrDefault(m => m.Name.Equals(args[0]));

            if (selectedItem != null)
            {
                try
                {
                    return selectedItem.Function(args);
                }
                catch (WPWithinException wpwe)
                {
                    _error.WriteLine(wpwe);
                    return CommandResult.NonCriticalError;
                }
                catch (Exception wpwe)
                {
                    _error.WriteLine(wpwe);
                    return CommandResult.CriticalError;
                }
            }

            _output.WriteLine($"Invalid command: \"{args[0]}\"");
            return CommandResult.NoSuchCommand;
        }
    }
}