import ttypes as types


class SampleConsumer:
    def __init__(self, client, hceCard):
        self.client = client
        self.hceCard = hceCard

    def discoverDevices(self):
        try:
            serviceMessages = self.client.deviceDiscovery(100000)
        except types.Error as err:
            print("deviceDiscovery.callback.err: " + err.message)
            raise err
        if len(serviceMessages) == 0:
            print("Did not discover any devices on the network.")
            return None

        print(serviceMessages)
        deviceCount = len(serviceMessages)
        print("Discovered {0} devices on the network.".format(len(serviceMessages)))
        print("Devices:")

        for i in range(deviceCount):
            deviceLog = '''Description: {0.deviceDescription}
            Hostname: {0.hostname}
            Port: {0.portNumber}
            Server ID: {0.serverId}
            URL Prefix: {0.urlPrefix}'''.format(serviceMessages[i])

            print(deviceLog)
        
        # Connect to first device
        return serviceMessages[0]

    def connectToDevice(self, serviceMessage):
        try:
            self.client.initConsumer('http://', serviceMessage.hostname, serviceMessage.portNumber, serviceMessage.urlPrefix, serviceMessage.serverId, self.hceCard)
        except types.Error as err:
            print("initConsumer.callback.err: " + err.message)
            raise err
        else:
            print("Initialised consumer.")

    def getAvailableServices(self):
        try:
            serviceDetails = self.client.requestServices()
        except types.Error as err:
            print("requestServices.callback.err: " + err.message)
            raise err
        if len(serviceDetails) == 0:
            service = serviceDetails[0]
            print("Services:")
            print("Id: " + service.serviceId)
            print("Description: " + service.serviceDescription)
            print("----------")
            return service.serviceId
            
    def getServicePrices(self, serviceId):
        try:
            prices = self.client.getServicePrices(serviceId)
        except types.Error as err:
            print("requestServicePrices.callback.err: " + err.message)
            raise err
        
        if len(prices) > 0:
            price = prices[0]
            printMessage = """Price details for serviceId {0}:
            Id: {1.id}
            Description: {1.description}
            UnitId: {1.unitId}
            Unit Description: {1.unitDescription}
            Price per Unit:
            \tAmount: {1.pricePerUnit.amount}
            \tCurrency Code: {1.pricePerUnit.currencyCode}
            ----------
            """.format(serviceId, price)

        else:
            raise types.Error("Did not receive any service prices :/")

        return price

    def getServicePriceQuote(self, serviceId, numberOfUnits, priceId):
        try:
            priceResponse = self.client.selectService(serviceId, numberOfUnits, priceId)
        except types.Error as err:
            print("selectService.callback.err: " + err.message)
            raise err

        if priceResponse is None:
            print("Did not receive total price response from selectService()")
            return None

        printMessage = """Total Price Response:
        Server ID: {0.serverId}
        Client ID: {0.clientId}
        Price ID: {0.priceId}
        Units to supply: {0.unitsToSupply}
        Total price: {0.totalPrice}
        Payment Reference ID: {0.paymentReferenceId}
        Merchant Client Key: {0.merchantClientKey}
        ------""".format(priceResponse)
        
        return priceResponse
        
    def purchaseFirstServiceFirstPrice(self, numberOfUnits):

        serviceId = self.getAvailableServices()

        priceId = self.getServicePrices(serviceId)

        priceResponse = self.getServicePriceQuote(serviceId, numberOfUnits, priceId)

        try:
            response = self.client.makePayment(priceResponse)
        except types.Error as err:
            print("makePayment.callback.err: " + err.message)
            raise err

        if response is None:
            print("Did not receive correct response to make payment")
            return None
        
        printMessage = """Response from make payment:
        Server ID: {0.serverId}
        Client ID: {0.clientId}
        Total Paid: {0.totalPaid}
        Service Delivery Token:
        \tKey: {1.key}
        \tIssued: {1.issues}
        \tExpiry: {1.expiry}
        \tRefund on expiry: {1.refundOnExpiry}
        \tSignature: {1.signature}
        Client UUID: {0.clientUUID}""".format(response, response.serviceDeliveryToken)

        print(printMessage)

        print("Shutting down...")
        self.client.proc.kill()