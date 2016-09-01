using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;
using Worldpay.Innovation.WPWithin.Sample.Commands;

namespace Worldpay.Innovation.WPWithin.Sample
{

    class Program
    {
        static void Main(string[] args)
        {
            new Program().Run(args);
        }

        private void WriteLine(string fmt, params object[] parameters)
        {
            Console.WriteLine(String.Format(fmt, parameters));
        }


        private void Run(string[] args)
        {
            bool terminate = false;
            CommandMenu menu = new CommandMenu();
            while (!terminate)
            {
                CommandResult result = menu.ProcessCommand();
                if (result == CommandResult.CriticalFailure || result == CommandResult.Exit)
                {
                    terminate = true;
                }
            }
        }
    }
}
