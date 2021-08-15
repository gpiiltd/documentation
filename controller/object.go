package controller

import (
	"time"
)

//FiriModel default models for GORM timestamps
type FiriModel struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

//UserAuthentication holds user authentication data
type UserAuthentication struct {
	FiriModel
	FullName         string `gorm:"type:varchar(100)" json:"full_name"`
	Email            string `gorm:"type:varchar(100); unique" json:"email"`
	Password         string `gorm:"type:varchar(100)" json:"password"`
	Status           string `gorm:"type:varchar(100)" json:"status"`
	VerificationCode string `gorm:"type:varchar(100)" json:"verification_code"`
}

type VerifyRegistration struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type VerifyToken struct {
	Validity bool   `json:"validity"`
	Token    string `json:"token_string"`
}

//RequestPasswordReset holds the data for getting a reset password
type RequestPasswordReset struct {
	FiriModel
	Email            string `gorm:"type:varchar(100)" json:"email"`
	VerificationCode string `gorm:"type:varchar(100)" json:"verification_code"`
}

//SocialMedia holds social media data
type SocialMedia struct {
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
}

//UserData holds user data
type UserData struct {
	FiriModel
	UUID           string      `gorm:"type:varchar(100)" json:"uuid"`
	FullName       string      `gorm:"type:varchar(100)" json:"full_name"`
	Email          string      `gorm:"type:varchar(100); unique" json:"email"`
	PhoneNumber    string      `gorm:"type:varchar(100)" json:"phone_number"`
	DisplayPicture string      `gorm:"type:varchar(100)" json:"display_picture"`
	Address        string      `gorm:"type:varchar(100)" json:"address"`
	Description    string      `gorm:"varchar(100)" json:"description"`
	SocialMedia    SocialMedia `gorm:"embedded"`
}

//LoginData takes user login information
type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//RegistrationData handles user registration
type RegistrationData struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

//UserRole holds user role object
type UserRole struct {
	FiriModel
	UUID     string `gorm:"type:varchar(100)" json:"uuid"`
	RoleCode string `gorm:"type:varchar(100)" json:"role_code"`
	Role     string `gorm:"type:varchar(100)" json:"role"`
}

//LoginObject show login response data structure
type LoginObject struct {
	UserData    UserData   `json:"user_data"`
	UserRole    []UserRole `json:"user_roles"`
	TokenString string     `json:"token_string"`
}

//RegistrationObject show registration response data structure
type RegistrationObject struct {
	UserData    UserData `json:"user_data"`
	UserRole    UserRole `json:"user_role"`
	TokenString string   `json:"token_string"`
}

//ResponseObject holds response object data
type ResponseObject struct {
	Code    int         `json:"code"`
	Body    interface{} `json:"body"`
	Message string      `json:"message"`
}

//BusinessProfile holds the data structure for enterprise users
type BusinessRegistration struct {
	Email                 string `json:"email"`
	BusinessName          string `json:"business_name"`
	BusinessDescription   string `json:"business_description"`
	BusinessAddress       string `json:"business_address"`
	Website               string `json:"website"`
	CACRegistrationNumber string `json:"cac_registration_number"`
	BusinessType          string `json:"business_type"`
	Industry              string `json:"industry"`
	YearsOfOperation      int    `json:"years_of_operation"`
	CoreBusinessFunctions string `json:"core_business_functions"`
	CoreBusinessTrainings string `json:"core_business_training"`
	OperationalSupport    string `json:"operational_support"`
}

//BusinessProfile holds the data structure for enterprise users
type BusinessProfile struct {
	FiriModel
	UserID                string `json:"user_id"`
	Email                 string `gorm:"type:varchar(100)" json:"email"`
	BusinessName          string `gorm:"type:varchar(100)" json:"business_name"`
	BusinessID            string `gorm:"type:varchar(100)" json:"business_id"`
	BusinessDescription   string `gorm:"type:text" json:"business_description"`
	BusinessAddress       string `gorm:"type:varchar(100)" json:"business_address"`
	Website               string `gorm:"type:varchar(100)" json:"website"`
	CACRegistrationNumber string `gorm:"type:varchar(100); unique" json:"cac_registration_number"`
	BusinessType          string `gorm:"type:varchar(100)" json:"business_type"`
	Industry              string `gorm:"type:varchar(100)" json:"industry"`
	YearsOfOperation      int    `json:"years_of_operation"`
	CoreBusinessFunctions string `gorm:"type:text" json:"core_business_functions"`
	CoreBusinessTrainings string `gorm:"type:text" json:"core_business_training"`
	OperationalSupport    string `gorm:"type:text" json:"operational_support"`
}

