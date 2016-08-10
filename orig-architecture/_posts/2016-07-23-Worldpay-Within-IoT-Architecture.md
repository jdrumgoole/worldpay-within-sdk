---
layout: post
title:  "3. Worldpay Within IoT Service Architecture "
date:   2016-07-27 12:13:03 +0100
categories: jekyll update
---
Worldpay Within IoT Service Architecture 
=========================================

![](media/image5.emf){width="6.497916666666667in" height="6.4375in"}The
provision of a service within the Worldpay IoT system is performed in 4
phases, as shown in Figure 3, these being: Service Discovery; Service
Negotiation; Payment; and Service Delivery. Each of these phases are
described in the following sections.

Service Discovery
-----------------

Each Thing that offers services, the service ‘supplier’ shall broadcast
it’s list of available services, as shown in Figure 4 below. When a
potential ‘consumer’ of the service connects with ‘supplier’ it can
request details of the services offered.

![](media/image6.emf){width="4.59375in"
height="2.3854166666666665in"}Providing a suitable service is
discovered, the consumer then requests the service from the supplier,
and price negotiations can begin.

### Service Discovery APIs

  **Key**              **Parameters**                   **Purpose**
  -------------------- -------------------------------- -----------------------------------------------------------------------------
  broadcast            server\_UUID                     Advertising services and identifying the sender
  request services     &lt;none&gt;                     Request a list of all services
  services\_response   list of services, server\_UUID   Provide client with a list of possible services that the sender can provide

### Service discovery messages

A broadcast message that includes Thing B’s UUID is sent.

Upon receiving the message Thing A connects to Thing B and requests the
list of available services.

Thing B responds with a list identifying the services available.

Service Negotiation
-------------------

![](media/image7.emf){width="4.802083333333333in"
height="2.6979166666666665in"}Once a suitable service has been
discovered, there will be a price negotiation. The provider may offer
the same service at different rates depending on the number of units of
service to be purchased. The process is outlined in Figure 5. The
outcome of the process is an agreement to purchase an amount of service
and a total price for the service to be provided. The service provider
can then request payment for the agreed service and price.

### Service Negotiation APIs

  -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  **Key**                   **Parameters**                                                                                                    **Purpose**
  ------------------------- ----------------------------------------------------------------------------------------------------------------- -----------------------------------------------------------------------------------------------------------
  price\_request            service\_id                                                                                                       Request a list of all prices for a given service.

  price\_response           server\_UUID, list of prices,                                                                                     Provide the client with a list of prices for a given service. A price object contains the per unit price.
                                                                                                                                              
                            (service\_id, price\_id, price\_per\_unit, unit\_ID, unit\_description, price\_description)                       

  price\_select             service\_id, price\_id, number\_of\_units, client\_UUID                                                           Select a price with price\_id, for service\_id for a number of units.

  price\_select\_response   price\_id, number\_of\_units, total\_price, server\_UUID, client\_UUID, payment\_ref\_ID, Merchant\_Client\_key   Communicate the expected total price to the client.
  -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

### Service negotiation messages

A price request is sent containing the selected service\_id.

The response from Thing B contains a list of price items; each item
should contain a price\_id, per unit price, unit\_ID and description
fields of both the unit and the price.

Thing A then selects an appropriate price\_id by sending a request with
its client\_UUID, the selected service\_id, the price\_id, and the
number of items required.

If the number of items falls within the correct number of items for the
price selected, then Thing B responds with a price select response
containing the service\_id, price\_id, the total price, the
service\_UUID and a reference for the payment and its Merchant Client
key. Otherwise Thing B shall return the number of units it can supply
along with the correct price, and additional details required to
initiate the payment.

Payment
-------

The payment process with Online.worldpay.com is a two stage process,
split between the consumer and merchant Things involved in the
transaction, these stages are:

-   Client Token Request, and

-   Payment Authorisation Request. (Also known as Order Request)

During the first stage, the consumer sends Online.worldpay.com their
payment credentials and the merchants Client Key. Online.worldpay.com
returns a Client Token, which the consumer passes to the Merchant,
allowing the merchant to perform the payment authorisation request with
Online.worldpay.com by providing the Client Token and transaction
details.

This payment process ensures that the consumer does not pass their
payment credentials to the merchant, only to Online.worldpay.com.

### Client token request

![](media/image8.emf){width="5.40625in"
height="2.6666666666666665in"}The first step in the payment process is
when Thing A receives the Merchant\_Client\_Key from Thing B. Thing B
passes their public Client Key to Thing A as part of the
price\_select\_response during the Service Negotiation phase. Upon
receiving the Client Key from Thing B, Thing A connects with
Online.worldpay.com to request the client token from
Online.worldpay.com. This request includes Thing A’s payment
credentials: Card PAN, expiry, and the client\_key of Thing B.
Online.worldpay.com will respond with a message that includes a
client\_token. This is shown in Figure 6.

### Client token request APIs

#### Thing A to Online.worldpay.com client token request

  ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  **Key**                   **Parameters**                                                                                                                                                                                                                                                           **Purpose**
  ------------------------- ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ ---------------------------------------------------------------------------------------------------------------------
  client\_token\_request    Payment\_method, reusable\_flag, Merchant\_client\_key                                                                                                                                                                                                                   Request a client token from Online.worldpay.com, whilst providing Online.worldpay.com with the payment credentials.
                                                                                                                                                                                                                                                                                                     
                            Payment\_method (name, PAN, expiryMonth, expiryYear, type)                                                                                                                                                                                                               

  client\_token\_response   client\_token, reusable\_flag, payment\_method\_response                                                                                                                                                                                                                 Response from Online.worldpay.com containing the client\_token.
                                                                                                                                                                                                                                                                                                     
                            (type, name, expiryMonth, expiryYear, card\_type, card\_scheme\_type, card\_scheme\_name, masked\_card\_number, card\_product\_type\_description\_non\_contactless, card\_product\_type\_description\_contactless, card\_issuer, country\_code, card\_class, pre-paid)   

  Payment\_request          client\_token, client\_UUID, payment\_ref\_ID                                                                                                                                                                                                                            The client\_token is passed to Thing B to allow the 2^nd^ part of the transaction process to take place.
  ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Thing A will connect to Online.worldpay.com using TLS. It will then
