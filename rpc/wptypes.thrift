##############################################
#
# Worldpay Within SDK Thrift definition
# Conor Hackett (conor.hackett@worldpay.com)
# June 3rd, 2016
#
#############################################

namespace csharp wptypes
namespace java com.worldpay.innovation.wpwithin.rpc.types

exception Error {
	
	1: string message
}

struct Service {
	
	1: i32 id
	2: string name
	3: optional string description
	4: optional map<i32, Price> prices
}

struct Price {
	
	1: i32 serviceId
	2: i32 id
	3: string description
	4: i32 pricePerUnit
	5: i32 unitId
	6: string unitDescription
}

struct HCECard {
	
	1: string FirstName
	2: string LastName
	3: i32 ExpMonth
	4: i32 ExpYear
	5: string CardNumber
	6: string Type
	7: optional string Cvc
}

struct Device {
	
	1: string uid
	2: string name
	3: optional string description
	4: optional map<i32, Service> services
	5: string ipv4Address
	6: string currencyCode
}

struct ServiceMessage {
	
	1: string deviceDescription
	2: string hostname
	3: i32 portNumber
	4: string serverId
	5: string urlPrefix
}

struct ServiceDetails {
	
	1: i32 serviceId
	2: string serviceDescription
}

struct TotalPriceResponse {
	
	1: string serverId
	2: string clientId
	3: i32 priceId
	4: i32 unitsToSupply
	5: i32 totalPrice
	6: string paymentReferenceId
	7: string merchantClientKey
}

struct PaymentResponse {
	
	1: string serverId
	2: string clientId
	3: i32 totalPaid
	4: string serviceDeliveryToken
	5: string ClientUUID
}