package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ThirdPartyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateThirdParty - creates a new db entry
//----------------------------------------------------------------------------
func CreateThirdParty(obj model.ThirdParty)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ThirdParty with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ThirdParty", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateThirdParty", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetThirdParty - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetThirdParty(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ThirdParty

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ThirdParty with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ThirdParty using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ThirdParty using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetThirdParty", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllThirdParty - returns all
//----------------------------------------------------------------------------
func GetAllThirdParty()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ThirdParty

	//----------------------------------------------------------------------------
	// Request the ORM to find all ThirdParty
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ThirdParty" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ThirdParty", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllThirdParty", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateThirdParty - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateThirdParty(obj model.ThirdParty)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ThirdParty using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ThirdParty using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateThirdParty", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteThirdParty - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteThirdParty(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ThirdParty with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetThirdParty(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdParty so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ThirdParty)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ThirdParty using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ThirdParty using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteThirdParty", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more subrogationsIds as a Subrogations to a ThirdParty
//----------------------------------------------------------------------------
func AddSubrogationsToThirdParty ( thirdPartyId uint64, subrogationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ThirdParty with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdParty(thirdPartyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdParty so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdParty)

		// slice the ids on comma with no spaces
		ids := strings.Split( subrogationsIds, ",")

		for _, subrogationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SubrogationRecovery

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SubrogationRecovery
			// with a matching subrogationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , subrogationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Subrogations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Subrogations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Subrogations", subrogationsId )
				return utils.RequestResult{false, msg, "unassignSubrogations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ThirdParty from the gorm
		//----------------------------------------------------------------------------
		return GetThirdParty(thirdPartyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more subrogationsIds as a Subrogations from a ThirdParty
//----------------------------------------------------------------------------
func RemoveSubrogationsFromThirdParty( thirdPartyId uint64, subrogationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ThirdParty with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdParty(thirdPartyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdParty so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdParty)

		// slice the ids on comma with no spaces
		ids := strings.Split( subrogationsIds, ",")

		for _, subrogationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SubrogationRecovery

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SubrogationRecovery
			// with a matching subrogationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , subrogationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SubrogationRecoveryObj from the Subrogations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Subrogations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Subrogations", subrogationsId )
				return utils.RequestResult{false, msg, "removeSubrogations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ThirdParty from the gorm
		//----------------------------------------------------------------------------
		return GetThirdParty(thirdPartyId)

	} else {
		return parentRequestResult
	}
}

