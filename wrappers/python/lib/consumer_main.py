from consumer import SampleConsumer
import ttypes as types
import wpwithin

hceCard = types.HCECard(FirstName='Samwise', LastName='Gamgee', ExpMonth=1, ExpYear=2018, CardNumber='3791421199999999', Type='Card', Cvc='865')

client = wpwithin.createClient("127.0.0.1", 9091, True)

consumer = SampleConsumer(client, hceCard)

consumer.client.setup("Python3-Device", "Sample Python3 consumer device")
    
print("Calling discover devices...")

serviceMessage = consumer.discoverDevices()

if serviceMessage != None:
    consumer.connectToDevice(serviceMessage)
    consumer.purchaseFirstServiceFirstPrice(1)
