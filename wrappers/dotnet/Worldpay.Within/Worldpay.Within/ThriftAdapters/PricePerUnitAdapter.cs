using System;
using ThriftPricePerUnit = Worldpay.Innovation.WPWithin.Rpc.Types.PricePerUnit;

namespace Worldpay.Innovation.WPWithin.ThriftAdapters
{
    internal class PricePerUnitAdapter
    {
        internal static ThriftPricePerUnit Create(PricePerUnit pricePerUnit)
        {
            return new ThriftPricePerUnit()
            {
                Amount = pricePerUnit.Amount,
                CurrencyCode = pricePerUnit.CurrencyCode
            };
        }
    }
}