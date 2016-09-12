package com.worldpay.innovation.wpwithin.rpc.launcher;

import com.worldpay.innovation.wpwithin.WPWithinGeneralException;

import java.io.*;
import java.util.Map;

public class Launcher {

    private Process processHandle;

    private StringBuilder errorOutput;
    private StringBuilder stdOutput;

    private static final String NEW_LINE = System.getProperty("line.separator", "\n");

    public void startProcess(Map<OS, PlatformConfig> launchConfig, final Listener listener) throws WPWithinGeneralException {

        OS hostOS = detectHostOS();
        Architecture hostArch = detectHostArchitecture();


        if (validateConfig(launchConfig, hostOS, hostArch)) {

            String command = launchConfig.get(hostOS).getCommand(hostArch);

            try {

                processHandle = Runtime.getRuntime().exec(command);

                doReadErrOutput();
                doReadStdOutput();

                new Thread(new Runnable() {
                    @Override
                    public void run() {

                        try {

                            int exitCode = processHandle.waitFor();

                            listener.onApplicationExit(exitCode, stdOutput.toString(), errorOutput.toString());

                        } catch (Exception e) {

                            e.printStackTrace();
                        }
                    }
                }).start();

            } catch (IOException ioe) {

                ioe.printStackTrace();

                throw new WPWithinGeneralException("Unable to launch process: " + ioe.getMessage());
            }

        } else {

            throw new WPWithinGeneralException("Invalid launch configuration detected, cannot launch application.");
        }
    }

    public void stopProcess() {

        if (processHandle != null) {

            processHandle.destroy();
        }
    }

    private OS detectHostOS() {

        String hostOS = System.getProperty("os.name");

        if (hostOS == null || hostOS.length() == 0) {

            return OS.UNKNOWN;
        } else if (hostOS.toLowerCase().contains("win")) {

            return OS.WINDOWS;

        } else if (hostOS.toLowerCase().contains("darwin") || hostOS.toLowerCase().contains("mac")) {

            return OS.MAC;
        } else if (hostOS.toLowerCase().contains("linux")) {

            return OS.LINUX;
        } else {

            return OS.UNKNOWN;
        }
    }

    private Architecture detectHostArchitecture() {

        String arch = System.getProperty("os.arch");

        if (arch == null || arch.length() == 0) {

            return Architecture.UNKNOWN;
        } else if (arch.toLowerCase().contains("64")) {

            return Architecture.X86_64;
        } else if (arch.toLowerCase().contains("arm")) {

            return Architecture.ARM;
        } else if (arch.toLowerCase().equals("x86")) {

            return Architecture.IA32;
        } else {

            return Architecture.UNKNOWN;
        }
    }

    private boolean validateConfig(Map<OS, PlatformConfig> launchConfig, OS hostOS, Architecture hostArchitecture) {

        return launchConfig.containsKey(hostOS) && (launchConfig.get(hostOS).getCommand(hostArchitecture) != null && launchConfig.get(hostOS).getCommand(hostArchitecture).length() > 0);
    }

    public String getErrorOutput() {
        return errorOutput.toString();
    }

    public String getStdOutput() {
        return stdOutput.toString();
    }

    private void doReadStdOutput() {

        new Thread(new Runnable() {

            @Override
            public void run() {

                try {

                    stdOutput = new StringBuilder();

                    BufferedReader br = new BufferedReader(new InputStreamReader(processHandle.getInputStream()));

                    String line;
                    while ((line = br.readLine()) != null) {

                        stdOutput.append(line + NEW_LINE);
                    }
                } catch (IOException e) {

                    if(!e.getMessage().toLowerCase().contains("stream closed")) {

                        e.printStackTrace();
                    }
                }

            }
        }).start();
    }

    private void doReadErrOutput() {

        new Thread(new Runnable() {

            @Override
            public void run() {

                try {

                    errorOutput = new StringBuilder();

                    BufferedReader br = new BufferedReader(new InputStreamReader(processHandle.getErrorStream()));

                    String line;
                    while ((line = br.readLine()) != null) {

                        errorOutput.append(line + NEW_LINE);
                    }
                } catch (IOException e) {

                    if(!e.getMessage().toLowerCase().contains("stream closed")) {

                        e.printStackTrace();
                    }
                }

            }
        }).start();
    }
}
