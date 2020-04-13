package model

import "angrymiao-go/punk/model"

type AuthType string
const (
	Wechat AuthType = "Wechat"
	QQ  AuthType = "QQ"
	Mobile  AuthType = "Mobile"
)

type User struct {
	model.BaseModel
	Email string   `json:"email" gorm:"email"`
	Phone string	`json:"phone" gorm:"phone"`
	Username string	`json:"username" gorm:"username"`
	Gender string	`json:"gender" gorm:"gender"`
	Nickname string	`json:"nickname" gorm:"nickname"`
	HeadPicurl string	`json:"headPicurl" gorm:"head_pic_url"`
	Password string	`json:"password" gorm:"password"`
	WxOpenId string	`json:"wxOpenId" gorm:"wx_open_id"`
	UnionId string	`json:"unionId" gorm:"union_id"`
	QQOpenId string	`json:"qqOpenId" gorm:"qq_open_id"`
	AuthType AuthType	`json:"authType" gorm:"type:ENUM('Wechat', 'QQ','Mobile')""`
	RegisterTime string	`json:"registerTime" gorm:"register_time"`
	LastLoginTime string	`json:"lastLoginTime" gorm:"last_login_time"`
}

func (u *User)ToUserViewResponse() (userViewResponse *UserViewResponse) {
	userViewResponse = &UserViewResponse{
		ID:           u.ID,
		Nickname:     u.Nickname,
		Phone:        u.Phone,
		AvatarUrl:    u.HeadPicurl,
		Balance:      "0.00",
		Gender:       0,
		DateJoined:   u.CreateTime,
		AuthType:     u.AuthType,
	}
	return
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
