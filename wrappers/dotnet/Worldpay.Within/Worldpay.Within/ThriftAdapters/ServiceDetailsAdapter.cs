using System.Collections.Generic;
using System.Linq;
using Thrift.Collections;
using ThriftServiceDetails = Worldpay.Innovation.WPWithin.Rpc.Types.ServiceDetails;

namespace Worldpay.Innovation.WPWithin.ThriftAdapters
{
    internal class ServiceDetailsAdapter
    {
        public static IEnumerable<ServiceDetails> Create(THashSet<ThriftServiceDetails> requestServices)
        {
            return requestServices.Select(Create);
        }

        public static ServiceDetails Create(ThriftServiceDetails tsd)
        {
            return new ServiceDetails
            {
                ServiceId = tsd.ServiceId,
                ServiceDescription = tsd.ServiceDescription
            };
        }
    }
}