using System;
using ThriftPrice = Worldpay.Innovation.WPWithin.Rpc.Types.Price;

namespace Worldpay.Innovation.WPWithin.ThriftAdapters
{
    internal class PriceAdapter
    {
        internal static ThriftPrice Create(Price price)
        {
            return new ThriftPrice()
            {
                Description = price.Description,
                Id = price.Id,
                PricePerUnit = PricePerUnitAdapter.Create(price.PricePerUnit),
                UnitDescription = price.UnitDescription,
                UnitId = price.UnitId
            };
        }
    }
}