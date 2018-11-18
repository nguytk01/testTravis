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
    public sealed class CreateAccountResultType
    {
        [DataMember]
        public string CreateAccountResult;
        [DataMember]
        public string SessionKey;
        private static DataContractJsonSerializer deserializerCreateAccountResultType =
                new DataContractJsonSerializer(typeof(CreateAccountResultType));
        public static CreateAccountResultType ReadFromStream(System.IO.Stream stream)
        {
            return (CreateAccountResultType)deserializerCreateAccountResultType.ReadObject(stream);
        }
    }
}
