using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Worldpay.Innovation.WPWithin
{
    public struct PricePerUnit
    {
        public int? Amount { get; set; }

        public string CurrencyCode { get; set; }

        public override bool Equals(object that)
        {
            return new EqualsBuilder<PricePerUnit>(this, that)
                .With(m => m.Amount)
            .With(m => m.CurrencyCode).Equals();
        }

        public override int GetHashCode()
        {
            return new HashCodeBuilder<PricePerUnit>(this)
                .With(m => m.Amount)
                .With(m => m.CurrencyCode).HashCode;
        }

        public override string ToString()
        {
            return new ToStringBuilder<PricePerUnit>(this)
                .Append(m => m.Amount)
                .Append(m => m.CurrencyCode)
                .ToString();
        }
    }
}