//LoginHistory handles Login Data
type LoginHistory struct {
	FiriModel
	Email   string `gorm:"type:varchar(100)" json:"email"`
	Status  string `gorm:"type:varchar(100)" json:"status"`
	Message string `gorm:"type:varchar(100)" json:"message"`
}

//RegistrationMailObject handles registration mail object
type RegistrationMailObject struct {
	Email            string `json:"email"`
	VerificationCode string `json:"code"`
	FullName         string `json:"full_name"`
}

//Industries holds industries object
type Industries struct {
	FiriModel
	Industry string `json:"industry"`
}

//Categories holds categories struct
type Categories struct {
	Category []SubCategories `json:"category"`
}

type TestCategory struct {
	Category map[string]string
}

type CategoryTable struct {
	FiriModel
	Category string `json:"category"`
}

type SubCategoryList struct {
	Category    string        `json:"category"`
	SubCategory []interface{} `json:"subcategories"`
}

type SubCategories struct {
	SubCategory string `json:"name"`
}

//Products holds product regsitration data
type Products struct {
	FiriModel
	ProductName        string `gorm:"type:varchar(100)" json:"product_name"`
	ProductDescription string `json:"product_description"`
	StockQuantity      int    `gorm:"type:varchar(100)" json:"stock_quantity"`
	DisplayImage       string `json:"photo"`
	Price              string `gorm:"type:varchar(100)" json:"price"`
	Industry           string `gorm:"type:varchar(100)" json:"industry"`
	ProductID          string `gorm:"type:varchar(100); unique" json:"product_id"`
	Tags               string `gorm:"type:varchar(100)" json:"tags"`
	BusinessID         string `gorm:"type:varchar(100)" json:"business_id"`
}

//UploadProducts holds the temporary data for uploading products
type UploadProducts struct {
	ProductName        string                `json:"product_name"`
	ProductDescription string                `json:"product_description"`
	StockQuantity      int                   `json:"stock_quantity"`
	DisplayImage       string                `json:"photo"`
	Price              string                `json:"price"`
	Industry           string                `json:"industry"`
	Section            string                `json:"section"`
	ProductID          string                `json:"product_id"`
	Tags               string                `json:"tags"`
	BusinessID         string                `json:"business_id"`
	Images             []UploadProductImages `json:"other_images"`
}

//UploadProductImages holds upload products array for additional products of a business
type UploadProductImages struct {
	ImageURL string `json:"photo"`
}

//Rating handles product ratings
type Rating struct {
	FiriModel
	ProductName  string `gorm:"type:varchar(100)" json:"product_name"`
	ProductID    string `gorm:"type:varchar(100)" json:"product_id"`
	UserName     string `gorm:"type:varchar(100)" json:"user_name"`
	UserRating   int    `json:"user_rating"`
	UserComments string `gorm:"type:varchar(100)" json:"user_comments"`
}

//RateObject handles ratings
type RateObject struct {
	ProductID string `json:"product_id"`
	Rating    int    `json:"rating"`
	Comments  string `json:"comments"`
}

//Tags holds hashtag data
type Tags struct {
	FiriModel
	Tag string `gorm:"type:varchar(100); unique" json:"tag"`
}

//ProductTag holds unique hashtags
type ProductTag struct {
	FiriModel
	ProductID string `gorm:"type:varchar(100)" json:"product_id"`
	TagID     int    `json:"tag_id"`
}

//ProductImage holds product images
type ProductImage struct {
	FiriModel
	ProductID int    `gorm:"type:varchar(100)" json:"product_id"`
	ImageURL  string `json:"image_url"`
}

