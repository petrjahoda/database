package database

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Alarm struct {
	gorm.Model
	Name          string `gorm:"unique"`
	WorkplaceID   int
	SqlCommand    string
	MessageHeader string
	MessageText   string
	Recipients    string
	Url           string
	Pdf           string
}

type AlarmRecord struct {
	gorm.Model
	DateTimeStart     time.Time `gorm:"uniqueIndex:unique_alarm_data"`
	DateTimeEnd       sql.NullTime
	DateTimeProcessed sql.NullTime
	AlarmID           int `gorm:"uniqueIndex:unique_alarm_data"`
	WorkplaceID       int `gorm:"uniqueIndex:unique_alarm_data"`
	UserID            sql.NullInt32
}

type SystemRecord struct {
	gorm.Model
	DatabaseSizeInMegaBytes     float32
	DatabaseGrowthInMegaBytes   float32
	DiscFreeSizeInMegaBytes     float32
	EstimatedDiscFreeSizeInDays float32
}

type State struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Color string
	Note  string
}

type PageCount struct {
	gorm.Model
	PageName string `gorm:"unique"`
	Count    int
	Note     string
}

type StateRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_state_data"`
	StateID       int       `gorm:"uniqueIndex:unique_state_data"`
	WorkplaceID   int       `gorm:"uniqueIndex:unique_state_data"`
	Note          string
}

type UserRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_user_data"`
	DateTimeEnd   sql.NullTime
	OrderRecordID int `gorm:"uniqueIndex:unique_user_data"`
	WorkplaceID   int `gorm:"uniqueIndex:unique_user_data"`
	UserID        int `gorm:"uniqueIndex:unique_user_data"`
	Note          string
}

type DowntimeRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_downtime_data"`
	DateTimeEnd   sql.NullTime
	WorkplaceID   int `gorm:"uniqueIndex:unique_downtime_data"`
	DowntimeID    int `gorm:"uniqueIndex:unique_downtime_data"`
	OrderRecordID sql.NullInt32
	UserID        sql.NullInt32
	Note          string
}

type BreakdownRecord struct {
	gorm.Model
	DateTimeStart time.Time `gorm:"uniqueIndex:unique_breakdown_data"`
	DateTimeEnd   sql.NullTime
	BreakdownID   int `gorm:"uniqueIndex:unique_breakdown_data"`
	WorkplaceID   int `gorm:"uniqueIndex:unique_breakdown_data"`
	UserID        int
	Note          string
}

type FaultRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_fault_data"`
	OrderRecordID int
	FaultID       int `gorm:"uniqueIndex:unique_fault_data"`
	WorkplaceID   int `gorm:"uniqueIndex:unique_fault_data"`
	UserID        int
	Count         int
	Note          string
}

type PackageRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_package_data"`
	OrderRecordID int
	PackageID     int `gorm:"uniqueIndex:unique_package_data"`
	WorkplaceID   int `gorm:"uniqueIndex:unique_package_data"`
	UserID        int
	Count         int
	Note          string
}

type PartRecord struct {
	gorm.Model
	DateTime      time.Time `gorm:"uniqueIndex:unique_part_data"`
	OrderRecordID int
	PartID        int `gorm:"uniqueIndex:unique_part_data"`
	WorkplaceID   int `gorm:"uniqueIndex:unique_part_data"`
	UserID        int
	Count         int
	Note          string
}

type OrderRecord struct {
	gorm.Model
	DateTimeStart   time.Time `gorm:"uniqueIndex:unique_order_data"`
	DateTimeEnd     sql.NullTime
	OrderID         int `gorm:"uniqueIndex:unique_order_data"`
	OperationID     int `gorm:"uniqueIndex:unique_order_data"`
	WorkplaceID     int `gorm:"uniqueIndex:unique_order_data"`
	UserId          sql.NullInt32
	WorkplaceModeID int
	WorkshiftID     int
	AverageCycle    float32
	Cavity          int
	CountOk         int
	CountNok        int
	Note            string
}

type Operation struct {
	gorm.Model
	Name    string `gorm:"unique"`
	OrderID int
	Barcode string
	Note    string
}

type WorkplaceSection struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}
type Order struct {
	gorm.Model
	Name            string `gorm:"unique"`
	ProductID       sql.NullInt32
	WorkplaceID     sql.NullInt32
	Barcode         string
	DateTimeRequest sql.NullTime
	Cavity          int
	CountRequest    int
	Note            string
}
type Product struct {
	gorm.Model
	Name             string `gorm:"unique"`
	Barcode          string
	CycleTime        int
	DownTimeDuration time.Duration
	Note             string
}

