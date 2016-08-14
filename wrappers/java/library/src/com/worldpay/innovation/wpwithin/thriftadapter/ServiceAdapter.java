package com.worldpay.innovation.wpwithin.thriftadapter;

import com.worldpay.innovation.wpwithin.rpc.types.Service;
import com.worldpay.innovation.wpwithin.types.WWService;

import java.util.HashMap;
import java.util.Map;
import java.util.Set;

/**
 * Created by conor on 13/08/2016.
 */
public class ServiceAdapter {

    public static HashMap<Integer, WWService> convertServices(Map<Integer, Service> services) {
        HashMap<Integer, WWService> wwServices = new HashMap();

        Set<Integer> keyset = services.keySet();

        for(Integer key : keyset) {

            wwServices.put(key, convertService(services.get(key)));

        }

        return wwServices;
    }

    public static WWService convertService(Service service) {
        WWService wwService = new WWService();
        wwService.setDescription(service.getDescription());
        wwService.setId(service.getId());
        wwService.setName(service.getName());
        wwService.setPrices(PriceAdapter.convertPrices(service.getPrices()));
        return wwService;
    }

    public static Service convertWWService(WWService wwService) {

        Service service = new Service();

        service.setDescription(wwService.getDescription());
        service.setId(wwService.getId());
        service.setName(wwService.getName());
        service.setPrices(PriceAdapter.convertWWPrices(wwService.getPrices()));

        return service;

    }
}
