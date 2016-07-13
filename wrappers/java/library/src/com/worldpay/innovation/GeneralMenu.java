/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

package com.worldpay.innovation;

import com.worldpay.innovation.wpwithin.rpc.WPWithin;
import com.worldpay.innovation.wpwithin.rpc.types.PaymentResponse;
import com.worldpay.innovation.wpwithin.rpc.types.Price;
import com.worldpay.innovation.wpwithin.rpc.types.Service;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceDetails;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceMessage;
import com.worldpay.innovation.wpwithin.rpc.types.TotalPriceResponse;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;
import java.util.logging.ConsoleHandler;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.apache.thrift.TException;


/**
 *
 * @author worldpay
 */
public class GeneralMenu extends MenuBase {

    private final WPWithin.Client sdk;
        
    public GeneralMenu(WPWithin.Client _client) {
        super(_client);
        this.sdk = _client;
        setupLog();
    } 

    protected static final Logger log = Logger.getLogger( GeneralMenu.class.getName() ); 
    public void setupLog() {
        log.setLevel(Level.FINE);
        ConsoleHandler handler = new ConsoleHandler();
        handler.setLevel(Level.FINE);
        log.addHandler(handler);
    }
        
    
    public MenuReturnStruct mGetDeviceInfo() {

            if(this.sdk == null) {
                return new MenuReturnStruct("ERR_DEVICE_NOT_INITIALISED", 0);
            }

            try {
                System.out.println("Uid of device: " + this.sdk.getDevice().getUid() + "\n");

                System.out.println("Name of device: " + this.sdk.getDevice().getName() + "\n");
                System.out.println("Description: " + this.sdk.getDevice().getDescription() + "\n");
                System.out.println("Services: \n");

                log.fine(this.sdk.getDevice().getServicesSize() + ": services available");
                
                for(int i=0; i < this.sdk.getDevice().getServicesSize(); i++) {

                        Service service = this.sdk.getDevice().getServices().get(i);
                        
                        if(service != null) {

                            try {
                                System.out.println("   " + i + ": Id:" + service.getId() + " Name:" + service.getName() + " Description:" + service.getDescription() + "\n");
                            } catch(NullPointerException npe) {
                                npe.printStackTrace();
                                return new MenuReturnStruct("Null pointer exception amongst the service info", 0);
                            }
                            System.out.println("   Prices: \n");


                            for(int j=0; j< service.getPricesSize(); j++) {
                                Price price = service.getPrices().get(j);
                                System.out.println("      " + j + ": ServiceID: " + service.getId() + " ID:" + price.getId() + " Description:" + price.getDescription() + " PricePerUnit:" + price.getPricePerUnit() + " UnitID:" + price.getUnitDescription() + " UnitDescription:%s\n");
                            }

                        } else {
                            log.fine(i + ": service not configured (was null)");
                        }
                }

                System.out.println("IPv4Address: " + this.sdk.getDevice().getIpv4Address() + "\n");
                System.out.println("CurrencyCode: " + this.sdk.getDevice().getCurrencyCode() + "\n");

                return new MenuReturnStruct(null, 0);
                        
            } catch (TException ex) {
                Logger.getLogger(GeneralMenu.class.getName()).log(Level.SEVERE, "sdk client call failed", ex);
                return new MenuReturnStruct("sdk client call failed", 1);
            }                        
    }    

//    public MenuReturnStruct mInitDefaultDevice() {
//
//            //_sdk, err := wpwithin.Initialise(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION)
//            
//            try {
//                super.sdk.setup(DEFAULT_DEVICE_NAME, DEFAULT_DEVICE_DESCRIPTION);
//            } catch(TException e) {
//                return new MenuReturnStruct("SDK setup failed", 1);
//            }
//        
//            return new MenuReturnStruct(null, 0);
//
//    }
//
//    public MenuReturnStruct mInitNewDevice()  {
//
//            System.out.println("Name of device: ");
//            
//            Scanner scanner = new Scanner(System.in);
//            String nameOfDevice = scanner.next();
//            if(null == nameOfDevice || "".equals(nameOfDevice)) {
//                    return new MenuReturnStruct("Name of device not set", 0);
//            }
//
//            System.out.println("Description: ");
//            String description = scanner.next();
//            if(null == description || "".equals(description)) {
//                    return new MenuReturnStruct("Description of device not set", 0);
//            }
//                    
//            try {
//                sdk.setup(nameOfDevice, description);
//            } catch(TException e) {
//                return new MenuReturnStruct("Setup of device unsucessful", 0);
//            }
//            
//            return new MenuReturnStruct(null, 0);
//            
//    }

    
    public MenuReturnStruct mCarWashDemoConsumer() {

	log.fine("testDiscoveryAndNegotiation");

	MenuReturnStruct rc = mInitDefaultDevice();
        if(rc.getMsg() != null) {
            return rc;
        }    

        ConsumerMenu consumerMenu = new ConsumerMenu(sdk);
        
        rc = consumerMenu.mDefaultHCECredential();
        if(rc.getMsg() != null) {
            return rc;
        }    
        
	if(sdk == null) {
		return new MenuReturnStruct("ERR_DEVICE_NOT_INITIALISED", 0);
	}

        HashSet services;
	log.fine("pre scan for services");
	try {
            services = (HashSet)sdk.serviceDiscovery(20000);
        } catch(TException e) {
            return new MenuReturnStruct("Something failed during service discovery", 0);
        }
	log.fine("end scan for services");


	if(services.size() >= 1) {

		ServiceMessage svc = (ServiceMessage)(services.toArray()[0]);

		System.out.println("# Service:: (" + svc.getHostname() + ":" + svc.getPortNumber() + "/" + svc.getUrlPrefix() + ") - " + svc.getDeviceDescription());

		log.fine("Init consumer");
                        
                try {        
                    sdk.initConsumer("http://", svc.getHostname(), svc.getPortNumber(), svc.getUrlPrefix(), svc.getServerId());
                } catch(TException e) {
                    return new MenuReturnStruct("Faild to init the consumer", 0);
                }

		log.fine("Client created..");
                Set<ServiceDetails> serviceDetails;
                
                try {
                    serviceDetails = sdk.requestServices();
                } catch(TException e) {
                    return new MenuReturnStruct("Failed to request services", 0);
                }

		if(serviceDetails.size() >= 1) {

			ServiceDetails svcDetails = serviceDetails.toArray(new ServiceDetails[]{})[0];

			System.out.println(svcDetails.getServiceId() + " - " + svcDetails.getServiceDescription() + "\n");

                        Set<Price> prices;
                        try {
        			        prices = sdk.getServicePrices(svcDetails.getServiceId());
                            
                        } catch(TException e) {
                            return new MenuReturnStruct("Failed to get prices", 0);
                        }

			System.out.println("------- Prices -------\n");
			if(prices.size() >= 1) {

				Price price = prices.toArray(new Price[]{})[0];

				System.out.println("(" + price.getId() + ") " + price.getDescription() + " @ " + price.getPricePerUnit() + ", " + price.getUnitDescription() + " (Unit id = " + price.getUnitId() +")\n");

                                TotalPriceResponse tpr;
                                try {

                                    tpr = sdk.selectService(svcDetails.getServiceId(), 2, price.getId());
                                    
                                } catch(TException e) {
                                    return new MenuReturnStruct("Failed to get total price response", 0);
                                }
                                
                                System.out.println("#Begin Request#");
				System.out.println("ServerID: " + tpr.getServerId() + "\n");
				System.out.println("PriceID = " + tpr.getPriceId() + " - " + tpr.getUnitsToSupply() + " units = " + tpr.getTotalPrice() + "\n");
				System.out.println("ClientID: " + tpr.getClientId() + ", MerchantClientKey: " + tpr.getMerchantClientKey() + ", PaymentRef: " + tpr.getPaymentReferenceId() + "\n");
				System.out.println("#End Request#");

				log.log(Level.FINE, "Making payment of {0}\n", tpr.getTotalPrice());

                                PaymentResponse payResp;
                                try {
                                    payResp = sdk.makePayment(tpr);
                                } catch(TException e) {
                                    return new MenuReturnStruct("Failed to make the payment unfortunately", 0);
                                }
                                
				System.out.println("Payment of " + payResp.getTotalPaid() + " made successfully\n");

				System.out.println("Service delivery token: " + payResp.getServiceDeliveryToken() + "\n");

			}
		}
	}
	return new MenuReturnStruct(null, 0);
}

    
public MenuReturnStruct mCarWashDemoProducer() {

    ProducerMenu producerMenu = new ProducerMenu(this.sdk);
    return producerMenu.mCarWashDemoProducer();
    
}
            
            
    /*
func mResetSessionState() (int, error) {

	sdk = nil

	return 0, nil
}
*/
    
    /*
func mLoadConfig() (int, error) {

	// Ask user for path to config file
	// (And password if secured)

	return 0, errors.New("Not implemented yet..")
}
*/
    
    /*
func mReadConfig() (int, error) {

	// Print out loaded configuration
	// Print out the path to file that was loaded (Need to keep reference during load stage)

	return 0, errors.New("Not implemented yet..")
}
*/
    /*
func mStartRPCService() (int, error) {

	config := rpc.Configuration{
		Protocol:   "binary",
		Framed:     false,
		Buffered:   false,
		Host:       "127.0.0.1",
		Port:       9091,
		Secure:     false,
		BufferSize: 8192,
	}

	rpc, err := rpc.NewService(config, sdk)

	if err != nil {

		return 0, err
	}

	if err := rpc.Start(); err != nil {

		return 0, err
	}

	return 0, nil
}
*/


}
