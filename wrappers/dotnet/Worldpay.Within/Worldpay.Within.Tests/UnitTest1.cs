using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Worldpay.Innovation.WPWithin;

namespace Worldpay.Within.Tests
{
    [TestClass]
    public class UnitTest1
    {
        [TestMethod]
        public void TestMethod1()
        {
            Main m = new Main();
            m.SendSimpleMessage();
        }
    }
}
