using ThriftServiceDeliveryToken=Worldpay.Innovation.WPWithin.Rpc.Types.ServiceDeliveryToken;

namespace Worldpay.Innovation.WPWithin.ThriftAdapters
{
    public class ServiceDeliveryTokenAdapter
    {
        public static ThriftServiceDeliveryToken Create(ServiceDeliveryToken serviceDeliveryToken)
        {
            return new ThriftServiceDeliveryToken
            {
                Key = serviceDeliveryToken.Key,
                RefundOnExpiry = serviceDeliveryToken.RefundOnExpiry,
                Signature = serviceDeliveryToken.Signature,
                Expiry = serviceDeliveryToken.Expiry,
                Issued = serviceDeliveryToken.Issued
            };
        }
    }
}