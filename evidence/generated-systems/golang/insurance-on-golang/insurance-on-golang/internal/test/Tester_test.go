package test

import ( 
	"testing"
    dao "insurance-on-golang/internal/dao"
	"insurance-on-golang/internal/model"
	"insurance-on-golang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestInsurerCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Insurer
	//----------------------------------------------------------------------------
	InsurerObj := model.Insurer                                                                                                                                                            {Name:"test value for Name",LegalName:"test value for LegalName",DomicileCountry:"test value for DomicileCountry",NaicNumber:"test value for NaicNumber",Website:"test value for Website"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInsurerRequestResult := dao.CreateInsurer( InsurerObj )
	
	if createInsurerRequestResult.Success == false {
		t.Errorf(createInsurerRequestResult.Msg)
	} else {
		fmt.Println("Check Create Insurer success...")
	}
	
	createInsurerObj,_ := createInsurerRequestResult.Data. (model.Insurer)

	// --------------------------------------------------------------
	// Check Insurer Obj ID
	// --------------------------------------------------------------	
	if createInsurerObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Insurer" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInsurerRequestResult := dao.GetInsurer( uint64(createInsurerObj.ID) )
	
	if getInsurerRequestResult.Success == false {
		t.Errorf(getInsurerRequestResult.Msg)
	} else {
		fmt.Println("Check Get Insurer success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInsurerObj,_ := getInsurerRequestResult.Data. (model.Insurer)
	compareInsurer := cmp.Equal(createInsurerObj.ID, getInsurerObj.ID)
	
	if  compareInsurer == false	{
		t.Errorf( "Created Insurer object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInsurerRequestResult := dao.GetAllInsurer()

	if getAllInsurerRequestResult.Success == false {
			t.Errorf(getAllInsurerRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Insurer success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInsurerObj []model.Insurer = getAllInsurerRequestResult.Data. ([]model.Insurer)
		
	equalInsurer := cmp.Equal(createInsurerObj.ID, getAllInsurerObj[len(getAllInsurerObj)-1].ID)
		
	if equalInsurer == false {
		t.Errorf( "Created object is not equal to the last entry in Insurer[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Insurer
	// --------------------------------------------------------------	
	deleteInsurerRequestResult := dao.DeleteInsurer(uint64(createInsurerObj.ID))

	if deleteInsurerRequestResult.Success == false {
			t.Errorf(deleteInsurerRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Insurer success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInsurerRequestResult = dao.GetInsurer( uint64(createInsurerObj.ID) )
	
	if getInsurerRequestResult.Success == true {
		t.Errorf(getInsurerRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInsuranceProductCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InsuranceProduct
	//----------------------------------------------------------------------------
	InsuranceProductObj := model.InsuranceProduct                                                                            {Name:"test value for Name",ProductCode:"test value for ProductCode",LineOfBusiness:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInsuranceProductRequestResult := dao.CreateInsuranceProduct( InsuranceProductObj )
	
	if createInsuranceProductRequestResult.Success == false {
		t.Errorf(createInsuranceProductRequestResult.Msg)
	} else {
		fmt.Println("Check Create InsuranceProduct success...")
	}
	
	createInsuranceProductObj,_ := createInsuranceProductRequestResult.Data. (model.InsuranceProduct)

	// --------------------------------------------------------------
	// Check InsuranceProduct Obj ID
	// --------------------------------------------------------------	
	if createInsuranceProductObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InsuranceProduct" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInsuranceProductRequestResult := dao.GetInsuranceProduct( uint64(createInsuranceProductObj.ID) )
	
	if getInsuranceProductRequestResult.Success == false {
		t.Errorf(getInsuranceProductRequestResult.Msg)
	} else {
		fmt.Println("Check Get InsuranceProduct success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInsuranceProductObj,_ := getInsuranceProductRequestResult.Data. (model.InsuranceProduct)
	compareInsuranceProduct := cmp.Equal(createInsuranceProductObj.ID, getInsuranceProductObj.ID)
	
	if  compareInsuranceProduct == false	{
		t.Errorf( "Created InsuranceProduct object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInsuranceProductRequestResult := dao.GetAllInsuranceProduct()

	if getAllInsuranceProductRequestResult.Success == false {
			t.Errorf(getAllInsuranceProductRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InsuranceProduct success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInsuranceProductObj []model.InsuranceProduct = getAllInsuranceProductRequestResult.Data. ([]model.InsuranceProduct)
		
	equalInsuranceProduct := cmp.Equal(createInsuranceProductObj.ID, getAllInsuranceProductObj[len(getAllInsuranceProductObj)-1].ID)
		
	if equalInsuranceProduct == false {
		t.Errorf( "Created object is not equal to the last entry in InsuranceProduct[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InsuranceProduct
	// --------------------------------------------------------------	
	deleteInsuranceProductRequestResult := dao.DeleteInsuranceProduct(uint64(createInsuranceProductObj.ID))

	if deleteInsuranceProductRequestResult.Success == false {
			t.Errorf(deleteInsuranceProductRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InsuranceProduct success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInsuranceProductRequestResult = dao.GetInsuranceProduct( uint64(createInsuranceProductObj.ID) )
	
	if getInsuranceProductRequestResult.Success == true {
		t.Errorf(getInsuranceProductRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCoverageDefinitionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CoverageDefinition
	//----------------------------------------------------------------------------
	CoverageDefinitionObj := model.CoverageDefinition                                                                                                            {Name:"test value for Name",DefaultLimit:new Money(),DefaultDeductible:new Money(),AsMandatory:true,CoverageType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCoverageDefinitionRequestResult := dao.CreateCoverageDefinition( CoverageDefinitionObj )
	
	if createCoverageDefinitionRequestResult.Success == false {
		t.Errorf(createCoverageDefinitionRequestResult.Msg)
	} else {
		fmt.Println("Check Create CoverageDefinition success...")
	}
	
	createCoverageDefinitionObj,_ := createCoverageDefinitionRequestResult.Data. (model.CoverageDefinition)

	// --------------------------------------------------------------
	// Check CoverageDefinition Obj ID
	// --------------------------------------------------------------	
	if createCoverageDefinitionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CoverageDefinition" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCoverageDefinitionRequestResult := dao.GetCoverageDefinition( uint64(createCoverageDefinitionObj.ID) )
	
	if getCoverageDefinitionRequestResult.Success == false {
		t.Errorf(getCoverageDefinitionRequestResult.Msg)
	} else {
		fmt.Println("Check Get CoverageDefinition success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCoverageDefinitionObj,_ := getCoverageDefinitionRequestResult.Data. (model.CoverageDefinition)
	compareCoverageDefinition := cmp.Equal(createCoverageDefinitionObj.ID, getCoverageDefinitionObj.ID)
	
	if  compareCoverageDefinition == false	{
		t.Errorf( "Created CoverageDefinition object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCoverageDefinitionRequestResult := dao.GetAllCoverageDefinition()

	if getAllCoverageDefinitionRequestResult.Success == false {
			t.Errorf(getAllCoverageDefinitionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CoverageDefinition success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCoverageDefinitionObj []model.CoverageDefinition = getAllCoverageDefinitionRequestResult.Data. ([]model.CoverageDefinition)
		
	equalCoverageDefinition := cmp.Equal(createCoverageDefinitionObj.ID, getAllCoverageDefinitionObj[len(getAllCoverageDefinitionObj)-1].ID)
		
	if equalCoverageDefinition == false {
		t.Errorf( "Created object is not equal to the last entry in CoverageDefinition[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CoverageDefinition
	// --------------------------------------------------------------	
	deleteCoverageDefinitionRequestResult := dao.DeleteCoverageDefinition(uint64(createCoverageDefinitionObj.ID))

	if deleteCoverageDefinitionRequestResult.Success == false {
			t.Errorf(deleteCoverageDefinitionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CoverageDefinition success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCoverageDefinitionRequestResult = dao.GetCoverageDefinition( uint64(createCoverageDefinitionObj.ID) )
	
	if getCoverageDefinitionRequestResult.Success == true {
		t.Errorf(getCoverageDefinitionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDistributorCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Distributor
	//----------------------------------------------------------------------------
	DistributorObj := model.Distributor                                                                                                            {Name:"test value for Name",LicenseNumber:"test value for LicenseNumber",Region:"test value for Region",DistributorType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDistributorRequestResult := dao.CreateDistributor( DistributorObj )
	
	if createDistributorRequestResult.Success == false {
		t.Errorf(createDistributorRequestResult.Msg)
	} else {
		fmt.Println("Check Create Distributor success...")
	}
	
	createDistributorObj,_ := createDistributorRequestResult.Data. (model.Distributor)

	// --------------------------------------------------------------
	// Check Distributor Obj ID
	// --------------------------------------------------------------	
	if createDistributorObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Distributor" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDistributorRequestResult := dao.GetDistributor( uint64(createDistributorObj.ID) )
	
	if getDistributorRequestResult.Success == false {
		t.Errorf(getDistributorRequestResult.Msg)
	} else {
		fmt.Println("Check Get Distributor success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDistributorObj,_ := getDistributorRequestResult.Data. (model.Distributor)
	compareDistributor := cmp.Equal(createDistributorObj.ID, getDistributorObj.ID)
	
	if  compareDistributor == false	{
		t.Errorf( "Created Distributor object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDistributorRequestResult := dao.GetAllDistributor()

	if getAllDistributorRequestResult.Success == false {
			t.Errorf(getAllDistributorRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Distributor success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDistributorObj []model.Distributor = getAllDistributorRequestResult.Data. ([]model.Distributor)
		
	equalDistributor := cmp.Equal(createDistributorObj.ID, getAllDistributorObj[len(getAllDistributorObj)-1].ID)
		
	if equalDistributor == false {
		t.Errorf( "Created object is not equal to the last entry in Distributor[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Distributor
	// --------------------------------------------------------------	
	deleteDistributorRequestResult := dao.DeleteDistributor(uint64(createDistributorObj.ID))

	if deleteDistributorRequestResult.Success == false {
			t.Errorf(deleteDistributorRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Distributor success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDistributorRequestResult = dao.GetDistributor( uint64(createDistributorObj.ID) )
	
	if getDistributorRequestResult.Success == true {
		t.Errorf(getDistributorRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAgentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Agent
	//----------------------------------------------------------------------------
	AgentObj := model.Agent                                                                                                            {FirstName:"test value for FirstName",LastName:"test value for LastName",LicenseId:"test value for LicenseId",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAgentRequestResult := dao.CreateAgent( AgentObj )
	
	if createAgentRequestResult.Success == false {
		t.Errorf(createAgentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Agent success...")
	}
	
	createAgentObj,_ := createAgentRequestResult.Data. (model.Agent)

	// --------------------------------------------------------------
	// Check Agent Obj ID
	// --------------------------------------------------------------	
	if createAgentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Agent" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAgentRequestResult := dao.GetAgent( uint64(createAgentObj.ID) )
	
	if getAgentRequestResult.Success == false {
		t.Errorf(getAgentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Agent success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAgentObj,_ := getAgentRequestResult.Data. (model.Agent)
	compareAgent := cmp.Equal(createAgentObj.ID, getAgentObj.ID)
	
	if  compareAgent == false	{
		t.Errorf( "Created Agent object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAgentRequestResult := dao.GetAllAgent()

	if getAllAgentRequestResult.Success == false {
			t.Errorf(getAllAgentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Agent success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAgentObj []model.Agent = getAllAgentRequestResult.Data. ([]model.Agent)
		
	equalAgent := cmp.Equal(createAgentObj.ID, getAllAgentObj[len(getAllAgentObj)-1].ID)
		
	if equalAgent == false {
		t.Errorf( "Created object is not equal to the last entry in Agent[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Agent
	// --------------------------------------------------------------	
	deleteAgentRequestResult := dao.DeleteAgent(uint64(createAgentObj.ID))

	if deleteAgentRequestResult.Success == false {
			t.Errorf(deleteAgentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Agent success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAgentRequestResult = dao.GetAgent( uint64(createAgentObj.ID) )
	
	if getAgentRequestResult.Success == true {
		t.Errorf(getAgentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCustomerCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Customer
	//----------------------------------------------------------------------------
	CustomerObj := model.Customer                                                                                                                                                                                                                    {FirstName:"test value for FirstName",LastName:"test value for LastName",OrganizationName:"test value for OrganizationName",TaxId:"test value for TaxId",DateOfBirth:time.Now(),PrimaryAddress:new Address(),CustomerType:0}

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


func TestApplicationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Application
	//----------------------------------------------------------------------------
	ApplicationObj := model.Application                                                                                                    {ApplicationNumber:"test value for ApplicationNumber",SubmissionDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createApplicationRequestResult := dao.CreateApplication( ApplicationObj )
	
	if createApplicationRequestResult.Success == false {
		t.Errorf(createApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Application success...")
	}
	
	createApplicationObj,_ := createApplicationRequestResult.Data. (model.Application)

	// --------------------------------------------------------------
	// Check Application Obj ID
	// --------------------------------------------------------------	
	if createApplicationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Application" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getApplicationRequestResult := dao.GetApplication( uint64(createApplicationObj.ID) )
	
	if getApplicationRequestResult.Success == false {
		t.Errorf(getApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Application success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getApplicationObj,_ := getApplicationRequestResult.Data. (model.Application)
	compareApplication := cmp.Equal(createApplicationObj.ID, getApplicationObj.ID)
	
	if  compareApplication == false	{
		t.Errorf( "Created Application object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllApplicationRequestResult := dao.GetAllApplication()

	if getAllApplicationRequestResult.Success == false {
			t.Errorf(getAllApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Application success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllApplicationObj []model.Application = getAllApplicationRequestResult.Data. ([]model.Application)
		
	equalApplication := cmp.Equal(createApplicationObj.ID, getAllApplicationObj[len(getAllApplicationObj)-1].ID)
		
	if equalApplication == false {
		t.Errorf( "Created object is not equal to the last entry in Application[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Application
	// --------------------------------------------------------------	
	deleteApplicationRequestResult := dao.DeleteApplication(uint64(createApplicationObj.ID))

	if deleteApplicationRequestResult.Success == false {
			t.Errorf(deleteApplicationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Application success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getApplicationRequestResult = dao.GetApplication( uint64(createApplicationObj.ID) )
	
	if getApplicationRequestResult.Success == true {
		t.Errorf(getApplicationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestQuoteCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Quote
	//----------------------------------------------------------------------------
	QuoteObj := model.Quote                                                                                                                                    {QuoteNumber:"test value for QuoteNumber",TotalPremium:new Money(),RatingDate:time.Now(),AsBound:true}

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


func TestUnderwritingDecisionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for UnderwritingDecision
	//----------------------------------------------------------------------------
	UnderwritingDecisionObj := model.UnderwritingDecision                                                                                                    {Notes:"test value for Notes",DecisionDate:time.Now(),Decision:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createUnderwritingDecisionRequestResult := dao.CreateUnderwritingDecision( UnderwritingDecisionObj )
	
	if createUnderwritingDecisionRequestResult.Success == false {
		t.Errorf(createUnderwritingDecisionRequestResult.Msg)
	} else {
		fmt.Println("Check Create UnderwritingDecision success...")
	}
	
	createUnderwritingDecisionObj,_ := createUnderwritingDecisionRequestResult.Data. (model.UnderwritingDecision)

	// --------------------------------------------------------------
	// Check UnderwritingDecision Obj ID
	// --------------------------------------------------------------	
	if createUnderwritingDecisionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for UnderwritingDecision" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getUnderwritingDecisionRequestResult := dao.GetUnderwritingDecision( uint64(createUnderwritingDecisionObj.ID) )
	
	if getUnderwritingDecisionRequestResult.Success == false {
		t.Errorf(getUnderwritingDecisionRequestResult.Msg)
	} else {
		fmt.Println("Check Get UnderwritingDecision success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getUnderwritingDecisionObj,_ := getUnderwritingDecisionRequestResult.Data. (model.UnderwritingDecision)
	compareUnderwritingDecision := cmp.Equal(createUnderwritingDecisionObj.ID, getUnderwritingDecisionObj.ID)
	
	if  compareUnderwritingDecision == false	{
		t.Errorf( "Created UnderwritingDecision object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllUnderwritingDecisionRequestResult := dao.GetAllUnderwritingDecision()

	if getAllUnderwritingDecisionRequestResult.Success == false {
			t.Errorf(getAllUnderwritingDecisionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll UnderwritingDecision success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllUnderwritingDecisionObj []model.UnderwritingDecision = getAllUnderwritingDecisionRequestResult.Data. ([]model.UnderwritingDecision)
		
	equalUnderwritingDecision := cmp.Equal(createUnderwritingDecisionObj.ID, getAllUnderwritingDecisionObj[len(getAllUnderwritingDecisionObj)-1].ID)
		
	if equalUnderwritingDecision == false {
		t.Errorf( "Created object is not equal to the last entry in UnderwritingDecision[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for UnderwritingDecision
	// --------------------------------------------------------------	
	deleteUnderwritingDecisionRequestResult := dao.DeleteUnderwritingDecision(uint64(createUnderwritingDecisionObj.ID))

	if deleteUnderwritingDecisionRequestResult.Success == false {
			t.Errorf(deleteUnderwritingDecisionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion UnderwritingDecision success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getUnderwritingDecisionRequestResult = dao.GetUnderwritingDecision( uint64(createUnderwritingDecisionObj.ID) )
	
	if getUnderwritingDecisionRequestResult.Success == true {
		t.Errorf(getUnderwritingDecisionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestUnderwriterCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Underwriter
	//----------------------------------------------------------------------------
	UnderwriterObj := model.Underwriter                                                                                                            {FirstName:"test value for FirstName",LastName:"test value for LastName",EmployeeId:"test value for EmployeeId",AuthorityLimit:new Money()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createUnderwriterRequestResult := dao.CreateUnderwriter( UnderwriterObj )
	
	if createUnderwriterRequestResult.Success == false {
		t.Errorf(createUnderwriterRequestResult.Msg)
	} else {
		fmt.Println("Check Create Underwriter success...")
	}
	
	createUnderwriterObj,_ := createUnderwriterRequestResult.Data. (model.Underwriter)

	// --------------------------------------------------------------
	// Check Underwriter Obj ID
	// --------------------------------------------------------------	
	if createUnderwriterObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Underwriter" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getUnderwriterRequestResult := dao.GetUnderwriter( uint64(createUnderwriterObj.ID) )
	
	if getUnderwriterRequestResult.Success == false {
		t.Errorf(getUnderwriterRequestResult.Msg)
	} else {
		fmt.Println("Check Get Underwriter success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getUnderwriterObj,_ := getUnderwriterRequestResult.Data. (model.Underwriter)
	compareUnderwriter := cmp.Equal(createUnderwriterObj.ID, getUnderwriterObj.ID)
	
	if  compareUnderwriter == false	{
		t.Errorf( "Created Underwriter object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllUnderwriterRequestResult := dao.GetAllUnderwriter()

	if getAllUnderwriterRequestResult.Success == false {
			t.Errorf(getAllUnderwriterRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Underwriter success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllUnderwriterObj []model.Underwriter = getAllUnderwriterRequestResult.Data. ([]model.Underwriter)
		
	equalUnderwriter := cmp.Equal(createUnderwriterObj.ID, getAllUnderwriterObj[len(getAllUnderwriterObj)-1].ID)
		
	if equalUnderwriter == false {
		t.Errorf( "Created object is not equal to the last entry in Underwriter[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Underwriter
	// --------------------------------------------------------------	
	deleteUnderwriterRequestResult := dao.DeleteUnderwriter(uint64(createUnderwriterObj.ID))

	if deleteUnderwriterRequestResult.Success == false {
			t.Errorf(deleteUnderwriterRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Underwriter success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getUnderwriterRequestResult = dao.GetUnderwriter( uint64(createUnderwriterObj.ID) )
	
	if getUnderwriterRequestResult.Success == true {
		t.Errorf(getUnderwriterRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPolicyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Policy
	//----------------------------------------------------------------------------
	PolicyObj := model.Policy                                                                            {PolicyNumber:new PolicyNumber(),EffectivePeriod:new DateRange(),TotalPremium:new Money(),Status:0,PaymentPlan:0}

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


func TestEndorsementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Endorsement
	//----------------------------------------------------------------------------
	EndorsementObj := model.Endorsement                                                                                                                    {EndorsementNumber:"test value for EndorsementNumber",EffectiveDate:time.Now(),Description:"test value for Description"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEndorsementRequestResult := dao.CreateEndorsement( EndorsementObj )
	
	if createEndorsementRequestResult.Success == false {
		t.Errorf(createEndorsementRequestResult.Msg)
	} else {
		fmt.Println("Check Create Endorsement success...")
	}
	
	createEndorsementObj,_ := createEndorsementRequestResult.Data. (model.Endorsement)

	// --------------------------------------------------------------
	// Check Endorsement Obj ID
	// --------------------------------------------------------------	
	if createEndorsementObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Endorsement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEndorsementRequestResult := dao.GetEndorsement( uint64(createEndorsementObj.ID) )
	
	if getEndorsementRequestResult.Success == false {
		t.Errorf(getEndorsementRequestResult.Msg)
	} else {
		fmt.Println("Check Get Endorsement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEndorsementObj,_ := getEndorsementRequestResult.Data. (model.Endorsement)
	compareEndorsement := cmp.Equal(createEndorsementObj.ID, getEndorsementObj.ID)
	
	if  compareEndorsement == false	{
		t.Errorf( "Created Endorsement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEndorsementRequestResult := dao.GetAllEndorsement()

	if getAllEndorsementRequestResult.Success == false {
			t.Errorf(getAllEndorsementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Endorsement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEndorsementObj []model.Endorsement = getAllEndorsementRequestResult.Data. ([]model.Endorsement)
		
	equalEndorsement := cmp.Equal(createEndorsementObj.ID, getAllEndorsementObj[len(getAllEndorsementObj)-1].ID)
		
	if equalEndorsement == false {
		t.Errorf( "Created object is not equal to the last entry in Endorsement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Endorsement
	// --------------------------------------------------------------	
	deleteEndorsementRequestResult := dao.DeleteEndorsement(uint64(createEndorsementObj.ID))

	if deleteEndorsementRequestResult.Success == false {
			t.Errorf(deleteEndorsementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Endorsement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEndorsementRequestResult = dao.GetEndorsement( uint64(createEndorsementObj.ID) )
	
	if getEndorsementRequestResult.Success == true {
		t.Errorf(getEndorsementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPolicyCoverageCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PolicyCoverage
	//----------------------------------------------------------------------------
	PolicyCoverageObj := model.PolicyCoverage                                                            {Limit:new Money(),Deductible:new Money(),Premium:new Money(),CoverageType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPolicyCoverageRequestResult := dao.CreatePolicyCoverage( PolicyCoverageObj )
	
	if createPolicyCoverageRequestResult.Success == false {
		t.Errorf(createPolicyCoverageRequestResult.Msg)
	} else {
		fmt.Println("Check Create PolicyCoverage success...")
	}
	
	createPolicyCoverageObj,_ := createPolicyCoverageRequestResult.Data. (model.PolicyCoverage)

	// --------------------------------------------------------------
	// Check PolicyCoverage Obj ID
	// --------------------------------------------------------------	
	if createPolicyCoverageObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PolicyCoverage" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPolicyCoverageRequestResult := dao.GetPolicyCoverage( uint64(createPolicyCoverageObj.ID) )
	
	if getPolicyCoverageRequestResult.Success == false {
		t.Errorf(getPolicyCoverageRequestResult.Msg)
	} else {
		fmt.Println("Check Get PolicyCoverage success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPolicyCoverageObj,_ := getPolicyCoverageRequestResult.Data. (model.PolicyCoverage)
	comparePolicyCoverage := cmp.Equal(createPolicyCoverageObj.ID, getPolicyCoverageObj.ID)
	
	if  comparePolicyCoverage == false	{
		t.Errorf( "Created PolicyCoverage object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPolicyCoverageRequestResult := dao.GetAllPolicyCoverage()

	if getAllPolicyCoverageRequestResult.Success == false {
			t.Errorf(getAllPolicyCoverageRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PolicyCoverage success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPolicyCoverageObj []model.PolicyCoverage = getAllPolicyCoverageRequestResult.Data. ([]model.PolicyCoverage)
		
	equalPolicyCoverage := cmp.Equal(createPolicyCoverageObj.ID, getAllPolicyCoverageObj[len(getAllPolicyCoverageObj)-1].ID)
		
	if equalPolicyCoverage == false {
		t.Errorf( "Created object is not equal to the last entry in PolicyCoverage[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PolicyCoverage
	// --------------------------------------------------------------	
	deletePolicyCoverageRequestResult := dao.DeletePolicyCoverage(uint64(createPolicyCoverageObj.ID))

	if deletePolicyCoverageRequestResult.Success == false {
			t.Errorf(deletePolicyCoverageRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PolicyCoverage success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPolicyCoverageRequestResult = dao.GetPolicyCoverage( uint64(createPolicyCoverageObj.ID) )
	
	if getPolicyCoverageRequestResult.Success == true {
		t.Errorf(getPolicyCoverageRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInsuredObjectCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InsuredObject
	//----------------------------------------------------------------------------
	InsuredObjectObj := model.InsuredObject                                                                                            {Description:"test value for Description",SerialOrId:"test value for SerialOrId",PrimaryAddress:new Address(),ObjectType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInsuredObjectRequestResult := dao.CreateInsuredObject( InsuredObjectObj )
	
	if createInsuredObjectRequestResult.Success == false {
		t.Errorf(createInsuredObjectRequestResult.Msg)
	} else {
		fmt.Println("Check Create InsuredObject success...")
	}
	
	createInsuredObjectObj,_ := createInsuredObjectRequestResult.Data. (model.InsuredObject)

	// --------------------------------------------------------------
	// Check InsuredObject Obj ID
	// --------------------------------------------------------------	
	if createInsuredObjectObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InsuredObject" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInsuredObjectRequestResult := dao.GetInsuredObject( uint64(createInsuredObjectObj.ID) )
	
	if getInsuredObjectRequestResult.Success == false {
		t.Errorf(getInsuredObjectRequestResult.Msg)
	} else {
		fmt.Println("Check Get InsuredObject success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInsuredObjectObj,_ := getInsuredObjectRequestResult.Data. (model.InsuredObject)
	compareInsuredObject := cmp.Equal(createInsuredObjectObj.ID, getInsuredObjectObj.ID)
	
	if  compareInsuredObject == false	{
		t.Errorf( "Created InsuredObject object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInsuredObjectRequestResult := dao.GetAllInsuredObject()

	if getAllInsuredObjectRequestResult.Success == false {
			t.Errorf(getAllInsuredObjectRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InsuredObject success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInsuredObjectObj []model.InsuredObject = getAllInsuredObjectRequestResult.Data. ([]model.InsuredObject)
		
	equalInsuredObject := cmp.Equal(createInsuredObjectObj.ID, getAllInsuredObjectObj[len(getAllInsuredObjectObj)-1].ID)
		
	if equalInsuredObject == false {
		t.Errorf( "Created object is not equal to the last entry in InsuredObject[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InsuredObject
	// --------------------------------------------------------------	
	deleteInsuredObjectRequestResult := dao.DeleteInsuredObject(uint64(createInsuredObjectObj.ID))

	if deleteInsuredObjectRequestResult.Success == false {
			t.Errorf(deleteInsuredObjectRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InsuredObject success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInsuredObjectRequestResult = dao.GetInsuredObject( uint64(createInsuredObjectObj.ID) )
	
	if getInsuredObjectRequestResult.Success == true {
		t.Errorf(getInsuredObjectRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBeneficiaryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Beneficiary
	//----------------------------------------------------------------------------
	BeneficiaryObj := model.Beneficiary                                                            {Name:"test value for Name",Share:new Percentage(),Relationship:0}

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


func TestBillingAccountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BillingAccount
	//----------------------------------------------------------------------------
	BillingAccountObj := model.BillingAccount                                                            {AccountNumber:"test value for AccountNumber",Balance:new Money(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBillingAccountRequestResult := dao.CreateBillingAccount( BillingAccountObj )
	
	if createBillingAccountRequestResult.Success == false {
		t.Errorf(createBillingAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Create BillingAccount success...")
	}
	
	createBillingAccountObj,_ := createBillingAccountRequestResult.Data. (model.BillingAccount)

	// --------------------------------------------------------------
	// Check BillingAccount Obj ID
	// --------------------------------------------------------------	
	if createBillingAccountObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BillingAccount" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBillingAccountRequestResult := dao.GetBillingAccount( uint64(createBillingAccountObj.ID) )
	
	if getBillingAccountRequestResult.Success == false {
		t.Errorf(getBillingAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Get BillingAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBillingAccountObj,_ := getBillingAccountRequestResult.Data. (model.BillingAccount)
	compareBillingAccount := cmp.Equal(createBillingAccountObj.ID, getBillingAccountObj.ID)
	
	if  compareBillingAccount == false	{
		t.Errorf( "Created BillingAccount object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBillingAccountRequestResult := dao.GetAllBillingAccount()

	if getAllBillingAccountRequestResult.Success == false {
			t.Errorf(getAllBillingAccountRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BillingAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBillingAccountObj []model.BillingAccount = getAllBillingAccountRequestResult.Data. ([]model.BillingAccount)
		
	equalBillingAccount := cmp.Equal(createBillingAccountObj.ID, getAllBillingAccountObj[len(getAllBillingAccountObj)-1].ID)
		
	if equalBillingAccount == false {
		t.Errorf( "Created object is not equal to the last entry in BillingAccount[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BillingAccount
	// --------------------------------------------------------------	
	deleteBillingAccountRequestResult := dao.DeleteBillingAccount(uint64(createBillingAccountObj.ID))

	if deleteBillingAccountRequestResult.Success == false {
			t.Errorf(deleteBillingAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BillingAccount success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBillingAccountRequestResult = dao.GetBillingAccount( uint64(createBillingAccountObj.ID) )
	
	if getBillingAccountRequestResult.Success == true {
		t.Errorf(getBillingAccountRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInvoiceCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Invoice
	//----------------------------------------------------------------------------
	InvoiceObj := model.Invoice                                                                                                                    {InvoiceNumber:"test value for InvoiceNumber",DueDate:time.Now(),TotalDue:new Money(),Status:0}

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
	PaymentObj := model.Payment                                                                                                                                    {PaymentReference:"test value for PaymentReference",Amount:new Money(),PaymentDate:time.Now(),Method:0,Status:0}

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


func TestClaimCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Claim
	//----------------------------------------------------------------------------
	ClaimObj := model.Claim                                                                                                                                                                                            {ClaimNumber:new ClaimNumber(),NoticeDate:time.Now(),LossDate:time.Now(),ReportedBy:"test value for ReportedBy",Status:0,LossCause:0}

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


func TestIncidentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Incident
	//----------------------------------------------------------------------------
	IncidentObj := model.Incident                                                            {Location:new Address(),Description:"test value for Description",IncidentType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createIncidentRequestResult := dao.CreateIncident( IncidentObj )
	
	if createIncidentRequestResult.Success == false {
		t.Errorf(createIncidentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Incident success...")
	}
	
	createIncidentObj,_ := createIncidentRequestResult.Data. (model.Incident)

	// --------------------------------------------------------------
	// Check Incident Obj ID
	// --------------------------------------------------------------	
	if createIncidentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Incident" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getIncidentRequestResult := dao.GetIncident( uint64(createIncidentObj.ID) )
	
	if getIncidentRequestResult.Success == false {
		t.Errorf(getIncidentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Incident success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getIncidentObj,_ := getIncidentRequestResult.Data. (model.Incident)
	compareIncident := cmp.Equal(createIncidentObj.ID, getIncidentObj.ID)
	
	if  compareIncident == false	{
		t.Errorf( "Created Incident object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllIncidentRequestResult := dao.GetAllIncident()

	if getAllIncidentRequestResult.Success == false {
			t.Errorf(getAllIncidentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Incident success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllIncidentObj []model.Incident = getAllIncidentRequestResult.Data. ([]model.Incident)
		
	equalIncident := cmp.Equal(createIncidentObj.ID, getAllIncidentObj[len(getAllIncidentObj)-1].ID)
		
	if equalIncident == false {
		t.Errorf( "Created object is not equal to the last entry in Incident[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Incident
	// --------------------------------------------------------------	
	deleteIncidentRequestResult := dao.DeleteIncident(uint64(createIncidentObj.ID))

	if deleteIncidentRequestResult.Success == false {
			t.Errorf(deleteIncidentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Incident success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getIncidentRequestResult = dao.GetIncident( uint64(createIncidentObj.ID) )
	
	if getIncidentRequestResult.Success == true {
		t.Errorf(getIncidentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestExposureCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Exposure
	//----------------------------------------------------------------------------
	ExposureObj := model.Exposure                            {ExposureType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createExposureRequestResult := dao.CreateExposure( ExposureObj )
	
	if createExposureRequestResult.Success == false {
		t.Errorf(createExposureRequestResult.Msg)
	} else {
		fmt.Println("Check Create Exposure success...")
	}
	
	createExposureObj,_ := createExposureRequestResult.Data. (model.Exposure)

	// --------------------------------------------------------------
	// Check Exposure Obj ID
	// --------------------------------------------------------------	
	if createExposureObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Exposure" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getExposureRequestResult := dao.GetExposure( uint64(createExposureObj.ID) )
	
	if getExposureRequestResult.Success == false {
		t.Errorf(getExposureRequestResult.Msg)
	} else {
		fmt.Println("Check Get Exposure success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getExposureObj,_ := getExposureRequestResult.Data. (model.Exposure)
	compareExposure := cmp.Equal(createExposureObj.ID, getExposureObj.ID)
	
	if  compareExposure == false	{
		t.Errorf( "Created Exposure object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllExposureRequestResult := dao.GetAllExposure()

	if getAllExposureRequestResult.Success == false {
			t.Errorf(getAllExposureRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Exposure success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllExposureObj []model.Exposure = getAllExposureRequestResult.Data. ([]model.Exposure)
		
	equalExposure := cmp.Equal(createExposureObj.ID, getAllExposureObj[len(getAllExposureObj)-1].ID)
		
	if equalExposure == false {
		t.Errorf( "Created object is not equal to the last entry in Exposure[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Exposure
	// --------------------------------------------------------------	
	deleteExposureRequestResult := dao.DeleteExposure(uint64(createExposureObj.ID))

	if deleteExposureRequestResult.Success == false {
			t.Errorf(deleteExposureRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Exposure success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getExposureRequestResult = dao.GetExposure( uint64(createExposureObj.ID) )
	
	if getExposureRequestResult.Success == true {
		t.Errorf(getExposureRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAdjusterCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Adjuster
	//----------------------------------------------------------------------------
	AdjusterObj := model.Adjuster                                                                                                            {FirstName:"test value for FirstName",LastName:"test value for LastName",LicenseNumber:"test value for LicenseNumber",AdjusterType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAdjusterRequestResult := dao.CreateAdjuster( AdjusterObj )
	
	if createAdjusterRequestResult.Success == false {
		t.Errorf(createAdjusterRequestResult.Msg)
	} else {
		fmt.Println("Check Create Adjuster success...")
	}
	
	createAdjusterObj,_ := createAdjusterRequestResult.Data. (model.Adjuster)

	// --------------------------------------------------------------
	// Check Adjuster Obj ID
	// --------------------------------------------------------------	
	if createAdjusterObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Adjuster" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAdjusterRequestResult := dao.GetAdjuster( uint64(createAdjusterObj.ID) )
	
	if getAdjusterRequestResult.Success == false {
		t.Errorf(getAdjusterRequestResult.Msg)
	} else {
		fmt.Println("Check Get Adjuster success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAdjusterObj,_ := getAdjusterRequestResult.Data. (model.Adjuster)
	compareAdjuster := cmp.Equal(createAdjusterObj.ID, getAdjusterObj.ID)
	
	if  compareAdjuster == false	{
		t.Errorf( "Created Adjuster object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAdjusterRequestResult := dao.GetAllAdjuster()

	if getAllAdjusterRequestResult.Success == false {
			t.Errorf(getAllAdjusterRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Adjuster success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAdjusterObj []model.Adjuster = getAllAdjusterRequestResult.Data. ([]model.Adjuster)
		
	equalAdjuster := cmp.Equal(createAdjusterObj.ID, getAllAdjusterObj[len(getAllAdjusterObj)-1].ID)
		
	if equalAdjuster == false {
		t.Errorf( "Created object is not equal to the last entry in Adjuster[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Adjuster
	// --------------------------------------------------------------	
	deleteAdjusterRequestResult := dao.DeleteAdjuster(uint64(createAdjusterObj.ID))

	if deleteAdjusterRequestResult.Success == false {
			t.Errorf(deleteAdjusterRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Adjuster success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAdjusterRequestResult = dao.GetAdjuster( uint64(createAdjusterObj.ID) )
	
	if getAdjusterRequestResult.Success == true {
		t.Errorf(getAdjusterRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestClaimReserveCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ClaimReserve
	//----------------------------------------------------------------------------
	ClaimReserveObj := model.ClaimReserve                                                                                                    {Amount:new Money(),SetDate:time.Now(),ReserveType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createClaimReserveRequestResult := dao.CreateClaimReserve( ClaimReserveObj )
	
	if createClaimReserveRequestResult.Success == false {
		t.Errorf(createClaimReserveRequestResult.Msg)
	} else {
		fmt.Println("Check Create ClaimReserve success...")
	}
	
	createClaimReserveObj,_ := createClaimReserveRequestResult.Data. (model.ClaimReserve)

	// --------------------------------------------------------------
	// Check ClaimReserve Obj ID
	// --------------------------------------------------------------	
	if createClaimReserveObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ClaimReserve" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getClaimReserveRequestResult := dao.GetClaimReserve( uint64(createClaimReserveObj.ID) )
	
	if getClaimReserveRequestResult.Success == false {
		t.Errorf(getClaimReserveRequestResult.Msg)
	} else {
		fmt.Println("Check Get ClaimReserve success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getClaimReserveObj,_ := getClaimReserveRequestResult.Data. (model.ClaimReserve)
	compareClaimReserve := cmp.Equal(createClaimReserveObj.ID, getClaimReserveObj.ID)
	
	if  compareClaimReserve == false	{
		t.Errorf( "Created ClaimReserve object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllClaimReserveRequestResult := dao.GetAllClaimReserve()

	if getAllClaimReserveRequestResult.Success == false {
			t.Errorf(getAllClaimReserveRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ClaimReserve success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllClaimReserveObj []model.ClaimReserve = getAllClaimReserveRequestResult.Data. ([]model.ClaimReserve)
		
	equalClaimReserve := cmp.Equal(createClaimReserveObj.ID, getAllClaimReserveObj[len(getAllClaimReserveObj)-1].ID)
		
	if equalClaimReserve == false {
		t.Errorf( "Created object is not equal to the last entry in ClaimReserve[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ClaimReserve
	// --------------------------------------------------------------	
	deleteClaimReserveRequestResult := dao.DeleteClaimReserve(uint64(createClaimReserveObj.ID))

	if deleteClaimReserveRequestResult.Success == false {
			t.Errorf(deleteClaimReserveRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ClaimReserve success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getClaimReserveRequestResult = dao.GetClaimReserve( uint64(createClaimReserveObj.ID) )
	
	if getClaimReserveRequestResult.Success == true {
		t.Errorf(getClaimReserveRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestClaimPaymentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ClaimPayment
	//----------------------------------------------------------------------------
	ClaimPaymentObj := model.ClaimPayment                                                                                                                                                    {PaymentNumber:"test value for PaymentNumber",Amount:new Money(),PaymentDate:time.Now(),PayeeType:0,Method:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createClaimPaymentRequestResult := dao.CreateClaimPayment( ClaimPaymentObj )
	
	if createClaimPaymentRequestResult.Success == false {
		t.Errorf(createClaimPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check Create ClaimPayment success...")
	}
	
	createClaimPaymentObj,_ := createClaimPaymentRequestResult.Data. (model.ClaimPayment)

	// --------------------------------------------------------------
	// Check ClaimPayment Obj ID
	// --------------------------------------------------------------	
	if createClaimPaymentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ClaimPayment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getClaimPaymentRequestResult := dao.GetClaimPayment( uint64(createClaimPaymentObj.ID) )
	
	if getClaimPaymentRequestResult.Success == false {
		t.Errorf(getClaimPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check Get ClaimPayment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getClaimPaymentObj,_ := getClaimPaymentRequestResult.Data. (model.ClaimPayment)
	compareClaimPayment := cmp.Equal(createClaimPaymentObj.ID, getClaimPaymentObj.ID)
	
	if  compareClaimPayment == false	{
		t.Errorf( "Created ClaimPayment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllClaimPaymentRequestResult := dao.GetAllClaimPayment()

	if getAllClaimPaymentRequestResult.Success == false {
			t.Errorf(getAllClaimPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ClaimPayment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllClaimPaymentObj []model.ClaimPayment = getAllClaimPaymentRequestResult.Data. ([]model.ClaimPayment)
		
	equalClaimPayment := cmp.Equal(createClaimPaymentObj.ID, getAllClaimPaymentObj[len(getAllClaimPaymentObj)-1].ID)
		
	if equalClaimPayment == false {
		t.Errorf( "Created object is not equal to the last entry in ClaimPayment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ClaimPayment
	// --------------------------------------------------------------	
	deleteClaimPaymentRequestResult := dao.DeleteClaimPayment(uint64(createClaimPaymentObj.ID))

	if deleteClaimPaymentRequestResult.Success == false {
			t.Errorf(deleteClaimPaymentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ClaimPayment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getClaimPaymentRequestResult = dao.GetClaimPayment( uint64(createClaimPaymentObj.ID) )
	
	if getClaimPaymentRequestResult.Success == true {
		t.Errorf(getClaimPaymentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestServiceProviderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ServiceProvider
	//----------------------------------------------------------------------------
	ServiceProviderObj := model.ServiceProvider                                                                                            {Name:"test value for Name",TaxId:"test value for TaxId",ProviderType:0,NetworkStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createServiceProviderRequestResult := dao.CreateServiceProvider( ServiceProviderObj )
	
	if createServiceProviderRequestResult.Success == false {
		t.Errorf(createServiceProviderRequestResult.Msg)
	} else {
		fmt.Println("Check Create ServiceProvider success...")
	}
	
	createServiceProviderObj,_ := createServiceProviderRequestResult.Data. (model.ServiceProvider)

	// --------------------------------------------------------------
	// Check ServiceProvider Obj ID
	// --------------------------------------------------------------	
	if createServiceProviderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ServiceProvider" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getServiceProviderRequestResult := dao.GetServiceProvider( uint64(createServiceProviderObj.ID) )
	
	if getServiceProviderRequestResult.Success == false {
		t.Errorf(getServiceProviderRequestResult.Msg)
	} else {
		fmt.Println("Check Get ServiceProvider success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getServiceProviderObj,_ := getServiceProviderRequestResult.Data. (model.ServiceProvider)
	compareServiceProvider := cmp.Equal(createServiceProviderObj.ID, getServiceProviderObj.ID)
	
	if  compareServiceProvider == false	{
		t.Errorf( "Created ServiceProvider object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllServiceProviderRequestResult := dao.GetAllServiceProvider()

	if getAllServiceProviderRequestResult.Success == false {
			t.Errorf(getAllServiceProviderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ServiceProvider success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllServiceProviderObj []model.ServiceProvider = getAllServiceProviderRequestResult.Data. ([]model.ServiceProvider)
		
	equalServiceProvider := cmp.Equal(createServiceProviderObj.ID, getAllServiceProviderObj[len(getAllServiceProviderObj)-1].ID)
		
	if equalServiceProvider == false {
		t.Errorf( "Created object is not equal to the last entry in ServiceProvider[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ServiceProvider
	// --------------------------------------------------------------	
	deleteServiceProviderRequestResult := dao.DeleteServiceProvider(uint64(createServiceProviderObj.ID))

	if deleteServiceProviderRequestResult.Success == false {
			t.Errorf(deleteServiceProviderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ServiceProvider success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getServiceProviderRequestResult = dao.GetServiceProvider( uint64(createServiceProviderObj.ID) )
	
	if getServiceProviderRequestResult.Success == true {
		t.Errorf(getServiceProviderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestReinsuranceAgreementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ReinsuranceAgreement
	//----------------------------------------------------------------------------
	ReinsuranceAgreementObj := model.ReinsuranceAgreement                                                                                                                            {AgreementNumber:"test value for AgreementNumber",EffectivePeriod:new DateRange(),Retention:new Money(),Limit:new Money(),CessionPercentage:new Percentage(),ReinsuranceType:0,TreatyType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createReinsuranceAgreementRequestResult := dao.CreateReinsuranceAgreement( ReinsuranceAgreementObj )
	
	if createReinsuranceAgreementRequestResult.Success == false {
		t.Errorf(createReinsuranceAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check Create ReinsuranceAgreement success...")
	}
	
	createReinsuranceAgreementObj,_ := createReinsuranceAgreementRequestResult.Data. (model.ReinsuranceAgreement)

	// --------------------------------------------------------------
	// Check ReinsuranceAgreement Obj ID
	// --------------------------------------------------------------	
	if createReinsuranceAgreementObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ReinsuranceAgreement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getReinsuranceAgreementRequestResult := dao.GetReinsuranceAgreement( uint64(createReinsuranceAgreementObj.ID) )
	
	if getReinsuranceAgreementRequestResult.Success == false {
		t.Errorf(getReinsuranceAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check Get ReinsuranceAgreement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getReinsuranceAgreementObj,_ := getReinsuranceAgreementRequestResult.Data. (model.ReinsuranceAgreement)
	compareReinsuranceAgreement := cmp.Equal(createReinsuranceAgreementObj.ID, getReinsuranceAgreementObj.ID)
	
	if  compareReinsuranceAgreement == false	{
		t.Errorf( "Created ReinsuranceAgreement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllReinsuranceAgreementRequestResult := dao.GetAllReinsuranceAgreement()

	if getAllReinsuranceAgreementRequestResult.Success == false {
			t.Errorf(getAllReinsuranceAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ReinsuranceAgreement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllReinsuranceAgreementObj []model.ReinsuranceAgreement = getAllReinsuranceAgreementRequestResult.Data. ([]model.ReinsuranceAgreement)
		
	equalReinsuranceAgreement := cmp.Equal(createReinsuranceAgreementObj.ID, getAllReinsuranceAgreementObj[len(getAllReinsuranceAgreementObj)-1].ID)
		
	if equalReinsuranceAgreement == false {
		t.Errorf( "Created object is not equal to the last entry in ReinsuranceAgreement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ReinsuranceAgreement
	// --------------------------------------------------------------	
	deleteReinsuranceAgreementRequestResult := dao.DeleteReinsuranceAgreement(uint64(createReinsuranceAgreementObj.ID))

	if deleteReinsuranceAgreementRequestResult.Success == false {
			t.Errorf(deleteReinsuranceAgreementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ReinsuranceAgreement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getReinsuranceAgreementRequestResult = dao.GetReinsuranceAgreement( uint64(createReinsuranceAgreementObj.ID) )
	
	if getReinsuranceAgreementRequestResult.Success == true {
		t.Errorf(getReinsuranceAgreementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSubrogationRecoveryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SubrogationRecovery
	//----------------------------------------------------------------------------
	SubrogationRecoveryObj := model.SubrogationRecovery                                                                                                                    {RecoveryReference:"test value for RecoveryReference",Amount:new Money(),RecoveryDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSubrogationRecoveryRequestResult := dao.CreateSubrogationRecovery( SubrogationRecoveryObj )
	
	if createSubrogationRecoveryRequestResult.Success == false {
		t.Errorf(createSubrogationRecoveryRequestResult.Msg)
	} else {
		fmt.Println("Check Create SubrogationRecovery success...")
	}
	
	createSubrogationRecoveryObj,_ := createSubrogationRecoveryRequestResult.Data. (model.SubrogationRecovery)

	// --------------------------------------------------------------
	// Check SubrogationRecovery Obj ID
	// --------------------------------------------------------------	
	if createSubrogationRecoveryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SubrogationRecovery" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSubrogationRecoveryRequestResult := dao.GetSubrogationRecovery( uint64(createSubrogationRecoveryObj.ID) )
	
	if getSubrogationRecoveryRequestResult.Success == false {
		t.Errorf(getSubrogationRecoveryRequestResult.Msg)
	} else {
		fmt.Println("Check Get SubrogationRecovery success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSubrogationRecoveryObj,_ := getSubrogationRecoveryRequestResult.Data. (model.SubrogationRecovery)
	compareSubrogationRecovery := cmp.Equal(createSubrogationRecoveryObj.ID, getSubrogationRecoveryObj.ID)
	
	if  compareSubrogationRecovery == false	{
		t.Errorf( "Created SubrogationRecovery object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSubrogationRecoveryRequestResult := dao.GetAllSubrogationRecovery()

	if getAllSubrogationRecoveryRequestResult.Success == false {
			t.Errorf(getAllSubrogationRecoveryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SubrogationRecovery success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSubrogationRecoveryObj []model.SubrogationRecovery = getAllSubrogationRecoveryRequestResult.Data. ([]model.SubrogationRecovery)
		
	equalSubrogationRecovery := cmp.Equal(createSubrogationRecoveryObj.ID, getAllSubrogationRecoveryObj[len(getAllSubrogationRecoveryObj)-1].ID)
		
	if equalSubrogationRecovery == false {
		t.Errorf( "Created object is not equal to the last entry in SubrogationRecovery[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SubrogationRecovery
	// --------------------------------------------------------------	
	deleteSubrogationRecoveryRequestResult := dao.DeleteSubrogationRecovery(uint64(createSubrogationRecoveryObj.ID))

	if deleteSubrogationRecoveryRequestResult.Success == false {
			t.Errorf(deleteSubrogationRecoveryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SubrogationRecovery success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSubrogationRecoveryRequestResult = dao.GetSubrogationRecovery( uint64(createSubrogationRecoveryObj.ID) )
	
	if getSubrogationRecoveryRequestResult.Success == true {
		t.Errorf(getSubrogationRecoveryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestThirdPartyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ThirdParty
	//----------------------------------------------------------------------------
	ThirdPartyObj := model.ThirdParty                                                                                            {Name:"test value for Name",TaxId:"test value for TaxId",Address:new Address(),PartyType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createThirdPartyRequestResult := dao.CreateThirdParty( ThirdPartyObj )
	
	if createThirdPartyRequestResult.Success == false {
		t.Errorf(createThirdPartyRequestResult.Msg)
	} else {
		fmt.Println("Check Create ThirdParty success...")
	}
	
	createThirdPartyObj,_ := createThirdPartyRequestResult.Data. (model.ThirdParty)

	// --------------------------------------------------------------
	// Check ThirdParty Obj ID
	// --------------------------------------------------------------	
	if createThirdPartyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ThirdParty" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getThirdPartyRequestResult := dao.GetThirdParty( uint64(createThirdPartyObj.ID) )
	
	if getThirdPartyRequestResult.Success == false {
		t.Errorf(getThirdPartyRequestResult.Msg)
	} else {
		fmt.Println("Check Get ThirdParty success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getThirdPartyObj,_ := getThirdPartyRequestResult.Data. (model.ThirdParty)
	compareThirdParty := cmp.Equal(createThirdPartyObj.ID, getThirdPartyObj.ID)
	
	if  compareThirdParty == false	{
		t.Errorf( "Created ThirdParty object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllThirdPartyRequestResult := dao.GetAllThirdParty()

	if getAllThirdPartyRequestResult.Success == false {
			t.Errorf(getAllThirdPartyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ThirdParty success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllThirdPartyObj []model.ThirdParty = getAllThirdPartyRequestResult.Data. ([]model.ThirdParty)
		
	equalThirdParty := cmp.Equal(createThirdPartyObj.ID, getAllThirdPartyObj[len(getAllThirdPartyObj)-1].ID)
		
	if equalThirdParty == false {
		t.Errorf( "Created object is not equal to the last entry in ThirdParty[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ThirdParty
	// --------------------------------------------------------------	
	deleteThirdPartyRequestResult := dao.DeleteThirdParty(uint64(createThirdPartyObj.ID))

	if deleteThirdPartyRequestResult.Success == false {
			t.Errorf(deleteThirdPartyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ThirdParty success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getThirdPartyRequestResult = dao.GetThirdParty( uint64(createThirdPartyObj.ID) )
	
	if getThirdPartyRequestResult.Success == true {
		t.Errorf(getThirdPartyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDocumentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Document
	//----------------------------------------------------------------------------
	DocumentObj := model.Document                                                                                                    {FileName:"test value for FileName",UploadedDate:time.Now(),DocumentType:0}

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

