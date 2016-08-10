using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Worldpay.Innovation.WPWithin
{
    public class Price
    {

        public int? ServiceId { get; set; }

        public int? Id { get; set; }

        public string Description { get; set; }

        public PricePerUnit PricePerUnit { get; set; }

        public int? UnitId { get; set; }

        public string UnitDescription { get; set; }

        public override bool Equals(object that)
        {
            return new EqualsBuilder<Price>(this, that)
                .With(m => m.ServiceId)
                .With(m => m.Id)
                .With(m => m.Description)
                .With(m => m.PricePerUnit)
                .With(m => m.UnitId)
                .With(m => m.UnitDescription)
                .Equals();
        }

        public override int GetHashCode()
        {
            return new HashCodeBuilder<Price>(this)
                .With(m => m.ServiceId)
                .With(m => m.Id)
                .With(m => m.Description)
                .With(m => m.PricePerUnit)
                .With(m => m.UnitId)
                .With(m => m.UnitDescription)
                .HashCode;
        }

        public override string ToString()
        {
            return new ToStringBuilder<Price>(this)
                .Append(m => m.ServiceId)
                .Append(m => m.Id)
                .Append(m => m.Description)
                .Append(m => m.PricePerUnit)
                .Append(m => m.UnitId)
                .Append(m => m.UnitDescription)
                .ToString();

        }

    }
}
