package test

import ( 
	"testing"
    dao "hr-on-golang/internal/dao"
	"hr-on-golang/internal/model"
	"hr-on-golang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestOrganizationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Organization
	//----------------------------------------------------------------------------
	OrganizationObj := model.Organization                                                                                                                            {Name:"test value for Name",LegalName:"test value for LegalName",RegistrationCountry:"test value for RegistrationCountry",Website:"test value for Website"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOrganizationRequestResult := dao.CreateOrganization( OrganizationObj )
	
	if createOrganizationRequestResult.Success == false {
		t.Errorf(createOrganizationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Organization success...")
	}
	
	createOrganizationObj,_ := createOrganizationRequestResult.Data. (model.Organization)

	// --------------------------------------------------------------
	// Check Organization Obj ID
	// --------------------------------------------------------------	
	if createOrganizationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Organization" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOrganizationRequestResult := dao.GetOrganization( uint64(createOrganizationObj.ID) )
	
	if getOrganizationRequestResult.Success == false {
		t.Errorf(getOrganizationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Organization success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOrganizationObj,_ := getOrganizationRequestResult.Data. (model.Organization)
	compareOrganization := cmp.Equal(createOrganizationObj.ID, getOrganizationObj.ID)
	
	if  compareOrganization == false	{
		t.Errorf( "Created Organization object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOrganizationRequestResult := dao.GetAllOrganization()

	if getAllOrganizationRequestResult.Success == false {
			t.Errorf(getAllOrganizationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Organization success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOrganizationObj []model.Organization = getAllOrganizationRequestResult.Data. ([]model.Organization)
		
	equalOrganization := cmp.Equal(createOrganizationObj.ID, getAllOrganizationObj[len(getAllOrganizationObj)-1].ID)
		
	if equalOrganization == false {
		t.Errorf( "Created object is not equal to the last entry in Organization[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Organization
	// --------------------------------------------------------------	
	deleteOrganizationRequestResult := dao.DeleteOrganization(uint64(createOrganizationObj.ID))

	if deleteOrganizationRequestResult.Success == false {
			t.Errorf(deleteOrganizationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Organization success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOrganizationRequestResult = dao.GetOrganization( uint64(createOrganizationObj.ID) )
	
	if getOrganizationRequestResult.Success == true {
		t.Errorf(getOrganizationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDepartmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Department
	//----------------------------------------------------------------------------
	DepartmentObj := model.Department                                                            {Name:"test value for Name",Code:"test value for Code"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDepartmentRequestResult := dao.CreateDepartment( DepartmentObj )
	
	if createDepartmentRequestResult.Success == false {
		t.Errorf(createDepartmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Department success...")
	}
	
	createDepartmentObj,_ := createDepartmentRequestResult.Data. (model.Department)

	// --------------------------------------------------------------
	// Check Department Obj ID
	// --------------------------------------------------------------	
	if createDepartmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Department" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDepartmentRequestResult := dao.GetDepartment( uint64(createDepartmentObj.ID) )
	
	if getDepartmentRequestResult.Success == false {
		t.Errorf(getDepartmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Department success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDepartmentObj,_ := getDepartmentRequestResult.Data. (model.Department)
	compareDepartment := cmp.Equal(createDepartmentObj.ID, getDepartmentObj.ID)
	
	if  compareDepartment == false	{
		t.Errorf( "Created Department object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDepartmentRequestResult := dao.GetAllDepartment()

	if getAllDepartmentRequestResult.Success == false {
			t.Errorf(getAllDepartmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Department success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDepartmentObj []model.Department = getAllDepartmentRequestResult.Data. ([]model.Department)
		
	equalDepartment := cmp.Equal(createDepartmentObj.ID, getAllDepartmentObj[len(getAllDepartmentObj)-1].ID)
		
	if equalDepartment == false {
		t.Errorf( "Created object is not equal to the last entry in Department[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Department
	// --------------------------------------------------------------	
	deleteDepartmentRequestResult := dao.DeleteDepartment(uint64(createDepartmentObj.ID))

	if deleteDepartmentRequestResult.Success == false {
			t.Errorf(deleteDepartmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Department success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDepartmentRequestResult = dao.GetDepartment( uint64(createDepartmentObj.ID) )
	
	if getDepartmentRequestResult.Success == true {
		t.Errorf(getDepartmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLocationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Location
	//----------------------------------------------------------------------------
	LocationObj := model.Location                                                                            {Name:"test value for Name",Address:new Address(),Timezone:"test value for Timezone"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLocationRequestResult := dao.CreateLocation( LocationObj )
	
	if createLocationRequestResult.Success == false {
		t.Errorf(createLocationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Location success...")
	}
	
	createLocationObj,_ := createLocationRequestResult.Data. (model.Location)

	// --------------------------------------------------------------
	// Check Location Obj ID
	// --------------------------------------------------------------	
	if createLocationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Location" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLocationRequestResult := dao.GetLocation( uint64(createLocationObj.ID) )
	
	if getLocationRequestResult.Success == false {
		t.Errorf(getLocationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Location success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLocationObj,_ := getLocationRequestResult.Data. (model.Location)
	compareLocation := cmp.Equal(createLocationObj.ID, getLocationObj.ID)
	
	if  compareLocation == false	{
		t.Errorf( "Created Location object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLocationRequestResult := dao.GetAllLocation()

	if getAllLocationRequestResult.Success == false {
			t.Errorf(getAllLocationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Location success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLocationObj []model.Location = getAllLocationRequestResult.Data. ([]model.Location)
		
	equalLocation := cmp.Equal(createLocationObj.ID, getAllLocationObj[len(getAllLocationObj)-1].ID)
		
	if equalLocation == false {
		t.Errorf( "Created object is not equal to the last entry in Location[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Location
	// --------------------------------------------------------------	
	deleteLocationRequestResult := dao.DeleteLocation(uint64(createLocationObj.ID))

	if deleteLocationRequestResult.Success == false {
			t.Errorf(deleteLocationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Location success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLocationRequestResult = dao.GetLocation( uint64(createLocationObj.ID) )
	
	if getLocationRequestResult.Success == true {
		t.Errorf(getLocationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCostCenterCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CostCenter
	//----------------------------------------------------------------------------
	CostCenterObj := model.CostCenter                                                            {Code:"test value for Code",Name:"test value for Name"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCostCenterRequestResult := dao.CreateCostCenter( CostCenterObj )
	
	if createCostCenterRequestResult.Success == false {
		t.Errorf(createCostCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Create CostCenter success...")
	}
	
	createCostCenterObj,_ := createCostCenterRequestResult.Data. (model.CostCenter)

	// --------------------------------------------------------------
	// Check CostCenter Obj ID
	// --------------------------------------------------------------	
	if createCostCenterObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CostCenter" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCostCenterRequestResult := dao.GetCostCenter( uint64(createCostCenterObj.ID) )
	
	if getCostCenterRequestResult.Success == false {
		t.Errorf(getCostCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Get CostCenter success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCostCenterObj,_ := getCostCenterRequestResult.Data. (model.CostCenter)
	compareCostCenter := cmp.Equal(createCostCenterObj.ID, getCostCenterObj.ID)
	
	if  compareCostCenter == false	{
		t.Errorf( "Created CostCenter object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCostCenterRequestResult := dao.GetAllCostCenter()

	if getAllCostCenterRequestResult.Success == false {
			t.Errorf(getAllCostCenterRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CostCenter success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCostCenterObj []model.CostCenter = getAllCostCenterRequestResult.Data. ([]model.CostCenter)
		
	equalCostCenter := cmp.Equal(createCostCenterObj.ID, getAllCostCenterObj[len(getAllCostCenterObj)-1].ID)
		
	if equalCostCenter == false {
		t.Errorf( "Created object is not equal to the last entry in CostCenter[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CostCenter
	// --------------------------------------------------------------	
	deleteCostCenterRequestResult := dao.DeleteCostCenter(uint64(createCostCenterObj.ID))

	if deleteCostCenterRequestResult.Success == false {
			t.Errorf(deleteCostCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CostCenter success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCostCenterRequestResult = dao.GetCostCenter( uint64(createCostCenterObj.ID) )
	
	if getCostCenterRequestResult.Success == true {
		t.Errorf(getCostCenterRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestJobFamilyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for JobFamily
	//----------------------------------------------------------------------------
	JobFamilyObj := model.JobFamily                                                            {Name:"test value for Name",Description:"test value for Description"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createJobFamilyRequestResult := dao.CreateJobFamily( JobFamilyObj )
	
	if createJobFamilyRequestResult.Success == false {
		t.Errorf(createJobFamilyRequestResult.Msg)
	} else {
		fmt.Println("Check Create JobFamily success...")
	}
	
	createJobFamilyObj,_ := createJobFamilyRequestResult.Data. (model.JobFamily)

	// --------------------------------------------------------------
	// Check JobFamily Obj ID
	// --------------------------------------------------------------	
	if createJobFamilyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for JobFamily" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getJobFamilyRequestResult := dao.GetJobFamily( uint64(createJobFamilyObj.ID) )
	
	if getJobFamilyRequestResult.Success == false {
		t.Errorf(getJobFamilyRequestResult.Msg)
	} else {
		fmt.Println("Check Get JobFamily success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getJobFamilyObj,_ := getJobFamilyRequestResult.Data. (model.JobFamily)
	compareJobFamily := cmp.Equal(createJobFamilyObj.ID, getJobFamilyObj.ID)
	
	if  compareJobFamily == false	{
		t.Errorf( "Created JobFamily object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllJobFamilyRequestResult := dao.GetAllJobFamily()

	if getAllJobFamilyRequestResult.Success == false {
			t.Errorf(getAllJobFamilyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll JobFamily success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllJobFamilyObj []model.JobFamily = getAllJobFamilyRequestResult.Data. ([]model.JobFamily)
		
	equalJobFamily := cmp.Equal(createJobFamilyObj.ID, getAllJobFamilyObj[len(getAllJobFamilyObj)-1].ID)
		
	if equalJobFamily == false {
		t.Errorf( "Created object is not equal to the last entry in JobFamily[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for JobFamily
	// --------------------------------------------------------------	
	deleteJobFamilyRequestResult := dao.DeleteJobFamily(uint64(createJobFamilyObj.ID))

	if deleteJobFamilyRequestResult.Success == false {
			t.Errorf(deleteJobFamilyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion JobFamily success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getJobFamilyRequestResult = dao.GetJobFamily( uint64(createJobFamilyObj.ID) )
	
	if getJobFamilyRequestResult.Success == true {
		t.Errorf(getJobFamilyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestJobProfileCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for JobProfile
	//----------------------------------------------------------------------------
	JobProfileObj := model.JobProfile                                                                                            {Title:"test value for Title",JobCode:"test value for JobCode",JobLevel:0,ExemptStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createJobProfileRequestResult := dao.CreateJobProfile( JobProfileObj )
	
	if createJobProfileRequestResult.Success == false {
		t.Errorf(createJobProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Create JobProfile success...")
	}
	
	createJobProfileObj,_ := createJobProfileRequestResult.Data. (model.JobProfile)

	// --------------------------------------------------------------
	// Check JobProfile Obj ID
	// --------------------------------------------------------------	
	if createJobProfileObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for JobProfile" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getJobProfileRequestResult := dao.GetJobProfile( uint64(createJobProfileObj.ID) )
	
	if getJobProfileRequestResult.Success == false {
		t.Errorf(getJobProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Get JobProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getJobProfileObj,_ := getJobProfileRequestResult.Data. (model.JobProfile)
	compareJobProfile := cmp.Equal(createJobProfileObj.ID, getJobProfileObj.ID)
	
	if  compareJobProfile == false	{
		t.Errorf( "Created JobProfile object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllJobProfileRequestResult := dao.GetAllJobProfile()

	if getAllJobProfileRequestResult.Success == false {
			t.Errorf(getAllJobProfileRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll JobProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllJobProfileObj []model.JobProfile = getAllJobProfileRequestResult.Data. ([]model.JobProfile)
		
	equalJobProfile := cmp.Equal(createJobProfileObj.ID, getAllJobProfileObj[len(getAllJobProfileObj)-1].ID)
		
	if equalJobProfile == false {
		t.Errorf( "Created object is not equal to the last entry in JobProfile[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for JobProfile
	// --------------------------------------------------------------	
	deleteJobProfileRequestResult := dao.DeleteJobProfile(uint64(createJobProfileObj.ID))

	if deleteJobProfileRequestResult.Success == false {
			t.Errorf(deleteJobProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion JobProfile success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getJobProfileRequestResult = dao.GetJobProfile( uint64(createJobProfileObj.ID) )
	
	if getJobProfileRequestResult.Success == true {
		t.Errorf(getJobProfileRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCompetencyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Competency
	//----------------------------------------------------------------------------
	CompetencyObj := model.Competency                                                            {Name:"test value for Name",Category:"test value for Category"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCompetencyRequestResult := dao.CreateCompetency( CompetencyObj )
	
	if createCompetencyRequestResult.Success == false {
		t.Errorf(createCompetencyRequestResult.Msg)
	} else {
		fmt.Println("Check Create Competency success...")
	}
	
	createCompetencyObj,_ := createCompetencyRequestResult.Data. (model.Competency)

	// --------------------------------------------------------------
	// Check Competency Obj ID
	// --------------------------------------------------------------	
	if createCompetencyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Competency" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCompetencyRequestResult := dao.GetCompetency( uint64(createCompetencyObj.ID) )
	
	if getCompetencyRequestResult.Success == false {
		t.Errorf(getCompetencyRequestResult.Msg)
	} else {
		fmt.Println("Check Get Competency success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCompetencyObj,_ := getCompetencyRequestResult.Data. (model.Competency)
	compareCompetency := cmp.Equal(createCompetencyObj.ID, getCompetencyObj.ID)
	
	if  compareCompetency == false	{
		t.Errorf( "Created Competency object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCompetencyRequestResult := dao.GetAllCompetency()

	if getAllCompetencyRequestResult.Success == false {
			t.Errorf(getAllCompetencyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Competency success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCompetencyObj []model.Competency = getAllCompetencyRequestResult.Data. ([]model.Competency)
		
	equalCompetency := cmp.Equal(createCompetencyObj.ID, getAllCompetencyObj[len(getAllCompetencyObj)-1].ID)
		
	if equalCompetency == false {
		t.Errorf( "Created object is not equal to the last entry in Competency[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Competency
	// --------------------------------------------------------------	
	deleteCompetencyRequestResult := dao.DeleteCompetency(uint64(createCompetencyObj.ID))

	if deleteCompetencyRequestResult.Success == false {
			t.Errorf(deleteCompetencyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Competency success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCompetencyRequestResult = dao.GetCompetency( uint64(createCompetencyObj.ID) )
	
	if getCompetencyRequestResult.Success == true {
		t.Errorf(getCompetencyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPositionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Position
	//----------------------------------------------------------------------------
	PositionObj := model.Position                                                                                                                    {PositionCode:"test value for PositionCode",Fte:"test value",Status:0,WorkLocationType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPositionRequestResult := dao.CreatePosition( PositionObj )
	
	if createPositionRequestResult.Success == false {
		t.Errorf(createPositionRequestResult.Msg)
	} else {
		fmt.Println("Check Create Position success...")
	}
	
	createPositionObj,_ := createPositionRequestResult.Data. (model.Position)

	// --------------------------------------------------------------
	// Check Position Obj ID
	// --------------------------------------------------------------	
	if createPositionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Position" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPositionRequestResult := dao.GetPosition( uint64(createPositionObj.ID) )
	
	if getPositionRequestResult.Success == false {
		t.Errorf(getPositionRequestResult.Msg)
	} else {
		fmt.Println("Check Get Position success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPositionObj,_ := getPositionRequestResult.Data. (model.Position)
	comparePosition := cmp.Equal(createPositionObj.ID, getPositionObj.ID)
	
	if  comparePosition == false	{
		t.Errorf( "Created Position object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPositionRequestResult := dao.GetAllPosition()

	if getAllPositionRequestResult.Success == false {
			t.Errorf(getAllPositionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Position success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPositionObj []model.Position = getAllPositionRequestResult.Data. ([]model.Position)
		
	equalPosition := cmp.Equal(createPositionObj.ID, getAllPositionObj[len(getAllPositionObj)-1].ID)
		
	if equalPosition == false {
		t.Errorf( "Created object is not equal to the last entry in Position[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Position
	// --------------------------------------------------------------	
	deletePositionRequestResult := dao.DeletePosition(uint64(createPositionObj.ID))

	if deletePositionRequestResult.Success == false {
			t.Errorf(deletePositionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Position success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPositionRequestResult = dao.GetPosition( uint64(createPositionObj.ID) )
	
	if getPositionRequestResult.Success == true {
		t.Errorf(getPositionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEmployeeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Employee
	//----------------------------------------------------------------------------
	EmployeeObj := model.Employee                                                                                                                                                                    {EmployeeNumber:"test value for EmployeeNumber",Name:new PersonName(),WorkEmail:new Email(),WorkPhone:new PhoneNumber(),DateOfHire:time.Now(),NationalId:new NationalID(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEmployeeRequestResult := dao.CreateEmployee( EmployeeObj )
	
	if createEmployeeRequestResult.Success == false {
		t.Errorf(createEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Check Create Employee success...")
	}
	
	createEmployeeObj,_ := createEmployeeRequestResult.Data. (model.Employee)

	// --------------------------------------------------------------
	// Check Employee Obj ID
	// --------------------------------------------------------------	
	if createEmployeeObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Employee" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEmployeeRequestResult := dao.GetEmployee( uint64(createEmployeeObj.ID) )
	
	if getEmployeeRequestResult.Success == false {
		t.Errorf(getEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Check Get Employee success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEmployeeObj,_ := getEmployeeRequestResult.Data. (model.Employee)
	compareEmployee := cmp.Equal(createEmployeeObj.ID, getEmployeeObj.ID)
	
	if  compareEmployee == false	{
		t.Errorf( "Created Employee object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEmployeeRequestResult := dao.GetAllEmployee()

	if getAllEmployeeRequestResult.Success == false {
			t.Errorf(getAllEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Employee success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEmployeeObj []model.Employee = getAllEmployeeRequestResult.Data. ([]model.Employee)
		
	equalEmployee := cmp.Equal(createEmployeeObj.ID, getAllEmployeeObj[len(getAllEmployeeObj)-1].ID)
		
	if equalEmployee == false {
		t.Errorf( "Created object is not equal to the last entry in Employee[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Employee
	// --------------------------------------------------------------	
	deleteEmployeeRequestResult := dao.DeleteEmployee(uint64(createEmployeeObj.ID))

	if deleteEmployeeRequestResult.Success == false {
			t.Errorf(deleteEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Employee success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEmployeeRequestResult = dao.GetEmployee( uint64(createEmployeeObj.ID) )
	
	if getEmployeeRequestResult.Success == true {
		t.Errorf(getEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEmploymentAssignmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for EmploymentAssignment
	//----------------------------------------------------------------------------
	EmploymentAssignmentObj := model.EmploymentAssignment                                                                                                                                                                            {StartDate:time.Now(),EndDate:time.Now(),Primary:true,AssignmentType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEmploymentAssignmentRequestResult := dao.CreateEmploymentAssignment( EmploymentAssignmentObj )
	
	if createEmploymentAssignmentRequestResult.Success == false {
		t.Errorf(createEmploymentAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create EmploymentAssignment success...")
	}
	
	createEmploymentAssignmentObj,_ := createEmploymentAssignmentRequestResult.Data. (model.EmploymentAssignment)

	// --------------------------------------------------------------
	// Check EmploymentAssignment Obj ID
	// --------------------------------------------------------------	
	if createEmploymentAssignmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for EmploymentAssignment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEmploymentAssignmentRequestResult := dao.GetEmploymentAssignment( uint64(createEmploymentAssignmentObj.ID) )
	
	if getEmploymentAssignmentRequestResult.Success == false {
		t.Errorf(getEmploymentAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get EmploymentAssignment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEmploymentAssignmentObj,_ := getEmploymentAssignmentRequestResult.Data. (model.EmploymentAssignment)
	compareEmploymentAssignment := cmp.Equal(createEmploymentAssignmentObj.ID, getEmploymentAssignmentObj.ID)
	
	if  compareEmploymentAssignment == false	{
		t.Errorf( "Created EmploymentAssignment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEmploymentAssignmentRequestResult := dao.GetAllEmploymentAssignment()

	if getAllEmploymentAssignmentRequestResult.Success == false {
			t.Errorf(getAllEmploymentAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll EmploymentAssignment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEmploymentAssignmentObj []model.EmploymentAssignment = getAllEmploymentAssignmentRequestResult.Data. ([]model.EmploymentAssignment)
		
	equalEmploymentAssignment := cmp.Equal(createEmploymentAssignmentObj.ID, getAllEmploymentAssignmentObj[len(getAllEmploymentAssignmentObj)-1].ID)
		
	if equalEmploymentAssignment == false {
		t.Errorf( "Created object is not equal to the last entry in EmploymentAssignment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for EmploymentAssignment
	// --------------------------------------------------------------	
	deleteEmploymentAssignmentRequestResult := dao.DeleteEmploymentAssignment(uint64(createEmploymentAssignmentObj.ID))

	if deleteEmploymentAssignmentRequestResult.Success == false {
			t.Errorf(deleteEmploymentAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion EmploymentAssignment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEmploymentAssignmentRequestResult = dao.GetEmploymentAssignment( uint64(createEmploymentAssignmentObj.ID) )
	
	if getEmploymentAssignmentRequestResult.Success == true {
		t.Errorf(getEmploymentAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEmploymentContractCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for EmploymentContract
	//----------------------------------------------------------------------------
	EmploymentContractObj := model.EmploymentContract                                                                                                                                                                                                                                                    {ContractNumber:"test value for ContractNumber",StartDate:time.Now(),EndDate:time.Now(),WorkHoursPerWeek:"test value",EmploymentType:0,Status:0,PayFrequency:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEmploymentContractRequestResult := dao.CreateEmploymentContract( EmploymentContractObj )
	
	if createEmploymentContractRequestResult.Success == false {
		t.Errorf(createEmploymentContractRequestResult.Msg)
	} else {
		fmt.Println("Check Create EmploymentContract success...")
	}
	
	createEmploymentContractObj,_ := createEmploymentContractRequestResult.Data. (model.EmploymentContract)

	// --------------------------------------------------------------
	// Check EmploymentContract Obj ID
	// --------------------------------------------------------------	
	if createEmploymentContractObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for EmploymentContract" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEmploymentContractRequestResult := dao.GetEmploymentContract( uint64(createEmploymentContractObj.ID) )
	
	if getEmploymentContractRequestResult.Success == false {
		t.Errorf(getEmploymentContractRequestResult.Msg)
	} else {
		fmt.Println("Check Get EmploymentContract success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEmploymentContractObj,_ := getEmploymentContractRequestResult.Data. (model.EmploymentContract)
	compareEmploymentContract := cmp.Equal(createEmploymentContractObj.ID, getEmploymentContractObj.ID)
	
	if  compareEmploymentContract == false	{
		t.Errorf( "Created EmploymentContract object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEmploymentContractRequestResult := dao.GetAllEmploymentContract()

	if getAllEmploymentContractRequestResult.Success == false {
			t.Errorf(getAllEmploymentContractRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll EmploymentContract success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEmploymentContractObj []model.EmploymentContract = getAllEmploymentContractRequestResult.Data. ([]model.EmploymentContract)
		
	equalEmploymentContract := cmp.Equal(createEmploymentContractObj.ID, getAllEmploymentContractObj[len(getAllEmploymentContractObj)-1].ID)
		
	if equalEmploymentContract == false {
		t.Errorf( "Created object is not equal to the last entry in EmploymentContract[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for EmploymentContract
	// --------------------------------------------------------------	
	deleteEmploymentContractRequestResult := dao.DeleteEmploymentContract(uint64(createEmploymentContractObj.ID))

	if deleteEmploymentContractRequestResult.Success == false {
			t.Errorf(deleteEmploymentContractRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion EmploymentContract success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEmploymentContractRequestResult = dao.GetEmploymentContract( uint64(createEmploymentContractObj.ID) )
	
	if getEmploymentContractRequestResult.Success == true {
		t.Errorf(getEmploymentContractRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWorkScheduleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for WorkSchedule
	//----------------------------------------------------------------------------
	WorkScheduleObj := model.WorkSchedule                                                                                                    {Name:"test value for Name",StandardHoursPerWeek:"test value",ScheduleType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWorkScheduleRequestResult := dao.CreateWorkSchedule( WorkScheduleObj )
	
	if createWorkScheduleRequestResult.Success == false {
		t.Errorf(createWorkScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Create WorkSchedule success...")
	}
	
	createWorkScheduleObj,_ := createWorkScheduleRequestResult.Data. (model.WorkSchedule)

	// --------------------------------------------------------------
	// Check WorkSchedule Obj ID
	// --------------------------------------------------------------	
	if createWorkScheduleObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for WorkSchedule" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWorkScheduleRequestResult := dao.GetWorkSchedule( uint64(createWorkScheduleObj.ID) )
	
	if getWorkScheduleRequestResult.Success == false {
		t.Errorf(getWorkScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Get WorkSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWorkScheduleObj,_ := getWorkScheduleRequestResult.Data. (model.WorkSchedule)
	compareWorkSchedule := cmp.Equal(createWorkScheduleObj.ID, getWorkScheduleObj.ID)
	
	if  compareWorkSchedule == false	{
		t.Errorf( "Created WorkSchedule object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWorkScheduleRequestResult := dao.GetAllWorkSchedule()

	if getAllWorkScheduleRequestResult.Success == false {
			t.Errorf(getAllWorkScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll WorkSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWorkScheduleObj []model.WorkSchedule = getAllWorkScheduleRequestResult.Data. ([]model.WorkSchedule)
		
	equalWorkSchedule := cmp.Equal(createWorkScheduleObj.ID, getAllWorkScheduleObj[len(getAllWorkScheduleObj)-1].ID)
		
	if equalWorkSchedule == false {
		t.Errorf( "Created object is not equal to the last entry in WorkSchedule[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for WorkSchedule
	// --------------------------------------------------------------	
	deleteWorkScheduleRequestResult := dao.DeleteWorkSchedule(uint64(createWorkScheduleObj.ID))

	if deleteWorkScheduleRequestResult.Success == false {
			t.Errorf(deleteWorkScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion WorkSchedule success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWorkScheduleRequestResult = dao.GetWorkSchedule( uint64(createWorkScheduleObj.ID) )
	
	if getWorkScheduleRequestResult.Success == true {
		t.Errorf(getWorkScheduleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWorkShiftCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for WorkShift
	//----------------------------------------------------------------------------
	WorkShiftObj := model.WorkShift                                                                            {StartTime:new LocalTime(),EndTime:new LocalTime(),BreakMinutes:100,DayOfWeek:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWorkShiftRequestResult := dao.CreateWorkShift( WorkShiftObj )
	
	if createWorkShiftRequestResult.Success == false {
		t.Errorf(createWorkShiftRequestResult.Msg)
	} else {
		fmt.Println("Check Create WorkShift success...")
	}
	
	createWorkShiftObj,_ := createWorkShiftRequestResult.Data. (model.WorkShift)

	// --------------------------------------------------------------
	// Check WorkShift Obj ID
	// --------------------------------------------------------------	
	if createWorkShiftObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for WorkShift" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWorkShiftRequestResult := dao.GetWorkShift( uint64(createWorkShiftObj.ID) )
	
	if getWorkShiftRequestResult.Success == false {
		t.Errorf(getWorkShiftRequestResult.Msg)
	} else {
		fmt.Println("Check Get WorkShift success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWorkShiftObj,_ := getWorkShiftRequestResult.Data. (model.WorkShift)
	compareWorkShift := cmp.Equal(createWorkShiftObj.ID, getWorkShiftObj.ID)
	
	if  compareWorkShift == false	{
		t.Errorf( "Created WorkShift object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWorkShiftRequestResult := dao.GetAllWorkShift()

	if getAllWorkShiftRequestResult.Success == false {
			t.Errorf(getAllWorkShiftRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll WorkShift success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWorkShiftObj []model.WorkShift = getAllWorkShiftRequestResult.Data. ([]model.WorkShift)
		
	equalWorkShift := cmp.Equal(createWorkShiftObj.ID, getAllWorkShiftObj[len(getAllWorkShiftObj)-1].ID)
		
	if equalWorkShift == false {
		t.Errorf( "Created object is not equal to the last entry in WorkShift[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for WorkShift
	// --------------------------------------------------------------	
	deleteWorkShiftRequestResult := dao.DeleteWorkShift(uint64(createWorkShiftObj.ID))

	if deleteWorkShiftRequestResult.Success == false {
			t.Errorf(deleteWorkShiftRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion WorkShift success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWorkShiftRequestResult = dao.GetWorkShift( uint64(createWorkShiftObj.ID) )
	
	if getWorkShiftRequestResult.Success == true {
		t.Errorf(getWorkShiftRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestScheduleExceptionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ScheduleException
	//----------------------------------------------------------------------------
	ScheduleExceptionObj := model.ScheduleException                                                                                                                                            {Date:time.Now(),Reason:"test value for Reason",Hours:"test value"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createScheduleExceptionRequestResult := dao.CreateScheduleException( ScheduleExceptionObj )
	
	if createScheduleExceptionRequestResult.Success == false {
		t.Errorf(createScheduleExceptionRequestResult.Msg)
	} else {
		fmt.Println("Check Create ScheduleException success...")
	}
	
	createScheduleExceptionObj,_ := createScheduleExceptionRequestResult.Data. (model.ScheduleException)

	// --------------------------------------------------------------
	// Check ScheduleException Obj ID
	// --------------------------------------------------------------	
	if createScheduleExceptionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ScheduleException" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getScheduleExceptionRequestResult := dao.GetScheduleException( uint64(createScheduleExceptionObj.ID) )
	
	if getScheduleExceptionRequestResult.Success == false {
		t.Errorf(getScheduleExceptionRequestResult.Msg)
	} else {
		fmt.Println("Check Get ScheduleException success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getScheduleExceptionObj,_ := getScheduleExceptionRequestResult.Data. (model.ScheduleException)
	compareScheduleException := cmp.Equal(createScheduleExceptionObj.ID, getScheduleExceptionObj.ID)
	
	if  compareScheduleException == false	{
		t.Errorf( "Created ScheduleException object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllScheduleExceptionRequestResult := dao.GetAllScheduleException()

	if getAllScheduleExceptionRequestResult.Success == false {
			t.Errorf(getAllScheduleExceptionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ScheduleException success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllScheduleExceptionObj []model.ScheduleException = getAllScheduleExceptionRequestResult.Data. ([]model.ScheduleException)
		
	equalScheduleException := cmp.Equal(createScheduleExceptionObj.ID, getAllScheduleExceptionObj[len(getAllScheduleExceptionObj)-1].ID)
		
	if equalScheduleException == false {
		t.Errorf( "Created object is not equal to the last entry in ScheduleException[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ScheduleException
	// --------------------------------------------------------------	
	deleteScheduleExceptionRequestResult := dao.DeleteScheduleException(uint64(createScheduleExceptionObj.ID))

	if deleteScheduleExceptionRequestResult.Success == false {
			t.Errorf(deleteScheduleExceptionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ScheduleException success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getScheduleExceptionRequestResult = dao.GetScheduleException( uint64(createScheduleExceptionObj.ID) )
	
	if getScheduleExceptionRequestResult.Success == true {
		t.Errorf(getScheduleExceptionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCompensationPackageCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CompensationPackage
	//----------------------------------------------------------------------------
	CompensationPackageObj := model.CompensationPackage                                                                                                                                            {EffectiveFrom:time.Now(),EffectiveTo:time.Now(),Currency:"test value for Currency"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCompensationPackageRequestResult := dao.CreateCompensationPackage( CompensationPackageObj )
	
	if createCompensationPackageRequestResult.Success == false {
		t.Errorf(createCompensationPackageRequestResult.Msg)
	} else {
		fmt.Println("Check Create CompensationPackage success...")
	}
	
	createCompensationPackageObj,_ := createCompensationPackageRequestResult.Data. (model.CompensationPackage)

	// --------------------------------------------------------------
	// Check CompensationPackage Obj ID
	// --------------------------------------------------------------	
	if createCompensationPackageObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CompensationPackage" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCompensationPackageRequestResult := dao.GetCompensationPackage( uint64(createCompensationPackageObj.ID) )
	
	if getCompensationPackageRequestResult.Success == false {
		t.Errorf(getCompensationPackageRequestResult.Msg)
	} else {
		fmt.Println("Check Get CompensationPackage success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCompensationPackageObj,_ := getCompensationPackageRequestResult.Data. (model.CompensationPackage)
	compareCompensationPackage := cmp.Equal(createCompensationPackageObj.ID, getCompensationPackageObj.ID)
	
	if  compareCompensationPackage == false	{
		t.Errorf( "Created CompensationPackage object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCompensationPackageRequestResult := dao.GetAllCompensationPackage()

	if getAllCompensationPackageRequestResult.Success == false {
			t.Errorf(getAllCompensationPackageRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CompensationPackage success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCompensationPackageObj []model.CompensationPackage = getAllCompensationPackageRequestResult.Data. ([]model.CompensationPackage)
		
	equalCompensationPackage := cmp.Equal(createCompensationPackageObj.ID, getAllCompensationPackageObj[len(getAllCompensationPackageObj)-1].ID)
		
	if equalCompensationPackage == false {
		t.Errorf( "Created object is not equal to the last entry in CompensationPackage[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CompensationPackage
	// --------------------------------------------------------------	
	deleteCompensationPackageRequestResult := dao.DeleteCompensationPackage(uint64(createCompensationPackageObj.ID))

	if deleteCompensationPackageRequestResult.Success == false {
			t.Errorf(deleteCompensationPackageRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CompensationPackage success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCompensationPackageRequestResult = dao.GetCompensationPackage( uint64(createCompensationPackageObj.ID) )
	
	if getCompensationPackageRequestResult.Success == true {
		t.Errorf(getCompensationPackageRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSalaryComponentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SalaryComponent
	//----------------------------------------------------------------------------
	SalaryComponentObj := model.SalaryComponent                                                            {Amount:new Money(),Recurring:true,ComponentType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSalaryComponentRequestResult := dao.CreateSalaryComponent( SalaryComponentObj )
	
	if createSalaryComponentRequestResult.Success == false {
		t.Errorf(createSalaryComponentRequestResult.Msg)
	} else {
		fmt.Println("Check Create SalaryComponent success...")
	}
	
	createSalaryComponentObj,_ := createSalaryComponentRequestResult.Data. (model.SalaryComponent)

	// --------------------------------------------------------------
	// Check SalaryComponent Obj ID
	// --------------------------------------------------------------	
	if createSalaryComponentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SalaryComponent" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSalaryComponentRequestResult := dao.GetSalaryComponent( uint64(createSalaryComponentObj.ID) )
	
	if getSalaryComponentRequestResult.Success == false {
		t.Errorf(getSalaryComponentRequestResult.Msg)
	} else {
		fmt.Println("Check Get SalaryComponent success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSalaryComponentObj,_ := getSalaryComponentRequestResult.Data. (model.SalaryComponent)
	compareSalaryComponent := cmp.Equal(createSalaryComponentObj.ID, getSalaryComponentObj.ID)
	
	if  compareSalaryComponent == false	{
		t.Errorf( "Created SalaryComponent object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSalaryComponentRequestResult := dao.GetAllSalaryComponent()

	if getAllSalaryComponentRequestResult.Success == false {
			t.Errorf(getAllSalaryComponentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SalaryComponent success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSalaryComponentObj []model.SalaryComponent = getAllSalaryComponentRequestResult.Data. ([]model.SalaryComponent)
		
	equalSalaryComponent := cmp.Equal(createSalaryComponentObj.ID, getAllSalaryComponentObj[len(getAllSalaryComponentObj)-1].ID)
		
	if equalSalaryComponent == false {
		t.Errorf( "Created object is not equal to the last entry in SalaryComponent[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SalaryComponent
	// --------------------------------------------------------------	
	deleteSalaryComponentRequestResult := dao.DeleteSalaryComponent(uint64(createSalaryComponentObj.ID))

	if deleteSalaryComponentRequestResult.Success == false {
			t.Errorf(deleteSalaryComponentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SalaryComponent success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSalaryComponentRequestResult = dao.GetSalaryComponent( uint64(createSalaryComponentObj.ID) )
	
	if getSalaryComponentRequestResult.Success == true {
		t.Errorf(getSalaryComponentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBonusPlanCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BonusPlan
	//----------------------------------------------------------------------------
	BonusPlanObj := model.BonusPlan                                            {Name:"test value for Name",TargetPercentage:new Percentage()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBonusPlanRequestResult := dao.CreateBonusPlan( BonusPlanObj )
	
	if createBonusPlanRequestResult.Success == false {
		t.Errorf(createBonusPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Create BonusPlan success...")
	}
	
	createBonusPlanObj,_ := createBonusPlanRequestResult.Data. (model.BonusPlan)

	// --------------------------------------------------------------
	// Check BonusPlan Obj ID
	// --------------------------------------------------------------	
	if createBonusPlanObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BonusPlan" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBonusPlanRequestResult := dao.GetBonusPlan( uint64(createBonusPlanObj.ID) )
	
	if getBonusPlanRequestResult.Success == false {
		t.Errorf(getBonusPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Get BonusPlan success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBonusPlanObj,_ := getBonusPlanRequestResult.Data. (model.BonusPlan)
	compareBonusPlan := cmp.Equal(createBonusPlanObj.ID, getBonusPlanObj.ID)
	
	if  compareBonusPlan == false	{
		t.Errorf( "Created BonusPlan object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBonusPlanRequestResult := dao.GetAllBonusPlan()

	if getAllBonusPlanRequestResult.Success == false {
			t.Errorf(getAllBonusPlanRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BonusPlan success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBonusPlanObj []model.BonusPlan = getAllBonusPlanRequestResult.Data. ([]model.BonusPlan)
		
	equalBonusPlan := cmp.Equal(createBonusPlanObj.ID, getAllBonusPlanObj[len(getAllBonusPlanObj)-1].ID)
		
	if equalBonusPlan == false {
		t.Errorf( "Created object is not equal to the last entry in BonusPlan[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BonusPlan
	// --------------------------------------------------------------	
	deleteBonusPlanRequestResult := dao.DeleteBonusPlan(uint64(createBonusPlanObj.ID))

	if deleteBonusPlanRequestResult.Success == false {
			t.Errorf(deleteBonusPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BonusPlan success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBonusPlanRequestResult = dao.GetBonusPlan( uint64(createBonusPlanObj.ID) )
	
	if getBonusPlanRequestResult.Success == true {
		t.Errorf(getBonusPlanRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEquityGrantCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for EquityGrant
	//----------------------------------------------------------------------------
	EquityGrantObj := model.EquityGrant                                                                                                                                    {GrantId:"test value for GrantId",GrantedUnits:100,VestingStart:time.Now(),GrantType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEquityGrantRequestResult := dao.CreateEquityGrant( EquityGrantObj )
	
	if createEquityGrantRequestResult.Success == false {
		t.Errorf(createEquityGrantRequestResult.Msg)
	} else {
		fmt.Println("Check Create EquityGrant success...")
	}
	
	createEquityGrantObj,_ := createEquityGrantRequestResult.Data. (model.EquityGrant)

	// --------------------------------------------------------------
	// Check EquityGrant Obj ID
	// --------------------------------------------------------------	
	if createEquityGrantObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for EquityGrant" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEquityGrantRequestResult := dao.GetEquityGrant( uint64(createEquityGrantObj.ID) )
	
	if getEquityGrantRequestResult.Success == false {
		t.Errorf(getEquityGrantRequestResult.Msg)
	} else {
		fmt.Println("Check Get EquityGrant success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEquityGrantObj,_ := getEquityGrantRequestResult.Data. (model.EquityGrant)
	compareEquityGrant := cmp.Equal(createEquityGrantObj.ID, getEquityGrantObj.ID)
	
	if  compareEquityGrant == false	{
		t.Errorf( "Created EquityGrant object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEquityGrantRequestResult := dao.GetAllEquityGrant()

	if getAllEquityGrantRequestResult.Success == false {
			t.Errorf(getAllEquityGrantRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll EquityGrant success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEquityGrantObj []model.EquityGrant = getAllEquityGrantRequestResult.Data. ([]model.EquityGrant)
		
	equalEquityGrant := cmp.Equal(createEquityGrantObj.ID, getAllEquityGrantObj[len(getAllEquityGrantObj)-1].ID)
		
	if equalEquityGrant == false {
		t.Errorf( "Created object is not equal to the last entry in EquityGrant[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for EquityGrant
	// --------------------------------------------------------------	
	deleteEquityGrantRequestResult := dao.DeleteEquityGrant(uint64(createEquityGrantObj.ID))

	if deleteEquityGrantRequestResult.Success == false {
			t.Errorf(deleteEquityGrantRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion EquityGrant success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEquityGrantRequestResult = dao.GetEquityGrant( uint64(createEquityGrantObj.ID) )
	
	if getEquityGrantRequestResult.Success == true {
		t.Errorf(getEquityGrantRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBenefitPlanCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BenefitPlan
	//----------------------------------------------------------------------------
	BenefitPlanObj := model.BenefitPlan                                                                                                                                            {Name:"test value for Name",ProviderName:"test value for ProviderName",EmployeeContributionRate:new Percentage(),EmployerContributionRate:new Percentage(),EligibilityRules:"test value for EligibilityRules",BenefitType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBenefitPlanRequestResult := dao.CreateBenefitPlan( BenefitPlanObj )
	
	if createBenefitPlanRequestResult.Success == false {
		t.Errorf(createBenefitPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Create BenefitPlan success...")
	}
	
	createBenefitPlanObj,_ := createBenefitPlanRequestResult.Data. (model.BenefitPlan)

	// --------------------------------------------------------------
	// Check BenefitPlan Obj ID
	// --------------------------------------------------------------	
	if createBenefitPlanObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BenefitPlan" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBenefitPlanRequestResult := dao.GetBenefitPlan( uint64(createBenefitPlanObj.ID) )
	
	if getBenefitPlanRequestResult.Success == false {
		t.Errorf(getBenefitPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Get BenefitPlan success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBenefitPlanObj,_ := getBenefitPlanRequestResult.Data. (model.BenefitPlan)
	compareBenefitPlan := cmp.Equal(createBenefitPlanObj.ID, getBenefitPlanObj.ID)
	
	if  compareBenefitPlan == false	{
		t.Errorf( "Created BenefitPlan object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBenefitPlanRequestResult := dao.GetAllBenefitPlan()

	if getAllBenefitPlanRequestResult.Success == false {
			t.Errorf(getAllBenefitPlanRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BenefitPlan success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBenefitPlanObj []model.BenefitPlan = getAllBenefitPlanRequestResult.Data. ([]model.BenefitPlan)
		
	equalBenefitPlan := cmp.Equal(createBenefitPlanObj.ID, getAllBenefitPlanObj[len(getAllBenefitPlanObj)-1].ID)
		
	if equalBenefitPlan == false {
		t.Errorf( "Created object is not equal to the last entry in BenefitPlan[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BenefitPlan
	// --------------------------------------------------------------	
	deleteBenefitPlanRequestResult := dao.DeleteBenefitPlan(uint64(createBenefitPlanObj.ID))

	if deleteBenefitPlanRequestResult.Success == false {
			t.Errorf(deleteBenefitPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BenefitPlan success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBenefitPlanRequestResult = dao.GetBenefitPlan( uint64(createBenefitPlanObj.ID) )
	
	if getBenefitPlanRequestResult.Success == true {
		t.Errorf(getBenefitPlanRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBenefitEnrollmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BenefitEnrollment
	//----------------------------------------------------------------------------
	BenefitEnrollmentObj := model.BenefitEnrollment                                                                                                                                                                            {EnrollmentId:"test value for EnrollmentId",EffectiveFrom:time.Now(),EffectiveTo:time.Now(),Status:0,CoverageLevel:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBenefitEnrollmentRequestResult := dao.CreateBenefitEnrollment( BenefitEnrollmentObj )
	
	if createBenefitEnrollmentRequestResult.Success == false {
		t.Errorf(createBenefitEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create BenefitEnrollment success...")
	}
	
	createBenefitEnrollmentObj,_ := createBenefitEnrollmentRequestResult.Data. (model.BenefitEnrollment)

	// --------------------------------------------------------------
	// Check BenefitEnrollment Obj ID
	// --------------------------------------------------------------	
	if createBenefitEnrollmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BenefitEnrollment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBenefitEnrollmentRequestResult := dao.GetBenefitEnrollment( uint64(createBenefitEnrollmentObj.ID) )
	
	if getBenefitEnrollmentRequestResult.Success == false {
		t.Errorf(getBenefitEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get BenefitEnrollment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBenefitEnrollmentObj,_ := getBenefitEnrollmentRequestResult.Data. (model.BenefitEnrollment)
	compareBenefitEnrollment := cmp.Equal(createBenefitEnrollmentObj.ID, getBenefitEnrollmentObj.ID)
	
	if  compareBenefitEnrollment == false	{
		t.Errorf( "Created BenefitEnrollment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBenefitEnrollmentRequestResult := dao.GetAllBenefitEnrollment()

	if getAllBenefitEnrollmentRequestResult.Success == false {
			t.Errorf(getAllBenefitEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BenefitEnrollment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBenefitEnrollmentObj []model.BenefitEnrollment = getAllBenefitEnrollmentRequestResult.Data. ([]model.BenefitEnrollment)
		
	equalBenefitEnrollment := cmp.Equal(createBenefitEnrollmentObj.ID, getAllBenefitEnrollmentObj[len(getAllBenefitEnrollmentObj)-1].ID)
		
	if equalBenefitEnrollment == false {
		t.Errorf( "Created object is not equal to the last entry in BenefitEnrollment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BenefitEnrollment
	// --------------------------------------------------------------	
	deleteBenefitEnrollmentRequestResult := dao.DeleteBenefitEnrollment(uint64(createBenefitEnrollmentObj.ID))

	if deleteBenefitEnrollmentRequestResult.Success == false {
			t.Errorf(deleteBenefitEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BenefitEnrollment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBenefitEnrollmentRequestResult = dao.GetBenefitEnrollment( uint64(createBenefitEnrollmentObj.ID) )
	
	if getBenefitEnrollmentRequestResult.Success == true {
		t.Errorf(getBenefitEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDependentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Dependent
	//----------------------------------------------------------------------------
	DependentObj := model.Dependent                                                                                                                                    {FirstName:"test value for FirstName",LastName:"test value for LastName",BirthDate:time.Now(),Relationship:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDependentRequestResult := dao.CreateDependent( DependentObj )
	
	if createDependentRequestResult.Success == false {
		t.Errorf(createDependentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Dependent success...")
	}
	
	createDependentObj,_ := createDependentRequestResult.Data. (model.Dependent)

	// --------------------------------------------------------------
	// Check Dependent Obj ID
	// --------------------------------------------------------------	
	if createDependentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Dependent" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDependentRequestResult := dao.GetDependent( uint64(createDependentObj.ID) )
	
	if getDependentRequestResult.Success == false {
		t.Errorf(getDependentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Dependent success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDependentObj,_ := getDependentRequestResult.Data. (model.Dependent)
	compareDependent := cmp.Equal(createDependentObj.ID, getDependentObj.ID)
	
	if  compareDependent == false	{
		t.Errorf( "Created Dependent object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDependentRequestResult := dao.GetAllDependent()

	if getAllDependentRequestResult.Success == false {
			t.Errorf(getAllDependentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Dependent success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDependentObj []model.Dependent = getAllDependentRequestResult.Data. ([]model.Dependent)
		
	equalDependent := cmp.Equal(createDependentObj.ID, getAllDependentObj[len(getAllDependentObj)-1].ID)
		
	if equalDependent == false {
		t.Errorf( "Created object is not equal to the last entry in Dependent[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Dependent
	// --------------------------------------------------------------	
	deleteDependentRequestResult := dao.DeleteDependent(uint64(createDependentObj.ID))

	if deleteDependentRequestResult.Success == false {
			t.Errorf(deleteDependentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Dependent success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDependentRequestResult = dao.GetDependent( uint64(createDependentObj.ID) )
	
	if getDependentRequestResult.Success == true {
		t.Errorf(getDependentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPayrollCalendarCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PayrollCalendar
	//----------------------------------------------------------------------------
	PayrollCalendarObj := model.PayrollCalendar                                                                            {Name:"test value for Name",Country:"test value for Country",PayFrequency:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPayrollCalendarRequestResult := dao.CreatePayrollCalendar( PayrollCalendarObj )
	
	if createPayrollCalendarRequestResult.Success == false {
		t.Errorf(createPayrollCalendarRequestResult.Msg)
	} else {
		fmt.Println("Check Create PayrollCalendar success...")
	}
	
	createPayrollCalendarObj,_ := createPayrollCalendarRequestResult.Data. (model.PayrollCalendar)

	// --------------------------------------------------------------
	// Check PayrollCalendar Obj ID
	// --------------------------------------------------------------	
	if createPayrollCalendarObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PayrollCalendar" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPayrollCalendarRequestResult := dao.GetPayrollCalendar( uint64(createPayrollCalendarObj.ID) )
	
	if getPayrollCalendarRequestResult.Success == false {
		t.Errorf(getPayrollCalendarRequestResult.Msg)
	} else {
		fmt.Println("Check Get PayrollCalendar success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPayrollCalendarObj,_ := getPayrollCalendarRequestResult.Data. (model.PayrollCalendar)
	comparePayrollCalendar := cmp.Equal(createPayrollCalendarObj.ID, getPayrollCalendarObj.ID)
	
	if  comparePayrollCalendar == false	{
		t.Errorf( "Created PayrollCalendar object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPayrollCalendarRequestResult := dao.GetAllPayrollCalendar()

	if getAllPayrollCalendarRequestResult.Success == false {
			t.Errorf(getAllPayrollCalendarRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PayrollCalendar success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPayrollCalendarObj []model.PayrollCalendar = getAllPayrollCalendarRequestResult.Data. ([]model.PayrollCalendar)
		
	equalPayrollCalendar := cmp.Equal(createPayrollCalendarObj.ID, getAllPayrollCalendarObj[len(getAllPayrollCalendarObj)-1].ID)
		
	if equalPayrollCalendar == false {
		t.Errorf( "Created object is not equal to the last entry in PayrollCalendar[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PayrollCalendar
	// --------------------------------------------------------------	
	deletePayrollCalendarRequestResult := dao.DeletePayrollCalendar(uint64(createPayrollCalendarObj.ID))

	if deletePayrollCalendarRequestResult.Success == false {
			t.Errorf(deletePayrollCalendarRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PayrollCalendar success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPayrollCalendarRequestResult = dao.GetPayrollCalendar( uint64(createPayrollCalendarObj.ID) )
	
	if getPayrollCalendarRequestResult.Success == true {
		t.Errorf(getPayrollCalendarRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPayrollRunCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PayrollRun
	//----------------------------------------------------------------------------
	PayrollRunObj := model.PayrollRun                                                                                                                                                                                                                    {RunNumber:"test value for RunNumber",PeriodStart:time.Now(),PeriodEnd:time.Now(),PaymentDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPayrollRunRequestResult := dao.CreatePayrollRun( PayrollRunObj )
	
	if createPayrollRunRequestResult.Success == false {
		t.Errorf(createPayrollRunRequestResult.Msg)
	} else {
		fmt.Println("Check Create PayrollRun success...")
	}
	
	createPayrollRunObj,_ := createPayrollRunRequestResult.Data. (model.PayrollRun)

	// --------------------------------------------------------------
	// Check PayrollRun Obj ID
	// --------------------------------------------------------------	
	if createPayrollRunObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PayrollRun" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPayrollRunRequestResult := dao.GetPayrollRun( uint64(createPayrollRunObj.ID) )
	
	if getPayrollRunRequestResult.Success == false {
		t.Errorf(getPayrollRunRequestResult.Msg)
	} else {
		fmt.Println("Check Get PayrollRun success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPayrollRunObj,_ := getPayrollRunRequestResult.Data. (model.PayrollRun)
	comparePayrollRun := cmp.Equal(createPayrollRunObj.ID, getPayrollRunObj.ID)
	
	if  comparePayrollRun == false	{
		t.Errorf( "Created PayrollRun object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPayrollRunRequestResult := dao.GetAllPayrollRun()

	if getAllPayrollRunRequestResult.Success == false {
			t.Errorf(getAllPayrollRunRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PayrollRun success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPayrollRunObj []model.PayrollRun = getAllPayrollRunRequestResult.Data. ([]model.PayrollRun)
		
	equalPayrollRun := cmp.Equal(createPayrollRunObj.ID, getAllPayrollRunObj[len(getAllPayrollRunObj)-1].ID)
		
	if equalPayrollRun == false {
		t.Errorf( "Created object is not equal to the last entry in PayrollRun[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PayrollRun
	// --------------------------------------------------------------	
	deletePayrollRunRequestResult := dao.DeletePayrollRun(uint64(createPayrollRunObj.ID))

	if deletePayrollRunRequestResult.Success == false {
			t.Errorf(deletePayrollRunRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PayrollRun success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPayrollRunRequestResult = dao.GetPayrollRun( uint64(createPayrollRunObj.ID) )
	
	if getPayrollRunRequestResult.Success == true {
		t.Errorf(getPayrollRunRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPayrollItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PayrollItem
	//----------------------------------------------------------------------------
	PayrollItemObj := model.PayrollItem                                                            {Amount:new Money(),Taxable:true,ItemType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPayrollItemRequestResult := dao.CreatePayrollItem( PayrollItemObj )
	
	if createPayrollItemRequestResult.Success == false {
		t.Errorf(createPayrollItemRequestResult.Msg)
	} else {
		fmt.Println("Check Create PayrollItem success...")
	}
	
	createPayrollItemObj,_ := createPayrollItemRequestResult.Data. (model.PayrollItem)

	// --------------------------------------------------------------
	// Check PayrollItem Obj ID
	// --------------------------------------------------------------	
	if createPayrollItemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PayrollItem" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPayrollItemRequestResult := dao.GetPayrollItem( uint64(createPayrollItemObj.ID) )
	
	if getPayrollItemRequestResult.Success == false {
		t.Errorf(getPayrollItemRequestResult.Msg)
	} else {
		fmt.Println("Check Get PayrollItem success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPayrollItemObj,_ := getPayrollItemRequestResult.Data. (model.PayrollItem)
	comparePayrollItem := cmp.Equal(createPayrollItemObj.ID, getPayrollItemObj.ID)
	
	if  comparePayrollItem == false	{
		t.Errorf( "Created PayrollItem object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPayrollItemRequestResult := dao.GetAllPayrollItem()

	if getAllPayrollItemRequestResult.Success == false {
			t.Errorf(getAllPayrollItemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PayrollItem success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPayrollItemObj []model.PayrollItem = getAllPayrollItemRequestResult.Data. ([]model.PayrollItem)
		
	equalPayrollItem := cmp.Equal(createPayrollItemObj.ID, getAllPayrollItemObj[len(getAllPayrollItemObj)-1].ID)
		
	if equalPayrollItem == false {
		t.Errorf( "Created object is not equal to the last entry in PayrollItem[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PayrollItem
	// --------------------------------------------------------------	
	deletePayrollItemRequestResult := dao.DeletePayrollItem(uint64(createPayrollItemObj.ID))

	if deletePayrollItemRequestResult.Success == false {
			t.Errorf(deletePayrollItemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PayrollItem success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPayrollItemRequestResult = dao.GetPayrollItem( uint64(createPayrollItemObj.ID) )
	
	if getPayrollItemRequestResult.Success == true {
		t.Errorf(getPayrollItemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTaxWithholdingCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TaxWithholding
	//----------------------------------------------------------------------------
	TaxWithholdingObj := model.TaxWithholding                                                                            {TaxId:new TaxId(),Allowances:100,AdditionalAmount:new Money(),FilingStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTaxWithholdingRequestResult := dao.CreateTaxWithholding( TaxWithholdingObj )
	
	if createTaxWithholdingRequestResult.Success == false {
		t.Errorf(createTaxWithholdingRequestResult.Msg)
	} else {
		fmt.Println("Check Create TaxWithholding success...")
	}
	
	createTaxWithholdingObj,_ := createTaxWithholdingRequestResult.Data. (model.TaxWithholding)

	// --------------------------------------------------------------
	// Check TaxWithholding Obj ID
	// --------------------------------------------------------------	
	if createTaxWithholdingObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TaxWithholding" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTaxWithholdingRequestResult := dao.GetTaxWithholding( uint64(createTaxWithholdingObj.ID) )
	
	if getTaxWithholdingRequestResult.Success == false {
		t.Errorf(getTaxWithholdingRequestResult.Msg)
	} else {
		fmt.Println("Check Get TaxWithholding success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTaxWithholdingObj,_ := getTaxWithholdingRequestResult.Data. (model.TaxWithholding)
	compareTaxWithholding := cmp.Equal(createTaxWithholdingObj.ID, getTaxWithholdingObj.ID)
	
	if  compareTaxWithholding == false	{
		t.Errorf( "Created TaxWithholding object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTaxWithholdingRequestResult := dao.GetAllTaxWithholding()

	if getAllTaxWithholdingRequestResult.Success == false {
			t.Errorf(getAllTaxWithholdingRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TaxWithholding success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTaxWithholdingObj []model.TaxWithholding = getAllTaxWithholdingRequestResult.Data. ([]model.TaxWithholding)
		
	equalTaxWithholding := cmp.Equal(createTaxWithholdingObj.ID, getAllTaxWithholdingObj[len(getAllTaxWithholdingObj)-1].ID)
		
	if equalTaxWithholding == false {
		t.Errorf( "Created object is not equal to the last entry in TaxWithholding[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TaxWithholding
	// --------------------------------------------------------------	
	deleteTaxWithholdingRequestResult := dao.DeleteTaxWithholding(uint64(createTaxWithholdingObj.ID))

	if deleteTaxWithholdingRequestResult.Success == false {
			t.Errorf(deleteTaxWithholdingRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TaxWithholding success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTaxWithholdingRequestResult = dao.GetTaxWithholding( uint64(createTaxWithholdingObj.ID) )
	
	if getTaxWithholdingRequestResult.Success == true {
		t.Errorf(getTaxWithholdingRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPaymentMethodCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PaymentMethod
	//----------------------------------------------------------------------------
	PaymentMethodObj := model.PaymentMethod                                            {Preferred:true,MethodType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPaymentMethodRequestResult := dao.CreatePaymentMethod( PaymentMethodObj )
	
	if createPaymentMethodRequestResult.Success == false {
		t.Errorf(createPaymentMethodRequestResult.Msg)
	} else {
		fmt.Println("Check Create PaymentMethod success...")
	}
	
	createPaymentMethodObj,_ := createPaymentMethodRequestResult.Data. (model.PaymentMethod)

	// --------------------------------------------------------------
	// Check PaymentMethod Obj ID
	// --------------------------------------------------------------	
	if createPaymentMethodObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PaymentMethod" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPaymentMethodRequestResult := dao.GetPaymentMethod( uint64(createPaymentMethodObj.ID) )
	
	if getPaymentMethodRequestResult.Success == false {
		t.Errorf(getPaymentMethodRequestResult.Msg)
	} else {
		fmt.Println("Check Get PaymentMethod success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPaymentMethodObj,_ := getPaymentMethodRequestResult.Data. (model.PaymentMethod)
	comparePaymentMethod := cmp.Equal(createPaymentMethodObj.ID, getPaymentMethodObj.ID)
	
	if  comparePaymentMethod == false	{
		t.Errorf( "Created PaymentMethod object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPaymentMethodRequestResult := dao.GetAllPaymentMethod()

	if getAllPaymentMethodRequestResult.Success == false {
			t.Errorf(getAllPaymentMethodRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PaymentMethod success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPaymentMethodObj []model.PaymentMethod = getAllPaymentMethodRequestResult.Data. ([]model.PaymentMethod)
		
	equalPaymentMethod := cmp.Equal(createPaymentMethodObj.ID, getAllPaymentMethodObj[len(getAllPaymentMethodObj)-1].ID)
		
	if equalPaymentMethod == false {
		t.Errorf( "Created object is not equal to the last entry in PaymentMethod[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PaymentMethod
	// --------------------------------------------------------------	
	deletePaymentMethodRequestResult := dao.DeletePaymentMethod(uint64(createPaymentMethodObj.ID))

	if deletePaymentMethodRequestResult.Success == false {
			t.Errorf(deletePaymentMethodRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PaymentMethod success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPaymentMethodRequestResult = dao.GetPaymentMethod( uint64(createPaymentMethodObj.ID) )
	
	if getPaymentMethodRequestResult.Success == true {
		t.Errorf(getPaymentMethodRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTimesheetCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Timesheet
	//----------------------------------------------------------------------------
	TimesheetObj := model.Timesheet                                                                                                                                                                                    {PeriodStart:time.Now(),PeriodEnd:time.Now(),SubmissionDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTimesheetRequestResult := dao.CreateTimesheet( TimesheetObj )
	
	if createTimesheetRequestResult.Success == false {
		t.Errorf(createTimesheetRequestResult.Msg)
	} else {
		fmt.Println("Check Create Timesheet success...")
	}
	
	createTimesheetObj,_ := createTimesheetRequestResult.Data. (model.Timesheet)

	// --------------------------------------------------------------
	// Check Timesheet Obj ID
	// --------------------------------------------------------------	
	if createTimesheetObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Timesheet" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTimesheetRequestResult := dao.GetTimesheet( uint64(createTimesheetObj.ID) )
	
	if getTimesheetRequestResult.Success == false {
		t.Errorf(getTimesheetRequestResult.Msg)
	} else {
		fmt.Println("Check Get Timesheet success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTimesheetObj,_ := getTimesheetRequestResult.Data. (model.Timesheet)
	compareTimesheet := cmp.Equal(createTimesheetObj.ID, getTimesheetObj.ID)
	
	if  compareTimesheet == false	{
		t.Errorf( "Created Timesheet object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTimesheetRequestResult := dao.GetAllTimesheet()

	if getAllTimesheetRequestResult.Success == false {
			t.Errorf(getAllTimesheetRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Timesheet success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTimesheetObj []model.Timesheet = getAllTimesheetRequestResult.Data. ([]model.Timesheet)
		
	equalTimesheet := cmp.Equal(createTimesheetObj.ID, getAllTimesheetObj[len(getAllTimesheetObj)-1].ID)
		
	if equalTimesheet == false {
		t.Errorf( "Created object is not equal to the last entry in Timesheet[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Timesheet
	// --------------------------------------------------------------	
	deleteTimesheetRequestResult := dao.DeleteTimesheet(uint64(createTimesheetObj.ID))

	if deleteTimesheetRequestResult.Success == false {
			t.Errorf(deleteTimesheetRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Timesheet success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTimesheetRequestResult = dao.GetTimesheet( uint64(createTimesheetObj.ID) )
	
	if getTimesheetRequestResult.Success == true {
		t.Errorf(getTimesheetRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTimeEntryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TimeEntry
	//----------------------------------------------------------------------------
	TimeEntryObj := model.TimeEntry                                                                                                                            {EntryDate:time.Now(),HoursWorked:"test value",EntryType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTimeEntryRequestResult := dao.CreateTimeEntry( TimeEntryObj )
	
	if createTimeEntryRequestResult.Success == false {
		t.Errorf(createTimeEntryRequestResult.Msg)
	} else {
		fmt.Println("Check Create TimeEntry success...")
	}
	
	createTimeEntryObj,_ := createTimeEntryRequestResult.Data. (model.TimeEntry)

	// --------------------------------------------------------------
	// Check TimeEntry Obj ID
	// --------------------------------------------------------------	
	if createTimeEntryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TimeEntry" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTimeEntryRequestResult := dao.GetTimeEntry( uint64(createTimeEntryObj.ID) )
	
	if getTimeEntryRequestResult.Success == false {
		t.Errorf(getTimeEntryRequestResult.Msg)
	} else {
		fmt.Println("Check Get TimeEntry success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTimeEntryObj,_ := getTimeEntryRequestResult.Data. (model.TimeEntry)
	compareTimeEntry := cmp.Equal(createTimeEntryObj.ID, getTimeEntryObj.ID)
	
	if  compareTimeEntry == false	{
		t.Errorf( "Created TimeEntry object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTimeEntryRequestResult := dao.GetAllTimeEntry()

	if getAllTimeEntryRequestResult.Success == false {
			t.Errorf(getAllTimeEntryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TimeEntry success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTimeEntryObj []model.TimeEntry = getAllTimeEntryRequestResult.Data. ([]model.TimeEntry)
		
	equalTimeEntry := cmp.Equal(createTimeEntryObj.ID, getAllTimeEntryObj[len(getAllTimeEntryObj)-1].ID)
		
	if equalTimeEntry == false {
		t.Errorf( "Created object is not equal to the last entry in TimeEntry[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TimeEntry
	// --------------------------------------------------------------	
	deleteTimeEntryRequestResult := dao.DeleteTimeEntry(uint64(createTimeEntryObj.ID))

	if deleteTimeEntryRequestResult.Success == false {
			t.Errorf(deleteTimeEntryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TimeEntry success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTimeEntryRequestResult = dao.GetTimeEntry( uint64(createTimeEntryObj.ID) )
	
	if getTimeEntryRequestResult.Success == true {
		t.Errorf(getTimeEntryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestApprovalCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Approval
	//----------------------------------------------------------------------------
	ApprovalObj := model.Approval                                                                                                    {ApproverComment:"test value for ApproverComment",ActionDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createApprovalRequestResult := dao.CreateApproval( ApprovalObj )
	
	if createApprovalRequestResult.Success == false {
		t.Errorf(createApprovalRequestResult.Msg)
	} else {
		fmt.Println("Check Create Approval success...")
	}
	
	createApprovalObj,_ := createApprovalRequestResult.Data. (model.Approval)

	// --------------------------------------------------------------
	// Check Approval Obj ID
	// --------------------------------------------------------------	
	if createApprovalObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Approval" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getApprovalRequestResult := dao.GetApproval( uint64(createApprovalObj.ID) )
	
	if getApprovalRequestResult.Success == false {
		t.Errorf(getApprovalRequestResult.Msg)
	} else {
		fmt.Println("Check Get Approval success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getApprovalObj,_ := getApprovalRequestResult.Data. (model.Approval)
	compareApproval := cmp.Equal(createApprovalObj.ID, getApprovalObj.ID)
	
	if  compareApproval == false	{
		t.Errorf( "Created Approval object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllApprovalRequestResult := dao.GetAllApproval()

	if getAllApprovalRequestResult.Success == false {
			t.Errorf(getAllApprovalRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Approval success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllApprovalObj []model.Approval = getAllApprovalRequestResult.Data. ([]model.Approval)
		
	equalApproval := cmp.Equal(createApprovalObj.ID, getAllApprovalObj[len(getAllApprovalObj)-1].ID)
		
	if equalApproval == false {
		t.Errorf( "Created object is not equal to the last entry in Approval[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Approval
	// --------------------------------------------------------------	
	deleteApprovalRequestResult := dao.DeleteApproval(uint64(createApprovalObj.ID))

	if deleteApprovalRequestResult.Success == false {
			t.Errorf(deleteApprovalRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Approval success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getApprovalRequestResult = dao.GetApproval( uint64(createApprovalObj.ID) )
	
	if getApprovalRequestResult.Success == true {
		t.Errorf(getApprovalRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLeavePolicyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LeavePolicy
	//----------------------------------------------------------------------------
	LeavePolicyObj := model.LeavePolicy                                                                                                                                                                                                            {Name:"test value for Name",AccrualRate:"test value",CarryoverAllowed:true,MaxBalance:"test value",LeaveCategory:0,AccrualUnit:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLeavePolicyRequestResult := dao.CreateLeavePolicy( LeavePolicyObj )
	
	if createLeavePolicyRequestResult.Success == false {
		t.Errorf(createLeavePolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Create LeavePolicy success...")
	}
	
	createLeavePolicyObj,_ := createLeavePolicyRequestResult.Data. (model.LeavePolicy)

	// --------------------------------------------------------------
	// Check LeavePolicy Obj ID
	// --------------------------------------------------------------	
	if createLeavePolicyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LeavePolicy" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLeavePolicyRequestResult := dao.GetLeavePolicy( uint64(createLeavePolicyObj.ID) )
	
	if getLeavePolicyRequestResult.Success == false {
		t.Errorf(getLeavePolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Get LeavePolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLeavePolicyObj,_ := getLeavePolicyRequestResult.Data. (model.LeavePolicy)
	compareLeavePolicy := cmp.Equal(createLeavePolicyObj.ID, getLeavePolicyObj.ID)
	
	if  compareLeavePolicy == false	{
		t.Errorf( "Created LeavePolicy object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLeavePolicyRequestResult := dao.GetAllLeavePolicy()

	if getAllLeavePolicyRequestResult.Success == false {
			t.Errorf(getAllLeavePolicyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LeavePolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLeavePolicyObj []model.LeavePolicy = getAllLeavePolicyRequestResult.Data. ([]model.LeavePolicy)
		
	equalLeavePolicy := cmp.Equal(createLeavePolicyObj.ID, getAllLeavePolicyObj[len(getAllLeavePolicyObj)-1].ID)
		
	if equalLeavePolicy == false {
		t.Errorf( "Created object is not equal to the last entry in LeavePolicy[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LeavePolicy
	// --------------------------------------------------------------	
	deleteLeavePolicyRequestResult := dao.DeleteLeavePolicy(uint64(createLeavePolicyObj.ID))

	if deleteLeavePolicyRequestResult.Success == false {
			t.Errorf(deleteLeavePolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LeavePolicy success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLeavePolicyRequestResult = dao.GetLeavePolicy( uint64(createLeavePolicyObj.ID) )
	
	if getLeavePolicyRequestResult.Success == true {
		t.Errorf(getLeavePolicyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLeaveRequestCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LeaveRequest
	//----------------------------------------------------------------------------
	LeaveRequestObj := model.LeaveRequest                                                                                                                                                                                                                                                    {RequestNumber:"test value for RequestNumber",StartDate:time.Now(),EndDate:time.Now(),Reason:"test value for Reason",Hours:"test value",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLeaveRequestRequestResult := dao.CreateLeaveRequest( LeaveRequestObj )
	
	if createLeaveRequestRequestResult.Success == false {
		t.Errorf(createLeaveRequestRequestResult.Msg)
	} else {
		fmt.Println("Check Create LeaveRequest success...")
	}
	
	createLeaveRequestObj,_ := createLeaveRequestRequestResult.Data. (model.LeaveRequest)

	// --------------------------------------------------------------
	// Check LeaveRequest Obj ID
	// --------------------------------------------------------------	
	if createLeaveRequestObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LeaveRequest" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLeaveRequestRequestResult := dao.GetLeaveRequest( uint64(createLeaveRequestObj.ID) )
	
	if getLeaveRequestRequestResult.Success == false {
		t.Errorf(getLeaveRequestRequestResult.Msg)
	} else {
		fmt.Println("Check Get LeaveRequest success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLeaveRequestObj,_ := getLeaveRequestRequestResult.Data. (model.LeaveRequest)
	compareLeaveRequest := cmp.Equal(createLeaveRequestObj.ID, getLeaveRequestObj.ID)
	
	if  compareLeaveRequest == false	{
		t.Errorf( "Created LeaveRequest object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLeaveRequestRequestResult := dao.GetAllLeaveRequest()

	if getAllLeaveRequestRequestResult.Success == false {
			t.Errorf(getAllLeaveRequestRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LeaveRequest success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLeaveRequestObj []model.LeaveRequest = getAllLeaveRequestRequestResult.Data. ([]model.LeaveRequest)
		
	equalLeaveRequest := cmp.Equal(createLeaveRequestObj.ID, getAllLeaveRequestObj[len(getAllLeaveRequestObj)-1].ID)
		
	if equalLeaveRequest == false {
		t.Errorf( "Created object is not equal to the last entry in LeaveRequest[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LeaveRequest
	// --------------------------------------------------------------	
	deleteLeaveRequestRequestResult := dao.DeleteLeaveRequest(uint64(createLeaveRequestObj.ID))

	if deleteLeaveRequestRequestResult.Success == false {
			t.Errorf(deleteLeaveRequestRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LeaveRequest success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLeaveRequestRequestResult = dao.GetLeaveRequest( uint64(createLeaveRequestObj.ID) )
	
	if getLeaveRequestRequestResult.Success == true {
		t.Errorf(getLeaveRequestRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPerformanceCycleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PerformanceCycle
	//----------------------------------------------------------------------------
	PerformanceCycleObj := model.PerformanceCycle                                                                                                                                                            {Name:"test value for Name",StartDate:time.Now(),EndDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPerformanceCycleRequestResult := dao.CreatePerformanceCycle( PerformanceCycleObj )
	
	if createPerformanceCycleRequestResult.Success == false {
		t.Errorf(createPerformanceCycleRequestResult.Msg)
	} else {
		fmt.Println("Check Create PerformanceCycle success...")
	}
	
	createPerformanceCycleObj,_ := createPerformanceCycleRequestResult.Data. (model.PerformanceCycle)

	// --------------------------------------------------------------
	// Check PerformanceCycle Obj ID
	// --------------------------------------------------------------	
	if createPerformanceCycleObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PerformanceCycle" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPerformanceCycleRequestResult := dao.GetPerformanceCycle( uint64(createPerformanceCycleObj.ID) )
	
	if getPerformanceCycleRequestResult.Success == false {
		t.Errorf(getPerformanceCycleRequestResult.Msg)
	} else {
		fmt.Println("Check Get PerformanceCycle success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPerformanceCycleObj,_ := getPerformanceCycleRequestResult.Data. (model.PerformanceCycle)
	comparePerformanceCycle := cmp.Equal(createPerformanceCycleObj.ID, getPerformanceCycleObj.ID)
	
	if  comparePerformanceCycle == false	{
		t.Errorf( "Created PerformanceCycle object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPerformanceCycleRequestResult := dao.GetAllPerformanceCycle()

	if getAllPerformanceCycleRequestResult.Success == false {
			t.Errorf(getAllPerformanceCycleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PerformanceCycle success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPerformanceCycleObj []model.PerformanceCycle = getAllPerformanceCycleRequestResult.Data. ([]model.PerformanceCycle)
		
	equalPerformanceCycle := cmp.Equal(createPerformanceCycleObj.ID, getAllPerformanceCycleObj[len(getAllPerformanceCycleObj)-1].ID)
		
	if equalPerformanceCycle == false {
		t.Errorf( "Created object is not equal to the last entry in PerformanceCycle[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PerformanceCycle
	// --------------------------------------------------------------	
	deletePerformanceCycleRequestResult := dao.DeletePerformanceCycle(uint64(createPerformanceCycleObj.ID))

	if deletePerformanceCycleRequestResult.Success == false {
			t.Errorf(deletePerformanceCycleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PerformanceCycle success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPerformanceCycleRequestResult = dao.GetPerformanceCycle( uint64(createPerformanceCycleObj.ID) )
	
	if getPerformanceCycleRequestResult.Success == true {
		t.Errorf(getPerformanceCycleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestGoalCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Goal
	//----------------------------------------------------------------------------
	GoalObj := model.Goal                                                                                                                                                    {Title:"test value for Title",Description:"test value for Description",TargetDate:time.Now(),Weight:new Percentage(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createGoalRequestResult := dao.CreateGoal( GoalObj )
	
	if createGoalRequestResult.Success == false {
		t.Errorf(createGoalRequestResult.Msg)
	} else {
		fmt.Println("Check Create Goal success...")
	}
	
	createGoalObj,_ := createGoalRequestResult.Data. (model.Goal)

	// --------------------------------------------------------------
	// Check Goal Obj ID
	// --------------------------------------------------------------	
	if createGoalObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Goal" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getGoalRequestResult := dao.GetGoal( uint64(createGoalObj.ID) )
	
	if getGoalRequestResult.Success == false {
		t.Errorf(getGoalRequestResult.Msg)
	} else {
		fmt.Println("Check Get Goal success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getGoalObj,_ := getGoalRequestResult.Data. (model.Goal)
	compareGoal := cmp.Equal(createGoalObj.ID, getGoalObj.ID)
	
	if  compareGoal == false	{
		t.Errorf( "Created Goal object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllGoalRequestResult := dao.GetAllGoal()

	if getAllGoalRequestResult.Success == false {
			t.Errorf(getAllGoalRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Goal success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllGoalObj []model.Goal = getAllGoalRequestResult.Data. ([]model.Goal)
		
	equalGoal := cmp.Equal(createGoalObj.ID, getAllGoalObj[len(getAllGoalObj)-1].ID)
		
	if equalGoal == false {
		t.Errorf( "Created object is not equal to the last entry in Goal[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Goal
	// --------------------------------------------------------------	
	deleteGoalRequestResult := dao.DeleteGoal(uint64(createGoalObj.ID))

	if deleteGoalRequestResult.Success == false {
			t.Errorf(deleteGoalRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Goal success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getGoalRequestResult = dao.GetGoal( uint64(createGoalObj.ID) )
	
	if getGoalRequestResult.Success == true {
		t.Errorf(getGoalRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPerformanceReviewCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PerformanceReview
	//----------------------------------------------------------------------------
	PerformanceReviewObj := model.PerformanceReview                                                                                                                                                    {ReviewNumber:"test value for ReviewNumber",ReviewDate:time.Now(),ReviewerComments:"test value for ReviewerComments",Rating:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPerformanceReviewRequestResult := dao.CreatePerformanceReview( PerformanceReviewObj )
	
	if createPerformanceReviewRequestResult.Success == false {
		t.Errorf(createPerformanceReviewRequestResult.Msg)
	} else {
		fmt.Println("Check Create PerformanceReview success...")
	}
	
	createPerformanceReviewObj,_ := createPerformanceReviewRequestResult.Data. (model.PerformanceReview)

	// --------------------------------------------------------------
	// Check PerformanceReview Obj ID
	// --------------------------------------------------------------	
	if createPerformanceReviewObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PerformanceReview" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPerformanceReviewRequestResult := dao.GetPerformanceReview( uint64(createPerformanceReviewObj.ID) )
	
	if getPerformanceReviewRequestResult.Success == false {
		t.Errorf(getPerformanceReviewRequestResult.Msg)
	} else {
		fmt.Println("Check Get PerformanceReview success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPerformanceReviewObj,_ := getPerformanceReviewRequestResult.Data. (model.PerformanceReview)
	comparePerformanceReview := cmp.Equal(createPerformanceReviewObj.ID, getPerformanceReviewObj.ID)
	
	if  comparePerformanceReview == false	{
		t.Errorf( "Created PerformanceReview object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPerformanceReviewRequestResult := dao.GetAllPerformanceReview()

	if getAllPerformanceReviewRequestResult.Success == false {
			t.Errorf(getAllPerformanceReviewRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PerformanceReview success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPerformanceReviewObj []model.PerformanceReview = getAllPerformanceReviewRequestResult.Data. ([]model.PerformanceReview)
		
	equalPerformanceReview := cmp.Equal(createPerformanceReviewObj.ID, getAllPerformanceReviewObj[len(getAllPerformanceReviewObj)-1].ID)
		
	if equalPerformanceReview == false {
		t.Errorf( "Created object is not equal to the last entry in PerformanceReview[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PerformanceReview
	// --------------------------------------------------------------	
	deletePerformanceReviewRequestResult := dao.DeletePerformanceReview(uint64(createPerformanceReviewObj.ID))

	if deletePerformanceReviewRequestResult.Success == false {
			t.Errorf(deletePerformanceReviewRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PerformanceReview success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPerformanceReviewRequestResult = dao.GetPerformanceReview( uint64(createPerformanceReviewObj.ID) )
	
	if getPerformanceReviewRequestResult.Success == true {
		t.Errorf(getPerformanceReviewRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCompetencyRatingCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CompetencyRating
	//----------------------------------------------------------------------------
	CompetencyRatingObj := model.CompetencyRating                                            {Comment:"test value for Comment",Rating:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCompetencyRatingRequestResult := dao.CreateCompetencyRating( CompetencyRatingObj )
	
	if createCompetencyRatingRequestResult.Success == false {
		t.Errorf(createCompetencyRatingRequestResult.Msg)
	} else {
		fmt.Println("Check Create CompetencyRating success...")
	}
	
	createCompetencyRatingObj,_ := createCompetencyRatingRequestResult.Data. (model.CompetencyRating)

	// --------------------------------------------------------------
	// Check CompetencyRating Obj ID
	// --------------------------------------------------------------	
	if createCompetencyRatingObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CompetencyRating" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCompetencyRatingRequestResult := dao.GetCompetencyRating( uint64(createCompetencyRatingObj.ID) )
	
	if getCompetencyRatingRequestResult.Success == false {
		t.Errorf(getCompetencyRatingRequestResult.Msg)
	} else {
		fmt.Println("Check Get CompetencyRating success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCompetencyRatingObj,_ := getCompetencyRatingRequestResult.Data. (model.CompetencyRating)
	compareCompetencyRating := cmp.Equal(createCompetencyRatingObj.ID, getCompetencyRatingObj.ID)
	
	if  compareCompetencyRating == false	{
		t.Errorf( "Created CompetencyRating object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCompetencyRatingRequestResult := dao.GetAllCompetencyRating()

	if getAllCompetencyRatingRequestResult.Success == false {
			t.Errorf(getAllCompetencyRatingRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CompetencyRating success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCompetencyRatingObj []model.CompetencyRating = getAllCompetencyRatingRequestResult.Data. ([]model.CompetencyRating)
		
	equalCompetencyRating := cmp.Equal(createCompetencyRatingObj.ID, getAllCompetencyRatingObj[len(getAllCompetencyRatingObj)-1].ID)
		
	if equalCompetencyRating == false {
		t.Errorf( "Created object is not equal to the last entry in CompetencyRating[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CompetencyRating
	// --------------------------------------------------------------	
	deleteCompetencyRatingRequestResult := dao.DeleteCompetencyRating(uint64(createCompetencyRatingObj.ID))

	if deleteCompetencyRatingRequestResult.Success == false {
			t.Errorf(deleteCompetencyRatingRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CompetencyRating success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCompetencyRatingRequestResult = dao.GetCompetencyRating( uint64(createCompetencyRatingObj.ID) )
	
	if getCompetencyRatingRequestResult.Success == true {
		t.Errorf(getCompetencyRatingRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTrainingCourseCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TrainingCourse
	//----------------------------------------------------------------------------
	TrainingCourseObj := model.TrainingCourse                                                                                                                                    {Code:"test value for Code",Title:"test value for Title",DurationHours:"test value",DeliveryMethod:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTrainingCourseRequestResult := dao.CreateTrainingCourse( TrainingCourseObj )
	
	if createTrainingCourseRequestResult.Success == false {
		t.Errorf(createTrainingCourseRequestResult.Msg)
	} else {
		fmt.Println("Check Create TrainingCourse success...")
	}
	
	createTrainingCourseObj,_ := createTrainingCourseRequestResult.Data. (model.TrainingCourse)

	// --------------------------------------------------------------
	// Check TrainingCourse Obj ID
	// --------------------------------------------------------------	
	if createTrainingCourseObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TrainingCourse" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTrainingCourseRequestResult := dao.GetTrainingCourse( uint64(createTrainingCourseObj.ID) )
	
	if getTrainingCourseRequestResult.Success == false {
		t.Errorf(getTrainingCourseRequestResult.Msg)
	} else {
		fmt.Println("Check Get TrainingCourse success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTrainingCourseObj,_ := getTrainingCourseRequestResult.Data. (model.TrainingCourse)
	compareTrainingCourse := cmp.Equal(createTrainingCourseObj.ID, getTrainingCourseObj.ID)
	
	if  compareTrainingCourse == false	{
		t.Errorf( "Created TrainingCourse object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTrainingCourseRequestResult := dao.GetAllTrainingCourse()

	if getAllTrainingCourseRequestResult.Success == false {
			t.Errorf(getAllTrainingCourseRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TrainingCourse success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTrainingCourseObj []model.TrainingCourse = getAllTrainingCourseRequestResult.Data. ([]model.TrainingCourse)
		
	equalTrainingCourse := cmp.Equal(createTrainingCourseObj.ID, getAllTrainingCourseObj[len(getAllTrainingCourseObj)-1].ID)
		
	if equalTrainingCourse == false {
		t.Errorf( "Created object is not equal to the last entry in TrainingCourse[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TrainingCourse
	// --------------------------------------------------------------	
	deleteTrainingCourseRequestResult := dao.DeleteTrainingCourse(uint64(createTrainingCourseObj.ID))

	if deleteTrainingCourseRequestResult.Success == false {
			t.Errorf(deleteTrainingCourseRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TrainingCourse success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTrainingCourseRequestResult = dao.GetTrainingCourse( uint64(createTrainingCourseObj.ID) )
	
	if getTrainingCourseRequestResult.Success == true {
		t.Errorf(getTrainingCourseRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTrainingEnrollmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TrainingEnrollment
	//----------------------------------------------------------------------------
	TrainingEnrollmentObj := model.TrainingEnrollment                                                                                                                                                            {EnrollmentNumber:"test value for EnrollmentNumber",CompletionDate:time.Now(),Score:"test value",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTrainingEnrollmentRequestResult := dao.CreateTrainingEnrollment( TrainingEnrollmentObj )
	
	if createTrainingEnrollmentRequestResult.Success == false {
		t.Errorf(createTrainingEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create TrainingEnrollment success...")
	}
	
	createTrainingEnrollmentObj,_ := createTrainingEnrollmentRequestResult.Data. (model.TrainingEnrollment)

	// --------------------------------------------------------------
	// Check TrainingEnrollment Obj ID
	// --------------------------------------------------------------	
	if createTrainingEnrollmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TrainingEnrollment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTrainingEnrollmentRequestResult := dao.GetTrainingEnrollment( uint64(createTrainingEnrollmentObj.ID) )
	
	if getTrainingEnrollmentRequestResult.Success == false {
		t.Errorf(getTrainingEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get TrainingEnrollment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTrainingEnrollmentObj,_ := getTrainingEnrollmentRequestResult.Data. (model.TrainingEnrollment)
	compareTrainingEnrollment := cmp.Equal(createTrainingEnrollmentObj.ID, getTrainingEnrollmentObj.ID)
	
	if  compareTrainingEnrollment == false	{
		t.Errorf( "Created TrainingEnrollment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTrainingEnrollmentRequestResult := dao.GetAllTrainingEnrollment()

	if getAllTrainingEnrollmentRequestResult.Success == false {
			t.Errorf(getAllTrainingEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TrainingEnrollment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTrainingEnrollmentObj []model.TrainingEnrollment = getAllTrainingEnrollmentRequestResult.Data. ([]model.TrainingEnrollment)
		
	equalTrainingEnrollment := cmp.Equal(createTrainingEnrollmentObj.ID, getAllTrainingEnrollmentObj[len(getAllTrainingEnrollmentObj)-1].ID)
		
	if equalTrainingEnrollment == false {
		t.Errorf( "Created object is not equal to the last entry in TrainingEnrollment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TrainingEnrollment
	// --------------------------------------------------------------	
	deleteTrainingEnrollmentRequestResult := dao.DeleteTrainingEnrollment(uint64(createTrainingEnrollmentObj.ID))

	if deleteTrainingEnrollmentRequestResult.Success == false {
			t.Errorf(deleteTrainingEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TrainingEnrollment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTrainingEnrollmentRequestResult = dao.GetTrainingEnrollment( uint64(createTrainingEnrollmentObj.ID) )
	
	if getTrainingEnrollmentRequestResult.Success == true {
		t.Errorf(getTrainingEnrollmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCertificationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Certification
	//----------------------------------------------------------------------------
	CertificationObj := model.Certification                                                                                                                                                                                                            {Name:"test value for Name",Issuer:"test value for Issuer",ValidFrom:time.Now(),ValidTo:time.Now(),CredentialId:"test value for CredentialId"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCertificationRequestResult := dao.CreateCertification( CertificationObj )
	
	if createCertificationRequestResult.Success == false {
		t.Errorf(createCertificationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Certification success...")
	}
	
	createCertificationObj,_ := createCertificationRequestResult.Data. (model.Certification)

	// --------------------------------------------------------------
	// Check Certification Obj ID
	// --------------------------------------------------------------	
	if createCertificationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Certification" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCertificationRequestResult := dao.GetCertification( uint64(createCertificationObj.ID) )
	
	if getCertificationRequestResult.Success == false {
		t.Errorf(getCertificationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Certification success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCertificationObj,_ := getCertificationRequestResult.Data. (model.Certification)
	compareCertification := cmp.Equal(createCertificationObj.ID, getCertificationObj.ID)
	
	if  compareCertification == false	{
		t.Errorf( "Created Certification object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCertificationRequestResult := dao.GetAllCertification()

	if getAllCertificationRequestResult.Success == false {
			t.Errorf(getAllCertificationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Certification success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCertificationObj []model.Certification = getAllCertificationRequestResult.Data. ([]model.Certification)
		
	equalCertification := cmp.Equal(createCertificationObj.ID, getAllCertificationObj[len(getAllCertificationObj)-1].ID)
		
	if equalCertification == false {
		t.Errorf( "Created object is not equal to the last entry in Certification[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Certification
	// --------------------------------------------------------------	
	deleteCertificationRequestResult := dao.DeleteCertification(uint64(createCertificationObj.ID))

	if deleteCertificationRequestResult.Success == false {
			t.Errorf(deleteCertificationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Certification success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCertificationRequestResult = dao.GetCertification( uint64(createCertificationObj.ID) )
	
	if getCertificationRequestResult.Success == true {
		t.Errorf(getCertificationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestJobRequisitionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for JobRequisition
	//----------------------------------------------------------------------------
	JobRequisitionObj := model.JobRequisition                                                                                                                                                                                    {RequisitionNumber:"test value for RequisitionNumber",Title:"test value for Title",Openings:100,TargetStartDate:time.Now(),Status:0,Priority:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createJobRequisitionRequestResult := dao.CreateJobRequisition( JobRequisitionObj )
	
	if createJobRequisitionRequestResult.Success == false {
		t.Errorf(createJobRequisitionRequestResult.Msg)
	} else {
		fmt.Println("Check Create JobRequisition success...")
	}
	
	createJobRequisitionObj,_ := createJobRequisitionRequestResult.Data. (model.JobRequisition)

	// --------------------------------------------------------------
	// Check JobRequisition Obj ID
	// --------------------------------------------------------------	
	if createJobRequisitionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for JobRequisition" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getJobRequisitionRequestResult := dao.GetJobRequisition( uint64(createJobRequisitionObj.ID) )
	
	if getJobRequisitionRequestResult.Success == false {
		t.Errorf(getJobRequisitionRequestResult.Msg)
	} else {
		fmt.Println("Check Get JobRequisition success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getJobRequisitionObj,_ := getJobRequisitionRequestResult.Data. (model.JobRequisition)
	compareJobRequisition := cmp.Equal(createJobRequisitionObj.ID, getJobRequisitionObj.ID)
	
	if  compareJobRequisition == false	{
		t.Errorf( "Created JobRequisition object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllJobRequisitionRequestResult := dao.GetAllJobRequisition()

	if getAllJobRequisitionRequestResult.Success == false {
			t.Errorf(getAllJobRequisitionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll JobRequisition success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllJobRequisitionObj []model.JobRequisition = getAllJobRequisitionRequestResult.Data. ([]model.JobRequisition)
		
	equalJobRequisition := cmp.Equal(createJobRequisitionObj.ID, getAllJobRequisitionObj[len(getAllJobRequisitionObj)-1].ID)
		
	if equalJobRequisition == false {
		t.Errorf( "Created object is not equal to the last entry in JobRequisition[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for JobRequisition
	// --------------------------------------------------------------	
	deleteJobRequisitionRequestResult := dao.DeleteJobRequisition(uint64(createJobRequisitionObj.ID))

	if deleteJobRequisitionRequestResult.Success == false {
			t.Errorf(deleteJobRequisitionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion JobRequisition success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getJobRequisitionRequestResult = dao.GetJobRequisition( uint64(createJobRequisitionObj.ID) )
	
	if getJobRequisitionRequestResult.Success == true {
		t.Errorf(getJobRequisitionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCandidateCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Candidate
	//----------------------------------------------------------------------------
	CandidateObj := model.Candidate                                                            {Name:new PersonName(),Email:new Email(),Phone:new PhoneNumber(),Source:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCandidateRequestResult := dao.CreateCandidate( CandidateObj )
	
	if createCandidateRequestResult.Success == false {
		t.Errorf(createCandidateRequestResult.Msg)
	} else {
		fmt.Println("Check Create Candidate success...")
	}
	
	createCandidateObj,_ := createCandidateRequestResult.Data. (model.Candidate)

	// --------------------------------------------------------------
	// Check Candidate Obj ID
	// --------------------------------------------------------------	
	if createCandidateObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Candidate" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCandidateRequestResult := dao.GetCandidate( uint64(createCandidateObj.ID) )
	
	if getCandidateRequestResult.Success == false {
		t.Errorf(getCandidateRequestResult.Msg)
	} else {
		fmt.Println("Check Get Candidate success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCandidateObj,_ := getCandidateRequestResult.Data. (model.Candidate)
	compareCandidate := cmp.Equal(createCandidateObj.ID, getCandidateObj.ID)
	
	if  compareCandidate == false	{
		t.Errorf( "Created Candidate object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCandidateRequestResult := dao.GetAllCandidate()

	if getAllCandidateRequestResult.Success == false {
			t.Errorf(getAllCandidateRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Candidate success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCandidateObj []model.Candidate = getAllCandidateRequestResult.Data. ([]model.Candidate)
		
	equalCandidate := cmp.Equal(createCandidateObj.ID, getAllCandidateObj[len(getAllCandidateObj)-1].ID)
		
	if equalCandidate == false {
		t.Errorf( "Created object is not equal to the last entry in Candidate[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Candidate
	// --------------------------------------------------------------	
	deleteCandidateRequestResult := dao.DeleteCandidate(uint64(createCandidateObj.ID))

	if deleteCandidateRequestResult.Success == false {
			t.Errorf(deleteCandidateRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Candidate success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCandidateRequestResult = dao.GetCandidate( uint64(createCandidateObj.ID) )
	
	if getCandidateRequestResult.Success == true {
		t.Errorf(getCandidateRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestJobApplicationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for JobApplication
	//----------------------------------------------------------------------------
	JobApplicationObj := model.JobApplication                                                                                                                                    {ApplicationNumber:"test value for ApplicationNumber",AppliedDate:time.Now(),ResumeUrl:"test value for ResumeUrl",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createJobApplicationRequestResult := dao.CreateJobApplication( JobApplicationObj )
	
	if createJobApplicationRequestResult.Success == false {
		t.Errorf(createJobApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check Create JobApplication success...")
	}
	
	createJobApplicationObj,_ := createJobApplicationRequestResult.Data. (model.JobApplication)

	// --------------------------------------------------------------
	// Check JobApplication Obj ID
	// --------------------------------------------------------------	
	if createJobApplicationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for JobApplication" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getJobApplicationRequestResult := dao.GetJobApplication( uint64(createJobApplicationObj.ID) )
	
	if getJobApplicationRequestResult.Success == false {
		t.Errorf(getJobApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check Get JobApplication success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getJobApplicationObj,_ := getJobApplicationRequestResult.Data. (model.JobApplication)
	compareJobApplication := cmp.Equal(createJobApplicationObj.ID, getJobApplicationObj.ID)
	
	if  compareJobApplication == false	{
		t.Errorf( "Created JobApplication object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllJobApplicationRequestResult := dao.GetAllJobApplication()

	if getAllJobApplicationRequestResult.Success == false {
			t.Errorf(getAllJobApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll JobApplication success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllJobApplicationObj []model.JobApplication = getAllJobApplicationRequestResult.Data. ([]model.JobApplication)
		
	equalJobApplication := cmp.Equal(createJobApplicationObj.ID, getAllJobApplicationObj[len(getAllJobApplicationObj)-1].ID)
		
	if equalJobApplication == false {
		t.Errorf( "Created object is not equal to the last entry in JobApplication[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for JobApplication
	// --------------------------------------------------------------	
	deleteJobApplicationRequestResult := dao.DeleteJobApplication(uint64(createJobApplicationObj.ID))

	if deleteJobApplicationRequestResult.Success == false {
			t.Errorf(deleteJobApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion JobApplication success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getJobApplicationRequestResult = dao.GetJobApplication( uint64(createJobApplicationObj.ID) )
	
	if getJobApplicationRequestResult.Success == true {
		t.Errorf(getJobApplicationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInterviewCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Interview
	//----------------------------------------------------------------------------
	InterviewObj := model.Interview                                                                                                                    {InterviewDate:time.Now(),Feedback:"test value for Feedback",Stage:0,Result:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInterviewRequestResult := dao.CreateInterview( InterviewObj )
	
	if createInterviewRequestResult.Success == false {
		t.Errorf(createInterviewRequestResult.Msg)
	} else {
		fmt.Println("Check Create Interview success...")
	}
	
	createInterviewObj,_ := createInterviewRequestResult.Data. (model.Interview)

	// --------------------------------------------------------------
	// Check Interview Obj ID
	// --------------------------------------------------------------	
	if createInterviewObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Interview" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInterviewRequestResult := dao.GetInterview( uint64(createInterviewObj.ID) )
	
	if getInterviewRequestResult.Success == false {
		t.Errorf(getInterviewRequestResult.Msg)
	} else {
		fmt.Println("Check Get Interview success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInterviewObj,_ := getInterviewRequestResult.Data. (model.Interview)
	compareInterview := cmp.Equal(createInterviewObj.ID, getInterviewObj.ID)
	
	if  compareInterview == false	{
		t.Errorf( "Created Interview object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInterviewRequestResult := dao.GetAllInterview()

	if getAllInterviewRequestResult.Success == false {
			t.Errorf(getAllInterviewRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Interview success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInterviewObj []model.Interview = getAllInterviewRequestResult.Data. ([]model.Interview)
		
	equalInterview := cmp.Equal(createInterviewObj.ID, getAllInterviewObj[len(getAllInterviewObj)-1].ID)
		
	if equalInterview == false {
		t.Errorf( "Created object is not equal to the last entry in Interview[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Interview
	// --------------------------------------------------------------	
	deleteInterviewRequestResult := dao.DeleteInterview(uint64(createInterviewObj.ID))

	if deleteInterviewRequestResult.Success == false {
			t.Errorf(deleteInterviewRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Interview success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInterviewRequestResult = dao.GetInterview( uint64(createInterviewObj.ID) )
	
	if getInterviewRequestResult.Success == true {
		t.Errorf(getInterviewRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestScreeningCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Screening
	//----------------------------------------------------------------------------
	ScreeningObj := model.Screening                                                                                                    {Name:"test value for Name",CompletedDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createScreeningRequestResult := dao.CreateScreening( ScreeningObj )
	
	if createScreeningRequestResult.Success == false {
		t.Errorf(createScreeningRequestResult.Msg)
	} else {
		fmt.Println("Check Create Screening success...")
	}
	
	createScreeningObj,_ := createScreeningRequestResult.Data. (model.Screening)

	// --------------------------------------------------------------
	// Check Screening Obj ID
	// --------------------------------------------------------------	
	if createScreeningObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Screening" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getScreeningRequestResult := dao.GetScreening( uint64(createScreeningObj.ID) )
	
	if getScreeningRequestResult.Success == false {
		t.Errorf(getScreeningRequestResult.Msg)
	} else {
		fmt.Println("Check Get Screening success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getScreeningObj,_ := getScreeningRequestResult.Data. (model.Screening)
	compareScreening := cmp.Equal(createScreeningObj.ID, getScreeningObj.ID)
	
	if  compareScreening == false	{
		t.Errorf( "Created Screening object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllScreeningRequestResult := dao.GetAllScreening()

	if getAllScreeningRequestResult.Success == false {
			t.Errorf(getAllScreeningRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Screening success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllScreeningObj []model.Screening = getAllScreeningRequestResult.Data. ([]model.Screening)
		
	equalScreening := cmp.Equal(createScreeningObj.ID, getAllScreeningObj[len(getAllScreeningObj)-1].ID)
		
	if equalScreening == false {
		t.Errorf( "Created object is not equal to the last entry in Screening[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Screening
	// --------------------------------------------------------------	
	deleteScreeningRequestResult := dao.DeleteScreening(uint64(createScreeningObj.ID))

	if deleteScreeningRequestResult.Success == false {
			t.Errorf(deleteScreeningRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Screening success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getScreeningRequestResult = dao.GetScreening( uint64(createScreeningObj.ID) )
	
	if getScreeningRequestResult.Success == true {
		t.Errorf(getScreeningRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestOfferCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Offer
	//----------------------------------------------------------------------------
	OfferObj := model.Offer                                                                                                                                    {OfferNumber:"test value for OfferNumber",ProposedStartDate:time.Now(),BaseSalary:new Money(),SignOnBonus:new Money(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOfferRequestResult := dao.CreateOffer( OfferObj )
	
	if createOfferRequestResult.Success == false {
		t.Errorf(createOfferRequestResult.Msg)
	} else {
		fmt.Println("Check Create Offer success...")
	}
	
	createOfferObj,_ := createOfferRequestResult.Data. (model.Offer)

	// --------------------------------------------------------------
	// Check Offer Obj ID
	// --------------------------------------------------------------	
	if createOfferObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Offer" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOfferRequestResult := dao.GetOffer( uint64(createOfferObj.ID) )
	
	if getOfferRequestResult.Success == false {
		t.Errorf(getOfferRequestResult.Msg)
	} else {
		fmt.Println("Check Get Offer success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOfferObj,_ := getOfferRequestResult.Data. (model.Offer)
	compareOffer := cmp.Equal(createOfferObj.ID, getOfferObj.ID)
	
	if  compareOffer == false	{
		t.Errorf( "Created Offer object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOfferRequestResult := dao.GetAllOffer()

	if getAllOfferRequestResult.Success == false {
			t.Errorf(getAllOfferRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Offer success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOfferObj []model.Offer = getAllOfferRequestResult.Data. ([]model.Offer)
		
	equalOffer := cmp.Equal(createOfferObj.ID, getAllOfferObj[len(getAllOfferObj)-1].ID)
		
	if equalOffer == false {
		t.Errorf( "Created object is not equal to the last entry in Offer[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Offer
	// --------------------------------------------------------------	
	deleteOfferRequestResult := dao.DeleteOffer(uint64(createOfferObj.ID))

	if deleteOfferRequestResult.Success == false {
			t.Errorf(deleteOfferRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Offer success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOfferRequestResult = dao.GetOffer( uint64(createOfferObj.ID) )
	
	if getOfferRequestResult.Success == true {
		t.Errorf(getOfferRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestOnboardingTaskCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for OnboardingTask
	//----------------------------------------------------------------------------
	OnboardingTaskObj := model.OnboardingTask                                                                                                                                    {TaskNumber:"test value for TaskNumber",Name:"test value for Name",DueDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOnboardingTaskRequestResult := dao.CreateOnboardingTask( OnboardingTaskObj )
	
	if createOnboardingTaskRequestResult.Success == false {
		t.Errorf(createOnboardingTaskRequestResult.Msg)
	} else {
		fmt.Println("Check Create OnboardingTask success...")
	}
	
	createOnboardingTaskObj,_ := createOnboardingTaskRequestResult.Data. (model.OnboardingTask)

	// --------------------------------------------------------------
	// Check OnboardingTask Obj ID
	// --------------------------------------------------------------	
	if createOnboardingTaskObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for OnboardingTask" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOnboardingTaskRequestResult := dao.GetOnboardingTask( uint64(createOnboardingTaskObj.ID) )
	
	if getOnboardingTaskRequestResult.Success == false {
		t.Errorf(getOnboardingTaskRequestResult.Msg)
	} else {
		fmt.Println("Check Get OnboardingTask success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOnboardingTaskObj,_ := getOnboardingTaskRequestResult.Data. (model.OnboardingTask)
	compareOnboardingTask := cmp.Equal(createOnboardingTaskObj.ID, getOnboardingTaskObj.ID)
	
	if  compareOnboardingTask == false	{
		t.Errorf( "Created OnboardingTask object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOnboardingTaskRequestResult := dao.GetAllOnboardingTask()

	if getAllOnboardingTaskRequestResult.Success == false {
			t.Errorf(getAllOnboardingTaskRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll OnboardingTask success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOnboardingTaskObj []model.OnboardingTask = getAllOnboardingTaskRequestResult.Data. ([]model.OnboardingTask)
		
	equalOnboardingTask := cmp.Equal(createOnboardingTaskObj.ID, getAllOnboardingTaskObj[len(getAllOnboardingTaskObj)-1].ID)
		
	if equalOnboardingTask == false {
		t.Errorf( "Created object is not equal to the last entry in OnboardingTask[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for OnboardingTask
	// --------------------------------------------------------------	
	deleteOnboardingTaskRequestResult := dao.DeleteOnboardingTask(uint64(createOnboardingTaskObj.ID))

	if deleteOnboardingTaskRequestResult.Success == false {
			t.Errorf(deleteOnboardingTaskRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion OnboardingTask success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOnboardingTaskRequestResult = dao.GetOnboardingTask( uint64(createOnboardingTaskObj.ID) )
	
	if getOnboardingTaskRequestResult.Success == true {
		t.Errorf(getOnboardingTaskRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBackgroundCheckCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BackgroundCheck
	//----------------------------------------------------------------------------
	BackgroundCheckObj := model.BackgroundCheck                                                                                                                                    {CheckNumber:"test value for CheckNumber",Provider:"test value for Provider",CompletedDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBackgroundCheckRequestResult := dao.CreateBackgroundCheck( BackgroundCheckObj )
	
	if createBackgroundCheckRequestResult.Success == false {
		t.Errorf(createBackgroundCheckRequestResult.Msg)
	} else {
		fmt.Println("Check Create BackgroundCheck success...")
	}
	
	createBackgroundCheckObj,_ := createBackgroundCheckRequestResult.Data. (model.BackgroundCheck)

	// --------------------------------------------------------------
	// Check BackgroundCheck Obj ID
	// --------------------------------------------------------------	
	if createBackgroundCheckObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BackgroundCheck" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBackgroundCheckRequestResult := dao.GetBackgroundCheck( uint64(createBackgroundCheckObj.ID) )
	
	if getBackgroundCheckRequestResult.Success == false {
		t.Errorf(getBackgroundCheckRequestResult.Msg)
	} else {
		fmt.Println("Check Get BackgroundCheck success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBackgroundCheckObj,_ := getBackgroundCheckRequestResult.Data. (model.BackgroundCheck)
	compareBackgroundCheck := cmp.Equal(createBackgroundCheckObj.ID, getBackgroundCheckObj.ID)
	
	if  compareBackgroundCheck == false	{
		t.Errorf( "Created BackgroundCheck object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBackgroundCheckRequestResult := dao.GetAllBackgroundCheck()

	if getAllBackgroundCheckRequestResult.Success == false {
			t.Errorf(getAllBackgroundCheckRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BackgroundCheck success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBackgroundCheckObj []model.BackgroundCheck = getAllBackgroundCheckRequestResult.Data. ([]model.BackgroundCheck)
		
	equalBackgroundCheck := cmp.Equal(createBackgroundCheckObj.ID, getAllBackgroundCheckObj[len(getAllBackgroundCheckObj)-1].ID)
		
	if equalBackgroundCheck == false {
		t.Errorf( "Created object is not equal to the last entry in BackgroundCheck[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BackgroundCheck
	// --------------------------------------------------------------	
	deleteBackgroundCheckRequestResult := dao.DeleteBackgroundCheck(uint64(createBackgroundCheckObj.ID))

	if deleteBackgroundCheckRequestResult.Success == false {
			t.Errorf(deleteBackgroundCheckRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BackgroundCheck success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBackgroundCheckRequestResult = dao.GetBackgroundCheck( uint64(createBackgroundCheckObj.ID) )
	
	if getBackgroundCheckRequestResult.Success == true {
		t.Errorf(getBackgroundCheckRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDocumentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Document
	//----------------------------------------------------------------------------
	DocumentObj := model.Document                                                                                                                                    {Name:"test value for Name",FileUrl:"test value for FileUrl",UploadedDate:time.Now(),DocumentType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDocumentRequestResult := dao.CreateDocument( DocumentObj )
	
	if createDocumentRequestResult.Success == false {
		t.Errorf(createDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Document success...")
	}
	
	createDocumentObj,_ := createDocumentRequestResult.Data. (model.Document)

	// --------------------------------------------------------------
	// Check Document Obj ID
	// --------------------------------------------------------------	
	if createDocumentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Document" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDocumentRequestResult := dao.GetDocument( uint64(createDocumentObj.ID) )
	
	if getDocumentRequestResult.Success == false {
		t.Errorf(getDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Document success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDocumentObj,_ := getDocumentRequestResult.Data. (model.Document)
	compareDocument := cmp.Equal(createDocumentObj.ID, getDocumentObj.ID)
	
	if  compareDocument == false	{
		t.Errorf( "Created Document object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDocumentRequestResult := dao.GetAllDocument()

	if getAllDocumentRequestResult.Success == false {
			t.Errorf(getAllDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Document success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDocumentObj []model.Document = getAllDocumentRequestResult.Data. ([]model.Document)
		
	equalDocument := cmp.Equal(createDocumentObj.ID, getAllDocumentObj[len(getAllDocumentObj)-1].ID)
		
	if equalDocument == false {
		t.Errorf( "Created object is not equal to the last entry in Document[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Document
	// --------------------------------------------------------------	
	deleteDocumentRequestResult := dao.DeleteDocument(uint64(createDocumentObj.ID))

	if deleteDocumentRequestResult.Success == false {
			t.Errorf(deleteDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Document success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDocumentRequestResult = dao.GetDocument( uint64(createDocumentObj.ID) )
	
	if getDocumentRequestResult.Success == true {
		t.Errorf(getDocumentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPolicyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Policy
	//----------------------------------------------------------------------------
	PolicyObj := model.Policy                                                                                                                                                    {PolicyNumber:"test value for PolicyNumber",Name:"test value for Name",EffectiveDate:time.Now(),Description:"test value for Description"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPolicyRequestResult := dao.CreatePolicy( PolicyObj )
	
	if createPolicyRequestResult.Success == false {
		t.Errorf(createPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Create Policy success...")
	}
	
	createPolicyObj,_ := createPolicyRequestResult.Data. (model.Policy)

	// --------------------------------------------------------------
	// Check Policy Obj ID
	// --------------------------------------------------------------	
	if createPolicyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Policy" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPolicyRequestResult := dao.GetPolicy( uint64(createPolicyObj.ID) )
	
	if getPolicyRequestResult.Success == false {
		t.Errorf(getPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Get Policy success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPolicyObj,_ := getPolicyRequestResult.Data. (model.Policy)
	comparePolicy := cmp.Equal(createPolicyObj.ID, getPolicyObj.ID)
	
	if  comparePolicy == false	{
		t.Errorf( "Created Policy object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPolicyRequestResult := dao.GetAllPolicy()

	if getAllPolicyRequestResult.Success == false {
			t.Errorf(getAllPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Policy success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPolicyObj []model.Policy = getAllPolicyRequestResult.Data. ([]model.Policy)
		
	equalPolicy := cmp.Equal(createPolicyObj.ID, getAllPolicyObj[len(getAllPolicyObj)-1].ID)
		
	if equalPolicy == false {
		t.Errorf( "Created object is not equal to the last entry in Policy[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Policy
	// --------------------------------------------------------------	
	deletePolicyRequestResult := dao.DeletePolicy(uint64(createPolicyObj.ID))

	if deletePolicyRequestResult.Success == false {
			t.Errorf(deletePolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Policy success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPolicyRequestResult = dao.GetPolicy( uint64(createPolicyObj.ID) )
	
	if getPolicyRequestResult.Success == true {
		t.Errorf(getPolicyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPolicyAcknowledgementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PolicyAcknowledgement
	//----------------------------------------------------------------------------
	PolicyAcknowledgementObj := model.PolicyAcknowledgement                                                                    {AcknowledgementDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPolicyAcknowledgementRequestResult := dao.CreatePolicyAcknowledgement( PolicyAcknowledgementObj )
	
	if createPolicyAcknowledgementRequestResult.Success == false {
		t.Errorf(createPolicyAcknowledgementRequestResult.Msg)
	} else {
		fmt.Println("Check Create PolicyAcknowledgement success...")
	}
	
	createPolicyAcknowledgementObj,_ := createPolicyAcknowledgementRequestResult.Data. (model.PolicyAcknowledgement)

	// --------------------------------------------------------------
	// Check PolicyAcknowledgement Obj ID
	// --------------------------------------------------------------	
	if createPolicyAcknowledgementObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PolicyAcknowledgement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPolicyAcknowledgementRequestResult := dao.GetPolicyAcknowledgement( uint64(createPolicyAcknowledgementObj.ID) )
	
	if getPolicyAcknowledgementRequestResult.Success == false {
		t.Errorf(getPolicyAcknowledgementRequestResult.Msg)
	} else {
		fmt.Println("Check Get PolicyAcknowledgement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPolicyAcknowledgementObj,_ := getPolicyAcknowledgementRequestResult.Data. (model.PolicyAcknowledgement)
	comparePolicyAcknowledgement := cmp.Equal(createPolicyAcknowledgementObj.ID, getPolicyAcknowledgementObj.ID)
	
	if  comparePolicyAcknowledgement == false	{
		t.Errorf( "Created PolicyAcknowledgement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPolicyAcknowledgementRequestResult := dao.GetAllPolicyAcknowledgement()

	if getAllPolicyAcknowledgementRequestResult.Success == false {
			t.Errorf(getAllPolicyAcknowledgementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PolicyAcknowledgement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPolicyAcknowledgementObj []model.PolicyAcknowledgement = getAllPolicyAcknowledgementRequestResult.Data. ([]model.PolicyAcknowledgement)
		
	equalPolicyAcknowledgement := cmp.Equal(createPolicyAcknowledgementObj.ID, getAllPolicyAcknowledgementObj[len(getAllPolicyAcknowledgementObj)-1].ID)
		
	if equalPolicyAcknowledgement == false {
		t.Errorf( "Created object is not equal to the last entry in PolicyAcknowledgement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PolicyAcknowledgement
	// --------------------------------------------------------------	
	deletePolicyAcknowledgementRequestResult := dao.DeletePolicyAcknowledgement(uint64(createPolicyAcknowledgementObj.ID))

	if deletePolicyAcknowledgementRequestResult.Success == false {
			t.Errorf(deletePolicyAcknowledgementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PolicyAcknowledgement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPolicyAcknowledgementRequestResult = dao.GetPolicyAcknowledgement( uint64(createPolicyAcknowledgementObj.ID) )
	
	if getPolicyAcknowledgementRequestResult.Success == true {
		t.Errorf(getPolicyAcknowledgementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTerminationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Termination
	//----------------------------------------------------------------------------
	TerminationObj := model.Termination                                                                                                                                                                                    {TerminationNumber:"test value for TerminationNumber",TerminationDate:time.Now(),Notes:"test value for Notes",EligibleForRehire:true,Reason:0,Type:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTerminationRequestResult := dao.CreateTermination( TerminationObj )
	
	if createTerminationRequestResult.Success == false {
		t.Errorf(createTerminationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Termination success...")
	}
	
	createTerminationObj,_ := createTerminationRequestResult.Data. (model.Termination)

	// --------------------------------------------------------------
	// Check Termination Obj ID
	// --------------------------------------------------------------	
	if createTerminationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Termination" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTerminationRequestResult := dao.GetTermination( uint64(createTerminationObj.ID) )
	
	if getTerminationRequestResult.Success == false {
		t.Errorf(getTerminationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Termination success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTerminationObj,_ := getTerminationRequestResult.Data. (model.Termination)
	compareTermination := cmp.Equal(createTerminationObj.ID, getTerminationObj.ID)
	
	if  compareTermination == false	{
		t.Errorf( "Created Termination object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTerminationRequestResult := dao.GetAllTermination()

	if getAllTerminationRequestResult.Success == false {
			t.Errorf(getAllTerminationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Termination success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTerminationObj []model.Termination = getAllTerminationRequestResult.Data. ([]model.Termination)
		
	equalTermination := cmp.Equal(createTerminationObj.ID, getAllTerminationObj[len(getAllTerminationObj)-1].ID)
		
	if equalTermination == false {
		t.Errorf( "Created object is not equal to the last entry in Termination[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Termination
	// --------------------------------------------------------------	
	deleteTerminationRequestResult := dao.DeleteTermination(uint64(createTerminationObj.ID))

	if deleteTerminationRequestResult.Success == false {
			t.Errorf(deleteTerminationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Termination success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTerminationRequestResult = dao.GetTermination( uint64(createTerminationObj.ID) )
	
	if getTerminationRequestResult.Success == true {
		t.Errorf(getTerminationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWorkAuthorizationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for WorkAuthorization
	//----------------------------------------------------------------------------
	WorkAuthorizationObj := model.WorkAuthorization                                                                                                    {Country:"test value for Country",ExpirationDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWorkAuthorizationRequestResult := dao.CreateWorkAuthorization( WorkAuthorizationObj )
	
	if createWorkAuthorizationRequestResult.Success == false {
		t.Errorf(createWorkAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Check Create WorkAuthorization success...")
	}
	
	createWorkAuthorizationObj,_ := createWorkAuthorizationRequestResult.Data. (model.WorkAuthorization)

	// --------------------------------------------------------------
	// Check WorkAuthorization Obj ID
	// --------------------------------------------------------------	
	if createWorkAuthorizationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for WorkAuthorization" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWorkAuthorizationRequestResult := dao.GetWorkAuthorization( uint64(createWorkAuthorizationObj.ID) )
	
	if getWorkAuthorizationRequestResult.Success == false {
		t.Errorf(getWorkAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Check Get WorkAuthorization success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWorkAuthorizationObj,_ := getWorkAuthorizationRequestResult.Data. (model.WorkAuthorization)
	compareWorkAuthorization := cmp.Equal(createWorkAuthorizationObj.ID, getWorkAuthorizationObj.ID)
	
	if  compareWorkAuthorization == false	{
		t.Errorf( "Created WorkAuthorization object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWorkAuthorizationRequestResult := dao.GetAllWorkAuthorization()

	if getAllWorkAuthorizationRequestResult.Success == false {
			t.Errorf(getAllWorkAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll WorkAuthorization success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWorkAuthorizationObj []model.WorkAuthorization = getAllWorkAuthorizationRequestResult.Data. ([]model.WorkAuthorization)
		
	equalWorkAuthorization := cmp.Equal(createWorkAuthorizationObj.ID, getAllWorkAuthorizationObj[len(getAllWorkAuthorizationObj)-1].ID)
		
	if equalWorkAuthorization == false {
		t.Errorf( "Created object is not equal to the last entry in WorkAuthorization[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for WorkAuthorization
	// --------------------------------------------------------------	
	deleteWorkAuthorizationRequestResult := dao.DeleteWorkAuthorization(uint64(createWorkAuthorizationObj.ID))

	if deleteWorkAuthorizationRequestResult.Success == false {
			t.Errorf(deleteWorkAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion WorkAuthorization success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWorkAuthorizationRequestResult = dao.GetWorkAuthorization( uint64(createWorkAuthorizationObj.ID) )
	
	if getWorkAuthorizationRequestResult.Success == true {
		t.Errorf(getWorkAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBankAccountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BankAccount
	//----------------------------------------------------------------------------
	BankAccountObj := model.BankAccount                                                                                                                                                                                            {AccountHolder:"test value for AccountHolder",BankName:"test value for BankName",Iban:"test value for Iban",Bic:"test value for Bic",AccountNumber:"test value for AccountNumber",RoutingNumber:"test value for RoutingNumber"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBankAccountRequestResult := dao.CreateBankAccount( BankAccountObj )
	
	if createBankAccountRequestResult.Success == false {
		t.Errorf(createBankAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Create BankAccount success...")
	}
	
	createBankAccountObj,_ := createBankAccountRequestResult.Data. (model.BankAccount)

	// --------------------------------------------------------------
	// Check BankAccount Obj ID
	// --------------------------------------------------------------	
	if createBankAccountObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BankAccount" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBankAccountRequestResult := dao.GetBankAccount( uint64(createBankAccountObj.ID) )
	
	if getBankAccountRequestResult.Success == false {
		t.Errorf(getBankAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Get BankAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBankAccountObj,_ := getBankAccountRequestResult.Data. (model.BankAccount)
	compareBankAccount := cmp.Equal(createBankAccountObj.ID, getBankAccountObj.ID)
	
	if  compareBankAccount == false	{
		t.Errorf( "Created BankAccount object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBankAccountRequestResult := dao.GetAllBankAccount()

	if getAllBankAccountRequestResult.Success == false {
			t.Errorf(getAllBankAccountRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BankAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBankAccountObj []model.BankAccount = getAllBankAccountRequestResult.Data. ([]model.BankAccount)
		
	equalBankAccount := cmp.Equal(createBankAccountObj.ID, getAllBankAccountObj[len(getAllBankAccountObj)-1].ID)
		
	if equalBankAccount == false {
		t.Errorf( "Created object is not equal to the last entry in BankAccount[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BankAccount
	// --------------------------------------------------------------	
	deleteBankAccountRequestResult := dao.DeleteBankAccount(uint64(createBankAccountObj.ID))

	if deleteBankAccountRequestResult.Success == false {
			t.Errorf(deleteBankAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BankAccount success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBankAccountRequestResult = dao.GetBankAccount( uint64(createBankAccountObj.ID) )
	
	if getBankAccountRequestResult.Success == true {
		t.Errorf(getBankAccountRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

