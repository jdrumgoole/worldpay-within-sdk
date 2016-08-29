package com.worldpay.innovation.wpwithin.rpc.launcher;

/**
 * Listen for events related to launched process
 */
public interface Listener {

    // Called when application exits with exitCode being the exit code of that application. By convention a zero exit code
    // means success and non zero means error
    void onApplicationExit(int exitCode, String stdOutput, String errOutput);
}
