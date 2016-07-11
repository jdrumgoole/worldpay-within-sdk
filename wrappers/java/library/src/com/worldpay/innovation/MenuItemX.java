/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation;

import java.lang.reflect.Method;

/**
 *
 * @author worldpay
 */
public class MenuItemX {
    
    final private String label;
    final private Method action;
    
    public MenuItemX(String _label, Method _action) {
        this.label = _label;
        this.action = _action;
    }
    
    public String getLabel() {
        return this.label;
    }
    
    public Method getAction() {
        return this.action;
    }
}
