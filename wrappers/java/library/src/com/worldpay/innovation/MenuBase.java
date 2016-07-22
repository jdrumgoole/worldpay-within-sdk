/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import java.util.Scanner;
import java.util.logging.ConsoleHandler;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.apache.thrift.TException;

/**
 *
 * @author worldpay
 */
public class MenuBase {
   
    protected final String DEFAULT_DEVICE_NAME = "conorhwp-macbook";
    protected final String DEFAULT_DEVICE_DESCRIPTION = "Conor H WP - Raspberry Pi";    
    private final WPWithin.Client sdk;
    
    public MenuBase(WPWithin.Client _client) {
        this.sdk = _client;
    }
    
    public MenuBase() {
        this.sdk = null;
    } 
    
    public MenuReturnStruct mInitDefaultDevice() {

            //_sdk, err := wpwithin.Initialise(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION)
            
            try {
                sdk.setup(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION);
            } catch(TException e) {
                return new MenuReturnStruct("SDK setup failed", 1);
            }
        
            return new MenuReturnStruct(null, 0);

    }
    
    public MenuReturnStruct mInitNewDevice()  {

            System.out.println("Name of device: ");
            
            Scanner scanner = new Scanner(System.in);
            String nameOfDevice = scanner.next();
            if(null == nameOfDevice || "".equals(nameOfDevice)) {
                    return new MenuReturnStruct("Name of device not set", 0);
            }

            System.out.println("Description: ");
            String description = scanner.next();
            if(null == description || "".equals(description)) {
                    return new MenuReturnStruct("Description of device not set", 0);
            }
                    
            try {
                sdk.setup(nameOfDevice, description);
            } catch(TException e) {
                return new MenuReturnStruct("Setup of device unsucessful", 0);
            }
            
            return new MenuReturnStruct(null, 0);
            
    }    
}
