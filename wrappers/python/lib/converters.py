try:
    from ttypes import *
except ImportError:
    from .ttypes import *

wptypes_thrift = thriftpy.load('wptypes.thrift', module_name="wptypes_thrift")

import wptypes_thrift as wpt

def ConvertToThriftPPU(ppu):
    return wpt.PricePerUnit(amount=ppu.amount, currencyCode=ppu.currencyCode)

def ConvertToThriftPrice(price):
    ppu = ConvertToThriftPPU(price.pricePerUnit)
    return wpt.Price(id=price.id, description=price.description, pricePerUnit=ppu, unitId=price.unitId, unitDescription=price.unitDescription)
        
def ConvertToThriftService(service):
    thriftPrices = {}
    for key, value in service.prices():
        thriftPrices[key] = ConvertoToThriftPrice(value)
    return wpt.Service(id=service.id, name=service.name, description=service.description, prices=thriftPrices)

def ConvertToThriftHCECard(card):
    return wpt.HCECard(firstName=card.firstName, lastName=card.lastName, expMonth=card.expMonth, expYear=card.expYear, cardNumber=card.cardNumber, cardType=card.cardType, cvc=card.cvc)

def ConvertToThriftTotalPriceResponse(response):
    return wpt.TotalPriceResponse(serverId=response.serverId, clientId=response.clientId, priceId=response.priceId, unitsToSupply=response.unitsToSupply, totalPrice=response.totalPrice, paymentReferenceId=response.paymentReferenceId, merchantClientKey=response.merchantClientKey)

def ConvertToThriftServiceDeliveryToken(token):
    return wpt.ServiceDeliveryToken(key=token.key, issued=token.issued, expiry=token.expiry, refundOnExpiry=token.refundOnExpiry, signature=token.signature)

def ConvertFromThriftPPU(ppu):
    return PricePerUnit(amount=ppu.amount, currencyCode=ppu.currencyCode)

def ConvertFromThriftPrice(price):
    ppu = ConvertFromThriftPPU(price.pricePerUnit)
    return Price(id=price.id, description=price.description, pricePerUnit=ppu, unitId=price.unitId, unitDescription=price.unitDescription)
        
def ConvertFromThriftService(service):
    prices = {}
    for key, value in service.prices():
        prices[key] = ConvertoFromThriftPrice(value)
    return Service(id=service.id, name=service.name, description=service.description, prices=prices)

def ConvertFromThriftDevice(device):
    services = {}
    for key, value in device.services():
        services[key] = ConvertoFromThriftService(value)
    return Device(uid=device.uid, name=device.name, description=device.description, services=services, ipv4Address=device.ipv4Address, currencyCode=device.currencyCode)

def ConvertFromThriftServiceMessage(message):
    return ServiceMessage(deviceDescription=message.deviceDescription, hostname=message.hostname, portNumber=message.portNumber, serverId=message.serverId, urlPrefix=message.urlPrefix, scheme=message.scheme)

def ConvertFromThriftServiceDetails(details):
    return ServiceDetails(serviceId=details.serviceId, serviceDescription=details.serviceDescription)

def ConvertFromThriftTotalPriceResponse(response):
    return TotalPriceResponse(serverId=response.serverId, clientId=response.clientId, priceId=response.priceId, unitsToSupply=response.unitsToSupply, totalPrice=response.totalPrice, paymentReferenceId=response.paymentReferenceId, merchantClientKey=response.merchantClientKey)

def ConvertFromThriftServiceDeliveryToken(token):
    return ServiceDeliveryToken(key=token.key, issued=token.issued, expiry=token.expiry, refundOnExpiry=token.refundOnExpiry, signature=token.signature)

def ConvertFromThriftPaymentResponse(response):
    token = ConvertFromThriftServiceDeliveryToken(token.serviceDeliveryToken)
    return PaymentResponse(serverId=response.serverId, clientId=response.clientId, totalPaid=response.totalPaid, serviceDeliveryToken=token)