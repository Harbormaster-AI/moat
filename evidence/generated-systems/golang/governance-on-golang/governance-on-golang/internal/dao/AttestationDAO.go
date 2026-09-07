package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AttestationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAttestation - creates a new db entry
//----------------------------------------------------------------------------
func CreateAttestation(obj model.Attestation)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Attestation with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Attestation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAttestation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAttestation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAttestation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Attestation

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Attestation with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Attestation using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Attestation using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAttestation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAttestation - returns all
//----------------------------------------------------------------------------
func GetAllAttestation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Attestation

	//----------------------------------------------------------------------------
	// Request the ORM to find all Attestation
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Attestation" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Attestation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAttestation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAttestation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAttestation(obj model.Attestation)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Attestation using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Attestation using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAttestation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAttestation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAttestation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Attestation with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAttestation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Attestation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Attestation)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Attestation using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Attestation using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAttestation", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Control on a Attestation
//----------------------------------------------------------------------------
func AssignControlToAttestation( attestationId uint64, controlId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Attestation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAttestation(attestationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Attestation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Attestation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Control

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Control with a
		// matching controlId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, controlId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Control	to the Attestation
			//----------------------------------------------------------------------------
			parentObj.Control = &childObj

			//----------------------------------------------------------------------------
			// save the Attestation
			//----------------------------------------------------------------------------
			return UpdateAttestation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Control", controlId )
			return utils.RequestResult{false, msg, "assignControl", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Control on a Attestation
//----------------------------------------------------------------------------
func UnassignControlFromAttestation(attestationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Attestation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAttestation(attestationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Attestation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Attestation)

		//----------------------------------------------------------------------------
		// assign an empty Control to the Control
		//----------------------------------------------------------------------------
		parentObj.Control = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Control
		//----------------------------------------------------------------------------
		parentObj.ControlId = nil;

		//----------------------------------------------------------------------------
		// save the Attestation
		//----------------------------------------------------------------------------
		return UpdateAttestation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Policy on a Attestation
//----------------------------------------------------------------------------
func AssignPolicyToAttestation( attestationId uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Attestation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAttestation(attestationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Attestation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Attestation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Policy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Policy with a
		// matching policyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, policyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Policy	to the Attestation
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the Attestation
			//----------------------------------------------------------------------------
			return UpdateAttestation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a Attestation
//----------------------------------------------------------------------------
func UnassignPolicyFromAttestation(attestationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Attestation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAttestation(attestationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Attestation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Attestation)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the Attestation
		//----------------------------------------------------------------------------
		return UpdateAttestation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ComplianceProgram on a Attestation
//----------------------------------------------------------------------------
func AssignComplianceProgramToAttestation( attestationId uint64, complianceProgramId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Attestation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAttestation(attestationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Attestation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Attestation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ComplianceProgram

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ComplianceProgram with a
		// matching complianceProgramId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, complianceProgramId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ComplianceProgram	to the Attestation
			//----------------------------------------------------------------------------
			parentObj.ComplianceProgram = &childObj

			//----------------------------------------------------------------------------
			// save the Attestation
			//----------------------------------------------------------------------------
			return UpdateAttestation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ComplianceProgram", complianceProgramId )
			return utils.RequestResult{false, msg, "assignComplianceProgram", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ComplianceProgram on a Attestation
//----------------------------------------------------------------------------
func UnassignComplianceProgramFromAttestation(attestationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Attestation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAttestation(attestationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Attestation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Attestation)

		//----------------------------------------------------------------------------
		// assign an empty ComplianceProgram to the ComplianceProgram
		//----------------------------------------------------------------------------
		parentObj.ComplianceProgram = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ComplianceProgram
		//----------------------------------------------------------------------------
		parentObj.ComplianceProgramId = nil;

		//----------------------------------------------------------------------------
		// save the Attestation
		//----------------------------------------------------------------------------
		return UpdateAttestation(parentObj)

	} else {
		return parentRequestResult
	}

}


