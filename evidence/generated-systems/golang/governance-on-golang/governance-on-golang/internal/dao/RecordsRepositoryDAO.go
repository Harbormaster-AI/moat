package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RecordsRepositoryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRecordsRepository - creates a new db entry
//----------------------------------------------------------------------------
func CreateRecordsRepository(obj model.RecordsRepository)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RecordsRepository with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RecordsRepository", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRecordsRepository", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRecordsRepository - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRecordsRepository(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RecordsRepository

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RecordsRepository with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RecordsRepository using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RecordsRepository using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRecordsRepository", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRecordsRepository - returns all
//----------------------------------------------------------------------------
func GetAllRecordsRepository()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RecordsRepository

	//----------------------------------------------------------------------------
	// Request the ORM to find all RecordsRepository
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all RecordsRepository" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RecordsRepository", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRecordsRepository", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRecordsRepository - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRecordsRepository(obj model.RecordsRepository)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RecordsRepository using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RecordsRepository using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRecordsRepository", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRecordsRepository - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRecordsRepository(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRecordsRepository(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.RecordsRepository)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RecordsRepository using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RecordsRepository using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRecordsRepository", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a RecordsRepository
//----------------------------------------------------------------------------
func AssignOrganizationToRecordsRepository( recordsRepositoryId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

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
			// assign the Organization	to the RecordsRepository
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the RecordsRepository
			//----------------------------------------------------------------------------
			return UpdateRecordsRepository(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a RecordsRepository
//----------------------------------------------------------------------------
func UnassignOrganizationFromRecordsRepository(recordsRepositoryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the RecordsRepository
		//----------------------------------------------------------------------------
		return UpdateRecordsRepository(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more recordsIds as a Records to a RecordsRepository
//----------------------------------------------------------------------------
func AddRecordsToRecordsRepository ( recordsRepositoryId uint64, recordsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

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
		// retrieve the modified RecordsRepository from the gorm
		//----------------------------------------------------------------------------
		return GetRecordsRepository(recordsRepositoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more recordsIds as a Records from a RecordsRepository
//----------------------------------------------------------------------------
func RemoveRecordsFromRecordsRepository( recordsRepositoryId uint64, recordsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

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
		// retrieve the modified RecordsRepository from the gorm
		//----------------------------------------------------------------------------
		return GetRecordsRepository(recordsRepositoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more systemsIds as a Systems to a RecordsRepository
//----------------------------------------------------------------------------
func AddSystemsToRecordsRepository ( recordsRepositoryId uint64, systemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

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
		// retrieve the modified RecordsRepository from the gorm
		//----------------------------------------------------------------------------
		return GetRecordsRepository(recordsRepositoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more systemsIds as a Systems from a RecordsRepository
//----------------------------------------------------------------------------
func RemoveSystemsFromRecordsRepository( recordsRepositoryId uint64, systemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

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
		// retrieve the modified RecordsRepository from the gorm
		//----------------------------------------------------------------------------
		return GetRecordsRepository(recordsRepositoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more retentionSchedulesIds as a RetentionSchedules to a RecordsRepository
//----------------------------------------------------------------------------
func AddRetentionSchedulesToRecordsRepository ( recordsRepositoryId uint64, retentionSchedulesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

		// slice the ids on comma with no spaces
		ids := strings.Split( retentionSchedulesIds, ",")

		for _, retentionSchedulesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RetentionSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RetentionSchedule
			// with a matching retentionSchedulesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , retentionSchedulesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RetentionSchedules using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RetentionSchedules").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RetentionSchedules", retentionSchedulesId )
				return utils.RequestResult{false, msg, "unassignRetentionSchedules", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecordsRepository from the gorm
		//----------------------------------------------------------------------------
		return GetRecordsRepository(recordsRepositoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more retentionSchedulesIds as a RetentionSchedules from a RecordsRepository
//----------------------------------------------------------------------------
func RemoveRetentionSchedulesFromRecordsRepository( recordsRepositoryId uint64, retentionSchedulesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

		// slice the ids on comma with no spaces
		ids := strings.Split( retentionSchedulesIds, ",")

		for _, retentionSchedulesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RetentionSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RetentionSchedule
			// with a matching retentionSchedulesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , retentionSchedulesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RetentionScheduleObj from the RetentionSchedules array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RetentionSchedules").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RetentionSchedules", retentionSchedulesId )
				return utils.RequestResult{false, msg, "removeRetentionSchedules", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RecordsRepository from the gorm
		//----------------------------------------------------------------------------
		return GetRecordsRepository(recordsRepositoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more legalHoldsIds as a LegalHolds to a RecordsRepository
//----------------------------------------------------------------------------
func AddLegalHoldsToRecordsRepository ( recordsRepositoryId uint64, legalHoldsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

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
		// retrieve the modified RecordsRepository from the gorm
		//----------------------------------------------------------------------------
		return GetRecordsRepository(recordsRepositoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more legalHoldsIds as a LegalHolds from a RecordsRepository
//----------------------------------------------------------------------------
func RemoveLegalHoldsFromRecordsRepository( recordsRepositoryId uint64, legalHoldsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RecordsRepository with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRecordsRepository(recordsRepositoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RecordsRepository so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RecordsRepository)

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
		// retrieve the modified RecordsRepository from the gorm
		//----------------------------------------------------------------------------
		return GetRecordsRepository(recordsRepositoryId)

	} else {
		return parentRequestResult
	}
}

