package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DispositionReviewDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDispositionReview - creates a new db entry
//----------------------------------------------------------------------------
func CreateDispositionReview(obj model.DispositionReview)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DispositionReview with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DispositionReview", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDispositionReview", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDispositionReview - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDispositionReview(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DispositionReview

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DispositionReview with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DispositionReview using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DispositionReview using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDispositionReview", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDispositionReview - returns all
//----------------------------------------------------------------------------
func GetAllDispositionReview()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DispositionReview

	//----------------------------------------------------------------------------
	// Request the ORM to find all DispositionReview
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DispositionReview" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DispositionReview", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDispositionReview", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDispositionReview - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDispositionReview(obj model.DispositionReview)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DispositionReview using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DispositionReview using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDispositionReview", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDispositionReview - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDispositionReview(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DispositionReview with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDispositionReview(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DispositionReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DispositionReview)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DispositionReview using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DispositionReview using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDispositionReview", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Record on a DispositionReview
//----------------------------------------------------------------------------
func AssignRecordToDispositionReview( dispositionReviewId uint64, recordId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DispositionReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispositionReview(dispositionReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DispositionReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DispositionReview)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Record_

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Record_ with a
		// matching recordId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, recordId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Record	to the DispositionReview
			//----------------------------------------------------------------------------
			parentObj.Record = &childObj

			//----------------------------------------------------------------------------
			// save the DispositionReview
			//----------------------------------------------------------------------------
			return UpdateDispositionReview(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Record", recordId )
			return utils.RequestResult{false, msg, "assignRecord", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Record on a DispositionReview
//----------------------------------------------------------------------------
func UnassignRecordFromDispositionReview(dispositionReviewId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DispositionReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispositionReview(dispositionReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DispositionReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DispositionReview)

		//----------------------------------------------------------------------------
		// assign an empty Record_ to the Record
		//----------------------------------------------------------------------------
		parentObj.Record = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Record
		//----------------------------------------------------------------------------
		parentObj.RecordId = nil;

		//----------------------------------------------------------------------------
		// save the DispositionReview
		//----------------------------------------------------------------------------
		return UpdateDispositionReview(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a RetentionSchedule on a DispositionReview
//----------------------------------------------------------------------------
func AssignRetentionScheduleToDispositionReview( dispositionReviewId uint64, retentionScheduleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DispositionReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispositionReview(dispositionReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DispositionReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DispositionReview)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.RetentionSchedule

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a RetentionSchedule with a
		// matching retentionScheduleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, retentionScheduleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the RetentionSchedule	to the DispositionReview
			//----------------------------------------------------------------------------
			parentObj.RetentionSchedule = &childObj

			//----------------------------------------------------------------------------
			// save the DispositionReview
			//----------------------------------------------------------------------------
			return UpdateDispositionReview(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RetentionSchedule", retentionScheduleId )
			return utils.RequestResult{false, msg, "assignRetentionSchedule", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a RetentionSchedule on a DispositionReview
//----------------------------------------------------------------------------
func UnassignRetentionScheduleFromDispositionReview(dispositionReviewId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DispositionReview with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDispositionReview(dispositionReviewId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DispositionReview so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DispositionReview)

		//----------------------------------------------------------------------------
		// assign an empty RetentionSchedule to the RetentionSchedule
		//----------------------------------------------------------------------------
		parentObj.RetentionSchedule = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the RetentionSchedule
		//----------------------------------------------------------------------------
		parentObj.RetentionScheduleId = nil;

		//----------------------------------------------------------------------------
		// save the DispositionReview
		//----------------------------------------------------------------------------
		return UpdateDispositionReview(parentObj)

	} else {
		return parentRequestResult
	}

}


