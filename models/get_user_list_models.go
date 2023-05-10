package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}

type GetUsersWithCategory struct {
	ID           int64   `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        *string `json:"email"`
	Mobile       string  `json:"mobile"`
	LastLogin    string  `json:"last_login"`
	SentOtp      string  `json:"sent_otp"`
	IsNewuser    int64   `json:"is_newuser"`
	IsSignup     int64   `json:"is_signup"`
	OtpVerified  int64   `json:"otp_verified"`
	IsWaitlisted int64   `json:"is_waitlisted"`
	IsCustomer   int64   `json:"is_customer"`
	FirebaseID   string  `json:"firebase_id"`
	TempID       string  `json:"temp_id"`
	InnerCircle  int64   `json:"inner_circle"`
	// PasswordResetToken string   `json:"password_reset_token"`
	IsActive         int64  `json:"is_active"`
	IsKyc            int64  `json:"is_kyc"`
	Gender           string `json:"gender"`
	Dob              string `json:"dob"`
	RoleID           int64  `json:"role_id"`
	AdminTypeID      string `json:"admin_type_id"`
	Username         string `json:"username"`
	UserProfileImage string `json:"user_profile_image"`
	Password         string `json:"password"`
	InvitedBy        string `json:"invited_by"`
	InvitedCount     int64  `json:"invited_count"`
	IsNotified       int64  `json:"is_notified"`
	// Source             interface{}   `json:"source"`
	// Medium             interface{}   `json:"medium"`
	// Campaign           interface{}   `json:"campaign"`
	AppVersion  string     `json:"app_version"`
	Platform    string     `json:"platform"`
	LoginAt     *string    `json:"login_at"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	Age         int64      `json:"age"`
	Height      string     `json:"height"`
	Weight      string     `json:"weight"`
	CustomerID  int64      `json:"customer_id"`
	OrderID     *int64     `json:"order_id"`
	ServiceID   *int64     `json:"service_id"`
	SessionType *string    `json:"session_type"`
	Type        *string    `json:"type"`
	Status      *string    `json:"status"`
	Category    []Category `gorm:"foreignKey:id"`
}

type GetAllUsers struct {
	ID           int64   `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        *string `json:"email"`
	Mobile       string  `json:"mobile"`
	LastLogin    string  `json:"last_login"`
	SentOtp      string  `json:"sent_otp"`
	IsNewuser    int64   `json:"is_newuser"`
	IsSignup     int64   `json:"is_signup"`
	OtpVerified  int64   `json:"otp_verified"`
	IsWaitlisted int64   `json:"is_waitlisted"`
	IsCustomer   int64   `json:"is_customer"`
	FirebaseID   string  `json:"firebase_id"`
	TempID       string  `json:"temp_id"`
	InnerCircle  int64   `json:"inner_circle"`
	// PasswordResetToken string   `json:"password_reset_token"`
	IsActive         int64  `json:"is_active"`
	IsKyc            int64  `json:"is_kyc"`
	Gender           string `json:"gender"`
	Dob              string `json:"dob"`
	RoleID           int64  `json:"role_id"`
	AdminTypeID      string `json:"admin_type_id"`
	Username         string `json:"username"`
	UserProfileImage string `json:"user_profile_image"`
	Password         string `json:"password"`
	InvitedBy        string `json:"invited_by"`
	InvitedCount     int64  `json:"invited_count"`
	IsNotified       int64  `json:"is_notified"`
	// Source             interface{}   `json:"source"`
	// Medium             interface{}   `json:"medium"`
	// Campaign           interface{}   `json:"campaign"`
	AppVersion  string  `json:"app_version"`
	Platform    string  `json:"platform"`
	LoginAt     *string `json:"login_at"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Age         int64   `json:"age"`
	Height      string  `json:"height"`
	Weight      string  `json:"weight"`
	CustomerID  int64   `json:"customer_id"`
	OrderID     *int64  `json:"order_id"`
	ServiceID   *int64  `json:"service_id"`
	SessionType *string `json:"session_type"`
	Type        *string `json:"type"`
	Status      *string `json:"status"`
	// Category    []Category `gorm:"foreignKey:id"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
