# Dev Client
The dev-client example console app gives some examples on how to use the sdk and ultimately how to enact a payment between a producer and a consumer. The app runs as a console app, and you select items by entering the item number:

```
----------------------------- Worldpay Within SDK Client ----------------------------
0 - -------------------- GENERAL  --------------------
1 - Init new device
2 - Start RPC Service
3 - Get device info
4 - Load device profile
5 - Reset session
6 - -------------------- PRODUCER --------------------
7 - Init new producer
8 - Add RoboWash service
9 - Add RoboAir service
10 - Start service broadcast
11 - Stop broadcast
12 - -------------------- CONSUMER --------------------
13 - Prepare new consumer
14 - Scan services
15 - Auto consume from profile info
16 - --------------------------------------------------
17 - Exit
-------------------------------------------------------------------------------------
Please select choice:
```

There are example profiles in the root of the dev-client application directory: test-consumer.profile and test-producer.profile.
To load a profile:
* Enter 4 followed by the name of the profile you wish to load, for example 'test-producer.profile'.
* Then press 'y' to continue
* Enter 10 to start the service broadcast
* Enter a timeout in milliseconds for the broadcast e.g 60000.

Then in another terminal (or on another device connected to the same wifi network):
* Enter 4 to load another profile, for example 'test-consumer.profile'.
* Enter 15 to auto consume from the profile info - after a few seconds a payment should now occur based on the selection criteria in the "autoConsume" section of the device profile.

As shown below:

```
Please select choice: 15
Starting auto consume...
Found Service:: (192.168.43.143:8080/) - Car services offered by robot
Selecting service: 1 - Car tyre pressure checked and topped up by robot
Selecting price: (2) Measure and adjust pressure - four tyres for the price of three @ 75GBP, 4 Tyre (Unit id = 2)
Payment of 75 made successfully
```

This happens because the "autoConsume" section of 'test-consumer.profile' has the following data:

```
"autoConsume": {
	"deviceUid": "3b0f50f5-c4e6-4d0f-5cbf-52466e858ea3",
	"serviceID": 1,
	"unitID": 2
}
```
This means that the device with Uid "3b0f50f5-c4e6-4d0f-5cbf-52466e858ea3" will be selected. (The uuid can be found in a text file 'uuid.txt' that gets written to the device upon initialisation, you can use the dev-client app to initialise the device, or simply place a 'uuid.txt' file with just the uuid in, at the same location as the dev-client app).
When the device Uid has been selected, it then selects the service with serviceID 1, and finally it will select the unit with unitID 2.