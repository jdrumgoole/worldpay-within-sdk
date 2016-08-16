using System.Collections.Generic;
using System.Linq;
using Thrift.Collections;
using ThriftServiceMessage = Worldpay.Innovation.WPWithin.Rpc.Types.ServiceMessage;

namespace Worldpay.Innovation.WPWithin.ThriftAdapters
{
    public class ServiceMessageAdapter
    {
        public static IEnumerable<ServiceMessage> Create(THashSet<Rpc.Types.ServiceMessage> deviceDiscovery)
        {
            return deviceDiscovery.Select(Create);
        }

        private static ServiceMessage Create(ThriftServiceMessage sm)
        {
            return new ServiceMessage()
            {
                DeviceDescription = sm.DeviceDescription,
                Hostname = sm.Hostname,
                PortNumber = sm.PortNumber,
                ServerId = sm.ServerId,
                UrlPrefix = sm.UrlPrefix
            };
        }
    }
}