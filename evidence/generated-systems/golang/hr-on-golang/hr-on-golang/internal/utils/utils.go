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
	"hr-on-golang/internal/model"
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
    GetDB().AutoMigrate(&model.Organization{})
    GetDB().AutoMigrate(&model.Department{})
    GetDB().AutoMigrate(&model.Location{})
    GetDB().AutoMigrate(&model.CostCenter{})
    GetDB().AutoMigrate(&model.JobFamily{})
    GetDB().AutoMigrate(&model.JobProfile{})
    GetDB().AutoMigrate(&model.Competency{})
    GetDB().AutoMigrate(&model.Position{})
    GetDB().AutoMigrate(&model.Employee{})
    GetDB().AutoMigrate(&model.EmploymentAssignment{})
    GetDB().AutoMigrate(&model.EmploymentContract{})
    GetDB().AutoMigrate(&model.WorkSchedule{})
    GetDB().AutoMigrate(&model.WorkShift{})
    GetDB().AutoMigrate(&model.ScheduleException{})
    GetDB().AutoMigrate(&model.CompensationPackage{})
    GetDB().AutoMigrate(&model.SalaryComponent{})
    GetDB().AutoMigrate(&model.BonusPlan{})
    GetDB().AutoMigrate(&model.EquityGrant{})
    GetDB().AutoMigrate(&model.BenefitPlan{})
    GetDB().AutoMigrate(&model.BenefitEnrollment{})
    GetDB().AutoMigrate(&model.Dependent{})
    GetDB().AutoMigrate(&model.PayrollCalendar{})
    GetDB().AutoMigrate(&model.PayrollRun{})
    GetDB().AutoMigrate(&model.PayrollItem{})
    GetDB().AutoMigrate(&model.TaxWithholding{})
    GetDB().AutoMigrate(&model.PaymentMethod{})
    GetDB().AutoMigrate(&model.Timesheet{})
    GetDB().AutoMigrate(&model.TimeEntry{})
    GetDB().AutoMigrate(&model.Approval{})
    GetDB().AutoMigrate(&model.LeavePolicy{})
    GetDB().AutoMigrate(&model.LeaveRequest{})
    GetDB().AutoMigrate(&model.PerformanceCycle{})
    GetDB().AutoMigrate(&model.Goal{})
    GetDB().AutoMigrate(&model.PerformanceReview{})
    GetDB().AutoMigrate(&model.CompetencyRating{})
    GetDB().AutoMigrate(&model.TrainingCourse{})
    GetDB().AutoMigrate(&model.TrainingEnrollment{})
    GetDB().AutoMigrate(&model.Certification{})
    GetDB().AutoMigrate(&model.JobRequisition{})
    GetDB().AutoMigrate(&model.Candidate{})
    GetDB().AutoMigrate(&model.JobApplication{})
    GetDB().AutoMigrate(&model.Interview{})
    GetDB().AutoMigrate(&model.Screening{})
    GetDB().AutoMigrate(&model.Offer{})
    GetDB().AutoMigrate(&model.OnboardingTask{})
    GetDB().AutoMigrate(&model.BackgroundCheck{})
    GetDB().AutoMigrate(&model.Document{})
    GetDB().AutoMigrate(&model.Policy{})
    GetDB().AutoMigrate(&model.PolicyAcknowledgement{})
    GetDB().AutoMigrate(&model.Termination{})
    GetDB().AutoMigrate(&model.WorkAuthorization{})
    GetDB().AutoMigrate(&model.BankAccount{})
}