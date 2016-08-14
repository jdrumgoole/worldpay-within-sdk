/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation.wpwithin.types;

import java.util.HashMap;

/**
 *
 * @author worldpay
 */
public class WWDevice {
// 	1: string uid
//	2: string name
//	3:  string description
//	4:  map<i32, Service> services
//	5: string ipv4Address
//	6: string currencyCode 
    
    String uid;
    String name;
    String description;
    HashMap<Integer, WWService> services;
    String ipv4Address;
    String currencyCode;

    public String getUid() {
        return uid;
    }

    public void setUid(String uid) {
        this.uid = uid;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public HashMap<Integer, WWService> getServices() {
        return services;
    }

    public void setServices(HashMap<Integer, WWService> services) {
        this.services = services;
    }

    public String getIpv4Address() {
        return ipv4Address;
    }

    public void setIpv4Address(String ipv4Address) {
        this.ipv4Address = ipv4Address;
    }

    public String getCurrencyCode() {
        return currencyCode;
    }

    public void setCurrencyCode(String currencyCode) {
        this.currencyCode = currencyCode;
    }
    
    
}
