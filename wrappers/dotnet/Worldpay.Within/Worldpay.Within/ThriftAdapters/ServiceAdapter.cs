using ThriftService = Worldpay.Innovation.WPWithin.Rpc.Types.Service;

namespace Worldpay.Innovation.WPWithin.ThriftAdapters
{
    internal class ServiceAdapter
    {
        public static ThriftService Create(Service service)
        {
            return new ThriftService()
            {
                Description = service.Description,
                Id = service.Id,
                Name = service.Name,
                Prices = CollectionUtils.Copy(service.Prices, PriceAdapter.Create)
            };
        }
        
    }
}