using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Worldpay.Innovation.WPWithin
{
    public class ServiceDetails
    {

        public int? ServiceId { get; set; }

        public string ServiceDescription { get; set; }

        public override bool Equals(object that)
        {
            return new EqualsBuilder<ServiceDetails>(this, that)
                .With(m => m.ServiceId)
                .With(m => m.ServiceDescription)
                .Equals();
        }

        public override int GetHashCode()
        {
            return new HashCodeBuilder<ServiceDetails>(this)
                .With(m => m.ServiceId)
                .With(m => m.ServiceDescription)
                .HashCode;
        }

        public override string ToString()
        {
            return new ToStringBuilder<ServiceDetails>(this)
                .Append(m => m.ServiceId)
                .Append(m => m.ServiceDescription)
                .ToString();
        }
    }
}
