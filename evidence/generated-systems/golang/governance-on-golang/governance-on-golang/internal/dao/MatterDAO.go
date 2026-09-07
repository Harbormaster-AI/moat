package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MatterDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMatter - creates a new db entry
//----------------------------------------------------------------------------
func CreateMatter(obj model.Matter)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Matter with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Matter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMatter", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMatter - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMatter(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Matter

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Matter with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Matter using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Matter using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMatter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMatter - returns all
//----------------------------------------------------------------------------
func GetAllMatter()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Matter

	//----------------------------------------------------------------------------
	// Request the ORM to find all Matter
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Matter" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Matter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMatter", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMatter - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMatter(obj model.Matter)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Matter using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Matter using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMatter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMatter - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMatter(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Matter with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMatter(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Matter)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Matter using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Matter using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMatter", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Matter
//----------------------------------------------------------------------------
func AssignOrganizationToMatter( matterId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Matter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatter(matterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Matter)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the Matter
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Matter
			//----------------------------------------------------------------------------
			return UpdateMatter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Matter
//----------------------------------------------------------------------------
func UnassignOrganizationFromMatter(matterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Matter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatter(matterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Matter)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Matter
		//----------------------------------------------------------------------------
		return UpdateMatter(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more legalHoldsIds as a LegalHolds to a Matter
//----------------------------------------------------------------------------
func AddLegalHoldsToMatter ( matterId uint64, legalHoldsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Matter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatter(matterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Matter)

		// slice the ids on comma with no spaces
		ids := strings.Split( legalHoldsIds, ",")

		for _, legalHoldsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LegalHold

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LegalHold
			// with a matching legalHoldsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , legalHoldsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LegalHolds using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LegalHolds").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LegalHolds", legalHoldsId )
				return utils.RequestResult{false, msg, "unassignLegalHolds", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Matter from the gorm
		//----------------------------------------------------------------------------
		return GetMatter(matterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more legalHoldsIds as a LegalHolds from a Matter
//----------------------------------------------------------------------------
func RemoveLegalHoldsFromMatter( matterId uint64, legalHoldsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Matter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatter(matterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Matter)

		// slice the ids on comma with no spaces
		ids := strings.Split( legalHoldsIds, ",")

		for _, legalHoldsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LegalHold

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LegalHold
			// with a matching legalHoldsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , legalHoldsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LegalHoldObj from the LegalHolds array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LegalHolds").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LegalHolds", legalHoldsId )
				return utils.RequestResult{false, msg, "removeLegalHolds", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Matter from the gorm
		//----------------------------------------------------------------------------
		return GetMatter(matterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataBreachesIds as a DataBreaches to a Matter
//----------------------------------------------------------------------------
func AddDataBreachesToMatter ( matterId uint64, dataBreachesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Matter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatter(matterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Matter)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataBreachesIds, ",")

		for _, dataBreachesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataBreach

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataBreach
			// with a matching dataBreachesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataBreachesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DataBreaches using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataBreaches").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataBreaches", dataBreachesId )
				return utils.RequestResult{false, msg, "unassignDataBreaches", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Matter from the gorm
		//----------------------------------------------------------------------------
		return GetMatter(matterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataBreachesIds as a DataBreaches from a Matter
//----------------------------------------------------------------------------
func RemoveDataBreachesFromMatter( matterId uint64, dataBreachesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Matter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatter(matterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Matter)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataBreachesIds, ",")

		for _, dataBreachesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataBreach

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataBreach
			// with a matching dataBreachesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataBreachesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataBreachObj from the DataBreaches array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataBreaches").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataBreaches", dataBreachesId )
				return utils.RequestResult{false, msg, "removeDataBreaches", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Matter from the gorm
		//----------------------------------------------------------------------------
		return GetMatter(matterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more contractsIds as a Contracts to a Matter
//----------------------------------------------------------------------------
func AddContractsToMatter ( matterId uint64, contractsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Matter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatter(matterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Matter)

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
		// retrieve the modified Matter from the gorm
		//----------------------------------------------------------------------------
		return GetMatter(matterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contractsIds as a Contracts from a Matter
//----------------------------------------------------------------------------
func RemoveContractsFromMatter( matterId uint64, contractsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Matter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMatter(matterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Matter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Matter)

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
		// retrieve the modified Matter from the gorm
		//----------------------------------------------------------------------------
		return GetMatter(matterId)

	} else {
		return parentRequestResult
	}
}

