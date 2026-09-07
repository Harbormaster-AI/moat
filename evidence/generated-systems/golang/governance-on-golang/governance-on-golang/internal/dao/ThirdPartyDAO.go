package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
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
// assigns a Organization on a ThirdParty
//----------------------------------------------------------------------------
func AssignOrganizationToThirdParty( thirdPartyId uint64, organizationId uint64 )(utils.RequestResult){

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
			// assign the Organization	to the ThirdParty
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the ThirdParty
			//----------------------------------------------------------------------------
			return UpdateThirdParty(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a ThirdParty
//----------------------------------------------------------------------------
func UnassignOrganizationFromThirdParty(thirdPartyId uint64)(utils.RequestResult) {

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

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the ThirdParty
		//----------------------------------------------------------------------------
		return UpdateThirdParty(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more processingActivitiesIds as a ProcessingActivities to a ThirdParty
//----------------------------------------------------------------------------
func AddProcessingActivitiesToThirdParty ( thirdPartyId uint64, processingActivitiesIds string )(utils.RequestResult) {

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
		ids := strings.Split( processingActivitiesIds, ",")

		for _, processingActivitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataProcessingActivity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataProcessingActivity
			// with a matching processingActivitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , processingActivitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ProcessingActivities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProcessingActivities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProcessingActivities", processingActivitiesId )
				return utils.RequestResult{false, msg, "unassignProcessingActivities", childObj}
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
// removes one or more processingActivitiesIds as a ProcessingActivities from a ThirdParty
//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromThirdParty( thirdPartyId uint64, processingActivitiesIds string )(utils.RequestResult) {
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
		ids := strings.Split( processingActivitiesIds, ",")

		for _, processingActivitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataProcessingActivity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataProcessingActivity
			// with a matching processingActivitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , processingActivitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataProcessingActivityObj from the ProcessingActivities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProcessingActivities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProcessingActivities", processingActivitiesId )
				return utils.RequestResult{false, msg, "removeProcessingActivities", childObj}
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
// adds one or more assessmentsIds as a Assessments to a ThirdParty
//----------------------------------------------------------------------------
func AddAssessmentsToThirdParty ( thirdPartyId uint64, assessmentsIds string )(utils.RequestResult) {

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
		ids := strings.Split( assessmentsIds, ",")

		for _, assessmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ThirdPartyAssessment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdPartyAssessment
			// with a matching assessmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assessmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Assessments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assessments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assessments", assessmentsId )
				return utils.RequestResult{false, msg, "unassignAssessments", childObj}
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
// removes one or more assessmentsIds as a Assessments from a ThirdParty
//----------------------------------------------------------------------------
func RemoveAssessmentsFromThirdParty( thirdPartyId uint64, assessmentsIds string )(utils.RequestResult) {
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
		ids := strings.Split( assessmentsIds, ",")

		for _, assessmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ThirdPartyAssessment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdPartyAssessment
			// with a matching assessmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assessmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ThirdPartyAssessmentObj from the Assessments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assessments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assessments", assessmentsId )
				return utils.RequestResult{false, msg, "removeAssessments", childObj}
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
// adds one or more contractsIds as a Contracts to a ThirdParty
//----------------------------------------------------------------------------
func AddContractsToThirdParty ( thirdPartyId uint64, contractsIds string )(utils.RequestResult) {

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
		// retrieve the modified ThirdParty from the gorm
		//----------------------------------------------------------------------------
		return GetThirdParty(thirdPartyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contractsIds as a Contracts from a ThirdParty
//----------------------------------------------------------------------------
func RemoveContractsFromThirdParty( thirdPartyId uint64, contractsIds string )(utils.RequestResult) {
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
		// retrieve the modified ThirdParty from the gorm
		//----------------------------------------------------------------------------
		return GetThirdParty(thirdPartyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more obligationsIds as a Obligations to a ThirdParty
//----------------------------------------------------------------------------
func AddObligationsToThirdParty ( thirdPartyId uint64, obligationsIds string )(utils.RequestResult) {

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
		// retrieve the modified ThirdParty from the gorm
		//----------------------------------------------------------------------------
		return GetThirdParty(thirdPartyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more obligationsIds as a Obligations from a ThirdParty
//----------------------------------------------------------------------------
func RemoveObligationsFromThirdParty( thirdPartyId uint64, obligationsIds string )(utils.RequestResult) {
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
		// retrieve the modified ThirdParty from the gorm
		//----------------------------------------------------------------------------
		return GetThirdParty(thirdPartyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataBreachesIds as a DataBreaches to a ThirdParty
//----------------------------------------------------------------------------
func AddDataBreachesToThirdParty ( thirdPartyId uint64, dataBreachesIds string )(utils.RequestResult) {

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
		// retrieve the modified ThirdParty from the gorm
		//----------------------------------------------------------------------------
		return GetThirdParty(thirdPartyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataBreachesIds as a DataBreaches from a ThirdParty
//----------------------------------------------------------------------------
func RemoveDataBreachesFromThirdParty( thirdPartyId uint64, dataBreachesIds string )(utils.RequestResult) {
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
		// retrieve the modified ThirdParty from the gorm
		//----------------------------------------------------------------------------
		return GetThirdParty(thirdPartyId)

	} else {
		return parentRequestResult
	}
}

