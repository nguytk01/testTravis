using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Net.Sockets;
using System.Diagnostics;
using System.Net.Http;
using System.Runtime.Serialization.Json;


namespace TimeTracker
{
    public class ServerProxy
    {
        public static ServerProxy instance = null;

        HttpClient httpClient = new HttpClient();

        public string host = "127.0.0.1";
        public int port = 8000;
        public string serverURL;
        private static readonly object mutex = new object();
        public ServerProxy Instance
        {
            get
            {
                lock (mutex)
                {
                    if (instance == null)
                    {
                        instance = new ServerProxy();
                    }
                    return instance;
                }
            }
        }

        public ServerProxy()
        {
            try
            {

                Console.WriteLine("Connecting.....");
                serverURL = "http://" + host + ":" + port;
            }
            catch (SocketException SE)
            {
                string error = "An error occured while connecting [" + SE.Message + "]\n";
            }
        }
        public async Task<SessionType> GetUnauthorizedSession()
        {
            var content = await httpClient.GetAsync("http://" + host + ":" + port + "/newUnauthorizedSession");
            System.IO.Stream stream = await content.Content.ReadAsStreamAsync();
            return SessionType.ReadFromStream(stream);
        }

        public async Task<LoginResultType> LogIn(LoginData data)
        {
            data.TimeZone = TimeZone.CurrentTimeZone.GetUtcOffset(DateTime.Today).Hours;
            HttpContent httpContent = new ByteArrayContent(data.GetMemoryStream().ToArray());
            var content = await httpClient.PostAsync(serverURL + "/Login", httpContent);
            return LoginResultType.ReadFromStream(await content.Content.ReadAsStreamAsync());
        }
        public async Task<CreateAccountResultType> CreateAccount(CreateAccountData data)
        {
            data.TimeZone = TimeZone.CurrentTimeZone.GetUtcOffset(DateTime.Today).Hours;
            HttpContent httpContent = new ByteArrayContent(data.GetMemoryStream().ToArray());
            var content = await httpClient.PostAsync(serverURL + "/CreateAccount", httpContent);
            return CreateAccountResultType.ReadFromStream(await content.Content.ReadAsStreamAsync());
        }
    }
}
