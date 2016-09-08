class Error(object):
    def __init__(self, message):
        super().__init__(message)
        self.message = message

class PricePerUnit(object):
    def __init__(self, amount, currencyCode):
        self.amount = amount
        self.currencyCode = currencyCode

class Price(object):
    def __init__(self, id, description, pricePerUnit, unitId, unitDescription):
        self.id = id
        self.description = description
        self.pricePerUnit = pricePerUnit
        self.unitId = unitId
        self.unitDescription = unitDescription

class Service(object):
    def __init__(self, id, name, description, prices=None):
        self.id = id
        self.name = name
        self.description = description
        self.prices = prices

class HCECard(object):
    def __init__(self, firstName, lastName, expMonth, expYear, cardNumber, cardType, cvc):
        self.firstName = firstName
        self.lastName = lastName
        self.expMonth = expMonth
        self.expYear = expYear
        self.cardNumber = cardNumber
        self.cardType = cardType
        self.cvc = cvc

class Device(object):
    def __init__(self, uid, name, description, services, ipv4Address, currencyCode):
        self.uid = uid
        self.name = name
        self.description = description 
        self.services = services
        self.ipv4Address = ipv4Address
        self.currencyCode = currencyCode

class ServiceMessage(object):
    def __init__(self, deviceDescription, hostname, portNumber, serverId, urlPrefix, scheme):
        self.deviceDescription = deviceDescription
        self.hostname = hostname
        self.portNumber = portNumber
        self.serverId = serverId
        self.urlPrefix = urlPrefix
        self.scheme = scheme

class ServiceDetails(object):
    def __init__(self, serviceId, serviceDescription):
        self.serviceId = serviceId
        self.serviceDescription = serviceDescription

class TotalPriceResponse(object):
    def __init__(self, serverId, clientId, priceId, unitsToSupply, totalPrice, paymentReferenceId, merchantClientKey):
        self.serverId = serverId
        self.clientId = clientId
        self.priceId = priceId
        self.unitsToSupply = unitsToSupply
        self.totalPrice = totalPrice
        self.paymentReferenceId = paymentReferenceId
        self.merchantClientKey = merchantClientKey

class ServiceDeliveryToken(object):
    def __init__(self, key, issued, expiry, refundOnExpiry, signature):
        self.key = key
        self.issued = issued
        self.expiry = expiry
        self.refundOnExpiry = refundOnExpiry
        self.signature = signature

class PaymentResponse(object):
    def __init__(self, serverId, clientId, totalPaid, serviceDeliveryToken):
        self.serverId = serverId
        self.clientId = clientId
        self.totalPaid = totalPaid
        self.serviceDeliveryToken = serviceDeliveryToken