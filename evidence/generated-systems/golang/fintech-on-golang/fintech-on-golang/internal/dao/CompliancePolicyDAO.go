package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CompliancePolicyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCompliancePolicy - creates a new db entry
//----------------------------------------------------------------------------
func CreateCompliancePolicy(obj model.CompliancePolicy)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var createMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a CompliancePolicy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CompliancePolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCompliancePolicy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCompliancePolicy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCompliancePolicy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CompliancePolicy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CompliancePolicy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CompliancePolicy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CompliancePolicy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCompliancePolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCompliancePolicy - returns all
//----------------------------------------------------------------------------
func GetAllCompliancePolicy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CompliancePolicy

	//----------------------------------------------------------------------------
	// Request the ORM to find all CompliancePolicy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CompliancePolicy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CompliancePolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCompliancePolicy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCompliancePolicy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCompliancePolicy(obj model.CompliancePolicy)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a CompliancePolicy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CompliancePolicy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCompliancePolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCompliancePolicy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCompliancePolicy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CompliancePolicy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCompliancePolicy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompliancePolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CompliancePolicy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CompliancePolicy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CompliancePolicy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCompliancePolicy", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Institution on a CompliancePolicy
//----------------------------------------------------------------------------
func AssignInstitutionToCompliancePolicy( compliancePolicyId uint64, institutionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CompliancePolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompliancePolicy(compliancePolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompliancePolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompliancePolicy)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.FinancialInstitution

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FinancialInstitution with a
		// matching institutionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, institutionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Institution	to the CompliancePolicy
			//----------------------------------------------------------------------------
			parentObj.Institution = &childObj

			//----------------------------------------------------------------------------
			// save the CompliancePolicy
			//----------------------------------------------------------------------------
			return UpdateCompliancePolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Institution", institutionId )
			return utils.RequestResult{false, msg, "assignInstitution", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Institution on a CompliancePolicy
//----------------------------------------------------------------------------
func UnassignInstitutionFromCompliancePolicy(compliancePolicyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CompliancePolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCompliancePolicy(compliancePolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CompliancePolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CompliancePolicy)

		//----------------------------------------------------------------------------
		// assign an empty FinancialInstitution to the Institution
		//----------------------------------------------------------------------------
		parentObj.Institution = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Institution
		//----------------------------------------------------------------------------
		parentObj.InstitutionId = nil;

		//----------------------------------------------------------------------------
		// save the CompliancePolicy
		//----------------------------------------------------------------------------
		return UpdateCompliancePolicy(parentObj)

	} else {
		return parentRequestResult
	}

}


