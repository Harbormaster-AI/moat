package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DataBreachDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDataBreach - creates a new db entry
//----------------------------------------------------------------------------
func CreateDataBreach(obj model.DataBreach)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DataBreach with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DataBreach", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDataBreach", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDataBreach - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDataBreach(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DataBreach

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DataBreach with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DataBreach using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DataBreach using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDataBreach", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDataBreach - returns all
//----------------------------------------------------------------------------
func GetAllDataBreach()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DataBreach

	//----------------------------------------------------------------------------
	// Request the ORM to find all DataBreach
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DataBreach" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DataBreach", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDataBreach", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDataBreach - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDataBreach(obj model.DataBreach)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DataBreach using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DataBreach using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDataBreach", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDataBreach - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDataBreach(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDataBreach(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DataBreach)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DataBreach using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DataBreach using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDataBreach", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a DataBreach
//----------------------------------------------------------------------------
func AssignOrganizationToDataBreach( dataBreachId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

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
			// assign the Organization	to the DataBreach
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the DataBreach
			//----------------------------------------------------------------------------
			return UpdateDataBreach(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a DataBreach
//----------------------------------------------------------------------------
func UnassignOrganizationFromDataBreach(dataBreachId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the DataBreach
		//----------------------------------------------------------------------------
		return UpdateDataBreach(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Matter on a DataBreach
//----------------------------------------------------------------------------
func AssignMatterToDataBreach( dataBreachId uint64, matterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

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
			// assign the Matter	to the DataBreach
			//----------------------------------------------------------------------------
			parentObj.Matter = &childObj

			//----------------------------------------------------------------------------
			// save the DataBreach
			//----------------------------------------------------------------------------
			return UpdateDataBreach(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Matter", matterId )
			return utils.RequestResult{false, msg, "assignMatter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Matter on a DataBreach
//----------------------------------------------------------------------------
func UnassignMatterFromDataBreach(dataBreachId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

		//----------------------------------------------------------------------------
		// assign an empty Matter to the Matter
		//----------------------------------------------------------------------------
		parentObj.Matter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Matter
		//----------------------------------------------------------------------------
		parentObj.MatterId = nil;

		//----------------------------------------------------------------------------
		// save the DataBreach
		//----------------------------------------------------------------------------
		return UpdateDataBreach(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more processingActivitiesIds as a ProcessingActivities to a DataBreach
//----------------------------------------------------------------------------
func AddProcessingActivitiesToDataBreach ( dataBreachId uint64, processingActivitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

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
		// retrieve the modified DataBreach from the gorm
		//----------------------------------------------------------------------------
		return GetDataBreach(dataBreachId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more processingActivitiesIds as a ProcessingActivities from a DataBreach
//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromDataBreach( dataBreachId uint64, processingActivitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

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
		// retrieve the modified DataBreach from the gorm
		//----------------------------------------------------------------------------
		return GetDataBreach(dataBreachId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataCategoriesIds as a DataCategories to a DataBreach
//----------------------------------------------------------------------------
func AddDataCategoriesToDataBreach ( dataBreachId uint64, dataCategoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataCategoriesIds, ",")

		for _, dataCategoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataCategory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataCategory
			// with a matching dataCategoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataCategoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DataCategories using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataCategories").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataCategories", dataCategoriesId )
				return utils.RequestResult{false, msg, "unassignDataCategories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataBreach from the gorm
		//----------------------------------------------------------------------------
		return GetDataBreach(dataBreachId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataCategoriesIds as a DataCategories from a DataBreach
//----------------------------------------------------------------------------
func RemoveDataCategoriesFromDataBreach( dataBreachId uint64, dataCategoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataCategoriesIds, ",")

		for _, dataCategoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataCategory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataCategory
			// with a matching dataCategoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataCategoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataCategoryObj from the DataCategories array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataCategories").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataCategories", dataCategoriesId )
				return utils.RequestResult{false, msg, "removeDataCategories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataBreach from the gorm
		//----------------------------------------------------------------------------
		return GetDataBreach(dataBreachId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more thirdPartiesIds as a ThirdParties to a DataBreach
//----------------------------------------------------------------------------
func AddThirdPartiesToDataBreach ( dataBreachId uint64, thirdPartiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

		// slice the ids on comma with no spaces
		ids := strings.Split( thirdPartiesIds, ",")

		for _, thirdPartiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ThirdParty

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdParty
			// with a matching thirdPartiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , thirdPartiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ThirdParties using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ThirdParties").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdParties", thirdPartiesId )
				return utils.RequestResult{false, msg, "unassignThirdParties", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataBreach from the gorm
		//----------------------------------------------------------------------------
		return GetDataBreach(dataBreachId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more thirdPartiesIds as a ThirdParties from a DataBreach
//----------------------------------------------------------------------------
func RemoveThirdPartiesFromDataBreach( dataBreachId uint64, thirdPartiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataBreach with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataBreach(dataBreachId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataBreach so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataBreach)

		// slice the ids on comma with no spaces
		ids := strings.Split( thirdPartiesIds, ",")

		for _, thirdPartiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ThirdParty

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdParty
			// with a matching thirdPartiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , thirdPartiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ThirdPartyObj from the ThirdParties array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ThirdParties").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdParties", thirdPartiesId )
				return utils.RequestResult{false, msg, "removeThirdParties", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataBreach from the gorm
		//----------------------------------------------------------------------------
		return GetDataBreach(dataBreachId)

	} else {
		return parentRequestResult
	}
}

