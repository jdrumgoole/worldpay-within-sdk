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
public class WWService {
    //struct Service {
    //	
    //	1: i32 id
    //	2: string name
    //	3:  string description
    //	4: optional map<i32, Price> prices  /* This should be optional now but these are stored as pointers and was causing an issue in Go - TODO CH - Conor to investigate and fix */
    //}
    
    int id;
    String name;
    String description;
    HashMap<Integer, WWPrice> prices;

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
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

    public HashMap<Integer, WWPrice> getPrices() {
        return prices;
    }

    public void setPrices(HashMap<Integer, WWPrice> prices) {
        this.prices = prices;
    }
    
    
    
}
