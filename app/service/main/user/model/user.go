package model

type AuthType string
const (
	Wechat AuthType = "Wechat"
	QQ  AuthType = "QQ"
	Mobile  AuthType = "Mobile"
)

type User struct {
	Email string   `json:"email" gorm:"email"`
	Phone string	`json:"phone" gorm:"phone"`
	Username string	`json:"username" gorm:"username"`
	Gender string	`json:"gender" gorm:"gender"`
	nickname string	`json:"nickname" gorm:"nickname"`
	headPicurl string	`json:"headPicurl" gorm:"head_pic_url"`
	password string	`json:"password" gorm:"password"`
	wxOpenId string	`json:"wxOpenId" gorm:"wx_open_id"`
	unionId string	`json:"unionId" gorm:"union_id"`
	qqOpenId string	`json:"qqOpenId" gorm:"qq_open_id"`
	authType AuthType	`json:"authType" gorm:"type:ENUM('Wechat', 'QQ','Mobile')""`
	registerTime string	`json:"registerTime" gorm:"register_time"`
	lastLoginTime string	`json:"lastLoginTime" gorm:"last_login_time"`
}

func (u *User) TableName() string {
	return "user"
}

type UserAddress struct {
	UserId int64	`json:"userId" gorm:"user_id"`
	Consignee string	`json:"consignee" gorm:"consignee"`
	Email string	`json:"email" gorm:"email"`
	CountryId string	`json:"countryId" gorm:"country_id"`
	Province string	`json:"province" gorm:"province"`
	City string	`json:"city" gorm:"city"`
	District string	`json:"district" gorm:"district"`
	Location string	`json:"location" gorm:"location"`
	FullAddress string	`json:"fullAddress" gorm:"full_address"`
	ZipCode string	`json:"zipCode" gorm:"zip_code"`
	Phone string	`json:"phone" gorm:"phone"`
	IdCode string	`json:"idCode" gorm:"id_code"`
	DefaultAddress int	`json:"defaultAddress" gorm:"default_address"`
}

func (u *UserAddress) TableName() string {
	return "user"
}

type UserFavorite struct {
	UserId int64	`json:"userId" gorm:"user_id"`
	ProductId int64	`json:"productId" gorm:"product_id"`
	ProductName int64	`json:"productName" gorm:"product_name"`
	ProductImgUrl int64	`json:"productImgUrl" gorm:"product_img_url"`
	Price int64	`json:"price" gorm:"price"`
	MarketPrice int64	`json:"marketPrice" gorm:"market_price"`
}

func (u *UserFavorite) TableName() string {
	return "user_favorite"
}


type UserProduct struct {
	UserId int64	`json:"userId" gorm:"user_id"`
	ProductId int64	`json:"productId" gorm:"product_id"`
}
func (u *UserProduct) TableName() string {
	return "user_favorite"
}


type UserRefreshToken struct {
	UserId int64	`json:"userId" gorm:"user_id"`
	Phone string	`json:"phone" gorm:"phone"`
	RefreshToken string	`json:"refreshToken" gorm:"refresh_token"`
}
func (u *UserRefreshToken) TableName() string {
	return "user_refresh_token"
}


type UserWantBuyProduct struct {
	UserId int64	`json:"userId" gorm:"user_id"`
	ProductId int64	`json:"productId" gorm:"product_id"`
}

func (u *UserWantBuyProduct) TableName() string {
	return "user_want_buy_product"
}
