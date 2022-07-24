package template

import "log"

/*
	什么是模版设计模式：
		1. 组成模版的方法都被定义在父类中，子类自己去实现
        2. 只要子类中实现具体的逻辑，当父类被调用时程序的行为就是一致的
*/

// 定义算法步骤，将具体处理叫个子类
type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

func Gen() {
	sms := &Sms{A: "a"}
	otp := Otp{sms}
	log.Println(otp.genAndSendOTP(4))

	email := &Email{"B"}
	otpEmail := Otp{email}
	log.Println(otpEmail.genAndSendOTP(4))
}
