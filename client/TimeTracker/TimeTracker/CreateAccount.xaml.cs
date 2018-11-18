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
    /// Interaction logic for CreateAccount.xaml
    /// </summary>
    public partial class CreateAccount : Page
    {
        private SessionType sessionObj;
        private CreateAccountResultType createAccountObj;
        public CreateAccount()
        {
            InitializeComponent();
        }

        private async  void CreateBtn_Click(object sender, RoutedEventArgs e)
        {
            bool correctInput = VerifyInput(FirstName.Text,MiddleName.Text,LastName.Text, Email.Text, Password.Text,RetypePassword.Text);
            bool correctEmail = VerifyEmail(Email.Text);
            CreateAccountData createAccount = new CreateAccountData {

                FirstName = FirstName.Text,
                MiddleName = MiddleName.Text,
                LastName = LastName.Text,
                Email = Email.Text,
                Password = Password.Text

            };
            if (correctEmail == true && correctInput == true) {

                sessionObj = await ServerProxySingleton.serverProxy.GetUnauthorizedSession();
                createAccount.SessionKey = sessionObj.Session;

                createAccountObj = await ServerProxySingleton.serverProxy.CreateAccount(createAccount);
                if (createAccountObj.CreateAccountResult.Equals("Success"))
                {

                    NavigationService.Navigate(new Home());
                }

                else { MessageBox.Show("Registration fail"); }
            }
            

        }
        private bool VerifyEmail(String email)
        {
            while (email != "")
            {
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
        private bool VerifyInput(String firstname,String middlename, String lastname, String email, String password,String retypepassword)
        {

            if (FirstName.Text == ""||MiddleName.Text == ""|| LastName.Text==""|| Email.Text == "" || Password.Text == "" || RetypePassword.Text == "")
            {

                MessageBox.Show("FullName,Email, password and Retype Password must be provided");
                return false;

                
            }

            else if (RetypePassword.Text != Password.Text)
            {

                MessageBox.Show("Password Doesn't match");

                return false;

            }
            else
            {
                return true;
            }
        }
    }
}
