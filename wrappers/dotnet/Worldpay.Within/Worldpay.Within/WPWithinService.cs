using System;
using System.Collections.Generic;
using Common.Logging;
using Thrift.Protocol;
using Thrift.Transport;
using Worldpay.Innovation.WPWithin.EventListener;
using Worldpay.Innovation.WPWithin.ThriftAdapters;
using ThriftWPWithinService = Worldpay.Innovation.WPWithin.Rpc.WPWithin;


namespace Worldpay.Innovation.WPWithin
{
    /// <summary>
    ///     Service wrapper that hides all references to underlying implementation (i.e. Thrift).
    /// </summary>
    public class WPWithinService : IDisposable
    {
        public delegate void BeginServiceDeliveryHandler(
            int serviceId, ServiceDeliveryToken serviceDeliveryToken, int unitsToSupply);

        public delegate void EndServiceDeliveryHandler(
            int serviceId, ServiceDeliveryToken serviceDeliveryToken, int unitsReceived);

        private static readonly ILog Log = LogManager.GetLogger<WPWithinService>();

        private ThriftWPWithinService.Client _client;
        private bool _isDisposed;

        private readonly CallbackServerManager _listener;
        private TTransport _transport;

        public WPWithinService(string host, int port) : this(host, port, 0)
        {
        }

        public WPWithinService(string host, int port, int callbackPort)
        {
            InitClient(host, port);
            _listener = new CallbackServerManager(callbackPort);
//            _listener.Start();
        }

        public void Dispose()
        {
            Dispose(true);
        }

        public event BeginServiceDeliveryHandler OnBeginServiceDelivery
        {
            add { _listener.BeginServiceDelivery += value; }
            remove { _listener.BeginServiceDelivery -= value; }
        }

        public event EndServiceDeliveryHandler OnEndServiceDelivery
        {
            add { _listener.EndServiceDelivery += value; }
            remove { _listener.EndServiceDelivery -= value; }
        }

        public void AddService(Service service)
        {
            _client.addService(ServiceAdapter.Create(service));
        }

        public void RemoveService(Service service)
        {
            _client.removeService(ServiceAdapter.Create(service));
        }

        public void InitConsumer(string scheme, string hostname, int port, string urlPrefix, string serviceId,
            HceCard hceCard)
        {
            _client.initConsumer(scheme, hostname, port, urlPrefix, serviceId, HceCardAdapter.Create(hceCard));
        }

        public void InitProducer(string merchantClientKey, string merchantServiceKey)
        {
            _client.initProducer(merchantClientKey, merchantServiceKey);
        }

        public Device GetDevice()
        {
            return DeviceAdapter.Create(_client.getDevice());
        }

        public void StopServiceBroadcast()
        {
            _client.stopServiceBroadcast();
        }

        public IEnumerable<ServiceMessage> DeviceDiscovery(int timeoutMillis)
        {
            return ServiceMessageAdapter.Create(_client.deviceDiscovery(timeoutMillis));
        }

        public IEnumerable<ServiceDetails> RequestServices()
        {
            return ServiceDetailsAdapter.Create(_client.requestServices());
        }

        public IEnumerable<Price> GetServicePrices(int serviceId)
        {
            return PriceAdapter.Create(_client.getServicePrices(serviceId));
        }

        public TotalPriceResponse SelectService(int serviceId, int numberOfUnits, int priceId)
        {
            return TotalPriceResponseAdapter.Create(_client.selectService(serviceId, numberOfUnits, priceId));
        }

        public PaymentResponse MakePayment(TotalPriceResponse request)
        {
            return PaymentResponseAdapter.Create(_client.makePayment(TotalPriceResponseAdapter.Create(request)));
        }

        public void BeginServiceDelivery(int serviceId, ServiceDeliveryToken serviceDeliveryToken, int unitsToSupply)
        {
            _client.beginServiceDelivery(serviceId, ServiceDeliveryTokenAdapter.Create(serviceDeliveryToken),
                unitsToSupply);
        }

        public void EndServiceDelivery(int serviceId, ServiceDeliveryToken serviceDeliveryToken, int unitsReceived)
        {
            _client.endServiceDelivery(serviceId, ServiceDeliveryTokenAdapter.Create(serviceDeliveryToken),
                unitsReceived);
        }

        public void StartServiceBroadcast(int timeoutMillis)
        {
            _client.startServiceBroadcast(timeoutMillis);
        }

        public void SetupDevice(string deviceName, string deviceDescription)
        {
            _client.setup(deviceName, deviceDescription);
        }

        private void InitClient(string host, int port)
        {
            Log.InfoFormat("Opening TSocket to {0}:{1}", host, port);
            TTransport transport = new TSocket(host, port);
            transport.Open();

            TProtocol protocol = new TBinaryProtocol(transport);
            ThriftWPWithinService.Client client = new ThriftWPWithinService.Client(protocol);

            _transport = transport;
            _client = client;
            Log.InfoFormat("Client connected to Thrift endpoint at {0}:{1}", host, port);
        }

        protected virtual void Dispose(bool disposing)
        {
            if (!_isDisposed)
            {
                if (disposing)
                {
                    GC.SuppressFinalize(this);
                }
            }
            try
            {
                _transport.Close();
            }
            catch (Exception e)
            {
                Log.Warn("Error closing connection to RPC Agent", e);
            }

            try
            {
                _listener.Stop();
            }
            catch (Exception e)
            {
                Log.Warn("Error stopping callback listener", e);
            }
            //Dispose of resources here
            _isDisposed = true;
        }

        ~WPWithinService()
        {
            Dispose(false);
        }
    }
}