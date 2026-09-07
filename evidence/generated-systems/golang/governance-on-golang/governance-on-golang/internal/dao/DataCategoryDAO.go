package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DataCategoryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDataCategory - creates a new db entry
//----------------------------------------------------------------------------
func CreateDataCategory(obj model.DataCategory)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DataCategory with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DataCategory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDataCategory", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDataCategory - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDataCategory(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DataCategory

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DataCategory with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DataCategory using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DataCategory using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDataCategory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDataCategory - returns all
//----------------------------------------------------------------------------
func GetAllDataCategory()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DataCategory

	//----------------------------------------------------------------------------
	// Request the ORM to find all DataCategory
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DataCategory" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DataCategory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDataCategory", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDataCategory - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDataCategory(obj model.DataCategory)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DataCategory using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DataCategory using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDataCategory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDataCategory - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDataCategory(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DataCategory with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDataCategory(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataCategory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DataCategory)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DataCategory using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DataCategory using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDataCategory", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more processingActivitiesIds as a ProcessingActivities to a DataCategory
//----------------------------------------------------------------------------
func AddProcessingActivitiesToDataCategory ( dataCategoryId uint64, processingActivitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataCategory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataCategory(dataCategoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataCategory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataCategory)

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
		// retrieve the modified DataCategory from the gorm
		//----------------------------------------------------------------------------
		return GetDataCategory(dataCategoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more processingActivitiesIds as a ProcessingActivities from a DataCategory
//----------------------------------------------------------------------------
func RemoveProcessingActivitiesFromDataCategory( dataCategoryId uint64, processingActivitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataCategory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataCategory(dataCategoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataCategory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataCategory)

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
		// retrieve the modified DataCategory from the gorm
		//----------------------------------------------------------------------------
		return GetDataCategory(dataCategoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more recordsIds as a Records to a DataCategory
//----------------------------------------------------------------------------
func AddRecordsToDataCategory ( dataCategoryId uint64, recordsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataCategory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataCategory(dataCategoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataCategory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataCategory)

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
		// retrieve the modified DataCategory from the gorm
		//----------------------------------------------------------------------------
		return GetDataCategory(dataCategoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more recordsIds as a Records from a DataCategory
//----------------------------------------------------------------------------
func RemoveRecordsFromDataCategory( dataCategoryId uint64, recordsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataCategory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataCategory(dataCategoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataCategory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataCategory)

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
		// retrieve the modified DataCategory from the gorm
		//----------------------------------------------------------------------------
		return GetDataCategory(dataCategoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataBreachesIds as a DataBreaches to a DataCategory
//----------------------------------------------------------------------------
func AddDataBreachesToDataCategory ( dataCategoryId uint64, dataBreachesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataCategory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataCategory(dataCategoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataCategory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataCategory)

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
		// retrieve the modified DataCategory from the gorm
		//----------------------------------------------------------------------------
		return GetDataCategory(dataCategoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataBreachesIds as a DataBreaches from a DataCategory
//----------------------------------------------------------------------------
func RemoveDataBreachesFromDataCategory( dataCategoryId uint64, dataBreachesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataCategory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataCategory(dataCategoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataCategory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataCategory)

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
		// retrieve the modified DataCategory from the gorm
		//----------------------------------------------------------------------------
		return GetDataCategory(dataCategoryId)

	} else {
		return parentRequestResult
	}
}

