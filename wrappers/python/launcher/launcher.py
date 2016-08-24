#!/usr/bin/python
## Python Launcher
## Sample ./launcher.py /home "{'linux-32':'rpc-agent-linux-32','linux-64':'rpc-agent-linux-64','win-32':'rpc-agent-win-32.exe','win-64':'rpc-agent-win-64.exe'}"
from subprocess import call
import sys
import ast
import platform

def os_arch():
    "This gives the architecture as 32 or 64"
    if sys.maxsize > 2**32:
        out = '64'
    else:
        out = '32'
    return out;

def os_platform():
    "This gives the platform as win, mac or linux"
    if platform.system().lower() == 'windows':
        out = 'win'
    elif platform.system().lower() == 'darwin':
        out = 'mac'
    else:
        out = platform.system().lower()
    return out;

def runRPCAgent(execPath,execMapStr,execParam):
    #print 'Number of arguments:', len(sys.argv), 'arguments.'
    #print 'Argument List:', str(sys.argv)
    #execPath=sys.argv[1]
    #print execPath
    #execMapStr=sys.argv[2]
    execMapDict= ast.literal_eval(execMapStr)
    print execPath+execMapDict[os_platform()+'-'+os_arch()]+' '+execParam
    call([execPath+execMapDict[os_platform()+'-'+os_arch()],execParam])

def main():
    print 'Platform: ' + os_platform() +'-'+ os_arch()
    runRPCAgent("../rpc-agent/","{'linux-32':'rpc-agent-linux-32','linux-64':'rpc-agent-linux-64','win-32':'rpc-agent-win-32.exe','win-64':'rpc-agent-win-64.exe','mac-32':'rpc-agent-mac-32','mac-64':'rpc-agent-mac-64'}","-port=9090")

if __name__ == "__main__":
    main()
