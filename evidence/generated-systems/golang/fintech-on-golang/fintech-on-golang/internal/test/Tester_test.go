package test

import ( 
	"testing"
    dao "fintech-on-golang/internal/dao"
	"fintech-on-golang/internal/model"
	"fintech-on-golang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestFinancialInstitutionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FinancialInstitution
	//----------------------------------------------------------------------------
	FinancialInstitutionObj := model.FinancialInstitution                                                                                                                                            {Name:"test value for Name",LegalName:"test value for LegalName",CountryOfIncorporation:"test value for CountryOfIncorporation",Bic:new BIC(),Website:"test value for Website"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFinancialInstitutionRequestResult := dao.CreateFinancialInstitution( FinancialInstitutionObj )
	
	if createFinancialInstitutionRequestResult.Success == false {
		t.Errorf(createFinancialInstitutionRequestResult.Msg)
	} else {
		fmt.Println("Check Create FinancialInstitution success...")
	}
	
	createFinancialInstitutionObj,_ := createFinancialInstitutionRequestResult.Data. (model.FinancialInstitution)

	// --------------------------------------------------------------
	// Check FinancialInstitution Obj ID
	// --------------------------------------------------------------	
	if createFinancialInstitutionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for FinancialInstitution" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFinancialInstitutionRequestResult := dao.GetFinancialInstitution( uint64(createFinancialInstitutionObj.ID) )
	
	if getFinancialInstitutionRequestResult.Success == false {
		t.Errorf(getFinancialInstitutionRequestResult.Msg)
	} else {
		fmt.Println("Check Get FinancialInstitution success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFinancialInstitutionObj,_ := getFinancialInstitutionRequestResult.Data. (model.FinancialInstitution)
	compareFinancialInstitution := cmp.Equal(createFinancialInstitutionObj.ID, getFinancialInstitutionObj.ID)
	
	if  compareFinancialInstitution == false	{
		t.Errorf( "Created FinancialInstitution object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFinancialInstitutionRequestResult := dao.GetAllFinancialInstitution()

	if getAllFinancialInstitutionRequestResult.Success == false {
			t.Errorf(getAllFinancialInstitutionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FinancialInstitution success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFinancialInstitutionObj []model.FinancialInstitution = getAllFinancialInstitutionRequestResult.Data. ([]model.FinancialInstitution)
		
	equalFinancialInstitution := cmp.Equal(createFinancialInstitutionObj.ID, getAllFinancialInstitutionObj[len(getAllFinancialInstitutionObj)-1].ID)
		
	if equalFinancialInstitution == false {
		t.Errorf( "Created object is not equal to the last entry in FinancialInstitution[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FinancialInstitution
	// --------------------------------------------------------------	
	deleteFinancialInstitutionRequestResult := dao.DeleteFinancialInstitution(uint64(createFinancialInstitutionObj.ID))

	if deleteFinancialInstitutionRequestResult.Success == false {
			t.Errorf(deleteFinancialInstitutionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FinancialInstitution success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFinancialInstitutionRequestResult = dao.GetFinancialInstitution( uint64(createFinancialInstitutionObj.ID) )
	
	if getFinancialInstitutionRequestResult.Success == true {
		t.Errorf(getFinancialInstitutionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBranchCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Branch
	//----------------------------------------------------------------------------
	BranchObj := model.Branch                                                                            {Name:"test value for Name",BranchCode:"test value for BranchCode",Address:new Address()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBranchRequestResult := dao.CreateBranch( BranchObj )
	
	if createBranchRequestResult.Success == false {
		t.Errorf(createBranchRequestResult.Msg)
	} else {
		fmt.Println("Check Create Branch success...")
	}
	
	createBranchObj,_ := createBranchRequestResult.Data. (model.Branch)

	// --------------------------------------------------------------
	// Check Branch Obj ID
	// --------------------------------------------------------------	
	if createBranchObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Branch" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBranchRequestResult := dao.GetBranch( uint64(createBranchObj.ID) )
	
	if getBranchRequestResult.Success == false {
		t.Errorf(getBranchRequestResult.Msg)
	} else {
		fmt.Println("Check Get Branch success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBranchObj,_ := getBranchRequestResult.Data. (model.Branch)
	compareBranch := cmp.Equal(createBranchObj.ID, getBranchObj.ID)
	
	if  compareBranch == false	{
		t.Errorf( "Created Branch object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBranchRequestResult := dao.GetAllBranch()

	if getAllBranchRequestResult.Success == false {
			t.Errorf(getAllBranchRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Branch success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBranchObj []model.Branch = getAllBranchRequestResult.Data. ([]model.Branch)
		
	equalBranch := cmp.Equal(createBranchObj.ID, getAllBranchObj[len(getAllBranchObj)-1].ID)
		
	if equalBranch == false {
		t.Errorf( "Created object is not equal to the last entry in Branch[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Branch
	// --------------------------------------------------------------	
	deleteBranchRequestResult := dao.DeleteBranch(uint64(createBranchObj.ID))

	if deleteBranchRequestResult.Success == false {
			t.Errorf(deleteBranchRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Branch success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBranchRequestResult = dao.GetBranch( uint64(createBranchObj.ID) )
	
	if getBranchRequestResult.Success == true {
		t.Errorf(getBranchRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestProductOfferingCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ProductOffering
	//----------------------------------------------------------------------------
	ProductOfferingObj := model.ProductOffering                                                                            {Name:"test value for Name",ProductCode:"test value for ProductCode",Category:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createProductOfferingRequestResult := dao.CreateProductOffering( ProductOfferingObj )
	
	if createProductOfferingRequestResult.Success == false {
		t.Errorf(createProductOfferingRequestResult.Msg)
	} else {
		fmt.Println("Check Create ProductOffering success...")
	}
	
	createProductOfferingObj,_ := createProductOfferingRequestResult.Data. (model.ProductOffering)

	// --------------------------------------------------------------
	// Check ProductOffering Obj ID
	// --------------------------------------------------------------	
	if createProductOfferingObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ProductOffering" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getProductOfferingRequestResult := dao.GetProductOffering( uint64(createProductOfferingObj.ID) )
	
	if getProductOfferingRequestResult.Success == false {
		t.Errorf(getProductOfferingRequestResult.Msg)
	} else {
		fmt.Println("Check Get ProductOffering success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getProductOfferingObj,_ := getProductOfferingRequestResult.Data. (model.ProductOffering)
	compareProductOffering := cmp.Equal(createProductOfferingObj.ID, getProductOfferingObj.ID)
	
	if  compareProductOffering == false	{
		t.Errorf( "Created ProductOffering object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllProductOfferingRequestResult := dao.GetAllProductOffering()

	if getAllProductOfferingRequestResult.Success == false {
			t.Errorf(getAllProductOfferingRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ProductOffering success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllProductOfferingObj []model.ProductOffering = getAllProductOfferingRequestResult.Data. ([]model.ProductOffering)
		
	equalProductOffering := cmp.Equal(createProductOfferingObj.ID, getAllProductOfferingObj[len(getAllProductOfferingObj)-1].ID)
		
	if equalProductOffering == false {
		t.Errorf( "Created object is not equal to the last entry in ProductOffering[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ProductOffering
	// --------------------------------------------------------------	
	deleteProductOfferingRequestResult := dao.DeleteProductOffering(uint64(createProductOfferingObj.ID))

	if deleteProductOfferingRequestResult.Success == false {
			t.Errorf(deleteProductOfferingRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ProductOffering success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getProductOfferingRequestResult = dao.GetProductOffering( uint64(createProductOfferingObj.ID) )
	
	if getProductOfferingRequestResult.Success == true {
		t.Errorf(getProductOfferingRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPricingPlanCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PricingPlan
	//----------------------------------------------------------------------------
	PricingPlanObj := model.PricingPlan                                                                                                            {Name:"test value for Name",PlanCode:"test value for PlanCode",BaseCurrency:"test value for BaseCurrency",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPricingPlanRequestResult := dao.CreatePricingPlan( PricingPlanObj )
	
	if createPricingPlanRequestResult.Success == false {
		t.Errorf(createPricingPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Create PricingPlan success...")
	}
	
	createPricingPlanObj,_ := createPricingPlanRequestResult.Data. (model.PricingPlan)

	// --------------------------------------------------------------
	// Check PricingPlan Obj ID
	// --------------------------------------------------------------	
	if createPricingPlanObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PricingPlan" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPricingPlanRequestResult := dao.GetPricingPlan( uint64(createPricingPlanObj.ID) )
	
	if getPricingPlanRequestResult.Success == false {
		t.Errorf(getPricingPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Get PricingPlan success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPricingPlanObj,_ := getPricingPlanRequestResult.Data. (model.PricingPlan)
	comparePricingPlan := cmp.Equal(createPricingPlanObj.ID, getPricingPlanObj.ID)
	
	if  comparePricingPlan == false	{
		t.Errorf( "Created PricingPlan object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPricingPlanRequestResult := dao.GetAllPricingPlan()

	if getAllPricingPlanRequestResult.Success == false {
			t.Errorf(getAllPricingPlanRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PricingPlan success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPricingPlanObj []model.PricingPlan = getAllPricingPlanRequestResult.Data. ([]model.PricingPlan)
		
	equalPricingPlan := cmp.Equal(createPricingPlanObj.ID, getAllPricingPlanObj[len(getAllPricingPlanObj)-1].ID)
		
	if equalPricingPlan == false {
		t.Errorf( "Created object is not equal to the last entry in PricingPlan[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PricingPlan
	// --------------------------------------------------------------	
	deletePricingPlanRequestResult := dao.DeletePricingPlan(uint64(createPricingPlanObj.ID))

	if deletePricingPlanRequestResult.Success == false {
			t.Errorf(deletePricingPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PricingPlan success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPricingPlanRequestResult = dao.GetPricingPlan( uint64(createPricingPlanObj.ID) )
	
	if getPricingPlanRequestResult.Success == true {
		t.Errorf(getPricingPlanRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFeeScheduleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FeeSchedule
	//----------------------------------------------------------------------------
	FeeScheduleObj := model.FeeSchedule                                                                                                                                                                    {Name:"test value for Name",Amount:new Money(),Percentage:"test value",Minimum:new Money(),Maximum:new Money(),FeeType:0,CalculationMethod:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFeeScheduleRequestResult := dao.CreateFeeSchedule( FeeScheduleObj )
	
	if createFeeScheduleRequestResult.Success == false {
		t.Errorf(createFeeScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Create FeeSchedule success...")
	}
	
	createFeeScheduleObj,_ := createFeeScheduleRequestResult.Data. (model.FeeSchedule)

	// --------------------------------------------------------------
	// Check FeeSchedule Obj ID
	// --------------------------------------------------------------	
	if createFeeScheduleObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for FeeSchedule" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFeeScheduleRequestResult := dao.GetFeeSchedule( uint64(createFeeScheduleObj.ID) )
	
	if getFeeScheduleRequestResult.Success == false {
		t.Errorf(getFeeScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Get FeeSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFeeScheduleObj,_ := getFeeScheduleRequestResult.Data. (model.FeeSchedule)
	compareFeeSchedule := cmp.Equal(createFeeScheduleObj.ID, getFeeScheduleObj.ID)
	
	if  compareFeeSchedule == false	{
		t.Errorf( "Created FeeSchedule object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFeeScheduleRequestResult := dao.GetAllFeeSchedule()

	if getAllFeeScheduleRequestResult.Success == false {
			t.Errorf(getAllFeeScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FeeSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFeeScheduleObj []model.FeeSchedule = getAllFeeScheduleRequestResult.Data. ([]model.FeeSchedule)
		
	equalFeeSchedule := cmp.Equal(createFeeScheduleObj.ID, getAllFeeScheduleObj[len(getAllFeeScheduleObj)-1].ID)
		
	if equalFeeSchedule == false {
		t.Errorf( "Created object is not equal to the last entry in FeeSchedule[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FeeSchedule
	// --------------------------------------------------------------	
	deleteFeeScheduleRequestResult := dao.DeleteFeeSchedule(uint64(createFeeScheduleObj.ID))

	if deleteFeeScheduleRequestResult.Success == false {
			t.Errorf(deleteFeeScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FeeSchedule success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFeeScheduleRequestResult = dao.GetFeeSchedule( uint64(createFeeScheduleObj.ID) )
	
	if getFeeScheduleRequestResult.Success == true {
		t.Errorf(getFeeScheduleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestUsageLimitCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for UsageLimit
	//----------------------------------------------------------------------------
	UsageLimitObj := model.UsageLimit                                                                                                            {Name:"test value for Name",Amount:new Money(),Count:100,Scope:0,Period:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createUsageLimitRequestResult := dao.CreateUsageLimit( UsageLimitObj )
	
	if createUsageLimitRequestResult.Success == false {
		t.Errorf(createUsageLimitRequestResult.Msg)
	} else {
		fmt.Println("Check Create UsageLimit success...")
	}
	
	createUsageLimitObj,_ := createUsageLimitRequestResult.Data. (model.UsageLimit)

	// --------------------------------------------------------------
	// Check UsageLimit Obj ID
	// --------------------------------------------------------------	
	if createUsageLimitObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for UsageLimit" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getUsageLimitRequestResult := dao.GetUsageLimit( uint64(createUsageLimitObj.ID) )
	
	if getUsageLimitRequestResult.Success == false {
		t.Errorf(getUsageLimitRequestResult.Msg)
	} else {
		fmt.Println("Check Get UsageLimit success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getUsageLimitObj,_ := getUsageLimitRequestResult.Data. (model.UsageLimit)
	compareUsageLimit := cmp.Equal(createUsageLimitObj.ID, getUsageLimitObj.ID)
	
	if  compareUsageLimit == false	{
		t.Errorf( "Created UsageLimit object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllUsageLimitRequestResult := dao.GetAllUsageLimit()

	if getAllUsageLimitRequestResult.Success == false {
			t.Errorf(getAllUsageLimitRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll UsageLimit success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllUsageLimitObj []model.UsageLimit = getAllUsageLimitRequestResult.Data. ([]model.UsageLimit)
		
	equalUsageLimit := cmp.Equal(createUsageLimitObj.ID, getAllUsageLimitObj[len(getAllUsageLimitObj)-1].ID)
		
	if equalUsageLimit == false {
		t.Errorf( "Created object is not equal to the last entry in UsageLimit[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for UsageLimit
	// --------------------------------------------------------------	
	deleteUsageLimitRequestResult := dao.DeleteUsageLimit(uint64(createUsageLimitObj.ID))

	if deleteUsageLimitRequestResult.Success == false {
			t.Errorf(deleteUsageLimitRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion UsageLimit success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getUsageLimitRequestResult = dao.GetUsageLimit( uint64(createUsageLimitObj.ID) )
	
	if getUsageLimitRequestResult.Success == true {
		t.Errorf(getUsageLimitRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCustomerCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Customer
	//----------------------------------------------------------------------------
	CustomerObj := model.Customer                                                                                                                                                                                                                    {FirstName:"test value for FirstName",LastName:"test value for LastName",DateOfBirth:time.Now(),Email:new Email(),Phone:new PhoneNumber(),Address:new Address(),TaxId:new TaxId(),RiskScore:new RiskScore(),CustomerType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCustomerRequestResult := dao.CreateCustomer( CustomerObj )
	
	if createCustomerRequestResult.Success == false {
		t.Errorf(createCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check Create Customer success...")
	}
	
	createCustomerObj,_ := createCustomerRequestResult.Data. (model.Customer)

	// --------------------------------------------------------------
	// Check Customer Obj ID
	// --------------------------------------------------------------	
	if createCustomerObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Customer" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCustomerRequestResult := dao.GetCustomer( uint64(createCustomerObj.ID) )
	
	if getCustomerRequestResult.Success == false {
		t.Errorf(getCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check Get Customer success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCustomerObj,_ := getCustomerRequestResult.Data. (model.Customer)
	compareCustomer := cmp.Equal(createCustomerObj.ID, getCustomerObj.ID)
	
	if  compareCustomer == false	{
		t.Errorf( "Created Customer object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCustomerRequestResult := dao.GetAllCustomer()

	if getAllCustomerRequestResult.Success == false {
			t.Errorf(getAllCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Customer success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCustomerObj []model.Customer = getAllCustomerRequestResult.Data. ([]model.Customer)
		
	equalCustomer := cmp.Equal(createCustomerObj.ID, getAllCustomerObj[len(getAllCustomerObj)-1].ID)
		
	if equalCustomer == false {
		t.Errorf( "Created object is not equal to the last entry in Customer[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Customer
	// --------------------------------------------------------------	
	deleteCustomerRequestResult := dao.DeleteCustomer(uint64(createCustomerObj.ID))

	if deleteCustomerRequestResult.Success == false {
			t.Errorf(deleteCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Customer success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCustomerRequestResult = dao.GetCustomer( uint64(createCustomerObj.ID) )
	
	if getCustomerRequestResult.Success == true {
		t.Errorf(getCustomerRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestKYCProfileCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for KYCProfile
	//----------------------------------------------------------------------------
	KYCProfileObj := model.KYCProfile                                                                            {ProfileId:"test value for ProfileId",CreatedAt:new DateTime(),Status:0,VerificationLevel:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createKYCProfileRequestResult := dao.CreateKYCProfile( KYCProfileObj )
	
	if createKYCProfileRequestResult.Success == false {
		t.Errorf(createKYCProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Create KYCProfile success...")
	}
	
	createKYCProfileObj,_ := createKYCProfileRequestResult.Data. (model.KYCProfile)

	// --------------------------------------------------------------
	// Check KYCProfile Obj ID
	// --------------------------------------------------------------	
	if createKYCProfileObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for KYCProfile" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getKYCProfileRequestResult := dao.GetKYCProfile( uint64(createKYCProfileObj.ID) )
	
	if getKYCProfileRequestResult.Success == false {
		t.Errorf(getKYCProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Get KYCProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getKYCProfileObj,_ := getKYCProfileRequestResult.Data. (model.KYCProfile)
	compareKYCProfile := cmp.Equal(createKYCProfileObj.ID, getKYCProfileObj.ID)
	
	if  compareKYCProfile == false	{
		t.Errorf( "Created KYCProfile object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllKYCProfileRequestResult := dao.GetAllKYCProfile()

	if getAllKYCProfileRequestResult.Success == false {
			t.Errorf(getAllKYCProfileRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll KYCProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllKYCProfileObj []model.KYCProfile = getAllKYCProfileRequestResult.Data. ([]model.KYCProfile)
		
	equalKYCProfile := cmp.Equal(createKYCProfileObj.ID, getAllKYCProfileObj[len(getAllKYCProfileObj)-1].ID)
		
	if equalKYCProfile == false {
		t.Errorf( "Created object is not equal to the last entry in KYCProfile[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for KYCProfile
	// --------------------------------------------------------------	
	deleteKYCProfileRequestResult := dao.DeleteKYCProfile(uint64(createKYCProfileObj.ID))

	if deleteKYCProfileRequestResult.Success == false {
			t.Errorf(deleteKYCProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion KYCProfile success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getKYCProfileRequestResult = dao.GetKYCProfile( uint64(createKYCProfileObj.ID) )
	
	if getKYCProfileRequestResult.Success == true {
		t.Errorf(getKYCProfileRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestKYCDocumentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for KYCDocument
	//----------------------------------------------------------------------------
	KYCDocumentObj := model.KYCDocument                                                                                                                                    {Reference:new DocumentReference(),IssuedCountry:"test value for IssuedCountry",ExpirationDate:time.Now(),DocumentType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createKYCDocumentRequestResult := dao.CreateKYCDocument( KYCDocumentObj )
	
	if createKYCDocumentRequestResult.Success == false {
		t.Errorf(createKYCDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check Create KYCDocument success...")
	}
	
	createKYCDocumentObj,_ := createKYCDocumentRequestResult.Data. (model.KYCDocument)

	// --------------------------------------------------------------
	// Check KYCDocument Obj ID
	// --------------------------------------------------------------	
	if createKYCDocumentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for KYCDocument" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getKYCDocumentRequestResult := dao.GetKYCDocument( uint64(createKYCDocumentObj.ID) )
	
	if getKYCDocumentRequestResult.Success == false {
		t.Errorf(getKYCDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check Get KYCDocument success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getKYCDocumentObj,_ := getKYCDocumentRequestResult.Data. (model.KYCDocument)
	compareKYCDocument := cmp.Equal(createKYCDocumentObj.ID, getKYCDocumentObj.ID)
	
	if  compareKYCDocument == false	{
		t.Errorf( "Created KYCDocument object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllKYCDocumentRequestResult := dao.GetAllKYCDocument()

	if getAllKYCDocumentRequestResult.Success == false {
			t.Errorf(getAllKYCDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll KYCDocument success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllKYCDocumentObj []model.KYCDocument = getAllKYCDocumentRequestResult.Data. ([]model.KYCDocument)
		
	equalKYCDocument := cmp.Equal(createKYCDocumentObj.ID, getAllKYCDocumentObj[len(getAllKYCDocumentObj)-1].ID)
		
	if equalKYCDocument == false {
		t.Errorf( "Created object is not equal to the last entry in KYCDocument[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for KYCDocument
	// --------------------------------------------------------------	
	deleteKYCDocumentRequestResult := dao.DeleteKYCDocument(uint64(createKYCDocumentObj.ID))

	if deleteKYCDocumentRequestResult.Success == false {
			t.Errorf(deleteKYCDocumentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion KYCDocument success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getKYCDocumentRequestResult = dao.GetKYCDocument( uint64(createKYCDocumentObj.ID) )
	
	if getKYCDocumentRequestResult.Success == true {
		t.Errorf(getKYCDocumentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestScreeningCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Screening
	//----------------------------------------------------------------------------
	ScreeningObj := model.Screening                                                            {Score:new RiskScore(),ScreenedAt:new DateTime(),ScreeningType:0,Status:0}

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


func TestVerifiedAddressCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for VerifiedAddress
	//----------------------------------------------------------------------------
	VerifiedAddressObj := model.VerifiedAddress                                            {Address:new Address(),VerifiedAt:new DateTime(),VerificationStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createVerifiedAddressRequestResult := dao.CreateVerifiedAddress( VerifiedAddressObj )
	
	if createVerifiedAddressRequestResult.Success == false {
		t.Errorf(createVerifiedAddressRequestResult.Msg)
	} else {
		fmt.Println("Check Create VerifiedAddress success...")
	}
	
	createVerifiedAddressObj,_ := createVerifiedAddressRequestResult.Data. (model.VerifiedAddress)

	// --------------------------------------------------------------
	// Check VerifiedAddress Obj ID
	// --------------------------------------------------------------	
	if createVerifiedAddressObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for VerifiedAddress" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getVerifiedAddressRequestResult := dao.GetVerifiedAddress( uint64(createVerifiedAddressObj.ID) )
	
	if getVerifiedAddressRequestResult.Success == false {
		t.Errorf(getVerifiedAddressRequestResult.Msg)
	} else {
		fmt.Println("Check Get VerifiedAddress success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getVerifiedAddressObj,_ := getVerifiedAddressRequestResult.Data. (model.VerifiedAddress)
	compareVerifiedAddress := cmp.Equal(createVerifiedAddressObj.ID, getVerifiedAddressObj.ID)
	
	if  compareVerifiedAddress == false	{
		t.Errorf( "Created VerifiedAddress object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllVerifiedAddressRequestResult := dao.GetAllVerifiedAddress()

	if getAllVerifiedAddressRequestResult.Success == false {
			t.Errorf(getAllVerifiedAddressRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll VerifiedAddress success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllVerifiedAddressObj []model.VerifiedAddress = getAllVerifiedAddressRequestResult.Data. ([]model.VerifiedAddress)
		
	equalVerifiedAddress := cmp.Equal(createVerifiedAddressObj.ID, getAllVerifiedAddressObj[len(getAllVerifiedAddressObj)-1].ID)
		
	if equalVerifiedAddress == false {
		t.Errorf( "Created object is not equal to the last entry in VerifiedAddress[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for VerifiedAddress
	// --------------------------------------------------------------	
	deleteVerifiedAddressRequestResult := dao.DeleteVerifiedAddress(uint64(createVerifiedAddressObj.ID))

	if deleteVerifiedAddressRequestResult.Success == false {
			t.Errorf(deleteVerifiedAddressRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion VerifiedAddress success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getVerifiedAddressRequestResult = dao.GetVerifiedAddress( uint64(createVerifiedAddressObj.ID) )
	
	if getVerifiedAddressRequestResult.Success == true {
		t.Errorf(getVerifiedAddressRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCompliancePolicyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CompliancePolicy
	//----------------------------------------------------------------------------
	CompliancePolicyObj := model.CompliancePolicy                                                                                                            {Name:"test value for Name",PolicyCode:"test value for PolicyCode",Description:"test value for Description",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCompliancePolicyRequestResult := dao.CreateCompliancePolicy( CompliancePolicyObj )
	
	if createCompliancePolicyRequestResult.Success == false {
		t.Errorf(createCompliancePolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Create CompliancePolicy success...")
	}
	
	createCompliancePolicyObj,_ := createCompliancePolicyRequestResult.Data. (model.CompliancePolicy)

	// --------------------------------------------------------------
	// Check CompliancePolicy Obj ID
	// --------------------------------------------------------------	
	if createCompliancePolicyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CompliancePolicy" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCompliancePolicyRequestResult := dao.GetCompliancePolicy( uint64(createCompliancePolicyObj.ID) )
	
	if getCompliancePolicyRequestResult.Success == false {
		t.Errorf(getCompliancePolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Get CompliancePolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCompliancePolicyObj,_ := getCompliancePolicyRequestResult.Data. (model.CompliancePolicy)
	compareCompliancePolicy := cmp.Equal(createCompliancePolicyObj.ID, getCompliancePolicyObj.ID)
	
	if  compareCompliancePolicy == false	{
		t.Errorf( "Created CompliancePolicy object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCompliancePolicyRequestResult := dao.GetAllCompliancePolicy()

	if getAllCompliancePolicyRequestResult.Success == false {
			t.Errorf(getAllCompliancePolicyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CompliancePolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCompliancePolicyObj []model.CompliancePolicy = getAllCompliancePolicyRequestResult.Data. ([]model.CompliancePolicy)
		
	equalCompliancePolicy := cmp.Equal(createCompliancePolicyObj.ID, getAllCompliancePolicyObj[len(getAllCompliancePolicyObj)-1].ID)
		
	if equalCompliancePolicy == false {
		t.Errorf( "Created object is not equal to the last entry in CompliancePolicy[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CompliancePolicy
	// --------------------------------------------------------------	
	deleteCompliancePolicyRequestResult := dao.DeleteCompliancePolicy(uint64(createCompliancePolicyObj.ID))

	if deleteCompliancePolicyRequestResult.Success == false {
			t.Errorf(deleteCompliancePolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CompliancePolicy success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCompliancePolicyRequestResult = dao.GetCompliancePolicy( uint64(createCompliancePolicyObj.ID) )
	
	if getCompliancePolicyRequestResult.Success == true {
		t.Errorf(getCompliancePolicyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestComplianceAlertCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ComplianceAlert
	//----------------------------------------------------------------------------
	ComplianceAlertObj := model.ComplianceAlert                                                                                                            {AlertCode:"test value for AlertCode",RaisedAt:new DateTime(),Notes:"test value for Notes",Severity:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createComplianceAlertRequestResult := dao.CreateComplianceAlert( ComplianceAlertObj )
	
	if createComplianceAlertRequestResult.Success == false {
		t.Errorf(createComplianceAlertRequestResult.Msg)
	} else {
		fmt.Println("Check Create ComplianceAlert success...")
	}
	
	createComplianceAlertObj,_ := createComplianceAlertRequestResult.Data. (model.ComplianceAlert)

	// --------------------------------------------------------------
	// Check ComplianceAlert Obj ID
	// --------------------------------------------------------------	
	if createComplianceAlertObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ComplianceAlert" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getComplianceAlertRequestResult := dao.GetComplianceAlert( uint64(createComplianceAlertObj.ID) )
	
	if getComplianceAlertRequestResult.Success == false {
		t.Errorf(getComplianceAlertRequestResult.Msg)
	} else {
		fmt.Println("Check Get ComplianceAlert success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getComplianceAlertObj,_ := getComplianceAlertRequestResult.Data. (model.ComplianceAlert)
	compareComplianceAlert := cmp.Equal(createComplianceAlertObj.ID, getComplianceAlertObj.ID)
	
	if  compareComplianceAlert == false	{
		t.Errorf( "Created ComplianceAlert object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllComplianceAlertRequestResult := dao.GetAllComplianceAlert()

	if getAllComplianceAlertRequestResult.Success == false {
			t.Errorf(getAllComplianceAlertRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ComplianceAlert success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllComplianceAlertObj []model.ComplianceAlert = getAllComplianceAlertRequestResult.Data. ([]model.ComplianceAlert)
		
	equalComplianceAlert := cmp.Equal(createComplianceAlertObj.ID, getAllComplianceAlertObj[len(getAllComplianceAlertObj)-1].ID)
		
	if equalComplianceAlert == false {
		t.Errorf( "Created object is not equal to the last entry in ComplianceAlert[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ComplianceAlert
	// --------------------------------------------------------------	
	deleteComplianceAlertRequestResult := dao.DeleteComplianceAlert(uint64(createComplianceAlertObj.ID))

	if deleteComplianceAlertRequestResult.Success == false {
			t.Errorf(deleteComplianceAlertRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ComplianceAlert success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getComplianceAlertRequestResult = dao.GetComplianceAlert( uint64(createComplianceAlertObj.ID) )
	
	if getComplianceAlertRequestResult.Success == true {
		t.Errorf(getComplianceAlertRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestConsentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Consent
	//----------------------------------------------------------------------------
	ConsentObj := model.Consent                                                                                            {GrantedAt:new DateTime(),ExpiresAt:new DateTime(),Scope:"test value for Scope",ConsentType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createConsentRequestResult := dao.CreateConsent( ConsentObj )
	
	if createConsentRequestResult.Success == false {
		t.Errorf(createConsentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Consent success...")
	}
	
	createConsentObj,_ := createConsentRequestResult.Data. (model.Consent)

	// --------------------------------------------------------------
	// Check Consent Obj ID
	// --------------------------------------------------------------	
	if createConsentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Consent" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getConsentRequestResult := dao.GetConsent( uint64(createConsentObj.ID) )
	
	if getConsentRequestResult.Success == false {
		t.Errorf(getConsentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Consent success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getConsentObj,_ := getConsentRequestResult.Data. (model.Consent)
	compareConsent := cmp.Equal(createConsentObj.ID, getConsentObj.ID)
	
	if  compareConsent == false	{
		t.Errorf( "Created Consent object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllConsentRequestResult := dao.GetAllConsent()

	if getAllConsentRequestResult.Success == false {
			t.Errorf(getAllConsentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Consent success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllConsentObj []model.Consent = getAllConsentRequestResult.Data. ([]model.Consent)
		
	equalConsent := cmp.Equal(createConsentObj.ID, getAllConsentObj[len(getAllConsentObj)-1].ID)
		
	if equalConsent == false {
		t.Errorf( "Created object is not equal to the last entry in Consent[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Consent
	// --------------------------------------------------------------	
	deleteConsentRequestResult := dao.DeleteConsent(uint64(createConsentObj.ID))

	if deleteConsentRequestResult.Success == false {
			t.Errorf(deleteConsentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Consent success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getConsentRequestResult = dao.GetConsent( uint64(createConsentObj.ID) )
	
	if getConsentRequestResult.Success == true {
		t.Errorf(getConsentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAPIClientCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for APIClient
	//----------------------------------------------------------------------------
	APIClientObj := model.APIClient                                                                                                            {Name:"test value for Name",ClientId:"test value for ClientId",RedirectUri:"test value for RedirectUri",ClientType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAPIClientRequestResult := dao.CreateAPIClient( APIClientObj )
	
	if createAPIClientRequestResult.Success == false {
		t.Errorf(createAPIClientRequestResult.Msg)
	} else {
		fmt.Println("Check Create APIClient success...")
	}
	
	createAPIClientObj,_ := createAPIClientRequestResult.Data. (model.APIClient)

	// --------------------------------------------------------------
	// Check APIClient Obj ID
	// --------------------------------------------------------------	
	if createAPIClientObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for APIClient" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAPIClientRequestResult := dao.GetAPIClient( uint64(createAPIClientObj.ID) )
	
	if getAPIClientRequestResult.Success == false {
		t.Errorf(getAPIClientRequestResult.Msg)
	} else {
		fmt.Println("Check Get APIClient success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAPIClientObj,_ := getAPIClientRequestResult.Data. (model.APIClient)
	compareAPIClient := cmp.Equal(createAPIClientObj.ID, getAPIClientObj.ID)
	
	if  compareAPIClient == false	{
		t.Errorf( "Created APIClient object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAPIClientRequestResult := dao.GetAllAPIClient()

	if getAllAPIClientRequestResult.Success == false {
			t.Errorf(getAllAPIClientRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll APIClient success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAPIClientObj []model.APIClient = getAllAPIClientRequestResult.Data. ([]model.APIClient)
		
	equalAPIClient := cmp.Equal(createAPIClientObj.ID, getAllAPIClientObj[len(getAllAPIClientObj)-1].ID)
		
	if equalAPIClient == false {
		t.Errorf( "Created object is not equal to the last entry in APIClient[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for APIClient
	// --------------------------------------------------------------	
	deleteAPIClientRequestResult := dao.DeleteAPIClient(uint64(createAPIClientObj.ID))

	if deleteAPIClientRequestResult.Success == false {
			t.Errorf(deleteAPIClientRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion APIClient success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAPIClientRequestResult = dao.GetAPIClient( uint64(createAPIClientObj.ID) )
	
	if getAPIClientRequestResult.Success == true {
		t.Errorf(getAPIClientRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAgreementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Agreement
	//----------------------------------------------------------------------------
	AgreementObj := model.Agreement                                                                                                                    {AgreementNumber:"test value for AgreementNumber",EffectiveDate:time.Now(),AgreementType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAgreementRequestResult := dao.CreateAgreement( AgreementObj )
	
	if createAgreementRequestResult.Success == false {
		t.Errorf(createAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check Create Agreement success...")
	}
	
	createAgreementObj,_ := createAgreementRequestResult.Data. (model.Agreement)

	// --------------------------------------------------------------
	// Check Agreement Obj ID
	// --------------------------------------------------------------	
	if createAgreementObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Agreement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAgreementRequestResult := dao.GetAgreement( uint64(createAgreementObj.ID) )
	
	if getAgreementRequestResult.Success == false {
		t.Errorf(getAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check Get Agreement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAgreementObj,_ := getAgreementRequestResult.Data. (model.Agreement)
	compareAgreement := cmp.Equal(createAgreementObj.ID, getAgreementObj.ID)
	
	if  compareAgreement == false	{
		t.Errorf( "Created Agreement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAgreementRequestResult := dao.GetAllAgreement()

	if getAllAgreementRequestResult.Success == false {
			t.Errorf(getAllAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Agreement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAgreementObj []model.Agreement = getAllAgreementRequestResult.Data. ([]model.Agreement)
		
	equalAgreement := cmp.Equal(createAgreementObj.ID, getAllAgreementObj[len(getAllAgreementObj)-1].ID)
		
	if equalAgreement == false {
		t.Errorf( "Created object is not equal to the last entry in Agreement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Agreement
	// --------------------------------------------------------------	
	deleteAgreementRequestResult := dao.DeleteAgreement(uint64(createAgreementObj.ID))

	if deleteAgreementRequestResult.Success == false {
			t.Errorf(deleteAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Agreement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAgreementRequestResult = dao.GetAgreement( uint64(createAgreementObj.ID) )
	
	if getAgreementRequestResult.Success == true {
		t.Errorf(getAgreementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAccountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Account
	//----------------------------------------------------------------------------
	AccountObj := model.Account                                                                                                                                                                                                    {AccountNumber:new AccountNumber(),Iban:new IBAN(),Bic:new BIC(),OpenedDate:time.Now(),Currency:"test value for Currency",Balance:new Money(),AvailableBalance:new Money(),AccountType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAccountRequestResult := dao.CreateAccount( AccountObj )
	
	if createAccountRequestResult.Success == false {
		t.Errorf(createAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Create Account success...")
	}
	
	createAccountObj,_ := createAccountRequestResult.Data. (model.Account)

	// --------------------------------------------------------------
	// Check Account Obj ID
	// --------------------------------------------------------------	
	if createAccountObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Account" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAccountRequestResult := dao.GetAccount( uint64(createAccountObj.ID) )
	
	if getAccountRequestResult.Success == false {
		t.Errorf(getAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Get Account success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAccountObj,_ := getAccountRequestResult.Data. (model.Account)
	compareAccount := cmp.Equal(createAccountObj.ID, getAccountObj.ID)
	
	if  compareAccount == false	{
		t.Errorf( "Created Account object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAccountRequestResult := dao.GetAllAccount()

	if getAllAccountRequestResult.Success == false {
			t.Errorf(getAllAccountRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Account success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAccountObj []model.Account = getAllAccountRequestResult.Data. ([]model.Account)
		
	equalAccount := cmp.Equal(createAccountObj.ID, getAllAccountObj[len(getAllAccountObj)-1].ID)
		
	if equalAccount == false {
		t.Errorf( "Created object is not equal to the last entry in Account[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Account
	// --------------------------------------------------------------	
	deleteAccountRequestResult := dao.DeleteAccount(uint64(createAccountObj.ID))

	if deleteAccountRequestResult.Success == false {
			t.Errorf(deleteAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Account success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAccountRequestResult = dao.GetAccount( uint64(createAccountObj.ID) )
	
	if getAccountRequestResult.Success == true {
		t.Errorf(getAccountRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWalletCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Wallet
	//----------------------------------------------------------------------------
	WalletObj := model.Wallet                                                            {Currency:"test value for Currency",Balance:new Money(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWalletRequestResult := dao.CreateWallet( WalletObj )
	
	if createWalletRequestResult.Success == false {
		t.Errorf(createWalletRequestResult.Msg)
	} else {
		fmt.Println("Check Create Wallet success...")
	}
	
	createWalletObj,_ := createWalletRequestResult.Data. (model.Wallet)

	// --------------------------------------------------------------
	// Check Wallet Obj ID
	// --------------------------------------------------------------	
	if createWalletObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Wallet" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWalletRequestResult := dao.GetWallet( uint64(createWalletObj.ID) )
	
	if getWalletRequestResult.Success == false {
		t.Errorf(getWalletRequestResult.Msg)
	} else {
		fmt.Println("Check Get Wallet success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWalletObj,_ := getWalletRequestResult.Data. (model.Wallet)
	compareWallet := cmp.Equal(createWalletObj.ID, getWalletObj.ID)
	
	if  compareWallet == false	{
		t.Errorf( "Created Wallet object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWalletRequestResult := dao.GetAllWallet()

	if getAllWalletRequestResult.Success == false {
			t.Errorf(getAllWalletRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Wallet success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWalletObj []model.Wallet = getAllWalletRequestResult.Data. ([]model.Wallet)
		
	equalWallet := cmp.Equal(createWalletObj.ID, getAllWalletObj[len(getAllWalletObj)-1].ID)
		
	if equalWallet == false {
		t.Errorf( "Created object is not equal to the last entry in Wallet[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Wallet
	// --------------------------------------------------------------	
	deleteWalletRequestResult := dao.DeleteWallet(uint64(createWalletObj.ID))

	if deleteWalletRequestResult.Success == false {
			t.Errorf(deleteWalletRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Wallet success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWalletRequestResult = dao.GetWallet( uint64(createWalletObj.ID) )
	
	if getWalletRequestResult.Success == true {
		t.Errorf(getWalletRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPaymentCardCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PaymentCard
	//----------------------------------------------------------------------------
	PaymentCardObj := model.PaymentCard                                                                                                                                                                            {CardToken:new CardNumberToken(),MaskedPan:"test value for MaskedPan",ExpiryMonth:100,ExpiryYear:100,CardholderName:"test value for CardholderName",Scheme:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPaymentCardRequestResult := dao.CreatePaymentCard( PaymentCardObj )
	
	if createPaymentCardRequestResult.Success == false {
		t.Errorf(createPaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Check Create PaymentCard success...")
	}
	
	createPaymentCardObj,_ := createPaymentCardRequestResult.Data. (model.PaymentCard)

	// --------------------------------------------------------------
	// Check PaymentCard Obj ID
	// --------------------------------------------------------------	
	if createPaymentCardObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PaymentCard" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPaymentCardRequestResult := dao.GetPaymentCard( uint64(createPaymentCardObj.ID) )
	
	if getPaymentCardRequestResult.Success == false {
		t.Errorf(getPaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Check Get PaymentCard success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPaymentCardObj,_ := getPaymentCardRequestResult.Data. (model.PaymentCard)
	comparePaymentCard := cmp.Equal(createPaymentCardObj.ID, getPaymentCardObj.ID)
	
	if  comparePaymentCard == false	{
		t.Errorf( "Created PaymentCard object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPaymentCardRequestResult := dao.GetAllPaymentCard()

	if getAllPaymentCardRequestResult.Success == false {
			t.Errorf(getAllPaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PaymentCard success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPaymentCardObj []model.PaymentCard = getAllPaymentCardRequestResult.Data. ([]model.PaymentCard)
		
	equalPaymentCard := cmp.Equal(createPaymentCardObj.ID, getAllPaymentCardObj[len(getAllPaymentCardObj)-1].ID)
		
	if equalPaymentCard == false {
		t.Errorf( "Created object is not equal to the last entry in PaymentCard[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PaymentCard
	// --------------------------------------------------------------	
	deletePaymentCardRequestResult := dao.DeletePaymentCard(uint64(createPaymentCardObj.ID))

	if deletePaymentCardRequestResult.Success == false {
			t.Errorf(deletePaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PaymentCard success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPaymentCardRequestResult = dao.GetPaymentCard( uint64(createPaymentCardObj.ID) )
	
	if getPaymentCardRequestResult.Success == true {
		t.Errorf(getPaymentCardRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCardTokenizationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CardTokenization
	//----------------------------------------------------------------------------
	CardTokenizationObj := model.CardTokenization                                                                            {TokenReference:"test value for TokenReference",CreatedAt:new DateTime(),WalletProvider:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCardTokenizationRequestResult := dao.CreateCardTokenization( CardTokenizationObj )
	
	if createCardTokenizationRequestResult.Success == false {
		t.Errorf(createCardTokenizationRequestResult.Msg)
	} else {
		fmt.Println("Check Create CardTokenization success...")
	}
	
	createCardTokenizationObj,_ := createCardTokenizationRequestResult.Data. (model.CardTokenization)

	// --------------------------------------------------------------
	// Check CardTokenization Obj ID
	// --------------------------------------------------------------	
	if createCardTokenizationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CardTokenization" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCardTokenizationRequestResult := dao.GetCardTokenization( uint64(createCardTokenizationObj.ID) )
	
	if getCardTokenizationRequestResult.Success == false {
		t.Errorf(getCardTokenizationRequestResult.Msg)
	} else {
		fmt.Println("Check Get CardTokenization success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCardTokenizationObj,_ := getCardTokenizationRequestResult.Data. (model.CardTokenization)
	compareCardTokenization := cmp.Equal(createCardTokenizationObj.ID, getCardTokenizationObj.ID)
	
	if  compareCardTokenization == false	{
		t.Errorf( "Created CardTokenization object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCardTokenizationRequestResult := dao.GetAllCardTokenization()

	if getAllCardTokenizationRequestResult.Success == false {
			t.Errorf(getAllCardTokenizationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CardTokenization success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCardTokenizationObj []model.CardTokenization = getAllCardTokenizationRequestResult.Data. ([]model.CardTokenization)
		
	equalCardTokenization := cmp.Equal(createCardTokenizationObj.ID, getAllCardTokenizationObj[len(getAllCardTokenizationObj)-1].ID)
		
	if equalCardTokenization == false {
		t.Errorf( "Created object is not equal to the last entry in CardTokenization[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CardTokenization
	// --------------------------------------------------------------	
	deleteCardTokenizationRequestResult := dao.DeleteCardTokenization(uint64(createCardTokenizationObj.ID))

	if deleteCardTokenizationRequestResult.Success == false {
			t.Errorf(deleteCardTokenizationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CardTokenization success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCardTokenizationRequestResult = dao.GetCardTokenization( uint64(createCardTokenizationObj.ID) )
	
	if getCardTokenizationRequestResult.Success == true {
		t.Errorf(getCardTokenizationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMerchantCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Merchant
	//----------------------------------------------------------------------------
	MerchantObj := model.Merchant                                                                                                                                                            {Name:"test value for Name",Mcc:"test value for Mcc",Url:"test value for Url",Country:"test value for Country",SettlementCurrency:"test value for SettlementCurrency"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMerchantRequestResult := dao.CreateMerchant( MerchantObj )
	
	if createMerchantRequestResult.Success == false {
		t.Errorf(createMerchantRequestResult.Msg)
	} else {
		fmt.Println("Check Create Merchant success...")
	}
	
	createMerchantObj,_ := createMerchantRequestResult.Data. (model.Merchant)

	// --------------------------------------------------------------
	// Check Merchant Obj ID
	// --------------------------------------------------------------	
	if createMerchantObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Merchant" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMerchantRequestResult := dao.GetMerchant( uint64(createMerchantObj.ID) )
	
	if getMerchantRequestResult.Success == false {
		t.Errorf(getMerchantRequestResult.Msg)
	} else {
		fmt.Println("Check Get Merchant success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMerchantObj,_ := getMerchantRequestResult.Data. (model.Merchant)
	compareMerchant := cmp.Equal(createMerchantObj.ID, getMerchantObj.ID)
	
	if  compareMerchant == false	{
		t.Errorf( "Created Merchant object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMerchantRequestResult := dao.GetAllMerchant()

	if getAllMerchantRequestResult.Success == false {
			t.Errorf(getAllMerchantRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Merchant success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMerchantObj []model.Merchant = getAllMerchantRequestResult.Data. ([]model.Merchant)
		
	equalMerchant := cmp.Equal(createMerchantObj.ID, getAllMerchantObj[len(getAllMerchantObj)-1].ID)
		
	if equalMerchant == false {
		t.Errorf( "Created object is not equal to the last entry in Merchant[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Merchant
	// --------------------------------------------------------------	
	deleteMerchantRequestResult := dao.DeleteMerchant(uint64(createMerchantObj.ID))

	if deleteMerchantRequestResult.Success == false {
			t.Errorf(deleteMerchantRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Merchant success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMerchantRequestResult = dao.GetMerchant( uint64(createMerchantObj.ID) )
	
	if getMerchantRequestResult.Success == true {
		t.Errorf(getMerchantRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTerminalCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Terminal
	//----------------------------------------------------------------------------
	TerminalObj := model.Terminal                                            {Location:new Address(),Type:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTerminalRequestResult := dao.CreateTerminal( TerminalObj )
	
	if createTerminalRequestResult.Success == false {
		t.Errorf(createTerminalRequestResult.Msg)
	} else {
		fmt.Println("Check Create Terminal success...")
	}
	
	createTerminalObj,_ := createTerminalRequestResult.Data. (model.Terminal)

	// --------------------------------------------------------------
	// Check Terminal Obj ID
	// --------------------------------------------------------------	
	if createTerminalObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Terminal" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTerminalRequestResult := dao.GetTerminal( uint64(createTerminalObj.ID) )
	
	if getTerminalRequestResult.Success == false {
		t.Errorf(getTerminalRequestResult.Msg)
	} else {
		fmt.Println("Check Get Terminal success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTerminalObj,_ := getTerminalRequestResult.Data. (model.Terminal)
	compareTerminal := cmp.Equal(createTerminalObj.ID, getTerminalObj.ID)
	
	if  compareTerminal == false	{
		t.Errorf( "Created Terminal object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTerminalRequestResult := dao.GetAllTerminal()

	if getAllTerminalRequestResult.Success == false {
			t.Errorf(getAllTerminalRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Terminal success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTerminalObj []model.Terminal = getAllTerminalRequestResult.Data. ([]model.Terminal)
		
	equalTerminal := cmp.Equal(createTerminalObj.ID, getAllTerminalObj[len(getAllTerminalObj)-1].ID)
		
	if equalTerminal == false {
		t.Errorf( "Created object is not equal to the last entry in Terminal[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Terminal
	// --------------------------------------------------------------	
	deleteTerminalRequestResult := dao.DeleteTerminal(uint64(createTerminalObj.ID))

	if deleteTerminalRequestResult.Success == false {
			t.Errorf(deleteTerminalRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Terminal success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTerminalRequestResult = dao.GetTerminal( uint64(createTerminalObj.ID) )
	
	if getTerminalRequestResult.Success == true {
		t.Errorf(getTerminalRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPaymentContractCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PaymentContract
	//----------------------------------------------------------------------------
	PaymentContractObj := model.PaymentContract                                                                            {ContractNumber:"test value for ContractNumber",PricingPlanCode:"test value for PricingPlanCode",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPaymentContractRequestResult := dao.CreatePaymentContract( PaymentContractObj )
	
	if createPaymentContractRequestResult.Success == false {
		t.Errorf(createPaymentContractRequestResult.Msg)
	} else {
		fmt.Println("Check Create PaymentContract success...")
	}
	
	createPaymentContractObj,_ := createPaymentContractRequestResult.Data. (model.PaymentContract)

	// --------------------------------------------------------------
	// Check PaymentContract Obj ID
	// --------------------------------------------------------------	
	if createPaymentContractObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PaymentContract" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPaymentContractRequestResult := dao.GetPaymentContract( uint64(createPaymentContractObj.ID) )
	
	if getPaymentContractRequestResult.Success == false {
		t.Errorf(getPaymentContractRequestResult.Msg)
	} else {
		fmt.Println("Check Get PaymentContract success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPaymentContractObj,_ := getPaymentContractRequestResult.Data. (model.PaymentContract)
	comparePaymentContract := cmp.Equal(createPaymentContractObj.ID, getPaymentContractObj.ID)
	
	if  comparePaymentContract == false	{
		t.Errorf( "Created PaymentContract object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPaymentContractRequestResult := dao.GetAllPaymentContract()

	if getAllPaymentContractRequestResult.Success == false {
			t.Errorf(getAllPaymentContractRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PaymentContract success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPaymentContractObj []model.PaymentContract = getAllPaymentContractRequestResult.Data. ([]model.PaymentContract)
		
	equalPaymentContract := cmp.Equal(createPaymentContractObj.ID, getAllPaymentContractObj[len(getAllPaymentContractObj)-1].ID)
		
	if equalPaymentContract == false {
		t.Errorf( "Created object is not equal to the last entry in PaymentContract[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PaymentContract
	// --------------------------------------------------------------	
	deletePaymentContractRequestResult := dao.DeletePaymentContract(uint64(createPaymentContractObj.ID))

	if deletePaymentContractRequestResult.Success == false {
			t.Errorf(deletePaymentContractRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PaymentContract success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPaymentContractRequestResult = dao.GetPaymentContract( uint64(createPaymentContractObj.ID) )
	
	if getPaymentContractRequestResult.Success == true {
		t.Errorf(getPaymentContractRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPaymentProcessorCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PaymentProcessor
	//----------------------------------------------------------------------------
	PaymentProcessorObj := model.PaymentProcessor                                                                                            {Name:"test value for Name",ProcessorCode:"test value for ProcessorCode",NetworkSupport:"test value for NetworkSupport"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPaymentProcessorRequestResult := dao.CreatePaymentProcessor( PaymentProcessorObj )
	
	if createPaymentProcessorRequestResult.Success == false {
		t.Errorf(createPaymentProcessorRequestResult.Msg)
	} else {
		fmt.Println("Check Create PaymentProcessor success...")
	}
	
	createPaymentProcessorObj,_ := createPaymentProcessorRequestResult.Data. (model.PaymentProcessor)

	// --------------------------------------------------------------
	// Check PaymentProcessor Obj ID
	// --------------------------------------------------------------	
	if createPaymentProcessorObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PaymentProcessor" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPaymentProcessorRequestResult := dao.GetPaymentProcessor( uint64(createPaymentProcessorObj.ID) )
	
	if getPaymentProcessorRequestResult.Success == false {
		t.Errorf(getPaymentProcessorRequestResult.Msg)
	} else {
		fmt.Println("Check Get PaymentProcessor success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPaymentProcessorObj,_ := getPaymentProcessorRequestResult.Data. (model.PaymentProcessor)
	comparePaymentProcessor := cmp.Equal(createPaymentProcessorObj.ID, getPaymentProcessorObj.ID)
	
	if  comparePaymentProcessor == false	{
		t.Errorf( "Created PaymentProcessor object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPaymentProcessorRequestResult := dao.GetAllPaymentProcessor()

	if getAllPaymentProcessorRequestResult.Success == false {
			t.Errorf(getAllPaymentProcessorRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PaymentProcessor success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPaymentProcessorObj []model.PaymentProcessor = getAllPaymentProcessorRequestResult.Data. ([]model.PaymentProcessor)
		
	equalPaymentProcessor := cmp.Equal(createPaymentProcessorObj.ID, getAllPaymentProcessorObj[len(getAllPaymentProcessorObj)-1].ID)
		
	if equalPaymentProcessor == false {
		t.Errorf( "Created object is not equal to the last entry in PaymentProcessor[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PaymentProcessor
	// --------------------------------------------------------------	
	deletePaymentProcessorRequestResult := dao.DeletePaymentProcessor(uint64(createPaymentProcessorObj.ID))

	if deletePaymentProcessorRequestResult.Success == false {
			t.Errorf(deletePaymentProcessorRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PaymentProcessor success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPaymentProcessorRequestResult = dao.GetPaymentProcessor( uint64(createPaymentProcessorObj.ID) )
	
	if getPaymentProcessorRequestResult.Success == true {
		t.Errorf(getPaymentProcessorRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTransactionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Transaction
	//----------------------------------------------------------------------------
	TransactionObj := model.Transaction                                                                                                                                                                                    {Amount:new Money(),Fee:new Money(),ExchangeRate:"test value",CreatedAt:new DateTime(),CompletedAt:new DateTime(),Narrative:"test value for Narrative",TransactionType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTransactionRequestResult := dao.CreateTransaction( TransactionObj )
	
	if createTransactionRequestResult.Success == false {
		t.Errorf(createTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Create Transaction success...")
	}
	
	createTransactionObj,_ := createTransactionRequestResult.Data. (model.Transaction)

	// --------------------------------------------------------------
	// Check Transaction Obj ID
	// --------------------------------------------------------------	
	if createTransactionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Transaction" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTransactionRequestResult := dao.GetTransaction( uint64(createTransactionObj.ID) )
	
	if getTransactionRequestResult.Success == false {
		t.Errorf(getTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Get Transaction success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTransactionObj,_ := getTransactionRequestResult.Data. (model.Transaction)
	compareTransaction := cmp.Equal(createTransactionObj.ID, getTransactionObj.ID)
	
	if  compareTransaction == false	{
		t.Errorf( "Created Transaction object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTransactionRequestResult := dao.GetAllTransaction()

	if getAllTransactionRequestResult.Success == false {
			t.Errorf(getAllTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Transaction success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTransactionObj []model.Transaction = getAllTransactionRequestResult.Data. ([]model.Transaction)
		
	equalTransaction := cmp.Equal(createTransactionObj.ID, getAllTransactionObj[len(getAllTransactionObj)-1].ID)
		
	if equalTransaction == false {
		t.Errorf( "Created object is not equal to the last entry in Transaction[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Transaction
	// --------------------------------------------------------------	
	deleteTransactionRequestResult := dao.DeleteTransaction(uint64(createTransactionObj.ID))

	if deleteTransactionRequestResult.Success == false {
			t.Errorf(deleteTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Transaction success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTransactionRequestResult = dao.GetTransaction( uint64(createTransactionObj.ID) )
	
	if getTransactionRequestResult.Success == true {
		t.Errorf(getTransactionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPaymentOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PaymentOrder
	//----------------------------------------------------------------------------
	PaymentOrderObj := model.PaymentOrder                                                                                                                                                                    {OrderReference:"test value for OrderReference",RequestedExecutionDate:time.Now(),Purpose:"test value for Purpose",PaymentMethod:0,Status:0,Priority:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPaymentOrderRequestResult := dao.CreatePaymentOrder( PaymentOrderObj )
	
	if createPaymentOrderRequestResult.Success == false {
		t.Errorf(createPaymentOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create PaymentOrder success...")
	}
	
	createPaymentOrderObj,_ := createPaymentOrderRequestResult.Data. (model.PaymentOrder)

	// --------------------------------------------------------------
	// Check PaymentOrder Obj ID
	// --------------------------------------------------------------	
	if createPaymentOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PaymentOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPaymentOrderRequestResult := dao.GetPaymentOrder( uint64(createPaymentOrderObj.ID) )
	
	if getPaymentOrderRequestResult.Success == false {
		t.Errorf(getPaymentOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get PaymentOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPaymentOrderObj,_ := getPaymentOrderRequestResult.Data. (model.PaymentOrder)
	comparePaymentOrder := cmp.Equal(createPaymentOrderObj.ID, getPaymentOrderObj.ID)
	
	if  comparePaymentOrder == false	{
		t.Errorf( "Created PaymentOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPaymentOrderRequestResult := dao.GetAllPaymentOrder()

	if getAllPaymentOrderRequestResult.Success == false {
			t.Errorf(getAllPaymentOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PaymentOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPaymentOrderObj []model.PaymentOrder = getAllPaymentOrderRequestResult.Data. ([]model.PaymentOrder)
		
	equalPaymentOrder := cmp.Equal(createPaymentOrderObj.ID, getAllPaymentOrderObj[len(getAllPaymentOrderObj)-1].ID)
		
	if equalPaymentOrder == false {
		t.Errorf( "Created object is not equal to the last entry in PaymentOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PaymentOrder
	// --------------------------------------------------------------	
	deletePaymentOrderRequestResult := dao.DeletePaymentOrder(uint64(createPaymentOrderObj.ID))

	if deletePaymentOrderRequestResult.Success == false {
			t.Errorf(deletePaymentOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PaymentOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPaymentOrderRequestResult = dao.GetPaymentOrder( uint64(createPaymentOrderObj.ID) )
	
	if getPaymentOrderRequestResult.Success == true {
		t.Errorf(getPaymentOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBeneficiaryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Beneficiary
	//----------------------------------------------------------------------------
	BeneficiaryObj := model.Beneficiary                                                                                            {Name:"test value for Name",AccountIdentifier:new AccountIdentifier(),Iban:new IBAN(),Bic:new BIC(),Address:new Address()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBeneficiaryRequestResult := dao.CreateBeneficiary( BeneficiaryObj )
	
	if createBeneficiaryRequestResult.Success == false {
		t.Errorf(createBeneficiaryRequestResult.Msg)
	} else {
		fmt.Println("Check Create Beneficiary success...")
	}
	
	createBeneficiaryObj,_ := createBeneficiaryRequestResult.Data. (model.Beneficiary)

	// --------------------------------------------------------------
	// Check Beneficiary Obj ID
	// --------------------------------------------------------------	
	if createBeneficiaryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Beneficiary" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBeneficiaryRequestResult := dao.GetBeneficiary( uint64(createBeneficiaryObj.ID) )
	
	if getBeneficiaryRequestResult.Success == false {
		t.Errorf(getBeneficiaryRequestResult.Msg)
	} else {
		fmt.Println("Check Get Beneficiary success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBeneficiaryObj,_ := getBeneficiaryRequestResult.Data. (model.Beneficiary)
	compareBeneficiary := cmp.Equal(createBeneficiaryObj.ID, getBeneficiaryObj.ID)
	
	if  compareBeneficiary == false	{
		t.Errorf( "Created Beneficiary object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBeneficiaryRequestResult := dao.GetAllBeneficiary()

	if getAllBeneficiaryRequestResult.Success == false {
			t.Errorf(getAllBeneficiaryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Beneficiary success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBeneficiaryObj []model.Beneficiary = getAllBeneficiaryRequestResult.Data. ([]model.Beneficiary)
		
	equalBeneficiary := cmp.Equal(createBeneficiaryObj.ID, getAllBeneficiaryObj[len(getAllBeneficiaryObj)-1].ID)
		
	if equalBeneficiary == false {
		t.Errorf( "Created object is not equal to the last entry in Beneficiary[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Beneficiary
	// --------------------------------------------------------------	
	deleteBeneficiaryRequestResult := dao.DeleteBeneficiary(uint64(createBeneficiaryObj.ID))

	if deleteBeneficiaryRequestResult.Success == false {
			t.Errorf(deleteBeneficiaryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Beneficiary success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBeneficiaryRequestResult = dao.GetBeneficiary( uint64(createBeneficiaryObj.ID) )
	
	if getBeneficiaryRequestResult.Success == true {
		t.Errorf(getBeneficiaryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAppliedFeeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AppliedFee
	//----------------------------------------------------------------------------
	AppliedFeeObj := model.AppliedFee                                                            {Amount:new Money(),Description:"test value for Description",FeeType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAppliedFeeRequestResult := dao.CreateAppliedFee( AppliedFeeObj )
	
	if createAppliedFeeRequestResult.Success == false {
		t.Errorf(createAppliedFeeRequestResult.Msg)
	} else {
		fmt.Println("Check Create AppliedFee success...")
	}
	
	createAppliedFeeObj,_ := createAppliedFeeRequestResult.Data. (model.AppliedFee)

	// --------------------------------------------------------------
	// Check AppliedFee Obj ID
	// --------------------------------------------------------------	
	if createAppliedFeeObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AppliedFee" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAppliedFeeRequestResult := dao.GetAppliedFee( uint64(createAppliedFeeObj.ID) )
	
	if getAppliedFeeRequestResult.Success == false {
		t.Errorf(getAppliedFeeRequestResult.Msg)
	} else {
		fmt.Println("Check Get AppliedFee success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAppliedFeeObj,_ := getAppliedFeeRequestResult.Data. (model.AppliedFee)
	compareAppliedFee := cmp.Equal(createAppliedFeeObj.ID, getAppliedFeeObj.ID)
	
	if  compareAppliedFee == false	{
		t.Errorf( "Created AppliedFee object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAppliedFeeRequestResult := dao.GetAllAppliedFee()

	if getAllAppliedFeeRequestResult.Success == false {
			t.Errorf(getAllAppliedFeeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AppliedFee success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAppliedFeeObj []model.AppliedFee = getAllAppliedFeeRequestResult.Data. ([]model.AppliedFee)
		
	equalAppliedFee := cmp.Equal(createAppliedFeeObj.ID, getAllAppliedFeeObj[len(getAllAppliedFeeObj)-1].ID)
		
	if equalAppliedFee == false {
		t.Errorf( "Created object is not equal to the last entry in AppliedFee[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AppliedFee
	// --------------------------------------------------------------	
	deleteAppliedFeeRequestResult := dao.DeleteAppliedFee(uint64(createAppliedFeeObj.ID))

	if deleteAppliedFeeRequestResult.Success == false {
			t.Errorf(deleteAppliedFeeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AppliedFee success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAppliedFeeRequestResult = dao.GetAppliedFee( uint64(createAppliedFeeObj.ID) )
	
	if getAppliedFeeRequestResult.Success == true {
		t.Errorf(getAppliedFeeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFXQuoteCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FXQuote
	//----------------------------------------------------------------------------
	FXQuoteObj := model.FXQuote                                                                                                                                                                    {BaseCurrency:"test value for BaseCurrency",QuoteCurrency:"test value for QuoteCurrency",Rate:"test value",QuotedAt:new DateTime(),ExpiresAt:new DateTime(),PriceType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFXQuoteRequestResult := dao.CreateFXQuote( FXQuoteObj )
	
	if createFXQuoteRequestResult.Success == false {
		t.Errorf(createFXQuoteRequestResult.Msg)
	} else {
		fmt.Println("Check Create FXQuote success...")
	}
	
	createFXQuoteObj,_ := createFXQuoteRequestResult.Data. (model.FXQuote)

	// --------------------------------------------------------------
	// Check FXQuote Obj ID
	// --------------------------------------------------------------	
	if createFXQuoteObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for FXQuote" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFXQuoteRequestResult := dao.GetFXQuote( uint64(createFXQuoteObj.ID) )
	
	if getFXQuoteRequestResult.Success == false {
		t.Errorf(getFXQuoteRequestResult.Msg)
	} else {
		fmt.Println("Check Get FXQuote success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFXQuoteObj,_ := getFXQuoteRequestResult.Data. (model.FXQuote)
	compareFXQuote := cmp.Equal(createFXQuoteObj.ID, getFXQuoteObj.ID)
	
	if  compareFXQuote == false	{
		t.Errorf( "Created FXQuote object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFXQuoteRequestResult := dao.GetAllFXQuote()

	if getAllFXQuoteRequestResult.Success == false {
			t.Errorf(getAllFXQuoteRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FXQuote success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFXQuoteObj []model.FXQuote = getAllFXQuoteRequestResult.Data. ([]model.FXQuote)
		
	equalFXQuote := cmp.Equal(createFXQuoteObj.ID, getAllFXQuoteObj[len(getAllFXQuoteObj)-1].ID)
		
	if equalFXQuote == false {
		t.Errorf( "Created object is not equal to the last entry in FXQuote[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FXQuote
	// --------------------------------------------------------------	
	deleteFXQuoteRequestResult := dao.DeleteFXQuote(uint64(createFXQuoteObj.ID))

	if deleteFXQuoteRequestResult.Success == false {
			t.Errorf(deleteFXQuoteRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FXQuote success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFXQuoteRequestResult = dao.GetFXQuote( uint64(createFXQuoteObj.ID) )
	
	if getFXQuoteRequestResult.Success == true {
		t.Errorf(getFXQuoteRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFXDealCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FXDeal
	//----------------------------------------------------------------------------
	FXDealObj := model.FXDeal                                                                                                                                                                                                                                            {DealReference:"test value for DealReference",BaseCurrency:"test value for BaseCurrency",QuoteCurrency:"test value for QuoteCurrency",Rate:"test value",Amount:new Money(),SettlementDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFXDealRequestResult := dao.CreateFXDeal( FXDealObj )
	
	if createFXDealRequestResult.Success == false {
		t.Errorf(createFXDealRequestResult.Msg)
	} else {
		fmt.Println("Check Create FXDeal success...")
	}
	
	createFXDealObj,_ := createFXDealRequestResult.Data. (model.FXDeal)

	// --------------------------------------------------------------
	// Check FXDeal Obj ID
	// --------------------------------------------------------------	
	if createFXDealObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for FXDeal" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFXDealRequestResult := dao.GetFXDeal( uint64(createFXDealObj.ID) )
	
	if getFXDealRequestResult.Success == false {
		t.Errorf(getFXDealRequestResult.Msg)
	} else {
		fmt.Println("Check Get FXDeal success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFXDealObj,_ := getFXDealRequestResult.Data. (model.FXDeal)
	compareFXDeal := cmp.Equal(createFXDealObj.ID, getFXDealObj.ID)
	
	if  compareFXDeal == false	{
		t.Errorf( "Created FXDeal object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFXDealRequestResult := dao.GetAllFXDeal()

	if getAllFXDealRequestResult.Success == false {
			t.Errorf(getAllFXDealRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FXDeal success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFXDealObj []model.FXDeal = getAllFXDealRequestResult.Data. ([]model.FXDeal)
		
	equalFXDeal := cmp.Equal(createFXDealObj.ID, getAllFXDealObj[len(getAllFXDealObj)-1].ID)
		
	if equalFXDeal == false {
		t.Errorf( "Created object is not equal to the last entry in FXDeal[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FXDeal
	// --------------------------------------------------------------	
	deleteFXDealRequestResult := dao.DeleteFXDeal(uint64(createFXDealObj.ID))

	if deleteFXDealRequestResult.Success == false {
			t.Errorf(deleteFXDealRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FXDeal success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFXDealRequestResult = dao.GetFXDeal( uint64(createFXDealObj.ID) )
	
	if getFXDealRequestResult.Success == true {
		t.Errorf(getFXDealRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSettlementBatchCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SettlementBatch
	//----------------------------------------------------------------------------
	SettlementBatchObj := model.SettlementBatch                                                                                                                            {BatchId:"test value for BatchId",PeriodStart:new DateTime(),PeriodEnd:new DateTime(),TotalVolume:new Money(),TotalCount:100,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSettlementBatchRequestResult := dao.CreateSettlementBatch( SettlementBatchObj )
	
	if createSettlementBatchRequestResult.Success == false {
		t.Errorf(createSettlementBatchRequestResult.Msg)
	} else {
		fmt.Println("Check Create SettlementBatch success...")
	}
	
	createSettlementBatchObj,_ := createSettlementBatchRequestResult.Data. (model.SettlementBatch)

	// --------------------------------------------------------------
	// Check SettlementBatch Obj ID
	// --------------------------------------------------------------	
	if createSettlementBatchObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SettlementBatch" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSettlementBatchRequestResult := dao.GetSettlementBatch( uint64(createSettlementBatchObj.ID) )
	
	if getSettlementBatchRequestResult.Success == false {
		t.Errorf(getSettlementBatchRequestResult.Msg)
	} else {
		fmt.Println("Check Get SettlementBatch success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSettlementBatchObj,_ := getSettlementBatchRequestResult.Data. (model.SettlementBatch)
	compareSettlementBatch := cmp.Equal(createSettlementBatchObj.ID, getSettlementBatchObj.ID)
	
	if  compareSettlementBatch == false	{
		t.Errorf( "Created SettlementBatch object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSettlementBatchRequestResult := dao.GetAllSettlementBatch()

	if getAllSettlementBatchRequestResult.Success == false {
			t.Errorf(getAllSettlementBatchRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SettlementBatch success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSettlementBatchObj []model.SettlementBatch = getAllSettlementBatchRequestResult.Data. ([]model.SettlementBatch)
		
	equalSettlementBatch := cmp.Equal(createSettlementBatchObj.ID, getAllSettlementBatchObj[len(getAllSettlementBatchObj)-1].ID)
		
	if equalSettlementBatch == false {
		t.Errorf( "Created object is not equal to the last entry in SettlementBatch[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SettlementBatch
	// --------------------------------------------------------------	
	deleteSettlementBatchRequestResult := dao.DeleteSettlementBatch(uint64(createSettlementBatchObj.ID))

	if deleteSettlementBatchRequestResult.Success == false {
			t.Errorf(deleteSettlementBatchRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SettlementBatch success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSettlementBatchRequestResult = dao.GetSettlementBatch( uint64(createSettlementBatchObj.ID) )
	
	if getSettlementBatchRequestResult.Success == true {
		t.Errorf(getSettlementBatchRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPayoutCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Payout
	//----------------------------------------------------------------------------
	PayoutObj := model.Payout                                                                                                                                                                                                            {PayoutReference:"test value for PayoutReference",Amount:new Money(),Currency:"test value for Currency",ScheduledDate:time.Now(),PaidDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPayoutRequestResult := dao.CreatePayout( PayoutObj )
	
	if createPayoutRequestResult.Success == false {
		t.Errorf(createPayoutRequestResult.Msg)
	} else {
		fmt.Println("Check Create Payout success...")
	}
	
	createPayoutObj,_ := createPayoutRequestResult.Data. (model.Payout)

	// --------------------------------------------------------------
	// Check Payout Obj ID
	// --------------------------------------------------------------	
	if createPayoutObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Payout" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPayoutRequestResult := dao.GetPayout( uint64(createPayoutObj.ID) )
	
	if getPayoutRequestResult.Success == false {
		t.Errorf(getPayoutRequestResult.Msg)
	} else {
		fmt.Println("Check Get Payout success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPayoutObj,_ := getPayoutRequestResult.Data. (model.Payout)
	comparePayout := cmp.Equal(createPayoutObj.ID, getPayoutObj.ID)
	
	if  comparePayout == false	{
		t.Errorf( "Created Payout object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPayoutRequestResult := dao.GetAllPayout()

	if getAllPayoutRequestResult.Success == false {
			t.Errorf(getAllPayoutRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Payout success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPayoutObj []model.Payout = getAllPayoutRequestResult.Data. ([]model.Payout)
		
	equalPayout := cmp.Equal(createPayoutObj.ID, getAllPayoutObj[len(getAllPayoutObj)-1].ID)
		
	if equalPayout == false {
		t.Errorf( "Created object is not equal to the last entry in Payout[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Payout
	// --------------------------------------------------------------	
	deletePayoutRequestResult := dao.DeletePayout(uint64(createPayoutObj.ID))

	if deletePayoutRequestResult.Success == false {
			t.Errorf(deletePayoutRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Payout success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPayoutRequestResult = dao.GetPayout( uint64(createPayoutObj.ID) )
	
	if getPayoutRequestResult.Success == true {
		t.Errorf(getPayoutRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDisputeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Dispute
	//----------------------------------------------------------------------------
	DisputeObj := model.Dispute                                                                                            {DisputeReference:"test value for DisputeReference",OpenedAt:new DateTime(),ClosedAt:new DateTime(),Reason:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDisputeRequestResult := dao.CreateDispute( DisputeObj )
	
	if createDisputeRequestResult.Success == false {
		t.Errorf(createDisputeRequestResult.Msg)
	} else {
		fmt.Println("Check Create Dispute success...")
	}
	
	createDisputeObj,_ := createDisputeRequestResult.Data. (model.Dispute)

	// --------------------------------------------------------------
	// Check Dispute Obj ID
	// --------------------------------------------------------------	
	if createDisputeObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Dispute" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDisputeRequestResult := dao.GetDispute( uint64(createDisputeObj.ID) )
	
	if getDisputeRequestResult.Success == false {
		t.Errorf(getDisputeRequestResult.Msg)
	} else {
		fmt.Println("Check Get Dispute success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDisputeObj,_ := getDisputeRequestResult.Data. (model.Dispute)
	compareDispute := cmp.Equal(createDisputeObj.ID, getDisputeObj.ID)
	
	if  compareDispute == false	{
		t.Errorf( "Created Dispute object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDisputeRequestResult := dao.GetAllDispute()

	if getAllDisputeRequestResult.Success == false {
			t.Errorf(getAllDisputeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Dispute success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDisputeObj []model.Dispute = getAllDisputeRequestResult.Data. ([]model.Dispute)
		
	equalDispute := cmp.Equal(createDisputeObj.ID, getAllDisputeObj[len(getAllDisputeObj)-1].ID)
		
	if equalDispute == false {
		t.Errorf( "Created object is not equal to the last entry in Dispute[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Dispute
	// --------------------------------------------------------------	
	deleteDisputeRequestResult := dao.DeleteDispute(uint64(createDisputeObj.ID))

	if deleteDisputeRequestResult.Success == false {
			t.Errorf(deleteDisputeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Dispute success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDisputeRequestResult = dao.GetDispute( uint64(createDisputeObj.ID) )
	
	if getDisputeRequestResult.Success == true {
		t.Errorf(getDisputeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestChargebackCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Chargeback
	//----------------------------------------------------------------------------
	ChargebackObj := model.Chargeback                                                                                            {ChargebackReference:"test value for ChargebackReference",Amount:new Money(),PostedAt:new DateTime(),Stage:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createChargebackRequestResult := dao.CreateChargeback( ChargebackObj )
	
	if createChargebackRequestResult.Success == false {
		t.Errorf(createChargebackRequestResult.Msg)
	} else {
		fmt.Println("Check Create Chargeback success...")
	}
	
	createChargebackObj,_ := createChargebackRequestResult.Data. (model.Chargeback)

	// --------------------------------------------------------------
	// Check Chargeback Obj ID
	// --------------------------------------------------------------	
	if createChargebackObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Chargeback" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getChargebackRequestResult := dao.GetChargeback( uint64(createChargebackObj.ID) )
	
	if getChargebackRequestResult.Success == false {
		t.Errorf(getChargebackRequestResult.Msg)
	} else {
		fmt.Println("Check Get Chargeback success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getChargebackObj,_ := getChargebackRequestResult.Data. (model.Chargeback)
	compareChargeback := cmp.Equal(createChargebackObj.ID, getChargebackObj.ID)
	
	if  compareChargeback == false	{
		t.Errorf( "Created Chargeback object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllChargebackRequestResult := dao.GetAllChargeback()

	if getAllChargebackRequestResult.Success == false {
			t.Errorf(getAllChargebackRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Chargeback success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllChargebackObj []model.Chargeback = getAllChargebackRequestResult.Data. ([]model.Chargeback)
		
	equalChargeback := cmp.Equal(createChargebackObj.ID, getAllChargebackObj[len(getAllChargebackObj)-1].ID)
		
	if equalChargeback == false {
		t.Errorf( "Created object is not equal to the last entry in Chargeback[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Chargeback
	// --------------------------------------------------------------	
	deleteChargebackRequestResult := dao.DeleteChargeback(uint64(createChargebackObj.ID))

	if deleteChargebackRequestResult.Success == false {
			t.Errorf(deleteChargebackRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Chargeback success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getChargebackRequestResult = dao.GetChargeback( uint64(createChargebackObj.ID) )
	
	if getChargebackRequestResult.Success == true {
		t.Errorf(getChargebackRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInvoiceCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Invoice
	//----------------------------------------------------------------------------
	InvoiceObj := model.Invoice                                                                                                                                                                                                            {InvoiceNumber:"test value for InvoiceNumber",IssueDate:time.Now(),DueDate:time.Now(),Total:new Money(),Currency:"test value for Currency",Status:0}

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


func TestAccountStatementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AccountStatement
	//----------------------------------------------------------------------------
	AccountStatementObj := model.AccountStatement                                                                                                                                                                                            {StatementNumber:"test value for StatementNumber",PeriodStart:time.Now(),PeriodEnd:time.Now(),OpeningBalance:new Money(),ClosingBalance:new Money(),GeneratedAt:new DateTime()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAccountStatementRequestResult := dao.CreateAccountStatement( AccountStatementObj )
	
	if createAccountStatementRequestResult.Success == false {
		t.Errorf(createAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Check Create AccountStatement success...")
	}
	
	createAccountStatementObj,_ := createAccountStatementRequestResult.Data. (model.AccountStatement)

	// --------------------------------------------------------------
	// Check AccountStatement Obj ID
	// --------------------------------------------------------------	
	if createAccountStatementObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AccountStatement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAccountStatementRequestResult := dao.GetAccountStatement( uint64(createAccountStatementObj.ID) )
	
	if getAccountStatementRequestResult.Success == false {
		t.Errorf(getAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Check Get AccountStatement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAccountStatementObj,_ := getAccountStatementRequestResult.Data. (model.AccountStatement)
	compareAccountStatement := cmp.Equal(createAccountStatementObj.ID, getAccountStatementObj.ID)
	
	if  compareAccountStatement == false	{
		t.Errorf( "Created AccountStatement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAccountStatementRequestResult := dao.GetAllAccountStatement()

	if getAllAccountStatementRequestResult.Success == false {
			t.Errorf(getAllAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AccountStatement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAccountStatementObj []model.AccountStatement = getAllAccountStatementRequestResult.Data. ([]model.AccountStatement)
		
	equalAccountStatement := cmp.Equal(createAccountStatementObj.ID, getAllAccountStatementObj[len(getAllAccountStatementObj)-1].ID)
		
	if equalAccountStatement == false {
		t.Errorf( "Created object is not equal to the last entry in AccountStatement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AccountStatement
	// --------------------------------------------------------------	
	deleteAccountStatementRequestResult := dao.DeleteAccountStatement(uint64(createAccountStatementObj.ID))

	if deleteAccountStatementRequestResult.Success == false {
			t.Errorf(deleteAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AccountStatement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAccountStatementRequestResult = dao.GetAccountStatement( uint64(createAccountStatementObj.ID) )
	
	if getAccountStatementRequestResult.Success == true {
		t.Errorf(getAccountStatementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDirectDebitMandateCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DirectDebitMandate
	//----------------------------------------------------------------------------
	DirectDebitMandateObj := model.DirectDebitMandate                                                                            {MandateId:"test value for MandateId",SignedAt:new DateTime(),Scheme:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDirectDebitMandateRequestResult := dao.CreateDirectDebitMandate( DirectDebitMandateObj )
	
	if createDirectDebitMandateRequestResult.Success == false {
		t.Errorf(createDirectDebitMandateRequestResult.Msg)
	} else {
		fmt.Println("Check Create DirectDebitMandate success...")
	}
	
	createDirectDebitMandateObj,_ := createDirectDebitMandateRequestResult.Data. (model.DirectDebitMandate)

	// --------------------------------------------------------------
	// Check DirectDebitMandate Obj ID
	// --------------------------------------------------------------	
	if createDirectDebitMandateObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DirectDebitMandate" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDirectDebitMandateRequestResult := dao.GetDirectDebitMandate( uint64(createDirectDebitMandateObj.ID) )
	
	if getDirectDebitMandateRequestResult.Success == false {
		t.Errorf(getDirectDebitMandateRequestResult.Msg)
	} else {
		fmt.Println("Check Get DirectDebitMandate success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDirectDebitMandateObj,_ := getDirectDebitMandateRequestResult.Data. (model.DirectDebitMandate)
	compareDirectDebitMandate := cmp.Equal(createDirectDebitMandateObj.ID, getDirectDebitMandateObj.ID)
	
	if  compareDirectDebitMandate == false	{
		t.Errorf( "Created DirectDebitMandate object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDirectDebitMandateRequestResult := dao.GetAllDirectDebitMandate()

	if getAllDirectDebitMandateRequestResult.Success == false {
			t.Errorf(getAllDirectDebitMandateRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DirectDebitMandate success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDirectDebitMandateObj []model.DirectDebitMandate = getAllDirectDebitMandateRequestResult.Data. ([]model.DirectDebitMandate)
		
	equalDirectDebitMandate := cmp.Equal(createDirectDebitMandateObj.ID, getAllDirectDebitMandateObj[len(getAllDirectDebitMandateObj)-1].ID)
		
	if equalDirectDebitMandate == false {
		t.Errorf( "Created object is not equal to the last entry in DirectDebitMandate[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DirectDebitMandate
	// --------------------------------------------------------------	
	deleteDirectDebitMandateRequestResult := dao.DeleteDirectDebitMandate(uint64(createDirectDebitMandateObj.ID))

	if deleteDirectDebitMandateRequestResult.Success == false {
			t.Errorf(deleteDirectDebitMandateRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DirectDebitMandate success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDirectDebitMandateRequestResult = dao.GetDirectDebitMandate( uint64(createDirectDebitMandateObj.ID) )
	
	if getDirectDebitMandateRequestResult.Success == true {
		t.Errorf(getDirectDebitMandateRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCreditorCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Creditor
	//----------------------------------------------------------------------------
	CreditorObj := model.Creditor                                                            {Name:"test value for Name",Bic:new BIC(),Address:new Address()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCreditorRequestResult := dao.CreateCreditor( CreditorObj )
	
	if createCreditorRequestResult.Success == false {
		t.Errorf(createCreditorRequestResult.Msg)
	} else {
		fmt.Println("Check Create Creditor success...")
	}
	
	createCreditorObj,_ := createCreditorRequestResult.Data. (model.Creditor)

	// --------------------------------------------------------------
	// Check Creditor Obj ID
	// --------------------------------------------------------------	
	if createCreditorObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Creditor" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCreditorRequestResult := dao.GetCreditor( uint64(createCreditorObj.ID) )
	
	if getCreditorRequestResult.Success == false {
		t.Errorf(getCreditorRequestResult.Msg)
	} else {
		fmt.Println("Check Get Creditor success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCreditorObj,_ := getCreditorRequestResult.Data. (model.Creditor)
	compareCreditor := cmp.Equal(createCreditorObj.ID, getCreditorObj.ID)
	
	if  compareCreditor == false	{
		t.Errorf( "Created Creditor object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCreditorRequestResult := dao.GetAllCreditor()

	if getAllCreditorRequestResult.Success == false {
			t.Errorf(getAllCreditorRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Creditor success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCreditorObj []model.Creditor = getAllCreditorRequestResult.Data. ([]model.Creditor)
		
	equalCreditor := cmp.Equal(createCreditorObj.ID, getAllCreditorObj[len(getAllCreditorObj)-1].ID)
		
	if equalCreditor == false {
		t.Errorf( "Created object is not equal to the last entry in Creditor[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Creditor
	// --------------------------------------------------------------	
	deleteCreditorRequestResult := dao.DeleteCreditor(uint64(createCreditorObj.ID))

	if deleteCreditorRequestResult.Success == false {
			t.Errorf(deleteCreditorRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Creditor success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCreditorRequestResult = dao.GetCreditor( uint64(createCreditorObj.ID) )
	
	if getCreditorRequestResult.Success == true {
		t.Errorf(getCreditorRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLoanApplicationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LoanApplication
	//----------------------------------------------------------------------------
	LoanApplicationObj := model.LoanApplication                                                                                                                                            {ApplicationNumber:"test value for ApplicationNumber",AmountRequested:new Money(),TermMonths:100,SubmittedAt:new DateTime(),Product:0,Purpose:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLoanApplicationRequestResult := dao.CreateLoanApplication( LoanApplicationObj )
	
	if createLoanApplicationRequestResult.Success == false {
		t.Errorf(createLoanApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check Create LoanApplication success...")
	}
	
	createLoanApplicationObj,_ := createLoanApplicationRequestResult.Data. (model.LoanApplication)

	// --------------------------------------------------------------
	// Check LoanApplication Obj ID
	// --------------------------------------------------------------	
	if createLoanApplicationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LoanApplication" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLoanApplicationRequestResult := dao.GetLoanApplication( uint64(createLoanApplicationObj.ID) )
	
	if getLoanApplicationRequestResult.Success == false {
		t.Errorf(getLoanApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check Get LoanApplication success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLoanApplicationObj,_ := getLoanApplicationRequestResult.Data. (model.LoanApplication)
	compareLoanApplication := cmp.Equal(createLoanApplicationObj.ID, getLoanApplicationObj.ID)
	
	if  compareLoanApplication == false	{
		t.Errorf( "Created LoanApplication object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLoanApplicationRequestResult := dao.GetAllLoanApplication()

	if getAllLoanApplicationRequestResult.Success == false {
			t.Errorf(getAllLoanApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LoanApplication success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLoanApplicationObj []model.LoanApplication = getAllLoanApplicationRequestResult.Data. ([]model.LoanApplication)
		
	equalLoanApplication := cmp.Equal(createLoanApplicationObj.ID, getAllLoanApplicationObj[len(getAllLoanApplicationObj)-1].ID)
		
	if equalLoanApplication == false {
		t.Errorf( "Created object is not equal to the last entry in LoanApplication[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LoanApplication
	// --------------------------------------------------------------	
	deleteLoanApplicationRequestResult := dao.DeleteLoanApplication(uint64(createLoanApplicationObj.ID))

	if deleteLoanApplicationRequestResult.Success == false {
			t.Errorf(deleteLoanApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LoanApplication success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLoanApplicationRequestResult = dao.GetLoanApplication( uint64(createLoanApplicationObj.ID) )
	
	if getLoanApplicationRequestResult.Success == true {
		t.Errorf(getLoanApplicationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRiskAssessmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RiskAssessment
	//----------------------------------------------------------------------------
	RiskAssessmentObj := model.RiskAssessment                                                                                                            {Score:new RiskScore(),AssessedAt:new DateTime(),ModelVersion:"test value for ModelVersion",Notes:"test value for Notes",Decision:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRiskAssessmentRequestResult := dao.CreateRiskAssessment( RiskAssessmentObj )
	
	if createRiskAssessmentRequestResult.Success == false {
		t.Errorf(createRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create RiskAssessment success...")
	}
	
	createRiskAssessmentObj,_ := createRiskAssessmentRequestResult.Data. (model.RiskAssessment)

	// --------------------------------------------------------------
	// Check RiskAssessment Obj ID
	// --------------------------------------------------------------	
	if createRiskAssessmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for RiskAssessment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRiskAssessmentRequestResult := dao.GetRiskAssessment( uint64(createRiskAssessmentObj.ID) )
	
	if getRiskAssessmentRequestResult.Success == false {
		t.Errorf(getRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get RiskAssessment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRiskAssessmentObj,_ := getRiskAssessmentRequestResult.Data. (model.RiskAssessment)
	compareRiskAssessment := cmp.Equal(createRiskAssessmentObj.ID, getRiskAssessmentObj.ID)
	
	if  compareRiskAssessment == false	{
		t.Errorf( "Created RiskAssessment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRiskAssessmentRequestResult := dao.GetAllRiskAssessment()

	if getAllRiskAssessmentRequestResult.Success == false {
			t.Errorf(getAllRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RiskAssessment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRiskAssessmentObj []model.RiskAssessment = getAllRiskAssessmentRequestResult.Data. ([]model.RiskAssessment)
		
	equalRiskAssessment := cmp.Equal(createRiskAssessmentObj.ID, getAllRiskAssessmentObj[len(getAllRiskAssessmentObj)-1].ID)
		
	if equalRiskAssessment == false {
		t.Errorf( "Created object is not equal to the last entry in RiskAssessment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RiskAssessment
	// --------------------------------------------------------------	
	deleteRiskAssessmentRequestResult := dao.DeleteRiskAssessment(uint64(createRiskAssessmentObj.ID))

	if deleteRiskAssessmentRequestResult.Success == false {
			t.Errorf(deleteRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RiskAssessment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRiskAssessmentRequestResult = dao.GetRiskAssessment( uint64(createRiskAssessmentObj.ID) )
	
	if getRiskAssessmentRequestResult.Success == true {
		t.Errorf(getRiskAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLoanCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Loan
	//----------------------------------------------------------------------------
	LoanObj := model.Loan                                                                                                                                                                                                                                                    {LoanNumber:"test value for LoanNumber",Principal:new Money(),InterestRate:"test value",OriginationDate:time.Now(),MaturityDate:time.Now(),RateType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLoanRequestResult := dao.CreateLoan( LoanObj )
	
	if createLoanRequestResult.Success == false {
		t.Errorf(createLoanRequestResult.Msg)
	} else {
		fmt.Println("Check Create Loan success...")
	}
	
	createLoanObj,_ := createLoanRequestResult.Data. (model.Loan)

	// --------------------------------------------------------------
	// Check Loan Obj ID
	// --------------------------------------------------------------	
	if createLoanObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Loan" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLoanRequestResult := dao.GetLoan( uint64(createLoanObj.ID) )
	
	if getLoanRequestResult.Success == false {
		t.Errorf(getLoanRequestResult.Msg)
	} else {
		fmt.Println("Check Get Loan success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLoanObj,_ := getLoanRequestResult.Data. (model.Loan)
	compareLoan := cmp.Equal(createLoanObj.ID, getLoanObj.ID)
	
	if  compareLoan == false	{
		t.Errorf( "Created Loan object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLoanRequestResult := dao.GetAllLoan()

	if getAllLoanRequestResult.Success == false {
			t.Errorf(getAllLoanRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Loan success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLoanObj []model.Loan = getAllLoanRequestResult.Data. ([]model.Loan)
		
	equalLoan := cmp.Equal(createLoanObj.ID, getAllLoanObj[len(getAllLoanObj)-1].ID)
		
	if equalLoan == false {
		t.Errorf( "Created object is not equal to the last entry in Loan[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Loan
	// --------------------------------------------------------------	
	deleteLoanRequestResult := dao.DeleteLoan(uint64(createLoanObj.ID))

	if deleteLoanRequestResult.Success == false {
			t.Errorf(deleteLoanRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Loan success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLoanRequestResult = dao.GetLoan( uint64(createLoanObj.ID) )
	
	if getLoanRequestResult.Success == true {
		t.Errorf(getLoanRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRepaymentScheduleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RepaymentSchedule
	//----------------------------------------------------------------------------
	RepaymentScheduleObj := model.RepaymentSchedule                                                                                                                                                    {InstallmentNumber:100,DueDate:time.Now(),AmountDue:new Money(),PrincipalDue:new Money(),InterestDue:new Money(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRepaymentScheduleRequestResult := dao.CreateRepaymentSchedule( RepaymentScheduleObj )
	
	if createRepaymentScheduleRequestResult.Success == false {
		t.Errorf(createRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Create RepaymentSchedule success...")
	}
	
	createRepaymentScheduleObj,_ := createRepaymentScheduleRequestResult.Data. (model.RepaymentSchedule)

	// --------------------------------------------------------------
	// Check RepaymentSchedule Obj ID
	// --------------------------------------------------------------	
	if createRepaymentScheduleObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for RepaymentSchedule" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRepaymentScheduleRequestResult := dao.GetRepaymentSchedule( uint64(createRepaymentScheduleObj.ID) )
	
	if getRepaymentScheduleRequestResult.Success == false {
		t.Errorf(getRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Get RepaymentSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRepaymentScheduleObj,_ := getRepaymentScheduleRequestResult.Data. (model.RepaymentSchedule)
	compareRepaymentSchedule := cmp.Equal(createRepaymentScheduleObj.ID, getRepaymentScheduleObj.ID)
	
	if  compareRepaymentSchedule == false	{
		t.Errorf( "Created RepaymentSchedule object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRepaymentScheduleRequestResult := dao.GetAllRepaymentSchedule()

	if getAllRepaymentScheduleRequestResult.Success == false {
			t.Errorf(getAllRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RepaymentSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRepaymentScheduleObj []model.RepaymentSchedule = getAllRepaymentScheduleRequestResult.Data. ([]model.RepaymentSchedule)
		
	equalRepaymentSchedule := cmp.Equal(createRepaymentScheduleObj.ID, getAllRepaymentScheduleObj[len(getAllRepaymentScheduleObj)-1].ID)
		
	if equalRepaymentSchedule == false {
		t.Errorf( "Created object is not equal to the last entry in RepaymentSchedule[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RepaymentSchedule
	// --------------------------------------------------------------	
	deleteRepaymentScheduleRequestResult := dao.DeleteRepaymentSchedule(uint64(createRepaymentScheduleObj.ID))

	if deleteRepaymentScheduleRequestResult.Success == false {
			t.Errorf(deleteRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RepaymentSchedule success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRepaymentScheduleRequestResult = dao.GetRepaymentSchedule( uint64(createRepaymentScheduleObj.ID) )
	
	if getRepaymentScheduleRequestResult.Success == true {
		t.Errorf(getRepaymentScheduleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCollateralCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Collateral
	//----------------------------------------------------------------------------
	CollateralObj := model.Collateral                                                            {Description:"test value for Description",Value:new Money(),CollateralType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCollateralRequestResult := dao.CreateCollateral( CollateralObj )
	
	if createCollateralRequestResult.Success == false {
		t.Errorf(createCollateralRequestResult.Msg)
	} else {
		fmt.Println("Check Create Collateral success...")
	}
	
	createCollateralObj,_ := createCollateralRequestResult.Data. (model.Collateral)

	// --------------------------------------------------------------
	// Check Collateral Obj ID
	// --------------------------------------------------------------	
	if createCollateralObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Collateral" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCollateralRequestResult := dao.GetCollateral( uint64(createCollateralObj.ID) )
	
	if getCollateralRequestResult.Success == false {
		t.Errorf(getCollateralRequestResult.Msg)
	} else {
		fmt.Println("Check Get Collateral success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCollateralObj,_ := getCollateralRequestResult.Data. (model.Collateral)
	compareCollateral := cmp.Equal(createCollateralObj.ID, getCollateralObj.ID)
	
	if  compareCollateral == false	{
		t.Errorf( "Created Collateral object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCollateralRequestResult := dao.GetAllCollateral()

	if getAllCollateralRequestResult.Success == false {
			t.Errorf(getAllCollateralRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Collateral success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCollateralObj []model.Collateral = getAllCollateralRequestResult.Data. ([]model.Collateral)
		
	equalCollateral := cmp.Equal(createCollateralObj.ID, getAllCollateralObj[len(getAllCollateralObj)-1].ID)
		
	if equalCollateral == false {
		t.Errorf( "Created object is not equal to the last entry in Collateral[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Collateral
	// --------------------------------------------------------------	
	deleteCollateralRequestResult := dao.DeleteCollateral(uint64(createCollateralObj.ID))

	if deleteCollateralRequestResult.Success == false {
			t.Errorf(deleteCollateralRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Collateral success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCollateralRequestResult = dao.GetCollateral( uint64(createCollateralObj.ID) )
	
	if getCollateralRequestResult.Success == true {
		t.Errorf(getCollateralRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLoanTransactionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LoanTransaction
	//----------------------------------------------------------------------------
	LoanTransactionObj := model.LoanTransaction                                                                                                                    {TransactionId:new TransactionId(),Amount:new Money(),PostingDate:time.Now(),Type:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLoanTransactionRequestResult := dao.CreateLoanTransaction( LoanTransactionObj )
	
	if createLoanTransactionRequestResult.Success == false {
		t.Errorf(createLoanTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Create LoanTransaction success...")
	}
	
	createLoanTransactionObj,_ := createLoanTransactionRequestResult.Data. (model.LoanTransaction)

	// --------------------------------------------------------------
	// Check LoanTransaction Obj ID
	// --------------------------------------------------------------	
	if createLoanTransactionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LoanTransaction" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLoanTransactionRequestResult := dao.GetLoanTransaction( uint64(createLoanTransactionObj.ID) )
	
	if getLoanTransactionRequestResult.Success == false {
		t.Errorf(getLoanTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Get LoanTransaction success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLoanTransactionObj,_ := getLoanTransactionRequestResult.Data. (model.LoanTransaction)
	compareLoanTransaction := cmp.Equal(createLoanTransactionObj.ID, getLoanTransactionObj.ID)
	
	if  compareLoanTransaction == false	{
		t.Errorf( "Created LoanTransaction object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLoanTransactionRequestResult := dao.GetAllLoanTransaction()

	if getAllLoanTransactionRequestResult.Success == false {
			t.Errorf(getAllLoanTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LoanTransaction success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLoanTransactionObj []model.LoanTransaction = getAllLoanTransactionRequestResult.Data. ([]model.LoanTransaction)
		
	equalLoanTransaction := cmp.Equal(createLoanTransactionObj.ID, getAllLoanTransactionObj[len(getAllLoanTransactionObj)-1].ID)
		
	if equalLoanTransaction == false {
		t.Errorf( "Created object is not equal to the last entry in LoanTransaction[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LoanTransaction
	// --------------------------------------------------------------	
	deleteLoanTransactionRequestResult := dao.DeleteLoanTransaction(uint64(createLoanTransactionObj.ID))

	if deleteLoanTransactionRequestResult.Success == false {
			t.Errorf(deleteLoanTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LoanTransaction success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLoanTransactionRequestResult = dao.GetLoanTransaction( uint64(createLoanTransactionObj.ID) )
	
	if getLoanTransactionRequestResult.Success == true {
		t.Errorf(getLoanTransactionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInvestmentPortfolioCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InvestmentPortfolio
	//----------------------------------------------------------------------------
	InvestmentPortfolioObj := model.InvestmentPortfolio                                                                                            {PortfolioCode:"test value for PortfolioCode",BaseCurrency:"test value for BaseCurrency",CreatedAt:new DateTime(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInvestmentPortfolioRequestResult := dao.CreateInvestmentPortfolio( InvestmentPortfolioObj )
	
	if createInvestmentPortfolioRequestResult.Success == false {
		t.Errorf(createInvestmentPortfolioRequestResult.Msg)
	} else {
		fmt.Println("Check Create InvestmentPortfolio success...")
	}
	
	createInvestmentPortfolioObj,_ := createInvestmentPortfolioRequestResult.Data. (model.InvestmentPortfolio)

	// --------------------------------------------------------------
	// Check InvestmentPortfolio Obj ID
	// --------------------------------------------------------------	
	if createInvestmentPortfolioObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InvestmentPortfolio" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInvestmentPortfolioRequestResult := dao.GetInvestmentPortfolio( uint64(createInvestmentPortfolioObj.ID) )
	
	if getInvestmentPortfolioRequestResult.Success == false {
		t.Errorf(getInvestmentPortfolioRequestResult.Msg)
	} else {
		fmt.Println("Check Get InvestmentPortfolio success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInvestmentPortfolioObj,_ := getInvestmentPortfolioRequestResult.Data. (model.InvestmentPortfolio)
	compareInvestmentPortfolio := cmp.Equal(createInvestmentPortfolioObj.ID, getInvestmentPortfolioObj.ID)
	
	if  compareInvestmentPortfolio == false	{
		t.Errorf( "Created InvestmentPortfolio object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInvestmentPortfolioRequestResult := dao.GetAllInvestmentPortfolio()

	if getAllInvestmentPortfolioRequestResult.Success == false {
			t.Errorf(getAllInvestmentPortfolioRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InvestmentPortfolio success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInvestmentPortfolioObj []model.InvestmentPortfolio = getAllInvestmentPortfolioRequestResult.Data. ([]model.InvestmentPortfolio)
		
	equalInvestmentPortfolio := cmp.Equal(createInvestmentPortfolioObj.ID, getAllInvestmentPortfolioObj[len(getAllInvestmentPortfolioObj)-1].ID)
		
	if equalInvestmentPortfolio == false {
		t.Errorf( "Created object is not equal to the last entry in InvestmentPortfolio[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InvestmentPortfolio
	// --------------------------------------------------------------	
	deleteInvestmentPortfolioRequestResult := dao.DeleteInvestmentPortfolio(uint64(createInvestmentPortfolioObj.ID))

	if deleteInvestmentPortfolioRequestResult.Success == false {
			t.Errorf(deleteInvestmentPortfolioRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InvestmentPortfolio success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInvestmentPortfolioRequestResult = dao.GetInvestmentPortfolio( uint64(createInvestmentPortfolioObj.ID) )
	
	if getInvestmentPortfolioRequestResult.Success == true {
		t.Errorf(getInvestmentPortfolioRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInvestmentAccountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InvestmentAccount
	//----------------------------------------------------------------------------
	InvestmentAccountObj := model.InvestmentAccount                                                                                            {AccountNumber:new AccountNumber(),BaseCurrency:"test value for BaseCurrency",Balance:new Money(),AccountType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInvestmentAccountRequestResult := dao.CreateInvestmentAccount( InvestmentAccountObj )
	
	if createInvestmentAccountRequestResult.Success == false {
		t.Errorf(createInvestmentAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Create InvestmentAccount success...")
	}
	
	createInvestmentAccountObj,_ := createInvestmentAccountRequestResult.Data. (model.InvestmentAccount)

	// --------------------------------------------------------------
	// Check InvestmentAccount Obj ID
	// --------------------------------------------------------------	
	if createInvestmentAccountObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InvestmentAccount" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInvestmentAccountRequestResult := dao.GetInvestmentAccount( uint64(createInvestmentAccountObj.ID) )
	
	if getInvestmentAccountRequestResult.Success == false {
		t.Errorf(getInvestmentAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Get InvestmentAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInvestmentAccountObj,_ := getInvestmentAccountRequestResult.Data. (model.InvestmentAccount)
	compareInvestmentAccount := cmp.Equal(createInvestmentAccountObj.ID, getInvestmentAccountObj.ID)
	
	if  compareInvestmentAccount == false	{
		t.Errorf( "Created InvestmentAccount object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInvestmentAccountRequestResult := dao.GetAllInvestmentAccount()

	if getAllInvestmentAccountRequestResult.Success == false {
			t.Errorf(getAllInvestmentAccountRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InvestmentAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInvestmentAccountObj []model.InvestmentAccount = getAllInvestmentAccountRequestResult.Data. ([]model.InvestmentAccount)
		
	equalInvestmentAccount := cmp.Equal(createInvestmentAccountObj.ID, getAllInvestmentAccountObj[len(getAllInvestmentAccountObj)-1].ID)
		
	if equalInvestmentAccount == false {
		t.Errorf( "Created object is not equal to the last entry in InvestmentAccount[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InvestmentAccount
	// --------------------------------------------------------------	
	deleteInvestmentAccountRequestResult := dao.DeleteInvestmentAccount(uint64(createInvestmentAccountObj.ID))

	if deleteInvestmentAccountRequestResult.Success == false {
			t.Errorf(deleteInvestmentAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InvestmentAccount success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInvestmentAccountRequestResult = dao.GetInvestmentAccount( uint64(createInvestmentAccountObj.ID) )
	
	if getInvestmentAccountRequestResult.Success == true {
		t.Errorf(getInvestmentAccountRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSecurityCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Security
	//----------------------------------------------------------------------------
	SecurityObj := model.Security                                                                                                                                            {Symbol:"test value for Symbol",Isin:"test value for Isin",Cusip:"test value for Cusip",Currency:"test value for Currency",SecurityType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSecurityRequestResult := dao.CreateSecurity( SecurityObj )
	
	if createSecurityRequestResult.Success == false {
		t.Errorf(createSecurityRequestResult.Msg)
	} else {
		fmt.Println("Check Create Security success...")
	}
	
	createSecurityObj,_ := createSecurityRequestResult.Data. (model.Security)

	// --------------------------------------------------------------
	// Check Security Obj ID
	// --------------------------------------------------------------	
	if createSecurityObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Security" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSecurityRequestResult := dao.GetSecurity( uint64(createSecurityObj.ID) )
	
	if getSecurityRequestResult.Success == false {
		t.Errorf(getSecurityRequestResult.Msg)
	} else {
		fmt.Println("Check Get Security success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSecurityObj,_ := getSecurityRequestResult.Data. (model.Security)
	compareSecurity := cmp.Equal(createSecurityObj.ID, getSecurityObj.ID)
	
	if  compareSecurity == false	{
		t.Errorf( "Created Security object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSecurityRequestResult := dao.GetAllSecurity()

	if getAllSecurityRequestResult.Success == false {
			t.Errorf(getAllSecurityRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Security success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSecurityObj []model.Security = getAllSecurityRequestResult.Data. ([]model.Security)
		
	equalSecurity := cmp.Equal(createSecurityObj.ID, getAllSecurityObj[len(getAllSecurityObj)-1].ID)
		
	if equalSecurity == false {
		t.Errorf( "Created object is not equal to the last entry in Security[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Security
	// --------------------------------------------------------------	
	deleteSecurityRequestResult := dao.DeleteSecurity(uint64(createSecurityObj.ID))

	if deleteSecurityRequestResult.Success == false {
			t.Errorf(deleteSecurityRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Security success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSecurityRequestResult = dao.GetSecurity( uint64(createSecurityObj.ID) )
	
	if getSecurityRequestResult.Success == true {
		t.Errorf(getSecurityRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPositionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Position
	//----------------------------------------------------------------------------
	PositionObj := model.Position                                                                                    {Quantity:"test value",AverageCost:new Money(),MarketValue:new Money()}

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


func TestTradeOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TradeOrder
	//----------------------------------------------------------------------------
	TradeOrderObj := model.TradeOrder                                                                                                                                                                                    {OrderId:"test value for OrderId",Quantity:"test value",LimitPrice:new Money(),PlacedAt:new DateTime(),Side:0,Type:0,Status:0,TimeInForce:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTradeOrderRequestResult := dao.CreateTradeOrder( TradeOrderObj )
	
	if createTradeOrderRequestResult.Success == false {
		t.Errorf(createTradeOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create TradeOrder success...")
	}
	
	createTradeOrderObj,_ := createTradeOrderRequestResult.Data. (model.TradeOrder)

	// --------------------------------------------------------------
	// Check TradeOrder Obj ID
	// --------------------------------------------------------------	
	if createTradeOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TradeOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTradeOrderRequestResult := dao.GetTradeOrder( uint64(createTradeOrderObj.ID) )
	
	if getTradeOrderRequestResult.Success == false {
		t.Errorf(getTradeOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get TradeOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTradeOrderObj,_ := getTradeOrderRequestResult.Data. (model.TradeOrder)
	compareTradeOrder := cmp.Equal(createTradeOrderObj.ID, getTradeOrderObj.ID)
	
	if  compareTradeOrder == false	{
		t.Errorf( "Created TradeOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTradeOrderRequestResult := dao.GetAllTradeOrder()

	if getAllTradeOrderRequestResult.Success == false {
			t.Errorf(getAllTradeOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TradeOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTradeOrderObj []model.TradeOrder = getAllTradeOrderRequestResult.Data. ([]model.TradeOrder)
		
	equalTradeOrder := cmp.Equal(createTradeOrderObj.ID, getAllTradeOrderObj[len(getAllTradeOrderObj)-1].ID)
		
	if equalTradeOrder == false {
		t.Errorf( "Created object is not equal to the last entry in TradeOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TradeOrder
	// --------------------------------------------------------------	
	deleteTradeOrderRequestResult := dao.DeleteTradeOrder(uint64(createTradeOrderObj.ID))

	if deleteTradeOrderRequestResult.Success == false {
			t.Errorf(deleteTradeOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TradeOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTradeOrderRequestResult = dao.GetTradeOrder( uint64(createTradeOrderObj.ID) )
	
	if getTradeOrderRequestResult.Success == true {
		t.Errorf(getTradeOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTradeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Trade
	//----------------------------------------------------------------------------
	TradeObj := model.Trade                                                                                                                                                            {ExecutedAt:new DateTime(),Quantity:"test value",Price:new Money(),Fees:new Money(),SettlementDate:time.Now()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTradeRequestResult := dao.CreateTrade( TradeObj )
	
	if createTradeRequestResult.Success == false {
		t.Errorf(createTradeRequestResult.Msg)
	} else {
		fmt.Println("Check Create Trade success...")
	}
	
	createTradeObj,_ := createTradeRequestResult.Data. (model.Trade)

	// --------------------------------------------------------------
	// Check Trade Obj ID
	// --------------------------------------------------------------	
	if createTradeObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Trade" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTradeRequestResult := dao.GetTrade( uint64(createTradeObj.ID) )
	
	if getTradeRequestResult.Success == false {
		t.Errorf(getTradeRequestResult.Msg)
	} else {
		fmt.Println("Check Get Trade success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTradeObj,_ := getTradeRequestResult.Data. (model.Trade)
	compareTrade := cmp.Equal(createTradeObj.ID, getTradeObj.ID)
	
	if  compareTrade == false	{
		t.Errorf( "Created Trade object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTradeRequestResult := dao.GetAllTrade()

	if getAllTradeRequestResult.Success == false {
			t.Errorf(getAllTradeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Trade success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTradeObj []model.Trade = getAllTradeRequestResult.Data. ([]model.Trade)
		
	equalTrade := cmp.Equal(createTradeObj.ID, getAllTradeObj[len(getAllTradeObj)-1].ID)
		
	if equalTrade == false {
		t.Errorf( "Created object is not equal to the last entry in Trade[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Trade
	// --------------------------------------------------------------	
	deleteTradeRequestResult := dao.DeleteTrade(uint64(createTradeObj.ID))

	if deleteTradeRequestResult.Success == false {
			t.Errorf(deleteTradeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Trade success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTradeRequestResult = dao.GetTrade( uint64(createTradeObj.ID) )
	
	if getTradeRequestResult.Success == true {
		t.Errorf(getTradeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestExchangeRateCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ExchangeRate
	//----------------------------------------------------------------------------
	ExchangeRateObj := model.ExchangeRate                                                                                                                                                                    {BaseCurrency:"test value for BaseCurrency",QuoteCurrency:"test value for QuoteCurrency",Rate:"test value",AsOf:new DateTime(),Source:"test value for Source"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createExchangeRateRequestResult := dao.CreateExchangeRate( ExchangeRateObj )
	
	if createExchangeRateRequestResult.Success == false {
		t.Errorf(createExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Check Create ExchangeRate success...")
	}
	
	createExchangeRateObj,_ := createExchangeRateRequestResult.Data. (model.ExchangeRate)

	// --------------------------------------------------------------
	// Check ExchangeRate Obj ID
	// --------------------------------------------------------------	
	if createExchangeRateObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ExchangeRate" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getExchangeRateRequestResult := dao.GetExchangeRate( uint64(createExchangeRateObj.ID) )
	
	if getExchangeRateRequestResult.Success == false {
		t.Errorf(getExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Check Get ExchangeRate success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getExchangeRateObj,_ := getExchangeRateRequestResult.Data. (model.ExchangeRate)
	compareExchangeRate := cmp.Equal(createExchangeRateObj.ID, getExchangeRateObj.ID)
	
	if  compareExchangeRate == false	{
		t.Errorf( "Created ExchangeRate object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllExchangeRateRequestResult := dao.GetAllExchangeRate()

	if getAllExchangeRateRequestResult.Success == false {
			t.Errorf(getAllExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ExchangeRate success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllExchangeRateObj []model.ExchangeRate = getAllExchangeRateRequestResult.Data. ([]model.ExchangeRate)
		
	equalExchangeRate := cmp.Equal(createExchangeRateObj.ID, getAllExchangeRateObj[len(getAllExchangeRateObj)-1].ID)
		
	if equalExchangeRate == false {
		t.Errorf( "Created object is not equal to the last entry in ExchangeRate[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ExchangeRate
	// --------------------------------------------------------------	
	deleteExchangeRateRequestResult := dao.DeleteExchangeRate(uint64(createExchangeRateObj.ID))

	if deleteExchangeRateRequestResult.Success == false {
			t.Errorf(deleteExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ExchangeRate success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getExchangeRateRequestResult = dao.GetExchangeRate( uint64(createExchangeRateObj.ID) )
	
	if getExchangeRateRequestResult.Success == true {
		t.Errorf(getExchangeRateRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

