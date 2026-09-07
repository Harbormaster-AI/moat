package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ObligationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateObligation - creates a new db entry
//----------------------------------------------------------------------------
func CreateObligation(obj model.Obligation)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Obligation with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Obligation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateObligation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetObligation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetObligation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Obligation

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Obligation with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Obligation using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Obligation using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetObligation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllObligation - returns all
//----------------------------------------------------------------------------
func GetAllObligation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Obligation

	//----------------------------------------------------------------------------
	// Request the ORM to find all Obligation
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Obligation" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Obligation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllObligation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateObligation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateObligation(obj model.Obligation)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Obligation using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Obligation using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateObligation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteObligation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteObligation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Obligation with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetObligation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Obligation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Obligation)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Obligation using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Obligation using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteObligation", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Regulation on a Obligation
//----------------------------------------------------------------------------
func AssignRegulationToObligation( obligationId uint64, regulationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Obligation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObligation(obligationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Obligation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Obligation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Regulation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Regulation with a
		// matching regulationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, regulationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Regulation	to the Obligation
			//----------------------------------------------------------------------------
			parentObj.Regulation = &childObj

			//----------------------------------------------------------------------------
			// save the Obligation
			//----------------------------------------------------------------------------
			return UpdateObligation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Regulation", regulationId )
			return utils.RequestResult{false, msg, "assignRegulation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Regulation on a Obligation
//----------------------------------------------------------------------------
func UnassignRegulationFromObligation(obligationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Obligation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObligation(obligationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Obligation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Obligation)

		//----------------------------------------------------------------------------
		// assign an empty Regulation to the Regulation
		//----------------------------------------------------------------------------
		parentObj.Regulation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Regulation
		//----------------------------------------------------------------------------
		parentObj.RegulationId = nil;

		//----------------------------------------------------------------------------
		// save the Obligation
		//----------------------------------------------------------------------------
		return UpdateObligation(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more controlsIds as a Controls to a Obligation
//----------------------------------------------------------------------------
func AddControlsToObligation ( obligationId uint64, controlsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Obligation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObligation(obligationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Obligation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Obligation)

		// slice the ids on comma with no spaces
		ids := strings.Split( controlsIds, ",")

		for _, controlsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Control

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Control
			// with a matching controlsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , controlsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Controls using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Controls").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Controls", controlsId )
				return utils.RequestResult{false, msg, "unassignControls", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Obligation from the gorm
		//----------------------------------------------------------------------------
		return GetObligation(obligationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more controlsIds as a Controls from a Obligation
//----------------------------------------------------------------------------
func RemoveControlsFromObligation( obligationId uint64, controlsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Obligation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObligation(obligationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Obligation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Obligation)

		// slice the ids on comma with no spaces
		ids := strings.Split( controlsIds, ",")

		for _, controlsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Control

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Control
			// with a matching controlsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , controlsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ControlObj from the Controls array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Controls").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Controls", controlsId )
				return utils.RequestResult{false, msg, "removeControls", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Obligation from the gorm
		//----------------------------------------------------------------------------
		return GetObligation(obligationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a Obligation
//----------------------------------------------------------------------------
func AddPoliciesToObligation ( obligationId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Obligation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObligation(obligationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Obligation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Obligation)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Policies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "unassignPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Obligation from the gorm
		//----------------------------------------------------------------------------
		return GetObligation(obligationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a Obligation
//----------------------------------------------------------------------------
func RemovePoliciesFromObligation( obligationId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Obligation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObligation(obligationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Obligation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Obligation)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyObj from the Policies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "removePolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Obligation from the gorm
		//----------------------------------------------------------------------------
		return GetObligation(obligationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more contractsIds as a Contracts to a Obligation
//----------------------------------------------------------------------------
func AddContractsToObligation ( obligationId uint64, contractsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Obligation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObligation(obligationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Obligation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Obligation)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Contracts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "unassignContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Obligation from the gorm
		//----------------------------------------------------------------------------
		return GetObligation(obligationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contractsIds as a Contracts from a Obligation
//----------------------------------------------------------------------------
func RemoveContractsFromObligation( obligationId uint64, contractsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Obligation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetObligation(obligationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Obligation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Obligation)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Contract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Contract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ContractObj from the Contracts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "removeContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Obligation from the gorm
		//----------------------------------------------------------------------------
		return GetObligation(obligationId)

	} else {
		return parentRequestResult
	}
}

