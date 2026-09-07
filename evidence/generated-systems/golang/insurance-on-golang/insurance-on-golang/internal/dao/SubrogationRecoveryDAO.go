package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SubrogationRecoveryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSubrogationRecovery - creates a new db entry
//----------------------------------------------------------------------------
func CreateSubrogationRecovery(obj model.SubrogationRecovery)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SubrogationRecovery with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SubrogationRecovery", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSubrogationRecovery", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSubrogationRecovery - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSubrogationRecovery(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SubrogationRecovery

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SubrogationRecovery with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SubrogationRecovery using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SubrogationRecovery using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSubrogationRecovery", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSubrogationRecovery - returns all
//----------------------------------------------------------------------------
func GetAllSubrogationRecovery()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SubrogationRecovery

	//----------------------------------------------------------------------------
	// Request the ORM to find all SubrogationRecovery
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SubrogationRecovery" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SubrogationRecovery", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSubrogationRecovery", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSubrogationRecovery - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSubrogationRecovery(obj model.SubrogationRecovery)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SubrogationRecovery using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SubrogationRecovery using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSubrogationRecovery", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSubrogationRecovery - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSubrogationRecovery(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SubrogationRecovery with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSubrogationRecovery(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SubrogationRecovery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SubrogationRecovery)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SubrogationRecovery using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SubrogationRecovery using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSubrogationRecovery", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Claim on a SubrogationRecovery
//----------------------------------------------------------------------------
func AssignClaimToSubrogationRecovery( subrogationRecoveryId uint64, claimId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SubrogationRecovery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSubrogationRecovery(subrogationRecoveryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SubrogationRecovery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SubrogationRecovery)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Claim

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Claim with a
		// matching claimId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, claimId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Claim	to the SubrogationRecovery
			//----------------------------------------------------------------------------
			parentObj.Claim = &childObj

			//----------------------------------------------------------------------------
			// save the SubrogationRecovery
			//----------------------------------------------------------------------------
			return UpdateSubrogationRecovery(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claim", claimId )
			return utils.RequestResult{false, msg, "assignClaim", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Claim on a SubrogationRecovery
//----------------------------------------------------------------------------
func UnassignClaimFromSubrogationRecovery(subrogationRecoveryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SubrogationRecovery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSubrogationRecovery(subrogationRecoveryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SubrogationRecovery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SubrogationRecovery)

		//----------------------------------------------------------------------------
		// assign an empty Claim to the Claim
		//----------------------------------------------------------------------------
		parentObj.Claim = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Claim
		//----------------------------------------------------------------------------
		parentObj.ClaimId = nil;

		//----------------------------------------------------------------------------
		// save the SubrogationRecovery
		//----------------------------------------------------------------------------
		return UpdateSubrogationRecovery(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Exposure on a SubrogationRecovery
//----------------------------------------------------------------------------
func AssignExposureToSubrogationRecovery( subrogationRecoveryId uint64, exposureId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SubrogationRecovery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSubrogationRecovery(subrogationRecoveryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SubrogationRecovery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SubrogationRecovery)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Exposure

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Exposure with a
		// matching exposureId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, exposureId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Exposure	to the SubrogationRecovery
			//----------------------------------------------------------------------------
			parentObj.Exposure = &childObj

			//----------------------------------------------------------------------------
			// save the SubrogationRecovery
			//----------------------------------------------------------------------------
			return UpdateSubrogationRecovery(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exposure", exposureId )
			return utils.RequestResult{false, msg, "assignExposure", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Exposure on a SubrogationRecovery
//----------------------------------------------------------------------------
func UnassignExposureFromSubrogationRecovery(subrogationRecoveryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SubrogationRecovery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSubrogationRecovery(subrogationRecoveryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SubrogationRecovery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SubrogationRecovery)

		//----------------------------------------------------------------------------
		// assign an empty Exposure to the Exposure
		//----------------------------------------------------------------------------
		parentObj.Exposure = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Exposure
		//----------------------------------------------------------------------------
		parentObj.ExposureId = nil;

		//----------------------------------------------------------------------------
		// save the SubrogationRecovery
		//----------------------------------------------------------------------------
		return UpdateSubrogationRecovery(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Counterparty on a SubrogationRecovery
//----------------------------------------------------------------------------
func AssignCounterpartyToSubrogationRecovery( subrogationRecoveryId uint64, counterpartyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SubrogationRecovery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSubrogationRecovery(subrogationRecoveryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SubrogationRecovery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SubrogationRecovery)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ThirdParty

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ThirdParty with a
		// matching counterpartyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, counterpartyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Counterparty	to the SubrogationRecovery
			//----------------------------------------------------------------------------
			parentObj.Counterparty = &childObj

			//----------------------------------------------------------------------------
			// save the SubrogationRecovery
			//----------------------------------------------------------------------------
			return UpdateSubrogationRecovery(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Counterparty", counterpartyId )
			return utils.RequestResult{false, msg, "assignCounterparty", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Counterparty on a SubrogationRecovery
//----------------------------------------------------------------------------
func UnassignCounterpartyFromSubrogationRecovery(subrogationRecoveryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SubrogationRecovery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSubrogationRecovery(subrogationRecoveryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SubrogationRecovery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SubrogationRecovery)

		//----------------------------------------------------------------------------
		// assign an empty ThirdParty to the Counterparty
		//----------------------------------------------------------------------------
		parentObj.Counterparty = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Counterparty
		//----------------------------------------------------------------------------
		parentObj.CounterpartyId = nil;

		//----------------------------------------------------------------------------
		// save the SubrogationRecovery
		//----------------------------------------------------------------------------
		return UpdateSubrogationRecovery(parentObj)

	} else {
		return parentRequestResult
	}

}


