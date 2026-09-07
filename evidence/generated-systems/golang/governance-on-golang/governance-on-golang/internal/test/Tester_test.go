package test

import ( 
	"testing"
    dao "governance-on-golang/internal/dao"
	"governance-on-golang/internal/model"
	"governance-on-golang/internal/utils"
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
	OrganizationObj := model.Organization                                                                                                                            {Name:"test value for Name",LegalName:"test value for LegalName",Jurisdiction:"test value for Jurisdiction",IndustrySector:"test value for IndustrySector"}

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


func TestGovernanceBodyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for GovernanceBody
	//----------------------------------------------------------------------------
	GovernanceBodyObj := model.GovernanceBody                                                                                            {Name:"test value for Name",CharterUrl:new URL(),Chair:"test value for Chair",BodyType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createGovernanceBodyRequestResult := dao.CreateGovernanceBody( GovernanceBodyObj )
	
	if createGovernanceBodyRequestResult.Success == false {
		t.Errorf(createGovernanceBodyRequestResult.Msg)
	} else {
		fmt.Println("Check Create GovernanceBody success...")
	}
	
	createGovernanceBodyObj,_ := createGovernanceBodyRequestResult.Data. (model.GovernanceBody)

	// --------------------------------------------------------------
	// Check GovernanceBody Obj ID
	// --------------------------------------------------------------	
	if createGovernanceBodyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for GovernanceBody" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getGovernanceBodyRequestResult := dao.GetGovernanceBody( uint64(createGovernanceBodyObj.ID) )
	
	if getGovernanceBodyRequestResult.Success == false {
		t.Errorf(getGovernanceBodyRequestResult.Msg)
	} else {
		fmt.Println("Check Get GovernanceBody success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getGovernanceBodyObj,_ := getGovernanceBodyRequestResult.Data. (model.GovernanceBody)
	compareGovernanceBody := cmp.Equal(createGovernanceBodyObj.ID, getGovernanceBodyObj.ID)
	
	if  compareGovernanceBody == false	{
		t.Errorf( "Created GovernanceBody object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllGovernanceBodyRequestResult := dao.GetAllGovernanceBody()

	if getAllGovernanceBodyRequestResult.Success == false {
			t.Errorf(getAllGovernanceBodyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll GovernanceBody success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllGovernanceBodyObj []model.GovernanceBody = getAllGovernanceBodyRequestResult.Data. ([]model.GovernanceBody)
		
	equalGovernanceBody := cmp.Equal(createGovernanceBodyObj.ID, getAllGovernanceBodyObj[len(getAllGovernanceBodyObj)-1].ID)
		
	if equalGovernanceBody == false {
		t.Errorf( "Created object is not equal to the last entry in GovernanceBody[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for GovernanceBody
	// --------------------------------------------------------------	
	deleteGovernanceBodyRequestResult := dao.DeleteGovernanceBody(uint64(createGovernanceBodyObj.ID))

	if deleteGovernanceBodyRequestResult.Success == false {
			t.Errorf(deleteGovernanceBodyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion GovernanceBody success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getGovernanceBodyRequestResult = dao.GetGovernanceBody( uint64(createGovernanceBodyObj.ID) )
	
	if getGovernanceBodyRequestResult.Success == true {
		t.Errorf(getGovernanceBodyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPersonCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Person
	//----------------------------------------------------------------------------
	PersonObj := model.Person                                                                                                            {FirstName:"test value for FirstName",LastName:"test value for LastName",Email:new EmailAddress(),Department:"test value for Department"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPersonRequestResult := dao.CreatePerson( PersonObj )
	
	if createPersonRequestResult.Success == false {
		t.Errorf(createPersonRequestResult.Msg)
	} else {
		fmt.Println("Check Create Person success...")
	}
	
	createPersonObj,_ := createPersonRequestResult.Data. (model.Person)

	// --------------------------------------------------------------
	// Check Person Obj ID
	// --------------------------------------------------------------	
	if createPersonObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Person" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPersonRequestResult := dao.GetPerson( uint64(createPersonObj.ID) )
	
	if getPersonRequestResult.Success == false {
		t.Errorf(getPersonRequestResult.Msg)
	} else {
		fmt.Println("Check Get Person success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPersonObj,_ := getPersonRequestResult.Data. (model.Person)
	comparePerson := cmp.Equal(createPersonObj.ID, getPersonObj.ID)
	
	if  comparePerson == false	{
		t.Errorf( "Created Person object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPersonRequestResult := dao.GetAllPerson()

	if getAllPersonRequestResult.Success == false {
			t.Errorf(getAllPersonRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Person success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPersonObj []model.Person = getAllPersonRequestResult.Data. ([]model.Person)
		
	equalPerson := cmp.Equal(createPersonObj.ID, getAllPersonObj[len(getAllPersonObj)-1].ID)
		
	if equalPerson == false {
		t.Errorf( "Created object is not equal to the last entry in Person[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Person
	// --------------------------------------------------------------	
	deletePersonRequestResult := dao.DeletePerson(uint64(createPersonObj.ID))

	if deletePersonRequestResult.Success == false {
			t.Errorf(deletePersonRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Person success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPersonRequestResult = dao.GetPerson( uint64(createPersonObj.ID) )
	
	if getPersonRequestResult.Success == true {
		t.Errorf(getPersonRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRoleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Role
	//----------------------------------------------------------------------------
	RoleObj := model.Role                                                            {Name:"test value for Name",Responsibility:"test value for Responsibility"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRoleRequestResult := dao.CreateRole( RoleObj )
	
	if createRoleRequestResult.Success == false {
		t.Errorf(createRoleRequestResult.Msg)
	} else {
		fmt.Println("Check Create Role success...")
	}
	
	createRoleObj,_ := createRoleRequestResult.Data. (model.Role)

	// --------------------------------------------------------------
	// Check Role Obj ID
	// --------------------------------------------------------------	
	if createRoleObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Role" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRoleRequestResult := dao.GetRole( uint64(createRoleObj.ID) )
	
	if getRoleRequestResult.Success == false {
		t.Errorf(getRoleRequestResult.Msg)
	} else {
		fmt.Println("Check Get Role success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRoleObj,_ := getRoleRequestResult.Data. (model.Role)
	compareRole := cmp.Equal(createRoleObj.ID, getRoleObj.ID)
	
	if  compareRole == false	{
		t.Errorf( "Created Role object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRoleRequestResult := dao.GetAllRole()

	if getAllRoleRequestResult.Success == false {
			t.Errorf(getAllRoleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Role success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRoleObj []model.Role = getAllRoleRequestResult.Data. ([]model.Role)
		
	equalRole := cmp.Equal(createRoleObj.ID, getAllRoleObj[len(getAllRoleObj)-1].ID)
		
	if equalRole == false {
		t.Errorf( "Created object is not equal to the last entry in Role[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Role
	// --------------------------------------------------------------	
	deleteRoleRequestResult := dao.DeleteRole(uint64(createRoleObj.ID))

	if deleteRoleRequestResult.Success == false {
			t.Errorf(deleteRoleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Role success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRoleRequestResult = dao.GetRole( uint64(createRoleObj.ID) )
	
	if getRoleRequestResult.Success == true {
		t.Errorf(getRoleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRoleAssignmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RoleAssignment
	//----------------------------------------------------------------------------
	RoleAssignmentObj := model.RoleAssignment                                                                                                            {EffectiveFrom:time.Now(),EffectiveTo:time.Now()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRoleAssignmentRequestResult := dao.CreateRoleAssignment( RoleAssignmentObj )
	
	if createRoleAssignmentRequestResult.Success == false {
		t.Errorf(createRoleAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create RoleAssignment success...")
	}
	
	createRoleAssignmentObj,_ := createRoleAssignmentRequestResult.Data. (model.RoleAssignment)

	// --------------------------------------------------------------
	// Check RoleAssignment Obj ID
	// --------------------------------------------------------------	
	if createRoleAssignmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for RoleAssignment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRoleAssignmentRequestResult := dao.GetRoleAssignment( uint64(createRoleAssignmentObj.ID) )
	
	if getRoleAssignmentRequestResult.Success == false {
		t.Errorf(getRoleAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get RoleAssignment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRoleAssignmentObj,_ := getRoleAssignmentRequestResult.Data. (model.RoleAssignment)
	compareRoleAssignment := cmp.Equal(createRoleAssignmentObj.ID, getRoleAssignmentObj.ID)
	
	if  compareRoleAssignment == false	{
		t.Errorf( "Created RoleAssignment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRoleAssignmentRequestResult := dao.GetAllRoleAssignment()

	if getAllRoleAssignmentRequestResult.Success == false {
			t.Errorf(getAllRoleAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RoleAssignment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRoleAssignmentObj []model.RoleAssignment = getAllRoleAssignmentRequestResult.Data. ([]model.RoleAssignment)
		
	equalRoleAssignment := cmp.Equal(createRoleAssignmentObj.ID, getAllRoleAssignmentObj[len(getAllRoleAssignmentObj)-1].ID)
		
	if equalRoleAssignment == false {
		t.Errorf( "Created object is not equal to the last entry in RoleAssignment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RoleAssignment
	// --------------------------------------------------------------	
	deleteRoleAssignmentRequestResult := dao.DeleteRoleAssignment(uint64(createRoleAssignmentObj.ID))

	if deleteRoleAssignmentRequestResult.Success == false {
			t.Errorf(deleteRoleAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RoleAssignment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRoleAssignmentRequestResult = dao.GetRoleAssignment( uint64(createRoleAssignmentObj.ID) )
	
	if getRoleAssignmentRequestResult.Success == true {
		t.Errorf(getRoleAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPolicyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Policy
	//----------------------------------------------------------------------------
	PolicyObj := model.Policy                                                                                                                                                                                                                            {Title:"test value for Title",VersionLabel:"test value for VersionLabel",ApprovalDate:time.Now(),NextReviewDate:time.Now(),DocumentUrl:new URL(),PolicyType:0,Status:0}

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


func TestProcedureCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Procedure
	//----------------------------------------------------------------------------
	ProcedureObj := model.Procedure                                                                            {Title:"test value for Title",VersionLabel:"test value for VersionLabel",Status:0}

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


func TestRegulationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Regulation
	//----------------------------------------------------------------------------
	RegulationObj := model.Regulation                                                                                                            {Name:"test value for Name",Citation:"test value for Citation",Jurisdiction:"test value for Jurisdiction",PublicationUrl:new URL()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRegulationRequestResult := dao.CreateRegulation( RegulationObj )
	
	if createRegulationRequestResult.Success == false {
		t.Errorf(createRegulationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Regulation success...")
	}
	
	createRegulationObj,_ := createRegulationRequestResult.Data. (model.Regulation)

	// --------------------------------------------------------------
	// Check Regulation Obj ID
	// --------------------------------------------------------------	
	if createRegulationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Regulation" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRegulationRequestResult := dao.GetRegulation( uint64(createRegulationObj.ID) )
	
	if getRegulationRequestResult.Success == false {
		t.Errorf(getRegulationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Regulation success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRegulationObj,_ := getRegulationRequestResult.Data. (model.Regulation)
	compareRegulation := cmp.Equal(createRegulationObj.ID, getRegulationObj.ID)
	
	if  compareRegulation == false	{
		t.Errorf( "Created Regulation object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRegulationRequestResult := dao.GetAllRegulation()

	if getAllRegulationRequestResult.Success == false {
			t.Errorf(getAllRegulationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Regulation success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRegulationObj []model.Regulation = getAllRegulationRequestResult.Data. ([]model.Regulation)
		
	equalRegulation := cmp.Equal(createRegulationObj.ID, getAllRegulationObj[len(getAllRegulationObj)-1].ID)
		
	if equalRegulation == false {
		t.Errorf( "Created object is not equal to the last entry in Regulation[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Regulation
	// --------------------------------------------------------------	
	deleteRegulationRequestResult := dao.DeleteRegulation(uint64(createRegulationObj.ID))

	if deleteRegulationRequestResult.Success == false {
			t.Errorf(deleteRegulationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Regulation success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRegulationRequestResult = dao.GetRegulation( uint64(createRegulationObj.ID) )
	
	if getRegulationRequestResult.Success == true {
		t.Errorf(getRegulationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestObligationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Obligation
	//----------------------------------------------------------------------------
	ObligationObj := model.Obligation                                                                                            {ReferenceNumber:"test value for ReferenceNumber",DescriptionText:"test value for DescriptionText",ObligationType:0,ReviewFrequency:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createObligationRequestResult := dao.CreateObligation( ObligationObj )
	
	if createObligationRequestResult.Success == false {
		t.Errorf(createObligationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Obligation success...")
	}
	
	createObligationObj,_ := createObligationRequestResult.Data. (model.Obligation)

	// --------------------------------------------------------------
	// Check Obligation Obj ID
	// --------------------------------------------------------------	
	if createObligationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Obligation" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getObligationRequestResult := dao.GetObligation( uint64(createObligationObj.ID) )
	
	if getObligationRequestResult.Success == false {
		t.Errorf(getObligationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Obligation success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getObligationObj,_ := getObligationRequestResult.Data. (model.Obligation)
	compareObligation := cmp.Equal(createObligationObj.ID, getObligationObj.ID)
	
	if  compareObligation == false	{
		t.Errorf( "Created Obligation object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllObligationRequestResult := dao.GetAllObligation()

	if getAllObligationRequestResult.Success == false {
			t.Errorf(getAllObligationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Obligation success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllObligationObj []model.Obligation = getAllObligationRequestResult.Data. ([]model.Obligation)
		
	equalObligation := cmp.Equal(createObligationObj.ID, getAllObligationObj[len(getAllObligationObj)-1].ID)
		
	if equalObligation == false {
		t.Errorf( "Created object is not equal to the last entry in Obligation[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Obligation
	// --------------------------------------------------------------	
	deleteObligationRequestResult := dao.DeleteObligation(uint64(createObligationObj.ID))

	if deleteObligationRequestResult.Success == false {
			t.Errorf(deleteObligationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Obligation success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getObligationRequestResult = dao.GetObligation( uint64(createObligationObj.ID) )
	
	if getObligationRequestResult.Success == true {
		t.Errorf(getObligationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestControlCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Control
	//----------------------------------------------------------------------------
	ControlObj := model.Control                                                                                                                                            {Name:"test value for Name",Objective:"test value for Objective",OwnerDepartment:"test value for OwnerDepartment",ControlType:0,Frequency:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createControlRequestResult := dao.CreateControl( ControlObj )
	
	if createControlRequestResult.Success == false {
		t.Errorf(createControlRequestResult.Msg)
	} else {
		fmt.Println("Check Create Control success...")
	}
	
	createControlObj,_ := createControlRequestResult.Data. (model.Control)

	// --------------------------------------------------------------
	// Check Control Obj ID
	// --------------------------------------------------------------	
	if createControlObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Control" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getControlRequestResult := dao.GetControl( uint64(createControlObj.ID) )
	
	if getControlRequestResult.Success == false {
		t.Errorf(getControlRequestResult.Msg)
	} else {
		fmt.Println("Check Get Control success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getControlObj,_ := getControlRequestResult.Data. (model.Control)
	compareControl := cmp.Equal(createControlObj.ID, getControlObj.ID)
	
	if  compareControl == false	{
		t.Errorf( "Created Control object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllControlRequestResult := dao.GetAllControl()

	if getAllControlRequestResult.Success == false {
			t.Errorf(getAllControlRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Control success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllControlObj []model.Control = getAllControlRequestResult.Data. ([]model.Control)
		
	equalControl := cmp.Equal(createControlObj.ID, getAllControlObj[len(getAllControlObj)-1].ID)
		
	if equalControl == false {
		t.Errorf( "Created object is not equal to the last entry in Control[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Control
	// --------------------------------------------------------------	
	deleteControlRequestResult := dao.DeleteControl(uint64(createControlObj.ID))

	if deleteControlRequestResult.Success == false {
			t.Errorf(deleteControlRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Control success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getControlRequestResult = dao.GetControl( uint64(createControlObj.ID) )
	
	if getControlRequestResult.Success == true {
		t.Errorf(getControlRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestControlTest_CRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ControlTest_
	//----------------------------------------------------------------------------
	ControlTest_Obj := model.ControlTest_                                                                                                                                                                                                                            {Name:"test value for Name",TestPeriodStart:time.Now(),TestPeriodEnd:time.Now(),SampleSize:100,TestType:0,Effectiveness:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createControlTest_RequestResult := dao.CreateControlTest_( ControlTest_Obj )
	
	if createControlTest_RequestResult.Success == false {
		t.Errorf(createControlTest_RequestResult.Msg)
	} else {
		fmt.Println("Check Create ControlTest_ success...")
	}
	
	createControlTest_Obj,_ := createControlTest_RequestResult.Data. (model.ControlTest_)

	// --------------------------------------------------------------
	// Check ControlTest_ Obj ID
	// --------------------------------------------------------------	
	if createControlTest_Obj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ControlTest_" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getControlTest_RequestResult := dao.GetControlTest_( uint64(createControlTest_Obj.ID) )
	
	if getControlTest_RequestResult.Success == false {
		t.Errorf(getControlTest_RequestResult.Msg)
	} else {
		fmt.Println("Check Get ControlTest_ success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getControlTest_Obj,_ := getControlTest_RequestResult.Data. (model.ControlTest_)
	compareControlTest_ := cmp.Equal(createControlTest_Obj.ID, getControlTest_Obj.ID)
	
	if  compareControlTest_ == false	{
		t.Errorf( "Created ControlTest_ object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllControlTest_RequestResult := dao.GetAllControlTest_()

	if getAllControlTest_RequestResult.Success == false {
			t.Errorf(getAllControlTest_RequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ControlTest_ success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllControlTest_Obj []model.ControlTest_ = getAllControlTest_RequestResult.Data. ([]model.ControlTest_)
		
	equalControlTest_ := cmp.Equal(createControlTest_Obj.ID, getAllControlTest_Obj[len(getAllControlTest_Obj)-1].ID)
		
	if equalControlTest_ == false {
		t.Errorf( "Created object is not equal to the last entry in ControlTest_[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ControlTest_
	// --------------------------------------------------------------	
	deleteControlTest_RequestResult := dao.DeleteControlTest_(uint64(createControlTest_Obj.ID))

	if deleteControlTest_RequestResult.Success == false {
			t.Errorf(deleteControlTest_RequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ControlTest_ success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getControlTest_RequestResult = dao.GetControlTest_( uint64(createControlTest_Obj.ID) )
	
	if getControlTest_RequestResult.Success == true {
		t.Errorf(getControlTest_RequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEvidenceCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Evidence
	//----------------------------------------------------------------------------
	EvidenceObj := model.Evidence                                                                                                                    {Title:"test value for Title",LocationUrl:new URL(),ReceivedDate:time.Now(),EvidenceType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEvidenceRequestResult := dao.CreateEvidence( EvidenceObj )
	
	if createEvidenceRequestResult.Success == false {
		t.Errorf(createEvidenceRequestResult.Msg)
	} else {
		fmt.Println("Check Create Evidence success...")
	}
	
	createEvidenceObj,_ := createEvidenceRequestResult.Data. (model.Evidence)

	// --------------------------------------------------------------
	// Check Evidence Obj ID
	// --------------------------------------------------------------	
	if createEvidenceObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Evidence" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEvidenceRequestResult := dao.GetEvidence( uint64(createEvidenceObj.ID) )
	
	if getEvidenceRequestResult.Success == false {
		t.Errorf(getEvidenceRequestResult.Msg)
	} else {
		fmt.Println("Check Get Evidence success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEvidenceObj,_ := getEvidenceRequestResult.Data. (model.Evidence)
	compareEvidence := cmp.Equal(createEvidenceObj.ID, getEvidenceObj.ID)
	
	if  compareEvidence == false	{
		t.Errorf( "Created Evidence object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEvidenceRequestResult := dao.GetAllEvidence()

	if getAllEvidenceRequestResult.Success == false {
			t.Errorf(getAllEvidenceRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Evidence success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEvidenceObj []model.Evidence = getAllEvidenceRequestResult.Data. ([]model.Evidence)
		
	equalEvidence := cmp.Equal(createEvidenceObj.ID, getAllEvidenceObj[len(getAllEvidenceObj)-1].ID)
		
	if equalEvidence == false {
		t.Errorf( "Created object is not equal to the last entry in Evidence[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Evidence
	// --------------------------------------------------------------	
	deleteEvidenceRequestResult := dao.DeleteEvidence(uint64(createEvidenceObj.ID))

	if deleteEvidenceRequestResult.Success == false {
			t.Errorf(deleteEvidenceRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Evidence success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEvidenceRequestResult = dao.GetEvidence( uint64(createEvidenceObj.ID) )
	
	if getEvidenceRequestResult.Success == true {
		t.Errorf(getEvidenceRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRiskCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Risk
	//----------------------------------------------------------------------------
	RiskObj := model.Risk                                                                                                                                                                                            {Name:"test value for Name",Description:"test value for Description",InherentRiskScore:100,ResidualRiskScore:100,Category:0,Impact:0,Likelihood:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRiskRequestResult := dao.CreateRisk( RiskObj )
	
	if createRiskRequestResult.Success == false {
		t.Errorf(createRiskRequestResult.Msg)
	} else {
		fmt.Println("Check Create Risk success...")
	}
	
	createRiskObj,_ := createRiskRequestResult.Data. (model.Risk)

	// --------------------------------------------------------------
	// Check Risk Obj ID
	// --------------------------------------------------------------	
	if createRiskObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Risk" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRiskRequestResult := dao.GetRisk( uint64(createRiskObj.ID) )
	
	if getRiskRequestResult.Success == false {
		t.Errorf(getRiskRequestResult.Msg)
	} else {
		fmt.Println("Check Get Risk success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRiskObj,_ := getRiskRequestResult.Data. (model.Risk)
	compareRisk := cmp.Equal(createRiskObj.ID, getRiskObj.ID)
	
	if  compareRisk == false	{
		t.Errorf( "Created Risk object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRiskRequestResult := dao.GetAllRisk()

	if getAllRiskRequestResult.Success == false {
			t.Errorf(getAllRiskRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Risk success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRiskObj []model.Risk = getAllRiskRequestResult.Data. ([]model.Risk)
		
	equalRisk := cmp.Equal(createRiskObj.ID, getAllRiskObj[len(getAllRiskObj)-1].ID)
		
	if equalRisk == false {
		t.Errorf( "Created object is not equal to the last entry in Risk[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Risk
	// --------------------------------------------------------------	
	deleteRiskRequestResult := dao.DeleteRisk(uint64(createRiskObj.ID))

	if deleteRiskRequestResult.Success == false {
			t.Errorf(deleteRiskRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Risk success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRiskRequestResult = dao.GetRisk( uint64(createRiskObj.ID) )
	
	if getRiskRequestResult.Success == true {
		t.Errorf(getRiskRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRiskAssessmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RiskAssessment
	//----------------------------------------------------------------------------
	RiskAssessmentObj := model.RiskAssessment                                                                                                                                    {AssessmentDate:time.Now(),Assessor:"test value for Assessor",Summary:"test value for Summary",AssessmentType:0}

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


func TestComplianceProgramCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ComplianceProgram
	//----------------------------------------------------------------------------
	ComplianceProgramObj := model.ComplianceProgram                                                                            {Name:"test value for Name",Framework:"test value for Framework",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createComplianceProgramRequestResult := dao.CreateComplianceProgram( ComplianceProgramObj )
	
	if createComplianceProgramRequestResult.Success == false {
		t.Errorf(createComplianceProgramRequestResult.Msg)
	} else {
		fmt.Println("Check Create ComplianceProgram success...")
	}
	
	createComplianceProgramObj,_ := createComplianceProgramRequestResult.Data. (model.ComplianceProgram)

	// --------------------------------------------------------------
	// Check ComplianceProgram Obj ID
	// --------------------------------------------------------------	
	if createComplianceProgramObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ComplianceProgram" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getComplianceProgramRequestResult := dao.GetComplianceProgram( uint64(createComplianceProgramObj.ID) )
	
	if getComplianceProgramRequestResult.Success == false {
		t.Errorf(getComplianceProgramRequestResult.Msg)
	} else {
		fmt.Println("Check Get ComplianceProgram success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getComplianceProgramObj,_ := getComplianceProgramRequestResult.Data. (model.ComplianceProgram)
	compareComplianceProgram := cmp.Equal(createComplianceProgramObj.ID, getComplianceProgramObj.ID)
	
	if  compareComplianceProgram == false	{
		t.Errorf( "Created ComplianceProgram object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllComplianceProgramRequestResult := dao.GetAllComplianceProgram()

	if getAllComplianceProgramRequestResult.Success == false {
			t.Errorf(getAllComplianceProgramRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ComplianceProgram success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllComplianceProgramObj []model.ComplianceProgram = getAllComplianceProgramRequestResult.Data. ([]model.ComplianceProgram)
		
	equalComplianceProgram := cmp.Equal(createComplianceProgramObj.ID, getAllComplianceProgramObj[len(getAllComplianceProgramObj)-1].ID)
		
	if equalComplianceProgram == false {
		t.Errorf( "Created object is not equal to the last entry in ComplianceProgram[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ComplianceProgram
	// --------------------------------------------------------------	
	deleteComplianceProgramRequestResult := dao.DeleteComplianceProgram(uint64(createComplianceProgramObj.ID))

	if deleteComplianceProgramRequestResult.Success == false {
			t.Errorf(deleteComplianceProgramRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ComplianceProgram success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getComplianceProgramRequestResult = dao.GetComplianceProgram( uint64(createComplianceProgramObj.ID) )
	
	if getComplianceProgramRequestResult.Success == true {
		t.Errorf(getComplianceProgramRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestComplianceRequirementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ComplianceRequirement
	//----------------------------------------------------------------------------
	ComplianceRequirementObj := model.ComplianceRequirement                                                                                                                            {Name:"test value for Name",Source:"test value for Source",Citation:"test value for Citation",Applicability:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createComplianceRequirementRequestResult := dao.CreateComplianceRequirement( ComplianceRequirementObj )
	
	if createComplianceRequirementRequestResult.Success == false {
		t.Errorf(createComplianceRequirementRequestResult.Msg)
	} else {
		fmt.Println("Check Create ComplianceRequirement success...")
	}
	
	createComplianceRequirementObj,_ := createComplianceRequirementRequestResult.Data. (model.ComplianceRequirement)

	// --------------------------------------------------------------
	// Check ComplianceRequirement Obj ID
	// --------------------------------------------------------------	
	if createComplianceRequirementObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ComplianceRequirement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getComplianceRequirementRequestResult := dao.GetComplianceRequirement( uint64(createComplianceRequirementObj.ID) )
	
	if getComplianceRequirementRequestResult.Success == false {
		t.Errorf(getComplianceRequirementRequestResult.Msg)
	} else {
		fmt.Println("Check Get ComplianceRequirement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getComplianceRequirementObj,_ := getComplianceRequirementRequestResult.Data. (model.ComplianceRequirement)
	compareComplianceRequirement := cmp.Equal(createComplianceRequirementObj.ID, getComplianceRequirementObj.ID)
	
	if  compareComplianceRequirement == false	{
		t.Errorf( "Created ComplianceRequirement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllComplianceRequirementRequestResult := dao.GetAllComplianceRequirement()

	if getAllComplianceRequirementRequestResult.Success == false {
			t.Errorf(getAllComplianceRequirementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ComplianceRequirement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllComplianceRequirementObj []model.ComplianceRequirement = getAllComplianceRequirementRequestResult.Data. ([]model.ComplianceRequirement)
		
	equalComplianceRequirement := cmp.Equal(createComplianceRequirementObj.ID, getAllComplianceRequirementObj[len(getAllComplianceRequirementObj)-1].ID)
		
	if equalComplianceRequirement == false {
		t.Errorf( "Created object is not equal to the last entry in ComplianceRequirement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ComplianceRequirement
	// --------------------------------------------------------------	
	deleteComplianceRequirementRequestResult := dao.DeleteComplianceRequirement(uint64(createComplianceRequirementObj.ID))

	if deleteComplianceRequirementRequestResult.Success == false {
			t.Errorf(deleteComplianceRequirementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ComplianceRequirement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getComplianceRequirementRequestResult = dao.GetComplianceRequirement( uint64(createComplianceRequirementObj.ID) )
	
	if getComplianceRequirementRequestResult.Success == true {
		t.Errorf(getComplianceRequirementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAttestationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Attestation
	//----------------------------------------------------------------------------
	AttestationObj := model.Attestation                                                                                                                                    {Statement:"test value for Statement",Attestor:"test value for Attestor",DateSigned:time.Now(),Result:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAttestationRequestResult := dao.CreateAttestation( AttestationObj )
	
	if createAttestationRequestResult.Success == false {
		t.Errorf(createAttestationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Attestation success...")
	}
	
	createAttestationObj,_ := createAttestationRequestResult.Data. (model.Attestation)

	// --------------------------------------------------------------
	// Check Attestation Obj ID
	// --------------------------------------------------------------	
	if createAttestationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Attestation" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAttestationRequestResult := dao.GetAttestation( uint64(createAttestationObj.ID) )
	
	if getAttestationRequestResult.Success == false {
		t.Errorf(getAttestationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Attestation success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAttestationObj,_ := getAttestationRequestResult.Data. (model.Attestation)
	compareAttestation := cmp.Equal(createAttestationObj.ID, getAttestationObj.ID)
	
	if  compareAttestation == false	{
		t.Errorf( "Created Attestation object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAttestationRequestResult := dao.GetAllAttestation()

	if getAllAttestationRequestResult.Success == false {
			t.Errorf(getAllAttestationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Attestation success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAttestationObj []model.Attestation = getAllAttestationRequestResult.Data. ([]model.Attestation)
		
	equalAttestation := cmp.Equal(createAttestationObj.ID, getAllAttestationObj[len(getAllAttestationObj)-1].ID)
		
	if equalAttestation == false {
		t.Errorf( "Created object is not equal to the last entry in Attestation[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Attestation
	// --------------------------------------------------------------	
	deleteAttestationRequestResult := dao.DeleteAttestation(uint64(createAttestationObj.ID))

	if deleteAttestationRequestResult.Success == false {
			t.Errorf(deleteAttestationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Attestation success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAttestationRequestResult = dao.GetAttestation( uint64(createAttestationObj.ID) )
	
	if getAttestationRequestResult.Success == true {
		t.Errorf(getAttestationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAuditProgramCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AuditProgram
	//----------------------------------------------------------------------------
	AuditProgramObj := model.AuditProgram                                                                                            {Name:"test value for Name",Scope:"test value for Scope",Cycle:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAuditProgramRequestResult := dao.CreateAuditProgram( AuditProgramObj )
	
	if createAuditProgramRequestResult.Success == false {
		t.Errorf(createAuditProgramRequestResult.Msg)
	} else {
		fmt.Println("Check Create AuditProgram success...")
	}
	
	createAuditProgramObj,_ := createAuditProgramRequestResult.Data. (model.AuditProgram)

	// --------------------------------------------------------------
	// Check AuditProgram Obj ID
	// --------------------------------------------------------------	
	if createAuditProgramObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AuditProgram" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAuditProgramRequestResult := dao.GetAuditProgram( uint64(createAuditProgramObj.ID) )
	
	if getAuditProgramRequestResult.Success == false {
		t.Errorf(getAuditProgramRequestResult.Msg)
	} else {
		fmt.Println("Check Get AuditProgram success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAuditProgramObj,_ := getAuditProgramRequestResult.Data. (model.AuditProgram)
	compareAuditProgram := cmp.Equal(createAuditProgramObj.ID, getAuditProgramObj.ID)
	
	if  compareAuditProgram == false	{
		t.Errorf( "Created AuditProgram object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAuditProgramRequestResult := dao.GetAllAuditProgram()

	if getAllAuditProgramRequestResult.Success == false {
			t.Errorf(getAllAuditProgramRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AuditProgram success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAuditProgramObj []model.AuditProgram = getAllAuditProgramRequestResult.Data. ([]model.AuditProgram)
		
	equalAuditProgram := cmp.Equal(createAuditProgramObj.ID, getAllAuditProgramObj[len(getAllAuditProgramObj)-1].ID)
		
	if equalAuditProgram == false {
		t.Errorf( "Created object is not equal to the last entry in AuditProgram[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AuditProgram
	// --------------------------------------------------------------	
	deleteAuditProgramRequestResult := dao.DeleteAuditProgram(uint64(createAuditProgramObj.ID))

	if deleteAuditProgramRequestResult.Success == false {
			t.Errorf(deleteAuditProgramRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AuditProgram success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAuditProgramRequestResult = dao.GetAuditProgram( uint64(createAuditProgramObj.ID) )
	
	if getAuditProgramRequestResult.Success == true {
		t.Errorf(getAuditProgramRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAuditEngagementCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AuditEngagement
	//----------------------------------------------------------------------------
	AuditEngagementObj := model.AuditEngagement                                                                                                                                                            {Title:"test value for Title",StartDate:time.Now(),EndDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAuditEngagementRequestResult := dao.CreateAuditEngagement( AuditEngagementObj )
	
	if createAuditEngagementRequestResult.Success == false {
		t.Errorf(createAuditEngagementRequestResult.Msg)
	} else {
		fmt.Println("Check Create AuditEngagement success...")
	}
	
	createAuditEngagementObj,_ := createAuditEngagementRequestResult.Data. (model.AuditEngagement)

	// --------------------------------------------------------------
	// Check AuditEngagement Obj ID
	// --------------------------------------------------------------	
	if createAuditEngagementObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AuditEngagement" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAuditEngagementRequestResult := dao.GetAuditEngagement( uint64(createAuditEngagementObj.ID) )
	
	if getAuditEngagementRequestResult.Success == false {
		t.Errorf(getAuditEngagementRequestResult.Msg)
	} else {
		fmt.Println("Check Get AuditEngagement success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAuditEngagementObj,_ := getAuditEngagementRequestResult.Data. (model.AuditEngagement)
	compareAuditEngagement := cmp.Equal(createAuditEngagementObj.ID, getAuditEngagementObj.ID)
	
	if  compareAuditEngagement == false	{
		t.Errorf( "Created AuditEngagement object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAuditEngagementRequestResult := dao.GetAllAuditEngagement()

	if getAllAuditEngagementRequestResult.Success == false {
			t.Errorf(getAllAuditEngagementRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AuditEngagement success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAuditEngagementObj []model.AuditEngagement = getAllAuditEngagementRequestResult.Data. ([]model.AuditEngagement)
		
	equalAuditEngagement := cmp.Equal(createAuditEngagementObj.ID, getAllAuditEngagementObj[len(getAllAuditEngagementObj)-1].ID)
		
	if equalAuditEngagement == false {
		t.Errorf( "Created object is not equal to the last entry in AuditEngagement[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AuditEngagement
	// --------------------------------------------------------------	
	deleteAuditEngagementRequestResult := dao.DeleteAuditEngagement(uint64(createAuditEngagementObj.ID))

	if deleteAuditEngagementRequestResult.Success == false {
			t.Errorf(deleteAuditEngagementRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AuditEngagement success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAuditEngagementRequestResult = dao.GetAuditEngagement( uint64(createAuditEngagementObj.ID) )
	
	if getAuditEngagementRequestResult.Success == true {
		t.Errorf(getAuditEngagementRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAuditWorkpaperCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AuditWorkpaper
	//----------------------------------------------------------------------------
	AuditWorkpaperObj := model.AuditWorkpaper                                                                            {WorkpaperRef:"test value for WorkpaperRef",Subject:"test value for Subject",WorkpaperUrl:new URL()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAuditWorkpaperRequestResult := dao.CreateAuditWorkpaper( AuditWorkpaperObj )
	
	if createAuditWorkpaperRequestResult.Success == false {
		t.Errorf(createAuditWorkpaperRequestResult.Msg)
	} else {
		fmt.Println("Check Create AuditWorkpaper success...")
	}
	
	createAuditWorkpaperObj,_ := createAuditWorkpaperRequestResult.Data. (model.AuditWorkpaper)

	// --------------------------------------------------------------
	// Check AuditWorkpaper Obj ID
	// --------------------------------------------------------------	
	if createAuditWorkpaperObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AuditWorkpaper" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAuditWorkpaperRequestResult := dao.GetAuditWorkpaper( uint64(createAuditWorkpaperObj.ID) )
	
	if getAuditWorkpaperRequestResult.Success == false {
		t.Errorf(getAuditWorkpaperRequestResult.Msg)
	} else {
		fmt.Println("Check Get AuditWorkpaper success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAuditWorkpaperObj,_ := getAuditWorkpaperRequestResult.Data. (model.AuditWorkpaper)
	compareAuditWorkpaper := cmp.Equal(createAuditWorkpaperObj.ID, getAuditWorkpaperObj.ID)
	
	if  compareAuditWorkpaper == false	{
		t.Errorf( "Created AuditWorkpaper object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAuditWorkpaperRequestResult := dao.GetAllAuditWorkpaper()

	if getAllAuditWorkpaperRequestResult.Success == false {
			t.Errorf(getAllAuditWorkpaperRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AuditWorkpaper success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAuditWorkpaperObj []model.AuditWorkpaper = getAllAuditWorkpaperRequestResult.Data. ([]model.AuditWorkpaper)
		
	equalAuditWorkpaper := cmp.Equal(createAuditWorkpaperObj.ID, getAllAuditWorkpaperObj[len(getAllAuditWorkpaperObj)-1].ID)
		
	if equalAuditWorkpaper == false {
		t.Errorf( "Created object is not equal to the last entry in AuditWorkpaper[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AuditWorkpaper
	// --------------------------------------------------------------	
	deleteAuditWorkpaperRequestResult := dao.DeleteAuditWorkpaper(uint64(createAuditWorkpaperObj.ID))

	if deleteAuditWorkpaperRequestResult.Success == false {
			t.Errorf(deleteAuditWorkpaperRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AuditWorkpaper success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAuditWorkpaperRequestResult = dao.GetAuditWorkpaper( uint64(createAuditWorkpaperObj.ID) )
	
	if getAuditWorkpaperRequestResult.Success == true {
		t.Errorf(getAuditWorkpaperRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAuditFindingCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AuditFinding
	//----------------------------------------------------------------------------
	AuditFindingObj := model.AuditFinding                                                                                                                                                    {Title:"test value for Title",Description:"test value for Description",DueDate:time.Now(),Severity:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAuditFindingRequestResult := dao.CreateAuditFinding( AuditFindingObj )
	
	if createAuditFindingRequestResult.Success == false {
		t.Errorf(createAuditFindingRequestResult.Msg)
	} else {
		fmt.Println("Check Create AuditFinding success...")
	}
	
	createAuditFindingObj,_ := createAuditFindingRequestResult.Data. (model.AuditFinding)

	// --------------------------------------------------------------
	// Check AuditFinding Obj ID
	// --------------------------------------------------------------	
	if createAuditFindingObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AuditFinding" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAuditFindingRequestResult := dao.GetAuditFinding( uint64(createAuditFindingObj.ID) )
	
	if getAuditFindingRequestResult.Success == false {
		t.Errorf(getAuditFindingRequestResult.Msg)
	} else {
		fmt.Println("Check Get AuditFinding success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAuditFindingObj,_ := getAuditFindingRequestResult.Data. (model.AuditFinding)
	compareAuditFinding := cmp.Equal(createAuditFindingObj.ID, getAuditFindingObj.ID)
	
	if  compareAuditFinding == false	{
		t.Errorf( "Created AuditFinding object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAuditFindingRequestResult := dao.GetAllAuditFinding()

	if getAllAuditFindingRequestResult.Success == false {
			t.Errorf(getAllAuditFindingRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AuditFinding success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAuditFindingObj []model.AuditFinding = getAllAuditFindingRequestResult.Data. ([]model.AuditFinding)
		
	equalAuditFinding := cmp.Equal(createAuditFindingObj.ID, getAllAuditFindingObj[len(getAllAuditFindingObj)-1].ID)
		
	if equalAuditFinding == false {
		t.Errorf( "Created object is not equal to the last entry in AuditFinding[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AuditFinding
	// --------------------------------------------------------------	
	deleteAuditFindingRequestResult := dao.DeleteAuditFinding(uint64(createAuditFindingObj.ID))

	if deleteAuditFindingRequestResult.Success == false {
			t.Errorf(deleteAuditFindingRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AuditFinding success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAuditFindingRequestResult = dao.GetAuditFinding( uint64(createAuditFindingObj.ID) )
	
	if getAuditFindingRequestResult.Success == true {
		t.Errorf(getAuditFindingRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCorrectiveActionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CorrectiveAction
	//----------------------------------------------------------------------------
	CorrectiveActionObj := model.CorrectiveAction                                                                                                                                    {ActionTitle:"test value for ActionTitle",Owner:"test value for Owner",TargetDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCorrectiveActionRequestResult := dao.CreateCorrectiveAction( CorrectiveActionObj )
	
	if createCorrectiveActionRequestResult.Success == false {
		t.Errorf(createCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Check Create CorrectiveAction success...")
	}
	
	createCorrectiveActionObj,_ := createCorrectiveActionRequestResult.Data. (model.CorrectiveAction)

	// --------------------------------------------------------------
	// Check CorrectiveAction Obj ID
	// --------------------------------------------------------------	
	if createCorrectiveActionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CorrectiveAction" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCorrectiveActionRequestResult := dao.GetCorrectiveAction( uint64(createCorrectiveActionObj.ID) )
	
	if getCorrectiveActionRequestResult.Success == false {
		t.Errorf(getCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Check Get CorrectiveAction success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCorrectiveActionObj,_ := getCorrectiveActionRequestResult.Data. (model.CorrectiveAction)
	compareCorrectiveAction := cmp.Equal(createCorrectiveActionObj.ID, getCorrectiveActionObj.ID)
	
	if  compareCorrectiveAction == false	{
		t.Errorf( "Created CorrectiveAction object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCorrectiveActionRequestResult := dao.GetAllCorrectiveAction()

	if getAllCorrectiveActionRequestResult.Success == false {
			t.Errorf(getAllCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CorrectiveAction success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCorrectiveActionObj []model.CorrectiveAction = getAllCorrectiveActionRequestResult.Data. ([]model.CorrectiveAction)
		
	equalCorrectiveAction := cmp.Equal(createCorrectiveActionObj.ID, getAllCorrectiveActionObj[len(getAllCorrectiveActionObj)-1].ID)
		
	if equalCorrectiveAction == false {
		t.Errorf( "Created object is not equal to the last entry in CorrectiveAction[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CorrectiveAction
	// --------------------------------------------------------------	
	deleteCorrectiveActionRequestResult := dao.DeleteCorrectiveAction(uint64(createCorrectiveActionObj.ID))

	if deleteCorrectiveActionRequestResult.Success == false {
			t.Errorf(deleteCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CorrectiveAction success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCorrectiveActionRequestResult = dao.GetCorrectiveAction( uint64(createCorrectiveActionObj.ID) )
	
	if getCorrectiveActionRequestResult.Success == true {
		t.Errorf(getCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestIssueCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Issue
	//----------------------------------------------------------------------------
	IssueObj := model.Issue                                                                                                                                                                                            {Title:"test value for Title",OpenedDate:time.Now(),ClosedDate:time.Now(),IssueType:0,Priority:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createIssueRequestResult := dao.CreateIssue( IssueObj )
	
	if createIssueRequestResult.Success == false {
		t.Errorf(createIssueRequestResult.Msg)
	} else {
		fmt.Println("Check Create Issue success...")
	}
	
	createIssueObj,_ := createIssueRequestResult.Data. (model.Issue)

	// --------------------------------------------------------------
	// Check Issue Obj ID
	// --------------------------------------------------------------	
	if createIssueObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Issue" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getIssueRequestResult := dao.GetIssue( uint64(createIssueObj.ID) )
	
	if getIssueRequestResult.Success == false {
		t.Errorf(getIssueRequestResult.Msg)
	} else {
		fmt.Println("Check Get Issue success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getIssueObj,_ := getIssueRequestResult.Data. (model.Issue)
	compareIssue := cmp.Equal(createIssueObj.ID, getIssueObj.ID)
	
	if  compareIssue == false	{
		t.Errorf( "Created Issue object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllIssueRequestResult := dao.GetAllIssue()

	if getAllIssueRequestResult.Success == false {
			t.Errorf(getAllIssueRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Issue success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllIssueObj []model.Issue = getAllIssueRequestResult.Data. ([]model.Issue)
		
	equalIssue := cmp.Equal(createIssueObj.ID, getAllIssueObj[len(getAllIssueObj)-1].ID)
		
	if equalIssue == false {
		t.Errorf( "Created object is not equal to the last entry in Issue[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Issue
	// --------------------------------------------------------------	
	deleteIssueRequestResult := dao.DeleteIssue(uint64(createIssueObj.ID))

	if deleteIssueRequestResult.Success == false {
			t.Errorf(deleteIssueRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Issue success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getIssueRequestResult = dao.GetIssue( uint64(createIssueObj.ID) )
	
	if getIssueRequestResult.Success == true {
		t.Errorf(getIssueRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBusinessUnitCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BusinessUnit
	//----------------------------------------------------------------------------
	BusinessUnitObj := model.BusinessUnit                                                            {Name:"test value for Name",Leader:"test value for Leader"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBusinessUnitRequestResult := dao.CreateBusinessUnit( BusinessUnitObj )
	
	if createBusinessUnitRequestResult.Success == false {
		t.Errorf(createBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Check Create BusinessUnit success...")
	}
	
	createBusinessUnitObj,_ := createBusinessUnitRequestResult.Data. (model.BusinessUnit)

	// --------------------------------------------------------------
	// Check BusinessUnit Obj ID
	// --------------------------------------------------------------	
	if createBusinessUnitObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BusinessUnit" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBusinessUnitRequestResult := dao.GetBusinessUnit( uint64(createBusinessUnitObj.ID) )
	
	if getBusinessUnitRequestResult.Success == false {
		t.Errorf(getBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Check Get BusinessUnit success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBusinessUnitObj,_ := getBusinessUnitRequestResult.Data. (model.BusinessUnit)
	compareBusinessUnit := cmp.Equal(createBusinessUnitObj.ID, getBusinessUnitObj.ID)
	
	if  compareBusinessUnit == false	{
		t.Errorf( "Created BusinessUnit object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBusinessUnitRequestResult := dao.GetAllBusinessUnit()

	if getAllBusinessUnitRequestResult.Success == false {
			t.Errorf(getAllBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BusinessUnit success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBusinessUnitObj []model.BusinessUnit = getAllBusinessUnitRequestResult.Data. ([]model.BusinessUnit)
		
	equalBusinessUnit := cmp.Equal(createBusinessUnitObj.ID, getAllBusinessUnitObj[len(getAllBusinessUnitObj)-1].ID)
		
	if equalBusinessUnit == false {
		t.Errorf( "Created object is not equal to the last entry in BusinessUnit[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BusinessUnit
	// --------------------------------------------------------------	
	deleteBusinessUnitRequestResult := dao.DeleteBusinessUnit(uint64(createBusinessUnitObj.ID))

	if deleteBusinessUnitRequestResult.Success == false {
			t.Errorf(deleteBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BusinessUnit success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBusinessUnitRequestResult = dao.GetBusinessUnit( uint64(createBusinessUnitObj.ID) )
	
	if getBusinessUnitRequestResult.Success == true {
		t.Errorf(getBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDataProcessingActivityCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DataProcessingActivity
	//----------------------------------------------------------------------------
	DataProcessingActivityObj := model.DataProcessingActivity                                                                                                                                    {Name:"test value for Name",Purpose:"test value for Purpose",StartDate:time.Now(),LawfulBasis:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDataProcessingActivityRequestResult := dao.CreateDataProcessingActivity( DataProcessingActivityObj )
	
	if createDataProcessingActivityRequestResult.Success == false {
		t.Errorf(createDataProcessingActivityRequestResult.Msg)
	} else {
		fmt.Println("Check Create DataProcessingActivity success...")
	}
	
	createDataProcessingActivityObj,_ := createDataProcessingActivityRequestResult.Data. (model.DataProcessingActivity)

	// --------------------------------------------------------------
	// Check DataProcessingActivity Obj ID
	// --------------------------------------------------------------	
	if createDataProcessingActivityObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DataProcessingActivity" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDataProcessingActivityRequestResult := dao.GetDataProcessingActivity( uint64(createDataProcessingActivityObj.ID) )
	
	if getDataProcessingActivityRequestResult.Success == false {
		t.Errorf(getDataProcessingActivityRequestResult.Msg)
	} else {
		fmt.Println("Check Get DataProcessingActivity success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDataProcessingActivityObj,_ := getDataProcessingActivityRequestResult.Data. (model.DataProcessingActivity)
	compareDataProcessingActivity := cmp.Equal(createDataProcessingActivityObj.ID, getDataProcessingActivityObj.ID)
	
	if  compareDataProcessingActivity == false	{
		t.Errorf( "Created DataProcessingActivity object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDataProcessingActivityRequestResult := dao.GetAllDataProcessingActivity()

	if getAllDataProcessingActivityRequestResult.Success == false {
			t.Errorf(getAllDataProcessingActivityRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DataProcessingActivity success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDataProcessingActivityObj []model.DataProcessingActivity = getAllDataProcessingActivityRequestResult.Data. ([]model.DataProcessingActivity)
		
	equalDataProcessingActivity := cmp.Equal(createDataProcessingActivityObj.ID, getAllDataProcessingActivityObj[len(getAllDataProcessingActivityObj)-1].ID)
		
	if equalDataProcessingActivity == false {
		t.Errorf( "Created object is not equal to the last entry in DataProcessingActivity[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DataProcessingActivity
	// --------------------------------------------------------------	
	deleteDataProcessingActivityRequestResult := dao.DeleteDataProcessingActivity(uint64(createDataProcessingActivityObj.ID))

	if deleteDataProcessingActivityRequestResult.Success == false {
			t.Errorf(deleteDataProcessingActivityRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DataProcessingActivity success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDataProcessingActivityRequestResult = dao.GetDataProcessingActivity( uint64(createDataProcessingActivityObj.ID) )
	
	if getDataProcessingActivityRequestResult.Success == true {
		t.Errorf(getDataProcessingActivityRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDataCategoryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DataCategory
	//----------------------------------------------------------------------------
	DataCategoryObj := model.DataCategory                                                                            {Name:"test value for Name",Description:"test value for Description",Classification:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDataCategoryRequestResult := dao.CreateDataCategory( DataCategoryObj )
	
	if createDataCategoryRequestResult.Success == false {
		t.Errorf(createDataCategoryRequestResult.Msg)
	} else {
		fmt.Println("Check Create DataCategory success...")
	}
	
	createDataCategoryObj,_ := createDataCategoryRequestResult.Data. (model.DataCategory)

	// --------------------------------------------------------------
	// Check DataCategory Obj ID
	// --------------------------------------------------------------	
	if createDataCategoryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DataCategory" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDataCategoryRequestResult := dao.GetDataCategory( uint64(createDataCategoryObj.ID) )
	
	if getDataCategoryRequestResult.Success == false {
		t.Errorf(getDataCategoryRequestResult.Msg)
	} else {
		fmt.Println("Check Get DataCategory success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDataCategoryObj,_ := getDataCategoryRequestResult.Data. (model.DataCategory)
	compareDataCategory := cmp.Equal(createDataCategoryObj.ID, getDataCategoryObj.ID)
	
	if  compareDataCategory == false	{
		t.Errorf( "Created DataCategory object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDataCategoryRequestResult := dao.GetAllDataCategory()

	if getAllDataCategoryRequestResult.Success == false {
			t.Errorf(getAllDataCategoryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DataCategory success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDataCategoryObj []model.DataCategory = getAllDataCategoryRequestResult.Data. ([]model.DataCategory)
		
	equalDataCategory := cmp.Equal(createDataCategoryObj.ID, getAllDataCategoryObj[len(getAllDataCategoryObj)-1].ID)
		
	if equalDataCategory == false {
		t.Errorf( "Created object is not equal to the last entry in DataCategory[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DataCategory
	// --------------------------------------------------------------	
	deleteDataCategoryRequestResult := dao.DeleteDataCategory(uint64(createDataCategoryObj.ID))

	if deleteDataCategoryRequestResult.Success == false {
			t.Errorf(deleteDataCategoryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DataCategory success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDataCategoryRequestResult = dao.GetDataCategory( uint64(createDataCategoryObj.ID) )
	
	if getDataCategoryRequestResult.Success == true {
		t.Errorf(getDataCategoryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSystem_CRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for System_
	//----------------------------------------------------------------------------
	System_Obj := model.System_                                                                            {Name:"test value for Name",OwnerDepartment:"test value for OwnerDepartment",SystemType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSystem_RequestResult := dao.CreateSystem_( System_Obj )
	
	if createSystem_RequestResult.Success == false {
		t.Errorf(createSystem_RequestResult.Msg)
	} else {
		fmt.Println("Check Create System_ success...")
	}
	
	createSystem_Obj,_ := createSystem_RequestResult.Data. (model.System_)

	// --------------------------------------------------------------
	// Check System_ Obj ID
	// --------------------------------------------------------------	
	if createSystem_Obj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for System_" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSystem_RequestResult := dao.GetSystem_( uint64(createSystem_Obj.ID) )
	
	if getSystem_RequestResult.Success == false {
		t.Errorf(getSystem_RequestResult.Msg)
	} else {
		fmt.Println("Check Get System_ success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSystem_Obj,_ := getSystem_RequestResult.Data. (model.System_)
	compareSystem_ := cmp.Equal(createSystem_Obj.ID, getSystem_Obj.ID)
	
	if  compareSystem_ == false	{
		t.Errorf( "Created System_ object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSystem_RequestResult := dao.GetAllSystem_()

	if getAllSystem_RequestResult.Success == false {
			t.Errorf(getAllSystem_RequestResult.Msg)
	} else {
		fmt.Println("Check GetAll System_ success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSystem_Obj []model.System_ = getAllSystem_RequestResult.Data. ([]model.System_)
		
	equalSystem_ := cmp.Equal(createSystem_Obj.ID, getAllSystem_Obj[len(getAllSystem_Obj)-1].ID)
		
	if equalSystem_ == false {
		t.Errorf( "Created object is not equal to the last entry in System_[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for System_
	// --------------------------------------------------------------	
	deleteSystem_RequestResult := dao.DeleteSystem_(uint64(createSystem_Obj.ID))

	if deleteSystem_RequestResult.Success == false {
			t.Errorf(deleteSystem_RequestResult.Msg)
	} else {
		fmt.Println("Check Deletion System_ success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSystem_RequestResult = dao.GetSystem_( uint64(createSystem_Obj.ID) )
	
	if getSystem_RequestResult.Success == true {
		t.Errorf(getSystem_RequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPrivacyNoticeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PrivacyNotice
	//----------------------------------------------------------------------------
	PrivacyNoticeObj := model.PrivacyNotice                                                                                                                                                                                    {Title:"test value for Title",Audience:"test value for Audience",VersionLabel:"test value for VersionLabel",PublicationDate:time.Now(),PublicationUrl:new URL(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPrivacyNoticeRequestResult := dao.CreatePrivacyNotice( PrivacyNoticeObj )
	
	if createPrivacyNoticeRequestResult.Success == false {
		t.Errorf(createPrivacyNoticeRequestResult.Msg)
	} else {
		fmt.Println("Check Create PrivacyNotice success...")
	}
	
	createPrivacyNoticeObj,_ := createPrivacyNoticeRequestResult.Data. (model.PrivacyNotice)

	// --------------------------------------------------------------
	// Check PrivacyNotice Obj ID
	// --------------------------------------------------------------	
	if createPrivacyNoticeObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PrivacyNotice" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPrivacyNoticeRequestResult := dao.GetPrivacyNotice( uint64(createPrivacyNoticeObj.ID) )
	
	if getPrivacyNoticeRequestResult.Success == false {
		t.Errorf(getPrivacyNoticeRequestResult.Msg)
	} else {
		fmt.Println("Check Get PrivacyNotice success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPrivacyNoticeObj,_ := getPrivacyNoticeRequestResult.Data. (model.PrivacyNotice)
	comparePrivacyNotice := cmp.Equal(createPrivacyNoticeObj.ID, getPrivacyNoticeObj.ID)
	
	if  comparePrivacyNotice == false	{
		t.Errorf( "Created PrivacyNotice object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPrivacyNoticeRequestResult := dao.GetAllPrivacyNotice()

	if getAllPrivacyNoticeRequestResult.Success == false {
			t.Errorf(getAllPrivacyNoticeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PrivacyNotice success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPrivacyNoticeObj []model.PrivacyNotice = getAllPrivacyNoticeRequestResult.Data. ([]model.PrivacyNotice)
		
	equalPrivacyNotice := cmp.Equal(createPrivacyNoticeObj.ID, getAllPrivacyNoticeObj[len(getAllPrivacyNoticeObj)-1].ID)
		
	if equalPrivacyNotice == false {
		t.Errorf( "Created object is not equal to the last entry in PrivacyNotice[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PrivacyNotice
	// --------------------------------------------------------------	
	deletePrivacyNoticeRequestResult := dao.DeletePrivacyNotice(uint64(createPrivacyNoticeObj.ID))

	if deletePrivacyNoticeRequestResult.Success == false {
			t.Errorf(deletePrivacyNoticeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PrivacyNotice success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPrivacyNoticeRequestResult = dao.GetPrivacyNotice( uint64(createPrivacyNoticeObj.ID) )
	
	if getPrivacyNoticeRequestResult.Success == true {
		t.Errorf(getPrivacyNoticeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDataSubjectRequestCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DataSubjectRequest
	//----------------------------------------------------------------------------
	DataSubjectRequestObj := model.DataSubjectRequest                                                                                                                                                                            {ReceivedDate:time.Now(),DueDate:time.Now(),RequesterCountry:"test value for RequesterCountry",RequestType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDataSubjectRequestRequestResult := dao.CreateDataSubjectRequest( DataSubjectRequestObj )
	
	if createDataSubjectRequestRequestResult.Success == false {
		t.Errorf(createDataSubjectRequestRequestResult.Msg)
	} else {
		fmt.Println("Check Create DataSubjectRequest success...")
	}
	
	createDataSubjectRequestObj,_ := createDataSubjectRequestRequestResult.Data. (model.DataSubjectRequest)

	// --------------------------------------------------------------
	// Check DataSubjectRequest Obj ID
	// --------------------------------------------------------------	
	if createDataSubjectRequestObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DataSubjectRequest" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDataSubjectRequestRequestResult := dao.GetDataSubjectRequest( uint64(createDataSubjectRequestObj.ID) )
	
	if getDataSubjectRequestRequestResult.Success == false {
		t.Errorf(getDataSubjectRequestRequestResult.Msg)
	} else {
		fmt.Println("Check Get DataSubjectRequest success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDataSubjectRequestObj,_ := getDataSubjectRequestRequestResult.Data. (model.DataSubjectRequest)
	compareDataSubjectRequest := cmp.Equal(createDataSubjectRequestObj.ID, getDataSubjectRequestObj.ID)
	
	if  compareDataSubjectRequest == false	{
		t.Errorf( "Created DataSubjectRequest object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDataSubjectRequestRequestResult := dao.GetAllDataSubjectRequest()

	if getAllDataSubjectRequestRequestResult.Success == false {
			t.Errorf(getAllDataSubjectRequestRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DataSubjectRequest success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDataSubjectRequestObj []model.DataSubjectRequest = getAllDataSubjectRequestRequestResult.Data. ([]model.DataSubjectRequest)
		
	equalDataSubjectRequest := cmp.Equal(createDataSubjectRequestObj.ID, getAllDataSubjectRequestObj[len(getAllDataSubjectRequestObj)-1].ID)
		
	if equalDataSubjectRequest == false {
		t.Errorf( "Created object is not equal to the last entry in DataSubjectRequest[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DataSubjectRequest
	// --------------------------------------------------------------	
	deleteDataSubjectRequestRequestResult := dao.DeleteDataSubjectRequest(uint64(createDataSubjectRequestObj.ID))

	if deleteDataSubjectRequestRequestResult.Success == false {
			t.Errorf(deleteDataSubjectRequestRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DataSubjectRequest success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDataSubjectRequestRequestResult = dao.GetDataSubjectRequest( uint64(createDataSubjectRequestObj.ID) )
	
	if getDataSubjectRequestRequestResult.Success == true {
		t.Errorf(getDataSubjectRequestRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRecordsRepositoryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RecordsRepository
	//----------------------------------------------------------------------------
	RecordsRepositoryObj := model.RecordsRepository                                                                                                            {Name:"test value for Name",Location:"test value for Location",OwnerDepartment:"test value for OwnerDepartment",RepositoryType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRecordsRepositoryRequestResult := dao.CreateRecordsRepository( RecordsRepositoryObj )
	
	if createRecordsRepositoryRequestResult.Success == false {
		t.Errorf(createRecordsRepositoryRequestResult.Msg)
	} else {
		fmt.Println("Check Create RecordsRepository success...")
	}
	
	createRecordsRepositoryObj,_ := createRecordsRepositoryRequestResult.Data. (model.RecordsRepository)

	// --------------------------------------------------------------
	// Check RecordsRepository Obj ID
	// --------------------------------------------------------------	
	if createRecordsRepositoryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for RecordsRepository" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRecordsRepositoryRequestResult := dao.GetRecordsRepository( uint64(createRecordsRepositoryObj.ID) )
	
	if getRecordsRepositoryRequestResult.Success == false {
		t.Errorf(getRecordsRepositoryRequestResult.Msg)
	} else {
		fmt.Println("Check Get RecordsRepository success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRecordsRepositoryObj,_ := getRecordsRepositoryRequestResult.Data. (model.RecordsRepository)
	compareRecordsRepository := cmp.Equal(createRecordsRepositoryObj.ID, getRecordsRepositoryObj.ID)
	
	if  compareRecordsRepository == false	{
		t.Errorf( "Created RecordsRepository object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRecordsRepositoryRequestResult := dao.GetAllRecordsRepository()

	if getAllRecordsRepositoryRequestResult.Success == false {
			t.Errorf(getAllRecordsRepositoryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RecordsRepository success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRecordsRepositoryObj []model.RecordsRepository = getAllRecordsRepositoryRequestResult.Data. ([]model.RecordsRepository)
		
	equalRecordsRepository := cmp.Equal(createRecordsRepositoryObj.ID, getAllRecordsRepositoryObj[len(getAllRecordsRepositoryObj)-1].ID)
		
	if equalRecordsRepository == false {
		t.Errorf( "Created object is not equal to the last entry in RecordsRepository[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RecordsRepository
	// --------------------------------------------------------------	
	deleteRecordsRepositoryRequestResult := dao.DeleteRecordsRepository(uint64(createRecordsRepositoryObj.ID))

	if deleteRecordsRepositoryRequestResult.Success == false {
			t.Errorf(deleteRecordsRepositoryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RecordsRepository success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRecordsRepositoryRequestResult = dao.GetRecordsRepository( uint64(createRecordsRepositoryObj.ID) )
	
	if getRecordsRepositoryRequestResult.Success == true {
		t.Errorf(getRecordsRepositoryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRecord_CRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Record_
	//----------------------------------------------------------------------------
	Record_Obj := model.Record_                                                                                                                                    {Title:"test value for Title",CreationDate:time.Now(),RecordType:0,Classification:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRecord_RequestResult := dao.CreateRecord_( Record_Obj )
	
	if createRecord_RequestResult.Success == false {
		t.Errorf(createRecord_RequestResult.Msg)
	} else {
		fmt.Println("Check Create Record_ success...")
	}
	
	createRecord_Obj,_ := createRecord_RequestResult.Data. (model.Record_)

	// --------------------------------------------------------------
	// Check Record_ Obj ID
	// --------------------------------------------------------------	
	if createRecord_Obj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Record_" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRecord_RequestResult := dao.GetRecord_( uint64(createRecord_Obj.ID) )
	
	if getRecord_RequestResult.Success == false {
		t.Errorf(getRecord_RequestResult.Msg)
	} else {
		fmt.Println("Check Get Record_ success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRecord_Obj,_ := getRecord_RequestResult.Data. (model.Record_)
	compareRecord_ := cmp.Equal(createRecord_Obj.ID, getRecord_Obj.ID)
	
	if  compareRecord_ == false	{
		t.Errorf( "Created Record_ object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRecord_RequestResult := dao.GetAllRecord_()

	if getAllRecord_RequestResult.Success == false {
			t.Errorf(getAllRecord_RequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Record_ success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRecord_Obj []model.Record_ = getAllRecord_RequestResult.Data. ([]model.Record_)
		
	equalRecord_ := cmp.Equal(createRecord_Obj.ID, getAllRecord_Obj[len(getAllRecord_Obj)-1].ID)
		
	if equalRecord_ == false {
		t.Errorf( "Created object is not equal to the last entry in Record_[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Record_
	// --------------------------------------------------------------	
	deleteRecord_RequestResult := dao.DeleteRecord_(uint64(createRecord_Obj.ID))

	if deleteRecord_RequestResult.Success == false {
			t.Errorf(deleteRecord_RequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Record_ success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRecord_RequestResult = dao.GetRecord_( uint64(createRecord_Obj.ID) )
	
	if getRecord_RequestResult.Success == true {
		t.Errorf(getRecord_RequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRetentionScheduleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RetentionSchedule
	//----------------------------------------------------------------------------
	RetentionScheduleObj := model.RetentionSchedule                                                                                                            {Name:"test value for Name",RetentionPeriodMonths:100,RetentionTrigger:0,DispositionAction:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRetentionScheduleRequestResult := dao.CreateRetentionSchedule( RetentionScheduleObj )
	
	if createRetentionScheduleRequestResult.Success == false {
		t.Errorf(createRetentionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Create RetentionSchedule success...")
	}
	
	createRetentionScheduleObj,_ := createRetentionScheduleRequestResult.Data. (model.RetentionSchedule)

	// --------------------------------------------------------------
	// Check RetentionSchedule Obj ID
	// --------------------------------------------------------------	
	if createRetentionScheduleObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for RetentionSchedule" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRetentionScheduleRequestResult := dao.GetRetentionSchedule( uint64(createRetentionScheduleObj.ID) )
	
	if getRetentionScheduleRequestResult.Success == false {
		t.Errorf(getRetentionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Get RetentionSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRetentionScheduleObj,_ := getRetentionScheduleRequestResult.Data. (model.RetentionSchedule)
	compareRetentionSchedule := cmp.Equal(createRetentionScheduleObj.ID, getRetentionScheduleObj.ID)
	
	if  compareRetentionSchedule == false	{
		t.Errorf( "Created RetentionSchedule object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRetentionScheduleRequestResult := dao.GetAllRetentionSchedule()

	if getAllRetentionScheduleRequestResult.Success == false {
			t.Errorf(getAllRetentionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RetentionSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRetentionScheduleObj []model.RetentionSchedule = getAllRetentionScheduleRequestResult.Data. ([]model.RetentionSchedule)
		
	equalRetentionSchedule := cmp.Equal(createRetentionScheduleObj.ID, getAllRetentionScheduleObj[len(getAllRetentionScheduleObj)-1].ID)
		
	if equalRetentionSchedule == false {
		t.Errorf( "Created object is not equal to the last entry in RetentionSchedule[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RetentionSchedule
	// --------------------------------------------------------------	
	deleteRetentionScheduleRequestResult := dao.DeleteRetentionSchedule(uint64(createRetentionScheduleObj.ID))

	if deleteRetentionScheduleRequestResult.Success == false {
			t.Errorf(deleteRetentionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RetentionSchedule success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRetentionScheduleRequestResult = dao.GetRetentionSchedule( uint64(createRetentionScheduleObj.ID) )
	
	if getRetentionScheduleRequestResult.Success == true {
		t.Errorf(getRetentionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDispositionReviewCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DispositionReview
	//----------------------------------------------------------------------------
	DispositionReviewObj := model.DispositionReview                                                                                                                                    {ReviewDate:time.Now(),Reviewer:"test value for Reviewer",Notes:"test value for Notes",Outcome:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDispositionReviewRequestResult := dao.CreateDispositionReview( DispositionReviewObj )
	
	if createDispositionReviewRequestResult.Success == false {
		t.Errorf(createDispositionReviewRequestResult.Msg)
	} else {
		fmt.Println("Check Create DispositionReview success...")
	}
	
	createDispositionReviewObj,_ := createDispositionReviewRequestResult.Data. (model.DispositionReview)

	// --------------------------------------------------------------
	// Check DispositionReview Obj ID
	// --------------------------------------------------------------	
	if createDispositionReviewObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DispositionReview" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDispositionReviewRequestResult := dao.GetDispositionReview( uint64(createDispositionReviewObj.ID) )
	
	if getDispositionReviewRequestResult.Success == false {
		t.Errorf(getDispositionReviewRequestResult.Msg)
	} else {
		fmt.Println("Check Get DispositionReview success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDispositionReviewObj,_ := getDispositionReviewRequestResult.Data. (model.DispositionReview)
	compareDispositionReview := cmp.Equal(createDispositionReviewObj.ID, getDispositionReviewObj.ID)
	
	if  compareDispositionReview == false	{
		t.Errorf( "Created DispositionReview object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDispositionReviewRequestResult := dao.GetAllDispositionReview()

	if getAllDispositionReviewRequestResult.Success == false {
			t.Errorf(getAllDispositionReviewRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DispositionReview success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDispositionReviewObj []model.DispositionReview = getAllDispositionReviewRequestResult.Data. ([]model.DispositionReview)
		
	equalDispositionReview := cmp.Equal(createDispositionReviewObj.ID, getAllDispositionReviewObj[len(getAllDispositionReviewObj)-1].ID)
		
	if equalDispositionReview == false {
		t.Errorf( "Created object is not equal to the last entry in DispositionReview[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DispositionReview
	// --------------------------------------------------------------	
	deleteDispositionReviewRequestResult := dao.DeleteDispositionReview(uint64(createDispositionReviewObj.ID))

	if deleteDispositionReviewRequestResult.Success == false {
			t.Errorf(deleteDispositionReviewRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DispositionReview success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDispositionReviewRequestResult = dao.GetDispositionReview( uint64(createDispositionReviewObj.ID) )
	
	if getDispositionReviewRequestResult.Success == true {
		t.Errorf(getDispositionReviewRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLegalHoldCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LegalHold
	//----------------------------------------------------------------------------
	LegalHoldObj := model.LegalHold                                                                                                                                                                                            {Name:"test value for Name",Reason:"test value for Reason",IssuedDate:time.Now(),ReleaseDate:time.Now(),HoldStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLegalHoldRequestResult := dao.CreateLegalHold( LegalHoldObj )
	
	if createLegalHoldRequestResult.Success == false {
		t.Errorf(createLegalHoldRequestResult.Msg)
	} else {
		fmt.Println("Check Create LegalHold success...")
	}
	
	createLegalHoldObj,_ := createLegalHoldRequestResult.Data. (model.LegalHold)

	// --------------------------------------------------------------
	// Check LegalHold Obj ID
	// --------------------------------------------------------------	
	if createLegalHoldObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LegalHold" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLegalHoldRequestResult := dao.GetLegalHold( uint64(createLegalHoldObj.ID) )
	
	if getLegalHoldRequestResult.Success == false {
		t.Errorf(getLegalHoldRequestResult.Msg)
	} else {
		fmt.Println("Check Get LegalHold success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLegalHoldObj,_ := getLegalHoldRequestResult.Data. (model.LegalHold)
	compareLegalHold := cmp.Equal(createLegalHoldObj.ID, getLegalHoldObj.ID)
	
	if  compareLegalHold == false	{
		t.Errorf( "Created LegalHold object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLegalHoldRequestResult := dao.GetAllLegalHold()

	if getAllLegalHoldRequestResult.Success == false {
			t.Errorf(getAllLegalHoldRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LegalHold success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLegalHoldObj []model.LegalHold = getAllLegalHoldRequestResult.Data. ([]model.LegalHold)
		
	equalLegalHold := cmp.Equal(createLegalHoldObj.ID, getAllLegalHoldObj[len(getAllLegalHoldObj)-1].ID)
		
	if equalLegalHold == false {
		t.Errorf( "Created object is not equal to the last entry in LegalHold[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LegalHold
	// --------------------------------------------------------------	
	deleteLegalHoldRequestResult := dao.DeleteLegalHold(uint64(createLegalHoldObj.ID))

	if deleteLegalHoldRequestResult.Success == false {
			t.Errorf(deleteLegalHoldRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LegalHold success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLegalHoldRequestResult = dao.GetLegalHold( uint64(createLegalHoldObj.ID) )
	
	if getLegalHoldRequestResult.Success == true {
		t.Errorf(getLegalHoldRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMatterCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Matter
	//----------------------------------------------------------------------------
	MatterObj := model.Matter                                                                                            {MatterName:"test value for MatterName",LeadCounsel:"test value for LeadCounsel",MatterType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMatterRequestResult := dao.CreateMatter( MatterObj )
	
	if createMatterRequestResult.Success == false {
		t.Errorf(createMatterRequestResult.Msg)
	} else {
		fmt.Println("Check Create Matter success...")
	}
	
	createMatterObj,_ := createMatterRequestResult.Data. (model.Matter)

	// --------------------------------------------------------------
	// Check Matter Obj ID
	// --------------------------------------------------------------	
	if createMatterObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Matter" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMatterRequestResult := dao.GetMatter( uint64(createMatterObj.ID) )
	
	if getMatterRequestResult.Success == false {
		t.Errorf(getMatterRequestResult.Msg)
	} else {
		fmt.Println("Check Get Matter success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMatterObj,_ := getMatterRequestResult.Data. (model.Matter)
	compareMatter := cmp.Equal(createMatterObj.ID, getMatterObj.ID)
	
	if  compareMatter == false	{
		t.Errorf( "Created Matter object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMatterRequestResult := dao.GetAllMatter()

	if getAllMatterRequestResult.Success == false {
			t.Errorf(getAllMatterRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Matter success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMatterObj []model.Matter = getAllMatterRequestResult.Data. ([]model.Matter)
		
	equalMatter := cmp.Equal(createMatterObj.ID, getAllMatterObj[len(getAllMatterObj)-1].ID)
		
	if equalMatter == false {
		t.Errorf( "Created object is not equal to the last entry in Matter[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Matter
	// --------------------------------------------------------------	
	deleteMatterRequestResult := dao.DeleteMatter(uint64(createMatterObj.ID))

	if deleteMatterRequestResult.Success == false {
			t.Errorf(deleteMatterRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Matter success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMatterRequestResult = dao.GetMatter( uint64(createMatterObj.ID) )
	
	if getMatterRequestResult.Success == true {
		t.Errorf(getMatterRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestThirdPartyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ThirdParty
	//----------------------------------------------------------------------------
	ThirdPartyObj := model.ThirdParty                                                                                                            {Name:"test value for Name",Country:"test value for Country",ContactEmail:new EmailAddress(),ThirdPartyType:0,Criticality:0}

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


func TestThirdPartyAssessmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ThirdPartyAssessment
	//----------------------------------------------------------------------------
	ThirdPartyAssessmentObj := model.ThirdPartyAssessment                                                                                                                    {AssessmentDate:time.Now(),Assessor:"test value for Assessor",AssessmentType:0,Result:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createThirdPartyAssessmentRequestResult := dao.CreateThirdPartyAssessment( ThirdPartyAssessmentObj )
	
	if createThirdPartyAssessmentRequestResult.Success == false {
		t.Errorf(createThirdPartyAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create ThirdPartyAssessment success...")
	}
	
	createThirdPartyAssessmentObj,_ := createThirdPartyAssessmentRequestResult.Data. (model.ThirdPartyAssessment)

	// --------------------------------------------------------------
	// Check ThirdPartyAssessment Obj ID
	// --------------------------------------------------------------	
	if createThirdPartyAssessmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ThirdPartyAssessment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getThirdPartyAssessmentRequestResult := dao.GetThirdPartyAssessment( uint64(createThirdPartyAssessmentObj.ID) )
	
	if getThirdPartyAssessmentRequestResult.Success == false {
		t.Errorf(getThirdPartyAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get ThirdPartyAssessment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getThirdPartyAssessmentObj,_ := getThirdPartyAssessmentRequestResult.Data. (model.ThirdPartyAssessment)
	compareThirdPartyAssessment := cmp.Equal(createThirdPartyAssessmentObj.ID, getThirdPartyAssessmentObj.ID)
	
	if  compareThirdPartyAssessment == false	{
		t.Errorf( "Created ThirdPartyAssessment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllThirdPartyAssessmentRequestResult := dao.GetAllThirdPartyAssessment()

	if getAllThirdPartyAssessmentRequestResult.Success == false {
			t.Errorf(getAllThirdPartyAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ThirdPartyAssessment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllThirdPartyAssessmentObj []model.ThirdPartyAssessment = getAllThirdPartyAssessmentRequestResult.Data. ([]model.ThirdPartyAssessment)
		
	equalThirdPartyAssessment := cmp.Equal(createThirdPartyAssessmentObj.ID, getAllThirdPartyAssessmentObj[len(getAllThirdPartyAssessmentObj)-1].ID)
		
	if equalThirdPartyAssessment == false {
		t.Errorf( "Created object is not equal to the last entry in ThirdPartyAssessment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ThirdPartyAssessment
	// --------------------------------------------------------------	
	deleteThirdPartyAssessmentRequestResult := dao.DeleteThirdPartyAssessment(uint64(createThirdPartyAssessmentObj.ID))

	if deleteThirdPartyAssessmentRequestResult.Success == false {
			t.Errorf(deleteThirdPartyAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ThirdPartyAssessment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getThirdPartyAssessmentRequestResult = dao.GetThirdPartyAssessment( uint64(createThirdPartyAssessmentObj.ID) )
	
	if getThirdPartyAssessmentRequestResult.Success == true {
		t.Errorf(getThirdPartyAssessmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestContractCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Contract
	//----------------------------------------------------------------------------
	ContractObj := model.Contract                                                                                                                                                                            {Title:"test value for Title",EffectiveDate:time.Now(),ExpiryDate:time.Now(),RepositoryUrl:new URL(),Status:0}

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


func TestException_CRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Exception_
	//----------------------------------------------------------------------------
	Exception_Obj := model.Exception_                                                                                                                                                                                                            {Title:"test value for Title",Justification:"test value for Justification",StartDate:time.Now(),EndDate:time.Now(),ExceptionType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createException_RequestResult := dao.CreateException_( Exception_Obj )
	
	if createException_RequestResult.Success == false {
		t.Errorf(createException_RequestResult.Msg)
	} else {
		fmt.Println("Check Create Exception_ success...")
	}
	
	createException_Obj,_ := createException_RequestResult.Data. (model.Exception_)

	// --------------------------------------------------------------
	// Check Exception_ Obj ID
	// --------------------------------------------------------------	
	if createException_Obj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Exception_" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getException_RequestResult := dao.GetException_( uint64(createException_Obj.ID) )
	
	if getException_RequestResult.Success == false {
		t.Errorf(getException_RequestResult.Msg)
	} else {
		fmt.Println("Check Get Exception_ success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getException_Obj,_ := getException_RequestResult.Data. (model.Exception_)
	compareException_ := cmp.Equal(createException_Obj.ID, getException_Obj.ID)
	
	if  compareException_ == false	{
		t.Errorf( "Created Exception_ object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllException_RequestResult := dao.GetAllException_()

	if getAllException_RequestResult.Success == false {
			t.Errorf(getAllException_RequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Exception_ success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllException_Obj []model.Exception_ = getAllException_RequestResult.Data. ([]model.Exception_)
		
	equalException_ := cmp.Equal(createException_Obj.ID, getAllException_Obj[len(getAllException_Obj)-1].ID)
		
	if equalException_ == false {
		t.Errorf( "Created object is not equal to the last entry in Exception_[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Exception_
	// --------------------------------------------------------------	
	deleteException_RequestResult := dao.DeleteException_(uint64(createException_Obj.ID))

	if deleteException_RequestResult.Success == false {
			t.Errorf(deleteException_RequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Exception_ success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getException_RequestResult = dao.GetException_( uint64(createException_Obj.ID) )
	
	if getException_RequestResult.Success == true {
		t.Errorf(getException_RequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestConsentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Consent
	//----------------------------------------------------------------------------
	ConsentObj := model.Consent                                                                                                                                                                            {SubjectIdentifier:"test value for SubjectIdentifier",CaptureDate:time.Now(),ExpiryDate:time.Now(),ConsentType:0,Status:0}

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


func TestDataBreachCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DataBreach
	//----------------------------------------------------------------------------
	DataBreachObj := model.DataBreach                                                                                                                                                                                    {IncidentDate:time.Now(),Description:"test value for Description",RecordsAffected:100,NotificationRequired:true,Severity:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDataBreachRequestResult := dao.CreateDataBreach( DataBreachObj )
	
	if createDataBreachRequestResult.Success == false {
		t.Errorf(createDataBreachRequestResult.Msg)
	} else {
		fmt.Println("Check Create DataBreach success...")
	}
	
	createDataBreachObj,_ := createDataBreachRequestResult.Data. (model.DataBreach)

	// --------------------------------------------------------------
	// Check DataBreach Obj ID
	// --------------------------------------------------------------	
	if createDataBreachObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DataBreach" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDataBreachRequestResult := dao.GetDataBreach( uint64(createDataBreachObj.ID) )
	
	if getDataBreachRequestResult.Success == false {
		t.Errorf(getDataBreachRequestResult.Msg)
	} else {
		fmt.Println("Check Get DataBreach success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDataBreachObj,_ := getDataBreachRequestResult.Data. (model.DataBreach)
	compareDataBreach := cmp.Equal(createDataBreachObj.ID, getDataBreachObj.ID)
	
	if  compareDataBreach == false	{
		t.Errorf( "Created DataBreach object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDataBreachRequestResult := dao.GetAllDataBreach()

	if getAllDataBreachRequestResult.Success == false {
			t.Errorf(getAllDataBreachRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DataBreach success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDataBreachObj []model.DataBreach = getAllDataBreachRequestResult.Data. ([]model.DataBreach)
		
	equalDataBreach := cmp.Equal(createDataBreachObj.ID, getAllDataBreachObj[len(getAllDataBreachObj)-1].ID)
		
	if equalDataBreach == false {
		t.Errorf( "Created object is not equal to the last entry in DataBreach[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DataBreach
	// --------------------------------------------------------------	
	deleteDataBreachRequestResult := dao.DeleteDataBreach(uint64(createDataBreachObj.ID))

	if deleteDataBreachRequestResult.Success == false {
			t.Errorf(deleteDataBreachRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DataBreach success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDataBreachRequestResult = dao.GetDataBreach( uint64(createDataBreachObj.ID) )
	
	if getDataBreachRequestResult.Success == true {
		t.Errorf(getDataBreachRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

