package test

import ( 
	"testing"
    dao "aerospace-on-golang/internal/dao"
	"aerospace-on-golang/internal/model"
	"aerospace-on-golang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestAerospaceManufacturerCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AerospaceManufacturer
	//----------------------------------------------------------------------------
	AerospaceManufacturerObj := model.AerospaceManufacturer                                                                                                                            {Name:"test value for Name",LegalName:"test value for LegalName",HeadquartersCountry:"test value for HeadquartersCountry",Website:"test value for Website"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAerospaceManufacturerRequestResult := dao.CreateAerospaceManufacturer( AerospaceManufacturerObj )
	
	if createAerospaceManufacturerRequestResult.Success == false {
		t.Errorf(createAerospaceManufacturerRequestResult.Msg)
	} else {
		fmt.Println("Check Create AerospaceManufacturer success...")
	}
	
	createAerospaceManufacturerObj,_ := createAerospaceManufacturerRequestResult.Data. (model.AerospaceManufacturer)

	// --------------------------------------------------------------
	// Check AerospaceManufacturer Obj ID
	// --------------------------------------------------------------	
	if createAerospaceManufacturerObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AerospaceManufacturer" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAerospaceManufacturerRequestResult := dao.GetAerospaceManufacturer( uint64(createAerospaceManufacturerObj.ID) )
	
	if getAerospaceManufacturerRequestResult.Success == false {
		t.Errorf(getAerospaceManufacturerRequestResult.Msg)
	} else {
		fmt.Println("Check Get AerospaceManufacturer success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAerospaceManufacturerObj,_ := getAerospaceManufacturerRequestResult.Data. (model.AerospaceManufacturer)
	compareAerospaceManufacturer := cmp.Equal(createAerospaceManufacturerObj.ID, getAerospaceManufacturerObj.ID)
	
	if  compareAerospaceManufacturer == false	{
		t.Errorf( "Created AerospaceManufacturer object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAerospaceManufacturerRequestResult := dao.GetAllAerospaceManufacturer()

	if getAllAerospaceManufacturerRequestResult.Success == false {
			t.Errorf(getAllAerospaceManufacturerRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AerospaceManufacturer success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAerospaceManufacturerObj []model.AerospaceManufacturer = getAllAerospaceManufacturerRequestResult.Data. ([]model.AerospaceManufacturer)
		
	equalAerospaceManufacturer := cmp.Equal(createAerospaceManufacturerObj.ID, getAllAerospaceManufacturerObj[len(getAllAerospaceManufacturerObj)-1].ID)
		
	if equalAerospaceManufacturer == false {
		t.Errorf( "Created object is not equal to the last entry in AerospaceManufacturer[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AerospaceManufacturer
	// --------------------------------------------------------------	
	deleteAerospaceManufacturerRequestResult := dao.DeleteAerospaceManufacturer(uint64(createAerospaceManufacturerObj.ID))

	if deleteAerospaceManufacturerRequestResult.Success == false {
			t.Errorf(deleteAerospaceManufacturerRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AerospaceManufacturer success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAerospaceManufacturerRequestResult = dao.GetAerospaceManufacturer( uint64(createAerospaceManufacturerObj.ID) )
	
	if getAerospaceManufacturerRequestResult.Success == true {
		t.Errorf(getAerospaceManufacturerRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAircraftProgramCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AircraftProgram
	//----------------------------------------------------------------------------
	AircraftProgramObj := model.AircraftProgram                                                                                                            {Name:"test value for Name",ProgramCode:"test value for ProgramCode",EntryIntoServiceYear:100,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAircraftProgramRequestResult := dao.CreateAircraftProgram( AircraftProgramObj )
	
	if createAircraftProgramRequestResult.Success == false {
		t.Errorf(createAircraftProgramRequestResult.Msg)
	} else {
		fmt.Println("Check Create AircraftProgram success...")
	}
	
	createAircraftProgramObj,_ := createAircraftProgramRequestResult.Data. (model.AircraftProgram)

	// --------------------------------------------------------------
	// Check AircraftProgram Obj ID
	// --------------------------------------------------------------	
	if createAircraftProgramObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AircraftProgram" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAircraftProgramRequestResult := dao.GetAircraftProgram( uint64(createAircraftProgramObj.ID) )
	
	if getAircraftProgramRequestResult.Success == false {
		t.Errorf(getAircraftProgramRequestResult.Msg)
	} else {
		fmt.Println("Check Get AircraftProgram success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAircraftProgramObj,_ := getAircraftProgramRequestResult.Data. (model.AircraftProgram)
	compareAircraftProgram := cmp.Equal(createAircraftProgramObj.ID, getAircraftProgramObj.ID)
	
	if  compareAircraftProgram == false	{
		t.Errorf( "Created AircraftProgram object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAircraftProgramRequestResult := dao.GetAllAircraftProgram()

	if getAllAircraftProgramRequestResult.Success == false {
			t.Errorf(getAllAircraftProgramRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AircraftProgram success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAircraftProgramObj []model.AircraftProgram = getAllAircraftProgramRequestResult.Data. ([]model.AircraftProgram)
		
	equalAircraftProgram := cmp.Equal(createAircraftProgramObj.ID, getAllAircraftProgramObj[len(getAllAircraftProgramObj)-1].ID)
		
	if equalAircraftProgram == false {
		t.Errorf( "Created object is not equal to the last entry in AircraftProgram[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AircraftProgram
	// --------------------------------------------------------------	
	deleteAircraftProgramRequestResult := dao.DeleteAircraftProgram(uint64(createAircraftProgramObj.ID))

	if deleteAircraftProgramRequestResult.Success == false {
			t.Errorf(deleteAircraftProgramRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AircraftProgram success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAircraftProgramRequestResult = dao.GetAircraftProgram( uint64(createAircraftProgramObj.ID) )
	
	if getAircraftProgramRequestResult.Success == true {
		t.Errorf(getAircraftProgramRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAircraftFamilyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AircraftFamily
	//----------------------------------------------------------------------------
	AircraftFamilyObj := model.AircraftFamily                                                            {Name:"test value for Name",FamilyCode:"test value for FamilyCode"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAircraftFamilyRequestResult := dao.CreateAircraftFamily( AircraftFamilyObj )
	
	if createAircraftFamilyRequestResult.Success == false {
		t.Errorf(createAircraftFamilyRequestResult.Msg)
	} else {
		fmt.Println("Check Create AircraftFamily success...")
	}
	
	createAircraftFamilyObj,_ := createAircraftFamilyRequestResult.Data. (model.AircraftFamily)

	// --------------------------------------------------------------
	// Check AircraftFamily Obj ID
	// --------------------------------------------------------------	
	if createAircraftFamilyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AircraftFamily" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAircraftFamilyRequestResult := dao.GetAircraftFamily( uint64(createAircraftFamilyObj.ID) )
	
	if getAircraftFamilyRequestResult.Success == false {
		t.Errorf(getAircraftFamilyRequestResult.Msg)
	} else {
		fmt.Println("Check Get AircraftFamily success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAircraftFamilyObj,_ := getAircraftFamilyRequestResult.Data. (model.AircraftFamily)
	compareAircraftFamily := cmp.Equal(createAircraftFamilyObj.ID, getAircraftFamilyObj.ID)
	
	if  compareAircraftFamily == false	{
		t.Errorf( "Created AircraftFamily object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAircraftFamilyRequestResult := dao.GetAllAircraftFamily()

	if getAllAircraftFamilyRequestResult.Success == false {
			t.Errorf(getAllAircraftFamilyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AircraftFamily success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAircraftFamilyObj []model.AircraftFamily = getAllAircraftFamilyRequestResult.Data. ([]model.AircraftFamily)
		
	equalAircraftFamily := cmp.Equal(createAircraftFamilyObj.ID, getAllAircraftFamilyObj[len(getAllAircraftFamilyObj)-1].ID)
		
	if equalAircraftFamily == false {
		t.Errorf( "Created object is not equal to the last entry in AircraftFamily[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AircraftFamily
	// --------------------------------------------------------------	
	deleteAircraftFamilyRequestResult := dao.DeleteAircraftFamily(uint64(createAircraftFamilyObj.ID))

	if deleteAircraftFamilyRequestResult.Success == false {
			t.Errorf(deleteAircraftFamilyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AircraftFamily success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAircraftFamilyRequestResult = dao.GetAircraftFamily( uint64(createAircraftFamilyObj.ID) )
	
	if getAircraftFamilyRequestResult.Success == true {
		t.Errorf(getAircraftFamilyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAircraftModelCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AircraftModel
	//----------------------------------------------------------------------------
	AircraftModelObj := model.AircraftModel                                                                            {Name:"test value for Name",ModelDesignation:"test value for ModelDesignation",AircraftType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAircraftModelRequestResult := dao.CreateAircraftModel( AircraftModelObj )
	
	if createAircraftModelRequestResult.Success == false {
		t.Errorf(createAircraftModelRequestResult.Msg)
	} else {
		fmt.Println("Check Create AircraftModel success...")
	}
	
	createAircraftModelObj,_ := createAircraftModelRequestResult.Data. (model.AircraftModel)

	// --------------------------------------------------------------
	// Check AircraftModel Obj ID
	// --------------------------------------------------------------	
	if createAircraftModelObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AircraftModel" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAircraftModelRequestResult := dao.GetAircraftModel( uint64(createAircraftModelObj.ID) )
	
	if getAircraftModelRequestResult.Success == false {
		t.Errorf(getAircraftModelRequestResult.Msg)
	} else {
		fmt.Println("Check Get AircraftModel success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAircraftModelObj,_ := getAircraftModelRequestResult.Data. (model.AircraftModel)
	compareAircraftModel := cmp.Equal(createAircraftModelObj.ID, getAircraftModelObj.ID)
	
	if  compareAircraftModel == false	{
		t.Errorf( "Created AircraftModel object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAircraftModelRequestResult := dao.GetAllAircraftModel()

	if getAllAircraftModelRequestResult.Success == false {
			t.Errorf(getAllAircraftModelRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AircraftModel success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAircraftModelObj []model.AircraftModel = getAllAircraftModelRequestResult.Data. ([]model.AircraftModel)
		
	equalAircraftModel := cmp.Equal(createAircraftModelObj.ID, getAllAircraftModelObj[len(getAllAircraftModelObj)-1].ID)
		
	if equalAircraftModel == false {
		t.Errorf( "Created object is not equal to the last entry in AircraftModel[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AircraftModel
	// --------------------------------------------------------------	
	deleteAircraftModelRequestResult := dao.DeleteAircraftModel(uint64(createAircraftModelObj.ID))

	if deleteAircraftModelRequestResult.Success == false {
			t.Errorf(deleteAircraftModelRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AircraftModel success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAircraftModelRequestResult = dao.GetAircraftModel( uint64(createAircraftModelObj.ID) )
	
	if getAircraftModelRequestResult.Success == true {
		t.Errorf(getAircraftModelRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEngineTypeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for EngineType
	//----------------------------------------------------------------------------
	EngineTypeObj := model.EngineType                                                                                                    {EngineModelCode:"test value for EngineModelCode",MaxThrustKn:"test value",Category:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEngineTypeRequestResult := dao.CreateEngineType( EngineTypeObj )
	
	if createEngineTypeRequestResult.Success == false {
		t.Errorf(createEngineTypeRequestResult.Msg)
	} else {
		fmt.Println("Check Create EngineType success...")
	}
	
	createEngineTypeObj,_ := createEngineTypeRequestResult.Data. (model.EngineType)

	// --------------------------------------------------------------
	// Check EngineType Obj ID
	// --------------------------------------------------------------	
	if createEngineTypeObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for EngineType" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEngineTypeRequestResult := dao.GetEngineType( uint64(createEngineTypeObj.ID) )
	
	if getEngineTypeRequestResult.Success == false {
		t.Errorf(getEngineTypeRequestResult.Msg)
	} else {
		fmt.Println("Check Get EngineType success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEngineTypeObj,_ := getEngineTypeRequestResult.Data. (model.EngineType)
	compareEngineType := cmp.Equal(createEngineTypeObj.ID, getEngineTypeObj.ID)
	
	if  compareEngineType == false	{
		t.Errorf( "Created EngineType object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEngineTypeRequestResult := dao.GetAllEngineType()

	if getAllEngineTypeRequestResult.Success == false {
			t.Errorf(getAllEngineTypeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll EngineType success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEngineTypeObj []model.EngineType = getAllEngineTypeRequestResult.Data. ([]model.EngineType)
		
	equalEngineType := cmp.Equal(createEngineTypeObj.ID, getAllEngineTypeObj[len(getAllEngineTypeObj)-1].ID)
		
	if equalEngineType == false {
		t.Errorf( "Created object is not equal to the last entry in EngineType[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for EngineType
	// --------------------------------------------------------------	
	deleteEngineTypeRequestResult := dao.DeleteEngineType(uint64(createEngineTypeObj.ID))

	if deleteEngineTypeRequestResult.Success == false {
			t.Errorf(deleteEngineTypeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion EngineType success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEngineTypeRequestResult = dao.GetEngineType( uint64(createEngineTypeObj.ID) )
	
	if getEngineTypeRequestResult.Success == true {
		t.Errorf(getEngineTypeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAircraftVariantCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AircraftVariant
	//----------------------------------------------------------------------------
	AircraftVariantObj := model.AircraftVariant                                                                                                                    {VariantCode:"test value for VariantCode",RangeNm:100,MaxTakeoffWeightKg:"test value"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAircraftVariantRequestResult := dao.CreateAircraftVariant( AircraftVariantObj )
	
	if createAircraftVariantRequestResult.Success == false {
		t.Errorf(createAircraftVariantRequestResult.Msg)
	} else {
		fmt.Println("Check Create AircraftVariant success...")
	}
	
	createAircraftVariantObj,_ := createAircraftVariantRequestResult.Data. (model.AircraftVariant)

	// --------------------------------------------------------------
	// Check AircraftVariant Obj ID
	// --------------------------------------------------------------	
	if createAircraftVariantObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AircraftVariant" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAircraftVariantRequestResult := dao.GetAircraftVariant( uint64(createAircraftVariantObj.ID) )
	
	if getAircraftVariantRequestResult.Success == false {
		t.Errorf(getAircraftVariantRequestResult.Msg)
	} else {
		fmt.Println("Check Get AircraftVariant success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAircraftVariantObj,_ := getAircraftVariantRequestResult.Data. (model.AircraftVariant)
	compareAircraftVariant := cmp.Equal(createAircraftVariantObj.ID, getAircraftVariantObj.ID)
	
	if  compareAircraftVariant == false	{
		t.Errorf( "Created AircraftVariant object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAircraftVariantRequestResult := dao.GetAllAircraftVariant()

	if getAllAircraftVariantRequestResult.Success == false {
			t.Errorf(getAllAircraftVariantRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AircraftVariant success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAircraftVariantObj []model.AircraftVariant = getAllAircraftVariantRequestResult.Data. ([]model.AircraftVariant)
		
	equalAircraftVariant := cmp.Equal(createAircraftVariantObj.ID, getAllAircraftVariantObj[len(getAllAircraftVariantObj)-1].ID)
		
	if equalAircraftVariant == false {
		t.Errorf( "Created object is not equal to the last entry in AircraftVariant[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AircraftVariant
	// --------------------------------------------------------------	
	deleteAircraftVariantRequestResult := dao.DeleteAircraftVariant(uint64(createAircraftVariantObj.ID))

	if deleteAircraftVariantRequestResult.Success == false {
			t.Errorf(deleteAircraftVariantRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AircraftVariant success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAircraftVariantRequestResult = dao.GetAircraftVariant( uint64(createAircraftVariantObj.ID) )
	
	if getAircraftVariantRequestResult.Success == true {
		t.Errorf(getAircraftVariantRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAvionicsSuiteCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AvionicsSuite
	//----------------------------------------------------------------------------
	AvionicsSuiteObj := model.AvionicsSuite                                                            {SuiteName:"test value for SuiteName",SoftwareBaseline:"test value for SoftwareBaseline"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAvionicsSuiteRequestResult := dao.CreateAvionicsSuite( AvionicsSuiteObj )
	
	if createAvionicsSuiteRequestResult.Success == false {
		t.Errorf(createAvionicsSuiteRequestResult.Msg)
	} else {
		fmt.Println("Check Create AvionicsSuite success...")
	}
	
	createAvionicsSuiteObj,_ := createAvionicsSuiteRequestResult.Data. (model.AvionicsSuite)

	// --------------------------------------------------------------
	// Check AvionicsSuite Obj ID
	// --------------------------------------------------------------	
	if createAvionicsSuiteObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AvionicsSuite" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAvionicsSuiteRequestResult := dao.GetAvionicsSuite( uint64(createAvionicsSuiteObj.ID) )
	
	if getAvionicsSuiteRequestResult.Success == false {
		t.Errorf(getAvionicsSuiteRequestResult.Msg)
	} else {
		fmt.Println("Check Get AvionicsSuite success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAvionicsSuiteObj,_ := getAvionicsSuiteRequestResult.Data. (model.AvionicsSuite)
	compareAvionicsSuite := cmp.Equal(createAvionicsSuiteObj.ID, getAvionicsSuiteObj.ID)
	
	if  compareAvionicsSuite == false	{
		t.Errorf( "Created AvionicsSuite object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAvionicsSuiteRequestResult := dao.GetAllAvionicsSuite()

	if getAllAvionicsSuiteRequestResult.Success == false {
			t.Errorf(getAllAvionicsSuiteRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AvionicsSuite success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAvionicsSuiteObj []model.AvionicsSuite = getAllAvionicsSuiteRequestResult.Data. ([]model.AvionicsSuite)
		
	equalAvionicsSuite := cmp.Equal(createAvionicsSuiteObj.ID, getAllAvionicsSuiteObj[len(getAllAvionicsSuiteObj)-1].ID)
		
	if equalAvionicsSuite == false {
		t.Errorf( "Created object is not equal to the last entry in AvionicsSuite[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AvionicsSuite
	// --------------------------------------------------------------	
	deleteAvionicsSuiteRequestResult := dao.DeleteAvionicsSuite(uint64(createAvionicsSuiteObj.ID))

	if deleteAvionicsSuiteRequestResult.Success == false {
			t.Errorf(deleteAvionicsSuiteRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AvionicsSuite success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAvionicsSuiteRequestResult = dao.GetAvionicsSuite( uint64(createAvionicsSuiteObj.ID) )
	
	if getAvionicsSuiteRequestResult.Success == true {
		t.Errorf(getAvionicsSuiteRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAPUCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for APU
	//----------------------------------------------------------------------------
	APUObj := model.APU                            {Model_:"test value for Model_"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAPURequestResult := dao.CreateAPU( APUObj )
	
	if createAPURequestResult.Success == false {
		t.Errorf(createAPURequestResult.Msg)
	} else {
		fmt.Println("Check Create APU success...")
	}
	
	createAPUObj,_ := createAPURequestResult.Data. (model.APU)

	// --------------------------------------------------------------
	// Check APU Obj ID
	// --------------------------------------------------------------	
	if createAPUObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for APU" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAPURequestResult := dao.GetAPU( uint64(createAPUObj.ID) )
	
	if getAPURequestResult.Success == false {
		t.Errorf(getAPURequestResult.Msg)
	} else {
		fmt.Println("Check Get APU success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAPUObj,_ := getAPURequestResult.Data. (model.APU)
	compareAPU := cmp.Equal(createAPUObj.ID, getAPUObj.ID)
	
	if  compareAPU == false	{
		t.Errorf( "Created APU object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAPURequestResult := dao.GetAllAPU()

	if getAllAPURequestResult.Success == false {
			t.Errorf(getAllAPURequestResult.Msg)
	} else {
		fmt.Println("Check GetAll APU success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAPUObj []model.APU = getAllAPURequestResult.Data. ([]model.APU)
		
	equalAPU := cmp.Equal(createAPUObj.ID, getAllAPUObj[len(getAllAPUObj)-1].ID)
		
	if equalAPU == false {
		t.Errorf( "Created object is not equal to the last entry in APU[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for APU
	// --------------------------------------------------------------	
	deleteAPURequestResult := dao.DeleteAPU(uint64(createAPUObj.ID))

	if deleteAPURequestResult.Success == false {
			t.Errorf(deleteAPURequestResult.Msg)
	} else {
		fmt.Println("Check Deletion APU success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAPURequestResult = dao.GetAPU( uint64(createAPUObj.ID) )
	
	if getAPURequestResult.Success == true {
		t.Errorf(getAPURequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLandingGearCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LandingGear
	//----------------------------------------------------------------------------
	LandingGearObj := model.LandingGear                                            {SupplierPartNumber:"test value for SupplierPartNumber",GearType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLandingGearRequestResult := dao.CreateLandingGear( LandingGearObj )
	
	if createLandingGearRequestResult.Success == false {
		t.Errorf(createLandingGearRequestResult.Msg)
	} else {
		fmt.Println("Check Create LandingGear success...")
	}
	
	createLandingGearObj,_ := createLandingGearRequestResult.Data. (model.LandingGear)

	// --------------------------------------------------------------
	// Check LandingGear Obj ID
	// --------------------------------------------------------------	
	if createLandingGearObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LandingGear" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLandingGearRequestResult := dao.GetLandingGear( uint64(createLandingGearObj.ID) )
	
	if getLandingGearRequestResult.Success == false {
		t.Errorf(getLandingGearRequestResult.Msg)
	} else {
		fmt.Println("Check Get LandingGear success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLandingGearObj,_ := getLandingGearRequestResult.Data. (model.LandingGear)
	compareLandingGear := cmp.Equal(createLandingGearObj.ID, getLandingGearObj.ID)
	
	if  compareLandingGear == false	{
		t.Errorf( "Created LandingGear object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLandingGearRequestResult := dao.GetAllLandingGear()

	if getAllLandingGearRequestResult.Success == false {
			t.Errorf(getAllLandingGearRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LandingGear success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLandingGearObj []model.LandingGear = getAllLandingGearRequestResult.Data. ([]model.LandingGear)
		
	equalLandingGear := cmp.Equal(createLandingGearObj.ID, getAllLandingGearObj[len(getAllLandingGearObj)-1].ID)
		
	if equalLandingGear == false {
		t.Errorf( "Created object is not equal to the last entry in LandingGear[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LandingGear
	// --------------------------------------------------------------	
	deleteLandingGearRequestResult := dao.DeleteLandingGear(uint64(createLandingGearObj.ID))

	if deleteLandingGearRequestResult.Success == false {
			t.Errorf(deleteLandingGearRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LandingGear success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLandingGearRequestResult = dao.GetLandingGear( uint64(createLandingGearObj.ID) )
	
	if getLandingGearRequestResult.Success == true {
		t.Errorf(getLandingGearRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAircraftOptionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AircraftOption
	//----------------------------------------------------------------------------
	AircraftOptionObj := model.AircraftOption                                                                            {Code:"test value for Code",Name:"test value for Name",OptionCategory:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAircraftOptionRequestResult := dao.CreateAircraftOption( AircraftOptionObj )
	
	if createAircraftOptionRequestResult.Success == false {
		t.Errorf(createAircraftOptionRequestResult.Msg)
	} else {
		fmt.Println("Check Create AircraftOption success...")
	}
	
	createAircraftOptionObj,_ := createAircraftOptionRequestResult.Data. (model.AircraftOption)

	// --------------------------------------------------------------
	// Check AircraftOption Obj ID
	// --------------------------------------------------------------	
	if createAircraftOptionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AircraftOption" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAircraftOptionRequestResult := dao.GetAircraftOption( uint64(createAircraftOptionObj.ID) )
	
	if getAircraftOptionRequestResult.Success == false {
		t.Errorf(getAircraftOptionRequestResult.Msg)
	} else {
		fmt.Println("Check Get AircraftOption success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAircraftOptionObj,_ := getAircraftOptionRequestResult.Data. (model.AircraftOption)
	compareAircraftOption := cmp.Equal(createAircraftOptionObj.ID, getAircraftOptionObj.ID)
	
	if  compareAircraftOption == false	{
		t.Errorf( "Created AircraftOption object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAircraftOptionRequestResult := dao.GetAllAircraftOption()

	if getAllAircraftOptionRequestResult.Success == false {
			t.Errorf(getAllAircraftOptionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AircraftOption success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAircraftOptionObj []model.AircraftOption = getAllAircraftOptionRequestResult.Data. ([]model.AircraftOption)
		
	equalAircraftOption := cmp.Equal(createAircraftOptionObj.ID, getAllAircraftOptionObj[len(getAllAircraftOptionObj)-1].ID)
		
	if equalAircraftOption == false {
		t.Errorf( "Created object is not equal to the last entry in AircraftOption[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AircraftOption
	// --------------------------------------------------------------	
	deleteAircraftOptionRequestResult := dao.DeleteAircraftOption(uint64(createAircraftOptionObj.ID))

	if deleteAircraftOptionRequestResult.Success == false {
			t.Errorf(deleteAircraftOptionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AircraftOption success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAircraftOptionRequestResult = dao.GetAircraftOption( uint64(createAircraftOptionObj.ID) )
	
	if getAircraftOptionRequestResult.Success == true {
		t.Errorf(getAircraftOptionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAircraftPackageCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AircraftPackage
	//----------------------------------------------------------------------------
	AircraftPackageObj := model.AircraftPackage                                            {Name:"test value for Name",PackageType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAircraftPackageRequestResult := dao.CreateAircraftPackage( AircraftPackageObj )
	
	if createAircraftPackageRequestResult.Success == false {
		t.Errorf(createAircraftPackageRequestResult.Msg)
	} else {
		fmt.Println("Check Create AircraftPackage success...")
	}
	
	createAircraftPackageObj,_ := createAircraftPackageRequestResult.Data. (model.AircraftPackage)

	// --------------------------------------------------------------
	// Check AircraftPackage Obj ID
	// --------------------------------------------------------------	
	if createAircraftPackageObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AircraftPackage" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAircraftPackageRequestResult := dao.GetAircraftPackage( uint64(createAircraftPackageObj.ID) )
	
	if getAircraftPackageRequestResult.Success == false {
		t.Errorf(getAircraftPackageRequestResult.Msg)
	} else {
		fmt.Println("Check Get AircraftPackage success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAircraftPackageObj,_ := getAircraftPackageRequestResult.Data. (model.AircraftPackage)
	compareAircraftPackage := cmp.Equal(createAircraftPackageObj.ID, getAircraftPackageObj.ID)
	
	if  compareAircraftPackage == false	{
		t.Errorf( "Created AircraftPackage object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAircraftPackageRequestResult := dao.GetAllAircraftPackage()

	if getAllAircraftPackageRequestResult.Success == false {
			t.Errorf(getAllAircraftPackageRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AircraftPackage success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAircraftPackageObj []model.AircraftPackage = getAllAircraftPackageRequestResult.Data. ([]model.AircraftPackage)
		
	equalAircraftPackage := cmp.Equal(createAircraftPackageObj.ID, getAllAircraftPackageObj[len(getAllAircraftPackageObj)-1].ID)
		
	if equalAircraftPackage == false {
		t.Errorf( "Created object is not equal to the last entry in AircraftPackage[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AircraftPackage
	// --------------------------------------------------------------	
	deleteAircraftPackageRequestResult := dao.DeleteAircraftPackage(uint64(createAircraftPackageObj.ID))

	if deleteAircraftPackageRequestResult.Success == false {
			t.Errorf(deleteAircraftPackageRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AircraftPackage success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAircraftPackageRequestResult = dao.GetAircraftPackage( uint64(createAircraftPackageObj.ID) )
	
	if getAircraftPackageRequestResult.Success == true {
		t.Errorf(getAircraftPackageRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSupplierCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Supplier
	//----------------------------------------------------------------------------
	SupplierObj := model.Supplier                                                            {Name:"test value for Name",SupplierType:0,ApprovalStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSupplierRequestResult := dao.CreateSupplier( SupplierObj )
	
	if createSupplierRequestResult.Success == false {
		t.Errorf(createSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check Create Supplier success...")
	}
	
	createSupplierObj,_ := createSupplierRequestResult.Data. (model.Supplier)

	// --------------------------------------------------------------
	// Check Supplier Obj ID
	// --------------------------------------------------------------	
	if createSupplierObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Supplier" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSupplierRequestResult := dao.GetSupplier( uint64(createSupplierObj.ID) )
	
	if getSupplierRequestResult.Success == false {
		t.Errorf(getSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check Get Supplier success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSupplierObj,_ := getSupplierRequestResult.Data. (model.Supplier)
	compareSupplier := cmp.Equal(createSupplierObj.ID, getSupplierObj.ID)
	
	if  compareSupplier == false	{
		t.Errorf( "Created Supplier object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSupplierRequestResult := dao.GetAllSupplier()

	if getAllSupplierRequestResult.Success == false {
			t.Errorf(getAllSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Supplier success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSupplierObj []model.Supplier = getAllSupplierRequestResult.Data. ([]model.Supplier)
		
	equalSupplier := cmp.Equal(createSupplierObj.ID, getAllSupplierObj[len(getAllSupplierObj)-1].ID)
		
	if equalSupplier == false {
		t.Errorf( "Created object is not equal to the last entry in Supplier[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Supplier
	// --------------------------------------------------------------	
	deleteSupplierRequestResult := dao.DeleteSupplier(uint64(createSupplierObj.ID))

	if deleteSupplierRequestResult.Success == false {
			t.Errorf(deleteSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Supplier success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSupplierRequestResult = dao.GetSupplier( uint64(createSupplierObj.ID) )
	
	if getSupplierRequestResult.Success == true {
		t.Errorf(getSupplierRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestComponent_CRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Component_
	//----------------------------------------------------------------------------
	Component_Obj := model.Component_                                                                                            {PartNumber:"test value for PartNumber",Name:"test value for Name",ComponentCategory:0,SerializationMethod:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createComponent_RequestResult := dao.CreateComponent_( Component_Obj )
	
	if createComponent_RequestResult.Success == false {
		t.Errorf(createComponent_RequestResult.Msg)
	} else {
		fmt.Println("Check Create Component_ success...")
	}
	
	createComponent_Obj,_ := createComponent_RequestResult.Data. (model.Component_)

	// --------------------------------------------------------------
	// Check Component_ Obj ID
	// --------------------------------------------------------------	
	if createComponent_Obj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Component_" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getComponent_RequestResult := dao.GetComponent_( uint64(createComponent_Obj.ID) )
	
	if getComponent_RequestResult.Success == false {
		t.Errorf(getComponent_RequestResult.Msg)
	} else {
		fmt.Println("Check Get Component_ success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getComponent_Obj,_ := getComponent_RequestResult.Data. (model.Component_)
	compareComponent_ := cmp.Equal(createComponent_Obj.ID, getComponent_Obj.ID)
	
	if  compareComponent_ == false	{
		t.Errorf( "Created Component_ object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllComponent_RequestResult := dao.GetAllComponent_()

	if getAllComponent_RequestResult.Success == false {
			t.Errorf(getAllComponent_RequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Component_ success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllComponent_Obj []model.Component_ = getAllComponent_RequestResult.Data. ([]model.Component_)
		
	equalComponent_ := cmp.Equal(createComponent_Obj.ID, getAllComponent_Obj[len(getAllComponent_Obj)-1].ID)
		
	if equalComponent_ == false {
		t.Errorf( "Created object is not equal to the last entry in Component_[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Component_
	// --------------------------------------------------------------	
	deleteComponent_RequestResult := dao.DeleteComponent_(uint64(createComponent_Obj.ID))

	if deleteComponent_RequestResult.Success == false {
			t.Errorf(deleteComponent_RequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Component_ success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getComponent_RequestResult = dao.GetComponent_( uint64(createComponent_Obj.ID) )
	
	if getComponent_RequestResult.Success == true {
		t.Errorf(getComponent_RequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPlantCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Plant
	//----------------------------------------------------------------------------
	PlantObj := model.Plant                                                                            {Name:"test value for Name",PlantCode:"test value for PlantCode",Address:new Address()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPlantRequestResult := dao.CreatePlant( PlantObj )
	
	if createPlantRequestResult.Success == false {
		t.Errorf(createPlantRequestResult.Msg)
	} else {
		fmt.Println("Check Create Plant success...")
	}
	
	createPlantObj,_ := createPlantRequestResult.Data. (model.Plant)

	// --------------------------------------------------------------
	// Check Plant Obj ID
	// --------------------------------------------------------------	
	if createPlantObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Plant" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPlantRequestResult := dao.GetPlant( uint64(createPlantObj.ID) )
	
	if getPlantRequestResult.Success == false {
		t.Errorf(getPlantRequestResult.Msg)
	} else {
		fmt.Println("Check Get Plant success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPlantObj,_ := getPlantRequestResult.Data. (model.Plant)
	comparePlant := cmp.Equal(createPlantObj.ID, getPlantObj.ID)
	
	if  comparePlant == false	{
		t.Errorf( "Created Plant object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPlantRequestResult := dao.GetAllPlant()

	if getAllPlantRequestResult.Success == false {
			t.Errorf(getAllPlantRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Plant success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPlantObj []model.Plant = getAllPlantRequestResult.Data. ([]model.Plant)
		
	equalPlant := cmp.Equal(createPlantObj.ID, getAllPlantObj[len(getAllPlantObj)-1].ID)
		
	if equalPlant == false {
		t.Errorf( "Created object is not equal to the last entry in Plant[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Plant
	// --------------------------------------------------------------	
	deletePlantRequestResult := dao.DeletePlant(uint64(createPlantObj.ID))

	if deletePlantRequestResult.Success == false {
			t.Errorf(deletePlantRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Plant success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPlantRequestResult = dao.GetPlant( uint64(createPlantObj.ID) )
	
	if getPlantRequestResult.Success == true {
		t.Errorf(getPlantRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestProductionLineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ProductionLine
	//----------------------------------------------------------------------------
	ProductionLineObj := model.ProductionLine                                            {Name:"test value for Name",LineType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createProductionLineRequestResult := dao.CreateProductionLine( ProductionLineObj )
	
	if createProductionLineRequestResult.Success == false {
		t.Errorf(createProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Check Create ProductionLine success...")
	}
	
	createProductionLineObj,_ := createProductionLineRequestResult.Data. (model.ProductionLine)

	// --------------------------------------------------------------
	// Check ProductionLine Obj ID
	// --------------------------------------------------------------	
	if createProductionLineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ProductionLine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getProductionLineRequestResult := dao.GetProductionLine( uint64(createProductionLineObj.ID) )
	
	if getProductionLineRequestResult.Success == false {
		t.Errorf(getProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Check Get ProductionLine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getProductionLineObj,_ := getProductionLineRequestResult.Data. (model.ProductionLine)
	compareProductionLine := cmp.Equal(createProductionLineObj.ID, getProductionLineObj.ID)
	
	if  compareProductionLine == false	{
		t.Errorf( "Created ProductionLine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllProductionLineRequestResult := dao.GetAllProductionLine()

	if getAllProductionLineRequestResult.Success == false {
			t.Errorf(getAllProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ProductionLine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllProductionLineObj []model.ProductionLine = getAllProductionLineRequestResult.Data. ([]model.ProductionLine)
		
	equalProductionLine := cmp.Equal(createProductionLineObj.ID, getAllProductionLineObj[len(getAllProductionLineObj)-1].ID)
		
	if equalProductionLine == false {
		t.Errorf( "Created object is not equal to the last entry in ProductionLine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ProductionLine
	// --------------------------------------------------------------	
	deleteProductionLineRequestResult := dao.DeleteProductionLine(uint64(createProductionLineObj.ID))

	if deleteProductionLineRequestResult.Success == false {
			t.Errorf(deleteProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ProductionLine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getProductionLineRequestResult = dao.GetProductionLine( uint64(createProductionLineObj.ID) )
	
	if getProductionLineRequestResult.Success == true {
		t.Errorf(getProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWorkCenterCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for WorkCenter
	//----------------------------------------------------------------------------
	WorkCenterObj := model.WorkCenter                                                            {Name:"test value for Name",Capability:"test value for Capability"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWorkCenterRequestResult := dao.CreateWorkCenter( WorkCenterObj )
	
	if createWorkCenterRequestResult.Success == false {
		t.Errorf(createWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Create WorkCenter success...")
	}
	
	createWorkCenterObj,_ := createWorkCenterRequestResult.Data. (model.WorkCenter)

	// --------------------------------------------------------------
	// Check WorkCenter Obj ID
	// --------------------------------------------------------------	
	if createWorkCenterObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for WorkCenter" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWorkCenterRequestResult := dao.GetWorkCenter( uint64(createWorkCenterObj.ID) )
	
	if getWorkCenterRequestResult.Success == false {
		t.Errorf(getWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Get WorkCenter success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWorkCenterObj,_ := getWorkCenterRequestResult.Data. (model.WorkCenter)
	compareWorkCenter := cmp.Equal(createWorkCenterObj.ID, getWorkCenterObj.ID)
	
	if  compareWorkCenter == false	{
		t.Errorf( "Created WorkCenter object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWorkCenterRequestResult := dao.GetAllWorkCenter()

	if getAllWorkCenterRequestResult.Success == false {
			t.Errorf(getAllWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll WorkCenter success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWorkCenterObj []model.WorkCenter = getAllWorkCenterRequestResult.Data. ([]model.WorkCenter)
		
	equalWorkCenter := cmp.Equal(createWorkCenterObj.ID, getAllWorkCenterObj[len(getAllWorkCenterObj)-1].ID)
		
	if equalWorkCenter == false {
		t.Errorf( "Created object is not equal to the last entry in WorkCenter[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for WorkCenter
	// --------------------------------------------------------------	
	deleteWorkCenterRequestResult := dao.DeleteWorkCenter(uint64(createWorkCenterObj.ID))

	if deleteWorkCenterRequestResult.Success == false {
			t.Errorf(deleteWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion WorkCenter success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWorkCenterRequestResult = dao.GetWorkCenter( uint64(createWorkCenterObj.ID) )
	
	if getWorkCenterRequestResult.Success == true {
		t.Errorf(getWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestProductionOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ProductionOrder
	//----------------------------------------------------------------------------
	ProductionOrderObj := model.ProductionOrder                                            {OrderNumber:"test value for OrderNumber",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createProductionOrderRequestResult := dao.CreateProductionOrder( ProductionOrderObj )
	
	if createProductionOrderRequestResult.Success == false {
		t.Errorf(createProductionOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create ProductionOrder success...")
	}
	
	createProductionOrderObj,_ := createProductionOrderRequestResult.Data. (model.ProductionOrder)

	// --------------------------------------------------------------
	// Check ProductionOrder Obj ID
	// --------------------------------------------------------------	
	if createProductionOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ProductionOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getProductionOrderRequestResult := dao.GetProductionOrder( uint64(createProductionOrderObj.ID) )
	
	if getProductionOrderRequestResult.Success == false {
		t.Errorf(getProductionOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get ProductionOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getProductionOrderObj,_ := getProductionOrderRequestResult.Data. (model.ProductionOrder)
	compareProductionOrder := cmp.Equal(createProductionOrderObj.ID, getProductionOrderObj.ID)
	
	if  compareProductionOrder == false	{
		t.Errorf( "Created ProductionOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllProductionOrderRequestResult := dao.GetAllProductionOrder()

	if getAllProductionOrderRequestResult.Success == false {
			t.Errorf(getAllProductionOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ProductionOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllProductionOrderObj []model.ProductionOrder = getAllProductionOrderRequestResult.Data. ([]model.ProductionOrder)
		
	equalProductionOrder := cmp.Equal(createProductionOrderObj.ID, getAllProductionOrderObj[len(getAllProductionOrderObj)-1].ID)
		
	if equalProductionOrder == false {
		t.Errorf( "Created object is not equal to the last entry in ProductionOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ProductionOrder
	// --------------------------------------------------------------	
	deleteProductionOrderRequestResult := dao.DeleteProductionOrder(uint64(createProductionOrderObj.ID))

	if deleteProductionOrderRequestResult.Success == false {
			t.Errorf(deleteProductionOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ProductionOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getProductionOrderRequestResult = dao.GetProductionOrder( uint64(createProductionOrderObj.ID) )
	
	if getProductionOrderRequestResult.Success == true {
		t.Errorf(getProductionOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBuildScheduleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BuildSchedule
	//----------------------------------------------------------------------------
	BuildScheduleObj := model.BuildSchedule                                            {ScheduleNumber:"test value for ScheduleNumber",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBuildScheduleRequestResult := dao.CreateBuildSchedule( BuildScheduleObj )
	
	if createBuildScheduleRequestResult.Success == false {
		t.Errorf(createBuildScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Create BuildSchedule success...")
	}
	
	createBuildScheduleObj,_ := createBuildScheduleRequestResult.Data. (model.BuildSchedule)

	// --------------------------------------------------------------
	// Check BuildSchedule Obj ID
	// --------------------------------------------------------------	
	if createBuildScheduleObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BuildSchedule" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBuildScheduleRequestResult := dao.GetBuildSchedule( uint64(createBuildScheduleObj.ID) )
	
	if getBuildScheduleRequestResult.Success == false {
		t.Errorf(getBuildScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Get BuildSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBuildScheduleObj,_ := getBuildScheduleRequestResult.Data. (model.BuildSchedule)
	compareBuildSchedule := cmp.Equal(createBuildScheduleObj.ID, getBuildScheduleObj.ID)
	
	if  compareBuildSchedule == false	{
		t.Errorf( "Created BuildSchedule object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBuildScheduleRequestResult := dao.GetAllBuildSchedule()

	if getAllBuildScheduleRequestResult.Success == false {
			t.Errorf(getAllBuildScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BuildSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBuildScheduleObj []model.BuildSchedule = getAllBuildScheduleRequestResult.Data. ([]model.BuildSchedule)
		
	equalBuildSchedule := cmp.Equal(createBuildScheduleObj.ID, getAllBuildScheduleObj[len(getAllBuildScheduleObj)-1].ID)
		
	if equalBuildSchedule == false {
		t.Errorf( "Created object is not equal to the last entry in BuildSchedule[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BuildSchedule
	// --------------------------------------------------------------	
	deleteBuildScheduleRequestResult := dao.DeleteBuildSchedule(uint64(createBuildScheduleObj.ID))

	if deleteBuildScheduleRequestResult.Success == false {
			t.Errorf(deleteBuildScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BuildSchedule success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBuildScheduleRequestResult = dao.GetBuildSchedule( uint64(createBuildScheduleObj.ID) )
	
	if getBuildScheduleRequestResult.Success == true {
		t.Errorf(getBuildScheduleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWarehouseCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Warehouse
	//----------------------------------------------------------------------------
	WarehouseObj := model.Warehouse                            {Name:"test value for Name"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWarehouseRequestResult := dao.CreateWarehouse( WarehouseObj )
	
	if createWarehouseRequestResult.Success == false {
		t.Errorf(createWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Check Create Warehouse success...")
	}
	
	createWarehouseObj,_ := createWarehouseRequestResult.Data. (model.Warehouse)

	// --------------------------------------------------------------
	// Check Warehouse Obj ID
	// --------------------------------------------------------------	
	if createWarehouseObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Warehouse" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWarehouseRequestResult := dao.GetWarehouse( uint64(createWarehouseObj.ID) )
	
	if getWarehouseRequestResult.Success == false {
		t.Errorf(getWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Check Get Warehouse success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWarehouseObj,_ := getWarehouseRequestResult.Data. (model.Warehouse)
	compareWarehouse := cmp.Equal(createWarehouseObj.ID, getWarehouseObj.ID)
	
	if  compareWarehouse == false	{
		t.Errorf( "Created Warehouse object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWarehouseRequestResult := dao.GetAllWarehouse()

	if getAllWarehouseRequestResult.Success == false {
			t.Errorf(getAllWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Warehouse success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWarehouseObj []model.Warehouse = getAllWarehouseRequestResult.Data. ([]model.Warehouse)
		
	equalWarehouse := cmp.Equal(createWarehouseObj.ID, getAllWarehouseObj[len(getAllWarehouseObj)-1].ID)
		
	if equalWarehouse == false {
		t.Errorf( "Created object is not equal to the last entry in Warehouse[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Warehouse
	// --------------------------------------------------------------	
	deleteWarehouseRequestResult := dao.DeleteWarehouse(uint64(createWarehouseObj.ID))

	if deleteWarehouseRequestResult.Success == false {
			t.Errorf(deleteWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Warehouse success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWarehouseRequestResult = dao.GetWarehouse( uint64(createWarehouseObj.ID) )
	
	if getWarehouseRequestResult.Success == true {
		t.Errorf(getWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInventoryItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InventoryItem
	//----------------------------------------------------------------------------
	InventoryItemObj := model.InventoryItem                                                                                            {QuantityOnHand:100,QuantityReserved:100,LotNumber:"test value for LotNumber"}

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


func TestOperatorCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Operator
	//----------------------------------------------------------------------------
	OperatorObj := model.Operator                                                                            {Name:"test value for Name",IcaoDesignator:"test value for IcaoDesignator",OperatorType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOperatorRequestResult := dao.CreateOperator( OperatorObj )
	
	if createOperatorRequestResult.Success == false {
		t.Errorf(createOperatorRequestResult.Msg)
	} else {
		fmt.Println("Check Create Operator success...")
	}
	
	createOperatorObj,_ := createOperatorRequestResult.Data. (model.Operator)

	// --------------------------------------------------------------
	// Check Operator Obj ID
	// --------------------------------------------------------------	
	if createOperatorObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Operator" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOperatorRequestResult := dao.GetOperator( uint64(createOperatorObj.ID) )
	
	if getOperatorRequestResult.Success == false {
		t.Errorf(getOperatorRequestResult.Msg)
	} else {
		fmt.Println("Check Get Operator success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOperatorObj,_ := getOperatorRequestResult.Data. (model.Operator)
	compareOperator := cmp.Equal(createOperatorObj.ID, getOperatorObj.ID)
	
	if  compareOperator == false	{
		t.Errorf( "Created Operator object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOperatorRequestResult := dao.GetAllOperator()

	if getAllOperatorRequestResult.Success == false {
			t.Errorf(getAllOperatorRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Operator success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOperatorObj []model.Operator = getAllOperatorRequestResult.Data. ([]model.Operator)
		
	equalOperator := cmp.Equal(createOperatorObj.ID, getAllOperatorObj[len(getAllOperatorObj)-1].ID)
		
	if equalOperator == false {
		t.Errorf( "Created object is not equal to the last entry in Operator[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Operator
	// --------------------------------------------------------------	
	deleteOperatorRequestResult := dao.DeleteOperator(uint64(createOperatorObj.ID))

	if deleteOperatorRequestResult.Success == false {
			t.Errorf(deleteOperatorRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Operator success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOperatorRequestResult = dao.GetOperator( uint64(createOperatorObj.ID) )
	
	if getOperatorRequestResult.Success == true {
		t.Errorf(getOperatorRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAircraftOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AircraftOrder
	//----------------------------------------------------------------------------
	AircraftOrderObj := model.AircraftOrder                                                            {OrderNumber:"test value for OrderNumber",TotalAmount:new Money(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAircraftOrderRequestResult := dao.CreateAircraftOrder( AircraftOrderObj )
	
	if createAircraftOrderRequestResult.Success == false {
		t.Errorf(createAircraftOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create AircraftOrder success...")
	}
	
	createAircraftOrderObj,_ := createAircraftOrderRequestResult.Data. (model.AircraftOrder)

	// --------------------------------------------------------------
	// Check AircraftOrder Obj ID
	// --------------------------------------------------------------	
	if createAircraftOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AircraftOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAircraftOrderRequestResult := dao.GetAircraftOrder( uint64(createAircraftOrderObj.ID) )
	
	if getAircraftOrderRequestResult.Success == false {
		t.Errorf(getAircraftOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get AircraftOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAircraftOrderObj,_ := getAircraftOrderRequestResult.Data. (model.AircraftOrder)
	compareAircraftOrder := cmp.Equal(createAircraftOrderObj.ID, getAircraftOrderObj.ID)
	
	if  compareAircraftOrder == false	{
		t.Errorf( "Created AircraftOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAircraftOrderRequestResult := dao.GetAllAircraftOrder()

	if getAllAircraftOrderRequestResult.Success == false {
			t.Errorf(getAllAircraftOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AircraftOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAircraftOrderObj []model.AircraftOrder = getAllAircraftOrderRequestResult.Data. ([]model.AircraftOrder)
		
	equalAircraftOrder := cmp.Equal(createAircraftOrderObj.ID, getAllAircraftOrderObj[len(getAllAircraftOrderObj)-1].ID)
		
	if equalAircraftOrder == false {
		t.Errorf( "Created object is not equal to the last entry in AircraftOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AircraftOrder
	// --------------------------------------------------------------	
	deleteAircraftOrderRequestResult := dao.DeleteAircraftOrder(uint64(createAircraftOrderObj.ID))

	if deleteAircraftOrderRequestResult.Success == false {
			t.Errorf(deleteAircraftOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AircraftOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAircraftOrderRequestResult = dao.GetAircraftOrder( uint64(createAircraftOrderObj.ID) )
	
	if getAircraftOrderRequestResult.Success == true {
		t.Errorf(getAircraftOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestQuoteCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Quote
	//----------------------------------------------------------------------------
	QuoteObj := model.Quote                                            {QuoteNumber:"test value for QuoteNumber",TotalAmount:new Money()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createQuoteRequestResult := dao.CreateQuote( QuoteObj )
	
	if createQuoteRequestResult.Success == false {
		t.Errorf(createQuoteRequestResult.Msg)
	} else {
		fmt.Println("Check Create Quote success...")
	}
	
	createQuoteObj,_ := createQuoteRequestResult.Data. (model.Quote)

	// --------------------------------------------------------------
	// Check Quote Obj ID
	// --------------------------------------------------------------	
	if createQuoteObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Quote" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getQuoteRequestResult := dao.GetQuote( uint64(createQuoteObj.ID) )
	
	if getQuoteRequestResult.Success == false {
		t.Errorf(getQuoteRequestResult.Msg)
	} else {
		fmt.Println("Check Get Quote success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getQuoteObj,_ := getQuoteRequestResult.Data. (model.Quote)
	compareQuote := cmp.Equal(createQuoteObj.ID, getQuoteObj.ID)
	
	if  compareQuote == false	{
		t.Errorf( "Created Quote object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllQuoteRequestResult := dao.GetAllQuote()

	if getAllQuoteRequestResult.Success == false {
			t.Errorf(getAllQuoteRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Quote success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllQuoteObj []model.Quote = getAllQuoteRequestResult.Data. ([]model.Quote)
		
	equalQuote := cmp.Equal(createQuoteObj.ID, getAllQuoteObj[len(getAllQuoteObj)-1].ID)
		
	if equalQuote == false {
		t.Errorf( "Created object is not equal to the last entry in Quote[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Quote
	// --------------------------------------------------------------	
	deleteQuoteRequestResult := dao.DeleteQuote(uint64(createQuoteObj.ID))

	if deleteQuoteRequestResult.Success == false {
			t.Errorf(deleteQuoteRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Quote success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getQuoteRequestResult = dao.GetQuote( uint64(createQuoteObj.ID) )
	
	if getQuoteRequestResult.Success == true {
		t.Errorf(getQuoteRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPurchaseAgreementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PurchaseAgreement
	//----------------------------------------------------------------------------
	PurchaseAgreementObj := model.PurchaseAgreement                                                                                    {AgreementNumber:"test value for AgreementNumber",EffectiveDate:time.Now()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPurchaseAgreementRequestResult := dao.CreatePurchaseAgreement( PurchaseAgreementObj )
	
	if createPurchaseAgreementRequestResult.Success == false {
		t.Errorf(createPurchaseAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check Create PurchaseAgreement success...")
	}
	
	createPurchaseAgreementObj,_ := createPurchaseAgreementRequestResult.Data. (model.PurchaseAgreement)

	// --------------------------------------------------------------
	// Check PurchaseAgreement Obj ID
	// --------------------------------------------------------------	
	if createPurchaseAgreementObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PurchaseAgreement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPurchaseAgreementRequestResult := dao.GetPurchaseAgreement( uint64(createPurchaseAgreementObj.ID) )
	
	if getPurchaseAgreementRequestResult.Success == false {
		t.Errorf(getPurchaseAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check Get PurchaseAgreement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPurchaseAgreementObj,_ := getPurchaseAgreementRequestResult.Data. (model.PurchaseAgreement)
	comparePurchaseAgreement := cmp.Equal(createPurchaseAgreementObj.ID, getPurchaseAgreementObj.ID)
	
	if  comparePurchaseAgreement == false	{
		t.Errorf( "Created PurchaseAgreement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPurchaseAgreementRequestResult := dao.GetAllPurchaseAgreement()

	if getAllPurchaseAgreementRequestResult.Success == false {
			t.Errorf(getAllPurchaseAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PurchaseAgreement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPurchaseAgreementObj []model.PurchaseAgreement = getAllPurchaseAgreementRequestResult.Data. ([]model.PurchaseAgreement)
		
	equalPurchaseAgreement := cmp.Equal(createPurchaseAgreementObj.ID, getAllPurchaseAgreementObj[len(getAllPurchaseAgreementObj)-1].ID)
		
	if equalPurchaseAgreement == false {
		t.Errorf( "Created object is not equal to the last entry in PurchaseAgreement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PurchaseAgreement
	// --------------------------------------------------------------	
	deletePurchaseAgreementRequestResult := dao.DeletePurchaseAgreement(uint64(createPurchaseAgreementObj.ID))

	if deletePurchaseAgreementRequestResult.Success == false {
			t.Errorf(deletePurchaseAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PurchaseAgreement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPurchaseAgreementRequestResult = dao.GetPurchaseAgreement( uint64(createPurchaseAgreementObj.ID) )
	
	if getPurchaseAgreementRequestResult.Success == true {
		t.Errorf(getPurchaseAgreementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAircraftCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Aircraft
	//----------------------------------------------------------------------------
	AircraftObj := model.Aircraft                                                                    {Msn:new MSN(),DeliveryDate:time.Now()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAircraftRequestResult := dao.CreateAircraft( AircraftObj )
	
	if createAircraftRequestResult.Success == false {
		t.Errorf(createAircraftRequestResult.Msg)
	} else {
		fmt.Println("Check Create Aircraft success...")
	}
	
	createAircraftObj,_ := createAircraftRequestResult.Data. (model.Aircraft)

	// --------------------------------------------------------------
	// Check Aircraft Obj ID
	// --------------------------------------------------------------	
	if createAircraftObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Aircraft" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAircraftRequestResult := dao.GetAircraft( uint64(createAircraftObj.ID) )
	
	if getAircraftRequestResult.Success == false {
		t.Errorf(getAircraftRequestResult.Msg)
	} else {
		fmt.Println("Check Get Aircraft success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAircraftObj,_ := getAircraftRequestResult.Data. (model.Aircraft)
	compareAircraft := cmp.Equal(createAircraftObj.ID, getAircraftObj.ID)
	
	if  compareAircraft == false	{
		t.Errorf( "Created Aircraft object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAircraftRequestResult := dao.GetAllAircraft()

	if getAllAircraftRequestResult.Success == false {
			t.Errorf(getAllAircraftRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Aircraft success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAircraftObj []model.Aircraft = getAllAircraftRequestResult.Data. ([]model.Aircraft)
		
	equalAircraft := cmp.Equal(createAircraftObj.ID, getAllAircraftObj[len(getAllAircraftObj)-1].ID)
		
	if equalAircraft == false {
		t.Errorf( "Created object is not equal to the last entry in Aircraft[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Aircraft
	// --------------------------------------------------------------	
	deleteAircraftRequestResult := dao.DeleteAircraft(uint64(createAircraftObj.ID))

	if deleteAircraftRequestResult.Success == false {
			t.Errorf(deleteAircraftRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Aircraft success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAircraftRequestResult = dao.GetAircraft( uint64(createAircraftObj.ID) )
	
	if getAircraftRequestResult.Success == true {
		t.Errorf(getAircraftRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRegistrationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Registration
	//----------------------------------------------------------------------------
	RegistrationObj := model.Registration                                            {TailNumber:new TailNumber(),RegistryCountry:"test value for RegistryCountry"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRegistrationRequestResult := dao.CreateRegistration( RegistrationObj )
	
	if createRegistrationRequestResult.Success == false {
		t.Errorf(createRegistrationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Registration success...")
	}
	
	createRegistrationObj,_ := createRegistrationRequestResult.Data. (model.Registration)

	// --------------------------------------------------------------
	// Check Registration Obj ID
	// --------------------------------------------------------------	
	if createRegistrationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Registration" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRegistrationRequestResult := dao.GetRegistration( uint64(createRegistrationObj.ID) )
	
	if getRegistrationRequestResult.Success == false {
		t.Errorf(getRegistrationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Registration success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRegistrationObj,_ := getRegistrationRequestResult.Data. (model.Registration)
	compareRegistration := cmp.Equal(createRegistrationObj.ID, getRegistrationObj.ID)
	
	if  compareRegistration == false	{
		t.Errorf( "Created Registration object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRegistrationRequestResult := dao.GetAllRegistration()

	if getAllRegistrationRequestResult.Success == false {
			t.Errorf(getAllRegistrationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Registration success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRegistrationObj []model.Registration = getAllRegistrationRequestResult.Data. ([]model.Registration)
		
	equalRegistration := cmp.Equal(createRegistrationObj.ID, getAllRegistrationObj[len(getAllRegistrationObj)-1].ID)
		
	if equalRegistration == false {
		t.Errorf( "Created object is not equal to the last entry in Registration[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Registration
	// --------------------------------------------------------------	
	deleteRegistrationRequestResult := dao.DeleteRegistration(uint64(createRegistrationObj.ID))

	if deleteRegistrationRequestResult.Success == false {
			t.Errorf(deleteRegistrationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Registration success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRegistrationRequestResult = dao.GetRegistration( uint64(createRegistrationObj.ID) )
	
	if getRegistrationRequestResult.Success == true {
		t.Errorf(getRegistrationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWarrantyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Warranty
	//----------------------------------------------------------------------------
	WarrantyObj := model.Warranty                                            {CoverageMonths:100,WarrantyType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWarrantyRequestResult := dao.CreateWarranty( WarrantyObj )
	
	if createWarrantyRequestResult.Success == false {
		t.Errorf(createWarrantyRequestResult.Msg)
	} else {
		fmt.Println("Check Create Warranty success...")
	}
	
	createWarrantyObj,_ := createWarrantyRequestResult.Data. (model.Warranty)

	// --------------------------------------------------------------
	// Check Warranty Obj ID
	// --------------------------------------------------------------	
	if createWarrantyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Warranty" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWarrantyRequestResult := dao.GetWarranty( uint64(createWarrantyObj.ID) )
	
	if getWarrantyRequestResult.Success == false {
		t.Errorf(getWarrantyRequestResult.Msg)
	} else {
		fmt.Println("Check Get Warranty success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWarrantyObj,_ := getWarrantyRequestResult.Data. (model.Warranty)
	compareWarranty := cmp.Equal(createWarrantyObj.ID, getWarrantyObj.ID)
	
	if  compareWarranty == false	{
		t.Errorf( "Created Warranty object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWarrantyRequestResult := dao.GetAllWarranty()

	if getAllWarrantyRequestResult.Success == false {
			t.Errorf(getAllWarrantyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Warranty success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWarrantyObj []model.Warranty = getAllWarrantyRequestResult.Data. ([]model.Warranty)
		
	equalWarranty := cmp.Equal(createWarrantyObj.ID, getAllWarrantyObj[len(getAllWarrantyObj)-1].ID)
		
	if equalWarranty == false {
		t.Errorf( "Created object is not equal to the last entry in Warranty[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Warranty
	// --------------------------------------------------------------	
	deleteWarrantyRequestResult := dao.DeleteWarranty(uint64(createWarrantyObj.ID))

	if deleteWarrantyRequestResult.Success == false {
			t.Errorf(deleteWarrantyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Warranty success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWarrantyRequestResult = dao.GetWarranty( uint64(createWarrantyObj.ID) )
	
	if getWarrantyRequestResult.Success == true {
		t.Errorf(getWarrantyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCabinLayoutCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CabinLayout
	//----------------------------------------------------------------------------
	CabinLayoutObj := model.CabinLayout                                                                                            {LayoutCode:"test value for LayoutCode",TotalSeats:100,ClassLayout:"test value for ClassLayout"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCabinLayoutRequestResult := dao.CreateCabinLayout( CabinLayoutObj )
	
	if createCabinLayoutRequestResult.Success == false {
		t.Errorf(createCabinLayoutRequestResult.Msg)
	} else {
		fmt.Println("Check Create CabinLayout success...")
	}
	
	createCabinLayoutObj,_ := createCabinLayoutRequestResult.Data. (model.CabinLayout)

	// --------------------------------------------------------------
	// Check CabinLayout Obj ID
	// --------------------------------------------------------------	
	if createCabinLayoutObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CabinLayout" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCabinLayoutRequestResult := dao.GetCabinLayout( uint64(createCabinLayoutObj.ID) )
	
	if getCabinLayoutRequestResult.Success == false {
		t.Errorf(getCabinLayoutRequestResult.Msg)
	} else {
		fmt.Println("Check Get CabinLayout success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCabinLayoutObj,_ := getCabinLayoutRequestResult.Data. (model.CabinLayout)
	compareCabinLayout := cmp.Equal(createCabinLayoutObj.ID, getCabinLayoutObj.ID)
	
	if  compareCabinLayout == false	{
		t.Errorf( "Created CabinLayout object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCabinLayoutRequestResult := dao.GetAllCabinLayout()

	if getAllCabinLayoutRequestResult.Success == false {
			t.Errorf(getAllCabinLayoutRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CabinLayout success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCabinLayoutObj []model.CabinLayout = getAllCabinLayoutRequestResult.Data. ([]model.CabinLayout)
		
	equalCabinLayout := cmp.Equal(createCabinLayoutObj.ID, getAllCabinLayoutObj[len(getAllCabinLayoutObj)-1].ID)
		
	if equalCabinLayout == false {
		t.Errorf( "Created object is not equal to the last entry in CabinLayout[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CabinLayout
	// --------------------------------------------------------------	
	deleteCabinLayoutRequestResult := dao.DeleteCabinLayout(uint64(createCabinLayoutObj.ID))

	if deleteCabinLayoutRequestResult.Success == false {
			t.Errorf(deleteCabinLayoutRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CabinLayout success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCabinLayoutRequestResult = dao.GetCabinLayout( uint64(createCabinLayoutObj.ID) )
	
	if getCabinLayoutRequestResult.Success == true {
		t.Errorf(getCabinLayoutRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMROFacilityCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MROFacility
	//----------------------------------------------------------------------------
	MROFacilityObj := model.MROFacility                                                                            {Name:"test value for Name",ApprovalScope:"test value for ApprovalScope",Address:new Address()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMROFacilityRequestResult := dao.CreateMROFacility( MROFacilityObj )
	
	if createMROFacilityRequestResult.Success == false {
		t.Errorf(createMROFacilityRequestResult.Msg)
	} else {
		fmt.Println("Check Create MROFacility success...")
	}
	
	createMROFacilityObj,_ := createMROFacilityRequestResult.Data. (model.MROFacility)

	// --------------------------------------------------------------
	// Check MROFacility Obj ID
	// --------------------------------------------------------------	
	if createMROFacilityObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MROFacility" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMROFacilityRequestResult := dao.GetMROFacility( uint64(createMROFacilityObj.ID) )
	
	if getMROFacilityRequestResult.Success == false {
		t.Errorf(getMROFacilityRequestResult.Msg)
	} else {
		fmt.Println("Check Get MROFacility success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMROFacilityObj,_ := getMROFacilityRequestResult.Data. (model.MROFacility)
	compareMROFacility := cmp.Equal(createMROFacilityObj.ID, getMROFacilityObj.ID)
	
	if  compareMROFacility == false	{
		t.Errorf( "Created MROFacility object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMROFacilityRequestResult := dao.GetAllMROFacility()

	if getAllMROFacilityRequestResult.Success == false {
			t.Errorf(getAllMROFacilityRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MROFacility success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMROFacilityObj []model.MROFacility = getAllMROFacilityRequestResult.Data. ([]model.MROFacility)
		
	equalMROFacility := cmp.Equal(createMROFacilityObj.ID, getAllMROFacilityObj[len(getAllMROFacilityObj)-1].ID)
		
	if equalMROFacility == false {
		t.Errorf( "Created object is not equal to the last entry in MROFacility[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MROFacility
	// --------------------------------------------------------------	
	deleteMROFacilityRequestResult := dao.DeleteMROFacility(uint64(createMROFacilityObj.ID))

	if deleteMROFacilityRequestResult.Success == false {
			t.Errorf(deleteMROFacilityRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MROFacility success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMROFacilityRequestResult = dao.GetMROFacility( uint64(createMROFacilityObj.ID) )
	
	if getMROFacilityRequestResult.Success == true {
		t.Errorf(getMROFacilityRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMaintenanceAppointmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MaintenanceAppointment
	//----------------------------------------------------------------------------
	MaintenanceAppointmentObj := model.MaintenanceAppointment                                                                    {AppointmentDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMaintenanceAppointmentRequestResult := dao.CreateMaintenanceAppointment( MaintenanceAppointmentObj )
	
	if createMaintenanceAppointmentRequestResult.Success == false {
		t.Errorf(createMaintenanceAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create MaintenanceAppointment success...")
	}
	
	createMaintenanceAppointmentObj,_ := createMaintenanceAppointmentRequestResult.Data. (model.MaintenanceAppointment)

	// --------------------------------------------------------------
	// Check MaintenanceAppointment Obj ID
	// --------------------------------------------------------------	
	if createMaintenanceAppointmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MaintenanceAppointment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMaintenanceAppointmentRequestResult := dao.GetMaintenanceAppointment( uint64(createMaintenanceAppointmentObj.ID) )
	
	if getMaintenanceAppointmentRequestResult.Success == false {
		t.Errorf(getMaintenanceAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get MaintenanceAppointment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMaintenanceAppointmentObj,_ := getMaintenanceAppointmentRequestResult.Data. (model.MaintenanceAppointment)
	compareMaintenanceAppointment := cmp.Equal(createMaintenanceAppointmentObj.ID, getMaintenanceAppointmentObj.ID)
	
	if  compareMaintenanceAppointment == false	{
		t.Errorf( "Created MaintenanceAppointment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMaintenanceAppointmentRequestResult := dao.GetAllMaintenanceAppointment()

	if getAllMaintenanceAppointmentRequestResult.Success == false {
			t.Errorf(getAllMaintenanceAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MaintenanceAppointment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMaintenanceAppointmentObj []model.MaintenanceAppointment = getAllMaintenanceAppointmentRequestResult.Data. ([]model.MaintenanceAppointment)
		
	equalMaintenanceAppointment := cmp.Equal(createMaintenanceAppointmentObj.ID, getAllMaintenanceAppointmentObj[len(getAllMaintenanceAppointmentObj)-1].ID)
		
	if equalMaintenanceAppointment == false {
		t.Errorf( "Created object is not equal to the last entry in MaintenanceAppointment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MaintenanceAppointment
	// --------------------------------------------------------------	
	deleteMaintenanceAppointmentRequestResult := dao.DeleteMaintenanceAppointment(uint64(createMaintenanceAppointmentObj.ID))

	if deleteMaintenanceAppointmentRequestResult.Success == false {
			t.Errorf(deleteMaintenanceAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MaintenanceAppointment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMaintenanceAppointmentRequestResult = dao.GetMaintenanceAppointment( uint64(createMaintenanceAppointmentObj.ID) )
	
	if getMaintenanceAppointmentRequestResult.Success == true {
		t.Errorf(getMaintenanceAppointmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMaintenanceWorkOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MaintenanceWorkOrder
	//----------------------------------------------------------------------------
	MaintenanceWorkOrderObj := model.MaintenanceWorkOrder                                            {WorkOrderNumber:"test value for WorkOrderNumber",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMaintenanceWorkOrderRequestResult := dao.CreateMaintenanceWorkOrder( MaintenanceWorkOrderObj )
	
	if createMaintenanceWorkOrderRequestResult.Success == false {
		t.Errorf(createMaintenanceWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create MaintenanceWorkOrder success...")
	}
	
	createMaintenanceWorkOrderObj,_ := createMaintenanceWorkOrderRequestResult.Data. (model.MaintenanceWorkOrder)

	// --------------------------------------------------------------
	// Check MaintenanceWorkOrder Obj ID
	// --------------------------------------------------------------	
	if createMaintenanceWorkOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MaintenanceWorkOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMaintenanceWorkOrderRequestResult := dao.GetMaintenanceWorkOrder( uint64(createMaintenanceWorkOrderObj.ID) )
	
	if getMaintenanceWorkOrderRequestResult.Success == false {
		t.Errorf(getMaintenanceWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get MaintenanceWorkOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMaintenanceWorkOrderObj,_ := getMaintenanceWorkOrderRequestResult.Data. (model.MaintenanceWorkOrder)
	compareMaintenanceWorkOrder := cmp.Equal(createMaintenanceWorkOrderObj.ID, getMaintenanceWorkOrderObj.ID)
	
	if  compareMaintenanceWorkOrder == false	{
		t.Errorf( "Created MaintenanceWorkOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMaintenanceWorkOrderRequestResult := dao.GetAllMaintenanceWorkOrder()

	if getAllMaintenanceWorkOrderRequestResult.Success == false {
			t.Errorf(getAllMaintenanceWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MaintenanceWorkOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMaintenanceWorkOrderObj []model.MaintenanceWorkOrder = getAllMaintenanceWorkOrderRequestResult.Data. ([]model.MaintenanceWorkOrder)
		
	equalMaintenanceWorkOrder := cmp.Equal(createMaintenanceWorkOrderObj.ID, getAllMaintenanceWorkOrderObj[len(getAllMaintenanceWorkOrderObj)-1].ID)
		
	if equalMaintenanceWorkOrder == false {
		t.Errorf( "Created object is not equal to the last entry in MaintenanceWorkOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MaintenanceWorkOrder
	// --------------------------------------------------------------	
	deleteMaintenanceWorkOrderRequestResult := dao.DeleteMaintenanceWorkOrder(uint64(createMaintenanceWorkOrderObj.ID))

	if deleteMaintenanceWorkOrderRequestResult.Success == false {
			t.Errorf(deleteMaintenanceWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MaintenanceWorkOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMaintenanceWorkOrderRequestResult = dao.GetMaintenanceWorkOrder( uint64(createMaintenanceWorkOrderObj.ID) )
	
	if getMaintenanceWorkOrderRequestResult.Success == true {
		t.Errorf(getMaintenanceWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAirworthinessDirectiveCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AirworthinessDirective
	//----------------------------------------------------------------------------
	AirworthinessDirectiveObj := model.AirworthinessDirective                                                            {DirectiveNumber:"test value for DirectiveNumber",Title:"test value for Title"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAirworthinessDirectiveRequestResult := dao.CreateAirworthinessDirective( AirworthinessDirectiveObj )
	
	if createAirworthinessDirectiveRequestResult.Success == false {
		t.Errorf(createAirworthinessDirectiveRequestResult.Msg)
	} else {
		fmt.Println("Check Create AirworthinessDirective success...")
	}
	
	createAirworthinessDirectiveObj,_ := createAirworthinessDirectiveRequestResult.Data. (model.AirworthinessDirective)

	// --------------------------------------------------------------
	// Check AirworthinessDirective Obj ID
	// --------------------------------------------------------------	
	if createAirworthinessDirectiveObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AirworthinessDirective" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAirworthinessDirectiveRequestResult := dao.GetAirworthinessDirective( uint64(createAirworthinessDirectiveObj.ID) )
	
	if getAirworthinessDirectiveRequestResult.Success == false {
		t.Errorf(getAirworthinessDirectiveRequestResult.Msg)
	} else {
		fmt.Println("Check Get AirworthinessDirective success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAirworthinessDirectiveObj,_ := getAirworthinessDirectiveRequestResult.Data. (model.AirworthinessDirective)
	compareAirworthinessDirective := cmp.Equal(createAirworthinessDirectiveObj.ID, getAirworthinessDirectiveObj.ID)
	
	if  compareAirworthinessDirective == false	{
		t.Errorf( "Created AirworthinessDirective object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAirworthinessDirectiveRequestResult := dao.GetAllAirworthinessDirective()

	if getAllAirworthinessDirectiveRequestResult.Success == false {
			t.Errorf(getAllAirworthinessDirectiveRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AirworthinessDirective success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAirworthinessDirectiveObj []model.AirworthinessDirective = getAllAirworthinessDirectiveRequestResult.Data. ([]model.AirworthinessDirective)
		
	equalAirworthinessDirective := cmp.Equal(createAirworthinessDirectiveObj.ID, getAllAirworthinessDirectiveObj[len(getAllAirworthinessDirectiveObj)-1].ID)
		
	if equalAirworthinessDirective == false {
		t.Errorf( "Created object is not equal to the last entry in AirworthinessDirective[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AirworthinessDirective
	// --------------------------------------------------------------	
	deleteAirworthinessDirectiveRequestResult := dao.DeleteAirworthinessDirective(uint64(createAirworthinessDirectiveObj.ID))

	if deleteAirworthinessDirectiveRequestResult.Success == false {
			t.Errorf(deleteAirworthinessDirectiveRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AirworthinessDirective success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAirworthinessDirectiveRequestResult = dao.GetAirworthinessDirective( uint64(createAirworthinessDirectiveObj.ID) )
	
	if getAirworthinessDirectiveRequestResult.Success == true {
		t.Errorf(getAirworthinessDirectiveRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestServiceBulletinCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ServiceBulletin
	//----------------------------------------------------------------------------
	ServiceBulletinObj := model.ServiceBulletin                                            {BulletinNumber:"test value for BulletinNumber",Category:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createServiceBulletinRequestResult := dao.CreateServiceBulletin( ServiceBulletinObj )
	
	if createServiceBulletinRequestResult.Success == false {
		t.Errorf(createServiceBulletinRequestResult.Msg)
	} else {
		fmt.Println("Check Create ServiceBulletin success...")
	}
	
	createServiceBulletinObj,_ := createServiceBulletinRequestResult.Data. (model.ServiceBulletin)

	// --------------------------------------------------------------
	// Check ServiceBulletin Obj ID
	// --------------------------------------------------------------	
	if createServiceBulletinObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ServiceBulletin" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getServiceBulletinRequestResult := dao.GetServiceBulletin( uint64(createServiceBulletinObj.ID) )
	
	if getServiceBulletinRequestResult.Success == false {
		t.Errorf(getServiceBulletinRequestResult.Msg)
	} else {
		fmt.Println("Check Get ServiceBulletin success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getServiceBulletinObj,_ := getServiceBulletinRequestResult.Data. (model.ServiceBulletin)
	compareServiceBulletin := cmp.Equal(createServiceBulletinObj.ID, getServiceBulletinObj.ID)
	
	if  compareServiceBulletin == false	{
		t.Errorf( "Created ServiceBulletin object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllServiceBulletinRequestResult := dao.GetAllServiceBulletin()

	if getAllServiceBulletinRequestResult.Success == false {
			t.Errorf(getAllServiceBulletinRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ServiceBulletin success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllServiceBulletinObj []model.ServiceBulletin = getAllServiceBulletinRequestResult.Data. ([]model.ServiceBulletin)
		
	equalServiceBulletin := cmp.Equal(createServiceBulletinObj.ID, getAllServiceBulletinObj[len(getAllServiceBulletinObj)-1].ID)
		
	if equalServiceBulletin == false {
		t.Errorf( "Created object is not equal to the last entry in ServiceBulletin[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ServiceBulletin
	// --------------------------------------------------------------	
	deleteServiceBulletinRequestResult := dao.DeleteServiceBulletin(uint64(createServiceBulletinObj.ID))

	if deleteServiceBulletinRequestResult.Success == false {
			t.Errorf(deleteServiceBulletinRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ServiceBulletin success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getServiceBulletinRequestResult = dao.GetServiceBulletin( uint64(createServiceBulletinObj.ID) )
	
	if getServiceBulletinRequestResult.Success == true {
		t.Errorf(getServiceBulletinRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestConnectedAircraftCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ConnectedAircraft
	//----------------------------------------------------------------------------
	ConnectedAircraftObj := model.ConnectedAircraft                                            {CommunicationsProvider:"test value for CommunicationsProvider",ConnectivityStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createConnectedAircraftRequestResult := dao.CreateConnectedAircraft( ConnectedAircraftObj )
	
	if createConnectedAircraftRequestResult.Success == false {
		t.Errorf(createConnectedAircraftRequestResult.Msg)
	} else {
		fmt.Println("Check Create ConnectedAircraft success...")
	}
	
	createConnectedAircraftObj,_ := createConnectedAircraftRequestResult.Data. (model.ConnectedAircraft)

	// --------------------------------------------------------------
	// Check ConnectedAircraft Obj ID
	// --------------------------------------------------------------	
	if createConnectedAircraftObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ConnectedAircraft" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getConnectedAircraftRequestResult := dao.GetConnectedAircraft( uint64(createConnectedAircraftObj.ID) )
	
	if getConnectedAircraftRequestResult.Success == false {
		t.Errorf(getConnectedAircraftRequestResult.Msg)
	} else {
		fmt.Println("Check Get ConnectedAircraft success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getConnectedAircraftObj,_ := getConnectedAircraftRequestResult.Data. (model.ConnectedAircraft)
	compareConnectedAircraft := cmp.Equal(createConnectedAircraftObj.ID, getConnectedAircraftObj.ID)
	
	if  compareConnectedAircraft == false	{
		t.Errorf( "Created ConnectedAircraft object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllConnectedAircraftRequestResult := dao.GetAllConnectedAircraft()

	if getAllConnectedAircraftRequestResult.Success == false {
			t.Errorf(getAllConnectedAircraftRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ConnectedAircraft success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllConnectedAircraftObj []model.ConnectedAircraft = getAllConnectedAircraftRequestResult.Data. ([]model.ConnectedAircraft)
		
	equalConnectedAircraft := cmp.Equal(createConnectedAircraftObj.ID, getAllConnectedAircraftObj[len(getAllConnectedAircraftObj)-1].ID)
		
	if equalConnectedAircraft == false {
		t.Errorf( "Created object is not equal to the last entry in ConnectedAircraft[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ConnectedAircraft
	// --------------------------------------------------------------	
	deleteConnectedAircraftRequestResult := dao.DeleteConnectedAircraft(uint64(createConnectedAircraftObj.ID))

	if deleteConnectedAircraftRequestResult.Success == false {
			t.Errorf(deleteConnectedAircraftRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ConnectedAircraft success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getConnectedAircraftRequestResult = dao.GetConnectedAircraft( uint64(createConnectedAircraftObj.ID) )
	
	if getConnectedAircraftRequestResult.Success == true {
		t.Errorf(getConnectedAircraftRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFlightHealthEventCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FlightHealthEvent
	//----------------------------------------------------------------------------
	FlightHealthEventObj := model.FlightHealthEvent                                            {EventCode:"test value for EventCode",Severity:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFlightHealthEventRequestResult := dao.CreateFlightHealthEvent( FlightHealthEventObj )
	
	if createFlightHealthEventRequestResult.Success == false {
		t.Errorf(createFlightHealthEventRequestResult.Msg)
	} else {
		fmt.Println("Check Create FlightHealthEvent success...")
	}
	
	createFlightHealthEventObj,_ := createFlightHealthEventRequestResult.Data. (model.FlightHealthEvent)

	// --------------------------------------------------------------
	// Check FlightHealthEvent Obj ID
	// --------------------------------------------------------------	
	if createFlightHealthEventObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for FlightHealthEvent" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFlightHealthEventRequestResult := dao.GetFlightHealthEvent( uint64(createFlightHealthEventObj.ID) )
	
	if getFlightHealthEventRequestResult.Success == false {
		t.Errorf(getFlightHealthEventRequestResult.Msg)
	} else {
		fmt.Println("Check Get FlightHealthEvent success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFlightHealthEventObj,_ := getFlightHealthEventRequestResult.Data. (model.FlightHealthEvent)
	compareFlightHealthEvent := cmp.Equal(createFlightHealthEventObj.ID, getFlightHealthEventObj.ID)
	
	if  compareFlightHealthEvent == false	{
		t.Errorf( "Created FlightHealthEvent object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFlightHealthEventRequestResult := dao.GetAllFlightHealthEvent()

	if getAllFlightHealthEventRequestResult.Success == false {
			t.Errorf(getAllFlightHealthEventRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FlightHealthEvent success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFlightHealthEventObj []model.FlightHealthEvent = getAllFlightHealthEventRequestResult.Data. ([]model.FlightHealthEvent)
		
	equalFlightHealthEvent := cmp.Equal(createFlightHealthEventObj.ID, getAllFlightHealthEventObj[len(getAllFlightHealthEventObj)-1].ID)
		
	if equalFlightHealthEvent == false {
		t.Errorf( "Created object is not equal to the last entry in FlightHealthEvent[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FlightHealthEvent
	// --------------------------------------------------------------	
	deleteFlightHealthEventRequestResult := dao.DeleteFlightHealthEvent(uint64(createFlightHealthEventObj.ID))

	if deleteFlightHealthEventRequestResult.Success == false {
			t.Errorf(deleteFlightHealthEventRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FlightHealthEvent success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFlightHealthEventRequestResult = dao.GetFlightHealthEvent( uint64(createFlightHealthEventObj.ID) )
	
	if getFlightHealthEventRequestResult.Success == true {
		t.Errorf(getFlightHealthEventRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSoftwareLoadCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SoftwareLoad
	//----------------------------------------------------------------------------
	SoftwareLoadObj := model.SoftwareLoad                                            {Version:"test value for Version",LoadType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSoftwareLoadRequestResult := dao.CreateSoftwareLoad( SoftwareLoadObj )
	
	if createSoftwareLoadRequestResult.Success == false {
		t.Errorf(createSoftwareLoadRequestResult.Msg)
	} else {
		fmt.Println("Check Create SoftwareLoad success...")
	}
	
	createSoftwareLoadObj,_ := createSoftwareLoadRequestResult.Data. (model.SoftwareLoad)

	// --------------------------------------------------------------
	// Check SoftwareLoad Obj ID
	// --------------------------------------------------------------	
	if createSoftwareLoadObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SoftwareLoad" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSoftwareLoadRequestResult := dao.GetSoftwareLoad( uint64(createSoftwareLoadObj.ID) )
	
	if getSoftwareLoadRequestResult.Success == false {
		t.Errorf(getSoftwareLoadRequestResult.Msg)
	} else {
		fmt.Println("Check Get SoftwareLoad success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSoftwareLoadObj,_ := getSoftwareLoadRequestResult.Data. (model.SoftwareLoad)
	compareSoftwareLoad := cmp.Equal(createSoftwareLoadObj.ID, getSoftwareLoadObj.ID)
	
	if  compareSoftwareLoad == false	{
		t.Errorf( "Created SoftwareLoad object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSoftwareLoadRequestResult := dao.GetAllSoftwareLoad()

	if getAllSoftwareLoadRequestResult.Success == false {
			t.Errorf(getAllSoftwareLoadRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SoftwareLoad success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSoftwareLoadObj []model.SoftwareLoad = getAllSoftwareLoadRequestResult.Data. ([]model.SoftwareLoad)
		
	equalSoftwareLoad := cmp.Equal(createSoftwareLoadObj.ID, getAllSoftwareLoadObj[len(getAllSoftwareLoadObj)-1].ID)
		
	if equalSoftwareLoad == false {
		t.Errorf( "Created object is not equal to the last entry in SoftwareLoad[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SoftwareLoad
	// --------------------------------------------------------------	
	deleteSoftwareLoadRequestResult := dao.DeleteSoftwareLoad(uint64(createSoftwareLoadObj.ID))

	if deleteSoftwareLoadRequestResult.Success == false {
			t.Errorf(deleteSoftwareLoadRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SoftwareLoad success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSoftwareLoadRequestResult = dao.GetSoftwareLoad( uint64(createSoftwareLoadObj.ID) )
	
	if getSoftwareLoadRequestResult.Success == true {
		t.Errorf(getSoftwareLoadRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTypeCertificateCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TypeCertificate
	//----------------------------------------------------------------------------
	TypeCertificateObj := model.TypeCertificate                                                            {CertificateNumber:"test value for CertificateNumber",Authority:"test value for Authority"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTypeCertificateRequestResult := dao.CreateTypeCertificate( TypeCertificateObj )
	
	if createTypeCertificateRequestResult.Success == false {
		t.Errorf(createTypeCertificateRequestResult.Msg)
	} else {
		fmt.Println("Check Create TypeCertificate success...")
	}
	
	createTypeCertificateObj,_ := createTypeCertificateRequestResult.Data. (model.TypeCertificate)

	// --------------------------------------------------------------
	// Check TypeCertificate Obj ID
	// --------------------------------------------------------------	
	if createTypeCertificateObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TypeCertificate" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTypeCertificateRequestResult := dao.GetTypeCertificate( uint64(createTypeCertificateObj.ID) )
	
	if getTypeCertificateRequestResult.Success == false {
		t.Errorf(getTypeCertificateRequestResult.Msg)
	} else {
		fmt.Println("Check Get TypeCertificate success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTypeCertificateObj,_ := getTypeCertificateRequestResult.Data. (model.TypeCertificate)
	compareTypeCertificate := cmp.Equal(createTypeCertificateObj.ID, getTypeCertificateObj.ID)
	
	if  compareTypeCertificate == false	{
		t.Errorf( "Created TypeCertificate object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTypeCertificateRequestResult := dao.GetAllTypeCertificate()

	if getAllTypeCertificateRequestResult.Success == false {
			t.Errorf(getAllTypeCertificateRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TypeCertificate success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTypeCertificateObj []model.TypeCertificate = getAllTypeCertificateRequestResult.Data. ([]model.TypeCertificate)
		
	equalTypeCertificate := cmp.Equal(createTypeCertificateObj.ID, getAllTypeCertificateObj[len(getAllTypeCertificateObj)-1].ID)
		
	if equalTypeCertificate == false {
		t.Errorf( "Created object is not equal to the last entry in TypeCertificate[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TypeCertificate
	// --------------------------------------------------------------	
	deleteTypeCertificateRequestResult := dao.DeleteTypeCertificate(uint64(createTypeCertificateObj.ID))

	if deleteTypeCertificateRequestResult.Success == false {
			t.Errorf(deleteTypeCertificateRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TypeCertificate success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTypeCertificateRequestResult = dao.GetTypeCertificate( uint64(createTypeCertificateObj.ID) )
	
	if getTypeCertificateRequestResult.Success == true {
		t.Errorf(getTypeCertificateRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestProductionCertificateCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ProductionCertificate
	//----------------------------------------------------------------------------
	ProductionCertificateObj := model.ProductionCertificate                                                            {CertificateNumber:"test value for CertificateNumber",Authority:"test value for Authority"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createProductionCertificateRequestResult := dao.CreateProductionCertificate( ProductionCertificateObj )
	
	if createProductionCertificateRequestResult.Success == false {
		t.Errorf(createProductionCertificateRequestResult.Msg)
	} else {
		fmt.Println("Check Create ProductionCertificate success...")
	}
	
	createProductionCertificateObj,_ := createProductionCertificateRequestResult.Data. (model.ProductionCertificate)

	// --------------------------------------------------------------
	// Check ProductionCertificate Obj ID
	// --------------------------------------------------------------	
	if createProductionCertificateObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ProductionCertificate" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getProductionCertificateRequestResult := dao.GetProductionCertificate( uint64(createProductionCertificateObj.ID) )
	
	if getProductionCertificateRequestResult.Success == false {
		t.Errorf(getProductionCertificateRequestResult.Msg)
	} else {
		fmt.Println("Check Get ProductionCertificate success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getProductionCertificateObj,_ := getProductionCertificateRequestResult.Data. (model.ProductionCertificate)
	compareProductionCertificate := cmp.Equal(createProductionCertificateObj.ID, getProductionCertificateObj.ID)
	
	if  compareProductionCertificate == false	{
		t.Errorf( "Created ProductionCertificate object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllProductionCertificateRequestResult := dao.GetAllProductionCertificate()

	if getAllProductionCertificateRequestResult.Success == false {
			t.Errorf(getAllProductionCertificateRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ProductionCertificate success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllProductionCertificateObj []model.ProductionCertificate = getAllProductionCertificateRequestResult.Data. ([]model.ProductionCertificate)
		
	equalProductionCertificate := cmp.Equal(createProductionCertificateObj.ID, getAllProductionCertificateObj[len(getAllProductionCertificateObj)-1].ID)
		
	if equalProductionCertificate == false {
		t.Errorf( "Created object is not equal to the last entry in ProductionCertificate[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ProductionCertificate
	// --------------------------------------------------------------	
	deleteProductionCertificateRequestResult := dao.DeleteProductionCertificate(uint64(createProductionCertificateObj.ID))

	if deleteProductionCertificateRequestResult.Success == false {
			t.Errorf(deleteProductionCertificateRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ProductionCertificate success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getProductionCertificateRequestResult = dao.GetProductionCertificate( uint64(createProductionCertificateObj.ID) )
	
	if getProductionCertificateRequestResult.Success == true {
		t.Errorf(getProductionCertificateRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSalesRegionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SalesRegion
	//----------------------------------------------------------------------------
	SalesRegionObj := model.SalesRegion                                                            {Name:"test value for Name",RegionCode:"test value for RegionCode"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSalesRegionRequestResult := dao.CreateSalesRegion( SalesRegionObj )
	
	if createSalesRegionRequestResult.Success == false {
		t.Errorf(createSalesRegionRequestResult.Msg)
	} else {
		fmt.Println("Check Create SalesRegion success...")
	}
	
	createSalesRegionObj,_ := createSalesRegionRequestResult.Data. (model.SalesRegion)

	// --------------------------------------------------------------
	// Check SalesRegion Obj ID
	// --------------------------------------------------------------	
	if createSalesRegionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SalesRegion" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSalesRegionRequestResult := dao.GetSalesRegion( uint64(createSalesRegionObj.ID) )
	
	if getSalesRegionRequestResult.Success == false {
		t.Errorf(getSalesRegionRequestResult.Msg)
	} else {
		fmt.Println("Check Get SalesRegion success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSalesRegionObj,_ := getSalesRegionRequestResult.Data. (model.SalesRegion)
	compareSalesRegion := cmp.Equal(createSalesRegionObj.ID, getSalesRegionObj.ID)
	
	if  compareSalesRegion == false	{
		t.Errorf( "Created SalesRegion object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSalesRegionRequestResult := dao.GetAllSalesRegion()

	if getAllSalesRegionRequestResult.Success == false {
			t.Errorf(getAllSalesRegionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SalesRegion success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSalesRegionObj []model.SalesRegion = getAllSalesRegionRequestResult.Data. ([]model.SalesRegion)
		
	equalSalesRegion := cmp.Equal(createSalesRegionObj.ID, getAllSalesRegionObj[len(getAllSalesRegionObj)-1].ID)
		
	if equalSalesRegion == false {
		t.Errorf( "Created object is not equal to the last entry in SalesRegion[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SalesRegion
	// --------------------------------------------------------------	
	deleteSalesRegionRequestResult := dao.DeleteSalesRegion(uint64(createSalesRegionObj.ID))

	if deleteSalesRegionRequestResult.Success == false {
			t.Errorf(deleteSalesRegionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SalesRegion success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSalesRegionRequestResult = dao.GetSalesRegion( uint64(createSalesRegionObj.ID) )
	
	if getSalesRegionRequestResult.Success == true {
		t.Errorf(getSalesRegionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSalesCampaignCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SalesCampaign
	//----------------------------------------------------------------------------
	SalesCampaignObj := model.SalesCampaign                                            {CampaignCode:"test value for CampaignCode",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSalesCampaignRequestResult := dao.CreateSalesCampaign( SalesCampaignObj )
	
	if createSalesCampaignRequestResult.Success == false {
		t.Errorf(createSalesCampaignRequestResult.Msg)
	} else {
		fmt.Println("Check Create SalesCampaign success...")
	}
	
	createSalesCampaignObj,_ := createSalesCampaignRequestResult.Data. (model.SalesCampaign)

	// --------------------------------------------------------------
	// Check SalesCampaign Obj ID
	// --------------------------------------------------------------	
	if createSalesCampaignObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SalesCampaign" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSalesCampaignRequestResult := dao.GetSalesCampaign( uint64(createSalesCampaignObj.ID) )
	
	if getSalesCampaignRequestResult.Success == false {
		t.Errorf(getSalesCampaignRequestResult.Msg)
	} else {
		fmt.Println("Check Get SalesCampaign success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSalesCampaignObj,_ := getSalesCampaignRequestResult.Data. (model.SalesCampaign)
	compareSalesCampaign := cmp.Equal(createSalesCampaignObj.ID, getSalesCampaignObj.ID)
	
	if  compareSalesCampaign == false	{
		t.Errorf( "Created SalesCampaign object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSalesCampaignRequestResult := dao.GetAllSalesCampaign()

	if getAllSalesCampaignRequestResult.Success == false {
			t.Errorf(getAllSalesCampaignRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SalesCampaign success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSalesCampaignObj []model.SalesCampaign = getAllSalesCampaignRequestResult.Data. ([]model.SalesCampaign)
		
	equalSalesCampaign := cmp.Equal(createSalesCampaignObj.ID, getAllSalesCampaignObj[len(getAllSalesCampaignObj)-1].ID)
		
	if equalSalesCampaign == false {
		t.Errorf( "Created object is not equal to the last entry in SalesCampaign[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SalesCampaign
	// --------------------------------------------------------------	
	deleteSalesCampaignRequestResult := dao.DeleteSalesCampaign(uint64(createSalesCampaignObj.ID))

	if deleteSalesCampaignRequestResult.Success == false {
			t.Errorf(deleteSalesCampaignRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SalesCampaign success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSalesCampaignRequestResult = dao.GetSalesCampaign( uint64(createSalesCampaignObj.ID) )
	
	if getSalesCampaignRequestResult.Success == true {
		t.Errorf(getSalesCampaignRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

