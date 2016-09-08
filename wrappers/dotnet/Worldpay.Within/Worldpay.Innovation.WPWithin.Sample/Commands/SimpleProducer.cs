using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Common.Logging;

namespace Worldpay.Innovation.WPWithin.Sample.Commands
{
    internal class SimpleProducer
    {
        private static readonly ILog Log = LogManager.GetLogger<SimpleProducer>();
        private TextWriter _error;
        private readonly TextWriter _output;
        private Task _task;
        private WPWithinService _wpWithinService;

        public SimpleProducer(TextWriter output, TextWriter error)
        {
            _output = output;
            _error = error;
        }

        public CommandResult Start()
        {
            _output.WriteLine("WorldpayWithin Sample Producer...");

            WPWithinService wpWithinService = new WPWithinService("127.0.0.1", 9091);

            wpWithinService.SetupDevice("Producer Example", "Example WorldpayWithin producer");

            Service svc = new Service
            {
                Name = "Car charger",
                Description = "Can charge your hybrid / electric car",
                Id = 1,
                Prices = new Dictionary<int, Price>
                {
                    {
                        1, new Price
                        {
                            Id = 1,
                            Description = "Kilowatt-hour",
                            UnitDescription = "One kilowatt-hour",
                            UnitId = 1,
                            PricePerUnit = new PricePerUnit
                            {
                                Amount = 25,
                                CurrencyCode = "GBP"
                            }
                        }
                    }
                }
            };

            wpWithinService.AddService(svc);

            wpWithinService.InitProducer("T_C_03eaa1d3-4642-4079-b030-b543ee04b5af", "T_S_f50ecb46-ca82-44a7-9c40-421818af5996");

            Log.Info("Starting service broadcast");
            _wpWithinService = wpWithinService;
            _task = Task.Run(() => wpWithinService.StartServiceBroadcast(20000));
            return CommandResult.Success;
        }

        public void Stop()
        {
            _output.WriteLine("Stopping service broadcast");
            _wpWithinService.StopServiceBroadcast();
            _output.WriteLine("Waiting for producer task to complete");
            _task.Wait();
            _output.WriteLine("Producer task terminated.");
        }
    }
}