package com.worldpay.innovation.types;

import com.worldpay.innovation.wpwithin.rpc.types.ServiceDetails;

/**
 * Created by conor on 29/07/2016.
 */
public class ServiceDeliveryToken {

    public String key;
    public String issued;
    public String expiry;
    public boolean refundOnExpiry;
    public byte[] signature;

    public ServiceDeliveryToken(String key, String issued, String expiry, boolean refundOnExpiry, byte[] signature) {

        this.key = key;
        this.issued = issued;
        this.expiry = expiry;
        this.refundOnExpiry = refundOnExpiry;
        this.signature = signature;
    }

    public String getKey() {
        return key;
    }

    public String getIssued() {
        return issued;
    }

    public String getExpiry() {
        return expiry;
    }

    public boolean isRefundOnExpiry() {
        return refundOnExpiry;
    }

    public byte[] getSignature() {
        return signature;
    }
}
