/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.worldpay.innovation.wpwithin.test;

import com.worldpay.innovation.wpwithin.WPWithinGeneralException;
import com.worldpay.innovation.wpwithin.WPWithinWrapper;
import com.worldpay.innovation.wpwithin.WPWithinWrapperImpl;
import com.worldpay.innovation.wpwithin.types.WWHCECard;
import com.worldpay.innovation.wpwithin.types.WWPrice;
import com.worldpay.innovation.wpwithin.types.WWServiceDetails;
import com.worldpay.innovation.wpwithin.types.WWServiceMessage;
import java.util.Set;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 *
 * @author worldpay
 */
public class WorldpayWithinWrapperTester {

    private static final Logger log = Logger.getLogger(WorldpayWithinWrapperTester.class.getName());

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {

        // ######### SETUP RPC SERVICES ############
        String host = "127.0.0.1";
        Integer port = 9091;
        WPWithinWrapper sdk = new WPWithinWrapperImpl(host, port, false);

        try {
            //        sdk.setupDefaultDevice();
//        sdk.initiateDefaultProducer();

            sdk.setup("Kevs device", "The device that doeseth belong to kev");
        } catch (WPWithinGeneralException ex) {
            Logger.getLogger(WorldpayWithinWrapperTester.class.getName()).log(Level.SEVERE, "Faield to setup device...", ex);
        }

//        WWHCECard wwHceCard = new WWHCECard();
//        wwHceCard.setCardNumber("555555555555555");
//        wwHceCard.setCvc("1234");
//        wwHceCard.setExpMonth(10);
//        wwHceCard.setExpYear(2017);
//        wwHceCard.setFirstName("Kevin");
//        wwHceCard.setLastName("Gordon");
//        wwHceCard.setType("VISA");
//        try {
//            sdk.initConsumer("http://", host, port, "", "kevserver", wwHceCard);
//        } catch (WPWithinGeneralException ex) {
//            Logger.getLogger(WorldpayWithinWrapper2.class.getName()).log(Level.SEVERE, "Failed to initiate consumer", ex);
//        }
//        String DEFAULT_HTE_MERCHANT_CLIENT_KEY = "T_C_03eaa1d3-4642-4079-b030-b543ee04b5af";
//        String DEFAULT_HTE_MERCHANT_SERVICE_KEY = "T_S_f50ecb46-ca82-44a7-9c40-421818af5996";
//        try {
//            sdk.initProducer(DEFAULT_HTE_MERCHANT_CLIENT_KEY, DEFAULT_HTE_MERCHANT_SERVICE_KEY);
//        } catch (WPWithinGeneralException ex) {
//            Logger.getLogger(WorldpayWithinWrapper2.class.getName()).log(Level.SEVERE, "Failed to initiate producer", ex);
//        }
//        WWService service = new WWService();
//        service.setDescription("Allows an autonomous car pay for car parking");
//        service.setId(1110874);
//        service.setName("Car Parking");
//        
//        WWPrice priceTarriffe1 = new WWPrice();
//        priceTarriffe1.setDescription("Tariffe 1 standard");
//        priceTarriffe1.setId(104);
//        priceTarriffe1.setUnitDescription("Cost per hour");
//        priceTarriffe1.setUnitId(47);
//        WWPricePerUnit pricePerUnit1 = new WWPricePerUnit();
//        pricePerUnit1.setAmount(2);
//        pricePerUnit1.setCurrentCode("GDP");
//        priceTarriffe1.setPricePerUnit(pricePerUnit1);
//        
//        WWPrice priceTarriffe2 = new WWPrice();
//        priceTarriffe2.setDescription("Tariffe 2 off peak");
//        priceTarriffe2.setId(104);
//        priceTarriffe2.setUnitDescription("Cost per hour");
//        priceTarriffe2.setUnitId(47);
//        WWPricePerUnit pricePerUnit2 = new WWPricePerUnit();
//        pricePerUnit2.setAmount(1);
//        pricePerUnit2.setCurrentCode("GDP");
//        priceTarriffe2.setPricePerUnit(pricePerUnit2);
//        
//        WWPrice priceTarriffe3 = new WWPrice();
//        priceTarriffe3.setDescription("Tariffe 3 premium");
//        priceTarriffe3.setId(104);
//        priceTarriffe3.setUnitDescription("Cost per hour");
//        priceTarriffe3.setUnitId(47);
//        WWPricePerUnit pricePerUnit3 = new WWPricePerUnit();
//        pricePerUnit3.setAmount(12);
//        pricePerUnit3.setCurrentCode("GDP");
//        priceTarriffe3.setPricePerUnit(pricePerUnit3);
//        
//        HashMap<Integer, WWPrice> prices = new HashMap();
//        prices.put(1, priceTarriffe1);
//        prices.put(2, priceTarriffe2);
//        prices.put(3, priceTarriffe3);
//        
//        service.setPrices(prices);
//        
//        try {
//            sdk.addService(service);
//        } catch (WPWithinGeneralException ex) {
//            Logger.getLogger(WorldpayWithinWrapper2.class.getName()).log(Level.SEVERE, "Test add service failed", ex);
//        }
//        try {
//            sdk.startServiceBroadcast(68000);
//        } catch (WPWithinGeneralException ex) {
//            Logger.getLogger(WorldpayWithinWrapper2.class.getName()).log(Level.SEVERE, "Test start service broadcast failed", ex);
//        }
        try {
            Set<WWServiceMessage> serviceMsgs = sdk.deviceDiscovery(port);

            for (WWServiceMessage svcMsg : serviceMsgs) {
                System.out.println(svcMsg.getDeviceDescription() + ":" + svcMsg.getHostname() + ":" + svcMsg.getServerId() + ":" + svcMsg.getUrlPrefix());

                WWHCECard wwHceCard = new WWHCECard();
                wwHceCard.setCardNumber("555555555555555");
                wwHceCard.setCvc("1234");
                wwHceCard.setExpMonth(10);
                wwHceCard.setExpYear(2017);
                wwHceCard.setFirstName("Kevin");
                wwHceCard.setLastName("Gordon");
                wwHceCard.setType("VISA");
                try {
                    sdk.initConsumer("http://", svcMsg.getHostname(), svcMsg.getPortNumber(), svcMsg.getUrlPrefix(), svcMsg.getServerId(), wwHceCard);
                } catch (WPWithinGeneralException ex) {
                    Logger.getLogger(WorldpayWithinWrapperTester.class.getName()).log(Level.SEVERE, "Failed to initiate consumer", ex);
                }
                
                Set<WWServiceDetails> svcDetails = sdk.requestServices();

                if (svcDetails != null && svcDetails.size() > 0) {

                    Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.INFO, "Ouputting service Details");

                    WWServiceDetails svcDetail = svcDetails.iterator().next();

                    Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.INFO, svcDetail.getServiceId() + " - " + svcDetail.getServiceDescription() + "\n");
                    Set<WWPrice> pricesReceived = sdk.getServicePrices(svcDetail.getServiceId());

                    if (pricesReceived != null && pricesReceived.size() > 0) {

                        Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.INFO, "Ouputting prices");

                            // Get the first price of the first service for tesitng purposes....
                        WWPrice price = pricesReceived.iterator().next();

                        Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.INFO, price.getDescription() + " - " + price.getId() + " - " + price.getUnitId() + " - " + price.getUnitDescription() + " - " + price.getPricePerUnit().getAmount() + " - " + price.getPricePerUnit().getCurrentCode() + "\n");

                    } else {
                        Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.INFO, "No prices found");
                    }

                } else {
                    Logger.getLogger(WPWithinWrapperImpl.class.getName()).log(Level.INFO, "No Service Details found");
                }
            }
        } catch (WPWithinGeneralException ex) {
            Logger.getLogger(WorldpayWithinWrapperTester.class.getName()).log(Level.SEVERE, "Test device Discovery failed", ex);
        }

    }

}
