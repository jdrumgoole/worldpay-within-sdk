using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using System.Threading.Tasks;

namespace Worldpay.Innovation.WPWithin
{
    /**
      * Taken from http://www.jroller.com/DhavalDalal/entry/equals_hashcode_and_tostring_builders
      * Code licensed under CC-3.0
      */
    public class HashCodeBuilder<T>
    {
        private readonly T target;
        private int hashCode = 17;

        public HashCodeBuilder(T target)
        {
            this.target = target;
        }

        public HashCodeBuilder<T> With<TProperty>(Expression<Func<T, TProperty>> propertyOrField)
        {
            var expression = propertyOrField.Body as MemberExpression;
            if (expression == null)
            {
                throw new ArgumentException("Expecting Property or Field Expression of an object");
            }

            var func = propertyOrField.Compile();
            var value = func(target);
            hashCode += 31 * hashCode + ((value == null) ? 0 : value.GetHashCode());
            return this;
        }

        public int HashCode
        {
            get
            {
                return hashCode;
            }
        }
    }
}
