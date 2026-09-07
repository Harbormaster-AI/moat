package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing System_DAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSystem_ - creates a new db entry
//----------------------------------------------------------------------------
func CreateSystem_(obj model.System_)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a System_ with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a System_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSystem_", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSystem_ - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSystem_(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.System_

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a System_ with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a System_ using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a System_ using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSystem_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSystem_ - returns all
//----------------------------------------------------------------------------
func GetAllSystem_()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.System_

	//----------------------------------------------------------------------------
	// Request the ORM to find all System_
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all System_" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all System_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSystem_", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSystem_ - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSystem_(obj model.System_)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a System_ using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a System_ using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSystem_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSystem_ - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSystem_(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the System_ with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSystem_(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.System_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.System_)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a System_ using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a System_ using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSystem_", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more processingActivitiesIds as a ProcessingActivities to a System_
//----------------------------------------------------------------------------
func AddProcessingActivitiesToSystem_ ( system_Id uint64, processingActivitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the System_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSystem_(system_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.System_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.System_)

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
		// retrieve the modified System_ from the gorm
		//----------------------------------------------------------------------------
		return GetSystem_(system_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more processingActivitiesIds as a ProcessingActivities from a System_
//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromSystem_( system_Id uint64, processingActivitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the System_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSystem_(system_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.System_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.System_)

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
		// retrieve the modified System_ from the gorm
		//----------------------------------------------------------------------------
		return GetSystem_(system_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more recordsRepositoriesIds as a RecordsRepositories to a System_
//----------------------------------------------------------------------------
func AddRecordsRepositoriesToSystem_ ( system_Id uint64, recordsRepositoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the System_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSystem_(system_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.System_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.System_)

		// slice the ids on comma with no spaces
		ids := strings.Split( recordsRepositoriesIds, ",")

		for _, recordsRepositoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RecordsRepository

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RecordsRepository
			// with a matching recordsRepositoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , recordsRepositoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RecordsRepositories using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RecordsRepositories").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RecordsRepositories", recordsRepositoriesId )
				return utils.RequestResult{false, msg, "unassignRecordsRepositories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified System_ from the gorm
		//----------------------------------------------------------------------------
		return GetSystem_(system_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more recordsRepositoriesIds as a RecordsRepositories from a System_
//----------------------------------------------------------------------------
func RemoveRecordsRepositoriesFromSystem_( system_Id uint64, recordsRepositoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the System_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSystem_(system_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.System_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.System_)

		// slice the ids on comma with no spaces
		ids := strings.Split( recordsRepositoriesIds, ",")

		for _, recordsRepositoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RecordsRepository

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RecordsRepository
			// with a matching recordsRepositoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , recordsRepositoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RecordsRepositoryObj from the RecordsRepositories array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RecordsRepositories").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RecordsRepositories", recordsRepositoriesId )
				return utils.RequestResult{false, msg, "removeRecordsRepositories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified System_ from the gorm
		//----------------------------------------------------------------------------
		return GetSystem_(system_Id)

	} else {
		return parentRequestResult
	}
}

