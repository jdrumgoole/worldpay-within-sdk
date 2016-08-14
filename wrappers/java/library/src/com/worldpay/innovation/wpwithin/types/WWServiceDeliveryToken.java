/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation.wpwithin.types;

import java.nio.ByteBuffer;

/**
 *
 * @author worldpay
 */
public class WWServiceDeliveryToken {

//    	1: string key
//	2: string issued
//	3: string expiry
//	4: bool refundOnExpiry
//	5: binary signature
                
    String key;
    String issued;
    String expiry;
    boolean refundOnExpiry;
    byte[] signature;

    public String getKey() {
        return key;
    }

    public void setKey(String key) {
        this.key = key;
    }

    public String getIssued() {
        return issued;
    }

    public void setIssued(String issued) {
        this.issued = issued;
    }

    public String getExpiry() {
        return expiry;
    }

    public void setExpiry(String expiry) {
        this.expiry = expiry;
    }

    public boolean isRefundOnExpiry() {
        return refundOnExpiry;
    }

    public void setRefundOnExpiry(boolean refundOnExpiry) {
        this.refundOnExpiry = refundOnExpiry;
    }

    public byte[] getSignature() {
        return signature;
    }

    public void setSignature(byte[] signature) {
        this.signature = signature;
    }
    
    
    
}
