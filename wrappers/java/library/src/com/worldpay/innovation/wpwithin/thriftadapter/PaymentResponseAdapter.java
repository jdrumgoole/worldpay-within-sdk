package com.worldpay.innovation.wpwithin.thriftadapter;

import com.worldpay.innovation.wpwithin.rpc.types.PaymentResponse;
import com.worldpay.innovation.wpwithin.types.WWPaymentResponse;

/**
 * Created by conor on 13/08/2016.
 */
public class PaymentResponseAdapter {

    public static PaymentResponse convertWWPaymentResponse(WWPaymentResponse wwPaymentResponse) {
        PaymentResponse paymentResponse = new PaymentResponse();
        paymentResponse.setClientId(wwPaymentResponse.getClientId());
        //paymentResponse.setClientUuid(wwPaymentResponse.getClientUUID());
        paymentResponse.setServerId(wwPaymentResponse.getServerId());
        paymentResponse.setServiceDeliveryToken(ServiceDeliveryTokenAdapter.convertWWServiceDeliveryToken(wwPaymentResponse.getServiceDeliveryToken()));
        paymentResponse.setTotalPaid(wwPaymentResponse.getTotalPaid());
        return paymentResponse;
    }

    public static WWPaymentResponse convertPaymentResponse(PaymentResponse paymentResponse) {
        WWPaymentResponse wwPaymentResponse = new WWPaymentResponse();
        wwPaymentResponse.setClientId(paymentResponse.getClientId());
        //wwPaymentResponse.setClientUuid(paymentResponse.getClientUUID());
        wwPaymentResponse.setServerId(paymentResponse.getServerId());
        wwPaymentResponse.setServiceDeliveryToken(ServiceDeliveryTokenAdapter.convertServiceDeliveryToken(paymentResponse.getServiceDeliveryToken()));
        wwPaymentResponse.setTotalPaid(paymentResponse.getTotalPaid());
        return wwPaymentResponse;
    }
}
