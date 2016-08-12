module.exports = {
  createClient: createClient
};

function WPWithin(thriftClient) {

  this.thriftClient = thriftClient;

  this.startRPC = fnStartRPC;
  this.setup = fnSetup;
  this.addService = fnAddService;
  this.removeService = fnRemoveService;
  this.initConsumer = fnInitConsumer;
  this.initProducer = fnInitProducer;
  this.getDevice = fnGetDevice;
  this.startServiceBroadcast = fnStartServiceBroadcast;
  this.stopServiceBroadcast = fnStopServiceBroadcast;
  this.deviceDiscovery = fnDeviceDiscovery;
  this.requestServices = fnRequestServices;
  this.getServicePrices = fnGetServicePrices;
  this.selectService = fnSelectService;
  this.makePayment = fnMakePayment;
  this.beginServiceDelivery = fnBeginServiceDelivery;
  this.endServiceDelivery = fnEndServiceDelivery;

};

var fnSetup = function(name, description, callback) {

  this.thriftClient.setup(name, description, function(err, response) {

    callback(err, response);
  });
};

var fnAddService = function(service, callback) {

  this.thriftClient.addService(service, function(err, result) {

    callback(err, result);
  });
};

var fnRemoveService = function(service, callback) {

  this.thriftClient.removeService(service, function(err, result) {

    callback(err, result);
  });
};

var fnInitConsumer = function(scheme, hostname, port, urlPrefix, serverId, hceCard, callback) {

  this.thriftClient.initConsumer(scheme, hostname, port, urlPrefix, serverId, hceCard, function(err, result) {

    callback(err, result);
  });
};

var fnInitProducer = function(merchantClientKey, merchantServiceKey, callback) {

  this.thriftClient.initProducer(merchantClientKey, merchantServiceKey, function(err, result) {

    callback(err, result);
  });
};

var fnGetDevice = function(callback) {

  this.thriftClient.getDevice(function(err, result) {

    callback(err, result);
  });
};

var fnStartServiceBroadcast = function(timeoutMillis, callback) {

  this.thriftClient.startServiceBroadcast(timeoutMillis, function(err, result) {

    callback(err, result);
  });
};

var fnStopServiceBroadcast = function(callback) {

  this.thriftClient.stopServiceBroadcast(function(err, result) {

    callback(err, result);
  });
};

var fnDeviceDiscovery = function(timeoutMillis, callback) {

  this.thriftClient.deviceDiscovery(timeoutMillis, function(err, response) {

    callback(err, response);
  });
};

var fnRequestServices = function(callback) {

  this.thriftClient.requestServices(function(err, result) {

    callback(err, result);
  });
};

var fnGetServicePrices = function(serviceId, callback) {

  this.thriftClient.getServicePrices(function(err, result) {

    callback(err, result);
  });
};

var fnSelectService = function(serviceId, numberOfUnits, priceId, callback) {

  this.thriftClient.selectService(serviceId, numberOfUnits, priceId, function(err, result) {

    callback(err, result);
  });
};

var fnMakePayment = function(request, callback) {

  this.thriftClient.makePayment(request, function(err, result) {

    callback(err, result);
  });
};

var fnBeginServiceDelivery = function(clientId, serviceDeliveryToken, unitsToSupply, callback) {

  this.thriftClient.beginServiceDelivery(clientId, serviceDeliveryToken, unitsToSupply, function(err, result) {

    callback(err, result);
  });
};

var fnEndServiceDelivery = function(clientId, serviceDeliveryToken, unitsReceived, callback) {

  this.thriftClient.endServiceDelivery(clientId, serviceDeliveryToken, unitsReceived, function(err, result) {

    callback(err, result);
  });
};

// void setup(1: string name, 2: string description) throws (1: wptypes.Error err),
// void addService(1: wptypes.Service svc) throws (1: wptypes.Error err),
// void removeService(1: wptypes.Service svc) throws (1: wptypes.Error err),
// void initConsumer(1: string scheme, 2: string hostname, 3: i32 port, 4: string urlPrefix, 5: string serverId, 6: wptypes.HCECard hceCard) throws (1: wptypes.Error err),
// void initProducer(1: string merchantClientKey, 2: string merchantServiceKey) throws (1: wptypes.Error err),
// wptypes.Device getDevice(),
// void startServiceBroadcast(1: i32 timeoutMillis) throws (1: wptypes.Error err),
// void stopServiceBroadcast() throws (1: wptypes.Error err),
// set<wptypes.ServiceMessage> deviceDiscovery(1: i32 timeoutMillis) throws (1: wptypes.Error err),
// set<wptypes.ServiceDetails> requestServices() throws (1: wptypes.Error err),
// set<wptypes.Price> getServicePrices(1: i32 serviceId) throws (1: wptypes.Error err),
// wptypes.TotalPriceResponse selectService(1: i32 serviceId, 2: i32 numberOfUnits, 3: i32 priceId) throws (1: wptypes.Error err),
// wptypes.PaymentResponse makePayment(1: wptypes.TotalPriceResponse request) throws (1: wptypes.Error err),
// void beginServiceDelivery(1: string clientId, 2: wptypes.ServiceDeliveryToken serviceDeliveryToken, 3: i32 unitsToSupply) throws (1: wptypes.Error err),
// void endServiceDelivery(1: string clientId, 2: wptypes.ServiceDeliveryToken serviceDeliveryToken, 3: i32 unitsReceived) throws (1: wptypes.Error err),

// Factory setup WPWithinClient
// Should return an instance of WPWithin
function createClient(host, port, callback) {

  try {

    fnStartRPC(port, function(err, stdout, stderr) {

      console.log("Err: ", err);
      console.log("STDOUT: ", stdout);
      console.log("STDERR: ", stderr);
    });

    var thrift = require('thrift');
    var WPWithinLib = require('./wpwithin-thrift/WPWithin');

    transport = thrift.TBufferedTransport;
    protocol = thrift.TBinaryProtocol;

    var connection = thrift.createConnection(host, port);

    connection.on('error', function(err) {

      callback(err, null);
    });

    tc = thrift.createClient(WPWithinLib, connection);

    return new WPWithin(tc);

  } catch (err) {

    console.log("Caught error: %s", err)

    callback(err, null);
  }
};

function fnStartRPC(port) {

  var rpc = require('./rpc');

  rpc.startRPC(port);
}
