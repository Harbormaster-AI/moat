package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EndorsementDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEndorsement - creates a new db entry
//----------------------------------------------------------------------------
func CreateEndorsement(obj model.Endorsement)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Endorsement with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Endorsement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEndorsement", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEndorsement - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEndorsement(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Endorsement

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Endorsement with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Endorsement using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Endorsement using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEndorsement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEndorsement - returns all
//----------------------------------------------------------------------------
func GetAllEndorsement()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Endorsement

	//----------------------------------------------------------------------------
	// Request the ORM to find all Endorsement
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Endorsement" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Endorsement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEndorsement", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEndorsement - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEndorsement(obj model.Endorsement)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Endorsement using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Endorsement using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEndorsement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEndorsement - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEndorsement(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Endorsement with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEndorsement(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Endorsement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Endorsement)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Endorsement using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Endorsement using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEndorsement", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Policy on a Endorsement
//----------------------------------------------------------------------------
func AssignPolicyToEndorsement( endorsementId uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Endorsement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEndorsement(endorsementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Endorsement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Endorsement)

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
			// assign the Policy	to the Endorsement
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the Endorsement
			//----------------------------------------------------------------------------
			return UpdateEndorsement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a Endorsement
//----------------------------------------------------------------------------
func UnassignPolicyFromEndorsement(endorsementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Endorsement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEndorsement(endorsementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Endorsement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Endorsement)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the Endorsement
		//----------------------------------------------------------------------------
		return UpdateEndorsement(parentObj)

	} else {
		return parentRequestResult
	}

}


