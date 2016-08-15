module.exports = {
  Error : Error,
  Service : Service,
  Price : Price,
  PricePerUnit : PricePerUnit,
  HCECard : HCECard,
  Device : Device,
  ServiceMessage : ServiceMessage,
  ServiceDetails : ServiceDetails,
  TotalPriceResponse : TotalPriceResponse,
  PaymentResponse : PaymentResponse,
  ServiceDeliveryToken : ServiceDeliveryToken
};

function Error() {

  var message;
};

function Service() {

  var id;
  var name;
  var description;
  var prices;
};

function Price() {

  var id;
  var description;
  var pricePerUnit;
  var unitId;
  var unitDescription;
};

function PricePerUnit() {

  var amount;
  var currencyCode;
};

function HCECard() {

  var firstname;
  var lastname;
  var expMonth;
  var expYear;
  var cardNumber;
  var type;
  var cvc;
};

function Device() {

  var uid;
  var name;
  var description;
  var services;
  var ipv4Address;
  var currencyCode;
};

function ServiceMessage() {

  var deviceDescription;
  var hostname;
  var portNumber;
  var serverId;
  var urlPrefix;
};

function ServiceDetails() {

  var serviceId;
  var serviceDescription;
};

function TotalPriceResponse() {

  var serverId;
  var clientId;
  var priceId;
  var unitsToSupply;
  var totalPrice;
  var paymentReferenceId;
  var merchantClientKey;
};

function PaymentResponse() {

  var serverId;
  var clientId;
  var totalPaid;
  var serviceDeliveryToken;
  var clientUUID;
}

function ServiceDeliveryToken() {

  var key;
  var issued;
  var expiry;
  var refundOnExpiry;
  var signature;
}
