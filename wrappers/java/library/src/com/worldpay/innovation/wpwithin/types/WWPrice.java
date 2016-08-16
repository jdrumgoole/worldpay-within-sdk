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
public class WWPrice {

//    	1: i32 id
//	2: string description
//	3: PricePerUnit pricePerUnit
//	4: i32 unitId
//	5: string unitDescription

    int id;
    String description;
    WWPricePerUnit pricePerUnit;
    int unitId;
    String unitDescription;

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public WWPricePerUnit getPricePerUnit() {
        return pricePerUnit;
    }

    public void setPricePerUnit(WWPricePerUnit pricePerUnit) {
        this.pricePerUnit = pricePerUnit;
    }

    public int getUnitId() {
        return unitId;
    }

    public void setUnitId(int unitId) {
        this.unitId = unitId;
    }

    public String getUnitDescription() {
        return unitDescription;
    }

    public void setUnitDescription(String unitDescription) {
        this.unitDescription = unitDescription;
    }

    
    
}
