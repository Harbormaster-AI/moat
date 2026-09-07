package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LegalHoldDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLegalHold - creates a new db entry
//----------------------------------------------------------------------------
func CreateLegalHold(obj model.LegalHold)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LegalHold with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LegalHold", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLegalHold", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLegalHold - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLegalHold(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LegalHold

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LegalHold with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LegalHold using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LegalHold using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLegalHold", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLegalHold - returns all
//----------------------------------------------------------------------------
func GetAllLegalHold()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LegalHold

	//----------------------------------------------------------------------------
	// Request the ORM to find all LegalHold
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LegalHold" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LegalHold", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLegalHold", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLegalHold - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLegalHold(obj model.LegalHold)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LegalHold using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LegalHold using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLegalHold", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLegalHold - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLegalHold(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LegalHold with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLegalHold(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LegalHold so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LegalHold)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LegalHold using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LegalHold using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLegalHold", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Matter on a LegalHold
//----------------------------------------------------------------------------
func AssignMatterToLegalHold( legalHoldId uint64, matterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LegalHold with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLegalHold(legalHoldId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LegalHold so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LegalHold)

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
			// assign the Matter	to the LegalHold
			//----------------------------------------------------------------------------
			parentObj.Matter = &childObj

			//----------------------------------------------------------------------------
			// save the LegalHold
			//----------------------------------------------------------------------------
			return UpdateLegalHold(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Matter", matterId )
			return utils.RequestResult{false, msg, "assignMatter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Matter on a LegalHold
//----------------------------------------------------------------------------
func UnassignMatterFromLegalHold(legalHoldId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LegalHold with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLegalHold(legalHoldId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LegalHold so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LegalHold)

		//----------------------------------------------------------------------------
		// assign an empty Matter to the Matter
		//----------------------------------------------------------------------------
		parentObj.Matter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Matter
		//----------------------------------------------------------------------------
		parentObj.MatterId = nil;

		//----------------------------------------------------------------------------
		// save the LegalHold
		//----------------------------------------------------------------------------
		return UpdateLegalHold(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more repositoriesIds as a Repositories to a LegalHold
//----------------------------------------------------------------------------
func AddRepositoriesToLegalHold ( legalHoldId uint64, repositoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LegalHold with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLegalHold(legalHoldId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LegalHold so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LegalHold)

		// slice the ids on comma with no spaces
		ids := strings.Split( repositoriesIds, ",")

		for _, repositoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RecordsRepository

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RecordsRepository
			// with a matching repositoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , repositoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Repositories using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Repositories").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Repositories", repositoriesId )
				return utils.RequestResult{false, msg, "unassignRepositories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LegalHold from the gorm
		//----------------------------------------------------------------------------
		return GetLegalHold(legalHoldId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more repositoriesIds as a Repositories from a LegalHold
//----------------------------------------------------------------------------
func RemoveRepositoriesFromLegalHold( legalHoldId uint64, repositoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LegalHold with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLegalHold(legalHoldId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LegalHold so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LegalHold)

		// slice the ids on comma with no spaces
		ids := strings.Split( repositoriesIds, ",")

		for _, repositoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RecordsRepository

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RecordsRepository
			// with a matching repositoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , repositoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RecordsRepositoryObj from the Repositories array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Repositories").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Repositories", repositoriesId )
				return utils.RequestResult{false, msg, "removeRepositories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LegalHold from the gorm
		//----------------------------------------------------------------------------
		return GetLegalHold(legalHoldId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more recordsIds as a Records to a LegalHold
//----------------------------------------------------------------------------
func AddRecordsToLegalHold ( legalHoldId uint64, recordsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LegalHold with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLegalHold(legalHoldId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LegalHold so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LegalHold)

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
		// retrieve the modified LegalHold from the gorm
		//----------------------------------------------------------------------------
		return GetLegalHold(legalHoldId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more recordsIds as a Records from a LegalHold
//----------------------------------------------------------------------------
func RemoveRecordsFromLegalHold( legalHoldId uint64, recordsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LegalHold with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLegalHold(legalHoldId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LegalHold so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LegalHold)

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
		// retrieve the modified LegalHold from the gorm
		//----------------------------------------------------------------------------
		return GetLegalHold(legalHoldId)

	} else {
		return parentRequestResult
	}
}

