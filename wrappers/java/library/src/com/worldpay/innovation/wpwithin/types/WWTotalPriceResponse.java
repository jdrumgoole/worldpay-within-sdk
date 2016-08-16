/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation.wpwithin.types;

/**
 *
 * @author worldpay
 */
public class WWTotalPriceResponse {
//	1: string serverId
//	2: string clientId
//	3: i32 priceId
//	4: i32 unitsToSupply
//	5: i32 totalPrice
//	6: string paymentReferenceId
//	7: string merchantClientKey
    
    String serverId;
    String clientId;
    int priceId;
    int unitsToSupply;
    int totalPrice;
    String paymentReferenceId;
    String merchantClientKey;

    public String getServerId() {
        return serverId;
    }

    public void setServerId(String serverId) {
        this.serverId = serverId;
    }

    public String getClientId() {
        return clientId;
    }

    public void setClientId(String clientId) {
        this.clientId = clientId;
    }

    public int getPriceId() {
        return priceId;
    }

    public void setPriceId(int priceId) {
        this.priceId = priceId;
    }

    public int getUnitsToSupply() {
        return unitsToSupply;
    }

    public void setUnitsToSupply(int unitsToSupply) {
        this.unitsToSupply = unitsToSupply;
    }

    public int getTotalPrice() {
        return totalPrice;
    }

    public void setTotalPrice(int totalPrice) {
        this.totalPrice = totalPrice;
    }

    public String getPaymentReferenceId() {
        return paymentReferenceId;
    }

    public void setPaymentReferenceId(String paymentReferenceId) {
        this.paymentReferenceId = paymentReferenceId;
    }

    public String getMerchantClientKey() {
        return merchantClientKey;
    }

    public void setMerchantClientKey(String merchantClientKey) {
        this.merchantClientKey = merchantClientKey;
    }
    
    
    
}
