package test

import ( 
	"testing"
    dao "healthcare-on-golang/internal/dao"
	"healthcare-on-golang/internal/model"
	"healthcare-on-golang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestHealthSystemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for HealthSystem
	//----------------------------------------------------------------------------
	HealthSystemObj := model.HealthSystem                                                                                                                            {Name:"test value for Name",LegalName:"test value for LegalName",HeadquartersCountry:"test value for HeadquartersCountry",Website:"test value for Website"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createHealthSystemRequestResult := dao.CreateHealthSystem( HealthSystemObj )
	
	if createHealthSystemRequestResult.Success == false {
		t.Errorf(createHealthSystemRequestResult.Msg)
	} else {
		fmt.Println("Check Create HealthSystem success...")
	}
	
	createHealthSystemObj,_ := createHealthSystemRequestResult.Data. (model.HealthSystem)

	// --------------------------------------------------------------
	// Check HealthSystem Obj ID
	// --------------------------------------------------------------	
	if createHealthSystemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for HealthSystem" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getHealthSystemRequestResult := dao.GetHealthSystem( uint64(createHealthSystemObj.ID) )
	
	if getHealthSystemRequestResult.Success == false {
		t.Errorf(getHealthSystemRequestResult.Msg)
	} else {
		fmt.Println("Check Get HealthSystem success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getHealthSystemObj,_ := getHealthSystemRequestResult.Data. (model.HealthSystem)
	compareHealthSystem := cmp.Equal(createHealthSystemObj.ID, getHealthSystemObj.ID)
	
	if  compareHealthSystem == false	{
		t.Errorf( "Created HealthSystem object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllHealthSystemRequestResult := dao.GetAllHealthSystem()

	if getAllHealthSystemRequestResult.Success == false {
			t.Errorf(getAllHealthSystemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll HealthSystem success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllHealthSystemObj []model.HealthSystem = getAllHealthSystemRequestResult.Data. ([]model.HealthSystem)
		
	equalHealthSystem := cmp.Equal(createHealthSystemObj.ID, getAllHealthSystemObj[len(getAllHealthSystemObj)-1].ID)
		
	if equalHealthSystem == false {
		t.Errorf( "Created object is not equal to the last entry in HealthSystem[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for HealthSystem
	// --------------------------------------------------------------	
	deleteHealthSystemRequestResult := dao.DeleteHealthSystem(uint64(createHealthSystemObj.ID))

	if deleteHealthSystemRequestResult.Success == false {
			t.Errorf(deleteHealthSystemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion HealthSystem success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getHealthSystemRequestResult = dao.GetHealthSystem( uint64(createHealthSystemObj.ID) )
	
	if getHealthSystemRequestResult.Success == true {
		t.Errorf(getHealthSystemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFacilityCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Facility
	//----------------------------------------------------------------------------
	FacilityObj := model.Facility                                                                                            {Name:"test value for Name",FacilityCode:"test value for FacilityCode",Address:new Address(),FacilityType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFacilityRequestResult := dao.CreateFacility( FacilityObj )
	
	if createFacilityRequestResult.Success == false {
		t.Errorf(createFacilityRequestResult.Msg)
	} else {
		fmt.Println("Check Create Facility success...")
	}
	
	createFacilityObj,_ := createFacilityRequestResult.Data. (model.Facility)

	// --------------------------------------------------------------
	// Check Facility Obj ID
	// --------------------------------------------------------------	
	if createFacilityObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Facility" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFacilityRequestResult := dao.GetFacility( uint64(createFacilityObj.ID) )
	
	if getFacilityRequestResult.Success == false {
		t.Errorf(getFacilityRequestResult.Msg)
	} else {
		fmt.Println("Check Get Facility success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFacilityObj,_ := getFacilityRequestResult.Data. (model.Facility)
	compareFacility := cmp.Equal(createFacilityObj.ID, getFacilityObj.ID)
	
	if  compareFacility == false	{
		t.Errorf( "Created Facility object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFacilityRequestResult := dao.GetAllFacility()

	if getAllFacilityRequestResult.Success == false {
			t.Errorf(getAllFacilityRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Facility success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFacilityObj []model.Facility = getAllFacilityRequestResult.Data. ([]model.Facility)
		
	equalFacility := cmp.Equal(createFacilityObj.ID, getAllFacilityObj[len(getAllFacilityObj)-1].ID)
		
	if equalFacility == false {
		t.Errorf( "Created object is not equal to the last entry in Facility[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Facility
	// --------------------------------------------------------------	
	deleteFacilityRequestResult := dao.DeleteFacility(uint64(createFacilityObj.ID))

	if deleteFacilityRequestResult.Success == false {
			t.Errorf(deleteFacilityRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Facility success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFacilityRequestResult = dao.GetFacility( uint64(createFacilityObj.ID) )
	
	if getFacilityRequestResult.Success == true {
		t.Errorf(getFacilityRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDepartmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Department
	//----------------------------------------------------------------------------
	DepartmentObj := model.Department                                            {Name:"test value for Name",DepartmentType:0}

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


func TestCareTeamCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CareTeam
	//----------------------------------------------------------------------------
	CareTeamObj := model.CareTeam                                            {Name:"test value for Name",CareSetting:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCareTeamRequestResult := dao.CreateCareTeam( CareTeamObj )
	
	if createCareTeamRequestResult.Success == false {
		t.Errorf(createCareTeamRequestResult.Msg)
	} else {
		fmt.Println("Check Create CareTeam success...")
	}
	
	createCareTeamObj,_ := createCareTeamRequestResult.Data. (model.CareTeam)

	// --------------------------------------------------------------
	// Check CareTeam Obj ID
	// --------------------------------------------------------------	
	if createCareTeamObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CareTeam" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCareTeamRequestResult := dao.GetCareTeam( uint64(createCareTeamObj.ID) )
	
	if getCareTeamRequestResult.Success == false {
		t.Errorf(getCareTeamRequestResult.Msg)
	} else {
		fmt.Println("Check Get CareTeam success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCareTeamObj,_ := getCareTeamRequestResult.Data. (model.CareTeam)
	compareCareTeam := cmp.Equal(createCareTeamObj.ID, getCareTeamObj.ID)
	
	if  compareCareTeam == false	{
		t.Errorf( "Created CareTeam object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCareTeamRequestResult := dao.GetAllCareTeam()

	if getAllCareTeamRequestResult.Success == false {
			t.Errorf(getAllCareTeamRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CareTeam success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCareTeamObj []model.CareTeam = getAllCareTeamRequestResult.Data. ([]model.CareTeam)
		
	equalCareTeam := cmp.Equal(createCareTeamObj.ID, getAllCareTeamObj[len(getAllCareTeamObj)-1].ID)
		
	if equalCareTeam == false {
		t.Errorf( "Created object is not equal to the last entry in CareTeam[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CareTeam
	// --------------------------------------------------------------	
	deleteCareTeamRequestResult := dao.DeleteCareTeam(uint64(createCareTeamObj.ID))

	if deleteCareTeamRequestResult.Success == false {
			t.Errorf(deleteCareTeamRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CareTeam success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCareTeamRequestResult = dao.GetCareTeam( uint64(createCareTeamObj.ID) )
	
	if getCareTeamRequestResult.Success == true {
		t.Errorf(getCareTeamRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestClinicianCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Clinician
	//----------------------------------------------------------------------------
	ClinicianObj := model.Clinician                                                                                                                            {FirstName:"test value for FirstName",LastName:"test value for LastName",LicenseNumber:"test value for LicenseNumber",ClinicianType:0,Specialty:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createClinicianRequestResult := dao.CreateClinician( ClinicianObj )
	
	if createClinicianRequestResult.Success == false {
		t.Errorf(createClinicianRequestResult.Msg)
	} else {
		fmt.Println("Check Create Clinician success...")
	}
	
	createClinicianObj,_ := createClinicianRequestResult.Data. (model.Clinician)

	// --------------------------------------------------------------
	// Check Clinician Obj ID
	// --------------------------------------------------------------	
	if createClinicianObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Clinician" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getClinicianRequestResult := dao.GetClinician( uint64(createClinicianObj.ID) )
	
	if getClinicianRequestResult.Success == false {
		t.Errorf(getClinicianRequestResult.Msg)
	} else {
		fmt.Println("Check Get Clinician success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getClinicianObj,_ := getClinicianRequestResult.Data. (model.Clinician)
	compareClinician := cmp.Equal(createClinicianObj.ID, getClinicianObj.ID)
	
	if  compareClinician == false	{
		t.Errorf( "Created Clinician object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllClinicianRequestResult := dao.GetAllClinician()

	if getAllClinicianRequestResult.Success == false {
			t.Errorf(getAllClinicianRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Clinician success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllClinicianObj []model.Clinician = getAllClinicianRequestResult.Data. ([]model.Clinician)
		
	equalClinician := cmp.Equal(createClinicianObj.ID, getAllClinicianObj[len(getAllClinicianObj)-1].ID)
		
	if equalClinician == false {
		t.Errorf( "Created object is not equal to the last entry in Clinician[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Clinician
	// --------------------------------------------------------------	
	deleteClinicianRequestResult := dao.DeleteClinician(uint64(createClinicianObj.ID))

	if deleteClinicianRequestResult.Success == false {
			t.Errorf(deleteClinicianRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Clinician success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getClinicianRequestResult = dao.GetClinician( uint64(createClinicianObj.ID) )
	
	if getClinicianRequestResult.Success == true {
		t.Errorf(getClinicianRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPatientCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Patient
	//----------------------------------------------------------------------------
	PatientObj := model.Patient                                                                                                                                                                                                                    {FirstName:"test value for FirstName",LastName:"test value for LastName",Mrn:new MRN(),DateOfBirth:time.Now(),Address:new Address(),PrimaryLanguage:"test value for PrimaryLanguage",SexAtBirth:0,BloodType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPatientRequestResult := dao.CreatePatient( PatientObj )
	
	if createPatientRequestResult.Success == false {
		t.Errorf(createPatientRequestResult.Msg)
	} else {
		fmt.Println("Check Create Patient success...")
	}
	
	createPatientObj,_ := createPatientRequestResult.Data. (model.Patient)

	// --------------------------------------------------------------
	// Check Patient Obj ID
	// --------------------------------------------------------------	
	if createPatientObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Patient" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPatientRequestResult := dao.GetPatient( uint64(createPatientObj.ID) )
	
	if getPatientRequestResult.Success == false {
		t.Errorf(getPatientRequestResult.Msg)
	} else {
		fmt.Println("Check Get Patient success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPatientObj,_ := getPatientRequestResult.Data. (model.Patient)
	comparePatient := cmp.Equal(createPatientObj.ID, getPatientObj.ID)
	
	if  comparePatient == false	{
		t.Errorf( "Created Patient object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPatientRequestResult := dao.GetAllPatient()

	if getAllPatientRequestResult.Success == false {
			t.Errorf(getAllPatientRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Patient success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPatientObj []model.Patient = getAllPatientRequestResult.Data. ([]model.Patient)
		
	equalPatient := cmp.Equal(createPatientObj.ID, getAllPatientObj[len(getAllPatientObj)-1].ID)
		
	if equalPatient == false {
		t.Errorf( "Created object is not equal to the last entry in Patient[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Patient
	// --------------------------------------------------------------	
	deletePatientRequestResult := dao.DeletePatient(uint64(createPatientObj.ID))

	if deletePatientRequestResult.Success == false {
			t.Errorf(deletePatientRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Patient success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPatientRequestResult = dao.GetPatient( uint64(createPatientObj.ID) )
	
	if getPatientRequestResult.Success == true {
		t.Errorf(getPatientRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAppointmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Appointment
	//----------------------------------------------------------------------------
	AppointmentObj := model.Appointment                                                                                                                    {AppointmentDate:time.Now(),Reason:"test value for Reason",Status:0,Priority:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAppointmentRequestResult := dao.CreateAppointment( AppointmentObj )
	
	if createAppointmentRequestResult.Success == false {
		t.Errorf(createAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Appointment success...")
	}
	
	createAppointmentObj,_ := createAppointmentRequestResult.Data. (model.Appointment)

	// --------------------------------------------------------------
	// Check Appointment Obj ID
	// --------------------------------------------------------------	
	if createAppointmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Appointment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAppointmentRequestResult := dao.GetAppointment( uint64(createAppointmentObj.ID) )
	
	if getAppointmentRequestResult.Success == false {
		t.Errorf(getAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Appointment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAppointmentObj,_ := getAppointmentRequestResult.Data. (model.Appointment)
	compareAppointment := cmp.Equal(createAppointmentObj.ID, getAppointmentObj.ID)
	
	if  compareAppointment == false	{
		t.Errorf( "Created Appointment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAppointmentRequestResult := dao.GetAllAppointment()

	if getAllAppointmentRequestResult.Success == false {
			t.Errorf(getAllAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Appointment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAppointmentObj []model.Appointment = getAllAppointmentRequestResult.Data. ([]model.Appointment)
		
	equalAppointment := cmp.Equal(createAppointmentObj.ID, getAllAppointmentObj[len(getAllAppointmentObj)-1].ID)
		
	if equalAppointment == false {
		t.Errorf( "Created object is not equal to the last entry in Appointment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Appointment
	// --------------------------------------------------------------	
	deleteAppointmentRequestResult := dao.DeleteAppointment(uint64(createAppointmentObj.ID))

	if deleteAppointmentRequestResult.Success == false {
			t.Errorf(deleteAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Appointment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAppointmentRequestResult = dao.GetAppointment( uint64(createAppointmentObj.ID) )
	
	if getAppointmentRequestResult.Success == true {
		t.Errorf(getAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEncounterCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Encounter
	//----------------------------------------------------------------------------
	EncounterObj := model.Encounter                                                                                                                                                                            {EncounterNumber:"test value for EncounterNumber",StartDateTime:time.Now(),EndDateTime:time.Now(),Status:0,EncounterType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEncounterRequestResult := dao.CreateEncounter( EncounterObj )
	
	if createEncounterRequestResult.Success == false {
		t.Errorf(createEncounterRequestResult.Msg)
	} else {
		fmt.Println("Check Create Encounter success...")
	}
	
	createEncounterObj,_ := createEncounterRequestResult.Data. (model.Encounter)

	// --------------------------------------------------------------
	// Check Encounter Obj ID
	// --------------------------------------------------------------	
	if createEncounterObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Encounter" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEncounterRequestResult := dao.GetEncounter( uint64(createEncounterObj.ID) )
	
	if getEncounterRequestResult.Success == false {
		t.Errorf(getEncounterRequestResult.Msg)
	} else {
		fmt.Println("Check Get Encounter success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEncounterObj,_ := getEncounterRequestResult.Data. (model.Encounter)
	compareEncounter := cmp.Equal(createEncounterObj.ID, getEncounterObj.ID)
	
	if  compareEncounter == false	{
		t.Errorf( "Created Encounter object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEncounterRequestResult := dao.GetAllEncounter()

	if getAllEncounterRequestResult.Success == false {
			t.Errorf(getAllEncounterRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Encounter success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEncounterObj []model.Encounter = getAllEncounterRequestResult.Data. ([]model.Encounter)
		
	equalEncounter := cmp.Equal(createEncounterObj.ID, getAllEncounterObj[len(getAllEncounterObj)-1].ID)
		
	if equalEncounter == false {
		t.Errorf( "Created object is not equal to the last entry in Encounter[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Encounter
	// --------------------------------------------------------------	
	deleteEncounterRequestResult := dao.DeleteEncounter(uint64(createEncounterObj.ID))

	if deleteEncounterRequestResult.Success == false {
			t.Errorf(deleteEncounterRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Encounter success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEncounterRequestResult = dao.GetEncounter( uint64(createEncounterObj.ID) )
	
	if getEncounterRequestResult.Success == true {
		t.Errorf(getEncounterRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAdmissionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Admission
	//----------------------------------------------------------------------------
	AdmissionObj := model.Admission                                                                                                    {AdmitDateTime:time.Now(),Bed:"test value for Bed",AdmissionType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAdmissionRequestResult := dao.CreateAdmission( AdmissionObj )
	
	if createAdmissionRequestResult.Success == false {
		t.Errorf(createAdmissionRequestResult.Msg)
	} else {
		fmt.Println("Check Create Admission success...")
	}
	
	createAdmissionObj,_ := createAdmissionRequestResult.Data. (model.Admission)

	// --------------------------------------------------------------
	// Check Admission Obj ID
	// --------------------------------------------------------------	
	if createAdmissionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Admission" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAdmissionRequestResult := dao.GetAdmission( uint64(createAdmissionObj.ID) )
	
	if getAdmissionRequestResult.Success == false {
		t.Errorf(getAdmissionRequestResult.Msg)
	} else {
		fmt.Println("Check Get Admission success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAdmissionObj,_ := getAdmissionRequestResult.Data. (model.Admission)
	compareAdmission := cmp.Equal(createAdmissionObj.ID, getAdmissionObj.ID)
	
	if  compareAdmission == false	{
		t.Errorf( "Created Admission object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAdmissionRequestResult := dao.GetAllAdmission()

	if getAllAdmissionRequestResult.Success == false {
			t.Errorf(getAllAdmissionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Admission success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAdmissionObj []model.Admission = getAllAdmissionRequestResult.Data. ([]model.Admission)
		
	equalAdmission := cmp.Equal(createAdmissionObj.ID, getAllAdmissionObj[len(getAllAdmissionObj)-1].ID)
		
	if equalAdmission == false {
		t.Errorf( "Created object is not equal to the last entry in Admission[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Admission
	// --------------------------------------------------------------	
	deleteAdmissionRequestResult := dao.DeleteAdmission(uint64(createAdmissionObj.ID))

	if deleteAdmissionRequestResult.Success == false {
			t.Errorf(deleteAdmissionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Admission success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAdmissionRequestResult = dao.GetAdmission( uint64(createAdmissionObj.ID) )
	
	if getAdmissionRequestResult.Success == true {
		t.Errorf(getAdmissionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDischargeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Discharge
	//----------------------------------------------------------------------------
	DischargeObj := model.Discharge                                                                    {DischargeDateTime:time.Now(),Disposition:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDischargeRequestResult := dao.CreateDischarge( DischargeObj )
	
	if createDischargeRequestResult.Success == false {
		t.Errorf(createDischargeRequestResult.Msg)
	} else {
		fmt.Println("Check Create Discharge success...")
	}
	
	createDischargeObj,_ := createDischargeRequestResult.Data. (model.Discharge)

	// --------------------------------------------------------------
	// Check Discharge Obj ID
	// --------------------------------------------------------------	
	if createDischargeObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Discharge" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDischargeRequestResult := dao.GetDischarge( uint64(createDischargeObj.ID) )
	
	if getDischargeRequestResult.Success == false {
		t.Errorf(getDischargeRequestResult.Msg)
	} else {
		fmt.Println("Check Get Discharge success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDischargeObj,_ := getDischargeRequestResult.Data. (model.Discharge)
	compareDischarge := cmp.Equal(createDischargeObj.ID, getDischargeObj.ID)
	
	if  compareDischarge == false	{
		t.Errorf( "Created Discharge object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDischargeRequestResult := dao.GetAllDischarge()

	if getAllDischargeRequestResult.Success == false {
			t.Errorf(getAllDischargeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Discharge success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDischargeObj []model.Discharge = getAllDischargeRequestResult.Data. ([]model.Discharge)
		
	equalDischarge := cmp.Equal(createDischargeObj.ID, getAllDischargeObj[len(getAllDischargeObj)-1].ID)
		
	if equalDischarge == false {
		t.Errorf( "Created object is not equal to the last entry in Discharge[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Discharge
	// --------------------------------------------------------------	
	deleteDischargeRequestResult := dao.DeleteDischarge(uint64(createDischargeObj.ID))

	if deleteDischargeRequestResult.Success == false {
			t.Errorf(deleteDischargeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Discharge success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDischargeRequestResult = dao.GetDischarge( uint64(createDischargeObj.ID) )
	
	if getDischargeRequestResult.Success == true {
		t.Errorf(getDischargeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestClinicalOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ClinicalOrder
	//----------------------------------------------------------------------------
	ClinicalOrderObj := model.ClinicalOrder                                                                            {OrderNumber:"test value for OrderNumber",Status:0,OrderType:0,Priority:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createClinicalOrderRequestResult := dao.CreateClinicalOrder( ClinicalOrderObj )
	
	if createClinicalOrderRequestResult.Success == false {
		t.Errorf(createClinicalOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create ClinicalOrder success...")
	}
	
	createClinicalOrderObj,_ := createClinicalOrderRequestResult.Data. (model.ClinicalOrder)

	// --------------------------------------------------------------
	// Check ClinicalOrder Obj ID
	// --------------------------------------------------------------	
	if createClinicalOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ClinicalOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getClinicalOrderRequestResult := dao.GetClinicalOrder( uint64(createClinicalOrderObj.ID) )
	
	if getClinicalOrderRequestResult.Success == false {
		t.Errorf(getClinicalOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get ClinicalOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getClinicalOrderObj,_ := getClinicalOrderRequestResult.Data. (model.ClinicalOrder)
	compareClinicalOrder := cmp.Equal(createClinicalOrderObj.ID, getClinicalOrderObj.ID)
	
	if  compareClinicalOrder == false	{
		t.Errorf( "Created ClinicalOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllClinicalOrderRequestResult := dao.GetAllClinicalOrder()

	if getAllClinicalOrderRequestResult.Success == false {
			t.Errorf(getAllClinicalOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ClinicalOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllClinicalOrderObj []model.ClinicalOrder = getAllClinicalOrderRequestResult.Data. ([]model.ClinicalOrder)
		
	equalClinicalOrder := cmp.Equal(createClinicalOrderObj.ID, getAllClinicalOrderObj[len(getAllClinicalOrderObj)-1].ID)
		
	if equalClinicalOrder == false {
		t.Errorf( "Created object is not equal to the last entry in ClinicalOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ClinicalOrder
	// --------------------------------------------------------------	
	deleteClinicalOrderRequestResult := dao.DeleteClinicalOrder(uint64(createClinicalOrderObj.ID))

	if deleteClinicalOrderRequestResult.Success == false {
			t.Errorf(deleteClinicalOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ClinicalOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getClinicalOrderRequestResult = dao.GetClinicalOrder( uint64(createClinicalOrderObj.ID) )
	
	if getClinicalOrderRequestResult.Success == true {
		t.Errorf(getClinicalOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMedicationOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MedicationOrder
	//----------------------------------------------------------------------------
	MedicationOrderObj := model.MedicationOrder                                                                                                                            {MedicationCode:"test value for MedicationCode",Dose:new Dose(),Frequency:"test value for Frequency",Duration:"test value for Duration",Route:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMedicationOrderRequestResult := dao.CreateMedicationOrder( MedicationOrderObj )
	
	if createMedicationOrderRequestResult.Success == false {
		t.Errorf(createMedicationOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create MedicationOrder success...")
	}
	
	createMedicationOrderObj,_ := createMedicationOrderRequestResult.Data. (model.MedicationOrder)

	// --------------------------------------------------------------
	// Check MedicationOrder Obj ID
	// --------------------------------------------------------------	
	if createMedicationOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MedicationOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMedicationOrderRequestResult := dao.GetMedicationOrder( uint64(createMedicationOrderObj.ID) )
	
	if getMedicationOrderRequestResult.Success == false {
		t.Errorf(getMedicationOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get MedicationOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMedicationOrderObj,_ := getMedicationOrderRequestResult.Data. (model.MedicationOrder)
	compareMedicationOrder := cmp.Equal(createMedicationOrderObj.ID, getMedicationOrderObj.ID)
	
	if  compareMedicationOrder == false	{
		t.Errorf( "Created MedicationOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMedicationOrderRequestResult := dao.GetAllMedicationOrder()

	if getAllMedicationOrderRequestResult.Success == false {
			t.Errorf(getAllMedicationOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MedicationOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMedicationOrderObj []model.MedicationOrder = getAllMedicationOrderRequestResult.Data. ([]model.MedicationOrder)
		
	equalMedicationOrder := cmp.Equal(createMedicationOrderObj.ID, getAllMedicationOrderObj[len(getAllMedicationOrderObj)-1].ID)
		
	if equalMedicationOrder == false {
		t.Errorf( "Created object is not equal to the last entry in MedicationOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MedicationOrder
	// --------------------------------------------------------------	
	deleteMedicationOrderRequestResult := dao.DeleteMedicationOrder(uint64(createMedicationOrderObj.ID))

	if deleteMedicationOrderRequestResult.Success == false {
			t.Errorf(deleteMedicationOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MedicationOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMedicationOrderRequestResult = dao.GetMedicationOrder( uint64(createMedicationOrderObj.ID) )
	
	if getMedicationOrderRequestResult.Success == true {
		t.Errorf(getMedicationOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLaboratoryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Laboratory
	//----------------------------------------------------------------------------
	LaboratoryObj := model.Laboratory                                                            {Name:"test value for Name",CliaNumber:"test value for CliaNumber"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLaboratoryRequestResult := dao.CreateLaboratory( LaboratoryObj )
	
	if createLaboratoryRequestResult.Success == false {
		t.Errorf(createLaboratoryRequestResult.Msg)
	} else {
		fmt.Println("Check Create Laboratory success...")
	}
	
	createLaboratoryObj,_ := createLaboratoryRequestResult.Data. (model.Laboratory)

	// --------------------------------------------------------------
	// Check Laboratory Obj ID
	// --------------------------------------------------------------	
	if createLaboratoryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Laboratory" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLaboratoryRequestResult := dao.GetLaboratory( uint64(createLaboratoryObj.ID) )
	
	if getLaboratoryRequestResult.Success == false {
		t.Errorf(getLaboratoryRequestResult.Msg)
	} else {
		fmt.Println("Check Get Laboratory success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLaboratoryObj,_ := getLaboratoryRequestResult.Data. (model.Laboratory)
	compareLaboratory := cmp.Equal(createLaboratoryObj.ID, getLaboratoryObj.ID)
	
	if  compareLaboratory == false	{
		t.Errorf( "Created Laboratory object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLaboratoryRequestResult := dao.GetAllLaboratory()

	if getAllLaboratoryRequestResult.Success == false {
			t.Errorf(getAllLaboratoryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Laboratory success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLaboratoryObj []model.Laboratory = getAllLaboratoryRequestResult.Data. ([]model.Laboratory)
		
	equalLaboratory := cmp.Equal(createLaboratoryObj.ID, getAllLaboratoryObj[len(getAllLaboratoryObj)-1].ID)
		
	if equalLaboratory == false {
		t.Errorf( "Created object is not equal to the last entry in Laboratory[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Laboratory
	// --------------------------------------------------------------	
	deleteLaboratoryRequestResult := dao.DeleteLaboratory(uint64(createLaboratoryObj.ID))

	if deleteLaboratoryRequestResult.Success == false {
			t.Errorf(deleteLaboratoryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Laboratory success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLaboratoryRequestResult = dao.GetLaboratory( uint64(createLaboratoryObj.ID) )
	
	if getLaboratoryRequestResult.Success == true {
		t.Errorf(getLaboratoryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLaboratoryOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LaboratoryOrder
	//----------------------------------------------------------------------------
	LaboratoryOrderObj := model.LaboratoryOrder                                                                            {TestCode:"test value for TestCode",FastingRequired:true,SpecimenType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLaboratoryOrderRequestResult := dao.CreateLaboratoryOrder( LaboratoryOrderObj )
	
	if createLaboratoryOrderRequestResult.Success == false {
		t.Errorf(createLaboratoryOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create LaboratoryOrder success...")
	}
	
	createLaboratoryOrderObj,_ := createLaboratoryOrderRequestResult.Data. (model.LaboratoryOrder)

	// --------------------------------------------------------------
	// Check LaboratoryOrder Obj ID
	// --------------------------------------------------------------	
	if createLaboratoryOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LaboratoryOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLaboratoryOrderRequestResult := dao.GetLaboratoryOrder( uint64(createLaboratoryOrderObj.ID) )
	
	if getLaboratoryOrderRequestResult.Success == false {
		t.Errorf(getLaboratoryOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get LaboratoryOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLaboratoryOrderObj,_ := getLaboratoryOrderRequestResult.Data. (model.LaboratoryOrder)
	compareLaboratoryOrder := cmp.Equal(createLaboratoryOrderObj.ID, getLaboratoryOrderObj.ID)
	
	if  compareLaboratoryOrder == false	{
		t.Errorf( "Created LaboratoryOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLaboratoryOrderRequestResult := dao.GetAllLaboratoryOrder()

	if getAllLaboratoryOrderRequestResult.Success == false {
			t.Errorf(getAllLaboratoryOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LaboratoryOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLaboratoryOrderObj []model.LaboratoryOrder = getAllLaboratoryOrderRequestResult.Data. ([]model.LaboratoryOrder)
		
	equalLaboratoryOrder := cmp.Equal(createLaboratoryOrderObj.ID, getAllLaboratoryOrderObj[len(getAllLaboratoryOrderObj)-1].ID)
		
	if equalLaboratoryOrder == false {
		t.Errorf( "Created object is not equal to the last entry in LaboratoryOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LaboratoryOrder
	// --------------------------------------------------------------	
	deleteLaboratoryOrderRequestResult := dao.DeleteLaboratoryOrder(uint64(createLaboratoryOrderObj.ID))

	if deleteLaboratoryOrderRequestResult.Success == false {
			t.Errorf(deleteLaboratoryOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LaboratoryOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLaboratoryOrderRequestResult = dao.GetLaboratoryOrder( uint64(createLaboratoryOrderObj.ID) )
	
	if getLaboratoryOrderRequestResult.Success == true {
		t.Errorf(getLaboratoryOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLabResultCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LabResult
	//----------------------------------------------------------------------------
	LabResultObj := model.LabResult                                                                                                    {ResultCode:"test value for ResultCode",IssuedDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLabResultRequestResult := dao.CreateLabResult( LabResultObj )
	
	if createLabResultRequestResult.Success == false {
		t.Errorf(createLabResultRequestResult.Msg)
	} else {
		fmt.Println("Check Create LabResult success...")
	}
	
	createLabResultObj,_ := createLabResultRequestResult.Data. (model.LabResult)

	// --------------------------------------------------------------
	// Check LabResult Obj ID
	// --------------------------------------------------------------	
	if createLabResultObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LabResult" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLabResultRequestResult := dao.GetLabResult( uint64(createLabResultObj.ID) )
	
	if getLabResultRequestResult.Success == false {
		t.Errorf(getLabResultRequestResult.Msg)
	} else {
		fmt.Println("Check Get LabResult success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLabResultObj,_ := getLabResultRequestResult.Data. (model.LabResult)
	compareLabResult := cmp.Equal(createLabResultObj.ID, getLabResultObj.ID)
	
	if  compareLabResult == false	{
		t.Errorf( "Created LabResult object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLabResultRequestResult := dao.GetAllLabResult()

	if getAllLabResultRequestResult.Success == false {
			t.Errorf(getAllLabResultRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LabResult success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLabResultObj []model.LabResult = getAllLabResultRequestResult.Data. ([]model.LabResult)
		
	equalLabResult := cmp.Equal(createLabResultObj.ID, getAllLabResultObj[len(getAllLabResultObj)-1].ID)
		
	if equalLabResult == false {
		t.Errorf( "Created object is not equal to the last entry in LabResult[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LabResult
	// --------------------------------------------------------------	
	deleteLabResultRequestResult := dao.DeleteLabResult(uint64(createLabResultObj.ID))

	if deleteLabResultRequestResult.Success == false {
			t.Errorf(deleteLabResultRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LabResult success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLabResultRequestResult = dao.GetLabResult( uint64(createLabResultObj.ID) )
	
	if getLabResultRequestResult.Success == true {
		t.Errorf(getLabResultRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestImagingCenterCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ImagingCenter
	//----------------------------------------------------------------------------
	ImagingCenterObj := model.ImagingCenter                            {Name:"test value for Name"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createImagingCenterRequestResult := dao.CreateImagingCenter( ImagingCenterObj )
	
	if createImagingCenterRequestResult.Success == false {
		t.Errorf(createImagingCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Create ImagingCenter success...")
	}
	
	createImagingCenterObj,_ := createImagingCenterRequestResult.Data. (model.ImagingCenter)

	// --------------------------------------------------------------
	// Check ImagingCenter Obj ID
	// --------------------------------------------------------------	
	if createImagingCenterObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ImagingCenter" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getImagingCenterRequestResult := dao.GetImagingCenter( uint64(createImagingCenterObj.ID) )
	
	if getImagingCenterRequestResult.Success == false {
		t.Errorf(getImagingCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Get ImagingCenter success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getImagingCenterObj,_ := getImagingCenterRequestResult.Data. (model.ImagingCenter)
	compareImagingCenter := cmp.Equal(createImagingCenterObj.ID, getImagingCenterObj.ID)
	
	if  compareImagingCenter == false	{
		t.Errorf( "Created ImagingCenter object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllImagingCenterRequestResult := dao.GetAllImagingCenter()

	if getAllImagingCenterRequestResult.Success == false {
			t.Errorf(getAllImagingCenterRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ImagingCenter success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllImagingCenterObj []model.ImagingCenter = getAllImagingCenterRequestResult.Data. ([]model.ImagingCenter)
		
	equalImagingCenter := cmp.Equal(createImagingCenterObj.ID, getAllImagingCenterObj[len(getAllImagingCenterObj)-1].ID)
		
	if equalImagingCenter == false {
		t.Errorf( "Created object is not equal to the last entry in ImagingCenter[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ImagingCenter
	// --------------------------------------------------------------	
	deleteImagingCenterRequestResult := dao.DeleteImagingCenter(uint64(createImagingCenterObj.ID))

	if deleteImagingCenterRequestResult.Success == false {
			t.Errorf(deleteImagingCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ImagingCenter success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getImagingCenterRequestResult = dao.GetImagingCenter( uint64(createImagingCenterObj.ID) )
	
	if getImagingCenterRequestResult.Success == true {
		t.Errorf(getImagingCenterRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestImagingOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ImagingOrder
	//----------------------------------------------------------------------------
	ImagingOrderObj := model.ImagingOrder                                                                            {BodySite:"test value for BodySite",Contrast:true,Modality:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createImagingOrderRequestResult := dao.CreateImagingOrder( ImagingOrderObj )
	
	if createImagingOrderRequestResult.Success == false {
		t.Errorf(createImagingOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create ImagingOrder success...")
	}
	
	createImagingOrderObj,_ := createImagingOrderRequestResult.Data. (model.ImagingOrder)

	// --------------------------------------------------------------
	// Check ImagingOrder Obj ID
	// --------------------------------------------------------------	
	if createImagingOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ImagingOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getImagingOrderRequestResult := dao.GetImagingOrder( uint64(createImagingOrderObj.ID) )
	
	if getImagingOrderRequestResult.Success == false {
		t.Errorf(getImagingOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get ImagingOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getImagingOrderObj,_ := getImagingOrderRequestResult.Data. (model.ImagingOrder)
	compareImagingOrder := cmp.Equal(createImagingOrderObj.ID, getImagingOrderObj.ID)
	
	if  compareImagingOrder == false	{
		t.Errorf( "Created ImagingOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllImagingOrderRequestResult := dao.GetAllImagingOrder()

	if getAllImagingOrderRequestResult.Success == false {
			t.Errorf(getAllImagingOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ImagingOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllImagingOrderObj []model.ImagingOrder = getAllImagingOrderRequestResult.Data. ([]model.ImagingOrder)
		
	equalImagingOrder := cmp.Equal(createImagingOrderObj.ID, getAllImagingOrderObj[len(getAllImagingOrderObj)-1].ID)
		
	if equalImagingOrder == false {
		t.Errorf( "Created object is not equal to the last entry in ImagingOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ImagingOrder
	// --------------------------------------------------------------	
	deleteImagingOrderRequestResult := dao.DeleteImagingOrder(uint64(createImagingOrderObj.ID))

	if deleteImagingOrderRequestResult.Success == false {
			t.Errorf(deleteImagingOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ImagingOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getImagingOrderRequestResult = dao.GetImagingOrder( uint64(createImagingOrderObj.ID) )
	
	if getImagingOrderRequestResult.Success == true {
		t.Errorf(getImagingOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestImagingReportCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ImagingReport
	//----------------------------------------------------------------------------
	ImagingReportObj := model.ImagingReport                                                                                                                                    {ReportNumber:"test value for ReportNumber",Impression:"test value for Impression",ReportedDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createImagingReportRequestResult := dao.CreateImagingReport( ImagingReportObj )
	
	if createImagingReportRequestResult.Success == false {
		t.Errorf(createImagingReportRequestResult.Msg)
	} else {
		fmt.Println("Check Create ImagingReport success...")
	}
	
	createImagingReportObj,_ := createImagingReportRequestResult.Data. (model.ImagingReport)

	// --------------------------------------------------------------
	// Check ImagingReport Obj ID
	// --------------------------------------------------------------	
	if createImagingReportObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ImagingReport" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getImagingReportRequestResult := dao.GetImagingReport( uint64(createImagingReportObj.ID) )
	
	if getImagingReportRequestResult.Success == false {
		t.Errorf(getImagingReportRequestResult.Msg)
	} else {
		fmt.Println("Check Get ImagingReport success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getImagingReportObj,_ := getImagingReportRequestResult.Data. (model.ImagingReport)
	compareImagingReport := cmp.Equal(createImagingReportObj.ID, getImagingReportObj.ID)
	
	if  compareImagingReport == false	{
		t.Errorf( "Created ImagingReport object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllImagingReportRequestResult := dao.GetAllImagingReport()

	if getAllImagingReportRequestResult.Success == false {
			t.Errorf(getAllImagingReportRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ImagingReport success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllImagingReportObj []model.ImagingReport = getAllImagingReportRequestResult.Data. ([]model.ImagingReport)
		
	equalImagingReport := cmp.Equal(createImagingReportObj.ID, getAllImagingReportObj[len(getAllImagingReportObj)-1].ID)
		
	if equalImagingReport == false {
		t.Errorf( "Created object is not equal to the last entry in ImagingReport[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ImagingReport
	// --------------------------------------------------------------	
	deleteImagingReportRequestResult := dao.DeleteImagingReport(uint64(createImagingReportObj.ID))

	if deleteImagingReportRequestResult.Success == false {
			t.Errorf(deleteImagingReportRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ImagingReport success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getImagingReportRequestResult = dao.GetImagingReport( uint64(createImagingReportObj.ID) )
	
	if getImagingReportRequestResult.Success == true {
		t.Errorf(getImagingReportRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestProcedureOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ProcedureOrder
	//----------------------------------------------------------------------------
	ProcedureOrderObj := model.ProcedureOrder                                                                            {ProcedureCode:"test value for ProcedureCode",ConsentObtained:true,AnesthesiaType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createProcedureOrderRequestResult := dao.CreateProcedureOrder( ProcedureOrderObj )
	
	if createProcedureOrderRequestResult.Success == false {
		t.Errorf(createProcedureOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create ProcedureOrder success...")
	}
	
	createProcedureOrderObj,_ := createProcedureOrderRequestResult.Data. (model.ProcedureOrder)

	// --------------------------------------------------------------
	// Check ProcedureOrder Obj ID
	// --------------------------------------------------------------	
	if createProcedureOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ProcedureOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getProcedureOrderRequestResult := dao.GetProcedureOrder( uint64(createProcedureOrderObj.ID) )
	
	if getProcedureOrderRequestResult.Success == false {
		t.Errorf(getProcedureOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get ProcedureOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getProcedureOrderObj,_ := getProcedureOrderRequestResult.Data. (model.ProcedureOrder)
	compareProcedureOrder := cmp.Equal(createProcedureOrderObj.ID, getProcedureOrderObj.ID)
	
	if  compareProcedureOrder == false	{
		t.Errorf( "Created ProcedureOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllProcedureOrderRequestResult := dao.GetAllProcedureOrder()

	if getAllProcedureOrderRequestResult.Success == false {
			t.Errorf(getAllProcedureOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ProcedureOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllProcedureOrderObj []model.ProcedureOrder = getAllProcedureOrderRequestResult.Data. ([]model.ProcedureOrder)
		
	equalProcedureOrder := cmp.Equal(createProcedureOrderObj.ID, getAllProcedureOrderObj[len(getAllProcedureOrderObj)-1].ID)
		
	if equalProcedureOrder == false {
		t.Errorf( "Created object is not equal to the last entry in ProcedureOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ProcedureOrder
	// --------------------------------------------------------------	
	deleteProcedureOrderRequestResult := dao.DeleteProcedureOrder(uint64(createProcedureOrderObj.ID))

	if deleteProcedureOrderRequestResult.Success == false {
			t.Errorf(deleteProcedureOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ProcedureOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getProcedureOrderRequestResult = dao.GetProcedureOrder( uint64(createProcedureOrderObj.ID) )
	
	if getProcedureOrderRequestResult.Success == true {
		t.Errorf(getProcedureOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestProcedureCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Procedure
	//----------------------------------------------------------------------------
	ProcedureObj := model.Procedure                                                                                                                                                            {ProcedureCode:"test value for ProcedureCode",StartDateTime:time.Now(),EndDateTime:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createProcedureRequestResult := dao.CreateProcedure( ProcedureObj )
	
	if createProcedureRequestResult.Success == false {
		t.Errorf(createProcedureRequestResult.Msg)
	} else {
		fmt.Println("Check Create Procedure success...")
	}
	
	createProcedureObj,_ := createProcedureRequestResult.Data. (model.Procedure)

	// --------------------------------------------------------------
	// Check Procedure Obj ID
	// --------------------------------------------------------------	
	if createProcedureObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Procedure" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getProcedureRequestResult := dao.GetProcedure( uint64(createProcedureObj.ID) )
	
	if getProcedureRequestResult.Success == false {
		t.Errorf(getProcedureRequestResult.Msg)
	} else {
		fmt.Println("Check Get Procedure success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getProcedureObj,_ := getProcedureRequestResult.Data. (model.Procedure)
	compareProcedure := cmp.Equal(createProcedureObj.ID, getProcedureObj.ID)
	
	if  compareProcedure == false	{
		t.Errorf( "Created Procedure object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllProcedureRequestResult := dao.GetAllProcedure()

	if getAllProcedureRequestResult.Success == false {
			t.Errorf(getAllProcedureRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Procedure success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllProcedureObj []model.Procedure = getAllProcedureRequestResult.Data. ([]model.Procedure)
		
	equalProcedure := cmp.Equal(createProcedureObj.ID, getAllProcedureObj[len(getAllProcedureObj)-1].ID)
		
	if equalProcedure == false {
		t.Errorf( "Created object is not equal to the last entry in Procedure[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Procedure
	// --------------------------------------------------------------	
	deleteProcedureRequestResult := dao.DeleteProcedure(uint64(createProcedureObj.ID))

	if deleteProcedureRequestResult.Success == false {
			t.Errorf(deleteProcedureRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Procedure success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getProcedureRequestResult = dao.GetProcedure( uint64(createProcedureObj.ID) )
	
	if getProcedureRequestResult.Success == true {
		t.Errorf(getProcedureRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPharmacyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Pharmacy
	//----------------------------------------------------------------------------
	PharmacyObj := model.Pharmacy                            {Name:"test value for Name"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPharmacyRequestResult := dao.CreatePharmacy( PharmacyObj )
	
	if createPharmacyRequestResult.Success == false {
		t.Errorf(createPharmacyRequestResult.Msg)
	} else {
		fmt.Println("Check Create Pharmacy success...")
	}
	
	createPharmacyObj,_ := createPharmacyRequestResult.Data. (model.Pharmacy)

	// --------------------------------------------------------------
	// Check Pharmacy Obj ID
	// --------------------------------------------------------------	
	if createPharmacyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Pharmacy" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPharmacyRequestResult := dao.GetPharmacy( uint64(createPharmacyObj.ID) )
	
	if getPharmacyRequestResult.Success == false {
		t.Errorf(getPharmacyRequestResult.Msg)
	} else {
		fmt.Println("Check Get Pharmacy success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPharmacyObj,_ := getPharmacyRequestResult.Data. (model.Pharmacy)
	comparePharmacy := cmp.Equal(createPharmacyObj.ID, getPharmacyObj.ID)
	
	if  comparePharmacy == false	{
		t.Errorf( "Created Pharmacy object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPharmacyRequestResult := dao.GetAllPharmacy()

	if getAllPharmacyRequestResult.Success == false {
			t.Errorf(getAllPharmacyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Pharmacy success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPharmacyObj []model.Pharmacy = getAllPharmacyRequestResult.Data. ([]model.Pharmacy)
		
	equalPharmacy := cmp.Equal(createPharmacyObj.ID, getAllPharmacyObj[len(getAllPharmacyObj)-1].ID)
		
	if equalPharmacy == false {
		t.Errorf( "Created object is not equal to the last entry in Pharmacy[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Pharmacy
	// --------------------------------------------------------------	
	deletePharmacyRequestResult := dao.DeletePharmacy(uint64(createPharmacyObj.ID))

	if deletePharmacyRequestResult.Success == false {
			t.Errorf(deletePharmacyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Pharmacy success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPharmacyRequestResult = dao.GetPharmacy( uint64(createPharmacyObj.ID) )
	
	if getPharmacyRequestResult.Success == true {
		t.Errorf(getPharmacyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMedicationDispenseCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MedicationDispense
	//----------------------------------------------------------------------------
	MedicationDispenseObj := model.MedicationDispense                                                                                                                                                            {DispenseNumber:"test value for DispenseNumber",Quantity:"test value",WhenPrepared:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMedicationDispenseRequestResult := dao.CreateMedicationDispense( MedicationDispenseObj )
	
	if createMedicationDispenseRequestResult.Success == false {
		t.Errorf(createMedicationDispenseRequestResult.Msg)
	} else {
		fmt.Println("Check Create MedicationDispense success...")
	}
	
	createMedicationDispenseObj,_ := createMedicationDispenseRequestResult.Data. (model.MedicationDispense)

	// --------------------------------------------------------------
	// Check MedicationDispense Obj ID
	// --------------------------------------------------------------	
	if createMedicationDispenseObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MedicationDispense" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMedicationDispenseRequestResult := dao.GetMedicationDispense( uint64(createMedicationDispenseObj.ID) )
	
	if getMedicationDispenseRequestResult.Success == false {
		t.Errorf(getMedicationDispenseRequestResult.Msg)
	} else {
		fmt.Println("Check Get MedicationDispense success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMedicationDispenseObj,_ := getMedicationDispenseRequestResult.Data. (model.MedicationDispense)
	compareMedicationDispense := cmp.Equal(createMedicationDispenseObj.ID, getMedicationDispenseObj.ID)
	
	if  compareMedicationDispense == false	{
		t.Errorf( "Created MedicationDispense object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMedicationDispenseRequestResult := dao.GetAllMedicationDispense()

	if getAllMedicationDispenseRequestResult.Success == false {
			t.Errorf(getAllMedicationDispenseRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MedicationDispense success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMedicationDispenseObj []model.MedicationDispense = getAllMedicationDispenseRequestResult.Data. ([]model.MedicationDispense)
		
	equalMedicationDispense := cmp.Equal(createMedicationDispenseObj.ID, getAllMedicationDispenseObj[len(getAllMedicationDispenseObj)-1].ID)
		
	if equalMedicationDispense == false {
		t.Errorf( "Created object is not equal to the last entry in MedicationDispense[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MedicationDispense
	// --------------------------------------------------------------	
	deleteMedicationDispenseRequestResult := dao.DeleteMedicationDispense(uint64(createMedicationDispenseObj.ID))

	if deleteMedicationDispenseRequestResult.Success == false {
			t.Errorf(deleteMedicationDispenseRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MedicationDispense success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMedicationDispenseRequestResult = dao.GetMedicationDispense( uint64(createMedicationDispenseObj.ID) )
	
	if getMedicationDispenseRequestResult.Success == true {
		t.Errorf(getMedicationDispenseRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDiagnosisCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Diagnosis
	//----------------------------------------------------------------------------
	DiagnosisObj := model.Diagnosis                                                                                                                                    {Code:"test value for Code",Description:"test value for Description",OnsetDate:time.Now(),Certainty:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDiagnosisRequestResult := dao.CreateDiagnosis( DiagnosisObj )
	
	if createDiagnosisRequestResult.Success == false {
		t.Errorf(createDiagnosisRequestResult.Msg)
	} else {
		fmt.Println("Check Create Diagnosis success...")
	}
	
	createDiagnosisObj,_ := createDiagnosisRequestResult.Data. (model.Diagnosis)

	// --------------------------------------------------------------
	// Check Diagnosis Obj ID
	// --------------------------------------------------------------	
	if createDiagnosisObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Diagnosis" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDiagnosisRequestResult := dao.GetDiagnosis( uint64(createDiagnosisObj.ID) )
	
	if getDiagnosisRequestResult.Success == false {
		t.Errorf(getDiagnosisRequestResult.Msg)
	} else {
		fmt.Println("Check Get Diagnosis success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDiagnosisObj,_ := getDiagnosisRequestResult.Data. (model.Diagnosis)
	compareDiagnosis := cmp.Equal(createDiagnosisObj.ID, getDiagnosisObj.ID)
	
	if  compareDiagnosis == false	{
		t.Errorf( "Created Diagnosis object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDiagnosisRequestResult := dao.GetAllDiagnosis()

	if getAllDiagnosisRequestResult.Success == false {
			t.Errorf(getAllDiagnosisRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Diagnosis success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDiagnosisObj []model.Diagnosis = getAllDiagnosisRequestResult.Data. ([]model.Diagnosis)
		
	equalDiagnosis := cmp.Equal(createDiagnosisObj.ID, getAllDiagnosisObj[len(getAllDiagnosisObj)-1].ID)
		
	if equalDiagnosis == false {
		t.Errorf( "Created object is not equal to the last entry in Diagnosis[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Diagnosis
	// --------------------------------------------------------------	
	deleteDiagnosisRequestResult := dao.DeleteDiagnosis(uint64(createDiagnosisObj.ID))

	if deleteDiagnosisRequestResult.Success == false {
			t.Errorf(deleteDiagnosisRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Diagnosis success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDiagnosisRequestResult = dao.GetDiagnosis( uint64(createDiagnosisObj.ID) )
	
	if getDiagnosisRequestResult.Success == true {
		t.Errorf(getDiagnosisRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestObservationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Observation
	//----------------------------------------------------------------------------
	ObservationObj := model.Observation                                                                                                                                                                    {Code:"test value for Code",Value:"test value for Value",Unit:"test value for Unit",EffectiveDateTime:time.Now(),Interpretation:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createObservationRequestResult := dao.CreateObservation( ObservationObj )
	
	if createObservationRequestResult.Success == false {
		t.Errorf(createObservationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Observation success...")
	}
	
	createObservationObj,_ := createObservationRequestResult.Data. (model.Observation)

	// --------------------------------------------------------------
	// Check Observation Obj ID
	// --------------------------------------------------------------	
	if createObservationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Observation" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getObservationRequestResult := dao.GetObservation( uint64(createObservationObj.ID) )
	
	if getObservationRequestResult.Success == false {
		t.Errorf(getObservationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Observation success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getObservationObj,_ := getObservationRequestResult.Data. (model.Observation)
	compareObservation := cmp.Equal(createObservationObj.ID, getObservationObj.ID)
	
	if  compareObservation == false	{
		t.Errorf( "Created Observation object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllObservationRequestResult := dao.GetAllObservation()

	if getAllObservationRequestResult.Success == false {
			t.Errorf(getAllObservationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Observation success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllObservationObj []model.Observation = getAllObservationRequestResult.Data. ([]model.Observation)
		
	equalObservation := cmp.Equal(createObservationObj.ID, getAllObservationObj[len(getAllObservationObj)-1].ID)
		
	if equalObservation == false {
		t.Errorf( "Created object is not equal to the last entry in Observation[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Observation
	// --------------------------------------------------------------	
	deleteObservationRequestResult := dao.DeleteObservation(uint64(createObservationObj.ID))

	if deleteObservationRequestResult.Success == false {
			t.Errorf(deleteObservationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Observation success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getObservationRequestResult = dao.GetObservation( uint64(createObservationObj.ID) )
	
	if getObservationRequestResult.Success == true {
		t.Errorf(getObservationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCarePlanCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CarePlan
	//----------------------------------------------------------------------------
	CarePlanObj := model.CarePlan                                                                            {PlanNumber:"test value for PlanNumber",GoalSummary:"test value for GoalSummary",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCarePlanRequestResult := dao.CreateCarePlan( CarePlanObj )
	
	if createCarePlanRequestResult.Success == false {
		t.Errorf(createCarePlanRequestResult.Msg)
	} else {
		fmt.Println("Check Create CarePlan success...")
	}
	
	createCarePlanObj,_ := createCarePlanRequestResult.Data. (model.CarePlan)

	// --------------------------------------------------------------
	// Check CarePlan Obj ID
	// --------------------------------------------------------------	
	if createCarePlanObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CarePlan" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCarePlanRequestResult := dao.GetCarePlan( uint64(createCarePlanObj.ID) )
	
	if getCarePlanRequestResult.Success == false {
		t.Errorf(getCarePlanRequestResult.Msg)
	} else {
		fmt.Println("Check Get CarePlan success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCarePlanObj,_ := getCarePlanRequestResult.Data. (model.CarePlan)
	compareCarePlan := cmp.Equal(createCarePlanObj.ID, getCarePlanObj.ID)
	
	if  compareCarePlan == false	{
		t.Errorf( "Created CarePlan object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCarePlanRequestResult := dao.GetAllCarePlan()

	if getAllCarePlanRequestResult.Success == false {
			t.Errorf(getAllCarePlanRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CarePlan success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCarePlanObj []model.CarePlan = getAllCarePlanRequestResult.Data. ([]model.CarePlan)
		
	equalCarePlan := cmp.Equal(createCarePlanObj.ID, getAllCarePlanObj[len(getAllCarePlanObj)-1].ID)
		
	if equalCarePlan == false {
		t.Errorf( "Created object is not equal to the last entry in CarePlan[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CarePlan
	// --------------------------------------------------------------	
	deleteCarePlanRequestResult := dao.DeleteCarePlan(uint64(createCarePlanObj.ID))

	if deleteCarePlanRequestResult.Success == false {
			t.Errorf(deleteCarePlanRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CarePlan success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCarePlanRequestResult = dao.GetCarePlan( uint64(createCarePlanObj.ID) )
	
	if getCarePlanRequestResult.Success == true {
		t.Errorf(getCarePlanRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCareTaskCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CareTask
	//----------------------------------------------------------------------------
	CareTaskObj := model.CareTask                                                                                                                    {Description:"test value for Description",DueDate:time.Now(),Status:0,Priority:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCareTaskRequestResult := dao.CreateCareTask( CareTaskObj )
	
	if createCareTaskRequestResult.Success == false {
		t.Errorf(createCareTaskRequestResult.Msg)
	} else {
		fmt.Println("Check Create CareTask success...")
	}
	
	createCareTaskObj,_ := createCareTaskRequestResult.Data. (model.CareTask)

	// --------------------------------------------------------------
	// Check CareTask Obj ID
	// --------------------------------------------------------------	
	if createCareTaskObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CareTask" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCareTaskRequestResult := dao.GetCareTask( uint64(createCareTaskObj.ID) )
	
	if getCareTaskRequestResult.Success == false {
		t.Errorf(getCareTaskRequestResult.Msg)
	} else {
		fmt.Println("Check Get CareTask success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCareTaskObj,_ := getCareTaskRequestResult.Data. (model.CareTask)
	compareCareTask := cmp.Equal(createCareTaskObj.ID, getCareTaskObj.ID)
	
	if  compareCareTask == false	{
		t.Errorf( "Created CareTask object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCareTaskRequestResult := dao.GetAllCareTask()

	if getAllCareTaskRequestResult.Success == false {
			t.Errorf(getAllCareTaskRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CareTask success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCareTaskObj []model.CareTask = getAllCareTaskRequestResult.Data. ([]model.CareTask)
		
	equalCareTask := cmp.Equal(createCareTaskObj.ID, getAllCareTaskObj[len(getAllCareTaskObj)-1].ID)
		
	if equalCareTask == false {
		t.Errorf( "Created object is not equal to the last entry in CareTask[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CareTask
	// --------------------------------------------------------------	
	deleteCareTaskRequestResult := dao.DeleteCareTask(uint64(createCareTaskObj.ID))

	if deleteCareTaskRequestResult.Success == false {
			t.Errorf(deleteCareTaskRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CareTask success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCareTaskRequestResult = dao.GetCareTask( uint64(createCareTaskObj.ID) )
	
	if getCareTaskRequestResult.Success == true {
		t.Errorf(getCareTaskRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAllergyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Allergy
	//----------------------------------------------------------------------------
	AllergyObj := model.Allergy                                                                                            {Substance:"test value for Substance",Reaction:"test value for Reaction",Severity:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAllergyRequestResult := dao.CreateAllergy( AllergyObj )
	
	if createAllergyRequestResult.Success == false {
		t.Errorf(createAllergyRequestResult.Msg)
	} else {
		fmt.Println("Check Create Allergy success...")
	}
	
	createAllergyObj,_ := createAllergyRequestResult.Data. (model.Allergy)

	// --------------------------------------------------------------
	// Check Allergy Obj ID
	// --------------------------------------------------------------	
	if createAllergyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Allergy" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAllergyRequestResult := dao.GetAllergy( uint64(createAllergyObj.ID) )
	
	if getAllergyRequestResult.Success == false {
		t.Errorf(getAllergyRequestResult.Msg)
	} else {
		fmt.Println("Check Get Allergy success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAllergyObj,_ := getAllergyRequestResult.Data. (model.Allergy)
	compareAllergy := cmp.Equal(createAllergyObj.ID, getAllergyObj.ID)
	
	if  compareAllergy == false	{
		t.Errorf( "Created Allergy object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAllergyRequestResult := dao.GetAllAllergy()

	if getAllAllergyRequestResult.Success == false {
			t.Errorf(getAllAllergyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Allergy success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAllergyObj []model.Allergy = getAllAllergyRequestResult.Data. ([]model.Allergy)
		
	equalAllergy := cmp.Equal(createAllergyObj.ID, getAllAllergyObj[len(getAllAllergyObj)-1].ID)
		
	if equalAllergy == false {
		t.Errorf( "Created object is not equal to the last entry in Allergy[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Allergy
	// --------------------------------------------------------------	
	deleteAllergyRequestResult := dao.DeleteAllergy(uint64(createAllergyObj.ID))

	if deleteAllergyRequestResult.Success == false {
			t.Errorf(deleteAllergyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Allergy success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAllergyRequestResult = dao.GetAllergy( uint64(createAllergyObj.ID) )
	
	if getAllergyRequestResult.Success == true {
		t.Errorf(getAllergyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestConditionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Condition
	//----------------------------------------------------------------------------
	ConditionObj := model.Condition                                                                                                                                                                            {Code:"test value for Code",OnsetDate:time.Now(),AbatementDate:time.Now(),ClinicalStatus:0,VerificationStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createConditionRequestResult := dao.CreateCondition( ConditionObj )
	
	if createConditionRequestResult.Success == false {
		t.Errorf(createConditionRequestResult.Msg)
	} else {
		fmt.Println("Check Create Condition success...")
	}
	
	createConditionObj,_ := createConditionRequestResult.Data. (model.Condition)

	// --------------------------------------------------------------
	// Check Condition Obj ID
	// --------------------------------------------------------------	
	if createConditionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Condition" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getConditionRequestResult := dao.GetCondition( uint64(createConditionObj.ID) )
	
	if getConditionRequestResult.Success == false {
		t.Errorf(getConditionRequestResult.Msg)
	} else {
		fmt.Println("Check Get Condition success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getConditionObj,_ := getConditionRequestResult.Data. (model.Condition)
	compareCondition := cmp.Equal(createConditionObj.ID, getConditionObj.ID)
	
	if  compareCondition == false	{
		t.Errorf( "Created Condition object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllConditionRequestResult := dao.GetAllCondition()

	if getAllConditionRequestResult.Success == false {
			t.Errorf(getAllConditionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Condition success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllConditionObj []model.Condition = getAllConditionRequestResult.Data. ([]model.Condition)
		
	equalCondition := cmp.Equal(createConditionObj.ID, getAllConditionObj[len(getAllConditionObj)-1].ID)
		
	if equalCondition == false {
		t.Errorf( "Created object is not equal to the last entry in Condition[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Condition
	// --------------------------------------------------------------	
	deleteConditionRequestResult := dao.DeleteCondition(uint64(createConditionObj.ID))

	if deleteConditionRequestResult.Success == false {
			t.Errorf(deleteConditionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Condition success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getConditionRequestResult = dao.GetCondition( uint64(createConditionObj.ID) )
	
	if getConditionRequestResult.Success == true {
		t.Errorf(getConditionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInsurancePayerCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InsurancePayer
	//----------------------------------------------------------------------------
	InsurancePayerObj := model.InsurancePayer                                                                            {Name:"test value for Name",Website:"test value for Website",PayerType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInsurancePayerRequestResult := dao.CreateInsurancePayer( InsurancePayerObj )
	
	if createInsurancePayerRequestResult.Success == false {
		t.Errorf(createInsurancePayerRequestResult.Msg)
	} else {
		fmt.Println("Check Create InsurancePayer success...")
	}
	
	createInsurancePayerObj,_ := createInsurancePayerRequestResult.Data. (model.InsurancePayer)

	// --------------------------------------------------------------
	// Check InsurancePayer Obj ID
	// --------------------------------------------------------------	
	if createInsurancePayerObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InsurancePayer" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInsurancePayerRequestResult := dao.GetInsurancePayer( uint64(createInsurancePayerObj.ID) )
	
	if getInsurancePayerRequestResult.Success == false {
		t.Errorf(getInsurancePayerRequestResult.Msg)
	} else {
		fmt.Println("Check Get InsurancePayer success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInsurancePayerObj,_ := getInsurancePayerRequestResult.Data. (model.InsurancePayer)
	compareInsurancePayer := cmp.Equal(createInsurancePayerObj.ID, getInsurancePayerObj.ID)
	
	if  compareInsurancePayer == false	{
		t.Errorf( "Created InsurancePayer object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInsurancePayerRequestResult := dao.GetAllInsurancePayer()

	if getAllInsurancePayerRequestResult.Success == false {
			t.Errorf(getAllInsurancePayerRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InsurancePayer success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInsurancePayerObj []model.InsurancePayer = getAllInsurancePayerRequestResult.Data. ([]model.InsurancePayer)
		
	equalInsurancePayer := cmp.Equal(createInsurancePayerObj.ID, getAllInsurancePayerObj[len(getAllInsurancePayerObj)-1].ID)
		
	if equalInsurancePayer == false {
		t.Errorf( "Created object is not equal to the last entry in InsurancePayer[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InsurancePayer
	// --------------------------------------------------------------	
	deleteInsurancePayerRequestResult := dao.DeleteInsurancePayer(uint64(createInsurancePayerObj.ID))

	if deleteInsurancePayerRequestResult.Success == false {
			t.Errorf(deleteInsurancePayerRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InsurancePayer success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInsurancePayerRequestResult = dao.GetInsurancePayer( uint64(createInsurancePayerObj.ID) )
	
	if getInsurancePayerRequestResult.Success == true {
		t.Errorf(getInsurancePayerRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInsurancePlanCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InsurancePlan
	//----------------------------------------------------------------------------
	InsurancePlanObj := model.InsurancePlan                                                                            {Name:"test value for Name",PlanCode:"test value for PlanCode",PlanType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInsurancePlanRequestResult := dao.CreateInsurancePlan( InsurancePlanObj )
	
	if createInsurancePlanRequestResult.Success == false {
		t.Errorf(createInsurancePlanRequestResult.Msg)
	} else {
		fmt.Println("Check Create InsurancePlan success...")
	}
	
	createInsurancePlanObj,_ := createInsurancePlanRequestResult.Data. (model.InsurancePlan)

	// --------------------------------------------------------------
	// Check InsurancePlan Obj ID
	// --------------------------------------------------------------	
	if createInsurancePlanObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InsurancePlan" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInsurancePlanRequestResult := dao.GetInsurancePlan( uint64(createInsurancePlanObj.ID) )
	
	if getInsurancePlanRequestResult.Success == false {
		t.Errorf(getInsurancePlanRequestResult.Msg)
	} else {
		fmt.Println("Check Get InsurancePlan success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInsurancePlanObj,_ := getInsurancePlanRequestResult.Data. (model.InsurancePlan)
	compareInsurancePlan := cmp.Equal(createInsurancePlanObj.ID, getInsurancePlanObj.ID)
	
	if  compareInsurancePlan == false	{
		t.Errorf( "Created InsurancePlan object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInsurancePlanRequestResult := dao.GetAllInsurancePlan()

	if getAllInsurancePlanRequestResult.Success == false {
			t.Errorf(getAllInsurancePlanRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InsurancePlan success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInsurancePlanObj []model.InsurancePlan = getAllInsurancePlanRequestResult.Data. ([]model.InsurancePlan)
		
	equalInsurancePlan := cmp.Equal(createInsurancePlanObj.ID, getAllInsurancePlanObj[len(getAllInsurancePlanObj)-1].ID)
		
	if equalInsurancePlan == false {
		t.Errorf( "Created object is not equal to the last entry in InsurancePlan[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InsurancePlan
	// --------------------------------------------------------------	
	deleteInsurancePlanRequestResult := dao.DeleteInsurancePlan(uint64(createInsurancePlanObj.ID))

	if deleteInsurancePlanRequestResult.Success == false {
			t.Errorf(deleteInsurancePlanRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InsurancePlan success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInsurancePlanRequestResult = dao.GetInsurancePlan( uint64(createInsurancePlanObj.ID) )
	
	if getInsurancePlanRequestResult.Success == true {
		t.Errorf(getInsurancePlanRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCoverageCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Coverage
	//----------------------------------------------------------------------------
	CoverageObj := model.Coverage                                                                                                                                                                                            {MemberId:"test value for MemberId",GroupNumber:"test value for GroupNumber",EffectiveDate:time.Now(),EndDate:time.Now(),CoverageType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCoverageRequestResult := dao.CreateCoverage( CoverageObj )
	
	if createCoverageRequestResult.Success == false {
		t.Errorf(createCoverageRequestResult.Msg)
	} else {
		fmt.Println("Check Create Coverage success...")
	}
	
	createCoverageObj,_ := createCoverageRequestResult.Data. (model.Coverage)

	// --------------------------------------------------------------
	// Check Coverage Obj ID
	// --------------------------------------------------------------	
	if createCoverageObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Coverage" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCoverageRequestResult := dao.GetCoverage( uint64(createCoverageObj.ID) )
	
	if getCoverageRequestResult.Success == false {
		t.Errorf(getCoverageRequestResult.Msg)
	} else {
		fmt.Println("Check Get Coverage success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCoverageObj,_ := getCoverageRequestResult.Data. (model.Coverage)
	compareCoverage := cmp.Equal(createCoverageObj.ID, getCoverageObj.ID)
	
	if  compareCoverage == false	{
		t.Errorf( "Created Coverage object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCoverageRequestResult := dao.GetAllCoverage()

	if getAllCoverageRequestResult.Success == false {
			t.Errorf(getAllCoverageRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Coverage success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCoverageObj []model.Coverage = getAllCoverageRequestResult.Data. ([]model.Coverage)
		
	equalCoverage := cmp.Equal(createCoverageObj.ID, getAllCoverageObj[len(getAllCoverageObj)-1].ID)
		
	if equalCoverage == false {
		t.Errorf( "Created object is not equal to the last entry in Coverage[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Coverage
	// --------------------------------------------------------------	
	deleteCoverageRequestResult := dao.DeleteCoverage(uint64(createCoverageObj.ID))

	if deleteCoverageRequestResult.Success == false {
			t.Errorf(deleteCoverageRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Coverage success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCoverageRequestResult = dao.GetCoverage( uint64(createCoverageObj.ID) )
	
	if getCoverageRequestResult.Success == true {
		t.Errorf(getCoverageRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestClaimCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Claim
	//----------------------------------------------------------------------------
	ClaimObj := model.Claim                                                            {ClaimNumber:"test value for ClaimNumber",TotalAmount:new Money(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createClaimRequestResult := dao.CreateClaim( ClaimObj )
	
	if createClaimRequestResult.Success == false {
		t.Errorf(createClaimRequestResult.Msg)
	} else {
		fmt.Println("Check Create Claim success...")
	}
	
	createClaimObj,_ := createClaimRequestResult.Data. (model.Claim)

	// --------------------------------------------------------------
	// Check Claim Obj ID
	// --------------------------------------------------------------	
	if createClaimObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Claim" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getClaimRequestResult := dao.GetClaim( uint64(createClaimObj.ID) )
	
	if getClaimRequestResult.Success == false {
		t.Errorf(getClaimRequestResult.Msg)
	} else {
		fmt.Println("Check Get Claim success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getClaimObj,_ := getClaimRequestResult.Data. (model.Claim)
	compareClaim := cmp.Equal(createClaimObj.ID, getClaimObj.ID)
	
	if  compareClaim == false	{
		t.Errorf( "Created Claim object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllClaimRequestResult := dao.GetAllClaim()

	if getAllClaimRequestResult.Success == false {
			t.Errorf(getAllClaimRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Claim success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllClaimObj []model.Claim = getAllClaimRequestResult.Data. ([]model.Claim)
		
	equalClaim := cmp.Equal(createClaimObj.ID, getAllClaimObj[len(getAllClaimObj)-1].ID)
		
	if equalClaim == false {
		t.Errorf( "Created object is not equal to the last entry in Claim[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Claim
	// --------------------------------------------------------------	
	deleteClaimRequestResult := dao.DeleteClaim(uint64(createClaimObj.ID))

	if deleteClaimRequestResult.Success == false {
			t.Errorf(deleteClaimRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Claim success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getClaimRequestResult = dao.GetClaim( uint64(createClaimObj.ID) )
	
	if getClaimRequestResult.Success == true {
		t.Errorf(getClaimRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAuthorizationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Authorization
	//----------------------------------------------------------------------------
	AuthorizationObj := model.Authorization                                                                            {AuthNumber:"test value for AuthNumber",RequestedService:"test value for RequestedService",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAuthorizationRequestResult := dao.CreateAuthorization( AuthorizationObj )
	
	if createAuthorizationRequestResult.Success == false {
		t.Errorf(createAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Authorization success...")
	}
	
	createAuthorizationObj,_ := createAuthorizationRequestResult.Data. (model.Authorization)

	// --------------------------------------------------------------
	// Check Authorization Obj ID
	// --------------------------------------------------------------	
	if createAuthorizationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Authorization" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAuthorizationRequestResult := dao.GetAuthorization( uint64(createAuthorizationObj.ID) )
	
	if getAuthorizationRequestResult.Success == false {
		t.Errorf(getAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Authorization success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAuthorizationObj,_ := getAuthorizationRequestResult.Data. (model.Authorization)
	compareAuthorization := cmp.Equal(createAuthorizationObj.ID, getAuthorizationObj.ID)
	
	if  compareAuthorization == false	{
		t.Errorf( "Created Authorization object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAuthorizationRequestResult := dao.GetAllAuthorization()

	if getAllAuthorizationRequestResult.Success == false {
			t.Errorf(getAllAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Authorization success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAuthorizationObj []model.Authorization = getAllAuthorizationRequestResult.Data. ([]model.Authorization)
		
	equalAuthorization := cmp.Equal(createAuthorizationObj.ID, getAllAuthorizationObj[len(getAllAuthorizationObj)-1].ID)
		
	if equalAuthorization == false {
		t.Errorf( "Created object is not equal to the last entry in Authorization[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Authorization
	// --------------------------------------------------------------	
	deleteAuthorizationRequestResult := dao.DeleteAuthorization(uint64(createAuthorizationObj.ID))

	if deleteAuthorizationRequestResult.Success == false {
			t.Errorf(deleteAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Authorization success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAuthorizationRequestResult = dao.GetAuthorization( uint64(createAuthorizationObj.ID) )
	
	if getAuthorizationRequestResult.Success == true {
		t.Errorf(getAuthorizationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInvoiceCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Invoice
	//----------------------------------------------------------------------------
	InvoiceObj := model.Invoice                                                                                                                    {InvoiceNumber:"test value for InvoiceNumber",TotalAmount:new Money(),DueDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInvoiceRequestResult := dao.CreateInvoice( InvoiceObj )
	
	if createInvoiceRequestResult.Success == false {
		t.Errorf(createInvoiceRequestResult.Msg)
	} else {
		fmt.Println("Check Create Invoice success...")
	}
	
	createInvoiceObj,_ := createInvoiceRequestResult.Data. (model.Invoice)

	// --------------------------------------------------------------
	// Check Invoice Obj ID
	// --------------------------------------------------------------	
	if createInvoiceObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Invoice" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInvoiceRequestResult := dao.GetInvoice( uint64(createInvoiceObj.ID) )
	
	if getInvoiceRequestResult.Success == false {
		t.Errorf(getInvoiceRequestResult.Msg)
	} else {
		fmt.Println("Check Get Invoice success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInvoiceObj,_ := getInvoiceRequestResult.Data. (model.Invoice)
	compareInvoice := cmp.Equal(createInvoiceObj.ID, getInvoiceObj.ID)
	
	if  compareInvoice == false	{
		t.Errorf( "Created Invoice object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInvoiceRequestResult := dao.GetAllInvoice()

	if getAllInvoiceRequestResult.Success == false {
			t.Errorf(getAllInvoiceRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Invoice success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInvoiceObj []model.Invoice = getAllInvoiceRequestResult.Data. ([]model.Invoice)
		
	equalInvoice := cmp.Equal(createInvoiceObj.ID, getAllInvoiceObj[len(getAllInvoiceObj)-1].ID)
		
	if equalInvoice == false {
		t.Errorf( "Created object is not equal to the last entry in Invoice[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Invoice
	// --------------------------------------------------------------	
	deleteInvoiceRequestResult := dao.DeleteInvoice(uint64(createInvoiceObj.ID))

	if deleteInvoiceRequestResult.Success == false {
			t.Errorf(deleteInvoiceRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Invoice success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInvoiceRequestResult = dao.GetInvoice( uint64(createInvoiceObj.ID) )
	
	if getInvoiceRequestResult.Success == true {
		t.Errorf(getInvoiceRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPaymentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Payment
	//----------------------------------------------------------------------------
	PaymentObj := model.Payment                                                                                                                    {PaymentNumber:"test value for PaymentNumber",Amount:new Money(),PaymentDate:time.Now(),Method:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPaymentRequestResult := dao.CreatePayment( PaymentObj )
	
	if createPaymentRequestResult.Success == false {
		t.Errorf(createPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Payment success...")
	}
	
	createPaymentObj,_ := createPaymentRequestResult.Data. (model.Payment)

	// --------------------------------------------------------------
	// Check Payment Obj ID
	// --------------------------------------------------------------	
	if createPaymentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Payment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPaymentRequestResult := dao.GetPayment( uint64(createPaymentObj.ID) )
	
	if getPaymentRequestResult.Success == false {
		t.Errorf(getPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Payment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPaymentObj,_ := getPaymentRequestResult.Data. (model.Payment)
	comparePayment := cmp.Equal(createPaymentObj.ID, getPaymentObj.ID)
	
	if  comparePayment == false	{
		t.Errorf( "Created Payment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPaymentRequestResult := dao.GetAllPayment()

	if getAllPaymentRequestResult.Success == false {
			t.Errorf(getAllPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Payment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPaymentObj []model.Payment = getAllPaymentRequestResult.Data. ([]model.Payment)
		
	equalPayment := cmp.Equal(createPaymentObj.ID, getAllPaymentObj[len(getAllPaymentObj)-1].ID)
		
	if equalPayment == false {
		t.Errorf( "Created object is not equal to the last entry in Payment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Payment
	// --------------------------------------------------------------	
	deletePaymentRequestResult := dao.DeletePayment(uint64(createPaymentObj.ID))

	if deletePaymentRequestResult.Success == false {
			t.Errorf(deletePaymentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Payment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPaymentRequestResult = dao.GetPayment( uint64(createPaymentObj.ID) )
	
	if getPaymentRequestResult.Success == true {
		t.Errorf(getPaymentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMedicalDeviceCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MedicalDevice
	//----------------------------------------------------------------------------
	MedicalDeviceObj := model.MedicalDevice                                                                                            {Udi:"test value for Udi",Manufacturer:"test value for Manufacturer",DeviceType:0,ConnectivityStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMedicalDeviceRequestResult := dao.CreateMedicalDevice( MedicalDeviceObj )
	
	if createMedicalDeviceRequestResult.Success == false {
		t.Errorf(createMedicalDeviceRequestResult.Msg)
	} else {
		fmt.Println("Check Create MedicalDevice success...")
	}
	
	createMedicalDeviceObj,_ := createMedicalDeviceRequestResult.Data. (model.MedicalDevice)

	// --------------------------------------------------------------
	// Check MedicalDevice Obj ID
	// --------------------------------------------------------------	
	if createMedicalDeviceObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MedicalDevice" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMedicalDeviceRequestResult := dao.GetMedicalDevice( uint64(createMedicalDeviceObj.ID) )
	
	if getMedicalDeviceRequestResult.Success == false {
		t.Errorf(getMedicalDeviceRequestResult.Msg)
	} else {
		fmt.Println("Check Get MedicalDevice success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMedicalDeviceObj,_ := getMedicalDeviceRequestResult.Data. (model.MedicalDevice)
	compareMedicalDevice := cmp.Equal(createMedicalDeviceObj.ID, getMedicalDeviceObj.ID)
	
	if  compareMedicalDevice == false	{
		t.Errorf( "Created MedicalDevice object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMedicalDeviceRequestResult := dao.GetAllMedicalDevice()

	if getAllMedicalDeviceRequestResult.Success == false {
			t.Errorf(getAllMedicalDeviceRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MedicalDevice success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMedicalDeviceObj []model.MedicalDevice = getAllMedicalDeviceRequestResult.Data. ([]model.MedicalDevice)
		
	equalMedicalDevice := cmp.Equal(createMedicalDeviceObj.ID, getAllMedicalDeviceObj[len(getAllMedicalDeviceObj)-1].ID)
		
	if equalMedicalDevice == false {
		t.Errorf( "Created object is not equal to the last entry in MedicalDevice[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MedicalDevice
	// --------------------------------------------------------------	
	deleteMedicalDeviceRequestResult := dao.DeleteMedicalDevice(uint64(createMedicalDeviceObj.ID))

	if deleteMedicalDeviceRequestResult.Success == false {
			t.Errorf(deleteMedicalDeviceRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MedicalDevice success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMedicalDeviceRequestResult = dao.GetMedicalDevice( uint64(createMedicalDeviceObj.ID) )
	
	if getMedicalDeviceRequestResult.Success == true {
		t.Errorf(getMedicalDeviceRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSoftwareUpdateCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SoftwareUpdate
	//----------------------------------------------------------------------------
	SoftwareUpdateObj := model.SoftwareUpdate                                                                                                    {Version:"test value for Version",AppliedDate:time.Now(),UpdateType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSoftwareUpdateRequestResult := dao.CreateSoftwareUpdate( SoftwareUpdateObj )
	
	if createSoftwareUpdateRequestResult.Success == false {
		t.Errorf(createSoftwareUpdateRequestResult.Msg)
	} else {
		fmt.Println("Check Create SoftwareUpdate success...")
	}
	
	createSoftwareUpdateObj,_ := createSoftwareUpdateRequestResult.Data. (model.SoftwareUpdate)

	// --------------------------------------------------------------
	// Check SoftwareUpdate Obj ID
	// --------------------------------------------------------------	
	if createSoftwareUpdateObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SoftwareUpdate" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSoftwareUpdateRequestResult := dao.GetSoftwareUpdate( uint64(createSoftwareUpdateObj.ID) )
	
	if getSoftwareUpdateRequestResult.Success == false {
		t.Errorf(getSoftwareUpdateRequestResult.Msg)
	} else {
		fmt.Println("Check Get SoftwareUpdate success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSoftwareUpdateObj,_ := getSoftwareUpdateRequestResult.Data. (model.SoftwareUpdate)
	compareSoftwareUpdate := cmp.Equal(createSoftwareUpdateObj.ID, getSoftwareUpdateObj.ID)
	
	if  compareSoftwareUpdate == false	{
		t.Errorf( "Created SoftwareUpdate object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSoftwareUpdateRequestResult := dao.GetAllSoftwareUpdate()

	if getAllSoftwareUpdateRequestResult.Success == false {
			t.Errorf(getAllSoftwareUpdateRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SoftwareUpdate success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSoftwareUpdateObj []model.SoftwareUpdate = getAllSoftwareUpdateRequestResult.Data. ([]model.SoftwareUpdate)
		
	equalSoftwareUpdate := cmp.Equal(createSoftwareUpdateObj.ID, getAllSoftwareUpdateObj[len(getAllSoftwareUpdateObj)-1].ID)
		
	if equalSoftwareUpdate == false {
		t.Errorf( "Created object is not equal to the last entry in SoftwareUpdate[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SoftwareUpdate
	// --------------------------------------------------------------	
	deleteSoftwareUpdateRequestResult := dao.DeleteSoftwareUpdate(uint64(createSoftwareUpdateObj.ID))

	if deleteSoftwareUpdateRequestResult.Success == false {
			t.Errorf(deleteSoftwareUpdateRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SoftwareUpdate success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSoftwareUpdateRequestResult = dao.GetSoftwareUpdate( uint64(createSoftwareUpdateObj.ID) )
	
	if getSoftwareUpdateRequestResult.Success == true {
		t.Errorf(getSoftwareUpdateRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMedicalSupplierCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MedicalSupplier
	//----------------------------------------------------------------------------
	MedicalSupplierObj := model.MedicalSupplier                                                                            {Name:"test value for Name",Website:"test value for Website",SupplierTier:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMedicalSupplierRequestResult := dao.CreateMedicalSupplier( MedicalSupplierObj )
	
	if createMedicalSupplierRequestResult.Success == false {
		t.Errorf(createMedicalSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check Create MedicalSupplier success...")
	}
	
	createMedicalSupplierObj,_ := createMedicalSupplierRequestResult.Data. (model.MedicalSupplier)

	// --------------------------------------------------------------
	// Check MedicalSupplier Obj ID
	// --------------------------------------------------------------	
	if createMedicalSupplierObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MedicalSupplier" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMedicalSupplierRequestResult := dao.GetMedicalSupplier( uint64(createMedicalSupplierObj.ID) )
	
	if getMedicalSupplierRequestResult.Success == false {
		t.Errorf(getMedicalSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check Get MedicalSupplier success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMedicalSupplierObj,_ := getMedicalSupplierRequestResult.Data. (model.MedicalSupplier)
	compareMedicalSupplier := cmp.Equal(createMedicalSupplierObj.ID, getMedicalSupplierObj.ID)
	
	if  compareMedicalSupplier == false	{
		t.Errorf( "Created MedicalSupplier object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMedicalSupplierRequestResult := dao.GetAllMedicalSupplier()

	if getAllMedicalSupplierRequestResult.Success == false {
			t.Errorf(getAllMedicalSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MedicalSupplier success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMedicalSupplierObj []model.MedicalSupplier = getAllMedicalSupplierRequestResult.Data. ([]model.MedicalSupplier)
		
	equalMedicalSupplier := cmp.Equal(createMedicalSupplierObj.ID, getAllMedicalSupplierObj[len(getAllMedicalSupplierObj)-1].ID)
		
	if equalMedicalSupplier == false {
		t.Errorf( "Created object is not equal to the last entry in MedicalSupplier[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MedicalSupplier
	// --------------------------------------------------------------	
	deleteMedicalSupplierRequestResult := dao.DeleteMedicalSupplier(uint64(createMedicalSupplierObj.ID))

	if deleteMedicalSupplierRequestResult.Success == false {
			t.Errorf(deleteMedicalSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MedicalSupplier success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMedicalSupplierRequestResult = dao.GetMedicalSupplier( uint64(createMedicalSupplierObj.ID) )
	
	if getMedicalSupplierRequestResult.Success == true {
		t.Errorf(getMedicalSupplierRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInventoryItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InventoryItem
	//----------------------------------------------------------------------------
	InventoryItemObj := model.InventoryItem                                                                                                                            {Sku:"test value for Sku",Name:"test value for Name",QuantityOnHand:100,QuantityReserved:100}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInventoryItemRequestResult := dao.CreateInventoryItem( InventoryItemObj )
	
	if createInventoryItemRequestResult.Success == false {
		t.Errorf(createInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Check Create InventoryItem success...")
	}
	
	createInventoryItemObj,_ := createInventoryItemRequestResult.Data. (model.InventoryItem)

	// --------------------------------------------------------------
	// Check InventoryItem Obj ID
	// --------------------------------------------------------------	
	if createInventoryItemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InventoryItem" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInventoryItemRequestResult := dao.GetInventoryItem( uint64(createInventoryItemObj.ID) )
	
	if getInventoryItemRequestResult.Success == false {
		t.Errorf(getInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Check Get InventoryItem success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInventoryItemObj,_ := getInventoryItemRequestResult.Data. (model.InventoryItem)
	compareInventoryItem := cmp.Equal(createInventoryItemObj.ID, getInventoryItemObj.ID)
	
	if  compareInventoryItem == false	{
		t.Errorf( "Created InventoryItem object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInventoryItemRequestResult := dao.GetAllInventoryItem()

	if getAllInventoryItemRequestResult.Success == false {
			t.Errorf(getAllInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InventoryItem success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInventoryItemObj []model.InventoryItem = getAllInventoryItemRequestResult.Data. ([]model.InventoryItem)
		
	equalInventoryItem := cmp.Equal(createInventoryItemObj.ID, getAllInventoryItemObj[len(getAllInventoryItemObj)-1].ID)
		
	if equalInventoryItem == false {
		t.Errorf( "Created object is not equal to the last entry in InventoryItem[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InventoryItem
	// --------------------------------------------------------------	
	deleteInventoryItemRequestResult := dao.DeleteInventoryItem(uint64(createInventoryItemObj.ID))

	if deleteInventoryItemRequestResult.Success == false {
			t.Errorf(deleteInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InventoryItem success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInventoryItemRequestResult = dao.GetInventoryItem( uint64(createInventoryItemObj.ID) )
	
	if getInventoryItemRequestResult.Success == true {
		t.Errorf(getInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

