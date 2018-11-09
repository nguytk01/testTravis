using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Runtime.Serialization;
namespace Project1
{
    [DataContract]
    class SessionType
    {
        [DataMember]
        internal string Session;
    }
}
