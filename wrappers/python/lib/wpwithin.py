import thriftpy
from thriftpy.rpc import make_client, make_server
from thriftpy.protocol.binary import TBinaryProtocolFactory
from thriftpy.transport.buffered import TBufferedTransportFactory
import launcher
import time
import wpwithincallbacks

try:
    from ttypes import *
except ImportError:
    from .ttypes import *

wptypes_thrift = thriftpy.load('wptypes.thrift', module_name="wptypes_thrift")

import wptypes_thrift as wpt

class WPWithin(object):
    def __init__(self, thriftClient):
        self.thriftClient = thriftClient

    def setup(self, name, description):
        try:
            self.thriftClient.setup(name, description)
        except wpt.Error as err:
            raise Error(err.message)

    def addService(self, svc):
        try:
            self.thriftClient.addService(svc)
        except wpt.Error as err:
            raise Error(err.message)

    def removeService(self, svc):
        try:
            self.thriftClient.removeService(svc)
        except wpt.Error as err:
            raise Error(err.message)

    def initConsumer(self, scheme, hostname, port, urlPrefix, clientId, hceCard):
        try:
            self.thriftClient.initConsumer(scheme, hostname, port, urlPrefix, clientId, hceCard)
        except wpt.Error as err:
            raise Error(err.message)

    def initProducer(self, merchantClientKey, merchantServiceKey):
        try:
            self.thriftClient.initProducer(merchantClientKey, merchantServiceKey)
        except wpt.Error as err:
            raise Error(err.message)

    def getDevice(self):
        return self.thriftClient.getDevice()

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
            return serviceMessages

    def requestServices(self):
        try:
            serviceDetails = self.thriftClient.requestServices()
        except wpt.Error as err:
            raise Error(err.message)
        else:
            return serviceDetails

    def getServicePrices(self, serviceId):
        try:
            prices = self.thriftClient.getServicePrices(serviceId)
        except wpt.Error as err:
            raise Error(err.message)
        else:
            return prices

    def selectServices(self, serviceId, numberOfUnits, priceId):
        try:
            service = self.thriftClient.selectServices(serviceId, numberOfUnits, priceId)
        except wpt.Error as err:
            raise Error(err.message)
        else:
            return service

    def makePayment(self, request):
        try:
            response = self.thriftClient.makePayment(request)
        except wpt.Error as err:
            raise Error(err.message)
        else:
            return response

    def beginServiceDelivery(self, serviceId, serviceDeliveryToken, unitsToSupply):
        try:
            serviceDeliveryToken = self.thriftClient.beginServiceDelivery(clientId, serviceDeliveryToken, unitsToSupply)
        except wpt.Error as err:
            raise Error(err.message)
        else:
            return serviceDeliveryToken

    def endServiceDelivery(self, serviceId, serviceDeliveryToken, unitsReceived):
        try:
            serviceDeliveryToken = self.thriftClient.endServiceDelivery(clientId, serviceDeliveryToken, unitsReceived)
        except wpt.Error as err:
            raise Error(err.message)
        else:
            return serviceDeliveryToken

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