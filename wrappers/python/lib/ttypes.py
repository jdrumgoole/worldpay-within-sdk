class Error(Exception):
    def __init__(self, message=None):
        super(ValidationError, self).__init__(message)
        self.message = message

class PricePerUnit(object):
    def __init__(self, amount=None, currencyCode=None):
        self.amount = amount
        self.currencyCode = currencyCode

class Price(object):
    def __init__(self, id=None, description=None, pricePerUnit=None, unitId=None, unitDescription=None):
        self.id = id
        self.description = description
        self.pricePerUnit = pricePerUnit
        self.unitId = unitId
        self.unitDescription = unitDescription

class Service(object):
    def __init__(self, id=None, name=None, description=None, prices=None):
        self.id = id
        self.name = name
        self.description = description
        self.prices = prices

class HCECard(object):
    def __init__(self, FirstName=None, LastName=None, ExpMonth=None, ExpYear=None, CardNumber=None, Type=None, Cvc=None):
        self.FirstName = FirstName
        self.LastName = LastName
        self.ExpMonth = ExpMonth
        self.ExpYear = ExpYear
        self.CardNumber = CardNumber
        self.Type = Type
        self.Cvc = Cvc

class Device(object):
    def __init__(self, uid=None, name=None, description=None, services=None, ipv4Address=None, currencyCode=None):
        self.uid = uid
        self.name = name
        self.description = description 
        self.services = services
        self.ipv4Address = ipv4Address
        self.currencyCode = currencyCode

class ServiceMessage(object):
    def __init__(self, deviceDescription=None, hostname=None, portNumber=None, serverId=None, urlPrefix=None):
        self.deviceDescription = deviceDescription
        self.hostname = hostname
        self.portNumber = portNumber
        self.serverId = serverId
        self.urlPrefix = urlPrefix

class ServiceDetails(object):
    def __init__(self, serviceId=None, serviceDescription=None):
        self.serviceId = serviceId
        self.serviceDescription = serviceDescription

class TotalPriceResponse(object):
    def __init__(self, serverId=None, clientId=None, priceId=None, unitsToSupply=None, totalPrice=None, paymentReferenceId=None, merchantClientKey=None):
        self.serverId = serverId
        self.clientId = clientId
        self.priceId = priceId
        self.unitsToSupply = unitsToSupply
        self.totalPrice = totalPrice
        self.paymentReferenceId = paymentReferenceId
        self.merchantClientKey = merchantClientKey

class ServiceDeliveryToken(object):
    def __init__(self, key=None, issues=None, expiry=None, refundOnExpiry=None, signature=None):
        self.key = key
        self.issued = issued
        self.expiry = expiry
        self.refundOnExpiry = refundOnExpiry
        self.signature = signature

class PaymentResponse(object):
    def __init__(self, serverId=None, clientId=None, totalPaid=None, serviceDeliveryToken=None, ClientUUID=None):
        self.serverId = serverId
        self.clientId = clientId
        self.totalPaid = totalPaid
        self.serviceDeliveryToken = serviceDeliveryToken 
        self.ClientUUID = ClientUUID