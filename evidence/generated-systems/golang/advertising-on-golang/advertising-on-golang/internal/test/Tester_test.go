package test

import ( 
	"testing"
    dao "advertising-on-golang/internal/dao"
	"advertising-on-golang/internal/model"
	"advertising-on-golang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestAgencyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Agency
	//----------------------------------------------------------------------------
	AgencyObj := model.Agency                                                                                                                            {Name:"test value for Name",LegalName:"test value for LegalName",HeadquartersCountry:"test value for HeadquartersCountry",Website:"test value for Website"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAgencyRequestResult := dao.CreateAgency( AgencyObj )
	
	if createAgencyRequestResult.Success == false {
		t.Errorf(createAgencyRequestResult.Msg)
	} else {
		fmt.Println("Check Create Agency success...")
	}
	
	createAgencyObj,_ := createAgencyRequestResult.Data. (model.Agency)

	// --------------------------------------------------------------
	// Check Agency Obj ID
	// --------------------------------------------------------------	
	if createAgencyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Agency" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAgencyRequestResult := dao.GetAgency( uint64(createAgencyObj.ID) )
	
	if getAgencyRequestResult.Success == false {
		t.Errorf(getAgencyRequestResult.Msg)
	} else {
		fmt.Println("Check Get Agency success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAgencyObj,_ := getAgencyRequestResult.Data. (model.Agency)
	compareAgency := cmp.Equal(createAgencyObj.ID, getAgencyObj.ID)
	
	if  compareAgency == false	{
		t.Errorf( "Created Agency object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAgencyRequestResult := dao.GetAllAgency()

	if getAllAgencyRequestResult.Success == false {
			t.Errorf(getAllAgencyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Agency success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAgencyObj []model.Agency = getAllAgencyRequestResult.Data. ([]model.Agency)
		
	equalAgency := cmp.Equal(createAgencyObj.ID, getAllAgencyObj[len(getAllAgencyObj)-1].ID)
		
	if equalAgency == false {
		t.Errorf( "Created object is not equal to the last entry in Agency[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Agency
	// --------------------------------------------------------------	
	deleteAgencyRequestResult := dao.DeleteAgency(uint64(createAgencyObj.ID))

	if deleteAgencyRequestResult.Success == false {
			t.Errorf(deleteAgencyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Agency success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAgencyRequestResult = dao.GetAgency( uint64(createAgencyObj.ID) )
	
	if getAgencyRequestResult.Success == true {
		t.Errorf(getAgencyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTeamCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Team
	//----------------------------------------------------------------------------
	TeamObj := model.Team                            {Name:"test value for Name"}

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


func TestUserCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for User
	//----------------------------------------------------------------------------
	UserObj := model.User                                                                                            {FirstName:"test value for FirstName",LastName:"test value for LastName",Email:new Email(),Role:0}

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


func TestAdvertiserCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Advertiser
	//----------------------------------------------------------------------------
	AdvertiserObj := model.Advertiser                                                                                                                            {Name:"test value for Name",LegalName:"test value for LegalName",Industry:"test value for Industry",Website:"test value for Website"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAdvertiserRequestResult := dao.CreateAdvertiser( AdvertiserObj )
	
	if createAdvertiserRequestResult.Success == false {
		t.Errorf(createAdvertiserRequestResult.Msg)
	} else {
		fmt.Println("Check Create Advertiser success...")
	}
	
	createAdvertiserObj,_ := createAdvertiserRequestResult.Data. (model.Advertiser)

	// --------------------------------------------------------------
	// Check Advertiser Obj ID
	// --------------------------------------------------------------	
	if createAdvertiserObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Advertiser" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAdvertiserRequestResult := dao.GetAdvertiser( uint64(createAdvertiserObj.ID) )
	
	if getAdvertiserRequestResult.Success == false {
		t.Errorf(getAdvertiserRequestResult.Msg)
	} else {
		fmt.Println("Check Get Advertiser success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAdvertiserObj,_ := getAdvertiserRequestResult.Data. (model.Advertiser)
	compareAdvertiser := cmp.Equal(createAdvertiserObj.ID, getAdvertiserObj.ID)
	
	if  compareAdvertiser == false	{
		t.Errorf( "Created Advertiser object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAdvertiserRequestResult := dao.GetAllAdvertiser()

	if getAllAdvertiserRequestResult.Success == false {
			t.Errorf(getAllAdvertiserRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Advertiser success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAdvertiserObj []model.Advertiser = getAllAdvertiserRequestResult.Data. ([]model.Advertiser)
		
	equalAdvertiser := cmp.Equal(createAdvertiserObj.ID, getAllAdvertiserObj[len(getAllAdvertiserObj)-1].ID)
		
	if equalAdvertiser == false {
		t.Errorf( "Created object is not equal to the last entry in Advertiser[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Advertiser
	// --------------------------------------------------------------	
	deleteAdvertiserRequestResult := dao.DeleteAdvertiser(uint64(createAdvertiserObj.ID))

	if deleteAdvertiserRequestResult.Success == false {
			t.Errorf(deleteAdvertiserRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Advertiser success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAdvertiserRequestResult = dao.GetAdvertiser( uint64(createAdvertiserObj.ID) )
	
	if getAdvertiserRequestResult.Success == true {
		t.Errorf(getAdvertiserRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBillingProfileCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BillingProfile
	//----------------------------------------------------------------------------
	BillingProfileObj := model.BillingProfile                                                                                            {BillingName:"test value for BillingName",TaxId:"test value for TaxId",BillingAddress:new Address(),PaymentTerms:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBillingProfileRequestResult := dao.CreateBillingProfile( BillingProfileObj )
	
	if createBillingProfileRequestResult.Success == false {
		t.Errorf(createBillingProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Create BillingProfile success...")
	}
	
	createBillingProfileObj,_ := createBillingProfileRequestResult.Data. (model.BillingProfile)

	// --------------------------------------------------------------
	// Check BillingProfile Obj ID
	// --------------------------------------------------------------	
	if createBillingProfileObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BillingProfile" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBillingProfileRequestResult := dao.GetBillingProfile( uint64(createBillingProfileObj.ID) )
	
	if getBillingProfileRequestResult.Success == false {
		t.Errorf(getBillingProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Get BillingProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBillingProfileObj,_ := getBillingProfileRequestResult.Data. (model.BillingProfile)
	compareBillingProfile := cmp.Equal(createBillingProfileObj.ID, getBillingProfileObj.ID)
	
	if  compareBillingProfile == false	{
		t.Errorf( "Created BillingProfile object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBillingProfileRequestResult := dao.GetAllBillingProfile()

	if getAllBillingProfileRequestResult.Success == false {
			t.Errorf(getAllBillingProfileRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BillingProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBillingProfileObj []model.BillingProfile = getAllBillingProfileRequestResult.Data. ([]model.BillingProfile)
		
	equalBillingProfile := cmp.Equal(createBillingProfileObj.ID, getAllBillingProfileObj[len(getAllBillingProfileObj)-1].ID)
		
	if equalBillingProfile == false {
		t.Errorf( "Created object is not equal to the last entry in BillingProfile[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BillingProfile
	// --------------------------------------------------------------	
	deleteBillingProfileRequestResult := dao.DeleteBillingProfile(uint64(createBillingProfileObj.ID))

	if deleteBillingProfileRequestResult.Success == false {
			t.Errorf(deleteBillingProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BillingProfile success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBillingProfileRequestResult = dao.GetBillingProfile( uint64(createBillingProfileObj.ID) )
	
	if getBillingProfileRequestResult.Success == true {
		t.Errorf(getBillingProfileRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPaymentMethodCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PaymentMethod
	//----------------------------------------------------------------------------
	PaymentMethodObj := model.PaymentMethod                                                                                            {Last4:"test value for Last4",CardholderName:"test value for CardholderName",BillingAddress:new Address(),MethodType:0}

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


func TestAdAccountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AdAccount
	//----------------------------------------------------------------------------
	AdAccountObj := model.AdAccount                                                                                                                            {Name:"test value for Name",AccountCode:"test value for AccountCode",DefaultCurrency:"test value for DefaultCurrency",DefaultTimezone:"test value for DefaultTimezone"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAdAccountRequestResult := dao.CreateAdAccount( AdAccountObj )
	
	if createAdAccountRequestResult.Success == false {
		t.Errorf(createAdAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Create AdAccount success...")
	}
	
	createAdAccountObj,_ := createAdAccountRequestResult.Data. (model.AdAccount)

	// --------------------------------------------------------------
	// Check AdAccount Obj ID
	// --------------------------------------------------------------	
	if createAdAccountObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AdAccount" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAdAccountRequestResult := dao.GetAdAccount( uint64(createAdAccountObj.ID) )
	
	if getAdAccountRequestResult.Success == false {
		t.Errorf(getAdAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Get AdAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAdAccountObj,_ := getAdAccountRequestResult.Data. (model.AdAccount)
	compareAdAccount := cmp.Equal(createAdAccountObj.ID, getAdAccountObj.ID)
	
	if  compareAdAccount == false	{
		t.Errorf( "Created AdAccount object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAdAccountRequestResult := dao.GetAllAdAccount()

	if getAllAdAccountRequestResult.Success == false {
			t.Errorf(getAllAdAccountRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AdAccount success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAdAccountObj []model.AdAccount = getAllAdAccountRequestResult.Data. ([]model.AdAccount)
		
	equalAdAccount := cmp.Equal(createAdAccountObj.ID, getAllAdAccountObj[len(getAllAdAccountObj)-1].ID)
		
	if equalAdAccount == false {
		t.Errorf( "Created object is not equal to the last entry in AdAccount[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AdAccount
	// --------------------------------------------------------------	
	deleteAdAccountRequestResult := dao.DeleteAdAccount(uint64(createAdAccountObj.ID))

	if deleteAdAccountRequestResult.Success == false {
			t.Errorf(deleteAdAccountRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AdAccount success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAdAccountRequestResult = dao.GetAdAccount( uint64(createAdAccountObj.ID) )
	
	if getAdAccountRequestResult.Success == true {
		t.Errorf(getAdAccountRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDSPCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DSP
	//----------------------------------------------------------------------------
	DSPObj := model.DSP                                                                                            {Name:"test value for Name",Website:"test value for Website",Region:"test value for Region"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDSPRequestResult := dao.CreateDSP( DSPObj )
	
	if createDSPRequestResult.Success == false {
		t.Errorf(createDSPRequestResult.Msg)
	} else {
		fmt.Println("Check Create DSP success...")
	}
	
	createDSPObj,_ := createDSPRequestResult.Data. (model.DSP)

	// --------------------------------------------------------------
	// Check DSP Obj ID
	// --------------------------------------------------------------	
	if createDSPObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DSP" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDSPRequestResult := dao.GetDSP( uint64(createDSPObj.ID) )
	
	if getDSPRequestResult.Success == false {
		t.Errorf(getDSPRequestResult.Msg)
	} else {
		fmt.Println("Check Get DSP success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDSPObj,_ := getDSPRequestResult.Data. (model.DSP)
	compareDSP := cmp.Equal(createDSPObj.ID, getDSPObj.ID)
	
	if  compareDSP == false	{
		t.Errorf( "Created DSP object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDSPRequestResult := dao.GetAllDSP()

	if getAllDSPRequestResult.Success == false {
			t.Errorf(getAllDSPRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DSP success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDSPObj []model.DSP = getAllDSPRequestResult.Data. ([]model.DSP)
		
	equalDSP := cmp.Equal(createDSPObj.ID, getAllDSPObj[len(getAllDSPObj)-1].ID)
		
	if equalDSP == false {
		t.Errorf( "Created object is not equal to the last entry in DSP[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DSP
	// --------------------------------------------------------------	
	deleteDSPRequestResult := dao.DeleteDSP(uint64(createDSPObj.ID))

	if deleteDSPRequestResult.Success == false {
			t.Errorf(deleteDSPRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DSP success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDSPRequestResult = dao.GetDSP( uint64(createDSPObj.ID) )
	
	if getDSPRequestResult.Success == true {
		t.Errorf(getDSPRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCampaignCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Campaign
	//----------------------------------------------------------------------------
	CampaignObj := model.Campaign                                                                                            {Name:"test value for Name",TotalBudget:new Money(),Flight:new DateRange(),Objective:0,Status:0}

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


func TestKPICRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for KPI
	//----------------------------------------------------------------------------
	KPIObj := model.KPI                                                                    {TargetValue:"test value",MetricType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createKPIRequestResult := dao.CreateKPI( KPIObj )
	
	if createKPIRequestResult.Success == false {
		t.Errorf(createKPIRequestResult.Msg)
	} else {
		fmt.Println("Check Create KPI success...")
	}
	
	createKPIObj,_ := createKPIRequestResult.Data. (model.KPI)

	// --------------------------------------------------------------
	// Check KPI Obj ID
	// --------------------------------------------------------------	
	if createKPIObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for KPI" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getKPIRequestResult := dao.GetKPI( uint64(createKPIObj.ID) )
	
	if getKPIRequestResult.Success == false {
		t.Errorf(getKPIRequestResult.Msg)
	} else {
		fmt.Println("Check Get KPI success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getKPIObj,_ := getKPIRequestResult.Data. (model.KPI)
	compareKPI := cmp.Equal(createKPIObj.ID, getKPIObj.ID)
	
	if  compareKPI == false	{
		t.Errorf( "Created KPI object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllKPIRequestResult := dao.GetAllKPI()

	if getAllKPIRequestResult.Success == false {
			t.Errorf(getAllKPIRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll KPI success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllKPIObj []model.KPI = getAllKPIRequestResult.Data. ([]model.KPI)
		
	equalKPI := cmp.Equal(createKPIObj.ID, getAllKPIObj[len(getAllKPIObj)-1].ID)
		
	if equalKPI == false {
		t.Errorf( "Created object is not equal to the last entry in KPI[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for KPI
	// --------------------------------------------------------------	
	deleteKPIRequestResult := dao.DeleteKPI(uint64(createKPIObj.ID))

	if deleteKPIRequestResult.Success == false {
			t.Errorf(deleteKPIRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion KPI success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getKPIRequestResult = dao.GetKPI( uint64(createKPIObj.ID) )
	
	if getKPIRequestResult.Success == true {
		t.Errorf(getKPIRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAudienceSegmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AudienceSegment
	//----------------------------------------------------------------------------
	AudienceSegmentObj := model.AudienceSegment                                                                                                            {Name:"test value for Name",EstimatedReach:100,Description:"test value for Description",ProviderType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAudienceSegmentRequestResult := dao.CreateAudienceSegment( AudienceSegmentObj )
	
	if createAudienceSegmentRequestResult.Success == false {
		t.Errorf(createAudienceSegmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create AudienceSegment success...")
	}
	
	createAudienceSegmentObj,_ := createAudienceSegmentRequestResult.Data. (model.AudienceSegment)

	// --------------------------------------------------------------
	// Check AudienceSegment Obj ID
	// --------------------------------------------------------------	
	if createAudienceSegmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AudienceSegment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAudienceSegmentRequestResult := dao.GetAudienceSegment( uint64(createAudienceSegmentObj.ID) )
	
	if getAudienceSegmentRequestResult.Success == false {
		t.Errorf(getAudienceSegmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get AudienceSegment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAudienceSegmentObj,_ := getAudienceSegmentRequestResult.Data. (model.AudienceSegment)
	compareAudienceSegment := cmp.Equal(createAudienceSegmentObj.ID, getAudienceSegmentObj.ID)
	
	if  compareAudienceSegment == false	{
		t.Errorf( "Created AudienceSegment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAudienceSegmentRequestResult := dao.GetAllAudienceSegment()

	if getAllAudienceSegmentRequestResult.Success == false {
			t.Errorf(getAllAudienceSegmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AudienceSegment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAudienceSegmentObj []model.AudienceSegment = getAllAudienceSegmentRequestResult.Data. ([]model.AudienceSegment)
		
	equalAudienceSegment := cmp.Equal(createAudienceSegmentObj.ID, getAllAudienceSegmentObj[len(getAllAudienceSegmentObj)-1].ID)
		
	if equalAudienceSegment == false {
		t.Errorf( "Created object is not equal to the last entry in AudienceSegment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AudienceSegment
	// --------------------------------------------------------------	
	deleteAudienceSegmentRequestResult := dao.DeleteAudienceSegment(uint64(createAudienceSegmentObj.ID))

	if deleteAudienceSegmentRequestResult.Success == false {
			t.Errorf(deleteAudienceSegmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AudienceSegment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAudienceSegmentRequestResult = dao.GetAudienceSegment( uint64(createAudienceSegmentObj.ID) )
	
	if getAudienceSegmentRequestResult.Success == true {
		t.Errorf(getAudienceSegmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDataProviderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DataProvider
	//----------------------------------------------------------------------------
	DataProviderObj := model.DataProvider                                                                            {Name:"test value for Name",Website:"test value for Website",ProviderType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDataProviderRequestResult := dao.CreateDataProvider( DataProviderObj )
	
	if createDataProviderRequestResult.Success == false {
		t.Errorf(createDataProviderRequestResult.Msg)
	} else {
		fmt.Println("Check Create DataProvider success...")
	}
	
	createDataProviderObj,_ := createDataProviderRequestResult.Data. (model.DataProvider)

	// --------------------------------------------------------------
	// Check DataProvider Obj ID
	// --------------------------------------------------------------	
	if createDataProviderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DataProvider" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDataProviderRequestResult := dao.GetDataProvider( uint64(createDataProviderObj.ID) )
	
	if getDataProviderRequestResult.Success == false {
		t.Errorf(getDataProviderRequestResult.Msg)
	} else {
		fmt.Println("Check Get DataProvider success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDataProviderObj,_ := getDataProviderRequestResult.Data. (model.DataProvider)
	compareDataProvider := cmp.Equal(createDataProviderObj.ID, getDataProviderObj.ID)
	
	if  compareDataProvider == false	{
		t.Errorf( "Created DataProvider object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDataProviderRequestResult := dao.GetAllDataProvider()

	if getAllDataProviderRequestResult.Success == false {
			t.Errorf(getAllDataProviderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DataProvider success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDataProviderObj []model.DataProvider = getAllDataProviderRequestResult.Data. ([]model.DataProvider)
		
	equalDataProvider := cmp.Equal(createDataProviderObj.ID, getAllDataProviderObj[len(getAllDataProviderObj)-1].ID)
		
	if equalDataProvider == false {
		t.Errorf( "Created object is not equal to the last entry in DataProvider[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DataProvider
	// --------------------------------------------------------------	
	deleteDataProviderRequestResult := dao.DeleteDataProvider(uint64(createDataProviderObj.ID))

	if deleteDataProviderRequestResult.Success == false {
			t.Errorf(deleteDataProviderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DataProvider success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDataProviderRequestResult = dao.GetDataProvider( uint64(createDataProviderObj.ID) )
	
	if getDataProviderRequestResult.Success == true {
		t.Errorf(getDataProviderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLineItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LineItem
	//----------------------------------------------------------------------------
	LineItemObj := model.LineItem                                                                                                                                            {Name:"test value for Name",BidAmount:new Money(),DailyBudget:new Money(),FrequencyCap:new FrequencyCap(),Status:0,PricingModel:0,BidStrategy:0,Pacing:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLineItemRequestResult := dao.CreateLineItem( LineItemObj )
	
	if createLineItemRequestResult.Success == false {
		t.Errorf(createLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check Create LineItem success...")
	}
	
	createLineItemObj,_ := createLineItemRequestResult.Data. (model.LineItem)

	// --------------------------------------------------------------
	// Check LineItem Obj ID
	// --------------------------------------------------------------	
	if createLineItemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LineItem" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLineItemRequestResult := dao.GetLineItem( uint64(createLineItemObj.ID) )
	
	if getLineItemRequestResult.Success == false {
		t.Errorf(getLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check Get LineItem success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLineItemObj,_ := getLineItemRequestResult.Data. (model.LineItem)
	compareLineItem := cmp.Equal(createLineItemObj.ID, getLineItemObj.ID)
	
	if  compareLineItem == false	{
		t.Errorf( "Created LineItem object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLineItemRequestResult := dao.GetAllLineItem()

	if getAllLineItemRequestResult.Success == false {
			t.Errorf(getAllLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LineItem success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLineItemObj []model.LineItem = getAllLineItemRequestResult.Data. ([]model.LineItem)
		
	equalLineItem := cmp.Equal(createLineItemObj.ID, getAllLineItemObj[len(getAllLineItemObj)-1].ID)
		
	if equalLineItem == false {
		t.Errorf( "Created object is not equal to the last entry in LineItem[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LineItem
	// --------------------------------------------------------------	
	deleteLineItemRequestResult := dao.DeleteLineItem(uint64(createLineItemObj.ID))

	if deleteLineItemRequestResult.Success == false {
			t.Errorf(deleteLineItemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LineItem success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLineItemRequestResult = dao.GetLineItem( uint64(createLineItemObj.ID) )
	
	if getLineItemRequestResult.Success == true {
		t.Errorf(getLineItemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTargetingProfileCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TargetingProfile
	//----------------------------------------------------------------------------
	TargetingProfileObj := model.TargetingProfile                            {Name:"test value for Name"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTargetingProfileRequestResult := dao.CreateTargetingProfile( TargetingProfileObj )
	
	if createTargetingProfileRequestResult.Success == false {
		t.Errorf(createTargetingProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Create TargetingProfile success...")
	}
	
	createTargetingProfileObj,_ := createTargetingProfileRequestResult.Data. (model.TargetingProfile)

	// --------------------------------------------------------------
	// Check TargetingProfile Obj ID
	// --------------------------------------------------------------	
	if createTargetingProfileObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TargetingProfile" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTargetingProfileRequestResult := dao.GetTargetingProfile( uint64(createTargetingProfileObj.ID) )
	
	if getTargetingProfileRequestResult.Success == false {
		t.Errorf(getTargetingProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Get TargetingProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTargetingProfileObj,_ := getTargetingProfileRequestResult.Data. (model.TargetingProfile)
	compareTargetingProfile := cmp.Equal(createTargetingProfileObj.ID, getTargetingProfileObj.ID)
	
	if  compareTargetingProfile == false	{
		t.Errorf( "Created TargetingProfile object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTargetingProfileRequestResult := dao.GetAllTargetingProfile()

	if getAllTargetingProfileRequestResult.Success == false {
			t.Errorf(getAllTargetingProfileRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TargetingProfile success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTargetingProfileObj []model.TargetingProfile = getAllTargetingProfileRequestResult.Data. ([]model.TargetingProfile)
		
	equalTargetingProfile := cmp.Equal(createTargetingProfileObj.ID, getAllTargetingProfileObj[len(getAllTargetingProfileObj)-1].ID)
		
	if equalTargetingProfile == false {
		t.Errorf( "Created object is not equal to the last entry in TargetingProfile[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TargetingProfile
	// --------------------------------------------------------------	
	deleteTargetingProfileRequestResult := dao.DeleteTargetingProfile(uint64(createTargetingProfileObj.ID))

	if deleteTargetingProfileRequestResult.Success == false {
			t.Errorf(deleteTargetingProfileRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TargetingProfile success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTargetingProfileRequestResult = dao.GetTargetingProfile( uint64(createTargetingProfileObj.ID) )
	
	if getTargetingProfileRequestResult.Success == true {
		t.Errorf(getTargetingProfileRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDeviceCriterionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DeviceCriterion
	//----------------------------------------------------------------------------
	DeviceCriterionObj := model.DeviceCriterion                                            {DeviceType:0,PlatformType:0,Operator:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDeviceCriterionRequestResult := dao.CreateDeviceCriterion( DeviceCriterionObj )
	
	if createDeviceCriterionRequestResult.Success == false {
		t.Errorf(createDeviceCriterionRequestResult.Msg)
	} else {
		fmt.Println("Check Create DeviceCriterion success...")
	}
	
	createDeviceCriterionObj,_ := createDeviceCriterionRequestResult.Data. (model.DeviceCriterion)

	// --------------------------------------------------------------
	// Check DeviceCriterion Obj ID
	// --------------------------------------------------------------	
	if createDeviceCriterionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DeviceCriterion" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDeviceCriterionRequestResult := dao.GetDeviceCriterion( uint64(createDeviceCriterionObj.ID) )
	
	if getDeviceCriterionRequestResult.Success == false {
		t.Errorf(getDeviceCriterionRequestResult.Msg)
	} else {
		fmt.Println("Check Get DeviceCriterion success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDeviceCriterionObj,_ := getDeviceCriterionRequestResult.Data. (model.DeviceCriterion)
	compareDeviceCriterion := cmp.Equal(createDeviceCriterionObj.ID, getDeviceCriterionObj.ID)
	
	if  compareDeviceCriterion == false	{
		t.Errorf( "Created DeviceCriterion object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDeviceCriterionRequestResult := dao.GetAllDeviceCriterion()

	if getAllDeviceCriterionRequestResult.Success == false {
			t.Errorf(getAllDeviceCriterionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DeviceCriterion success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDeviceCriterionObj []model.DeviceCriterion = getAllDeviceCriterionRequestResult.Data. ([]model.DeviceCriterion)
		
	equalDeviceCriterion := cmp.Equal(createDeviceCriterionObj.ID, getAllDeviceCriterionObj[len(getAllDeviceCriterionObj)-1].ID)
		
	if equalDeviceCriterion == false {
		t.Errorf( "Created object is not equal to the last entry in DeviceCriterion[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DeviceCriterion
	// --------------------------------------------------------------	
	deleteDeviceCriterionRequestResult := dao.DeleteDeviceCriterion(uint64(createDeviceCriterionObj.ID))

	if deleteDeviceCriterionRequestResult.Success == false {
			t.Errorf(deleteDeviceCriterionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DeviceCriterion success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDeviceCriterionRequestResult = dao.GetDeviceCriterion( uint64(createDeviceCriterionObj.ID) )
	
	if getDeviceCriterionRequestResult.Success == true {
		t.Errorf(getDeviceCriterionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBrandSafetyPolicyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BrandSafetyPolicy
	//----------------------------------------------------------------------------
	BrandSafetyPolicyObj := model.BrandSafetyPolicy                            {Level:0,ContentRatingThreshold:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBrandSafetyPolicyRequestResult := dao.CreateBrandSafetyPolicy( BrandSafetyPolicyObj )
	
	if createBrandSafetyPolicyRequestResult.Success == false {
		t.Errorf(createBrandSafetyPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Create BrandSafetyPolicy success...")
	}
	
	createBrandSafetyPolicyObj,_ := createBrandSafetyPolicyRequestResult.Data. (model.BrandSafetyPolicy)

	// --------------------------------------------------------------
	// Check BrandSafetyPolicy Obj ID
	// --------------------------------------------------------------	
	if createBrandSafetyPolicyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BrandSafetyPolicy" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBrandSafetyPolicyRequestResult := dao.GetBrandSafetyPolicy( uint64(createBrandSafetyPolicyObj.ID) )
	
	if getBrandSafetyPolicyRequestResult.Success == false {
		t.Errorf(getBrandSafetyPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Get BrandSafetyPolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBrandSafetyPolicyObj,_ := getBrandSafetyPolicyRequestResult.Data. (model.BrandSafetyPolicy)
	compareBrandSafetyPolicy := cmp.Equal(createBrandSafetyPolicyObj.ID, getBrandSafetyPolicyObj.ID)
	
	if  compareBrandSafetyPolicy == false	{
		t.Errorf( "Created BrandSafetyPolicy object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBrandSafetyPolicyRequestResult := dao.GetAllBrandSafetyPolicy()

	if getAllBrandSafetyPolicyRequestResult.Success == false {
			t.Errorf(getAllBrandSafetyPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BrandSafetyPolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBrandSafetyPolicyObj []model.BrandSafetyPolicy = getAllBrandSafetyPolicyRequestResult.Data. ([]model.BrandSafetyPolicy)
		
	equalBrandSafetyPolicy := cmp.Equal(createBrandSafetyPolicyObj.ID, getAllBrandSafetyPolicyObj[len(getAllBrandSafetyPolicyObj)-1].ID)
		
	if equalBrandSafetyPolicy == false {
		t.Errorf( "Created object is not equal to the last entry in BrandSafetyPolicy[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BrandSafetyPolicy
	// --------------------------------------------------------------	
	deleteBrandSafetyPolicyRequestResult := dao.DeleteBrandSafetyPolicy(uint64(createBrandSafetyPolicyObj.ID))

	if deleteBrandSafetyPolicyRequestResult.Success == false {
			t.Errorf(deleteBrandSafetyPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BrandSafetyPolicy success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBrandSafetyPolicyRequestResult = dao.GetBrandSafetyPolicy( uint64(createBrandSafetyPolicyObj.ID) )
	
	if getBrandSafetyPolicyRequestResult.Success == true {
		t.Errorf(getBrandSafetyPolicyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestContentCategoryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ContentCategory
	//----------------------------------------------------------------------------
	ContentCategoryObj := model.ContentCategory                                                            {Code:"test value for Code",Name:"test value for Name"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createContentCategoryRequestResult := dao.CreateContentCategory( ContentCategoryObj )
	
	if createContentCategoryRequestResult.Success == false {
		t.Errorf(createContentCategoryRequestResult.Msg)
	} else {
		fmt.Println("Check Create ContentCategory success...")
	}
	
	createContentCategoryObj,_ := createContentCategoryRequestResult.Data. (model.ContentCategory)

	// --------------------------------------------------------------
	// Check ContentCategory Obj ID
	// --------------------------------------------------------------	
	if createContentCategoryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ContentCategory" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getContentCategoryRequestResult := dao.GetContentCategory( uint64(createContentCategoryObj.ID) )
	
	if getContentCategoryRequestResult.Success == false {
		t.Errorf(getContentCategoryRequestResult.Msg)
	} else {
		fmt.Println("Check Get ContentCategory success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getContentCategoryObj,_ := getContentCategoryRequestResult.Data. (model.ContentCategory)
	compareContentCategory := cmp.Equal(createContentCategoryObj.ID, getContentCategoryObj.ID)
	
	if  compareContentCategory == false	{
		t.Errorf( "Created ContentCategory object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllContentCategoryRequestResult := dao.GetAllContentCategory()

	if getAllContentCategoryRequestResult.Success == false {
			t.Errorf(getAllContentCategoryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ContentCategory success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllContentCategoryObj []model.ContentCategory = getAllContentCategoryRequestResult.Data. ([]model.ContentCategory)
		
	equalContentCategory := cmp.Equal(createContentCategoryObj.ID, getAllContentCategoryObj[len(getAllContentCategoryObj)-1].ID)
		
	if equalContentCategory == false {
		t.Errorf( "Created object is not equal to the last entry in ContentCategory[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ContentCategory
	// --------------------------------------------------------------	
	deleteContentCategoryRequestResult := dao.DeleteContentCategory(uint64(createContentCategoryObj.ID))

	if deleteContentCategoryRequestResult.Success == false {
			t.Errorf(deleteContentCategoryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ContentCategory success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getContentCategoryRequestResult = dao.GetContentCategory( uint64(createContentCategoryObj.ID) )
	
	if getContentCategoryRequestResult.Success == true {
		t.Errorf(getContentCategoryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPublisherCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Publisher
	//----------------------------------------------------------------------------
	PublisherObj := model.Publisher                                                                            {Name:"test value for Name",Website:"test value for Website",PublisherType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPublisherRequestResult := dao.CreatePublisher( PublisherObj )
	
	if createPublisherRequestResult.Success == false {
		t.Errorf(createPublisherRequestResult.Msg)
	} else {
		fmt.Println("Check Create Publisher success...")
	}
	
	createPublisherObj,_ := createPublisherRequestResult.Data. (model.Publisher)

	// --------------------------------------------------------------
	// Check Publisher Obj ID
	// --------------------------------------------------------------	
	if createPublisherObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Publisher" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPublisherRequestResult := dao.GetPublisher( uint64(createPublisherObj.ID) )
	
	if getPublisherRequestResult.Success == false {
		t.Errorf(getPublisherRequestResult.Msg)
	} else {
		fmt.Println("Check Get Publisher success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPublisherObj,_ := getPublisherRequestResult.Data. (model.Publisher)
	comparePublisher := cmp.Equal(createPublisherObj.ID, getPublisherObj.ID)
	
	if  comparePublisher == false	{
		t.Errorf( "Created Publisher object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPublisherRequestResult := dao.GetAllPublisher()

	if getAllPublisherRequestResult.Success == false {
			t.Errorf(getAllPublisherRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Publisher success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPublisherObj []model.Publisher = getAllPublisherRequestResult.Data. ([]model.Publisher)
		
	equalPublisher := cmp.Equal(createPublisherObj.ID, getAllPublisherObj[len(getAllPublisherObj)-1].ID)
		
	if equalPublisher == false {
		t.Errorf( "Created object is not equal to the last entry in Publisher[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Publisher
	// --------------------------------------------------------------	
	deletePublisherRequestResult := dao.DeletePublisher(uint64(createPublisherObj.ID))

	if deletePublisherRequestResult.Success == false {
			t.Errorf(deletePublisherRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Publisher success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPublisherRequestResult = dao.GetPublisher( uint64(createPublisherObj.ID) )
	
	if getPublisherRequestResult.Success == true {
		t.Errorf(getPublisherRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInventorySourceCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InventorySource
	//----------------------------------------------------------------------------
	InventorySourceObj := model.InventorySource                                                                                            {Name:"test value for Name",Domain:"test value for Domain",Channel:0,PrimaryFormat:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInventorySourceRequestResult := dao.CreateInventorySource( InventorySourceObj )
	
	if createInventorySourceRequestResult.Success == false {
		t.Errorf(createInventorySourceRequestResult.Msg)
	} else {
		fmt.Println("Check Create InventorySource success...")
	}
	
	createInventorySourceObj,_ := createInventorySourceRequestResult.Data. (model.InventorySource)

	// --------------------------------------------------------------
	// Check InventorySource Obj ID
	// --------------------------------------------------------------	
	if createInventorySourceObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InventorySource" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInventorySourceRequestResult := dao.GetInventorySource( uint64(createInventorySourceObj.ID) )
	
	if getInventorySourceRequestResult.Success == false {
		t.Errorf(getInventorySourceRequestResult.Msg)
	} else {
		fmt.Println("Check Get InventorySource success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInventorySourceObj,_ := getInventorySourceRequestResult.Data. (model.InventorySource)
	compareInventorySource := cmp.Equal(createInventorySourceObj.ID, getInventorySourceObj.ID)
	
	if  compareInventorySource == false	{
		t.Errorf( "Created InventorySource object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInventorySourceRequestResult := dao.GetAllInventorySource()

	if getAllInventorySourceRequestResult.Success == false {
			t.Errorf(getAllInventorySourceRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InventorySource success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInventorySourceObj []model.InventorySource = getAllInventorySourceRequestResult.Data. ([]model.InventorySource)
		
	equalInventorySource := cmp.Equal(createInventorySourceObj.ID, getAllInventorySourceObj[len(getAllInventorySourceObj)-1].ID)
		
	if equalInventorySource == false {
		t.Errorf( "Created object is not equal to the last entry in InventorySource[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InventorySource
	// --------------------------------------------------------------	
	deleteInventorySourceRequestResult := dao.DeleteInventorySource(uint64(createInventorySourceObj.ID))

	if deleteInventorySourceRequestResult.Success == false {
			t.Errorf(deleteInventorySourceRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InventorySource success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInventorySourceRequestResult = dao.GetInventorySource( uint64(createInventorySourceObj.ID) )
	
	if getInventorySourceRequestResult.Success == true {
		t.Errorf(getInventorySourceRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAdSlotCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AdSlot
	//----------------------------------------------------------------------------
	AdSlotObj := model.AdSlot                                                                                                                            {SlotCode:"test value for SlotCode",Width:100,Height:100,FloorPrice:new Money(),Format:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAdSlotRequestResult := dao.CreateAdSlot( AdSlotObj )
	
	if createAdSlotRequestResult.Success == false {
		t.Errorf(createAdSlotRequestResult.Msg)
	} else {
		fmt.Println("Check Create AdSlot success...")
	}
	
	createAdSlotObj,_ := createAdSlotRequestResult.Data. (model.AdSlot)

	// --------------------------------------------------------------
	// Check AdSlot Obj ID
	// --------------------------------------------------------------	
	if createAdSlotObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AdSlot" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAdSlotRequestResult := dao.GetAdSlot( uint64(createAdSlotObj.ID) )
	
	if getAdSlotRequestResult.Success == false {
		t.Errorf(getAdSlotRequestResult.Msg)
	} else {
		fmt.Println("Check Get AdSlot success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAdSlotObj,_ := getAdSlotRequestResult.Data. (model.AdSlot)
	compareAdSlot := cmp.Equal(createAdSlotObj.ID, getAdSlotObj.ID)
	
	if  compareAdSlot == false	{
		t.Errorf( "Created AdSlot object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAdSlotRequestResult := dao.GetAllAdSlot()

	if getAllAdSlotRequestResult.Success == false {
			t.Errorf(getAllAdSlotRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AdSlot success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAdSlotObj []model.AdSlot = getAllAdSlotRequestResult.Data. ([]model.AdSlot)
		
	equalAdSlot := cmp.Equal(createAdSlotObj.ID, getAllAdSlotObj[len(getAllAdSlotObj)-1].ID)
		
	if equalAdSlot == false {
		t.Errorf( "Created object is not equal to the last entry in AdSlot[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AdSlot
	// --------------------------------------------------------------	
	deleteAdSlotRequestResult := dao.DeleteAdSlot(uint64(createAdSlotObj.ID))

	if deleteAdSlotRequestResult.Success == false {
			t.Errorf(deleteAdSlotRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AdSlot success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAdSlotRequestResult = dao.GetAdSlot( uint64(createAdSlotObj.ID) )
	
	if getAdSlotRequestResult.Success == true {
		t.Errorf(getAdSlotRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDealCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Deal
	//----------------------------------------------------------------------------
	DealObj := model.Deal                            {FloorPrice:new Money(),DealType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDealRequestResult := dao.CreateDeal( DealObj )
	
	if createDealRequestResult.Success == false {
		t.Errorf(createDealRequestResult.Msg)
	} else {
		fmt.Println("Check Create Deal success...")
	}
	
	createDealObj,_ := createDealRequestResult.Data. (model.Deal)

	// --------------------------------------------------------------
	// Check Deal Obj ID
	// --------------------------------------------------------------	
	if createDealObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Deal" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDealRequestResult := dao.GetDeal( uint64(createDealObj.ID) )
	
	if getDealRequestResult.Success == false {
		t.Errorf(getDealRequestResult.Msg)
	} else {
		fmt.Println("Check Get Deal success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDealObj,_ := getDealRequestResult.Data. (model.Deal)
	compareDeal := cmp.Equal(createDealObj.ID, getDealObj.ID)
	
	if  compareDeal == false	{
		t.Errorf( "Created Deal object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDealRequestResult := dao.GetAllDeal()

	if getAllDealRequestResult.Success == false {
			t.Errorf(getAllDealRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Deal success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDealObj []model.Deal = getAllDealRequestResult.Data. ([]model.Deal)
		
	equalDeal := cmp.Equal(createDealObj.ID, getAllDealObj[len(getAllDealObj)-1].ID)
		
	if equalDeal == false {
		t.Errorf( "Created object is not equal to the last entry in Deal[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Deal
	// --------------------------------------------------------------	
	deleteDealRequestResult := dao.DeleteDeal(uint64(createDealObj.ID))

	if deleteDealRequestResult.Success == false {
			t.Errorf(deleteDealRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Deal success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDealRequestResult = dao.GetDeal( uint64(createDealObj.ID) )
	
	if getDealRequestResult.Success == true {
		t.Errorf(getDealRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPlacementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Placement
	//----------------------------------------------------------------------------
	PlacementObj := model.Placement                                                                            {Name:"test value for Name",Flight:new DateRange(),GoalImpressions:100}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPlacementRequestResult := dao.CreatePlacement( PlacementObj )
	
	if createPlacementRequestResult.Success == false {
		t.Errorf(createPlacementRequestResult.Msg)
	} else {
		fmt.Println("Check Create Placement success...")
	}
	
	createPlacementObj,_ := createPlacementRequestResult.Data. (model.Placement)

	// --------------------------------------------------------------
	// Check Placement Obj ID
	// --------------------------------------------------------------	
	if createPlacementObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Placement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPlacementRequestResult := dao.GetPlacement( uint64(createPlacementObj.ID) )
	
	if getPlacementRequestResult.Success == false {
		t.Errorf(getPlacementRequestResult.Msg)
	} else {
		fmt.Println("Check Get Placement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPlacementObj,_ := getPlacementRequestResult.Data. (model.Placement)
	comparePlacement := cmp.Equal(createPlacementObj.ID, getPlacementObj.ID)
	
	if  comparePlacement == false	{
		t.Errorf( "Created Placement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPlacementRequestResult := dao.GetAllPlacement()

	if getAllPlacementRequestResult.Success == false {
			t.Errorf(getAllPlacementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Placement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPlacementObj []model.Placement = getAllPlacementRequestResult.Data. ([]model.Placement)
		
	equalPlacement := cmp.Equal(createPlacementObj.ID, getAllPlacementObj[len(getAllPlacementObj)-1].ID)
		
	if equalPlacement == false {
		t.Errorf( "Created object is not equal to the last entry in Placement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Placement
	// --------------------------------------------------------------	
	deletePlacementRequestResult := dao.DeletePlacement(uint64(createPlacementObj.ID))

	if deletePlacementRequestResult.Success == false {
			t.Errorf(deletePlacementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Placement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPlacementRequestResult = dao.GetPlacement( uint64(createPlacementObj.ID) )
	
	if getPlacementRequestResult.Success == true {
		t.Errorf(getPlacementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCreativeAssetCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CreativeAsset
	//----------------------------------------------------------------------------
	CreativeAssetObj := model.CreativeAsset                                                                                                                                                                                            {Name:"test value for Name",ClickUrl:new URL(),LandingPage:new URL(),Width:100,Height:100,DurationSeconds:100,CreativeType:0,AdFormat:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCreativeAssetRequestResult := dao.CreateCreativeAsset( CreativeAssetObj )
	
	if createCreativeAssetRequestResult.Success == false {
		t.Errorf(createCreativeAssetRequestResult.Msg)
	} else {
		fmt.Println("Check Create CreativeAsset success...")
	}
	
	createCreativeAssetObj,_ := createCreativeAssetRequestResult.Data. (model.CreativeAsset)

	// --------------------------------------------------------------
	// Check CreativeAsset Obj ID
	// --------------------------------------------------------------	
	if createCreativeAssetObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CreativeAsset" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCreativeAssetRequestResult := dao.GetCreativeAsset( uint64(createCreativeAssetObj.ID) )
	
	if getCreativeAssetRequestResult.Success == false {
		t.Errorf(getCreativeAssetRequestResult.Msg)
	} else {
		fmt.Println("Check Get CreativeAsset success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCreativeAssetObj,_ := getCreativeAssetRequestResult.Data. (model.CreativeAsset)
	compareCreativeAsset := cmp.Equal(createCreativeAssetObj.ID, getCreativeAssetObj.ID)
	
	if  compareCreativeAsset == false	{
		t.Errorf( "Created CreativeAsset object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCreativeAssetRequestResult := dao.GetAllCreativeAsset()

	if getAllCreativeAssetRequestResult.Success == false {
			t.Errorf(getAllCreativeAssetRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CreativeAsset success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCreativeAssetObj []model.CreativeAsset = getAllCreativeAssetRequestResult.Data. ([]model.CreativeAsset)
		
	equalCreativeAsset := cmp.Equal(createCreativeAssetObj.ID, getAllCreativeAssetObj[len(getAllCreativeAssetObj)-1].ID)
		
	if equalCreativeAsset == false {
		t.Errorf( "Created object is not equal to the last entry in CreativeAsset[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CreativeAsset
	// --------------------------------------------------------------	
	deleteCreativeAssetRequestResult := dao.DeleteCreativeAsset(uint64(createCreativeAssetObj.ID))

	if deleteCreativeAssetRequestResult.Success == false {
			t.Errorf(deleteCreativeAssetRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CreativeAsset success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCreativeAssetRequestResult = dao.GetCreativeAsset( uint64(createCreativeAssetObj.ID) )
	
	if getCreativeAssetRequestResult.Success == true {
		t.Errorf(getCreativeAssetRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCreativeFileCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CreativeFile
	//----------------------------------------------------------------------------
	CreativeFileObj := model.CreativeFile                                                                                                            {Uri:new URL(),FileSizeKB:100,MimeType:"test value for MimeType",Checksum:"test value for Checksum"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCreativeFileRequestResult := dao.CreateCreativeFile( CreativeFileObj )
	
	if createCreativeFileRequestResult.Success == false {
		t.Errorf(createCreativeFileRequestResult.Msg)
	} else {
		fmt.Println("Check Create CreativeFile success...")
	}
	
	createCreativeFileObj,_ := createCreativeFileRequestResult.Data. (model.CreativeFile)

	// --------------------------------------------------------------
	// Check CreativeFile Obj ID
	// --------------------------------------------------------------	
	if createCreativeFileObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CreativeFile" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCreativeFileRequestResult := dao.GetCreativeFile( uint64(createCreativeFileObj.ID) )
	
	if getCreativeFileRequestResult.Success == false {
		t.Errorf(getCreativeFileRequestResult.Msg)
	} else {
		fmt.Println("Check Get CreativeFile success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCreativeFileObj,_ := getCreativeFileRequestResult.Data. (model.CreativeFile)
	compareCreativeFile := cmp.Equal(createCreativeFileObj.ID, getCreativeFileObj.ID)
	
	if  compareCreativeFile == false	{
		t.Errorf( "Created CreativeFile object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCreativeFileRequestResult := dao.GetAllCreativeFile()

	if getAllCreativeFileRequestResult.Success == false {
			t.Errorf(getAllCreativeFileRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CreativeFile success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCreativeFileObj []model.CreativeFile = getAllCreativeFileRequestResult.Data. ([]model.CreativeFile)
		
	equalCreativeFile := cmp.Equal(createCreativeFileObj.ID, getAllCreativeFileObj[len(getAllCreativeFileObj)-1].ID)
		
	if equalCreativeFile == false {
		t.Errorf( "Created object is not equal to the last entry in CreativeFile[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CreativeFile
	// --------------------------------------------------------------	
	deleteCreativeFileRequestResult := dao.DeleteCreativeFile(uint64(createCreativeFileObj.ID))

	if deleteCreativeFileRequestResult.Success == false {
			t.Errorf(deleteCreativeFileRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CreativeFile success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCreativeFileRequestResult = dao.GetCreativeFile( uint64(createCreativeFileObj.ID) )
	
	if getCreativeFileRequestResult.Success == true {
		t.Errorf(getCreativeFileRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCreativeVariationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CreativeVariation
	//----------------------------------------------------------------------------
	CreativeVariationObj := model.CreativeVariation                                                                                                                                                            {Name:"test value for Name",Language:"test value for Language",Headline:"test value for Headline",BodyText:"test value for BodyText",CallToAction:"test value for CallToAction"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCreativeVariationRequestResult := dao.CreateCreativeVariation( CreativeVariationObj )
	
	if createCreativeVariationRequestResult.Success == false {
		t.Errorf(createCreativeVariationRequestResult.Msg)
	} else {
		fmt.Println("Check Create CreativeVariation success...")
	}
	
	createCreativeVariationObj,_ := createCreativeVariationRequestResult.Data. (model.CreativeVariation)

	// --------------------------------------------------------------
	// Check CreativeVariation Obj ID
	// --------------------------------------------------------------	
	if createCreativeVariationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CreativeVariation" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCreativeVariationRequestResult := dao.GetCreativeVariation( uint64(createCreativeVariationObj.ID) )
	
	if getCreativeVariationRequestResult.Success == false {
		t.Errorf(getCreativeVariationRequestResult.Msg)
	} else {
		fmt.Println("Check Get CreativeVariation success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCreativeVariationObj,_ := getCreativeVariationRequestResult.Data. (model.CreativeVariation)
	compareCreativeVariation := cmp.Equal(createCreativeVariationObj.ID, getCreativeVariationObj.ID)
	
	if  compareCreativeVariation == false	{
		t.Errorf( "Created CreativeVariation object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCreativeVariationRequestResult := dao.GetAllCreativeVariation()

	if getAllCreativeVariationRequestResult.Success == false {
			t.Errorf(getAllCreativeVariationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CreativeVariation success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCreativeVariationObj []model.CreativeVariation = getAllCreativeVariationRequestResult.Data. ([]model.CreativeVariation)
		
	equalCreativeVariation := cmp.Equal(createCreativeVariationObj.ID, getAllCreativeVariationObj[len(getAllCreativeVariationObj)-1].ID)
		
	if equalCreativeVariation == false {
		t.Errorf( "Created object is not equal to the last entry in CreativeVariation[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CreativeVariation
	// --------------------------------------------------------------	
	deleteCreativeVariationRequestResult := dao.DeleteCreativeVariation(uint64(createCreativeVariationObj.ID))

	if deleteCreativeVariationRequestResult.Success == false {
			t.Errorf(deleteCreativeVariationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CreativeVariation success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCreativeVariationRequestResult = dao.GetCreativeVariation( uint64(createCreativeVariationObj.ID) )
	
	if getCreativeVariationRequestResult.Success == true {
		t.Errorf(getCreativeVariationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCreativeApprovalCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CreativeApproval
	//----------------------------------------------------------------------------
	CreativeApprovalObj := model.CreativeApproval                                                                                                    {Reviewer:"test value for Reviewer",ReviewedAt:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCreativeApprovalRequestResult := dao.CreateCreativeApproval( CreativeApprovalObj )
	
	if createCreativeApprovalRequestResult.Success == false {
		t.Errorf(createCreativeApprovalRequestResult.Msg)
	} else {
		fmt.Println("Check Create CreativeApproval success...")
	}
	
	createCreativeApprovalObj,_ := createCreativeApprovalRequestResult.Data. (model.CreativeApproval)

	// --------------------------------------------------------------
	// Check CreativeApproval Obj ID
	// --------------------------------------------------------------	
	if createCreativeApprovalObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CreativeApproval" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCreativeApprovalRequestResult := dao.GetCreativeApproval( uint64(createCreativeApprovalObj.ID) )
	
	if getCreativeApprovalRequestResult.Success == false {
		t.Errorf(getCreativeApprovalRequestResult.Msg)
	} else {
		fmt.Println("Check Get CreativeApproval success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCreativeApprovalObj,_ := getCreativeApprovalRequestResult.Data. (model.CreativeApproval)
	compareCreativeApproval := cmp.Equal(createCreativeApprovalObj.ID, getCreativeApprovalObj.ID)
	
	if  compareCreativeApproval == false	{
		t.Errorf( "Created CreativeApproval object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCreativeApprovalRequestResult := dao.GetAllCreativeApproval()

	if getAllCreativeApprovalRequestResult.Success == false {
			t.Errorf(getAllCreativeApprovalRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CreativeApproval success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCreativeApprovalObj []model.CreativeApproval = getAllCreativeApprovalRequestResult.Data. ([]model.CreativeApproval)
		
	equalCreativeApproval := cmp.Equal(createCreativeApprovalObj.ID, getAllCreativeApprovalObj[len(getAllCreativeApprovalObj)-1].ID)
		
	if equalCreativeApproval == false {
		t.Errorf( "Created object is not equal to the last entry in CreativeApproval[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CreativeApproval
	// --------------------------------------------------------------	
	deleteCreativeApprovalRequestResult := dao.DeleteCreativeApproval(uint64(createCreativeApprovalObj.ID))

	if deleteCreativeApprovalRequestResult.Success == false {
			t.Errorf(deleteCreativeApprovalRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CreativeApproval success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCreativeApprovalRequestResult = dao.GetCreativeApproval( uint64(createCreativeApprovalObj.ID) )
	
	if getCreativeApprovalRequestResult.Success == true {
		t.Errorf(getCreativeApprovalRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTrackingPixelCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TrackingPixel
	//----------------------------------------------------------------------------
	TrackingPixelObj := model.TrackingPixel                                                                            {Name:"test value for Name",Url:new URL(),EventType:0,PixelType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTrackingPixelRequestResult := dao.CreateTrackingPixel( TrackingPixelObj )
	
	if createTrackingPixelRequestResult.Success == false {
		t.Errorf(createTrackingPixelRequestResult.Msg)
	} else {
		fmt.Println("Check Create TrackingPixel success...")
	}
	
	createTrackingPixelObj,_ := createTrackingPixelRequestResult.Data. (model.TrackingPixel)

	// --------------------------------------------------------------
	// Check TrackingPixel Obj ID
	// --------------------------------------------------------------	
	if createTrackingPixelObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TrackingPixel" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTrackingPixelRequestResult := dao.GetTrackingPixel( uint64(createTrackingPixelObj.ID) )
	
	if getTrackingPixelRequestResult.Success == false {
		t.Errorf(getTrackingPixelRequestResult.Msg)
	} else {
		fmt.Println("Check Get TrackingPixel success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTrackingPixelObj,_ := getTrackingPixelRequestResult.Data. (model.TrackingPixel)
	compareTrackingPixel := cmp.Equal(createTrackingPixelObj.ID, getTrackingPixelObj.ID)
	
	if  compareTrackingPixel == false	{
		t.Errorf( "Created TrackingPixel object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTrackingPixelRequestResult := dao.GetAllTrackingPixel()

	if getAllTrackingPixelRequestResult.Success == false {
			t.Errorf(getAllTrackingPixelRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TrackingPixel success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTrackingPixelObj []model.TrackingPixel = getAllTrackingPixelRequestResult.Data. ([]model.TrackingPixel)
		
	equalTrackingPixel := cmp.Equal(createTrackingPixelObj.ID, getAllTrackingPixelObj[len(getAllTrackingPixelObj)-1].ID)
		
	if equalTrackingPixel == false {
		t.Errorf( "Created object is not equal to the last entry in TrackingPixel[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TrackingPixel
	// --------------------------------------------------------------	
	deleteTrackingPixelRequestResult := dao.DeleteTrackingPixel(uint64(createTrackingPixelObj.ID))

	if deleteTrackingPixelRequestResult.Success == false {
			t.Errorf(deleteTrackingPixelRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TrackingPixel success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTrackingPixelRequestResult = dao.GetTrackingPixel( uint64(createTrackingPixelObj.ID) )
	
	if getTrackingPixelRequestResult.Success == true {
		t.Errorf(getTrackingPixelRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestConversionEventCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ConversionEvent
	//----------------------------------------------------------------------------
	ConversionEventObj := model.ConversionEvent                                                                                                    {Timestamp:time.Now(),Value:new Money(),EventType:0,AttributionModel:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createConversionEventRequestResult := dao.CreateConversionEvent( ConversionEventObj )
	
	if createConversionEventRequestResult.Success == false {
		t.Errorf(createConversionEventRequestResult.Msg)
	} else {
		fmt.Println("Check Create ConversionEvent success...")
	}
	
	createConversionEventObj,_ := createConversionEventRequestResult.Data. (model.ConversionEvent)

	// --------------------------------------------------------------
	// Check ConversionEvent Obj ID
	// --------------------------------------------------------------	
	if createConversionEventObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ConversionEvent" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getConversionEventRequestResult := dao.GetConversionEvent( uint64(createConversionEventObj.ID) )
	
	if getConversionEventRequestResult.Success == false {
		t.Errorf(getConversionEventRequestResult.Msg)
	} else {
		fmt.Println("Check Get ConversionEvent success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getConversionEventObj,_ := getConversionEventRequestResult.Data. (model.ConversionEvent)
	compareConversionEvent := cmp.Equal(createConversionEventObj.ID, getConversionEventObj.ID)
	
	if  compareConversionEvent == false	{
		t.Errorf( "Created ConversionEvent object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllConversionEventRequestResult := dao.GetAllConversionEvent()

	if getAllConversionEventRequestResult.Success == false {
			t.Errorf(getAllConversionEventRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ConversionEvent success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllConversionEventObj []model.ConversionEvent = getAllConversionEventRequestResult.Data. ([]model.ConversionEvent)
		
	equalConversionEvent := cmp.Equal(createConversionEventObj.ID, getAllConversionEventObj[len(getAllConversionEventObj)-1].ID)
		
	if equalConversionEvent == false {
		t.Errorf( "Created object is not equal to the last entry in ConversionEvent[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ConversionEvent
	// --------------------------------------------------------------	
	deleteConversionEventRequestResult := dao.DeleteConversionEvent(uint64(createConversionEventObj.ID))

	if deleteConversionEventRequestResult.Success == false {
			t.Errorf(deleteConversionEventRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ConversionEvent success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getConversionEventRequestResult = dao.GetConversionEvent( uint64(createConversionEventObj.ID) )
	
	if getConversionEventRequestResult.Success == true {
		t.Errorf(getConversionEventRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPerformanceMetricCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PerformanceMetric
	//----------------------------------------------------------------------------
	PerformanceMetricObj := model.PerformanceMetric                                                                                                                            {Date:time.Now(),Value:"test value",MetricType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPerformanceMetricRequestResult := dao.CreatePerformanceMetric( PerformanceMetricObj )
	
	if createPerformanceMetricRequestResult.Success == false {
		t.Errorf(createPerformanceMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Create PerformanceMetric success...")
	}
	
	createPerformanceMetricObj,_ := createPerformanceMetricRequestResult.Data. (model.PerformanceMetric)

	// --------------------------------------------------------------
	// Check PerformanceMetric Obj ID
	// --------------------------------------------------------------	
	if createPerformanceMetricObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PerformanceMetric" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPerformanceMetricRequestResult := dao.GetPerformanceMetric( uint64(createPerformanceMetricObj.ID) )
	
	if getPerformanceMetricRequestResult.Success == false {
		t.Errorf(getPerformanceMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Get PerformanceMetric success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPerformanceMetricObj,_ := getPerformanceMetricRequestResult.Data. (model.PerformanceMetric)
	comparePerformanceMetric := cmp.Equal(createPerformanceMetricObj.ID, getPerformanceMetricObj.ID)
	
	if  comparePerformanceMetric == false	{
		t.Errorf( "Created PerformanceMetric object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPerformanceMetricRequestResult := dao.GetAllPerformanceMetric()

	if getAllPerformanceMetricRequestResult.Success == false {
			t.Errorf(getAllPerformanceMetricRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PerformanceMetric success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPerformanceMetricObj []model.PerformanceMetric = getAllPerformanceMetricRequestResult.Data. ([]model.PerformanceMetric)
		
	equalPerformanceMetric := cmp.Equal(createPerformanceMetricObj.ID, getAllPerformanceMetricObj[len(getAllPerformanceMetricObj)-1].ID)
		
	if equalPerformanceMetric == false {
		t.Errorf( "Created object is not equal to the last entry in PerformanceMetric[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PerformanceMetric
	// --------------------------------------------------------------	
	deletePerformanceMetricRequestResult := dao.DeletePerformanceMetric(uint64(createPerformanceMetricObj.ID))

	if deletePerformanceMetricRequestResult.Success == false {
			t.Errorf(deletePerformanceMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PerformanceMetric success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPerformanceMetricRequestResult = dao.GetPerformanceMetric( uint64(createPerformanceMetricObj.ID) )
	
	if getPerformanceMetricRequestResult.Success == true {
		t.Errorf(getPerformanceMetricRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestReportCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Report
	//----------------------------------------------------------------------------
	ReportObj := model.Report                                                                                                                    {ReportName:"test value for ReportName",GeneratedAt:time.Now(),FileUrl:new URL(),ReportType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createReportRequestResult := dao.CreateReport( ReportObj )
	
	if createReportRequestResult.Success == false {
		t.Errorf(createReportRequestResult.Msg)
	} else {
		fmt.Println("Check Create Report success...")
	}
	
	createReportObj,_ := createReportRequestResult.Data. (model.Report)

	// --------------------------------------------------------------
	// Check Report Obj ID
	// --------------------------------------------------------------	
	if createReportObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Report" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getReportRequestResult := dao.GetReport( uint64(createReportObj.ID) )
	
	if getReportRequestResult.Success == false {
		t.Errorf(getReportRequestResult.Msg)
	} else {
		fmt.Println("Check Get Report success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getReportObj,_ := getReportRequestResult.Data. (model.Report)
	compareReport := cmp.Equal(createReportObj.ID, getReportObj.ID)
	
	if  compareReport == false	{
		t.Errorf( "Created Report object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllReportRequestResult := dao.GetAllReport()

	if getAllReportRequestResult.Success == false {
			t.Errorf(getAllReportRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Report success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllReportObj []model.Report = getAllReportRequestResult.Data. ([]model.Report)
		
	equalReport := cmp.Equal(createReportObj.ID, getAllReportObj[len(getAllReportObj)-1].ID)
		
	if equalReport == false {
		t.Errorf( "Created object is not equal to the last entry in Report[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Report
	// --------------------------------------------------------------	
	deleteReportRequestResult := dao.DeleteReport(uint64(createReportObj.ID))

	if deleteReportRequestResult.Success == false {
			t.Errorf(deleteReportRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Report success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getReportRequestResult = dao.GetReport( uint64(createReportObj.ID) )
	
	if getReportRequestResult.Success == true {
		t.Errorf(getReportRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInsertionOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InsertionOrder
	//----------------------------------------------------------------------------
	InsertionOrderObj := model.InsertionOrder                                                                            {IoNumber:"test value for IoNumber",AgreedBudget:new Money(),Flight:new DateRange(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInsertionOrderRequestResult := dao.CreateInsertionOrder( InsertionOrderObj )
	
	if createInsertionOrderRequestResult.Success == false {
		t.Errorf(createInsertionOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create InsertionOrder success...")
	}
	
	createInsertionOrderObj,_ := createInsertionOrderRequestResult.Data. (model.InsertionOrder)

	// --------------------------------------------------------------
	// Check InsertionOrder Obj ID
	// --------------------------------------------------------------	
	if createInsertionOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InsertionOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInsertionOrderRequestResult := dao.GetInsertionOrder( uint64(createInsertionOrderObj.ID) )
	
	if getInsertionOrderRequestResult.Success == false {
		t.Errorf(getInsertionOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get InsertionOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInsertionOrderObj,_ := getInsertionOrderRequestResult.Data. (model.InsertionOrder)
	compareInsertionOrder := cmp.Equal(createInsertionOrderObj.ID, getInsertionOrderObj.ID)
	
	if  compareInsertionOrder == false	{
		t.Errorf( "Created InsertionOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInsertionOrderRequestResult := dao.GetAllInsertionOrder()

	if getAllInsertionOrderRequestResult.Success == false {
			t.Errorf(getAllInsertionOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InsertionOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInsertionOrderObj []model.InsertionOrder = getAllInsertionOrderRequestResult.Data. ([]model.InsertionOrder)
		
	equalInsertionOrder := cmp.Equal(createInsertionOrderObj.ID, getAllInsertionOrderObj[len(getAllInsertionOrderObj)-1].ID)
		
	if equalInsertionOrder == false {
		t.Errorf( "Created object is not equal to the last entry in InsertionOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InsertionOrder
	// --------------------------------------------------------------	
	deleteInsertionOrderRequestResult := dao.DeleteInsertionOrder(uint64(createInsertionOrderObj.ID))

	if deleteInsertionOrderRequestResult.Success == false {
			t.Errorf(deleteInsertionOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InsertionOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInsertionOrderRequestResult = dao.GetInsertionOrder( uint64(createInsertionOrderObj.ID) )
	
	if getInsertionOrderRequestResult.Success == true {
		t.Errorf(getInsertionOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRateCardCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RateCard
	//----------------------------------------------------------------------------
	RateCardObj := model.RateCard                                                                                                                    {Name:"test value for Name",EffectiveDate:time.Now(),Currency:"test value for Currency"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRateCardRequestResult := dao.CreateRateCard( RateCardObj )
	
	if createRateCardRequestResult.Success == false {
		t.Errorf(createRateCardRequestResult.Msg)
	} else {
		fmt.Println("Check Create RateCard success...")
	}
	
	createRateCardObj,_ := createRateCardRequestResult.Data. (model.RateCard)

	// --------------------------------------------------------------
	// Check RateCard Obj ID
	// --------------------------------------------------------------	
	if createRateCardObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for RateCard" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRateCardRequestResult := dao.GetRateCard( uint64(createRateCardObj.ID) )
	
	if getRateCardRequestResult.Success == false {
		t.Errorf(getRateCardRequestResult.Msg)
	} else {
		fmt.Println("Check Get RateCard success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRateCardObj,_ := getRateCardRequestResult.Data. (model.RateCard)
	compareRateCard := cmp.Equal(createRateCardObj.ID, getRateCardObj.ID)
	
	if  compareRateCard == false	{
		t.Errorf( "Created RateCard object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRateCardRequestResult := dao.GetAllRateCard()

	if getAllRateCardRequestResult.Success == false {
			t.Errorf(getAllRateCardRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RateCard success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRateCardObj []model.RateCard = getAllRateCardRequestResult.Data. ([]model.RateCard)
		
	equalRateCard := cmp.Equal(createRateCardObj.ID, getAllRateCardObj[len(getAllRateCardObj)-1].ID)
		
	if equalRateCard == false {
		t.Errorf( "Created object is not equal to the last entry in RateCard[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RateCard
	// --------------------------------------------------------------	
	deleteRateCardRequestResult := dao.DeleteRateCard(uint64(createRateCardObj.ID))

	if deleteRateCardRequestResult.Success == false {
			t.Errorf(deleteRateCardRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RateCard success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRateCardRequestResult = dao.GetRateCard( uint64(createRateCardObj.ID) )
	
	if getRateCardRequestResult.Success == true {
		t.Errorf(getRateCardRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRateCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Rate
	//----------------------------------------------------------------------------
	RateObj := model.Rate                                            {UnitPrice:new Money(),AdFormat:0,PricingModel:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRateRequestResult := dao.CreateRate( RateObj )
	
	if createRateRequestResult.Success == false {
		t.Errorf(createRateRequestResult.Msg)
	} else {
		fmt.Println("Check Create Rate success...")
	}
	
	createRateObj,_ := createRateRequestResult.Data. (model.Rate)

	// --------------------------------------------------------------
	// Check Rate Obj ID
	// --------------------------------------------------------------	
	if createRateObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Rate" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRateRequestResult := dao.GetRate( uint64(createRateObj.ID) )
	
	if getRateRequestResult.Success == false {
		t.Errorf(getRateRequestResult.Msg)
	} else {
		fmt.Println("Check Get Rate success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRateObj,_ := getRateRequestResult.Data. (model.Rate)
	compareRate := cmp.Equal(createRateObj.ID, getRateObj.ID)
	
	if  compareRate == false	{
		t.Errorf( "Created Rate object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRateRequestResult := dao.GetAllRate()

	if getAllRateRequestResult.Success == false {
			t.Errorf(getAllRateRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Rate success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRateObj []model.Rate = getAllRateRequestResult.Data. ([]model.Rate)
		
	equalRate := cmp.Equal(createRateObj.ID, getAllRateObj[len(getAllRateObj)-1].ID)
		
	if equalRate == false {
		t.Errorf( "Created object is not equal to the last entry in Rate[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Rate
	// --------------------------------------------------------------	
	deleteRateRequestResult := dao.DeleteRate(uint64(createRateObj.ID))

	if deleteRateRequestResult.Success == false {
			t.Errorf(deleteRateRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Rate success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRateRequestResult = dao.GetRate( uint64(createRateObj.ID) )
	
	if getRateRequestResult.Success == true {
		t.Errorf(getRateRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestExperimentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Experiment
	//----------------------------------------------------------------------------
	ExperimentObj := model.Experiment                                                                                                                                                                                            {Name:"test value for Name",Hypothesis:"test value for Hypothesis",StartDate:time.Now(),EndDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createExperimentRequestResult := dao.CreateExperiment( ExperimentObj )
	
	if createExperimentRequestResult.Success == false {
		t.Errorf(createExperimentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Experiment success...")
	}
	
	createExperimentObj,_ := createExperimentRequestResult.Data. (model.Experiment)

	// --------------------------------------------------------------
	// Check Experiment Obj ID
	// --------------------------------------------------------------	
	if createExperimentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Experiment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getExperimentRequestResult := dao.GetExperiment( uint64(createExperimentObj.ID) )
	
	if getExperimentRequestResult.Success == false {
		t.Errorf(getExperimentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Experiment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getExperimentObj,_ := getExperimentRequestResult.Data. (model.Experiment)
	compareExperiment := cmp.Equal(createExperimentObj.ID, getExperimentObj.ID)
	
	if  compareExperiment == false	{
		t.Errorf( "Created Experiment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllExperimentRequestResult := dao.GetAllExperiment()

	if getAllExperimentRequestResult.Success == false {
			t.Errorf(getAllExperimentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Experiment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllExperimentObj []model.Experiment = getAllExperimentRequestResult.Data. ([]model.Experiment)
		
	equalExperiment := cmp.Equal(createExperimentObj.ID, getAllExperimentObj[len(getAllExperimentObj)-1].ID)
		
	if equalExperiment == false {
		t.Errorf( "Created object is not equal to the last entry in Experiment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Experiment
	// --------------------------------------------------------------	
	deleteExperimentRequestResult := dao.DeleteExperiment(uint64(createExperimentObj.ID))

	if deleteExperimentRequestResult.Success == false {
			t.Errorf(deleteExperimentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Experiment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getExperimentRequestResult = dao.GetExperiment( uint64(createExperimentObj.ID) )
	
	if getExperimentRequestResult.Success == true {
		t.Errorf(getExperimentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestExperimentVariantCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ExperimentVariant
	//----------------------------------------------------------------------------
	ExperimentVariantObj := model.ExperimentVariant                                            {Name:"test value for Name",Allocation:new Percentage()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createExperimentVariantRequestResult := dao.CreateExperimentVariant( ExperimentVariantObj )
	
	if createExperimentVariantRequestResult.Success == false {
		t.Errorf(createExperimentVariantRequestResult.Msg)
	} else {
		fmt.Println("Check Create ExperimentVariant success...")
	}
	
	createExperimentVariantObj,_ := createExperimentVariantRequestResult.Data. (model.ExperimentVariant)

	// --------------------------------------------------------------
	// Check ExperimentVariant Obj ID
	// --------------------------------------------------------------	
	if createExperimentVariantObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ExperimentVariant" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getExperimentVariantRequestResult := dao.GetExperimentVariant( uint64(createExperimentVariantObj.ID) )
	
	if getExperimentVariantRequestResult.Success == false {
		t.Errorf(getExperimentVariantRequestResult.Msg)
	} else {
		fmt.Println("Check Get ExperimentVariant success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getExperimentVariantObj,_ := getExperimentVariantRequestResult.Data. (model.ExperimentVariant)
	compareExperimentVariant := cmp.Equal(createExperimentVariantObj.ID, getExperimentVariantObj.ID)
	
	if  compareExperimentVariant == false	{
		t.Errorf( "Created ExperimentVariant object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllExperimentVariantRequestResult := dao.GetAllExperimentVariant()

	if getAllExperimentVariantRequestResult.Success == false {
			t.Errorf(getAllExperimentVariantRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ExperimentVariant success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllExperimentVariantObj []model.ExperimentVariant = getAllExperimentVariantRequestResult.Data. ([]model.ExperimentVariant)
		
	equalExperimentVariant := cmp.Equal(createExperimentVariantObj.ID, getAllExperimentVariantObj[len(getAllExperimentVariantObj)-1].ID)
		
	if equalExperimentVariant == false {
		t.Errorf( "Created object is not equal to the last entry in ExperimentVariant[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ExperimentVariant
	// --------------------------------------------------------------	
	deleteExperimentVariantRequestResult := dao.DeleteExperimentVariant(uint64(createExperimentVariantObj.ID))

	if deleteExperimentVariantRequestResult.Success == false {
			t.Errorf(deleteExperimentVariantRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ExperimentVariant success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getExperimentVariantRequestResult = dao.GetExperimentVariant( uint64(createExperimentVariantObj.ID) )
	
	if getExperimentVariantRequestResult.Success == true {
		t.Errorf(getExperimentVariantRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestGeoRegionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for GeoRegion
	//----------------------------------------------------------------------------
	GeoRegionObj := model.GeoRegion                                                                            {Code:"test value for Code",Name:"test value for Name",RegionType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createGeoRegionRequestResult := dao.CreateGeoRegion( GeoRegionObj )
	
	if createGeoRegionRequestResult.Success == false {
		t.Errorf(createGeoRegionRequestResult.Msg)
	} else {
		fmt.Println("Check Create GeoRegion success...")
	}
	
	createGeoRegionObj,_ := createGeoRegionRequestResult.Data. (model.GeoRegion)

	// --------------------------------------------------------------
	// Check GeoRegion Obj ID
	// --------------------------------------------------------------	
	if createGeoRegionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for GeoRegion" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getGeoRegionRequestResult := dao.GetGeoRegion( uint64(createGeoRegionObj.ID) )
	
	if getGeoRegionRequestResult.Success == false {
		t.Errorf(getGeoRegionRequestResult.Msg)
	} else {
		fmt.Println("Check Get GeoRegion success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getGeoRegionObj,_ := getGeoRegionRequestResult.Data. (model.GeoRegion)
	compareGeoRegion := cmp.Equal(createGeoRegionObj.ID, getGeoRegionObj.ID)
	
	if  compareGeoRegion == false	{
		t.Errorf( "Created GeoRegion object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllGeoRegionRequestResult := dao.GetAllGeoRegion()

	if getAllGeoRegionRequestResult.Success == false {
			t.Errorf(getAllGeoRegionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll GeoRegion success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllGeoRegionObj []model.GeoRegion = getAllGeoRegionRequestResult.Data. ([]model.GeoRegion)
		
	equalGeoRegion := cmp.Equal(createGeoRegionObj.ID, getAllGeoRegionObj[len(getAllGeoRegionObj)-1].ID)
		
	if equalGeoRegion == false {
		t.Errorf( "Created object is not equal to the last entry in GeoRegion[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for GeoRegion
	// --------------------------------------------------------------	
	deleteGeoRegionRequestResult := dao.DeleteGeoRegion(uint64(createGeoRegionObj.ID))

	if deleteGeoRegionRequestResult.Success == false {
			t.Errorf(deleteGeoRegionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion GeoRegion success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getGeoRegionRequestResult = dao.GetGeoRegion( uint64(createGeoRegionObj.ID) )
	
	if getGeoRegionRequestResult.Success == true {
		t.Errorf(getGeoRegionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

