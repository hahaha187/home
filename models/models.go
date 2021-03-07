package models

import (
	"time"

	"github.com/astaxie/beego/client/orm"
	//
	_ "github.com/go-sql-driver/mysql"
)

//User biao
type User struct {
	ID        int           `json:"user_id"`                    //用户编号
	Name      string        `orm:"size(32);unique" json:"name"` //用户昵称
	Password  string        `orm:"size(128)" json:"password"`   //用户密码加密的
	Mobile    string        `orm:"size(11)" json:"mobile"`      //手机号
	RealName  string        `orm:"size(32)" json:"real_name"`   //真实姓名
	IDCard    string        `orm:"size(20)" json:"id_card"`     //身份证号
	AvatarURL string        `orm:"size(256)" json:"avatar_url"` //用户头像路径
	Houses    []*House      `orm:"reverse(many)" json:"houses"` //用户发布的房屋信息
	Orders    []*OrderHouse `orm:"reverse(many)" json:"orders"` //用户下的订单
}

//House 户层信息 table_name = house
type House struct {
	ID            int           `json:"house_id"`                                          //房屋编号
	User          *User         `orm:"rel(fk)" json:"user_id"`                             //房屋主人的用户编号
	Area          *Area         `orm:"rel(fk)" json:"area_id"`                             //归属地的区域编号
	Title         string        `orm:"size(64)" json:"title"`                              //房屋标题
	Price         int           `orm:"default(0)" json:"price"`                            //单价，单位：分
	Address       string        `orm:"size(512)" orm:"default('')" json:"address"`         //地址
	RoomCount     int           `orm:"default(1)" json:"room_count"`                       //房间数目
	Acreage       int           `orm:"default(0)" json:"acreage"`                          //房屋总面积
	Unit          string        `orm:"size(32)" orm:"default('')" json:"unit"`             //房屋单元，如 几室几厅
	Capacity      int           `orm:"default(1)" json:"capacity"`                         //房屋容纳的总人数
	Beds          string        `orm:"size(64)" orm:"default('')" json:"beds"`             //房屋床铺的位置
	Deposit       int           `orm:"default(0)" json:"deposit"`                          //押金
	MinDays       int           `orm:"default(1)" json:"min_days"`                         //最少入住天数
	MaxDays       int           `orm:"default(0)" json:"max_days"`                         //最多住天数 0表示不限制
	OrderCount    int           `orm:"default(0)" json:"order_count"`                      //预定完成的该房屋的订单数
	IndexImageURL string        `orm:"size(256)" orm:"default('')" json:"index_image_url"` //房屋主图片路径
	Facilities    []*Facility   `orm:"reverse(many)" json:"facilities"`                    //房屋设施
	Images        []*HouseImage `orm:"reverse(many)" json:"img_urls"`                      //房屋的图片
	Orders        []*OrderHouse `orm:"reverse(many)" json:"orders"`                        //房屋的订单
	Ctime         time.Time     `orm:"auto_now_add;type(datetime)" json:"ctime"`
}

//HomeMax 首页最高展示的房屋数量
var HomeMax int = 5

//HouseNum 房屋列表页面每页显示条目数
var HouseNum int = 2

//Area 区域信息 table_name = area
type Area struct {
	ID   int    `json:"aid"`                  //区域编号
	Name string `orm:"size(32)" json:"aname"` //区域名字
}

//Facility 设施信息 table_name = "facility"
type Facility struct {
	ID     int      `json:"fid"`     //设施编号
	Name   string   `orm:"size(32)"` //设施名字
	Houses []*House `orm:"rel(m2m)"` //都有哪些房屋有此设施
}

//HouseImage 房屋图片 table_name = "house_image"
type HouseImage struct {
	ID    int    `json:"house_image_id"`         //图片id
	URL   string `orm:"size(256)" json:"url"`    //图片url
	House *House `orm:"rel(fk)" json:"house_id"` //图片所属房屋编号
}

/* 订单状态常量 */
const (
	OrderStatusWaitAccept  = "WAIT_ACCEPT"  //待接单
	OrderStatusWaitPayment = "WAIT_PAYMENT" //待支付
	OrderStatusPaid        = "PAID"         //已支付
	OrderStatusWaitComment = "COMMENT"      //待评价
	OrderStatusComplete    = "COMPLETE"     //已完成
	OrderStatusCanceled    = "CANCELED"     //已取消
	OrderStatusRejected    = "REJECTED"     //已拒单
)

//OrderHouse 订单 table_name = order
type OrderHouse struct {
	ID         int       `json:"order_id"`               //订单编号
	User       *User     `orm:"rel(fk)" json:"user_id"`  //下单的用户编号
	House      *House    `orm:"rel(fk)" json:"house_id"` //预定的房间编号
	BeginData  time.Time `orm:"type(datetime)"`          //预定的起始时间
	EndData    time.Time `orm:"type(datetime)"`          //预定的结束时间
	Days       int       //预定总天数
	HousePrice int       //房屋的单价
	Amount     int       //订单总金额
	Status     string    `orm:"default(WAIT_ACCEPT)"`                     //订单状态
	Comment    string    `orm:"size(512)"`                                //订单评论
	Ctime      time.Time `orm:"auto_now_add;type(datetime)" json:"ctime"` //
}

func init() {
	// 设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:lhy871601318@tcp(127.0.0.1:3306)/loveHome2?charset=utf8")

	// register model映射model数据
	orm.RegisterModel(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse))

	// create table 生成表
	//orm.RunSyncdb("default", false, true)
}
