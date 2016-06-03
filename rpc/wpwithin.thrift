##############################################
#
# Worldpay Within SDK Thrift definition
# Conor Hackett (conor.hackett@worldpay.com)
# June 3rd, 2016
#
#############################################

include "wptypes.thrift"

namespace csharp wpwithin
namespace java com.worldpay.innovation.wpwithin.rpc
namespace go worldpay_within.wpwithin.rpc

/**
 * WorldpayWithin Service - exposing all WorldpayWithin SDK functionality
 */
service WPWithin {

  /**
   * A method definition looks like C code. It has a return type, arguments,
   * and optionally a list of exceptions that it may throw. Note that argument
   * lists and exception lists are specified using the exact same syntax as
   * field lists in struct or exception definitions.
   */

   void addService(1: wptypes.Service svc) throws (1: wptypes.Error err),
   void removeService(1: wptypes.Service svc) throws (1: wptypes.Error err),
   void initHCE(1: wptypes.HCECard hceCard) throws (1: wptypes.Error err),
   void initHTE(1: string merchantClientKey, 2: string merchantServiceKey) throws (1: wptypes.Error err),
   void initConsumer(1: string scheme, 2: string hostname, 3: i32 port, 4: string urlPrefix, 5: string serviceId) throws (1: wptypes.Error err),
   void initProducer() throws (1: wptypes.Error err),
   wptypes.Device getDevice(),
   void startServiceBroadcast(1: i32 timeoutMillis) throws (1: wptypes.Error err),
   void stopServiceBroadcast() throws (1: wptypes.Error err),
   set<wptypes.ServiceMessage> serviceDiscovery(1: i32 timeoutMillis) throws (1: wptypes.Error err),
   set<wptypes.ServiceDetails> requestServices() throws (1: wptypes.Error err),
   set<wptypes.Price> getServicePrices(1: i32 serviceId) throws (1: wptypes.Error err),
   wptypes.TotalPriceResponse selectService(1: i32 serviceId, 2: i32 numberOfUnits, 3: i32 priceId) throws (1: wptypes.Error err),
   wptypes.PaymentResponse makePayment(1: wptypes.TotalPriceResponse request) throws (1: wptypes.Error err),

}