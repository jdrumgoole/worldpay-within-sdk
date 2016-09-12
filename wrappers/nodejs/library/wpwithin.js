var util = require('util');
var thrift = require('thrift');
var wpwithinThrift = require('./wpwithin-thrift/WPWithin');
var eventServer = require('./eventlistener/eventserver');

module.exports = {
  createClient: createClient
};

function WPWithin(thriftClient) {

  this.converter = require('./types/converter');

  this.thriftClient = thriftClient;

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
function createClient(host, port, startRPC, callback, eventListener, callbackPort) {

  try {

    // First, we validate the callback parameters. There are two and both need to be set
    // to setup the callback server.
    if(eventListener != null) {

      // If eventListener is set, then need to validate the port
      if(callbackPort <= 0 || callbackPort > 65535) {

        callback(util.format("callbackPort (%d) is invalid should be > 0 and <= 65535", callbackPort), null)

        return
      }

      new eventServer.EventServer().start(eventListener, callbackPort);

    } else {

      // So the event listener is not set, meaning the developer doesn't want any feedback of events
      // in this case there is no need start the callback server. We can do this by setting the port to 0
      callbackPort = 0
    }

    if(startRPC) {

      launchRPCAgent(port, callbackPort, function(error, stdout, stderr){

          if(error == null) {

            createThriftClient(host, port, callback)
          } else {

            var strErr = util.format("%s \n %s", error, stderr)

            callback(strErr, null);
          }
      });

    } else {

      createThriftClient(host, port, callback)
    }
  } catch (err) {

    console.log("Caught error: %s", err)

    callback(err, null);
  }
};

function launchRPCAgent(port, callbackPort, callback) {

  var launcher = require('./launcher');

  var flagLogFile = "wpwithin.log"
  var flagLogLevels = "debug,error,info,warn,fatal"
  var flagCallbackPort = callbackPort > 0 ? "-callbackport="+callbackPort : ""
  var binBase = process.env.WPWBIN == "" ? "./rpc-agent-bin" : process.env.WPWBIN

  var config = {
  	"windows": {
  		"x64": util.format("%s/rpc-agent-win-64 -port=%d -logfile=%s -loglevel=%s %s", binBase, port, flagLogFile, flagLogLevels, flagCallbackPort),
  		"ia32": util.format("%s/rpc-agent-win-32 -port=%d -logfile=%s -loglevel=%s %s", binBase, port, flagLogFile, flagLogLevels, flagCallbackPort),
  		"arm": null
  	},
  	"darwin": {
  		"x64": util.format("%s/rpc-agent-mac-64 -port %d -logfile %s -loglevel %s %s", binBase, port, flagLogFile, flagLogLevels, flagCallbackPort),
  		"ia32": util.format("%s/rpc-agent-mac-32 -port %d -logfile %s -loglevel %s %s", binBase, port, flagLogFile, flagLogLevels, flagCallbackPort),
  		"arm": null
  	},
  	"linux": {
  		"x64": util.format("%s/rpc-agent-linux-64 -port %d -logfile %s -loglevel %s %s",binBase, port, flagLogFile, flagLogFile, flagCallbackPort),
  		"ia32": util.format("%s/rpc-agent-linux-32 -port %d -logfile %s -loglevel %s %s",binBase, port, flagLogFile, flagLogFile, flagCallbackPort),
  		"arm": util.format("%s/rpc-agent-linux-arm -port %d -logfile %s -loglevel %s %s",binBase, port, flagLogFile, flagLogFile, flagCallbackPort),
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

  var sleep = require('sleep');
  sleep.usleep(750);
  callback(null, null, null);
};

// Create a WPWithin Thrift client
function createThriftClient(host, port, callback) {

  transport = thrift.TBufferedTransport;
  protocol = thrift.TBinaryProtocol;

  var connection = thrift.createConnection(host, port);

  connection.on('error', function(err) {

    callback(err, null);
  });

  client = thrift.createClient(wpwithinThrift, connection);

  callback(null, new WPWithin(client));
}
