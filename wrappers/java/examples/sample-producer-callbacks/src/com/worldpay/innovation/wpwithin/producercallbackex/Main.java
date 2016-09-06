package com.worldpay.innovation.wpwithin.producercallbackex;


import com.worldpay.innovation.wpwithin.WPWithinGeneralException;
import com.worldpay.innovation.wpwithin.WPWithinWrapper;
import com.worldpay.innovation.wpwithin.WPWithinWrapperImpl;
import com.worldpay.innovation.wpwithin.eventlistener.EventListener;
import com.worldpay.innovation.wpwithin.types.WWPrice;
import com.worldpay.innovation.wpwithin.types.WWPricePerUnit;
import com.worldpay.innovation.wpwithin.types.WWService;
import com.worldpay.innovation.wpwithin.types.WWServiceDeliveryToken;

import java.util.HashMap;

public class Main {

    public static void main(String[] args) {

        try {

            System.out.println("WorldpayWithin Sample Producer...");

            WPWithinWrapper wpw = new WPWithinWrapperImpl("127.0.0.1", 9089, wpWithinEventListener, 9095);

            wpw.setup("Producer Example", "Example WorldpayWithin producer");

            WWService svc = new WWService();
            svc.setName("Car charger");
            svc.setDescription("Can charge your hybrid / electric car");
            svc.setId(1);

            WWPrice ccPrice = new WWPrice();
            ccPrice.setId(1);
            ccPrice.setDescription("Kilowatt-hour");
            ccPrice.setUnitDescription("One kilowatt-hour");
            ccPrice.setUnitId(1);
            WWPricePerUnit ppu = new WWPricePerUnit();
            ppu.setAmount(25);
            ppu.setCurrentCode("GBP");
            ccPrice.setPricePerUnit(ppu);

            HashMap<Integer, WWPrice> prices = new HashMap<>(1);
            prices.put(ccPrice.getId(), ccPrice);

            svc.setPrices(prices);

            wpw.addService(svc);

            wpw.initProducer("T_C_03eaa1d3-4642-4079-b030-b543ee04b5af", "T_S_f50ecb46-ca82-44a7-9c40-421818af5996");

            wpw.startServiceBroadcast(1000 * 9999);

        } catch (WPWithinGeneralException e) {

            e.printStackTrace();
        }
    }

    private static EventListener wpWithinEventListener = new EventListener() {

        @Override
        public void onBeginServiceDelivery(int serviceID, WWServiceDeliveryToken wwServiceDeliveryToken, int unitsToSupply) throws WPWithinGeneralException {

            try {
                System.out.println("event from core - onBeginServiceDelivery()");
                System.out.printf("ServiceID: %d\n", serviceID);
                System.out.printf("UnitsToSupply: %d\n", unitsToSupply);
                System.out.printf("SDT.Key: %s\n", wwServiceDeliveryToken.getKey());
                System.out.printf("SDT.Expiry: %s\n", wwServiceDeliveryToken.getExpiry());
                System.out.printf("SDT.Issued: %s\n", wwServiceDeliveryToken.getIssued());
                System.out.printf("SDT.Signature: %s\n", wwServiceDeliveryToken.getSignature());
                System.out.printf("SDT.RefundOnExpiry: %b\n", wwServiceDeliveryToken.isRefundOnExpiry());
            } catch (Exception e) {

                e.printStackTrace();
            }
        }

        @Override
        public void onEndServiceDelivery(int serviceID, WWServiceDeliveryToken wwServiceDeliveryToken, int unitsReceived) throws WPWithinGeneralException {

            try {

                System.out.println("event from core - onEndServiceDelivery()");
                System.out.printf("ServiceID: %d\n", serviceID);
                System.out.printf("UnitsReceived: %d\n", unitsReceived);
                System.out.printf("SDT.Key: %s\n", wwServiceDeliveryToken.getKey());
                System.out.printf("SDT.Expiry: %s\n", wwServiceDeliveryToken.getExpiry());
                System.out.printf("SDT.Issued: %s\n", wwServiceDeliveryToken.getIssued());
                System.out.printf("SDT.Signature: %s\n", wwServiceDeliveryToken.getSignature());
                System.out.printf("SDT.RefundOnExpiry: %b\n", wwServiceDeliveryToken.isRefundOnExpiry());
            } catch (Exception e) {

                e.printStackTrace();
            }
        }
    };
}
