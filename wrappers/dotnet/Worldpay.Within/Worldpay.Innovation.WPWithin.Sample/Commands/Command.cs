using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Worldpay.Innovation.WPWithin.Sample.Commands
{

    /// <summary>
    /// Represents a single menu item that the user can invoke.
    /// </summary>
    public class Command
    {
        public string Name { get; private set; }
        public string Description { get; private set; }
        public Func<string[], CommandResult> Function { get; private set; }

        public Command(string name, string description, Func<string[], CommandResult> function)
        {
            Name = name;
            Description = description;
            Function = function;
        }
    }



}
