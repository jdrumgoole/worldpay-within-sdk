/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.types.Price;
import com.worldpay.innovation.wpwithin.rpc.types.PricePerUnit;
import com.worldpay.innovation.wpwithin.rpc.types.Service;
import java.util.Scanner;
import java.util.logging.ConsoleHandler;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.apache.thrift.TException;

/**
 *
 * @author worldpay
 */
public class ProducerMenu {
    
    // TODO: put this somewhere sensible
    final private String ERR_DEVICE_NOT_INITIALISED = "Error: Device not initialised";
    final private String DEFAULT_HTE_MERCHANT_CLIENT_KEY = "T_C_03eaa1d3-4642-4079-b030-b543ee04b5af";
    final private String DEFAULT_HTE_MERCHANT_SERVICE_KEY = "T_S_f50ecb46-ca82-44a7-9c40-421818af5996";

    private static final Logger log= Logger.getLogger( ProducerMenu.class.getName() );    
    private final WPWithin.Client sdk;
    
    public ProducerMenu(WPWithin.Client _client) {
        log.setLevel(Level.FINE);
        ConsoleHandler handler = new ConsoleHandler();
        handler.setLevel(Level.FINE);
        log.addHandler(handler);
        this.sdk = _client;
    }
    
    public MenuReturnStruct mBroadcast() {
	System.out.print("Broadcast timeout in milliseconds: ");
        Scanner scanner = new Scanner(System.in);
	String inputStr = scanner.next();
        try {
            int input = (new Integer(inputStr));
        } catch(NumberFormatException e) {
            return new MenuReturnStruct("Could not get correct input number", 0);
        }
        return new MenuReturnStruct(null, 0);
    }

    
    public MenuReturnStruct mProducerStatus() {

	// Show all services
	// Show all prices
	// Status of broadcast

	return new MenuReturnStruct("Not implemented yet..", 0);
    }
    
    public MenuReturnStruct mDefaultProducer(){

	try {
            mInitDefaultDevice();
        } catch(Exception e) {
            return new MenuReturnStruct("failed to init default device", 0);
	}

        try {
            mDefaultHTECredentials();
        } catch(Exception e) {
            return new MenuReturnStruct("failed to default HTE credentials", 0);
	}


        if(sdk == null) {
            return new MenuReturnStruct(ERR_DEVICE_NOT_INITIALISED, 0);
	}

        try {
            sdk.initProducer();
        } catch(TException e) {
            return new MenuReturnStruct("failed to init producer via SDK", 0);
	}

	return new MenuReturnStruct(null, 0);
    }
    
    
    public MenuReturnStruct mNewProducer() {
        
            try {
                mInitNewDevice();
            } catch(Exception e) {
                return new MenuReturnStruct("failed to init new device", 0);
            }

            try {
                mNewHTECredentials();
            } catch(Exception e) {
                return new MenuReturnStruct("failed to new HTE credentials", 0);
            }

            if(sdk == null) {
                return new MenuReturnStruct(ERR_DEVICE_NOT_INITIALISED, 0);
            }

            return new MenuReturnStruct(null, 0);
    }
    
    public MenuReturnStruct mDefaultHTECredentials() {

        if(sdk == null) {
            return new MenuReturnStruct(ERR_DEVICE_NOT_INITIALISED, 0);
        }
        
        try {
            sdk.initHTE(DEFAULT_HTE_MERCHANT_CLIENT_KEY, DEFAULT_HTE_MERCHANT_SERVICE_KEY);
        } catch(TException e) {
            return new MenuReturnStruct("Failed to initiate HTE", 0);
        }    
            
        return new MenuReturnStruct(null, 0);
        
    }

    public MenuReturnStruct mNewHTECredentials() {

	System.out.print("Merchant Client Key: ");
	String merchantClientKey;
        String merchantServiceKey;
        
        
        try {
            Scanner scanner = new Scanner(System.in);
            merchantClientKey = scanner.next();
        } catch(Exception e) {
            return new MenuReturnStruct("Could not read in merchant client key", 0);
        }
        
        try {
            Scanner scanner = new Scanner(System.in);
            merchantServiceKey = scanner.next();
        } catch(Exception e) {
            return new MenuReturnStruct("Could not read in merchant secret key", 0);
        }

        if(sdk == null) {
            return new MenuReturnStruct(ERR_DEVICE_NOT_INITIALISED, 0);
        }
                
        try {
            sdk.initHTE(merchantClientKey, merchantServiceKey);
        } catch(TException e) {
            return new MenuReturnStruct("Failed to initiate HTE", 0);
        }    
            
        return new MenuReturnStruct(null, 0);
    }
    
