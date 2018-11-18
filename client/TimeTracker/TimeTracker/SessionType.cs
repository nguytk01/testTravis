using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Runtime.Serialization.Json;
using System.Runtime.Serialization;
namespace TimeTracker
{
    public class SessionType
    {
        [DataMember]
        public string Session;

        private static DataContractJsonSerializer deserializerSessionType =
            new DataContractJsonSerializer(typeof(SessionType));

        public byte[] getSessionKey()
        {
            return Convert.FromBase64String(Session);
        }

        public static SessionType ReadFromStream(System.IO.Stream stream)
        {
            return (SessionType)deserializerSessionType.ReadObject(stream);
        }

    }
}
