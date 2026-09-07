package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing VerifiedAddressDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateVerifiedAddress - creates a new db entry
//----------------------------------------------------------------------------
func CreateVerifiedAddress(obj model.VerifiedAddress)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a VerifiedAddress with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a VerifiedAddress", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateVerifiedAddress", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetVerifiedAddress - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetVerifiedAddress(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.VerifiedAddress

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a VerifiedAddress with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a VerifiedAddress using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a VerifiedAddress using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetVerifiedAddress", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllVerifiedAddress - returns all
//----------------------------------------------------------------------------
func GetAllVerifiedAddress()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.VerifiedAddress

	//----------------------------------------------------------------------------
	// Request the ORM to find all VerifiedAddress
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all VerifiedAddress" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all VerifiedAddress", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllVerifiedAddress", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateVerifiedAddress - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateVerifiedAddress(obj model.VerifiedAddress)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a VerifiedAddress using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a VerifiedAddress using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateVerifiedAddress", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteVerifiedAddress - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteVerifiedAddress(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the VerifiedAddress with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetVerifiedAddress(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.VerifiedAddress so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.VerifiedAddress)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a VerifiedAddress using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a VerifiedAddress using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteVerifiedAddress", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a KycProfile on a VerifiedAddress
//----------------------------------------------------------------------------
func AssignKycProfileToVerifiedAddress( verifiedAddressId uint64, kycProfileId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the VerifiedAddress with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVerifiedAddress(verifiedAddressId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.VerifiedAddress so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.VerifiedAddress)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.KYCProfile

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a KYCProfile with a
		// matching kycProfileId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, kycProfileId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the KycProfile	to the VerifiedAddress
			//----------------------------------------------------------------------------
			parentObj.KycProfile = &childObj

			//----------------------------------------------------------------------------
			// save the VerifiedAddress
			//----------------------------------------------------------------------------
			return UpdateVerifiedAddress(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "KycProfile", kycProfileId )
			return utils.RequestResult{false, msg, "assignKycProfile", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a KycProfile on a VerifiedAddress
//----------------------------------------------------------------------------
func UnassignKycProfileFromVerifiedAddress(verifiedAddressId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the VerifiedAddress with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVerifiedAddress(verifiedAddressId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.VerifiedAddress so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.VerifiedAddress)

		//----------------------------------------------------------------------------
		// assign an empty KYCProfile to the KycProfile
		//----------------------------------------------------------------------------
		parentObj.KycProfile = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the KycProfile
		//----------------------------------------------------------------------------
		parentObj.KycProfileId = nil;

		//----------------------------------------------------------------------------
		// save the VerifiedAddress
		//----------------------------------------------------------------------------
		return UpdateVerifiedAddress(parentObj)

	} else {
		return parentRequestResult
	}

}


