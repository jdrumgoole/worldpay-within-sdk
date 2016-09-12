namespace Worldpay.Innovation.WPWithin.ThriftAdapters
{
    internal class PaymentResponseAdapter
    {
        public static PaymentResponse Create(Rpc.Types.PaymentResponse makePayment)
        {
            return new PaymentResponse()
            {
                ClientId = makePayment.ClientId,
                ClientUuid = makePayment.ClientUUID,
                ServerId = makePayment.ServerId,
                ServiceDeliveryToken = ServiceDeliveryTokenAdapter.Create(makePayment.ServiceDeliveryToken),
                TotalPaid = makePayment.TotalPaid
            };
        }
    }
}