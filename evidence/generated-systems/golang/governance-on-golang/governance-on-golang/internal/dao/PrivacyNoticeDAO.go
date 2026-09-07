package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PrivacyNoticeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePrivacyNotice - creates a new db entry
//----------------------------------------------------------------------------
func CreatePrivacyNotice(obj model.PrivacyNotice)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PrivacyNotice with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PrivacyNotice", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePrivacyNotice", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPrivacyNotice - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPrivacyNotice(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PrivacyNotice

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PrivacyNotice with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PrivacyNotice using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PrivacyNotice using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPrivacyNotice", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPrivacyNotice - returns all
//----------------------------------------------------------------------------
func GetAllPrivacyNotice()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PrivacyNotice

	//----------------------------------------------------------------------------
	// Request the ORM to find all PrivacyNotice
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PrivacyNotice" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PrivacyNotice", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPrivacyNotice", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePrivacyNotice - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePrivacyNotice(obj model.PrivacyNotice)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PrivacyNotice using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PrivacyNotice using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePrivacyNotice", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePrivacyNotice - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePrivacyNotice(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PrivacyNotice with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPrivacyNotice(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PrivacyNotice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PrivacyNotice)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PrivacyNotice using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PrivacyNotice using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePrivacyNotice", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a PrivacyNotice
//----------------------------------------------------------------------------
func AssignOrganizationToPrivacyNotice( privacyNoticeId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PrivacyNotice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrivacyNotice(privacyNoticeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PrivacyNotice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PrivacyNotice)

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
			// assign the Organization	to the PrivacyNotice
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the PrivacyNotice
			//----------------------------------------------------------------------------
			return UpdatePrivacyNotice(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a PrivacyNotice
//----------------------------------------------------------------------------
func UnassignOrganizationFromPrivacyNotice(privacyNoticeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PrivacyNotice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrivacyNotice(privacyNoticeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PrivacyNotice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PrivacyNotice)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the PrivacyNotice
		//----------------------------------------------------------------------------
		return UpdatePrivacyNotice(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more processingActivitiesIds as a ProcessingActivities to a PrivacyNotice
//----------------------------------------------------------------------------
func AddProcessingActivitiesToPrivacyNotice ( privacyNoticeId uint64, processingActivitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PrivacyNotice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrivacyNotice(privacyNoticeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PrivacyNotice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PrivacyNotice)

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
		// retrieve the modified PrivacyNotice from the gorm
		//----------------------------------------------------------------------------
		return GetPrivacyNotice(privacyNoticeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more processingActivitiesIds as a ProcessingActivities from a PrivacyNotice
//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromPrivacyNotice( privacyNoticeId uint64, processingActivitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PrivacyNotice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrivacyNotice(privacyNoticeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PrivacyNotice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PrivacyNotice)

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
		// retrieve the modified PrivacyNotice from the gorm
		//----------------------------------------------------------------------------
		return GetPrivacyNotice(privacyNoticeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more consentsIds as a Consents to a PrivacyNotice
//----------------------------------------------------------------------------
func AddConsentsToPrivacyNotice ( privacyNoticeId uint64, consentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PrivacyNotice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrivacyNotice(privacyNoticeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PrivacyNotice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PrivacyNotice)

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
		// retrieve the modified PrivacyNotice from the gorm
		//----------------------------------------------------------------------------
		return GetPrivacyNotice(privacyNoticeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more consentsIds as a Consents from a PrivacyNotice
//----------------------------------------------------------------------------
func RemoveConsentsFromPrivacyNotice( privacyNoticeId uint64, consentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PrivacyNotice with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrivacyNotice(privacyNoticeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PrivacyNotice so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PrivacyNotice)

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
		// retrieve the modified PrivacyNotice from the gorm
		//----------------------------------------------------------------------------
		return GetPrivacyNotice(privacyNoticeId)

	} else {
		return parentRequestResult
	}
}

