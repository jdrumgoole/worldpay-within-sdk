package com.worldpay.innovation.wpwithin.thriftadapter;

import com.worldpay.innovation.wpwithin.rpc.types.TotalPriceResponse;
import com.worldpay.innovation.wpwithin.types.WWTotalPriceResponse;

/**
 * Created by conor on 13/08/2016.
 */
public class TotalPriceResponseAdapter {

    public static WWTotalPriceResponse convertTotalPriceResponse(TotalPriceResponse totalPriceResponse) {

        WWTotalPriceResponse wwTotalPriceResponse = new WWTotalPriceResponse();
        wwTotalPriceResponse.setClientId(totalPriceResponse.getClientId());
        wwTotalPriceResponse.setMerchantClientKey(totalPriceResponse.getMerchantClientKey());
        wwTotalPriceResponse.setPaymentReferenceId(totalPriceResponse.getPaymentReferenceId());
        wwTotalPriceResponse.setPriceId(totalPriceResponse.getPriceId());
        wwTotalPriceResponse.setServerId(totalPriceResponse.getServerId());
        wwTotalPriceResponse.setUnitsToSupply(totalPriceResponse.getUnitsToSupply());

        return wwTotalPriceResponse;

    }

    public static TotalPriceResponse convertWWTotalPriceResponse(WWTotalPriceResponse wwTotalPriceResponse) {

        TotalPriceResponse totalPriceResponse = new TotalPriceResponse();
        totalPriceResponse.setClientId(wwTotalPriceResponse.getClientId());
        totalPriceResponse.setMerchantClientKey(wwTotalPriceResponse.getMerchantClientKey());
        totalPriceResponse.setPaymentReferenceId(wwTotalPriceResponse.getPaymentReferenceId());
        totalPriceResponse.setPriceId(wwTotalPriceResponse.getPriceId());
        totalPriceResponse.setServerId(wwTotalPriceResponse.getServerId());
        totalPriceResponse.setUnitsToSupply(wwTotalPriceResponse.getUnitsToSupply());

        return totalPriceResponse;

    }
}
