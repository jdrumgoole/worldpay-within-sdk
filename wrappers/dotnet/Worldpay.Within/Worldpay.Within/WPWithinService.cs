using System;
using Common.Logging;
using Thrift.Protocol;
using Thrift.Transport;
using Worldpay.Innovation.WPWithin.ThriftAdapters;
using ThriftWPWithinService = Worldpay.Innovation.WPWithin.Rpc.WPWithin;


namespace Worldpay.Innovation.WPWithin
{
    /**
     * Service wrapper that hides all references to underlying implementation (i.e. Thrift).
     */

    public class WPWithinService : IDisposable
    {
        /*
   void addService(1: wptypes.Service svc) throws (1: wptypes.Error err),
   void removeService(1: wptypes.Service svc) throws (1: wptypes.Error err),
   void initHCE(1: wptypes.HCECard hceCard) throws (1: wptypes.Error err),
   void initHTE(1: string merchantClientKey, 2: string merchantServiceKey) throws (1: wptypes.Error err),
   void initConsumer(1: string scheme, 2: string hostname, 3: i32 port, 4: string urlPrefix, 5: string serviceId) throws (1: wptypes.Error err),
   void initProducer() throws (1: wptypes.Error err),
   wptypes.Device getDevice(),
   void startServiceBroadcast(1: i32 timeoutMillis) throws (1: wptypes.Error err),
   void stopServiceBroadcast() throws (1: wptypes.Error err),
   set<wptypes.ServiceMessage> serviceDiscovery(1: i32 timeoutMillis) throws (1: wptypes.Error err),
   set<wptypes.ServiceDetails> requestServices() throws (1: wptypes.Error err),
   set<wptypes.Price> getServicePrices(1: i32 serviceId) throws (1: wptypes.Error err),
   wptypes.TotalPriceResponse selectService(1: i32 serviceId, 2: i32 numberOfUnits, 3: i32 priceId) throws (1: wptypes.Error err),
   wptypes.PaymentResponse makePayment(1: wptypes.TotalPriceResponse request) throws (1: wptypes.Error err),
*/

        public void StartServiceBroadcast(int timeoutMillis)
        {
            _client.startServiceBroadcast(timeoutMillis);
        }

        public void SetupDevice(string deviceName, string deviceDescription)
        {
            _client.setup(deviceName, deviceDescription);
        }

        public void InitProducer(string clientKey, string serviceKey)
        {
            _client.initHTE("cl_key", "srv_key");
            _client.initProducer();
        }

        public void InitConsumer(HceCard card)
        {
            _client.initHCE(new HceCardAdapter().Create(card));
            // TODO Identify where we get the parameter for initConsumer from.  Can't be hard-coded, must be from a discovered service
//          _client.initConsumer("");
        }
        
        private static readonly ILog Log = LogManager.GetLogger<WPWithinService>();
        private ThriftWPWithinService.Client _client;
        private bool _isDisposed;
        private TTransport _transport;


        public WPWithinService(string host, int port)
        {
            Init(host, port);
        }

        public void Dispose()
        {
            Dispose(true);
        }

        private void Init(string host, int port)
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
                Log.Warn("Error closing transport.", e);
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