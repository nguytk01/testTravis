using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using NUnit.Framework;
using System.Runtime.Serialization.Json;
using System.Net.Http;

namespace Project1
{
    
    [TestFixture]
    public class TestStudent
    {
        string receivedSessionKey;
        const string serverIP = "127.0.0.1:8000";

        [TestCase]
        public void TestServer(){
            request().Wait();
            byte[] arr = Convert.FromBase64String(receivedSessionKey);
            Assert.AreEqual(30, arr.Length);

        }

        public async Task request(){
            HttpClient httpClient = new HttpClient();
            DataContractJsonSerializer ser = new DataContractJsonSerializer(typeof(SessionType));
            var content = await httpClient.GetAsync("http://127.0.0.1:8000/newUnauthorizedSession");
            System.IO.Stream stream = await content.Content.ReadAsStreamAsync();
            SessionType obj = (SessionType)ser.ReadObject(stream);
            receivedSessionKey = obj.Session;
        }

        
    }
}