type Part struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Barcode string
	Note    string
}

type Workplace struct {
	gorm.Model
	Name               string `gorm:"unique"`
	Code               string
	WorkplaceSectionID int
	WorkplaceModeID    int
	Note               string
}

type WorkplacePort struct {
	gorm.Model
	Name         string `gorm:"uniqueIndex:unique_workplace_port_data"`
	DevicePortID int    `gorm:"uniqueIndex:unique_workplace_port_data"`
	StateID      sql.NullInt32
	WorkplaceID  int `gorm:"uniqueIndex:unique_workplace_port_data"`
	Color        string
	CounterOK    bool
	CounterNOK   bool
	HighValue    float32
	LowValue     float32
	Note         string
}

type WorkplaceMode struct {
	gorm.Model
	Name             string `gorm:"unique"`
	DowntimeDuration time.Duration
	PoweroffDuration time.Duration
	Note             string
}

type WorkplaceWorkshift struct {
	gorm.Model
	WorkplaceID int
	WorkshiftID int
}

type Workshift struct {
	gorm.Model
	Name           string `gorm:"unique"`
	WorkshiftStart time.Time
	WorkshiftEnd   time.Time
	Note           string
}

type User struct {
	gorm.Model
	FirstName  string `gorm:"uniqueIndex:username_data"`
	SecondName string `gorm:"uniqueIndex:username_data"`
	UserRoleID int
	UserTypeID int
	Barcode    string
	Email      string `gorm:"uniqueIndex:username_data"`
	Password   string
	Phone      string
	Pin        string
	Position   string
	Rfid       string
	Locale     string
	Note       string
}

type UserRole struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type UserType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Downtime struct {
	gorm.Model
	Name           string `gorm:"unique"`
	DowntimeTypeID int
	Barcode        string
	Color          string
	Note           string
}

type DowntimeType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Breakdown struct {
	gorm.Model
	Name            string `gorm:"unique"`
	BreakdownTypeID int
	Barcode         string
	Color           string
	Note            string
}

type BreakdownType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Fault struct {
	gorm.Model
	Name        string `gorm:"unique"`
	FaultTypeID int
	Barcode     string
	Note        string
}

type FaultType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Package struct {
	gorm.Model
	Name          string `gorm:"unique"`
	OrderID       int
	PackageTypeID int
	Barcode       string
	Note          string
}

type PackageType struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Count int
	Note  string
}

type DevicePortType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Setting struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Value   string
	Enabled bool
	Note    string
}

type DeviceType struct {
	gorm.Model
	Name string `gorm:"unique"`
	Note string
}

type Device struct {
	gorm.Model
	Name         string `gorm:"unique"`
	DeviceTypeID int
	Activated    bool
	IpAddress    string `gorm:"unique"`
	MacAddress   string
	Settings     string
	TypeName     string
	Note         string
}

type DeviceWorkplaceRecord struct {
	gorm.Model
	DeviceID    int `gorm:"uniqueIndex:device_workplace_data"`
	WorkplaceID int `gorm:"uniqueIndex:device_workplace_data"`
}

type DevicePort struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex:device_port_data"`
	DeviceID         int    `gorm:"uniqueIndex:device_port_data"`
	DevicePortTypeID int
	PortNumber       int
	PlcDataType      string
	PlcDataAddress   string
	Settings         string
	Unit             string
	Virtual          bool `gorm:"default:false"`
	Note             string
}

type DevicePortAnalogRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_analog_data"`
	DevicePortID int       `gorm:"uniqueIndex:unique_analog_data"`
	Data         float32
}

type DevicePortDigitalRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_digital_data"`
	DevicePortID int       `gorm:"uniqueIndex:unique_digital_data"`
	Data         int
}

type DevicePortSerialRecord struct {
	ID           int       `gorm:"primaryKey"`
	DateTime     time.Time `gorm:"uniqueIndex:unique_serial_data"`
	DevicePortID int       `gorm:"uniqueIndex:unique_serial_data"`
	Data         float32
}

type Locale struct {
	gorm.Model
	Name string `gorm:"unique"`
	CsCZ string
	DeDE string
	EnUS string
	EsES string
	FrFR string
	ItIT string
	PlPL string
	PtPT string
	SkSK string
	RuRU string
}

type WebUserRecord struct {
	gorm.Model
	UserEmail string
	WebPage   string
	DateTime  time.Time
}
