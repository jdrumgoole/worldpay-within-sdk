var wpwithin = require('../../library/wpwithin');
var types = require('../../library/types/types');
var typesConverter = require('../../library/types/converter');
var client;

wpwithin.createClient("127.0.0.1", 9090, true, function(err, response){

  console.log("createClient.callback")
  console.log("createClient.callback.err: " + err)
  console.log("createClient.callback.response: %j", response);

  if(err == null) {

      client = response;

      setup();
  }
});

function setup() {

  client.setup("NodeJS-Device", "Sample NodeJS consumer device", function(err, response){

    console.log("setup.callback.err: " + err);
    console.log("setup.callback.response: %j", response);

    console.log("Calling discover devices..");
    discoverDevices();
  })
};

function discoverDevices() {

  client.deviceDiscovery(10000, function(err, response) {

    console.log("deviceDiscovery.callback.err: %s" + err);
    console.log("deviceDiscovery.callback.response: %j", response);

    if(response != null && response.length > 0) {

      console.log("Discovered %d devices on the network.", response.length);
      console.log("Devices:");

      for(var i = 0; i < response.length; i++) {

        console.log("Description: %s", response[i].deviceDescription);
        console.log("Hostname: %s", response[i].hostname);
        console.log("Port: %d", response[i].portNumber);
        console.log("Server ID: %s", response[i].serverId);
        console.log("URL Prefix: %s", response[i].urlPrefix);

        console.log("-------")
      }

      // Connect to the first device
      var serviceMessage = response[0];

      connectToDevice(serviceMessage);

    } else {

      console.log("Did not discover any devices on the network.");
    }

  });
};

function connectToDevice(serviceMessage) {

  var hceCard = new types.HCECard();
  hceCard.firstname = "Bilbo";
  hceCard.lastname = "Baggins";
  hceCard.expMonth = 01;
  hceCard.expYear = 2018;
  hceCard.cardNumber = "5555555555554444";
  hceCard.type = "Card";
  hceCard.cvc = "123";

  client.initConsumer("http://", serviceMessage.hostname, serviceMessage.portNumber,
  serviceMessage.urlPrefix, serviceMessage.serverId, hceCard, function(err, response){

    console.log("initConsumer.callback.err: %s" + err);
    console.log("initConsumer.callback.response: %j", response);

    if(err == null) {

      console.log("Did initialise consumer.")

      getAvailableServices();
    }
  });
}

function getAvailableServices() {

  client.requestServices(function(err, response) {

    console.log("requestServices.callback.err: %s" + err);
    console.log("requestServices.callback.response: %j", response);

    if(err == null && response != null && response.length > 0) {

      var svc = response[0];

      console.log("Services:");
      console.log("Id: %s", svc.serviceId);
      console.log("Description: %s", svc.serviceDescription);
      console.log("----------");

      getServicePrices(svc.serviceId);
    }
  });
}

function getServicePrices(serviceId) {

  client.getServicePrices(serviceId, function(err, response) {

    console.log("requestServicePrices.callback.err: %s" + err);
    console.log("requestServicePrices.callback.response: %j", response);

    if(err == null && response != null && response.length > 0) {

        var price = response[0];

        console.log("Price details for ServiceId: %d", serviceId);
        console.log("Id: %d", price.id);
        console.log("Description: %s", price.description);
        console.log("UnitId: %d", price.unitId);
        console.log("unitDescription: %s", price.unitDescription);
        console.log("PricePerUnit:");
        console.log("\tAmount: %d", price.pricePerUnit.amount);
        console.log("\tCurrency Code: %s", price.pricePerUnit.currencyCode);
        console.log("----------");

        getServicePriceQuote(serviceId, 1, price.id);
    } else {

      console.log("Did not receive any service prices :/");
    }
  });
}

function getServicePriceQuote(serviceId, numberOfUnits, priceId) {

  client.selectService(serviceId, numberOfUnits, priceId, function(err, response) {

    console.log("selectService.callback.err: %s" + err);
    console.log("selectService.callback.response: %j", response);

    if(err == null && response != null) {

      console.log("TotalPriceResponse:");
      console.log("ServerId: %s", response.serverId);
      console.log("ClientId: %s", response.clientId);
      console.log("PriceId: %d", response.priceId);
      console.log("UnitsToSupply: %d", response.unitsToSupply);
      console.log("TotalPrice: %d", response.totalPrice);
      console.log("PaymentReferenceId: %s", response.paymentReferenceId);
      console.log("MerchantClientKey: %s", response.merchantClientKey);
      console.log("------");

      purchaseService(response);

    } else {

      console.log("Did not receive total price response from selectService()");
    }
  });
}

function purchaseService(totalPriceResponse) {

  client.makePayment(totalPriceResponse, function(err, response) {

    console.log("makePayment.callback.err: %s" + err);
    console.log("makePayment.callback.response: %j", response);

    if(err == null && response != null) {

      console.log("Resonse from make payment:")
      console.log("ServerID: %s", response.serverId);
      console.log("ClientID: %s", response.clientId);
      console.log("TotalPaid: %d", response.totalPaid);
      console.log("ServiceDeliveryToken:");
      console.log("\tKey: %s", response.serviceDeliveryToken.key);
      console.log("\tIssued: %s", response.serviceDeliveryToken.issued);
      console.log("\tExpiry: %s", response.serviceDeliveryToken.expiry);
      console.log("\tRefundOnExpiry: %b", response.serviceDeliveryToken.refundOnExpiry);
      console.log("\tSignature: %s", response.serviceDeliveryToken.signature);
      console.log("ClientUUID: %s", response.clientUUID);
      console.log("----------");

    } else {

      console.log("Did not receive correct response to make payment..");
    }
  });
}
