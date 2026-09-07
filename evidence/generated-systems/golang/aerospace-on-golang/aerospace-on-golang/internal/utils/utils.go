package utils

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
	"strconv"
	"log"
	"aerospace-on-golang/internal/model"
)

//----------------------------------------------------------------------------
// global level var declarations
//----------------------------------------------------------------------------
var db *gorm.DB

type RequestResult struct {
    Success  	bool
	Msg 		string
	Call 		string
	Data 		interface{}
}

//----------------------------------------------------------------------------
// function initialze the database and environment
//----------------------------------------------------------------------------
func InitializeEnvironment() {

	//----------------------------------------------------------------------------
    // load .env file
    //----------------------------------------------------------------------------
    err := godotenv.Load()

    if err != nil {
        log.Println("No .env file found. Using environment variables.")
    }

    fmt.Println("DB_USER_NAME =", os.Getenv("DB_USER_NAME"))
    fmt.Println("DB_PASSWORD SET =", os.Getenv("DB_PASSWORD") != "")
    fmt.Println("DB_HOST =", os.Getenv("DB_HOST"))
    fmt.Println("DB_PORT =", os.Getenv("DB_PORT"))
    fmt.Println("DB_NAME =", os.Getenv("DB_NAME"))

	//----------------------------------------------------------------------------
	// Open the mysql database and initialize the ORM
	//----------------------------------------------------------------------------

    var dsn string
    switch dbVersion := os.Getenv("DB_VERSION"); dbVersion {
	    case "postgres":
		    // example: host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
		    dsn = fmt.Sprintf( "host=%s user=%s password=%s dbname=%s port=%s %s", 
						os.Getenv("DB_HOST"),
                        os.Getenv("DB_USER_NAME"),
                        os.Getenv("DB_PASSWORD"),
                        os.Getenv("DB_NAME"),
                        os.Getenv("DB_PORT"),
                        os.Getenv("DB_ARGS") )
	    case "sqlite" :
		    dsn = fmt.Sprintf( "%s %s", os.Getenv("DB_NAME"), os.Getenv("DB_ARGS") )

	    case "sqlserver" :
		    // example: sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm
           dsn = fmt.Sprintf( "sqlserver://%s:%s@%s:%s?database=%s&%s", 
    						os.Getenv("DB_USER_NAME"), 
    						os.Getenv("DB_PASSWORD"), 
    						os.Getenv("DB_HOST"), 
    						os.Getenv("DB_PORT"), 
    						os.Getenv("DB_NAME"), 
    						os.Getenv("DB_ARGS") )
	    default:
		    // default to mysql
		    // example: gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local
            dsn = fmt.Sprintf( "%s:%s@tcp(%s:%s)/%s?%s", 
        						os.Getenv("DB_USER_NAME"), 
        						os.Getenv("DB_PASSWORD"), 
        						os.Getenv("DB_HOST"), 
        						os.Getenv("DB_PORT"), 
        						os.Getenv("DB_NAME"), 
        						os.Getenv("DB_ARGS") )
	}
	
	fmt.Println( "Connecting to the database using DSN ", dsn )
		
	disableFKConstraint,_ := strconv.ParseBool(os.Getenv("DB_DISABLE_FK_CONSTRAINTS"));
	
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: disableFKConstraint,
    })

    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
	//----------------------------------------------------------------------------
	// Handle schema creation or update
	//----------------------------------------------------------------------------
	AutoMigrate()
}

//----------------------------------------------------------------------------
// Returns the database instance
//----------------------------------------------------------------------------
func GetDB() *gorm.DB {
	return db
}

//----------------------------------------------------------------------------
// Parses the body of the HTTP Request and unmarshal it via JSON into the
// provided interface
//----------------------------------------------------------------------------
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

//----------------------------------------------------------------------------
// Handler to AutoMigrate schema to gorm for each model struct
//----------------------------------------------------------------------------
func AutoMigrate() {
    GetDB().AutoMigrate(&model.AerospaceManufacturer{})
    GetDB().AutoMigrate(&model.AircraftProgram{})
    GetDB().AutoMigrate(&model.AircraftFamily{})
    GetDB().AutoMigrate(&model.AircraftModel{})
    GetDB().AutoMigrate(&model.EngineType{})
    GetDB().AutoMigrate(&model.AircraftVariant{})
    GetDB().AutoMigrate(&model.AvionicsSuite{})
    GetDB().AutoMigrate(&model.APU{})
    GetDB().AutoMigrate(&model.LandingGear{})
    GetDB().AutoMigrate(&model.AircraftOption{})
    GetDB().AutoMigrate(&model.AircraftPackage{})
    GetDB().AutoMigrate(&model.Supplier{})
    GetDB().AutoMigrate(&model.Component_{})
    GetDB().AutoMigrate(&model.Plant{})
    GetDB().AutoMigrate(&model.ProductionLine{})
    GetDB().AutoMigrate(&model.WorkCenter{})
    GetDB().AutoMigrate(&model.ProductionOrder{})
    GetDB().AutoMigrate(&model.BuildSchedule{})
    GetDB().AutoMigrate(&model.Warehouse{})
    GetDB().AutoMigrate(&model.InventoryItem{})
    GetDB().AutoMigrate(&model.Operator{})
    GetDB().AutoMigrate(&model.AircraftOrder{})
    GetDB().AutoMigrate(&model.Quote{})
    GetDB().AutoMigrate(&model.PurchaseAgreement{})
    GetDB().AutoMigrate(&model.Aircraft{})
    GetDB().AutoMigrate(&model.Registration{})
    GetDB().AutoMigrate(&model.Warranty{})
    GetDB().AutoMigrate(&model.CabinLayout{})
    GetDB().AutoMigrate(&model.MROFacility{})
    GetDB().AutoMigrate(&model.MaintenanceAppointment{})
    GetDB().AutoMigrate(&model.MaintenanceWorkOrder{})
    GetDB().AutoMigrate(&model.AirworthinessDirective{})
    GetDB().AutoMigrate(&model.ServiceBulletin{})
    GetDB().AutoMigrate(&model.ConnectedAircraft{})
    GetDB().AutoMigrate(&model.FlightHealthEvent{})
    GetDB().AutoMigrate(&model.SoftwareLoad{})
    GetDB().AutoMigrate(&model.TypeCertificate{})
    GetDB().AutoMigrate(&model.ProductionCertificate{})
    GetDB().AutoMigrate(&model.SalesRegion{})
    GetDB().AutoMigrate(&model.SalesCampaign{})
}