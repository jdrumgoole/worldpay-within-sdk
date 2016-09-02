var util = require('util');

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

function fnStartRPC(port, callback) {

  var rpc = require('./rpc');

  rpc.startRPC(port, callback);
}

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

var fnBeginServiceDelivery = function(serviceId, serviceDeliveryToken, unitsToSupply, callback) {

  this.thriftClient.beginServiceDelivery(serviceId, serviceDeliveryToken, unitsToSupply, function(err, result) {

    callback(err, result);
  });
};

var fnEndServiceDelivery = function(serviceId, serviceDeliveryToken, unitsReceived, callback) {

  this.thriftClient.endServiceDelivery(serviceId, serviceDeliveryToken, unitsReceived, function(err, result) {

    callback(err, result);
  });
};

// Factory setup WPWithinClient
// Should return an instance of WPWithin
function createClient(host, port, startRPCAgent, callback) {

  createClient(host, port, startRPCAgent, callback, null, 0);
}

// Factory setup WPWithinClient
// Should return an instance of WPWithin
function createClient(host, port, startRPCAgent, callback, eventListener, callbackPort) {

  try {

    var thrift = require('thrift');
    var WPWithinLib = require('./wpwithin-thrift/WPWithin');
    var evServer = require('./eventlistener/eventserver');

    var doCreate = function() {

      transport = thrift.TBufferedTransport;
      protocol = thrift.TBinaryProtocol;

      var connection = thrift.createConnection(host, port);

      connection.on('error', function(err) {

        callback(err, null);
      });

      tc = thrift.createClient(WPWithinLib, connection);

      callback(null, new WPWithin(tc));
    }

    if(eventListener != null) {

      new evServer.EventServer().start(eventListener, callbackPort);
    }

    if(callbackPort > 0) {

      launchRPCAgent(port, function(error, stdout, stderr){

          if(error == null) {

            return doCreate();
          } else {

            var strErr = util.format("%s \n %s", error, stderr)

            callback(strErr, null);
          }
      });

    } else {

      doCreate()
    }
  } catch (err) {

    console.log("Caught error: %s", err)

    callback(err, null);
  }
};

function launchRPCAgent(port, callback) {

  var launcher = require('./launcher');

  var config = {
  	"windows": {
  		"x64": null,
  		"ia32": null,
  		"arm": null
  	},
  	"darwin": {
  		"x64": util.format("/Users/conor/Repositories/GoLang/src/github.com/wptechinnovation/worldpay-within-sdk/applications/rpc-agent/rpc-agent -port %d -logfile wpwithin.log", port),
  		"ia32": util.format("/Users/conor/Repositories/GoLang/src/github.com/wptechinnovation/worldpay-within-sdk/applications/rpc-agent/rpc-agent -port %d -logfile wpwithin.log", port),
  		"arm": null
  	},
  	"linux": {
  		"x64": null,
  		"ia32": null,
  		"arm": null
  	}
  };

  var launchCallback = function(error, stdout, stderr) {

    console.log("-------------Launcher Process Event-------------");
    console.log("error: " + error);
    console.log("stdout: " + stdout);
    console.log("stderr: " + stderr);
    console.log("------------------------------------------------");

    callback(error, stdout, stderr);
  };

  launcher.startProcess(config, launchCallback);

  callback(null, null, null);
};
