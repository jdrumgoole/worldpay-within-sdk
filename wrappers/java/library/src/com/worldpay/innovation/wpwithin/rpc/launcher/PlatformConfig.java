package com.worldpay.innovation.wpwithin.rpc.launcher;

import java.util.HashMap;
import java.util.Map;

/**
 * Details of how to config for a specific platform
 */
public class PlatformConfig {

    private Map<Architecture, String> commands;

    public PlatformConfig() {

        commands = new HashMap<Architecture, String>(3);
    }

    public void setCommand(Architecture architecture, String command) {

        commands.put(architecture, command);
    }

    public String getCommand(Architecture architecture) {

        return commands.get(architecture);
    }
}
