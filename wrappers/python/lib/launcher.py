#!/usr/bin/python
## Python Launcher
## Sample ./launcher.py /home "{'linux-32':'rpc-agent-linux-32','linux-64':'rpc-agent-linux-64','win-32':'rpc-agent-win-32.exe','win-64':'rpc-agent-win-64.exe'}"
from subprocess import call
import sys
import ast
import platform

def os_arch():
    """Give the architecture as 32 or 64"""
    if sys.maxsize > 2**32:
        out = '64'
    else:
        out = '32'
    return out;

def os_platform():
    """Give the platform as win, mac or linux"""
    os = platform.system().lower()
    if os == 'windows':
        out = 'win'
    elif os == 'darwin':
        out = 'mac'
    else:
        out = os
    return out;

def runRPCAgent(execPath, port):
    """Run RPC Agent
    
    Args:
        execPath (string): path to directory with rpc agent launchers
        port (integer): port to run RPC agent on
    """
    os = os_platform()
    agent = 'rpc-agent-' + os + '-' + os_arch()
    if os == 'win':
        agent += '.exe'
    call([execPath + agent,'-port='+port)