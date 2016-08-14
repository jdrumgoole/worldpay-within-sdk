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
public class WWPaymentResponse {

//    	1: string serverId
//	2: string clientId
//	3: i32 totalPaid
//	4: ServiceDeliveryToken serviceDeliveryToken
//	5: string ClientUUID
                
    String serverId;
    String clientId;
    int totalPaid;
    WWServiceDeliveryToken serviceDeliveryToken;
    String clientUuid;

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

    public int getTotalPaid() {
        return totalPaid;
    }

    public void setTotalPaid(int totalPaid) {
        this.totalPaid = totalPaid;
    }

    public WWServiceDeliveryToken getServiceDeliveryToken() {
        return serviceDeliveryToken;
    }

    public void setServiceDeliveryToken(WWServiceDeliveryToken serviceDeliveryToken) {
        this.serviceDeliveryToken = serviceDeliveryToken;
    }

    public String getClientUuid() {
        return clientUuid;
    }

    public void setClientUuid(String clientUuid) {
        this.clientUuid = clientUuid;
    }
    
    
    
}
