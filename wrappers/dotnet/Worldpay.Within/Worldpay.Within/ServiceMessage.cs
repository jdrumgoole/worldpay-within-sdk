using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Worldpay.Innovation.WPWithin
{
    public class ServiceMessage
    {
        public string DeviceDescription { get; set; }

        public string Hostname { get; set; }

        public int? PortNumber { get; set; }

        public string ServerId { get; set; }

        public string UrlPrefix { get; set; }

        public override bool Equals(object that)
        {
            return new EqualsBuilder<ServiceMessage>(this, that)
                .With(m => m.DeviceDescription)
                .With(m => m.Hostname)
                .With(m => m.PortNumber)
                .With(m => m.ServerId)
                .With(m => m.UrlPrefix)
                .Equals();
        }

        public override int GetHashCode()
        {
            return new HashCodeBuilder<ServiceMessage>(this)
                .With(m => m.DeviceDescription)
                .With(m => m.Hostname)
                .With(m => m.PortNumber)
                .With(m => m.ServerId)
                .With(m => m.UrlPrefix)
                .HashCode;
        }

        public override string ToString()
        {
            return new ToStringBuilder<ServiceMessage>(this)
                .Append(m => m.DeviceDescription)
                .Append(m => m.Hostname)
                .Append(m => m.PortNumber)
                .Append(m => m.ServerId)
                .Append(m => m.UrlPrefix)
                .ToString();
        }
    }
}
