/*
 * Worldpay Within Java Wrapper
 */
package com.worldpay.innovation.wpwithin;

import com.worldpay.innovation.wpwithin.eventlistener.EventListener;
import com.worldpay.innovation.wpwithin.types.WWDevice;
import com.worldpay.innovation.wpwithin.types.WWHCECard;
import com.worldpay.innovation.wpwithin.types.WWPaymentResponse;
import com.worldpay.innovation.wpwithin.types.WWPrice;
import com.worldpay.innovation.wpwithin.types.WWService;
import com.worldpay.innovation.wpwithin.types.WWServiceDeliveryToken;
import com.worldpay.innovation.wpwithin.types.WWServiceDetails;
import com.worldpay.innovation.wpwithin.types.WWServiceMessage;
import com.worldpay.innovation.wpwithin.types.WWTotalPriceResponse;
import java.util.Set;

/**
 * Worldpay Within Java Wrapper Interface to be coding against and will give
 * you access to the underlying Golang SDK, via RPC calls made by Thrift.
 *
 * @author worldpay
 */
public interface WPWithinWrapper {

    /**
     * Sets up the wrapper to be able to start communicating with the
     * underlying SDK.
     *
     * @param name The name of the device
     * @param description The description of the device
     * @throws WPWithinGeneralException
     */
    public void setup(String name, String description) throws WPWithinGeneralException;

    /**
     * Adds a service of type WWSerive to the producer, used if the device you
     * are operating on is a producer, if added to a device you intend as a
     * consumer this will give that device producer functionality.
     *
     * @param svc This is the service to be added
     * @throws WPWithinGeneralException
     */
    public void addService(WWService svc) throws WPWithinGeneralException;

    /**
     * This removes the service from the producer
     *
     * @param svc The service to remove
     * @throws WPWithinGeneralException
     */
    public void removeService(WWService svc) throws WPWithinGeneralException;

    /**
     * This initiates the device as a consumer, which enables it to find
     * services, negotiate prices, make payments and receive services.
     *
     * @param scheme This should be 'http://' or 'https://'
     * @param hostname This will be host name, find this from the service messages
     * @param port This will be the port for the device, find this from the service message
     * @param urlPrefix This will be the urlPrefix for which the restful interface of the producer is located at
     * @param serverId This is the unique Id of the server of the producer, provided by the service message
     * @param hceCard This is the payment credentials to configure the consumer with, which it will use to make payments
     * @throws WPWithinGeneralException
     */
    public void initConsumer(String scheme, String hostname, Integer port, String urlPrefix, String serverId, WWHCECard hceCard) throws WPWithinGeneralException;

    /**
     * This initiates the device as a producer / or initialises the devices producer
     * capability.
     *
     * @param merchantClientKey This is the online.worldpay.com client key for your test account
     * @param merchantServiceKey This is the online.worlpday.com merchant service key (secret key) from your test account
     * @throws WPWithinGeneralException
     */
    public void initProducer(String merchantClientKey, String merchantServiceKey) throws WPWithinGeneralException;

    /**
     * This is able to provide back details of the the current device that the
     * SDK is running on, and it credentials / information
     *
     * @return The actual device object with it's details
     * @throws WPWithinGeneralException
     */
    public WWDevice getDevice() throws WPWithinGeneralException;

    /**
     * This enables the producer device to start broadcasting itself via UDP
     * broadcast over the network to notifiy devices it is available to be
     * consumed.
     *
     * @param timeoutMillis This is the timeout for the braodcast e.g. 68000
     * @throws WPWithinGeneralException
     */
    public void startServiceBroadcast(Integer timeoutMillis) throws WPWithinGeneralException;

    /**
     * This method stops the SDK from broadcasting the current service messages
     * that it is broadcasting
     *
     * @throws WPWithinGeneralException
     */
    public void stopServiceBroadcast() throws WPWithinGeneralException;

    /**
     * This enables the consumer device to discovery other devices (producers)
     * on the network that are UDP broadcasting.
     *
     * @param timeoutMillis This is the timeout in milli seconds e.g. 68000
     * @return The Set of service messages that are found to be broadcasting
     * are returned
     * @throws WPWithinGeneralException
     */
    public Set<WWServiceMessage> deviceDiscovery(Integer timeoutMillis) throws WPWithinGeneralException;

    /**
     * TODO: What does this do? This requests a services that are available???
     * @return Returns the ServiceDetails object which contains all the details
     * about the service.
     * @throws WPWithinGeneralException
     */
    public Set<WWServiceDetails> requestServices() throws WPWithinGeneralException;

    /**
     * This is used by the consumer to get the list of prices associated with
     * a particular serviceId
     *
     * @param serviceId The serviceId of the service the consumer would like to
     * get the prices for.
     * @return Returns a set of prices that the producer offers for this service
     * @throws WPWithinGeneralException
     */
    public Set<WWPrice> getServicePrices(Integer serviceId) throws WPWithinGeneralException;

    /**
     * Selection of a service is performed by the consumer, providing details of
     * the service, the amount and at what price point it wants to purchase
     * the service.
     *
     * @param serviceId The id of the service to be purchased
     * @param numberOfUnits The number of units to purchase
     * @param priceId The PriceId of the price that is to be chosen
     * @return The TotalPriceResponse object is provided back which details how
     * much the service will cost based on the requested service details.
     * @throws WPWithinGeneralException
     */
    public WWTotalPriceResponse selectService(Integer serviceId, Integer numberOfUnits, Integer priceId) throws WPWithinGeneralException;

    /**
     * This allows the consumer to request a payment be made at the producer
     * device, by providing the total price response object as the request. The
     * producer will then make the payment (or attempt to) and send back a
     * Payment Response detailing whether it was successful or not.
     *
     * @param request The TotalPriceResponse object detailing the agreement of
     * the service that will be purchased.
     * @return Returns the Payment Response as to whether it was successful
     * or not
     * @throws WPWithinGeneralException
     */
    public WWPaymentResponse makePayment(WWTotalPriceResponse request) throws WPWithinGeneralException;

    /**
     * This begins the service delivery, and is requested by the consumer, and
     * will proceed as long as the correct information is provided to the
     * producer. If the correct credentials are passed through, then the
     * producer will start releasing the service known as a 'trusted trigger'
     *
     * @param serviceId Id of service being delivered
     * service
     * @param serviceDeliveryToken The secure token that details the usage of
     * the service, and the producer will know if it has been paid for and is
     * still valid
     * @param unitsToSupply This is how many units of service being requested
     * at this time.
     * @throws WPWithinGeneralException
     */
    public void beginServiceDelivery(int serviceId, WWServiceDeliveryToken serviceDeliveryToken, Integer unitsToSupply) throws WPWithinGeneralException;

    /**
     * This ends the service delivery, a request initiated by the consumer.
     *
     * @param serviceID Id of service being delivered
     * @param serviceDeliveryToken The secure token that details the usage of the
     * service
     * @param unitsReceived The number of units received - to confirm that this is
     * consistent with what the produce understands
     * @throws WPWithinGeneralException
     */
    public void endServiceDelivery(int serviceID, WWServiceDeliveryToken serviceDeliveryToken, Integer unitsReceived) throws WPWithinGeneralException;

    void stopRPCAgent();
}
