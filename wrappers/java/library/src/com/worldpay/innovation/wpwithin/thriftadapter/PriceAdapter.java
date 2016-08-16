package com.worldpay.innovation.wpwithin.thriftadapter;

import com.worldpay.innovation.wpwithin.rpc.types.Price;
import com.worldpay.innovation.wpwithin.rpc.types.PricePerUnit;
import com.worldpay.innovation.wpwithin.types.WWPrice;
import com.worldpay.innovation.wpwithin.types.WWPricePerUnit;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

/**
 * Created by conor on 13/08/2016.
 */
public class PriceAdapter {

    public static HashMap<Integer, WWPrice> convertPrices(Map<Integer, Price> prices) {
        HashMap<Integer, WWPrice> wwPrices = new HashMap();

        Set<Integer> keyset = prices.keySet();

        for(Integer key : keyset) {

            wwPrices.put(key, convertPrice(prices.get(key)));

        }

        return wwPrices;
    }

    public static WWPrice convertPrice(Price price) {
        WWPrice wwPrice = new WWPrice();
        wwPrice.setDescription(price.getDescription());
        wwPrice.setId(price.getId());
        wwPrice.setPricePerUnit(PricePerUnitAdapter.convertPricePerUnit(price.getPricePerUnit()));
        wwPrice.setUnitDescription(price.getUnitDescription());
        wwPrice.setUnitId(price.getUnitId());
        return wwPrice;
    }

    public static HashMap<Integer, Price> convertWWPrices(HashMap<Integer, WWPrice> wwPrices) {
        HashMap<Integer, Price> prices = new HashMap();

        Set<Integer> keySet = wwPrices.keySet();
        if(keySet != null && keySet.size() > 0) {
            for(Integer i : keySet) {
                WWPrice wwPrice = wwPrices.get(i);
                Price price = new Price();
                price.setDescription(wwPrice.getDescription());
                price.setId(wwPrice.getId());
                price.setUnitDescription(wwPrice.getUnitDescription());
                price.setUnitId(wwPrice.getUnitId());
                price.setPricePerUnit(PricePerUnitAdapter.convertWWPricePerUnit(wwPrice.getPricePerUnit()));

                prices.put(i, price);
            }
        }

        return prices;
    }

    public static Set<WWPrice> convertServicePrices(Set<Price> prices) {
        Set<WWPrice> wwPrices = new HashSet();

        if (prices != null && prices.size() > 0) {

            for (Price price : prices) {

                WWPrice wwPrice = new WWPrice();
                wwPrice.setDescription(price.getDescription());
                wwPrice.setId(price.getId());
                wwPrice.setPricePerUnit(PricePerUnitAdapter.convertPricePerUnit(price.getPricePerUnit()));
                wwPrice.setUnitDescription(price.getUnitDescription());
                wwPrice.setUnitId(price.getUnitId());
                wwPrices.add(wwPrice);

            }
        }
        return wwPrices;
    }
}
