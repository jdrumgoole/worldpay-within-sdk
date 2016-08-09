function WPWithin() {

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

var fnSetup = function(name, description) {
}

var fnAddService = function(service) {

}

var fnRemoveService = function(service) {

}

var fnInitConsumer = function(scheme, hostname, port, urlPrefix, serverId, hceCard) {

}

var fnInitProducer = function(merchantClientKey, merchantServiceKey) {

}

var fnGetDevice = function() {

}

var fnStartServiceBroadcast = function(timeoutMillis) {

}

var fnStopServiceBroadcast = function() {

}

var fnDeviceDiscovery = function(timeoutMillis) {

}

var fnRequestServices = function() {

}

var fnGetServicePrices = function(serviceId) {

}

var fnSelectService = function(serviceId, numberOfUnits, priceId) {

}

var fnMakePayment = function(request) {

}

var fnBeginServiceDelivery = function(clientId, serviceDeliveryToken, unitsToSupply) {

}

var fnEndServiceDelivery = function(clientId, serviceDeliveryToken, unitsReceived) {

}
