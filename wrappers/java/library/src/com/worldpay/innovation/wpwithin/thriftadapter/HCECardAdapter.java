package com.worldpay.innovation.wpwithin.thriftadapter;

import com.worldpay.innovation.wpwithin.rpc.types.HCECard;
import com.worldpay.innovation.wpwithin.types.WWHCECard;

/**
 * Created by conor on 13/08/2016.
 */
public class HCECardAdapter {

    public static HCECard convertWWHCECard(WWHCECard wwHceCard) {
        HCECard hceCard = new HCECard();

        hceCard.setCardNumber(wwHceCard.getCardNumber());
        hceCard.setCvc(wwHceCard.getCvc());
        hceCard.setExpMonth(wwHceCard.getExpMonth());
        hceCard.setExpYear(wwHceCard.getExpYear());
        hceCard.setFirstName(wwHceCard.getFirstName());
        hceCard.setLastName(wwHceCard.getLastName());
        hceCard.setType(wwHceCard.getType());

        return hceCard;
    }
}