//Discount discounts
type ProductDiscount struct {
	FiriModel
	ProductName        string  `gorm:"type:varchar(100)" json:"product_name"`
	ProductID          int     `json:"product_id"`
	DiscountPercentage float32 `son:"discount_percent"`
}

//ProductResponseObject response data for when a product is returned
type ProductResponseObject struct {
	ProductID          string          `json:"product_id"`
	ProductName        string          `json:"product_name"`
	ItemLeft           int             `json:"item_left"`
	ProductDescription string          `json:"product_description"`
	Discount           ProductDiscount `json:"discount"`
	Rating             []Rating        `json:"ratings"`
	Tags               []Tags          `json:"tags"`
	Industry           string          `json:"industry"`
	Section            string          `json:"section"`
	Price              string          `json:"price"`
	Images             []ProductImage  `json:"product_images"`
}

type BusinessResponse struct {
	BusinessID            string                  `json:"business_id"`
	Email                 string                  `json:"email"`
	BusinessName          string                  `json:"business_name"`
	BusinessDescription   string                  `json:"business_description"`
	BusinessAddress       string                  `json:"business_address"`
	Website               string                  `json:"website"`
	CACRegistrationNumber string                  `json:"cac_registration_number"`
	BusinessType          string                  `json:"business_type"`
	Industry              string                  `json:"industry"`
	YearsOfOperation      int                     `json:"years_of_operation"`
	CoreBusinessFunctions string                  `json:"core_business_functions"`
	CoreBusinessTrainings string                  `json:"core_business_training"`
	OperationalSupport    string                  `json:"operational_support"`
	Products              []ProductResponseObject `json:"products"`
}

type BusinessResponseObject struct {
	Business BusinessProfile         `json:"business"`
	Products []ProductResponseObject `json:"products"`
}

type ProductQuantity struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CreateProductOrder struct {
	FiriModel
	ProductOrderDetails []ProductQuantity `gorm:"-" json:"products_order_details"`
	UserID              string            `json:"user_id"`
	OrderStatus         string            `json:"order_status"`
	OrderID             string            `json:"order_id"`
	OrderTrackingNumber string            `json:"tracking_number"`
	PaymentReferenceID  string            `json:"payment_id"`
	CustomAddress       bool              `json:"custom_address"`
	Address             string            `json:"address"`
}

type OrderDetails struct {
	FiriModel
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

//CreateOrder holds the data for creating a new order order
type CreateOrder struct {
	FiriModel
	ProductID            int    `json:"product_id"`
	UserID               int    `json:"user_id"`
	OrderQuantity        int    `json:"order_quantity"`
	DeliveryInstructions string `json:"delivery_instructions"`
	Street               string `json:"street"`
	State                string `json:"state"`
	Country              string `json:"country"`
	OrderStatus          string `json:"order_status"`
	OrderID              string `json:"order_id"`
}

type DeliveryAddress struct {
	Street  string `json:"street"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type OrderResponseBody struct {
	Product       ProductResponseObject `json:"product"`
	User          UserData              `json:"user"`
	OrderQuantity int                   `json:"order_quantity"`
	Address       DeliveryAddress       `json:"address"`
}

//EditOrder holds data to cacel order detail
type EditOrder struct {
	OrderID int    `json:"order_id"`
	Message string `json:"edit_message"`
}

type Pay struct {
	FiriModel
	PaymentID     string `json:"payment_id"`
	TransactionID string `json:"transaction_id"`
	BusinessID    string `json:"business_id"`
	ProductID     string `json:"product_id"`
	OrderID       string `json:"order_id"`
	Cost          string `json:"cost"`
}

type Cart struct {
	FiriModel
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type SearchResult struct {
	ProductMatch  bool                     `json:"product_match"`
	Product       []ProductResponseObject  `json:"products"`
	BusinessMatch bool                     `json:"business_match"`
	Businesses    []BusinessResponseObject `json:"business_array"`
	NoMatch       bool                     `json:"no_match"`
}
