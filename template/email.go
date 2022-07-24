package template

import "fmt"

type Email struct {
	B string
}

func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