request a client\_token by securely (see 2.2.1) sending a JSON message
containing the paymentMethod, its payment credentials (PAN, expiry) to
Online.worldpay.com along with the client\_ key from Thing B. In
addition a flag indicating if the client details can be used in future
is sent, for IoT this should always be set ‘reusable’:’false’ in order
to force generation of a new client token for each transaction.

A successful response will be an HTTP POST response containing fields:
client\_token, reusable\_flag and the payment\_method\_response. Once
received, the client\_token shall be passed to Thing B

A sample request is shown in Appendix B: Sample Service Messaging.

See Online.worldpay.com documentation for client\_token\_request &
client\_token\_repsonse APIs data descriptions.

### Payment authorisation request

Thing B will process the order and request the payment from
Online.worldpay.com providing its Service key, client\_token,
transaction currency and payment amount. This is transmitted to
Online.worldpay.com over TLS. After successful processing
Online.worldpay.com will provide a payment response. Thing B shall then
generate a service token, which Thing A may use in future to obtain the
services that the payment has been made for. This is shown in Figure 7.

### ![](media/image9.emf){width="6.25in" height="2.4166666666666665in"}Payment authorisation request APIs

#### Thing B to Online.worldpay.com payment authorisation request

  --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
  **Key**           **Parameters**                                                                                                                                                                                                                                                           **Purpose**
  ----------------- ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ -------------------------------------------------------------------------------------------
  order\_request    client\_service\_key, client\_token, currency\_code, amount, order\_description, customer\_order\_code                                                                                                                                                                   Request payment from Online.worldpay.com.

  order\_response   order\_code, client\_token, order\_description, amount, currency\_code, payment\_status, customer\_order\_code, environment, risk\_score, payment\_response                                                                                                              Payment response indicating a successful transaction on the Online.worldpay.com platform.
                                                                                                                                                                                                                                                                                             
                    (type, name, expiryMonth, expiryYear, card\_type, card\_scheme\_type, card\_scheme\_name, masked\_card\_number, card\_product\_type\_description\_non\_contactless, card\_product\_type\_description\_contactless, card\_issuer, country\_code, card\_class, pre-paid)   
  --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Thing B shall assemble a message to be posted to Online.worldpay.com
that contains the client token, Service key, the amount, currency and
transaction description. Online.worldpay.com shall then perform an
authorisation using the payment credentials identified by the
client\_token. A successful authorisation will result in a
payment\_status of SUCCESS being returned to Thing B.

#### Thing B to Thing A service token

  **Key**                      **Parameters**                                                      **Purpose**
  ---------------------------- ------------------------------------------------------------------- -----------------------------------------------
  payment\_request\_response   service\_delivery\_token, server\_UUID, client\_UUID, total\_paid   service\_delivery\_token is passed to ThingB.

Thing B shall then generate a cryptographically secure
service\_delivery\_token, which can be used by Thing A to request
provision of services from Thing B.

Service Delivery
----------------

![](media/image10.emf){width="3.8020833333333335in"
height="3.3333333333333335in"}Once the payment has been made, Thing B
shall return to broadcasting its available services. Thing A will now be
able to consume the service from Thing B by providing the
service\_delivery\_token. The service delivery may be in a single step,
or over time. An overview of service delivery is shown in Figure 8.

Once in possession of a service\_token, Thing A may then request the
service be provided. The service could be consumed in one session, or in
several sessions over time, depending on the nature of the service and
number of units purchased. Thing A may repeatedly send service delivery
requests until Thing B indicates that the service has been delivered.

### Service Delivery APIs

  **Key**                     **Parameters**                                                                                                          **Purpose**
  --------------------------- ----------------------------------------------------------------------------------------------------------------------- -------------------------------------------------------------------------------------------------------------------------------------------
  broadcast                   server\_UUID                                                                                                            Advertising services and identifying the sender.
  delivery\_begin\_request    service\_delivery\_token, client\_UUID, number\_of\_units\_to\_supply                                                   Request the service item, with the service\_delivery\_token providing right to receive the service, and amount of service to be supplied.
  delivery\_begin\_response   server\_UUID, service\_delivery\_token, client\_UUID, number\_of\_units\_to\_be\_supplied                               Response for the service delivery. Confirmation of number of service units to be supplied (Allowing for less units than requested).
  delivery\_end               client\_UUID, number\_of\_units\_received                                                                               Confirmation of service received.
  delivery\_end\_response     server\_UUID, service\_delivery\_token, client\_UUID, number\_of\_units\_just\_supplied, number\_of\_units\_remaining   Service end indicating outstanding service credits and token for subsequent delivery.

Thing A sends a message with the service\_delivery\_token to Thing B,
along with the amount of service it wishes to consume. The response
shall confirm the amount of service units that Thing B can supply to
Thing A at that time. Once the service has been delivered, Thing A shall
confirm the amount of service units it has received, with Thing B
responding, stating the number of units still remaining to Thing A, if
any.