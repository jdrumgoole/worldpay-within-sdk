package com.worldpay.innovation.wpwithin.thriftadapter;

import com.worldpay.innovation.wpwithin.rpc.types.ServiceDeliveryToken;
import com.worldpay.innovation.wpwithin.types.WWServiceDeliveryToken;

/**
 * Created by conor on 13/08/2016.
 */
public class ServiceDeliveryTokenAdapter {

    public static WWServiceDeliveryToken convertServiceDeliveryToken(ServiceDeliveryToken serviceDeliveryToken) {
        WWServiceDeliveryToken wwServiceDeliveryToken = new WWServiceDeliveryToken();
        wwServiceDeliveryToken.setExpiry(serviceDeliveryToken.getExpiry());
        wwServiceDeliveryToken.setIssued(serviceDeliveryToken.getIssued());
        wwServiceDeliveryToken.setKey(serviceDeliveryToken.getKey());
        //wwServiceDeliveryToken.setRefundOnExpiry(serviceDeliveryToken.get); // TODO What to do about this?
        wwServiceDeliveryToken.setSignature(serviceDeliveryToken.getSignature());
        return wwServiceDeliveryToken;
    }

    public static ServiceDeliveryToken convertWWServiceDeliveryToken(WWServiceDeliveryToken wwServiceDeliveryToken) {
        ServiceDeliveryToken serviceDeliveryToken = new ServiceDeliveryToken();
        serviceDeliveryToken.setExpiry(wwServiceDeliveryToken.getExpiry());
        serviceDeliveryToken.setIssued(wwServiceDeliveryToken.getIssued());
        serviceDeliveryToken.setKey(wwServiceDeliveryToken.getKey());
        //wwServiceDeliveryToken.setRefundOnExpiry(serviceDeliveryToken.get); // TODO What to do about this?
        serviceDeliveryToken.setSignature(wwServiceDeliveryToken.getSignature());
        return serviceDeliveryToken;
    }
}
