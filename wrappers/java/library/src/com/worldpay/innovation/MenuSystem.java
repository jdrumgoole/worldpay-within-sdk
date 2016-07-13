/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.logging.ConsoleHandler;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 *
 * @author worldpay
 */
public class MenuSystem extends MenuBase {


    private ArrayList menuItems;
    private GeneralMenu generalMenu;
    private ProducerMenu producerMenu;
    private ConsumerMenu consumerMenu;

    public MenuSystem() {
        super();
    } 
        
    public MenuReturnStruct mInvalidSelection() {
        return new MenuReturnStruct("*** Invalid menu selection - please choose another item ***", 0);
    }

    public void doUI(WPWithin.Client client) {
        try {
            menuItems = new ArrayList();

            generalMenu = new GeneralMenu(client);
            producerMenu = new ProducerMenu(client);
            consumerMenu = new ConsumerMenu(client);

            menuItems.add(new MenuItemX("-------------------- GENERAL  --------------------", MenuSystem.class.getMethod("mInvalidSelection")));
            menuItems.add(new MenuItemX("Init default device", GeneralMenu.class.getMethod("mInitDefaultDevice")));
//            menuItems.add(new MenuItemX("Start RPC Service", MenuSystem.class.getMethod("mStartRPCService"))); // doesn't make sense...?
            menuItems.add(new MenuItemX("Init new device", GeneralMenu.class.getMethod("mInitNewDevice")));
            menuItems.add(new MenuItemX("Get device info", GeneralMenu.class.getMethod("mGetDeviceInfo")));
            menuItems.add(new MenuItemX("Sample demo, car wash (Producer)", GeneralMenu.class.getMethod("mCarWashDemoProducer")));
            menuItems.add(new MenuItemX("Sample demo, car wash (Consumer)", GeneralMenu.class.getMethod("mCarWashDemoConsumer")));
//            menuItems.add(new MenuItemX("Reset session", MenuSystem.class.getMethod("mResetSessionState")));
//            menuItems.add(new MenuItemX("Load configuration", MenuSystem.class.getMethod("mLoadConfig")));
//            menuItems.add(new MenuItemX("Read loaded configuration", MenuSystem.class.getMethod("mReadConfig")));
            menuItems.add(new MenuItemX("-------------------- PRODUCER --------------------", MenuSystem.class.getMethod("mInvalidSelection")));
            menuItems.add(new MenuItemX("Create default producer", ProducerMenu.class.getMethod("mDefaultProducer")));
            menuItems.add(new MenuItemX("Create new producer", ProducerMenu.class.getMethod("mNewProducer")));
            menuItems.add(new MenuItemX("Add default HTE credentials", ProducerMenu.class.getMethod("mDefaultHTECredentials")));
            menuItems.add(new MenuItemX("Add new HTE credentials", ProducerMenu.class.getMethod("mNewHTECredentials")));
            menuItems.add(new MenuItemX("Add RoboWash service", ProducerMenu.class.getMethod("mAddRoboWashService")));
            menuItems.add(new MenuItemX("Add RoboAir service", ProducerMenu.class.getMethod("mAddRoboAirService")));
            menuItems.add(new MenuItemX("Initialise producer", ProducerMenu.class.getMethod("mBroadcast")));
            menuItems.add(new MenuItemX("Start service broadcast", ProducerMenu.class.getMethod("mStartBroadcast")));
            menuItems.add(new MenuItemX("Stop broadcast", ProducerMenu.class.getMethod("mStopBroadcast")));
            menuItems.add(new MenuItemX("Producer status", ProducerMenu.class.getMethod("mProducerStatus")));
            menuItems.add(new MenuItemX("-------------------- CONSUMER --------------------", MenuSystem.class.getMethod("mInvalidSelection")));
            menuItems.add(new MenuItemX("Create default consumer", ConsumerMenu.class.getMethod("mDefaultConsumer")));
            menuItems.add(new MenuItemX("Create new consumer", ConsumerMenu.class.getMethod("mNewConsumer")));
            menuItems.add(new MenuItemX("Scan services", ConsumerMenu.class.getMethod("mScanService")));
            menuItems.add(new MenuItemX("Create default HCE credential", ConsumerMenu.class.getMethod("mDefaultHCECredential")));
            menuItems.add(new MenuItemX("Create new HCE credential", ConsumerMenu.class.getMethod("mNewHCECredential")));
            menuItems.add(new MenuItemX("Discover services", ConsumerMenu.class.getMethod("mDiscoverSvcs")));
            menuItems.add(new MenuItemX("Get service prices", ConsumerMenu.class.getMethod("mGetSvcPrices")));
            menuItems.add(new MenuItemX("Select service", ConsumerMenu.class.getMethod("mSelectService")));
            menuItems.add(new MenuItemX("Make payment", ConsumerMenu.class.getMethod("mMakePayment")));
            menuItems.add(new MenuItemX("Consumer status", ConsumerMenu.class.getMethod("mConsumerStatus")));
            menuItems.add(new MenuItemX("--------------------------------------------------", MenuSystem.class.getMethod("mInvalidSelection")));
            menuItems.add(new MenuItemX("Exit", MenuSystem.class.getMethod("mQuit")));
        } catch (NoSuchMethodException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "Could not find method", ex);
        } catch (SecurityException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "Security issue with method call", ex);
        }
        renderMenu();
    }

    public void renderMenu() {
        System.out.println("----------------------------- Worldpay Within SDK Client ----------------------------");

        for (int i = 0; i < menuItems.size(); i++) {

            System.out.printf("%d - %s\n", i, ((MenuItemX) menuItems.get(i)).getLabel());

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

        if (inputInt >= menuItems.size()) {
            System.out.println("Index out of bounds!");
            renderMenu();
            return;
        }

        MenuReturnStruct rc;
        try {
            //rc = (MenuReturnStruct)((MenuItemX)menuItems.get(inputInt)).getAction().invoke(this);
            Method action = ((MenuItemX) menuItems.get(inputInt)).getAction();

            if ((GeneralMenu.class.getMethod("mInitDefaultDevice")).equals(action) || (GeneralMenu.class.getMethod("mInitNewDevice")).equals(action) || (GeneralMenu.class.getMethod("mGetDeviceInfo").equals(action) || (GeneralMenu.class.getMethod("mCarWashDemoConsumer")).equals(action)) || (GeneralMenu.class.getMethod("mCarWashDemoProducer").equals(action))) {
                log.log(Level.INFO, "Inialising: {0}", action.toString());
                rc = (MenuReturnStruct) action.invoke(this.generalMenu);
            } else if ((ProducerMenu.class.getMethod("mDefaultProducer")).equals(action)
                    || (ProducerMenu.class.getMethod("mNewProducer")).equals(action)
                    || (ProducerMenu.class.getMethod("mDefaultHTECredentials")).equals(action)
                    || (ProducerMenu.class.getMethod("mNewHTECredentials")).equals(action)
                    || (ProducerMenu.class.getMethod("mAddRoboWashService")).equals(action)
                    || (ProducerMenu.class.getMethod("mAddRoboAirService")).equals(action)
                    || (ProducerMenu.class.getMethod("mBroadcast")).equals(action)
                    || (ProducerMenu.class.getMethod("mStartBroadcast")).equals(action)
                    || (ProducerMenu.class.getMethod("mStopBroadcast")).equals(action)
                    || (ProducerMenu.class.getMethod("mProducerStatus")).equals(action)) {
                log.log(Level.INFO, "Inialising: {0}", action.toString());
                rc = (MenuReturnStruct) action.invoke(this.producerMenu);
            } else if (ConsumerMenu.class.getMethod("mDefaultConsumer").equals(action)
                    || ConsumerMenu.class.getMethod("mNewConsumer").equals(action)
                    || ConsumerMenu.class.getMethod("mScanService").equals(action)
                    || ConsumerMenu.class.getMethod("mDefaultHCECredential").equals(action)
                    || ConsumerMenu.class.getMethod("mNewHCECredential").equals(action)
                    || ConsumerMenu.class.getMethod("mDiscoverSvcs").equals(action)
                    || ConsumerMenu.class.getMethod("mGetSvcPrices").equals(action)
                    || ConsumerMenu.class.getMethod("mSelectService").equals(action)
                    || ConsumerMenu.class.getMethod("mMakePayment").equals(action)
                    || ConsumerMenu.class.getMethod("mConsumerStatus").equals(action)) {
                log.log(Level.INFO, "Inialising: {0}", action.toString());
                rc = (MenuReturnStruct) action.invoke(this.consumerMenu);                
            } else {

                log.log(Level.INFO, "Inialising: {0}", action.toString());
                rc = (MenuReturnStruct) action.invoke(this);
            }

            if (rc.getReturnCode() != 1) {
                System.out.println("Action seemed to have some degree of success");
                renderMenu();
            } else {
                System.out.println(rc.getMsg());
            }

        } catch (IllegalArgumentException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "The arguments provided to the method weren't allowed", ex);
        } catch (InvocationTargetException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "The method call was invoked on somethign that it can't be invoked on", ex);
        } catch (IllegalAccessException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "Were unable to access this method at this point", ex);
        } catch (NoSuchMethodException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "No such method exists in cmoparison", ex);
        } catch (SecurityException ex) {
            Logger.getLogger(MenuSystem.class.getName()).log(Level.SEVERE, "Security exception with comparison", ex);
        }

    }

    public MenuReturnStruct mQuit() {

        System.out.println("");
        System.out.println("Goodbye...");
        return new MenuReturnStruct(null, 1);

    }

}
