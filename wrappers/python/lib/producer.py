import ttypes as types
import wpwithin


out = wpwithin.createClient("127.0.0.1", 9090, True)

client = out['client']
agent = out['rpc']

client.setup("Python3 Device", "Sample Python3 producer device")

pricePerUnit = types.PricePerUnit(amount=650, currencyCode="GBP")
rwPrice = types.Price(id=1, description="Car Wash", pricePerUnit=pricePerUnit, unitId=1, unitDescription="Single wash")
service = types.Service(id=1, name="RoboWash", description="Car washed by robot", prices={1: rwPrice})

client.addService(service)

client.initProducer(merchantClientKey="A_P_03eaa1d3-4642-4079-b030-b543ee04b5af", merchantServiceKey="A_P_f50ecb46-ca82-44a7-9c40-421818af5996")

client.startServiceBroadcast(20000)