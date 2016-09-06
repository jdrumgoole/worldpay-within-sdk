#!/usr/bin/python
import launcher
import time

def main():
    proc=launcher.runRPCAgent("./rpc-agent/", "9092")
    print("FIRST TASK")
    time.sleep(20)
    print("SECOND TASK")
    proc.kill()
    print("KILLED")
    time.sleep(20)


if __name__ == "__main__":
    main()
