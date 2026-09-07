package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ClaimReserveDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateClaimReserve - creates a new db entry
//----------------------------------------------------------------------------
func CreateClaimReserve(obj model.ClaimReserve)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ClaimReserve with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ClaimReserve", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateClaimReserve", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetClaimReserve - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetClaimReserve(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ClaimReserve

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ClaimReserve with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ClaimReserve using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ClaimReserve using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetClaimReserve", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllClaimReserve - returns all
//----------------------------------------------------------------------------
func GetAllClaimReserve()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ClaimReserve

	//----------------------------------------------------------------------------
	// Request the ORM to find all ClaimReserve
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ClaimReserve" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ClaimReserve", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllClaimReserve", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateClaimReserve - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateClaimReserve(obj model.ClaimReserve)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ClaimReserve using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ClaimReserve using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateClaimReserve", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteClaimReserve - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteClaimReserve(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ClaimReserve with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetClaimReserve(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimReserve so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ClaimReserve)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ClaimReserve using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ClaimReserve using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteClaimReserve", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Claim on a ClaimReserve
//----------------------------------------------------------------------------
func AssignClaimToClaimReserve( claimReserveId uint64, claimId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClaimReserve with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimReserve(claimReserveId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimReserve so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimReserve)

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
			// assign the Claim	to the ClaimReserve
			//----------------------------------------------------------------------------
			parentObj.Claim = &childObj

			//----------------------------------------------------------------------------
			// save the ClaimReserve
			//----------------------------------------------------------------------------
			return UpdateClaimReserve(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claim", claimId )
			return utils.RequestResult{false, msg, "assignClaim", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Claim on a ClaimReserve
//----------------------------------------------------------------------------
func UnassignClaimFromClaimReserve(claimReserveId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClaimReserve with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimReserve(claimReserveId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimReserve so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimReserve)

		//----------------------------------------------------------------------------
		// assign an empty Claim to the Claim
		//----------------------------------------------------------------------------
		parentObj.Claim = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Claim
		//----------------------------------------------------------------------------
		parentObj.ClaimId = nil;

		//----------------------------------------------------------------------------
		// save the ClaimReserve
		//----------------------------------------------------------------------------
		return UpdateClaimReserve(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Exposure on a ClaimReserve
//----------------------------------------------------------------------------
func AssignExposureToClaimReserve( claimReserveId uint64, exposureId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClaimReserve with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimReserve(claimReserveId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimReserve so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimReserve)

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
			// assign the Exposure	to the ClaimReserve
			//----------------------------------------------------------------------------
			parentObj.Exposure = &childObj

			//----------------------------------------------------------------------------
			// save the ClaimReserve
			//----------------------------------------------------------------------------
			return UpdateClaimReserve(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exposure", exposureId )
			return utils.RequestResult{false, msg, "assignExposure", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Exposure on a ClaimReserve
//----------------------------------------------------------------------------
func UnassignExposureFromClaimReserve(claimReserveId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClaimReserve with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimReserve(claimReserveId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimReserve so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimReserve)

		//----------------------------------------------------------------------------
		// assign an empty Exposure to the Exposure
		//----------------------------------------------------------------------------
		parentObj.Exposure = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Exposure
		//----------------------------------------------------------------------------
		parentObj.ExposureId = nil;

		//----------------------------------------------------------------------------
		// save the ClaimReserve
		//----------------------------------------------------------------------------
		return UpdateClaimReserve(parentObj)

	} else {
		return parentRequestResult
	}

}


