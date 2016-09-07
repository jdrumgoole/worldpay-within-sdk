
using Worldpay.Innovation.WPWithin.Sample.Commands;

namespace Worldpay.Innovation.WPWithin.Sample
{
    internal class Program
    {
        private static void Main(string[] args)
        {
            new Program().Run(args);
        }

        private void Run(string[] args)
        {
            bool terminate = false;
            CommandMenu menu = new CommandMenu();
            while (!terminate)
            {
                CommandResult result = menu.ReadEvalPrint();
                if (result == CommandResult.CriticalFailure || result == CommandResult.Exit)
                {
                    terminate = true;
                }
            }
        }
    }
}