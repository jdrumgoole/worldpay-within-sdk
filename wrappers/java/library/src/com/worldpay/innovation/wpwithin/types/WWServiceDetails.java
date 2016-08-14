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
public class WWServiceDetails {

//    	1: i32 serviceId
//	2: string serviceDescription
                
    int serviceId;
    String serviceDescription;

    public int getServiceId() {
        return serviceId;
    }

    public void setServiceId(int serviceId) {
        this.serviceId = serviceId;
    }

    public String getServiceDescription() {
        return serviceDescription;
    }

    public void setServiceDescription(String serviceDescription) {
        this.serviceDescription = serviceDescription;
    }
    
    
}
