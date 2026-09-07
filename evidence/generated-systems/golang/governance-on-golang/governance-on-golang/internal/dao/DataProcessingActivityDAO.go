package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DataProcessingActivityDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDataProcessingActivity - creates a new db entry
//----------------------------------------------------------------------------
func CreateDataProcessingActivity(obj model.DataProcessingActivity)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DataProcessingActivity with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DataProcessingActivity", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDataProcessingActivity", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDataProcessingActivity - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDataProcessingActivity(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DataProcessingActivity

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DataProcessingActivity with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DataProcessingActivity using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DataProcessingActivity using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDataProcessingActivity", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDataProcessingActivity - returns all
//----------------------------------------------------------------------------
func GetAllDataProcessingActivity()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DataProcessingActivity

	//----------------------------------------------------------------------------
	// Request the ORM to find all DataProcessingActivity
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DataProcessingActivity" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DataProcessingActivity", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDataProcessingActivity", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDataProcessingActivity - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDataProcessingActivity(obj model.DataProcessingActivity)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DataProcessingActivity using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DataProcessingActivity using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDataProcessingActivity", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDataProcessingActivity - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDataProcessingActivity(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDataProcessingActivity(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DataProcessingActivity)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DataProcessingActivity using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DataProcessingActivity using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDataProcessingActivity", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a DataProcessingActivity
//----------------------------------------------------------------------------
func AssignOrganizationToDataProcessingActivity( dataProcessingActivityId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

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
			// assign the Organization	to the DataProcessingActivity
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the DataProcessingActivity
			//----------------------------------------------------------------------------
			return UpdateDataProcessingActivity(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a DataProcessingActivity
//----------------------------------------------------------------------------
func UnassignOrganizationFromDataProcessingActivity(dataProcessingActivityId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the DataProcessingActivity
		//----------------------------------------------------------------------------
		return UpdateDataProcessingActivity(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more dataCategoriesIds as a DataCategories to a DataProcessingActivity
//----------------------------------------------------------------------------
func AddDataCategoriesToDataProcessingActivity ( dataProcessingActivityId uint64, dataCategoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

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
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataCategoriesIds as a DataCategories from a DataProcessingActivity
//----------------------------------------------------------------------------
func RemoveDataCategoriesFromDataProcessingActivity( dataProcessingActivityId uint64, dataCategoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

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
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more systemsIds as a Systems to a DataProcessingActivity
//----------------------------------------------------------------------------
func AddSystemsToDataProcessingActivity ( dataProcessingActivityId uint64, systemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( systemsIds, ",")

		for _, systemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.System_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a System_
			// with a matching systemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , systemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Systems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Systems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Systems", systemsId )
				return utils.RequestResult{false, msg, "unassignSystems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more systemsIds as a Systems from a DataProcessingActivity
//----------------------------------------------------------------------------
func RemoveSystemsFromDataProcessingActivity( dataProcessingActivityId uint64, systemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( systemsIds, ",")

		for _, systemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.System_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a System_
			// with a matching systemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , systemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove System_Obj from the Systems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Systems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Systems", systemsId )
				return utils.RequestResult{false, msg, "removeSystems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more recordsIds as a Records to a DataProcessingActivity
//----------------------------------------------------------------------------
func AddRecordsToDataProcessingActivity ( dataProcessingActivityId uint64, recordsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( recordsIds, ",")

		for _, recordsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Record_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Record_
			// with a matching recordsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , recordsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Records using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Records").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Records", recordsId )
				return utils.RequestResult{false, msg, "unassignRecords", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more recordsIds as a Records from a DataProcessingActivity
//----------------------------------------------------------------------------
func RemoveRecordsFromDataProcessingActivity( dataProcessingActivityId uint64, recordsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( recordsIds, ",")

		for _, recordsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Record_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Record_
			// with a matching recordsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , recordsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Record_Obj from the Records array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Records").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Records", recordsId )
				return utils.RequestResult{false, msg, "removeRecords", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more privacyNoticesIds as a PrivacyNotices to a DataProcessingActivity
//----------------------------------------------------------------------------
func AddPrivacyNoticesToDataProcessingActivity ( dataProcessingActivityId uint64, privacyNoticesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( privacyNoticesIds, ",")

		for _, privacyNoticesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PrivacyNotice

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PrivacyNotice
			// with a matching privacyNoticesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , privacyNoticesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PrivacyNotices using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PrivacyNotices").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PrivacyNotices", privacyNoticesId )
				return utils.RequestResult{false, msg, "unassignPrivacyNotices", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more privacyNoticesIds as a PrivacyNotices from a DataProcessingActivity
//----------------------------------------------------------------------------
func RemovePrivacyNoticesFromDataProcessingActivity( dataProcessingActivityId uint64, privacyNoticesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( privacyNoticesIds, ",")

		for _, privacyNoticesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PrivacyNotice

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PrivacyNotice
			// with a matching privacyNoticesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , privacyNoticesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PrivacyNoticeObj from the PrivacyNotices array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PrivacyNotices").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PrivacyNotices", privacyNoticesId )
				return utils.RequestResult{false, msg, "removePrivacyNotices", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more thirdPartiesIds as a ThirdParties to a DataProcessingActivity
//----------------------------------------------------------------------------
func AddThirdPartiesToDataProcessingActivity ( dataProcessingActivityId uint64, thirdPartiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

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
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more thirdPartiesIds as a ThirdParties from a DataProcessingActivity
//----------------------------------------------------------------------------
func RemoveThirdPartiesFromDataProcessingActivity( dataProcessingActivityId uint64, thirdPartiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

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
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more consentsIds as a Consents to a DataProcessingActivity
//----------------------------------------------------------------------------
func AddConsentsToDataProcessingActivity ( dataProcessingActivityId uint64, consentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Consents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "unassignConsents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more consentsIds as a Consents from a DataProcessingActivity
//----------------------------------------------------------------------------
func RemoveConsentsFromDataProcessingActivity( dataProcessingActivityId uint64, consentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( consentsIds, ",")

		for _, consentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Consent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Consent
			// with a matching consentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , consentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ConsentObj from the Consents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Consents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Consents", consentsId )
				return utils.RequestResult{false, msg, "removeConsents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataBreachesIds as a DataBreaches to a DataProcessingActivity
//----------------------------------------------------------------------------
func AddDataBreachesToDataProcessingActivity ( dataProcessingActivityId uint64, dataBreachesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

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
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataBreachesIds as a DataBreaches from a DataProcessingActivity
//----------------------------------------------------------------------------
func RemoveDataBreachesFromDataProcessingActivity( dataProcessingActivityId uint64, dataBreachesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

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
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataSubjectRequestsIds as a DataSubjectRequests to a DataProcessingActivity
//----------------------------------------------------------------------------
func AddDataSubjectRequestsToDataProcessingActivity ( dataProcessingActivityId uint64, dataSubjectRequestsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataSubjectRequestsIds, ",")

		for _, dataSubjectRequestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSubjectRequest

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSubjectRequest
			// with a matching dataSubjectRequestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataSubjectRequestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DataSubjectRequests using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataSubjectRequests").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataSubjectRequests", dataSubjectRequestsId )
				return utils.RequestResult{false, msg, "unassignDataSubjectRequests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataSubjectRequestsIds as a DataSubjectRequests from a DataProcessingActivity
//----------------------------------------------------------------------------
func RemoveDataSubjectRequestsFromDataProcessingActivity( dataProcessingActivityId uint64, dataSubjectRequestsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataProcessingActivity with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataProcessingActivity(dataProcessingActivityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataProcessingActivity so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataProcessingActivity)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataSubjectRequestsIds, ",")

		for _, dataSubjectRequestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSubjectRequest

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSubjectRequest
			// with a matching dataSubjectRequestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataSubjectRequestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSubjectRequestObj from the DataSubjectRequests array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataSubjectRequests").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataSubjectRequests", dataSubjectRequestsId )
				return utils.RequestResult{false, msg, "removeDataSubjectRequests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataProcessingActivity from the gorm
		//----------------------------------------------------------------------------
		return GetDataProcessingActivity(dataProcessingActivityId)

	} else {
		return parentRequestResult
	}
}

