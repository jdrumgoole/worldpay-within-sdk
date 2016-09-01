using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Worldpay.Innovation.WPWithin.Sample.Commands
{
    public enum CommandResult
    {
        NoOp,
        NoSuchCommand,
        Success,
        Failure,
        CriticalFailure,
        Exit
    }

}
