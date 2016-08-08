package com.worldpay.innovation;

import com.worldpay.innovation.types.*;
import com.worldpay.innovation.wpwithin.rpc.WPWithinCallback;
import com.worldpay.innovation.wpwithin.rpc.types.*;
import com.worldpay.innovation.wpwithin.rpc.types.ServiceDeliveryToken;
import org.apache.thrift.TException;

/**
 * Created by conor on 29/07/2016.
 */
public class WPWithinHandler implements WPWithinCallback.Iface {

    private Callbacks callbacks;

    public WPWithinHandler(Callbacks callbacks) {

        this.callbacks = callbacks;
    }

    @Override
    public void beginServiceDelivery(String clientId, ServiceDeliveryToken serviceDeliveryToken, int unitsToSupply) throws com.worldpay.innovation.wpwithin.rpc.types.Error, TException {

        System.out.printf("beginServiceDelivery(%s, %s, %d)\n", clientId, serviceDeliveryToken.getKey(), unitsToSupply);

        com.worldpay.innovation.types.ServiceDeliveryToken sdt = new com.worldpay.innovation.types.ServiceDeliveryToken(
                serviceDeliveryToken.key,
                serviceDeliveryToken.getIssued(),
                serviceDeliveryToken.getExpiry(),
                serviceDeliveryToken.isRefundOnExpiry(),
                serviceDeliveryToken.getSignature());

        callbacks.beginServiceDelivery(clientId, serviceDeliveryToken, unitsToSupply);
    }

    @Override
    public void endServiceDelivery(String clientId, ServiceDeliveryToken serviceDeliveryToken, int unitsReceived) throws com.worldpay.innovation.wpwithin.rpc.types.Error, TException {

        System.out.printf("endServiceDelivery(%s, %s, %d)\n", clientId, serviceDeliveryToken.getKey(), unitsReceived);

        com.worldpay.innovation.types.ServiceDeliveryToken sdt = new com.worldpay.innovation.types.ServiceDeliveryToken(
                serviceDeliveryToken.key,
                serviceDeliveryToken.getIssued(),
                serviceDeliveryToken.getExpiry(),
                serviceDeliveryToken.isRefundOnExpiry(),
                serviceDeliveryToken.getSignature());

        callbacks.beginServiceDelivery(clientId, serviceDeliveryToken, unitsReceived);
    }

    public interface Callbacks {

        void beginServiceDelivery(String clientId, ServiceDeliveryToken serviceDeliveryToken, int unitsToSupply) throws com.worldpay.innovation.wpwithin.rpc.types.Error, TException;
        void endServiceDelivery(String clientId, ServiceDeliveryToken serviceDeliveryToken, int unitsReceived) throws com.worldpay.innovation.wpwithin.rpc.types.Error, TException;
    }
}
