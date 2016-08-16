package com.worldpay.innovation.wpwithin.thriftadapter;

import com.worldpay.innovation.wpwithin.rpc.types.Device;
import com.worldpay.innovation.wpwithin.types.WWDevice;

/**
 * Created by conor on 13/08/2016.
 */
public class DeviceAdapter {

    public static WWDevice convertDevice(Device device) {
        WWDevice wwDevice = new WWDevice();
        wwDevice.setCurrencyCode(device.getCurrencyCode());
        wwDevice.setDescription(device.getDescription());
        wwDevice.setIpv4Address(device.getIpv4Address());
        wwDevice.setName(device.getName());
        wwDevice.setServices(ServiceAdapter.convertServices(device.getServices()));
        wwDevice.setUid(device.getUid());
        return wwDevice;
    }
}
