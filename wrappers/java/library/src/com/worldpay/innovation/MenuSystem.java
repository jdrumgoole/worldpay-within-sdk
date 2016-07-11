/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation;

import static java.lang.System.err;
import java.lang.reflect.InvocationTargetException;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.logging.Level;
import java.util.logging.Logger;


/**
 *
 * @author worldpay
 */
public class MenuSystem {
    private static final Logger log= Logger.getLogger( MenuSystem.class.getName() );

    private ArrayList menuItems;
    
    public MenuSystem() {
    }
    
    public MenuReturnStruct mInvalidSelection() {
	return new MenuReturnStruct("*** Invalid menu selection - please choose another item ***", 0);
    }
    
    public void doUI() {
        try {
            menuItems = new ArrayList();
 
            menuItems.add(new MenuItemX("-------------------- GENERAL  --------------------", MenuSystem.class.getMethod("mInvalidSelection")));
//            menuItems.add(new MenuItemX("Init default device", MenuSystem.class.getMethod("mInitDefaultDevice")));
//            menuItems.add(new MenuItemX("Start RPC Service", MenuSystem.class.getMethod("mStartRPCService")));
//            menuItems.add(new MenuItemX("Init new device", MenuSystem.class.getMethod("mInitNewDevice")));
//            menuItems.add(new MenuItemX("Get device info", MenuSystem.class.getMethod("mGetDeviceInfo")));
//            menuItems.add(new MenuItemX("Sample demo, car wash (Producer)", MenuSystem.class.getMethod("mCarWashDemoProducer")));
//            menuItems.add(new MenuItemX("Sample demo, car wash (Consumer)", MenuSystem.class.getMethod("mCarWashDemoConsumer")));
//            menuItems.add(new MenuItemX("Reset session", MenuSystem.class.getMethod("mResetSessionState")));
//            menuItems.add(new MenuItemX("Load configuration", MenuSystem.class.getMethod("mLoadConfig")));
//            menuItems.add(new MenuItemX("Read loaded configuration", MenuSystem.class.getMethod("mReadConfig")));
            menuItems.add(new MenuItemX("-------------------- PRODUCER --------------------", MenuSystem.class.getMethod("mInvalidSelection")));
//            menuItems.add(new MenuItemX("Create default producer", MenuSystem.class.getMethod("mDefaultProducer")));
//            menuItems.add(new MenuItemX("Create new producer", MenuSystem.class.getMethod("mNewProducer")));
//            menuItems.add(new MenuItemX("Add default HTE credentials", MenuSystem.class.getMethod("mDefaultHTECredentials")));
//            menuItems.add(new MenuItemX("Add new HTE credentials", MenuSystem.class.getMethod("mNewHTECredentials")));
//            menuItems.add(new MenuItemX("Add RoboWash service", MenuSystem.class.getMethod("mAddRoboWashService")));
//            menuItems.add(new MenuItemX("Add RoboAir service", MenuSystem.class.getMethod("mAddRoboAirService")));
//            //menuItems.add(new MenuItemX("Initialise producer", MenuSystem.class.getMethod("mBroadcast);
//            menuItems.add(new MenuItemX("Start service broadcast", MenuSystem.class.getMethod("mStartBroadcast")));
//            menuItems.add(new MenuItemX("Stop broadcast", MenuSystem.class.getMethod("mStopBroadcast")));
//            menuItems.add(new MenuItemX("Producer status", MenuSystem.class.getMethod("mProducerStatus")));
            menuItems.add(new MenuItemX("-------------------- CONSUMER --------------------", MenuSystem.class.getMethod("mInvalidSelection")));
//            menuItems.add(new MenuItemX("Create default consumer", MenuSystem.class.getMethod("mDefaultConsumer")));
//            menuItems.add(new MenuItemX("Create new consumer", MenuSystem.class.getMethod("mNewConsumer")));
//            menuItems.add(new MenuItemX("Scan services", MenuSystem.class.getMethod("mScanService")));
//            menuItems.add(new MenuItemX("Create default HCE credential", MenuSystem.class.getMethod("mDefaultHCECredential")));
//            menuItems.add(new MenuItemX("Create new HCE credential", MenuSystem.class.getMethod("mNewHCECredential")));
//            menuItems.add(new MenuItemX("Discover services", MenuSystem.class.getMethod("mDiscoverSvcs")));
//            menuItems.add(new MenuItemX("Get service prices", MenuSystem.class.getMethod("mGetSvcPrices")));
//            menuItems.add(new MenuItemX("Select service", MenuSystem.class.getMethod("mSelectService")));
//            menuItems.add(new MenuItemX("Make payment", MenuSystem.class.getMethod("mMakePayment")));
//            menuItems.add(new MenuItemX("Consumer status", MenuSystem.class.getMethod("mConsumerStatus")));
            menuItems.add(new MenuItemX("--------------------------------------------------", MenuSystem.class.getMethod("mInvalidSelection")));
//            menuItems.add(new MenuItemX("Exit", MenuSystem.class.getMethod("mQuit")));
        } catch (NoSuchMethodException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "Could not find method", ex);
        } catch (SecurityException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "Security issue with method call", ex);
        }
	renderMenu();        
    }
    
    public void renderMenu() {
	System.out.println("----------------------------- Worldpay Within SDK Client ----------------------------");

	for(int i = 0; i<menuItems.size(); i++) {

		System.out.printf("%d - %s\n", i, ((MenuItemX)menuItems.get(i)).getLabel());
	
        }

	System.out.println("-------------------------------------------------------------------------------------");

	System.out.print("Please select choice: ");
        Scanner scanner = new Scanner(System.in);
	String input = scanner.next();

//	if _, err := fmt.Scanln(&input); err != nil {
//
//		fmt.Printf("Selection error: %q\n", err.Error())
//		renderMenu()
//		return
//	}
//
//	inputInt, err := strconv.Atoi(input)
//
//	if err != nil {
//		fmt.Println("Please type an integer choice!")
//		renderMenu()
//		return
//	}
        
        int inputInt = new Integer(input).intValue();

	if(inputInt >= menuItems.size()) {
		System.out.println("Index out of bounds!");
		renderMenu();
		return;
	}

        
        MenuReturnStruct rc;
        try {
            rc = (MenuReturnStruct)((MenuItemX)menuItems.get(inputInt)).getAction().invoke(this);
        
            if(rc.getReturnCode() != 1) {

                System.out.println(rc.getMsg());
                renderMenu();
            }        
        
        
        } catch (IllegalArgumentException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "The arguments provided to the method weren't allowed", ex);
        } catch (InvocationTargetException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "The method call was invoked on somethign that it can't be invoked on", ex);
        } catch (IllegalAccessException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "Were unable to access this method at this point", ex);
        }

       


    }
    
}
