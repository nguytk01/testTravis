using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using NUnit.Framework;
using System.Runtime.Serialization.Json;
using System.Net.Http;
using System.Diagnostics;
using System.IO;

namespace Project1
{
    [TestFixture]
    public class TestStudent
    {
        string receivedSessionKey;
        const string serverIP = "127.0.0.1";
        const string serverPort = "8000";
        Process server;

        [OneTimeSetUp]
        public void RunBeforeAnyTests()
        {
            Console.WriteLine(Path.Combine(
            Directory.GetParent(Directory.GetCurrentDirectory()).FullName,
            "server.exe"));

            server = Process.Start(
            Path.Combine(
            Directory.GetParent(Directory.GetCurrentDirectory()).FullName,
            "server.exe"));

            Console.WriteLine("server ID : " + server.Id);
        }

        [TestCase]
        public void TestServer(){
            request().Wait();
            byte[] arr = Convert.FromBase64String(receivedSessionKey);
            Assert.AreEqual(30, arr.Length);
        }

        public async Task request(){
            HttpClient httpClient = new HttpClient();
            DataContractJsonSerializer ser = new DataContractJsonSerializer(typeof(SessionType));
            var content = await httpClient.GetAsync("http://" + serverIP + ":"+ serverPort+"/newUnauthorizedSession");
            System.IO.Stream stream = await content.Content.ReadAsStreamAsync();
            SessionType obj = (SessionType)ser.ReadObject(stream);
            receivedSessionKey = obj.Session;
        }

        [OneTimeTearDown]
        public void TearDown()
        {
            server.Kill();
        }
    }
}