    public MenuReturnStruct mAddRoboWashService() {

            Service roboWash = new Service();
            roboWash.setName("RoboWash");
            roboWash.setDescription("Car washed by robot");
            roboWash.setId(1);
            
            Price washPriceCar = new Price();
            washPriceCar.setServiceId(roboWash.getId());
            washPriceCar.setUnitId(1);
            washPriceCar.setId(1);
            washPriceCar.setDescription("Car wash");
            washPriceCar.setUnitDescription("Single wash");
            washPriceCar.setPricePerUnit(new PricePerUnit(500, "GBP"));
//            washPriceCar.setPricePerUnit(new PricePerUnit(500));

            Price washPriceSUV = new Price();
            washPriceSUV.setServiceId(roboWash.getId());
            washPriceSUV.setUnitId(1);
            washPriceSUV.setId(2);
            washPriceSUV.setDescription("SUV wash");
            washPriceSUV.setUnitDescription("Single wash");
            washPriceSUV.setPricePerUnit(new PricePerUnit(650, "GBP"));
            //washPriceSUV.setPricePerUnit(new PricePerUnit(650));
            
            roboWash.putToPrices(0, washPriceCar);
            roboWash.putToPrices(1, washPriceSUV);

            if(sdk == null) {
                return new MenuReturnStruct(ERR_DEVICE_NOT_INITIALISED, 0);
            }

            try {
                sdk.addService(roboWash); 
            } catch(TException e) {
                return new MenuReturnStruct("Failed to add Service for roboWash", 0);
            }  
            
            return new MenuReturnStruct(null, 0);
    }
    
    
    public MenuReturnStruct mAddRoboAirService() {

	Service roboAir = new Service();
	roboAir.setName("RoboAir");
	roboAir.setDescription("Car tyre pressure checked and topped up by robot");
	roboAir.setId(2);

        Price airSinglePrice = new Price();
        airSinglePrice.setServiceId(roboAir.getId());
        airSinglePrice.setUnitId(1);
        airSinglePrice.setId(1);
        airSinglePrice.setDescription("Measure and adjust pressue");
        airSinglePrice.setUnitDescription("Tyre");
        airSinglePrice.setPricePerUnit(new PricePerUnit(25, "GBP"));
        //airSinglePrice.setPricePerUnit(25);

        
        Price airFourPrice = new Price();
        airFourPrice.setServiceId(roboAir.getId());
        airFourPrice.setUnitId(2);
        airFourPrice.setId(2);
        airFourPrice.setDescription("Measure and adjust pressure - four tyres for the price of three");
        airFourPrice.setUnitDescription("4 Tyre");
        airFourPrice.setPricePerUnit(new PricePerUnit(airSinglePrice.getPricePerUnit().getAmount() * 3, "GBP"));

	roboAir.putToPrices(0, airSinglePrice);
	roboAir.putToPrices(1, airFourPrice);

        if(sdk == null) {
            return new MenuReturnStruct(ERR_DEVICE_NOT_INITIALISED, 0);
        }

        try {
            sdk.addService(roboAir); 
        } catch(TException e) {
            return new MenuReturnStruct("Failed to addService for roboAir", 0);
        }  

        return new MenuReturnStruct(null, 0);                
                
                
    }
    
    
    
    public MenuReturnStruct mStartBroadcast(){

        if(sdk == null) {
            return new MenuReturnStruct(ERR_DEVICE_NOT_INITIALISED, 0);
        }

        
	System.out.print("Broadcast timeout in milliseconds: ");
        
        Scanner scanner = new Scanner(System.in);
	String timeoutStr = scanner.next();
        int timeout;
        try {
            timeout = (new Integer(timeoutStr));
        } catch(NumberFormatException e) {
            return new MenuReturnStruct("Could not get correct input number for timeout", 0);
        }
               
        try {
            sdk.startServiceBroadcast(timeout); 
        } catch(TException e) {
            return new MenuReturnStruct("Failed to start the service broadcast", 0);
        }  

        return new MenuReturnStruct(null, 0);
        
    }
    
    public MenuReturnStruct mStopBroadcast() {

        return new MenuReturnStruct("Not implemented yet...", 0);
                
    }
    
    public MenuReturnStruct mCarWashDemoProducer() {

        
        MenuReturnStruct rc = mDefaultProducer();  
        if(rc.getMsg() != null) return rc;
        
        mAddRoboWashService();
        
        MenuReturnStruct rc2 = mAddRoboAirService();  
        if(rc2.getMsg() != null) return rc2;
        
        try {
            sdk.startServiceBroadcast(20000); 
        } catch(TException e) {
            return new MenuReturnStruct("Start service broadcast failed", 0);
        }          
        
	return new MenuReturnStruct(null, 0);
    }
    
    // TODO: To be moved to superclass
    // TODO: put these somewhere sensible
    // TODO: What do these do and what should they be?
    private final String DEFAULT_DEVICE_NAME = "conorhwp-macbook";
    private final String DEFAULT_DEVICE_DESCRIPTION = "Conor H WP - Raspberry Pi";    
    public MenuReturnStruct mInitDefaultDevice() {

            //_sdk, err := wpwithin.Initialise(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION)
            
            try {
                sdk.setup(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION);
            } catch(TException e) {
                return new MenuReturnStruct("SDK setup failed", 1);
            }
        
            return new MenuReturnStruct(null, 0);

    }
    
    // TODO: To be moved to superclass
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