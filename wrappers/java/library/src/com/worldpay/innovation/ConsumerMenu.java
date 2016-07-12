/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.types.HCECard;
import java.util.logging.ConsoleHandler;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.apache.thrift.TException;

/**
 *
 * @author worldpay
 */
public class ConsumerMenu {
    private static final Logger log = Logger.getLogger( ConsumerMenu.class.getName() );
        
    private final WPWithin.Client sdk;
    
    public ConsumerMenu(WPWithin.Client _client) {
        log.setLevel(Level.FINE);
        ConsoleHandler handler = new ConsoleHandler();
        handler.setLevel(Level.FINE);
        log.addHandler(handler);
        this.sdk = _client;
    }    
    
    public MenuReturnStruct mDefaultHCECredential() {

        HCECard card = new HCECard("Bilbo", "Baggins", 11, 2018, "5555555555554444", "Card", "113");

	if(this.sdk == null) {
            return new MenuReturnStruct("ERR_DEVICE_NOT_INITIALISED", 0);
	}

        try {
            sdk.initHCE(card);
        } catch(TException e) {
            return new MenuReturnStruct("Issue initialising the HCE card", 0);    
        }
        
	return new MenuReturnStruct(null, 0);
    }
}
