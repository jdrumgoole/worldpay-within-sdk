package com.worldpay.innovation.wpwithin.thriftadapter;

import com.worldpay.innovation.wpwithin.rpc.types.PricePerUnit;
import com.worldpay.innovation.wpwithin.types.WWPricePerUnit;

/**
 * Created by conor on 13/08/2016.
 */
public class PricePerUnitAdapter {

    public static PricePerUnit convertWWPricePerUnit(WWPricePerUnit wwPricePerUnit) {

        PricePerUnit pricePerUnit = new PricePerUnit();

        pricePerUnit.setAmount(wwPricePerUnit.getAmount());
        pricePerUnit.setCurrencyCode(wwPricePerUnit.getCurrentCode());

        return pricePerUnit;

    }

    public static WWPricePerUnit convertPricePerUnit(PricePerUnit pricePerUnit) {

        WWPricePerUnit wwPricePerUnit = new WWPricePerUnit();
        wwPricePerUnit.setAmount(pricePerUnit.getAmount());
        wwPricePerUnit.setCurrentCode(pricePerUnit.getCurrencyCode());
        return wwPricePerUnit;

    }
}
