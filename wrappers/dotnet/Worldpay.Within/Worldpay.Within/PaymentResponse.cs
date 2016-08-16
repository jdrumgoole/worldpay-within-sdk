using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Worldpay.Innovation.WPWithin
{
    public class PaymentResponse
    {

        public string ServerId { get; set; }

        public string ClientId { get; set; }

        public int? TotalPaid { get; set; }

        public string ServiceDeliveryToken { get; set; }

        public string Clientuid { get; set; }

        public override bool Equals(object that)
        {
            return new EqualsBuilder<PaymentResponse>(this, that)
                .With(m => m.ServerId)
                .With(m => m.ClientId)
                .With(m => m.TotalPaid)
                .With(m => m.ServiceDeliveryToken)
                .With(m => m.Clientuid)
                .Equals();
        }

        public override int GetHashCode()
        {
            return new HashCodeBuilder<PaymentResponse>(this)
                .With(m => m.ServerId)
                .With(m => m.ClientId)
                .With(m => m.TotalPaid)
                .With(m => m.ServiceDeliveryToken)
                .With(m => m.Clientuid)
                .HashCode;
        }

        public override string ToString()
        {
            return new ToStringBuilder<PaymentResponse>(this)
                .Append(m => m.ServerId)
                .Append(m => m.ClientId)
                .Append(m => m.TotalPaid)
                .Append(m => m.ServiceDeliveryToken)
                .Append(m => m.Clientuid)
                .ToString();
        }
    }
}
