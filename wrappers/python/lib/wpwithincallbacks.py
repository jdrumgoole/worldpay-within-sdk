import thriftpy

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