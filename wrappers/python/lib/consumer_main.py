from consumer import SampleConsumer
import ttypes as types
import wpwithin

hceCard = types.HCECard(firstName='Samwise', lastName='Gamgee', expMonth=1, expYear=2018, cardNumber='3791421199999999', cardType='Card', cvc='865')

out = wpwithin.createClient("127.0.0.1", 9091, True)

client = out['client']
agent = out['rpc']

consumer = SampleConsumer(client, hceCard)

consumer.client.setup("Python3-Device", "Sample Python3 consumer device")
    
print("Calling discover devices...")

serviceMessage = consumer.discoverDevices()

if serviceMessage is not None:
    consumer.connectToDevice(serviceMessage)
    consumer.purchaseFirstServiceFirstPrice(1)
