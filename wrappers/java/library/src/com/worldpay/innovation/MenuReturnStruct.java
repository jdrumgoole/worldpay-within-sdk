/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation;

/**
 *
 * @author worldpay
 */
public class MenuReturnStruct {
    final private String msg;
    final private int returnCode;
    
    public MenuReturnStruct(String _msg, int _returnCode) {
        this.msg = _msg;
        this.returnCode = _returnCode;
    }
    
    public String getMsg() {
        return this.msg;
    }
    
    public int getReturnCode() {
        return this.returnCode;
    }
}
