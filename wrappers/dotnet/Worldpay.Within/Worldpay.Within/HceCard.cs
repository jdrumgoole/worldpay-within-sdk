using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Worldpay.Innovation.WPWithin
{
    public class HceCard
    {
        public string FirstName { get; set; }

        public string LastName { get; set; }

        public int? ExpMonth { get; set; }

        public int? ExpYear { get; set; }

        public string CardNumber { get; set; }

        public string Type { get; set; }

        public string Cvc { get; set; }

        public override bool Equals(object that)
        {
            return new EqualsBuilder<HceCard>(this, that)
                .With(m => m.FirstName)
                .With(m => m.LastName)
                .With(m => m.ExpMonth)
                .With(m => m.ExpYear)
                .With(m => m.CardNumber)
                .With(m => m.Type)
                .With(m => m.Cvc)
                .Equals();
        }

        public override int GetHashCode()
        {
            return new HashCodeBuilder<HceCard>(this)
                .With(m => m.FirstName)
                .With(m => m.LastName)
                .With(m => m.ExpMonth)
                .With(m => m.ExpYear)
                .With(m => m.CardNumber)
                .With(m => m.Type)
                .With(m => m.Cvc)
                .HashCode;
        }

        public override string ToString()
        {
            return new ToStringBuilder<HceCard>(this)
                .Append(m => m.FirstName)
                .Append(m => m.LastName)
                .Append(m => m.ExpMonth)
                .Append(m => m.ExpYear)
                .Append(m => m.CardNumber)
                .Append(m => m.Type)
                .Append(m => m.Cvc)
                .ToString();
        }
    }
}
