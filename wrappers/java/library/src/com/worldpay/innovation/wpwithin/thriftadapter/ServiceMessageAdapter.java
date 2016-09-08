package com.worldpay.innovation.wpwithin.thriftadapter;

import com.worldpay.innovation.wpwithin.rpc.types.ServiceMessage;
import com.worldpay.innovation.wpwithin.types.WWServiceMessage;

import java.util.HashSet;
import java.util.Set;

/**
 * Created by conor on 13/08/2016.
 */
public class ServiceMessageAdapter {

    public static WWServiceMessage convertServiceMessage(ServiceMessage svcMsg) {

        WWServiceMessage wwSvcMsg = new WWServiceMessage();
        wwSvcMsg.setDeviceDescription(svcMsg.getDeviceDescription());
        wwSvcMsg.setHostname(svcMsg.getHostname());
        wwSvcMsg.setPortNumber(svcMsg.getPortNumber());
        wwSvcMsg.setServerId(svcMsg.getServerId());
        wwSvcMsg.setUrlPrefix(svcMsg.getUrlPrefix());
        wwSvcMsg.setScheme(svcMsg.getScheme());
        return wwSvcMsg;

    }

    public static Set<WWServiceMessage> convertServiceMessages(Set<ServiceMessage> svcMsgs) {

        Set<WWServiceMessage> wwServiceMessages = new HashSet();

        if (svcMsgs != null && svcMsgs.size() > 0) {

            for (ServiceMessage svcMsg : svcMsgs) {
                WWServiceMessage wwSvcMsg = convertServiceMessage(svcMsg);
                wwServiceMessages.add(wwSvcMsg);
            }
        }
        return wwServiceMessages;

    }
}
