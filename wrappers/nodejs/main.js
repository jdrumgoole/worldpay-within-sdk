var wpwithin = require('./wpwithin');
var types = require('./types/types');
var typesConverter = require('./types/converter');

client = wpwithin.createClient("localhost", 9090, null);

client.setup("conor-njs", "conors node.js client", function(err, response){

  console.log("setup.callback.err: %j", err);
  console.log("setup.callback.response: %j", response);

  discoverDevices();
});

function discoverDevices() {

  client.deviceDiscovery(10000, function(err, response) {

    for(var i = 0; i < response.length; i++) {

        console.log("Hostname: " + response[i].hostname);
    }

    console.log("deviceDiscovery.callback.err: %j", err);
    console.log("deviceDiscovery.callback.response: %j", response);
  });
};

function addService() {

  var service = new types.Service();
  console.log("%j", service);

  service.id = 1;
  service.name = "svcNAME";
  service.description = "svcDESCRIPTION";
  service.prices = null;

  result = new wpthrift_types.Service();
  result.id = service.id;
  result.name = service.name;
  result.description = service.description;
  result.prices = null;

  var ft = typesConverter.fromThrift();
  resultConv = ft.service(result);

  console.log("result: %j", resultConv);

  var tt = typesConverter.toThrift();

  svcConverted = tt.service(service);

  console.log("converted: %j", svcConverted);


  client.addService(svcConverted, function(err, response) {

      console.log("addService().callback");
      console.log("err: %j", err)
      console.log("response: %j", response)
  });
}

function getDevice() {

  client.getDevice(function(err, response) {

    console.log("getDevice().callback");
    console.log("err: %j", err)
    console.log("response: %j", response)
  });
}
