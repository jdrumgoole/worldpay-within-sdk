module.exports = {
  createClient: createClient
};

function WPWithin(thriftClient) {

  this.converter = require('./types/converter');

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

  var convSvc = this.converter.toThrift().service(service);

  this.thriftClient.addService(convSvc, function(err, result) {

    callback(err, result);
  });
};

var fnRemoveService = function(service, callback) {

  var convSvc = this.converter.toThrift().service(service);

  this.thriftClient.removeService(convSvc, function(err, result) {

    callback(err, result);
  });
};

var fnInitConsumer = function(scheme, hostname, port, urlPrefix, serverId, hceCard, callback) {

  tHCECard = this.converter.toThrift().hceCard(hceCard);

  this.thriftClient.initConsumer(scheme, hostname, port, urlPrefix, serverId, tHCECard, function(err, result) {

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

  this.thriftClient.getServicePrices(serviceId, function(err, result) {

    callback(err, result);
  });
};

var fnSelectService = function(serviceId, numberOfUnits, priceId, callback) {

  this.thriftClient.selectService(serviceId, numberOfUnits, priceId, function(err, result) {

    callback(err, result);
  });
};

var fnMakePayment = function(request, callback) {

  var convRequest = this.converter.toThrift().totalPriceResponse(request);

  this.thriftClient.makePayment(convRequest, function(err, result) {

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

// Factory setup WPWithinClient
// Should return an instance of WPWithin
function createClient(host, port, callback) {

  try {

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

function fnStartRPC(port, callback) {

  var rpc = require('./rpc');

  rpc.startRPC(port, callback);
}
