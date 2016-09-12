package com.worldpay.innovation.wpwithin.eventlistener;

import com.worldpay.innovation.wpwithin.WPWithinGeneralException;
import com.worldpay.innovation.wpwithin.types.WWServiceDeliveryToken;

/**
 * Events returned from Worldpay Within Core SDK
 */
public interface EventListener {

    /**
     * Begin service delivery requested by consumer.
     *
     * @param serviceId Id of service being delivered
     * @param serviceDeliveryToken service delivery token used for service delivery
     * @param unitsToSupply the number of units the consumer wishes to receive
     */
    void onBeginServiceDelivery(int serviceId, WWServiceDeliveryToken serviceDeliveryToken, int unitsToSupply) throws WPWithinGeneralException;

    /**
     * End service delivery requested by consumer
     *
     * @param serviceId Id of service being delivered
     * @param serviceDeliveryToken service delivery token used for delivery
     * @param unitsReceived the number of units that the consumer says it received
     */
    void onEndServiceDelivery(int serviceId, WWServiceDeliveryToken serviceDeliveryToken, int unitsReceived) throws WPWithinGeneralException;
}
