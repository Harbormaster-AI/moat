package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ContractDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateContract - creates a new db entry
//----------------------------------------------------------------------------
func CreateContract(obj model.Contract)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Contract with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Contract", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateContract", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetContract - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetContract(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Contract

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Contract with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Contract using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Contract using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetContract", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllContract - returns all
//----------------------------------------------------------------------------
func GetAllContract()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Contract

	//----------------------------------------------------------------------------
	// Request the ORM to find all Contract
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Contract" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Contract", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllContract", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateContract - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateContract(obj model.Contract)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Contract using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Contract using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateContract", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteContract - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteContract(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetContract(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Contract using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Contract using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteContract", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ThirdParty on a Contract
//----------------------------------------------------------------------------
func AssignThirdPartyToContract( contractId uint64, thirdPartyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ThirdParty

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ThirdParty with a
		// matching thirdPartyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, thirdPartyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ThirdParty	to the Contract
			//----------------------------------------------------------------------------
			parentObj.ThirdParty = &childObj

			//----------------------------------------------------------------------------
			// save the Contract
			//----------------------------------------------------------------------------
			return UpdateContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdParty", thirdPartyId )
			return utils.RequestResult{false, msg, "assignThirdParty", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ThirdParty on a Contract
//----------------------------------------------------------------------------
func UnassignThirdPartyFromContract(contractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// assign an empty ThirdParty to the ThirdParty
		//----------------------------------------------------------------------------
		parentObj.ThirdParty = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ThirdParty
		//----------------------------------------------------------------------------
		parentObj.ThirdPartyId = nil;

		//----------------------------------------------------------------------------
		// save the Contract
		//----------------------------------------------------------------------------
		return UpdateContract(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Matter on a Contract
//----------------------------------------------------------------------------
func AssignMatterToContract( contractId uint64, matterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Matter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Matter with a
		// matching matterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, matterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Matter	to the Contract
			//----------------------------------------------------------------------------
			parentObj.Matter = &childObj

			//----------------------------------------------------------------------------
			// save the Contract
			//----------------------------------------------------------------------------
			return UpdateContract(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Matter", matterId )
			return utils.RequestResult{false, msg, "assignMatter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Matter on a Contract
//----------------------------------------------------------------------------
func UnassignMatterFromContract(contractId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		//----------------------------------------------------------------------------
		// assign an empty Matter to the Matter
		//----------------------------------------------------------------------------
		parentObj.Matter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Matter
		//----------------------------------------------------------------------------
		parentObj.MatterId = nil;

		//----------------------------------------------------------------------------
		// save the Contract
		//----------------------------------------------------------------------------
		return UpdateContract(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more obligationsIds as a Obligations to a Contract
//----------------------------------------------------------------------------
func AddObligationsToContract ( contractId uint64, obligationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		// slice the ids on comma with no spaces
		ids := strings.Split( obligationsIds, ",")

		for _, obligationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Obligation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Obligation
			// with a matching obligationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , obligationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Obligations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Obligations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Obligations", obligationsId )
				return utils.RequestResult{false, msg, "unassignObligations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contract from the gorm
		//----------------------------------------------------------------------------
		return GetContract(contractId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more obligationsIds as a Obligations from a Contract
//----------------------------------------------------------------------------
func RemoveObligationsFromContract( contractId uint64, obligationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		// slice the ids on comma with no spaces
		ids := strings.Split( obligationsIds, ",")

		for _, obligationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Obligation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Obligation
			// with a matching obligationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , obligationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ObligationObj from the Obligations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Obligations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Obligations", obligationsId )
				return utils.RequestResult{false, msg, "removeObligations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contract from the gorm
		//----------------------------------------------------------------------------
		return GetContract(contractId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataProcessingActivitiesIds as a DataProcessingActivities to a Contract
//----------------------------------------------------------------------------
func AddDataProcessingActivitiesToContract ( contractId uint64, dataProcessingActivitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataProcessingActivitiesIds, ",")

		for _, dataProcessingActivitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataProcessingActivity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataProcessingActivity
			// with a matching dataProcessingActivitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataProcessingActivitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DataProcessingActivities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataProcessingActivities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataProcessingActivities", dataProcessingActivitiesId )
				return utils.RequestResult{false, msg, "unassignDataProcessingActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contract from the gorm
		//----------------------------------------------------------------------------
		return GetContract(contractId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataProcessingActivitiesIds as a DataProcessingActivities from a Contract
//----------------------------------------------------------------------------
func RemoveDataProcessingActivitiesFromContract( contractId uint64, dataProcessingActivitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Contract with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetContract(contractId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Contract so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Contract)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataProcessingActivitiesIds, ",")

		for _, dataProcessingActivitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataProcessingActivity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataProcessingActivity
			// with a matching dataProcessingActivitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataProcessingActivitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataProcessingActivityObj from the DataProcessingActivities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataProcessingActivities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataProcessingActivities", dataProcessingActivitiesId )
				return utils.RequestResult{false, msg, "removeDataProcessingActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Contract from the gorm
		//----------------------------------------------------------------------------
		return GetContract(contractId)

	} else {
		return parentRequestResult
	}
}

