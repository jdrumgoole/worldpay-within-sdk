import thriftpy
from thriftpy.rpc import make_client, make_server
from thriftpy.protocol.binary import TBinaryProtocolFactory
from thriftpy.transport.buffered import TBufferedTransportFactory
import launcher
import time
import wpwithincallbacks

wptypes_thrift = thriftpy.load('wptypes.thrift', module_name="wptypes_thrift")

import wptypes_thrift as wpt

from ttypes import *
from converters import *


class WPWithin(object):
    def __init__(self, thriftClient):
        self.thriftClient = thriftClient

    def setup(self, name, description):
        try:
            self.thriftClient.setup(name, description)
        except wpt.Error as err:
            raise Error(err.message)

    def addService(self, svc):
        service = ConvertToThriftService(svc)
        try:
            self.thriftClient.addService(service)
        except wpt.Error as err:
            raise Error(err.message)

    def removeService(self, svc):
        service = ConvertToThriftService(svc)
        try:
            self.thriftClient.removeService(service)
        except wpt.Error as err:
            raise Error(err.message)

    def initConsumer(self, scheme, hostname, port, urlPrefix, clientId, hceCard):
        card = ConvertToThriftHCECard(hceCard)
        try:
            self.thriftClient.initConsumer(scheme, hostname, port, urlPrefix, clientId, card)
        except wpt.Error as err:
            raise Error(err.message)

    def initProducer(self, merchantClientKey, merchantServiceKey):
        try:
            self.thriftClient.initProducer(merchantClientKey, merchantServiceKey)
        except wpt.Error as err:
            raise Error(err.message)

    def getDevice(self):
        return ConvertFromThriftDevice(self.thriftClient.getDevice())

    def startServiceBroadcast(self, timeoutMillis):
        try:
            self.thriftClient.startServiceBroadcast(timeoutMillis)
        except wpt.Error as err:
            raise Error(err.message)

    def stopServiceBroadcast(self):
        try:
            self.thriftClient.stopServiceBroadcast
        except wpt.Error as err:
            raise Error(err.message)

    def deviceDiscovery(self, timeoutMillis):
        try:
            serviceMessages = self.thriftClient.deviceDiscovery(timeoutMillis)
        except wpt.Error as err:
            raise Error(err.message)
        else:
            svcMessages = []
            for val in serviceMessages:
                svcMessages.append(ConvertFromThriftServiceMessage(val))
            return svcMessages

    def requestServices(self):
        try:
            serviceDetails = self.thriftClient.requestServices()
        except wpt.Error as err:
            raise Error(err.message)
        else:
            svcDetails = []
            for val in serviceDetails:
                svcDetails.append(ConvertFromThriftServiceDetails(val))
            return svcDetails

    def getServicePrices(self, serviceId):
        try:
            prices = self.thriftClient.getServicePrices(serviceId)
        except wpt.Error as err:
            raise Error(err.message)
        else:
            wprices = []
            for val in prices:
                wprices.append(ConvertFromThriftPrice(val))
            return wprices

    def selectServices(self, serviceId, numberOfUnits, priceId):
        try:
            service = self.thriftClient.selectServices(serviceId, numberOfUnits, priceId)
        except wpt.Error as err:
            raise Error(err.message)
        else:
            return ConvertFromThriftTotalPriceResponse(service)

    def makePayment(self, request):
        trequest = ConvertToThriftTotalPriceResponse(request)
        try:
            response = self.thriftClient.makePayment(trequest)
        except wpt.Error as err:
            raise Error(err.message)
        else:
            return ConvertFromThriftPaymentResponse(response)

    def beginServiceDelivery(self, serviceId, serviceDeliveryToken, unitsToSupply):
        token = ConvertToThriftServiceDeliveryToken(serviceDeliveryToken)
        try:
            serviceDeliveryToken = self.thriftClient.beginServiceDelivery(clientId, token, unitsToSupply)
        except wpt.Error as err:
            raise Error(err.message)

    def endServiceDelivery(self, serviceId, serviceDeliveryToken, unitsReceived):
        token = ConvertToThriftServiceDeliveryToken(serviceDeliveryToken)
        try:
            serviceDeliveryToken = self.thriftClient.endServiceDelivery(clientId, token, unitsReceived)
        except wpt.Error as err:
            raise Error(err.message)

def runRPCAgent(port, dir="./rpc-agent/", callbackPort=None):
    return launcher.runRPCAgent(dir, port, callbackPort=callbackPort)

def createClient(host, port, startRPC, startCallbackServer=False, callbackPort=None, callbackService=None, rpcDir=None):
    
    if (startCallbackServer == True) and (callbackPort == None or callbackService == None):
        raise ValueError('No callback port or service provided')

    wpw_thrift = thriftpy.load('wpwithin.thrift', module_name="wpw_thrift")

    returnDict = {}
    
    if startRPC:
        if rpcDir == None and startCallbackServer == None:
            proc = runRPCAgent(port)
        elif rpcDir == None:
            proc = runRPCAgent(port, callbackPort=callbackPort)
        elif startCallbackServer == None:
            proc = runRPCAgent(port, dir=rpcDir)
        else:
            proc = runRPCAgent(port, dir=rpcDir, callbackPort=callbackPort)
        returnDict['rpc'] = proc

    time.sleep(2)
    # add try ...
    TClient = make_client(wpw_thrift.WPWithin, host=host, port=port, proto_factory=TBinaryProtocolFactory(), trans_factory=TBufferedTransportFactory())
    
    if startCallbackServer:
        server = make_server(callbackService, wpwithincallbacks, host=host, port=callbackPort, proto_factory=TBinaryProtocolFactory(), trans_factory=TBufferedTransportFactory()) 
        returnDict['server'] = server

    if len(returnDict) > 0:
        returnDict['client'] = WPWithin(TClient)
        return returnDict

    return WPWithin(TClient)