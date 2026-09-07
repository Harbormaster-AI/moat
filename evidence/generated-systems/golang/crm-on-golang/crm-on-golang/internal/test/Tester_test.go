package test

import ( 
	"testing"
    dao "crm-on-golang/internal/dao"
	"crm-on-golang/internal/model"
	"crm-on-golang/internal/utils"
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
	OrganizationObj := model.Organization                                                                                            {Name:"test value for Name",DefaultCurrency:"test value for DefaultCurrency",DefaultLocale:new _Locale(),Website:new URL()}

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


func TestUserCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for User
	//----------------------------------------------------------------------------
	UserObj := model.User                                                                                                                            {Username:"test value for Username",FullName:"test value for FullName",Email:new EmailAddress(),Locale:new _Locale(),Role:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createUserRequestResult := dao.CreateUser( UserObj )
	
	if createUserRequestResult.Success == false {
		t.Errorf(createUserRequestResult.Msg)
	} else {
		fmt.Println("Check Create User success...")
	}
	
	createUserObj,_ := createUserRequestResult.Data. (model.User)

	// --------------------------------------------------------------
	// Check User Obj ID
	// --------------------------------------------------------------	
	if createUserObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for User" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getUserRequestResult := dao.GetUser( uint64(createUserObj.ID) )
	
	if getUserRequestResult.Success == false {
		t.Errorf(getUserRequestResult.Msg)
	} else {
		fmt.Println("Check Get User success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getUserObj,_ := getUserRequestResult.Data. (model.User)
	compareUser := cmp.Equal(createUserObj.ID, getUserObj.ID)
	
	if  compareUser == false	{
		t.Errorf( "Created User object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllUserRequestResult := dao.GetAllUser()

	if getAllUserRequestResult.Success == false {
			t.Errorf(getAllUserRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll User success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllUserObj []model.User = getAllUserRequestResult.Data. ([]model.User)
		
	equalUser := cmp.Equal(createUserObj.ID, getAllUserObj[len(getAllUserObj)-1].ID)
		
	if equalUser == false {
		t.Errorf( "Created object is not equal to the last entry in User[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for User
	// --------------------------------------------------------------	
	deleteUserRequestResult := dao.DeleteUser(uint64(createUserObj.ID))

	if deleteUserRequestResult.Success == false {
			t.Errorf(deleteUserRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion User success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getUserRequestResult = dao.GetUser( uint64(createUserObj.ID) )
	
	if getUserRequestResult.Success == true {
		t.Errorf(getUserRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTeamCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Team
	//----------------------------------------------------------------------------
	TeamObj := model.Team                                            {Name:"test value for Name",TeamType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTeamRequestResult := dao.CreateTeam( TeamObj )
	
	if createTeamRequestResult.Success == false {
		t.Errorf(createTeamRequestResult.Msg)
	} else {
		fmt.Println("Check Create Team success...")
	}
	
	createTeamObj,_ := createTeamRequestResult.Data. (model.Team)

	// --------------------------------------------------------------
	// Check Team Obj ID
	// --------------------------------------------------------------	
	if createTeamObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Team" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTeamRequestResult := dao.GetTeam( uint64(createTeamObj.ID) )
	
	if getTeamRequestResult.Success == false {
		t.Errorf(getTeamRequestResult.Msg)
	} else {
		fmt.Println("Check Get Team success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTeamObj,_ := getTeamRequestResult.Data. (model.Team)
	compareTeam := cmp.Equal(createTeamObj.ID, getTeamObj.ID)
	
	if  compareTeam == false	{
		t.Errorf( "Created Team object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTeamRequestResult := dao.GetAllTeam()

	if getAllTeamRequestResult.Success == false {
			t.Errorf(getAllTeamRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Team success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTeamObj []model.Team = getAllTeamRequestResult.Data. ([]model.Team)
		
	equalTeam := cmp.Equal(createTeamObj.ID, getAllTeamObj[len(getAllTeamObj)-1].ID)
		
	if equalTeam == false {
		t.Errorf( "Created object is not equal to the last entry in Team[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Team
	// --------------------------------------------------------------	
	deleteTeamRequestResult := dao.DeleteTeam(uint64(createTeamObj.ID))

	if deleteTeamRequestResult.Success == false {
			t.Errorf(deleteTeamRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Team success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTeamRequestResult = dao.GetTeam( uint64(createTeamObj.ID) )
	
	if getTeamRequestResult.Success == true {
		t.Errorf(getTeamRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTerritoryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Territory
	//----------------------------------------------------------------------------
	TerritoryObj := model.Territory                                                                            {Name:"test value for Name",Region:"test value for Region",TerritoryType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTerritoryRequestResult := dao.CreateTerritory( TerritoryObj )
	
	if createTerritoryRequestResult.Success == false {
		t.Errorf(createTerritoryRequestResult.Msg)
	} else {
		fmt.Println("Check Create Territory success...")
	}
	
	createTerritoryObj,_ := createTerritoryRequestResult.Data. (model.Territory)

	// --------------------------------------------------------------
	// Check Territory Obj ID
	// --------------------------------------------------------------	
	if createTerritoryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Territory" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTerritoryRequestResult := dao.GetTerritory( uint64(createTerritoryObj.ID) )
	
	if getTerritoryRequestResult.Success == false {
		t.Errorf(getTerritoryRequestResult.Msg)
	} else {
		fmt.Println("Check Get Territory success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTerritoryObj,_ := getTerritoryRequestResult.Data. (model.Territory)
	compareTerritory := cmp.Equal(createTerritoryObj.ID, getTerritoryObj.ID)
	
	if  compareTerritory == false	{
		t.Errorf( "Created Territory object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTerritoryRequestResult := dao.GetAllTerritory()

	if getAllTerritoryRequestResult.Success == false {
			t.Errorf(getAllTerritoryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Territory success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTerritoryObj []model.Territory = getAllTerritoryRequestResult.Data. ([]model.Territory)
		
	equalTerritory := cmp.Equal(createTerritoryObj.ID, getAllTerritoryObj[len(getAllTerritoryObj)-1].ID)
		
	if equalTerritory == false {
		t.Errorf( "Created object is not equal to the last entry in Territory[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Territory
	// --------------------------------------------------------------	
	deleteTerritoryRequestResult := dao.DeleteTerritory(uint64(createTerritoryObj.ID))

	if deleteTerritoryRequestResult.Success == false {
			t.Errorf(deleteTerritoryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Territory success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTerritoryRequestResult = dao.GetTerritory( uint64(createTerritoryObj.ID) )
	
	if getTerritoryRequestResult.Success == true {
		t.Errorf(getTerritoryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAccountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Account
	//----------------------------------------------------------------------------
	AccountObj := model.Account                                                                                                                                                                                                                            {Name:"test value for Name",AccountNumber:"test value for AccountNumber",Industry:"test value for Industry",BillingAddress:new Address(),ShippingAddress:new Address(),Website:new URL(),Phone:new PhoneNumber(),AsActive:true,AccountType:0,LifecycleStage:0}

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


func TestContactCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Contact
	//----------------------------------------------------------------------------
	ContactObj := model.Contact                                                                                                                                                                            {FirstName:"test value for FirstName",LastName:"test value for LastName",Title:"test value for Title",Email:new EmailAddress(),Phone:new PhoneNumber(),Mobile:new PhoneNumber(),MailingAddress:new Address(),PreferredContactMethod:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createContactRequestResult := dao.CreateContact( ContactObj )
	
	if createContactRequestResult.Success == false {
		t.Errorf(createContactRequestResult.Msg)
	} else {
		fmt.Println("Check Create Contact success...")
	}
	
	createContactObj,_ := createContactRequestResult.Data. (model.Contact)

	// --------------------------------------------------------------
	// Check Contact Obj ID
	// --------------------------------------------------------------	
	if createContactObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Contact" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getContactRequestResult := dao.GetContact( uint64(createContactObj.ID) )
	
	if getContactRequestResult.Success == false {
		t.Errorf(getContactRequestResult.Msg)
	} else {
		fmt.Println("Check Get Contact success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getContactObj,_ := getContactRequestResult.Data. (model.Contact)
	compareContact := cmp.Equal(createContactObj.ID, getContactObj.ID)
	
	if  compareContact == false	{
		t.Errorf( "Created Contact object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllContactRequestResult := dao.GetAllContact()

	if getAllContactRequestResult.Success == false {
			t.Errorf(getAllContactRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Contact success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllContactObj []model.Contact = getAllContactRequestResult.Data. ([]model.Contact)
		
	equalContact := cmp.Equal(createContactObj.ID, getAllContactObj[len(getAllContactObj)-1].ID)
		
	if equalContact == false {
		t.Errorf( "Created object is not equal to the last entry in Contact[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Contact
	// --------------------------------------------------------------	
	deleteContactRequestResult := dao.DeleteContact(uint64(createContactObj.ID))

	if deleteContactRequestResult.Success == false {
			t.Errorf(deleteContactRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Contact success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getContactRequestResult = dao.GetContact( uint64(createContactObj.ID) )
	
	if getContactRequestResult.Success == true {
		t.Errorf(getContactRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLeadCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Lead
	//----------------------------------------------------------------------------
	LeadObj := model.Lead                                                                                                                                                                                                            {FirstName:"test value for FirstName",LastName:"test value for LastName",Company:"test value for Company",Email:new EmailAddress(),Phone:new PhoneNumber(),Converted:true,Status:0,Source:0,Rating:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLeadRequestResult := dao.CreateLead( LeadObj )
	
	if createLeadRequestResult.Success == false {
		t.Errorf(createLeadRequestResult.Msg)
	} else {
		fmt.Println("Check Create Lead success...")
	}
	
	createLeadObj,_ := createLeadRequestResult.Data. (model.Lead)

	// --------------------------------------------------------------
	// Check Lead Obj ID
	// --------------------------------------------------------------	
	if createLeadObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Lead" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLeadRequestResult := dao.GetLead( uint64(createLeadObj.ID) )
	
	if getLeadRequestResult.Success == false {
		t.Errorf(getLeadRequestResult.Msg)
	} else {
		fmt.Println("Check Get Lead success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLeadObj,_ := getLeadRequestResult.Data. (model.Lead)
	compareLead := cmp.Equal(createLeadObj.ID, getLeadObj.ID)
	
	if  compareLead == false	{
		t.Errorf( "Created Lead object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLeadRequestResult := dao.GetAllLead()

	if getAllLeadRequestResult.Success == false {
			t.Errorf(getAllLeadRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Lead success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLeadObj []model.Lead = getAllLeadRequestResult.Data. ([]model.Lead)
		
	equalLead := cmp.Equal(createLeadObj.ID, getAllLeadObj[len(getAllLeadObj)-1].ID)
		
	if equalLead == false {
		t.Errorf( "Created object is not equal to the last entry in Lead[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Lead
	// --------------------------------------------------------------	
	deleteLeadRequestResult := dao.DeleteLead(uint64(createLeadObj.ID))

	if deleteLeadRequestResult.Success == false {
			t.Errorf(deleteLeadRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Lead success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLeadRequestResult = dao.GetLead( uint64(createLeadObj.ID) )
	
	if getLeadRequestResult.Success == true {
		t.Errorf(getLeadRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestOpportunityCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Opportunity
	//----------------------------------------------------------------------------
	OpportunityObj := model.Opportunity                                                                                                                                                                                                                                            {Name:"test value for Name",Amount:new Money(),CloseDate:time.Now(),Probability:"test value",Description:"test value for Description",Stage:0,Type:0,ForecastCategory:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOpportunityRequestResult := dao.CreateOpportunity( OpportunityObj )
	
	if createOpportunityRequestResult.Success == false {
		t.Errorf(createOpportunityRequestResult.Msg)
	} else {
		fmt.Println("Check Create Opportunity success...")
	}
	
	createOpportunityObj,_ := createOpportunityRequestResult.Data. (model.Opportunity)

	// --------------------------------------------------------------
	// Check Opportunity Obj ID
	// --------------------------------------------------------------	
	if createOpportunityObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Opportunity" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOpportunityRequestResult := dao.GetOpportunity( uint64(createOpportunityObj.ID) )
	
	if getOpportunityRequestResult.Success == false {
		t.Errorf(getOpportunityRequestResult.Msg)
	} else {
		fmt.Println("Check Get Opportunity success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOpportunityObj,_ := getOpportunityRequestResult.Data. (model.Opportunity)
	compareOpportunity := cmp.Equal(createOpportunityObj.ID, getOpportunityObj.ID)
	
	if  compareOpportunity == false	{
		t.Errorf( "Created Opportunity object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOpportunityRequestResult := dao.GetAllOpportunity()

	if getAllOpportunityRequestResult.Success == false {
			t.Errorf(getAllOpportunityRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Opportunity success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOpportunityObj []model.Opportunity = getAllOpportunityRequestResult.Data. ([]model.Opportunity)
		
	equalOpportunity := cmp.Equal(createOpportunityObj.ID, getAllOpportunityObj[len(getAllOpportunityObj)-1].ID)
		
	if equalOpportunity == false {
		t.Errorf( "Created object is not equal to the last entry in Opportunity[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Opportunity
	// --------------------------------------------------------------	
	deleteOpportunityRequestResult := dao.DeleteOpportunity(uint64(createOpportunityObj.ID))

	if deleteOpportunityRequestResult.Success == false {
			t.Errorf(deleteOpportunityRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Opportunity success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOpportunityRequestResult = dao.GetOpportunity( uint64(createOpportunityObj.ID) )
	
	if getOpportunityRequestResult.Success == true {
		t.Errorf(getOpportunityRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestOpportunityLineItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for OpportunityLineItem
	//----------------------------------------------------------------------------
	OpportunityLineItemObj := model.OpportunityLineItem                                                                                                                                            {Quantity:"test value",UnitPrice:new Money(),DiscountPercent:"test value",TotalPrice:new Money()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOpportunityLineItemRequestResult := dao.CreateOpportunityLineItem( OpportunityLineItemObj )
	
	if createOpportunityLineItemRequestResult.Success == false {
		t.Errorf(createOpportunityLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check Create OpportunityLineItem success...")
	}
	
	createOpportunityLineItemObj,_ := createOpportunityLineItemRequestResult.Data. (model.OpportunityLineItem)

	// --------------------------------------------------------------
	// Check OpportunityLineItem Obj ID
	// --------------------------------------------------------------	
	if createOpportunityLineItemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for OpportunityLineItem" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOpportunityLineItemRequestResult := dao.GetOpportunityLineItem( uint64(createOpportunityLineItemObj.ID) )
	
	if getOpportunityLineItemRequestResult.Success == false {
		t.Errorf(getOpportunityLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check Get OpportunityLineItem success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOpportunityLineItemObj,_ := getOpportunityLineItemRequestResult.Data. (model.OpportunityLineItem)
	compareOpportunityLineItem := cmp.Equal(createOpportunityLineItemObj.ID, getOpportunityLineItemObj.ID)
	
	if  compareOpportunityLineItem == false	{
		t.Errorf( "Created OpportunityLineItem object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOpportunityLineItemRequestResult := dao.GetAllOpportunityLineItem()

	if getAllOpportunityLineItemRequestResult.Success == false {
			t.Errorf(getAllOpportunityLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll OpportunityLineItem success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOpportunityLineItemObj []model.OpportunityLineItem = getAllOpportunityLineItemRequestResult.Data. ([]model.OpportunityLineItem)
		
	equalOpportunityLineItem := cmp.Equal(createOpportunityLineItemObj.ID, getAllOpportunityLineItemObj[len(getAllOpportunityLineItemObj)-1].ID)
		
	if equalOpportunityLineItem == false {
		t.Errorf( "Created object is not equal to the last entry in OpportunityLineItem[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for OpportunityLineItem
	// --------------------------------------------------------------	
	deleteOpportunityLineItemRequestResult := dao.DeleteOpportunityLineItem(uint64(createOpportunityLineItemObj.ID))

	if deleteOpportunityLineItemRequestResult.Success == false {
			t.Errorf(deleteOpportunityLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion OpportunityLineItem success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOpportunityLineItemRequestResult = dao.GetOpportunityLineItem( uint64(createOpportunityLineItemObj.ID) )
	
	if getOpportunityLineItemRequestResult.Success == true {
		t.Errorf(getOpportunityLineItemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestOpportunityStageHistoryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for OpportunityStageHistory
	//----------------------------------------------------------------------------
	OpportunityStageHistoryObj := model.OpportunityStageHistory                                                                                                                    {ChangedAt:time.Now(),Comment:"test value for Comment",FromStage:0,ToStage:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOpportunityStageHistoryRequestResult := dao.CreateOpportunityStageHistory( OpportunityStageHistoryObj )
	
	if createOpportunityStageHistoryRequestResult.Success == false {
		t.Errorf(createOpportunityStageHistoryRequestResult.Msg)
	} else {
		fmt.Println("Check Create OpportunityStageHistory success...")
	}
	
	createOpportunityStageHistoryObj,_ := createOpportunityStageHistoryRequestResult.Data. (model.OpportunityStageHistory)

	// --------------------------------------------------------------
	// Check OpportunityStageHistory Obj ID
	// --------------------------------------------------------------	
	if createOpportunityStageHistoryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for OpportunityStageHistory" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOpportunityStageHistoryRequestResult := dao.GetOpportunityStageHistory( uint64(createOpportunityStageHistoryObj.ID) )
	
	if getOpportunityStageHistoryRequestResult.Success == false {
		t.Errorf(getOpportunityStageHistoryRequestResult.Msg)
	} else {
		fmt.Println("Check Get OpportunityStageHistory success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOpportunityStageHistoryObj,_ := getOpportunityStageHistoryRequestResult.Data. (model.OpportunityStageHistory)
	compareOpportunityStageHistory := cmp.Equal(createOpportunityStageHistoryObj.ID, getOpportunityStageHistoryObj.ID)
	
	if  compareOpportunityStageHistory == false	{
		t.Errorf( "Created OpportunityStageHistory object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOpportunityStageHistoryRequestResult := dao.GetAllOpportunityStageHistory()

	if getAllOpportunityStageHistoryRequestResult.Success == false {
			t.Errorf(getAllOpportunityStageHistoryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll OpportunityStageHistory success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOpportunityStageHistoryObj []model.OpportunityStageHistory = getAllOpportunityStageHistoryRequestResult.Data. ([]model.OpportunityStageHistory)
		
	equalOpportunityStageHistory := cmp.Equal(createOpportunityStageHistoryObj.ID, getAllOpportunityStageHistoryObj[len(getAllOpportunityStageHistoryObj)-1].ID)
		
	if equalOpportunityStageHistory == false {
		t.Errorf( "Created object is not equal to the last entry in OpportunityStageHistory[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for OpportunityStageHistory
	// --------------------------------------------------------------	
	deleteOpportunityStageHistoryRequestResult := dao.DeleteOpportunityStageHistory(uint64(createOpportunityStageHistoryObj.ID))

	if deleteOpportunityStageHistoryRequestResult.Success == false {
			t.Errorf(deleteOpportunityStageHistoryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion OpportunityStageHistory success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOpportunityStageHistoryRequestResult = dao.GetOpportunityStageHistory( uint64(createOpportunityStageHistoryObj.ID) )
	
	if getOpportunityStageHistoryRequestResult.Success == true {
		t.Errorf(getOpportunityStageHistoryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestProductCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Product
	//----------------------------------------------------------------------------
	ProductObj := model.Product                                                                                                                                                                            {Sku:"test value for Sku",Name:"test value for Name",AsActive:true,StandardPrice:new Money(),Description:"test value for Description",ProductType:0,Uom:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createProductRequestResult := dao.CreateProduct( ProductObj )
	
	if createProductRequestResult.Success == false {
		t.Errorf(createProductRequestResult.Msg)
	} else {
		fmt.Println("Check Create Product success...")
	}
	
	createProductObj,_ := createProductRequestResult.Data. (model.Product)

	// --------------------------------------------------------------
	// Check Product Obj ID
	// --------------------------------------------------------------	
	if createProductObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Product" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getProductRequestResult := dao.GetProduct( uint64(createProductObj.ID) )
	
	if getProductRequestResult.Success == false {
		t.Errorf(getProductRequestResult.Msg)
	} else {
		fmt.Println("Check Get Product success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getProductObj,_ := getProductRequestResult.Data. (model.Product)
	compareProduct := cmp.Equal(createProductObj.ID, getProductObj.ID)
	
	if  compareProduct == false	{
		t.Errorf( "Created Product object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllProductRequestResult := dao.GetAllProduct()

	if getAllProductRequestResult.Success == false {
			t.Errorf(getAllProductRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Product success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllProductObj []model.Product = getAllProductRequestResult.Data. ([]model.Product)
		
	equalProduct := cmp.Equal(createProductObj.ID, getAllProductObj[len(getAllProductObj)-1].ID)
		
	if equalProduct == false {
		t.Errorf( "Created object is not equal to the last entry in Product[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Product
	// --------------------------------------------------------------	
	deleteProductRequestResult := dao.DeleteProduct(uint64(createProductObj.ID))

	if deleteProductRequestResult.Success == false {
			t.Errorf(deleteProductRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Product success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getProductRequestResult = dao.GetProduct( uint64(createProductObj.ID) )
	
	if getProductRequestResult.Success == true {
		t.Errorf(getProductRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPriceBookCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PriceBook
	//----------------------------------------------------------------------------
	PriceBookObj := model.PriceBook                                                                                            {Name:"test value for Name",AsActive:true,Description:"test value for Description"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPriceBookRequestResult := dao.CreatePriceBook( PriceBookObj )
	
	if createPriceBookRequestResult.Success == false {
		t.Errorf(createPriceBookRequestResult.Msg)
	} else {
		fmt.Println("Check Create PriceBook success...")
	}
	
	createPriceBookObj,_ := createPriceBookRequestResult.Data. (model.PriceBook)

	// --------------------------------------------------------------
	// Check PriceBook Obj ID
	// --------------------------------------------------------------	
	if createPriceBookObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PriceBook" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPriceBookRequestResult := dao.GetPriceBook( uint64(createPriceBookObj.ID) )
	
	if getPriceBookRequestResult.Success == false {
		t.Errorf(getPriceBookRequestResult.Msg)
	} else {
		fmt.Println("Check Get PriceBook success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPriceBookObj,_ := getPriceBookRequestResult.Data. (model.PriceBook)
	comparePriceBook := cmp.Equal(createPriceBookObj.ID, getPriceBookObj.ID)
	
	if  comparePriceBook == false	{
		t.Errorf( "Created PriceBook object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPriceBookRequestResult := dao.GetAllPriceBook()

	if getAllPriceBookRequestResult.Success == false {
			t.Errorf(getAllPriceBookRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PriceBook success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPriceBookObj []model.PriceBook = getAllPriceBookRequestResult.Data. ([]model.PriceBook)
		
	equalPriceBook := cmp.Equal(createPriceBookObj.ID, getAllPriceBookObj[len(getAllPriceBookObj)-1].ID)
		
	if equalPriceBook == false {
		t.Errorf( "Created object is not equal to the last entry in PriceBook[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PriceBook
	// --------------------------------------------------------------	
	deletePriceBookRequestResult := dao.DeletePriceBook(uint64(createPriceBookObj.ID))

	if deletePriceBookRequestResult.Success == false {
			t.Errorf(deletePriceBookRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PriceBook success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPriceBookRequestResult = dao.GetPriceBook( uint64(createPriceBookObj.ID) )
	
	if getPriceBookRequestResult.Success == true {
		t.Errorf(getPriceBookRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPriceBookEntryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PriceBookEntry
	//----------------------------------------------------------------------------
	PriceBookEntryObj := model.PriceBookEntry                                                                                                                                                            {UnitPrice:new Money(),EffectiveDate:time.Now(),ExpirationDate:time.Now(),AsActive:true}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPriceBookEntryRequestResult := dao.CreatePriceBookEntry( PriceBookEntryObj )
	
	if createPriceBookEntryRequestResult.Success == false {
		t.Errorf(createPriceBookEntryRequestResult.Msg)
	} else {
		fmt.Println("Check Create PriceBookEntry success...")
	}
	
	createPriceBookEntryObj,_ := createPriceBookEntryRequestResult.Data. (model.PriceBookEntry)

	// --------------------------------------------------------------
	// Check PriceBookEntry Obj ID
	// --------------------------------------------------------------	
	if createPriceBookEntryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PriceBookEntry" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPriceBookEntryRequestResult := dao.GetPriceBookEntry( uint64(createPriceBookEntryObj.ID) )
	
	if getPriceBookEntryRequestResult.Success == false {
		t.Errorf(getPriceBookEntryRequestResult.Msg)
	} else {
		fmt.Println("Check Get PriceBookEntry success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPriceBookEntryObj,_ := getPriceBookEntryRequestResult.Data. (model.PriceBookEntry)
	comparePriceBookEntry := cmp.Equal(createPriceBookEntryObj.ID, getPriceBookEntryObj.ID)
	
	if  comparePriceBookEntry == false	{
		t.Errorf( "Created PriceBookEntry object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPriceBookEntryRequestResult := dao.GetAllPriceBookEntry()

	if getAllPriceBookEntryRequestResult.Success == false {
			t.Errorf(getAllPriceBookEntryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PriceBookEntry success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPriceBookEntryObj []model.PriceBookEntry = getAllPriceBookEntryRequestResult.Data. ([]model.PriceBookEntry)
		
	equalPriceBookEntry := cmp.Equal(createPriceBookEntryObj.ID, getAllPriceBookEntryObj[len(getAllPriceBookEntryObj)-1].ID)
		
	if equalPriceBookEntry == false {
		t.Errorf( "Created object is not equal to the last entry in PriceBookEntry[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PriceBookEntry
	// --------------------------------------------------------------	
	deletePriceBookEntryRequestResult := dao.DeletePriceBookEntry(uint64(createPriceBookEntryObj.ID))

	if deletePriceBookEntryRequestResult.Success == false {
			t.Errorf(deletePriceBookEntryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PriceBookEntry success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPriceBookEntryRequestResult = dao.GetPriceBookEntry( uint64(createPriceBookEntryObj.ID) )
	
	if getPriceBookEntryRequestResult.Success == true {
		t.Errorf(getPriceBookEntryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestQuoteCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Quote
	//----------------------------------------------------------------------------
	QuoteObj := model.Quote                                                                                                                                                                                                                                                                    {QuoteNumber:"test value for QuoteNumber",ValidityStart:time.Now(),ValidityEnd:time.Now(),TotalAmount:new Money(),DiscountPercent:"test value",TaxAmount:new Money(),ShippingAmount:new Money(),Status:0}

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


func TestQuoteLineItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for QuoteLineItem
	//----------------------------------------------------------------------------
	QuoteLineItemObj := model.QuoteLineItem                                                                                                                    {Quantity:"test value",UnitPrice:new Money(),DiscountAmount:new Money(),TaxAmount:new Money(),TotalAmount:new Money()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createQuoteLineItemRequestResult := dao.CreateQuoteLineItem( QuoteLineItemObj )
	
	if createQuoteLineItemRequestResult.Success == false {
		t.Errorf(createQuoteLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check Create QuoteLineItem success...")
	}
	
	createQuoteLineItemObj,_ := createQuoteLineItemRequestResult.Data. (model.QuoteLineItem)

	// --------------------------------------------------------------
	// Check QuoteLineItem Obj ID
	// --------------------------------------------------------------	
	if createQuoteLineItemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for QuoteLineItem" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getQuoteLineItemRequestResult := dao.GetQuoteLineItem( uint64(createQuoteLineItemObj.ID) )
	
	if getQuoteLineItemRequestResult.Success == false {
		t.Errorf(getQuoteLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check Get QuoteLineItem success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getQuoteLineItemObj,_ := getQuoteLineItemRequestResult.Data. (model.QuoteLineItem)
	compareQuoteLineItem := cmp.Equal(createQuoteLineItemObj.ID, getQuoteLineItemObj.ID)
	
	if  compareQuoteLineItem == false	{
		t.Errorf( "Created QuoteLineItem object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllQuoteLineItemRequestResult := dao.GetAllQuoteLineItem()

	if getAllQuoteLineItemRequestResult.Success == false {
			t.Errorf(getAllQuoteLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll QuoteLineItem success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllQuoteLineItemObj []model.QuoteLineItem = getAllQuoteLineItemRequestResult.Data. ([]model.QuoteLineItem)
		
	equalQuoteLineItem := cmp.Equal(createQuoteLineItemObj.ID, getAllQuoteLineItemObj[len(getAllQuoteLineItemObj)-1].ID)
		
	if equalQuoteLineItem == false {
		t.Errorf( "Created object is not equal to the last entry in QuoteLineItem[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for QuoteLineItem
	// --------------------------------------------------------------	
	deleteQuoteLineItemRequestResult := dao.DeleteQuoteLineItem(uint64(createQuoteLineItemObj.ID))

	if deleteQuoteLineItemRequestResult.Success == false {
			t.Errorf(deleteQuoteLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion QuoteLineItem success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getQuoteLineItemRequestResult = dao.GetQuoteLineItem( uint64(createQuoteLineItemObj.ID) )
	
	if getQuoteLineItemRequestResult.Success == true {
		t.Errorf(getQuoteLineItemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Order
	//----------------------------------------------------------------------------
	OrderObj := model.Order                                                                                                                                                    {OrderNumber:"test value for OrderNumber",OrderDate:time.Now(),TotalAmount:new Money(),TaxAmount:new Money(),ShippingAmount:new Money(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOrderRequestResult := dao.CreateOrder( OrderObj )
	
	if createOrderRequestResult.Success == false {
		t.Errorf(createOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create Order success...")
	}
	
	createOrderObj,_ := createOrderRequestResult.Data. (model.Order)

	// --------------------------------------------------------------
	// Check Order Obj ID
	// --------------------------------------------------------------	
	if createOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Order" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOrderRequestResult := dao.GetOrder( uint64(createOrderObj.ID) )
	
	if getOrderRequestResult.Success == false {
		t.Errorf(getOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get Order success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOrderObj,_ := getOrderRequestResult.Data. (model.Order)
	compareOrder := cmp.Equal(createOrderObj.ID, getOrderObj.ID)
	
	if  compareOrder == false	{
		t.Errorf( "Created Order object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOrderRequestResult := dao.GetAllOrder()

	if getAllOrderRequestResult.Success == false {
			t.Errorf(getAllOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Order success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOrderObj []model.Order = getAllOrderRequestResult.Data. ([]model.Order)
		
	equalOrder := cmp.Equal(createOrderObj.ID, getAllOrderObj[len(getAllOrderObj)-1].ID)
		
	if equalOrder == false {
		t.Errorf( "Created object is not equal to the last entry in Order[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Order
	// --------------------------------------------------------------	
	deleteOrderRequestResult := dao.DeleteOrder(uint64(createOrderObj.ID))

	if deleteOrderRequestResult.Success == false {
			t.Errorf(deleteOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Order success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOrderRequestResult = dao.GetOrder( uint64(createOrderObj.ID) )
	
	if getOrderRequestResult.Success == true {
		t.Errorf(getOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestOrderItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for OrderItem
	//----------------------------------------------------------------------------
	OrderItemObj := model.OrderItem                                                                                                                    {Quantity:"test value",UnitPrice:new Money(),DiscountAmount:new Money(),TaxAmount:new Money(),TotalAmount:new Money()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOrderItemRequestResult := dao.CreateOrderItem( OrderItemObj )
	
	if createOrderItemRequestResult.Success == false {
		t.Errorf(createOrderItemRequestResult.Msg)
	} else {
		fmt.Println("Check Create OrderItem success...")
	}
	
	createOrderItemObj,_ := createOrderItemRequestResult.Data. (model.OrderItem)

	// --------------------------------------------------------------
	// Check OrderItem Obj ID
	// --------------------------------------------------------------	
	if createOrderItemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for OrderItem" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOrderItemRequestResult := dao.GetOrderItem( uint64(createOrderItemObj.ID) )
	
	if getOrderItemRequestResult.Success == false {
		t.Errorf(getOrderItemRequestResult.Msg)
	} else {
		fmt.Println("Check Get OrderItem success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOrderItemObj,_ := getOrderItemRequestResult.Data. (model.OrderItem)
	compareOrderItem := cmp.Equal(createOrderItemObj.ID, getOrderItemObj.ID)
	
	if  compareOrderItem == false	{
		t.Errorf( "Created OrderItem object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOrderItemRequestResult := dao.GetAllOrderItem()

	if getAllOrderItemRequestResult.Success == false {
			t.Errorf(getAllOrderItemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll OrderItem success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOrderItemObj []model.OrderItem = getAllOrderItemRequestResult.Data. ([]model.OrderItem)
		
	equalOrderItem := cmp.Equal(createOrderItemObj.ID, getAllOrderItemObj[len(getAllOrderItemObj)-1].ID)
		
	if equalOrderItem == false {
		t.Errorf( "Created object is not equal to the last entry in OrderItem[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for OrderItem
	// --------------------------------------------------------------	
	deleteOrderItemRequestResult := dao.DeleteOrderItem(uint64(createOrderItemObj.ID))

	if deleteOrderItemRequestResult.Success == false {
			t.Errorf(deleteOrderItemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion OrderItem success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOrderItemRequestResult = dao.GetOrderItem( uint64(createOrderItemObj.ID) )
	
	if getOrderItemRequestResult.Success == true {
		t.Errorf(getOrderItemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestContractCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Contract
	//----------------------------------------------------------------------------
	ContractObj := model.Contract                                                                                                                                                                                                                            {ContractNumber:"test value for ContractNumber",StartDate:time.Now(),EndDate:time.Now(),RenewalTermMonths:100,AutoRenew:true,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createContractRequestResult := dao.CreateContract( ContractObj )
	
	if createContractRequestResult.Success == false {
		t.Errorf(createContractRequestResult.Msg)
	} else {
		fmt.Println("Check Create Contract success...")
	}
	
	createContractObj,_ := createContractRequestResult.Data. (model.Contract)

	// --------------------------------------------------------------
	// Check Contract Obj ID
	// --------------------------------------------------------------	
	if createContractObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Contract" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getContractRequestResult := dao.GetContract( uint64(createContractObj.ID) )
	
	if getContractRequestResult.Success == false {
		t.Errorf(getContractRequestResult.Msg)
	} else {
		fmt.Println("Check Get Contract success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getContractObj,_ := getContractRequestResult.Data. (model.Contract)
	compareContract := cmp.Equal(createContractObj.ID, getContractObj.ID)
	
	if  compareContract == false	{
		t.Errorf( "Created Contract object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllContractRequestResult := dao.GetAllContract()

	if getAllContractRequestResult.Success == false {
			t.Errorf(getAllContractRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Contract success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllContractObj []model.Contract = getAllContractRequestResult.Data. ([]model.Contract)
		
	equalContract := cmp.Equal(createContractObj.ID, getAllContractObj[len(getAllContractObj)-1].ID)
		
	if equalContract == false {
		t.Errorf( "Created object is not equal to the last entry in Contract[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Contract
	// --------------------------------------------------------------	
	deleteContractRequestResult := dao.DeleteContract(uint64(createContractObj.ID))

	if deleteContractRequestResult.Success == false {
			t.Errorf(deleteContractRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Contract success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getContractRequestResult = dao.GetContract( uint64(createContractObj.ID) )
	
	if getContractRequestResult.Success == true {
		t.Errorf(getContractRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCase_CRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Case_
	//----------------------------------------------------------------------------
	Case_Obj := model.Case_                                                                                                                                                                                                                    {CaseNumber:"test value for CaseNumber",Subject:"test value for Subject",Description:"test value for Description",SlaDue:time.Now(),Status:0,Priority:0,Origin:0,Severity:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCase_RequestResult := dao.CreateCase_( Case_Obj )
	
	if createCase_RequestResult.Success == false {
		t.Errorf(createCase_RequestResult.Msg)
	} else {
		fmt.Println("Check Create Case_ success...")
	}
	
	createCase_Obj,_ := createCase_RequestResult.Data. (model.Case_)

	// --------------------------------------------------------------
	// Check Case_ Obj ID
	// --------------------------------------------------------------	
	if createCase_Obj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Case_" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCase_RequestResult := dao.GetCase_( uint64(createCase_Obj.ID) )
	
	if getCase_RequestResult.Success == false {
		t.Errorf(getCase_RequestResult.Msg)
	} else {
		fmt.Println("Check Get Case_ success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCase_Obj,_ := getCase_RequestResult.Data. (model.Case_)
	compareCase_ := cmp.Equal(createCase_Obj.ID, getCase_Obj.ID)
	
	if  compareCase_ == false	{
		t.Errorf( "Created Case_ object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCase_RequestResult := dao.GetAllCase_()

	if getAllCase_RequestResult.Success == false {
			t.Errorf(getAllCase_RequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Case_ success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCase_Obj []model.Case_ = getAllCase_RequestResult.Data. ([]model.Case_)
		
	equalCase_ := cmp.Equal(createCase_Obj.ID, getAllCase_Obj[len(getAllCase_Obj)-1].ID)
		
	if equalCase_ == false {
		t.Errorf( "Created object is not equal to the last entry in Case_[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Case_
	// --------------------------------------------------------------	
	deleteCase_RequestResult := dao.DeleteCase_(uint64(createCase_Obj.ID))

	if deleteCase_RequestResult.Success == false {
			t.Errorf(deleteCase_RequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Case_ success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCase_RequestResult = dao.GetCase_( uint64(createCase_Obj.ID) )
	
	if getCase_RequestResult.Success == true {
		t.Errorf(getCase_RequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestActivityCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Activity
	//----------------------------------------------------------------------------
	ActivityObj := model.Activity                                                                                                                                                                                                                                                                                    {Subject:"test value for Subject",DueDate:time.Now(),StartAt:time.Now(),EndAt:time.Now(),Location:"test value for Location",ActivityType:0,Status:0,Priority:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createActivityRequestResult := dao.CreateActivity( ActivityObj )
	
	if createActivityRequestResult.Success == false {
		t.Errorf(createActivityRequestResult.Msg)
	} else {
		fmt.Println("Check Create Activity success...")
	}
	
	createActivityObj,_ := createActivityRequestResult.Data. (model.Activity)

	// --------------------------------------------------------------
	// Check Activity Obj ID
	// --------------------------------------------------------------	
	if createActivityObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Activity" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getActivityRequestResult := dao.GetActivity( uint64(createActivityObj.ID) )
	
	if getActivityRequestResult.Success == false {
		t.Errorf(getActivityRequestResult.Msg)
	} else {
		fmt.Println("Check Get Activity success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getActivityObj,_ := getActivityRequestResult.Data. (model.Activity)
	compareActivity := cmp.Equal(createActivityObj.ID, getActivityObj.ID)
	
	if  compareActivity == false	{
		t.Errorf( "Created Activity object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllActivityRequestResult := dao.GetAllActivity()

	if getAllActivityRequestResult.Success == false {
			t.Errorf(getAllActivityRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Activity success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllActivityObj []model.Activity = getAllActivityRequestResult.Data. ([]model.Activity)
		
	equalActivity := cmp.Equal(createActivityObj.ID, getAllActivityObj[len(getAllActivityObj)-1].ID)
		
	if equalActivity == false {
		t.Errorf( "Created object is not equal to the last entry in Activity[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Activity
	// --------------------------------------------------------------	
	deleteActivityRequestResult := dao.DeleteActivity(uint64(createActivityObj.ID))

	if deleteActivityRequestResult.Success == false {
			t.Errorf(deleteActivityRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Activity success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getActivityRequestResult = dao.GetActivity( uint64(createActivityObj.ID) )
	
	if getActivityRequestResult.Success == true {
		t.Errorf(getActivityRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCampaignCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Campaign
	//----------------------------------------------------------------------------
	CampaignObj := model.Campaign                                                                                                                                                                                                                            {Name:"test value for Name",StartDate:time.Now(),EndDate:time.Now(),Budget:new Money(),ActualCost:new Money(),ExpectedRevenue:new Money(),Status:0,Type:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCampaignRequestResult := dao.CreateCampaign( CampaignObj )
	
	if createCampaignRequestResult.Success == false {
		t.Errorf(createCampaignRequestResult.Msg)
	} else {
		fmt.Println("Check Create Campaign success...")
	}
	
	createCampaignObj,_ := createCampaignRequestResult.Data. (model.Campaign)

	// --------------------------------------------------------------
	// Check Campaign Obj ID
	// --------------------------------------------------------------	
	if createCampaignObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Campaign" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCampaignRequestResult := dao.GetCampaign( uint64(createCampaignObj.ID) )
	
	if getCampaignRequestResult.Success == false {
		t.Errorf(getCampaignRequestResult.Msg)
	} else {
		fmt.Println("Check Get Campaign success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCampaignObj,_ := getCampaignRequestResult.Data. (model.Campaign)
	compareCampaign := cmp.Equal(createCampaignObj.ID, getCampaignObj.ID)
	
	if  compareCampaign == false	{
		t.Errorf( "Created Campaign object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCampaignRequestResult := dao.GetAllCampaign()

	if getAllCampaignRequestResult.Success == false {
			t.Errorf(getAllCampaignRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Campaign success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCampaignObj []model.Campaign = getAllCampaignRequestResult.Data. ([]model.Campaign)
		
	equalCampaign := cmp.Equal(createCampaignObj.ID, getAllCampaignObj[len(getAllCampaignObj)-1].ID)
		
	if equalCampaign == false {
		t.Errorf( "Created object is not equal to the last entry in Campaign[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Campaign
	// --------------------------------------------------------------	
	deleteCampaignRequestResult := dao.DeleteCampaign(uint64(createCampaignObj.ID))

	if deleteCampaignRequestResult.Success == false {
			t.Errorf(deleteCampaignRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Campaign success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCampaignRequestResult = dao.GetCampaign( uint64(createCampaignObj.ID) )
	
	if getCampaignRequestResult.Success == true {
		t.Errorf(getCampaignRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCampaignMemberCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CampaignMember
	//----------------------------------------------------------------------------
	CampaignMemberObj := model.CampaignMember                                                            {Responded:true,Status:0,MemberType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCampaignMemberRequestResult := dao.CreateCampaignMember( CampaignMemberObj )
	
	if createCampaignMemberRequestResult.Success == false {
		t.Errorf(createCampaignMemberRequestResult.Msg)
	} else {
		fmt.Println("Check Create CampaignMember success...")
	}
	
	createCampaignMemberObj,_ := createCampaignMemberRequestResult.Data. (model.CampaignMember)

	// --------------------------------------------------------------
	// Check CampaignMember Obj ID
	// --------------------------------------------------------------	
	if createCampaignMemberObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CampaignMember" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCampaignMemberRequestResult := dao.GetCampaignMember( uint64(createCampaignMemberObj.ID) )
	
	if getCampaignMemberRequestResult.Success == false {
		t.Errorf(getCampaignMemberRequestResult.Msg)
	} else {
		fmt.Println("Check Get CampaignMember success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCampaignMemberObj,_ := getCampaignMemberRequestResult.Data. (model.CampaignMember)
	compareCampaignMember := cmp.Equal(createCampaignMemberObj.ID, getCampaignMemberObj.ID)
	
	if  compareCampaignMember == false	{
		t.Errorf( "Created CampaignMember object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCampaignMemberRequestResult := dao.GetAllCampaignMember()

	if getAllCampaignMemberRequestResult.Success == false {
			t.Errorf(getAllCampaignMemberRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CampaignMember success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCampaignMemberObj []model.CampaignMember = getAllCampaignMemberRequestResult.Data. ([]model.CampaignMember)
		
	equalCampaignMember := cmp.Equal(createCampaignMemberObj.ID, getAllCampaignMemberObj[len(getAllCampaignMemberObj)-1].ID)
		
	if equalCampaignMember == false {
		t.Errorf( "Created object is not equal to the last entry in CampaignMember[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CampaignMember
	// --------------------------------------------------------------	
	deleteCampaignMemberRequestResult := dao.DeleteCampaignMember(uint64(createCampaignMemberObj.ID))

	if deleteCampaignMemberRequestResult.Success == false {
			t.Errorf(deleteCampaignMemberRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CampaignMember success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCampaignMemberRequestResult = dao.GetCampaignMember( uint64(createCampaignMemberObj.ID) )
	
	if getCampaignMemberRequestResult.Success == true {
		t.Errorf(getCampaignMemberRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestNoteCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Note
	//----------------------------------------------------------------------------
	NoteObj := model.Note                                                                                                                                                                            {Title:"test value for Title",Content:"test value for Content",CreatedAt:time.Now(),UpdatedAt:time.Now()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createNoteRequestResult := dao.CreateNote( NoteObj )
	
	if createNoteRequestResult.Success == false {
		t.Errorf(createNoteRequestResult.Msg)
	} else {
		fmt.Println("Check Create Note success...")
	}
	
	createNoteObj,_ := createNoteRequestResult.Data. (model.Note)

	// --------------------------------------------------------------
	// Check Note Obj ID
	// --------------------------------------------------------------	
	if createNoteObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Note" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getNoteRequestResult := dao.GetNote( uint64(createNoteObj.ID) )
	
	if getNoteRequestResult.Success == false {
		t.Errorf(getNoteRequestResult.Msg)
	} else {
		fmt.Println("Check Get Note success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getNoteObj,_ := getNoteRequestResult.Data. (model.Note)
	compareNote := cmp.Equal(createNoteObj.ID, getNoteObj.ID)
	
	if  compareNote == false	{
		t.Errorf( "Created Note object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllNoteRequestResult := dao.GetAllNote()

	if getAllNoteRequestResult.Success == false {
			t.Errorf(getAllNoteRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Note success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllNoteObj []model.Note = getAllNoteRequestResult.Data. ([]model.Note)
		
	equalNote := cmp.Equal(createNoteObj.ID, getAllNoteObj[len(getAllNoteObj)-1].ID)
		
	if equalNote == false {
		t.Errorf( "Created object is not equal to the last entry in Note[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Note
	// --------------------------------------------------------------	
	deleteNoteRequestResult := dao.DeleteNote(uint64(createNoteObj.ID))

	if deleteNoteRequestResult.Success == false {
			t.Errorf(deleteNoteRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Note success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getNoteRequestResult = dao.GetNote( uint64(createNoteObj.ID) )
	
	if getNoteRequestResult.Success == true {
		t.Errorf(getNoteRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEmailMessageCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for EmailMessage
	//----------------------------------------------------------------------------
	EmailMessageObj := model.EmailMessage                                                                                                                                                                                    {Subject:"test value for Subject",Body:"test value for Body",SentAt:time.Now(),MessageId:"test value for MessageId",Direction:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEmailMessageRequestResult := dao.CreateEmailMessage( EmailMessageObj )
	
	if createEmailMessageRequestResult.Success == false {
		t.Errorf(createEmailMessageRequestResult.Msg)
	} else {
		fmt.Println("Check Create EmailMessage success...")
	}
	
	createEmailMessageObj,_ := createEmailMessageRequestResult.Data. (model.EmailMessage)

	// --------------------------------------------------------------
	// Check EmailMessage Obj ID
	// --------------------------------------------------------------	
	if createEmailMessageObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for EmailMessage" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEmailMessageRequestResult := dao.GetEmailMessage( uint64(createEmailMessageObj.ID) )
	
	if getEmailMessageRequestResult.Success == false {
		t.Errorf(getEmailMessageRequestResult.Msg)
	} else {
		fmt.Println("Check Get EmailMessage success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEmailMessageObj,_ := getEmailMessageRequestResult.Data. (model.EmailMessage)
	compareEmailMessage := cmp.Equal(createEmailMessageObj.ID, getEmailMessageObj.ID)
	
	if  compareEmailMessage == false	{
		t.Errorf( "Created EmailMessage object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEmailMessageRequestResult := dao.GetAllEmailMessage()

	if getAllEmailMessageRequestResult.Success == false {
			t.Errorf(getAllEmailMessageRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll EmailMessage success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEmailMessageObj []model.EmailMessage = getAllEmailMessageRequestResult.Data. ([]model.EmailMessage)
		
	equalEmailMessage := cmp.Equal(createEmailMessageObj.ID, getAllEmailMessageObj[len(getAllEmailMessageObj)-1].ID)
		
	if equalEmailMessage == false {
		t.Errorf( "Created object is not equal to the last entry in EmailMessage[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for EmailMessage
	// --------------------------------------------------------------	
	deleteEmailMessageRequestResult := dao.DeleteEmailMessage(uint64(createEmailMessageObj.ID))

	if deleteEmailMessageRequestResult.Success == false {
			t.Errorf(deleteEmailMessageRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion EmailMessage success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEmailMessageRequestResult = dao.GetEmailMessage( uint64(createEmailMessageObj.ID) )
	
	if getEmailMessageRequestResult.Success == true {
		t.Errorf(getEmailMessageRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

