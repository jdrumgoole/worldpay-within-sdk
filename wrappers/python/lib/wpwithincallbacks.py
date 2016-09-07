import thriftpy
from thriftpy.rpc import make_client
from thriftpy.protocol.binary import TBinaryProtocolFactory
from thriftpy.transport.buffered import TBufferedTransportFactory
import launcher
import time

try:
    from ttypes import *
except ImportError:
    from .ttypes import *

wptypes_thrift = thriftpy.load('wptypes.thrift', module_name="wptypes_thrift")

import wptypes_thrift as wpt

class WPWithinCallback(object):
    def __init__(self, thriftClient):
        self.thriftClient = thriftClient

    def beginServiceDelivery(self, serviceId, serviceDeliveryToken, unitsToSupply):
        try:
            self.thriftClient.beginServiceDelivery(self, serviceId, serviceDeliveryToken, unitsToSupply)
        except wpt.Error as err:
            raise Error(err.message)

    def endServiceDelivery(self, serviceId, serviceDeliveryToken, unitsToSupply):
        try:
            self.thriftClient.beginServiceDelivery(self, serviceId, serviceDeliveryToken, unitsToSupply)
        except wpt.Error as err:
            raise Error(err.message)

def createClient(host, port):
    wpw_thrift = thriftpy.load('wpwithin.thrift', module_name="wpw_thrift")

    # add try ...
    TClient = make_client(wpw_thrift.WPWithinCallbacks, host, port, proto_factory=TBinaryProtocolFactory(), trans_factory=TBufferedTransportFactory())
    
    return WPWithinCallback(TClient)