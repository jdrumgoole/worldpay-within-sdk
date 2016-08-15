var wpwithin = require('../../library/wpwithin');
var types = require('../../library/types/types');
var typesConverter = require('../../library/types/converter');

client = wpwithin.createClient("127.0.0.1", 9090, function(err, response){

  console.log("createClient.callback")
  console.log("createClient.callback.err: " + err)
  console.log("createClient.callback.response: %j", response);
});

client.setup("NodeJS-Device", "Sample NodeJS producer device", function(err, response){

  console.log("setup.callback.err: " + err);
  console.log("setup.callback.response: %j", response);

  if(err == null) {

    addService();
  }
});

function addService() {

  var service = new types.Service();

  service.id = 1;
  service.name = "RoboWash";
  service.description = "Car washed by robot";

  var rwPrice = new types.Price();
  rwPrice.id = 1;
  rwPrice.description = "Car wash";
  rwPrice.unitId = 1;
  rwPrice.unitDescription = "Single wash";
  rwPrice.pricePerUnit = new types.PricePerUnit();
  rwPrice.pricePerUnit.amount = 650;
  rwPrice.pricePerUnit.currencyCode = "GBP";
  service.prices = rwPrice;

  client.addService(service, function(err, response) {

      console.log("addService.callback");
      console.log("err: " + err)
      console.log("response: %j", response)

      if(err == null) {

        initProducer();
      }
  });
}

function initProducer() {

  client.initProducer("T_C_03eaa1d3-4642-4079-b030-b543ee04b5af", "T_S_f50ecb46-ca82-44a7-9c40-421818af5996", function(err, response) {

    console.log("initProducer.callback");
    console.log("initProducer.err: " + err)
    console.log("initProducer.response: %j", response)

    if(err == null) {

      startBroadcast();
    }
  });
}

function startBroadcast() {

  client.startServiceBroadcast(20000, function(err, response){

    console.log("startServiceBroadcast.callback");
    console.log("startServiceBroadcast.err: " + err)
    console.log("startServiceBroadcast.response: %j", response)
  });
}
