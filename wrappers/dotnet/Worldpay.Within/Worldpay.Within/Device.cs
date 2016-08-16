using System.Collections.Generic;

namespace Worldpay.Innovation.WPWithin
{
    public class Device
    {

        public string Uid { get; set; }

        public string Name { get; set; }

        public string Description { get; set; }

        public Dictionary<int, Service> Services { get; set; }

        public System.Net.IPAddress Ipv4Address { get; set; }

        public string CurrencyCode { get; set; }

        public override bool Equals(object that)
        {
            return new EqualsBuilder<Device>(this, that)
                .With(m => m.Uid)
                .With(m => m.Name)
                .With(m => m.Description)
                .With(m => m.Services)
                .With(m => m.Ipv4Address)
                .With(m => m.CurrencyCode)
                .Equals();
        }

        public override int GetHashCode()
        {
            return new HashCodeBuilder<Device>(this)
                .With(m => m.Uid)
                .With(m => m.Name)
                .With(m => m.Description)
                .With(m => m.Services)
                .With(m => m.Ipv4Address)
                .With(m => m.CurrencyCode)
                .HashCode;
        }

        public override string ToString()
        {
            return new ToStringBuilder<Device>(this)
                .Append(m => m.Uid)
                .Append(m => m.Name)
                .Append(m => m.Description)
                .Append(m => m.Services)
                .Append(m => m.Ipv4Address)
                .Append(m => m.CurrencyCode)
                .ToString();
        }
    }
}
