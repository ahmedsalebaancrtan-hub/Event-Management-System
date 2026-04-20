package helpers

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendOTPEmail(toEmail string, otp string) error {
	// 1. Setup credentials (Use Environment Variables for security!)
	from := os.Getenv("EMAIL_USER")     // Your Gmail: e.g. "example@gmail.com"
	password := os.Getenv("EMAIL_PASS") // Your App Password (not your login password)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// 2. Create the message
	subject := "Subject: Your Password Reset OTP\n"
	body := fmt.Sprintf("Your one-time password for resetting your account is: %s\nThis code expires in 10 minutes.", otp)
	message := []byte(subject + "\n" + body)

	// 3. Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// 4. Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	if err != nil {
		return err
	}
	return nil
}
