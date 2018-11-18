using System;
using System.Threading.Tasks;
using System.Diagnostics;
using NUnit.Framework;
using System.IO;
using System.Net.Http;
using System.Runtime.Serialization.Json;

namespace timetracker
{
    class ServerProxyTest
    {
        const string serverIP = "127.0.0.1";
        const string serverPort = "8000";
        Process server;
        ServerProxy serverProxy;

       /* [OneTimeSetUp]
        public void RunBeforeAnyTests()
        {
            Console.WriteLine("Server binary path: " + 
            Directory.GetParent(
            System.AppDomain.CurrentDomain.SetupInformation.ApplicationBase)
            .Parent.Parent.Parent.FullName,
            "server.exe");

            server = Process.Start(
            Path.Combine(
            Directory.GetParent(
            System.AppDomain.CurrentDomain.SetupInformation.ApplicationBase)
            .Parent.Parent.Parent.FullName,
            "server.exe"));

            serverProxy = new ServerProxy();
            Console.WriteLine("server PID : " + server.Id);
        }*/

        [TestCase]
        public async Task TestGetUnauthorizedSession()
        {
            SessionType session = await serverProxy.GetUnauthorizedSession();
            Assert.AreEqual(30, session.getSessionKey().Length);
        }

        [TestCase]
        public async Task TestLogin()
        {
            LoginResultType loginResultType = await serverProxy.LogIn(new LoginData()
            {
                Email = "email",
                Password = "password",
                SessionKey = "zaaaaaa"
            });
            Assert.AreEqual("Success", loginResultType.LoginResult);
        }
        [TestCase]
        public async Task CreateAccountForm()
        {
            CreateAccountResultType createAccountResult = await serverProxy.CreateAc(new CreateAccount()
            {
                FirstName = "Blah",
                MiddleName = "Blah",
                LastName ="Blah",
                Email = "email",
                Password = "password",
                SessionKey = "zaaaaaa"
            });
            Assert.AreEqual("Success", createAccountResult.CreateAccountResult);
        }
        /*[OneTimeTearDown]
        public void TearDown()
        {
            server.Kill();
        }*/
    }
}
