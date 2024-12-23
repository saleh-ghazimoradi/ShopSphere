package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUp struct {
	UserLogin
	Phone string `json:"phone"`
}

type VerificationCodeInput struct {
	Code int `json:"code"`
}

type SellerInput struct {
	Firstname         string `json:"first_name"`
	Lastname          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	BankAccountNumber string `json:"bank_account_number"`
	SwiftCode         string `json:"swift_code"`
	PaymentType       string `json:"payment_type"`
}
