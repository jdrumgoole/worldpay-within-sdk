using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Common.Logging;

namespace Worldpay.Innovation.WPWithin.Sample.Commands
{
    internal class SimpleProducer
    {
        private static readonly ILog Log = LogManager.GetLogger<SimpleProducer>();
        private readonly TextWriter _error;
        private readonly WPWithinService _service;
        private readonly TextWriter _output;
        private Task _task;

        public SimpleProducer(TextWriter output, TextWriter error, WPWithinService service)
        {
            _output = output;
            _error = error;
            _service = service;
        }

        public CommandResult Start()
        {
            _output.WriteLine("WorldpayWithin Sample Producer...");

            _service.SetupDevice("Producer Example", "Example WorldpayWithin producer");

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

            _service.AddService(svc);

            _service.InitProducer("T_C_03eaa1d3-4642-4079-b030-b543ee04b5af", "T_S_f50ecb46-ca82-44a7-9c40-421818af5996");

            Log.Info("Starting service broadcast");
            _task = Task.Run(() => _service.StartServiceBroadcast(20000));
            return CommandResult.Success;
        }

        public void Stop()
        {
            _output.WriteLine("Stopping service broadcast");
            _service.StopServiceBroadcast();
            _output.WriteLine("Waiting for producer task to complete");
            _task.Wait();
            _output.WriteLine("Producer task terminated.");
        }
    }
}