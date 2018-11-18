using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Runtime.Serialization;
using System.Runtime.Serialization.Json;

namespace TimeTracker
{
    [DataContract]
    public sealed class LoginResultType
    {

            private static DataContractJsonSerializer deserializerLoginResultType =
                new DataContractJsonSerializer(typeof(LoginResultType));

            [DataMember]
            public string LoginResult;
            [DataMember]
            public string SessionKey;

            public static LoginResultType ReadFromStream(System.IO.Stream stream)
            {
                return (LoginResultType)deserializerLoginResultType.ReadObject(stream);
            }
        }

    
}
