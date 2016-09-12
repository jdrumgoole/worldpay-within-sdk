package com.worldpay.innovation.wpwithin.eventlistener;

import com.worldpay.innovation.wpwithin.WPWithinGeneralException;
import com.worldpay.innovation.wpwithin.rpc.WPWithinCallback;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceDeliveryToken;
import com.worldpay.innovation.wpwithin.types.WWServiceDeliveryToken;
import org.apache.thrift.TException;

class EventHandler implements WPWithinCallback.Iface {

    private EventListener listener;

    public EventHandler(EventListener listener) {

        this.listener = listener;
    }

    @Override
    public void beginServiceDelivery(int serviceId, ServiceDeliveryToken serviceDeliveryToken, int unitsToSupply) throws TException {

        try {

            WWServiceDeliveryToken sdt = new WWServiceDeliveryToken();
            sdt.setKey(serviceDeliveryToken.getKey());
            sdt.setIssued(serviceDeliveryToken.getIssued());
            sdt.setExpiry(serviceDeliveryToken.getExpiry());
            sdt.setRefundOnExpiry(serviceDeliveryToken.isRefundOnExpiry());
            sdt.setSignature(serviceDeliveryToken.getSignature());

            listener.onBeginServiceDelivery(serviceId, sdt, unitsToSupply);

        } catch (WPWithinGeneralException e) {

            e.printStackTrace();

            throw new TException(e.getMessage(), e.getCause());
        }
    }

    @Override
    public void endServiceDelivery(int serviceId, ServiceDeliveryToken serviceDeliveryToken, int unitsReceived) throws TException {

        try {

            WWServiceDeliveryToken sdt = new WWServiceDeliveryToken();
            sdt.setKey(serviceDeliveryToken.getKey());
            sdt.setIssued(serviceDeliveryToken.getIssued());
            sdt.setExpiry(serviceDeliveryToken.getExpiry());
            sdt.setRefundOnExpiry(serviceDeliveryToken.isRefundOnExpiry());
            sdt.setSignature(serviceDeliveryToken.getSignature());

            listener.onEndServiceDelivery(serviceId, sdt, unitsReceived);

        } catch (WPWithinGeneralException e) {

            e.printStackTrace();

            throw new TException(e.getMessage(), e.getCause());
        }
    }
}
