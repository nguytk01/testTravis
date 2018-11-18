using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace TimeTracker
{
    /// <summary>
    /// Interaction logic for Login.xaml
    /// </summary>
    public partial class Login : Page
    {
        private SessionType sessionObj;
        private LoginResultType resultObj;
        public Login()
        {
            InitializeComponent();
        }

        private void CreateAccountBtn_Click(object sender, RoutedEventArgs e)
        {
            NavigationService.Navigate(new CreateAccount());
        }

        private async void LoginBtn_Click(object sender, RoutedEventArgs e)
        {


            bool correctInput = VerifyInput(Email.Text, Password.Text);
            bool correctEmail = VerifyEmail(Email.Text);
            LoginData log = new LoginData {

                Email = Email.Text,
                Password = Password.Text
            };
            
           
            if (correctEmail == true && correctInput == true) {
                sessionObj = await ServerProxySingleton.serverProxy.GetUnauthorizedSession();
                log.SessionKey = sessionObj.Session;

                resultObj = await ServerProxySingleton.serverProxy.LogIn(log);
                if (resultObj.LoginResult.Equals("Success")) 

                
                { NavigationService.Navigate(new Home()); }

                else { MessageBox.Show("Login fail! Please Provide correct user name and password"); }
            }

           
        }
        private bool VerifyEmail(String email) {
          while(email != "") { 
            try
            {

                

                    var addr = new System.Net.Mail.MailAddress(email);
                    return addr.Address == email;

            }

            catch
            {

                MessageBox.Show("Invalid Email");
                    return false;

            }
          }

            return true;
        }
        private bool VerifyInput(String email, String password)
        {

            if (Email.Text == "" || Password.Text == "")
            {

                MessageBox.Show("Email & password must be provided");
                return false;
            }

            else {
                return true;
            }
        }

        private void ForgotPasswordBtn_Click(object sender, RoutedEventArgs e)
        {

        }
    }
}
