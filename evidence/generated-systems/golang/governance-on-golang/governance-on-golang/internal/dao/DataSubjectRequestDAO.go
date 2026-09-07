package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DataSubjectRequestDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDataSubjectRequest - creates a new db entry
//----------------------------------------------------------------------------
func CreateDataSubjectRequest(obj model.DataSubjectRequest)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DataSubjectRequest with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DataSubjectRequest", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDataSubjectRequest", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDataSubjectRequest - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDataSubjectRequest(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DataSubjectRequest

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DataSubjectRequest with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DataSubjectRequest using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DataSubjectRequest using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDataSubjectRequest", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDataSubjectRequest - returns all
//----------------------------------------------------------------------------
func GetAllDataSubjectRequest()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DataSubjectRequest

	//----------------------------------------------------------------------------
	// Request the ORM to find all DataSubjectRequest
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DataSubjectRequest" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DataSubjectRequest", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDataSubjectRequest", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDataSubjectRequest - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDataSubjectRequest(obj model.DataSubjectRequest)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DataSubjectRequest using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DataSubjectRequest using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDataSubjectRequest", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDataSubjectRequest - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDataSubjectRequest(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DataSubjectRequest with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDataSubjectRequest(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSubjectRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DataSubjectRequest)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DataSubjectRequest using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DataSubjectRequest using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDataSubjectRequest", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a DataSubjectRequest
//----------------------------------------------------------------------------
func AssignOrganizationToDataSubjectRequest( dataSubjectRequestId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataSubjectRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSubjectRequest(dataSubjectRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSubjectRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSubjectRequest)

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
			// assign the Organization	to the DataSubjectRequest
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the DataSubjectRequest
			//----------------------------------------------------------------------------
			return UpdateDataSubjectRequest(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a DataSubjectRequest
//----------------------------------------------------------------------------
func UnassignOrganizationFromDataSubjectRequest(dataSubjectRequestId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSubjectRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSubjectRequest(dataSubjectRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSubjectRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSubjectRequest)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the DataSubjectRequest
		//----------------------------------------------------------------------------
		return UpdateDataSubjectRequest(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more processingActivitiesIds as a ProcessingActivities to a DataSubjectRequest
//----------------------------------------------------------------------------
func AddProcessingActivitiesToDataSubjectRequest ( dataSubjectRequestId uint64, processingActivitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSubjectRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSubjectRequest(dataSubjectRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSubjectRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSubjectRequest)

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
		// retrieve the modified DataSubjectRequest from the gorm
		//----------------------------------------------------------------------------
		return GetDataSubjectRequest(dataSubjectRequestId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more processingActivitiesIds as a ProcessingActivities from a DataSubjectRequest
//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromDataSubjectRequest( dataSubjectRequestId uint64, processingActivitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSubjectRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSubjectRequest(dataSubjectRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSubjectRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSubjectRequest)

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
		// retrieve the modified DataSubjectRequest from the gorm
		//----------------------------------------------------------------------------
		return GetDataSubjectRequest(dataSubjectRequestId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more recordsIds as a Records to a DataSubjectRequest
//----------------------------------------------------------------------------
func AddRecordsToDataSubjectRequest ( dataSubjectRequestId uint64, recordsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSubjectRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSubjectRequest(dataSubjectRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSubjectRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSubjectRequest)

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
		// retrieve the modified DataSubjectRequest from the gorm
		//----------------------------------------------------------------------------
		return GetDataSubjectRequest(dataSubjectRequestId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more recordsIds as a Records from a DataSubjectRequest
//----------------------------------------------------------------------------
func RemoveRecordsFromDataSubjectRequest( dataSubjectRequestId uint64, recordsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSubjectRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSubjectRequest(dataSubjectRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSubjectRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSubjectRequest)

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
		// retrieve the modified DataSubjectRequest from the gorm
		//----------------------------------------------------------------------------
		return GetDataSubjectRequest(dataSubjectRequestId)

	} else {
		return parentRequestResult
	}
}

