package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing WorkScheduleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateWorkSchedule - creates a new db entry
//----------------------------------------------------------------------------
func CreateWorkSchedule(obj model.WorkSchedule)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a WorkSchedule with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a WorkSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateWorkSchedule", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetWorkSchedule - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetWorkSchedule(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.WorkSchedule

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a WorkSchedule with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a WorkSchedule using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a WorkSchedule using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetWorkSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllWorkSchedule - returns all
//----------------------------------------------------------------------------
func GetAllWorkSchedule()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.WorkSchedule

	//----------------------------------------------------------------------------
	// Request the ORM to find all WorkSchedule
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all WorkSchedule" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all WorkSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllWorkSchedule", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateWorkSchedule - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateWorkSchedule(obj model.WorkSchedule)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a WorkSchedule using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a WorkSchedule using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateWorkSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteWorkSchedule - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteWorkSchedule(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the WorkSchedule with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetWorkSchedule(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.WorkSchedule)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a WorkSchedule using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a WorkSchedule using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteWorkSchedule", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more contractsIds as a Contracts to a WorkSchedule
//----------------------------------------------------------------------------
func AddContractsToWorkSchedule ( workScheduleId uint64, contractsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkSchedule(workScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmploymentContract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmploymentContract
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
		// retrieve the modified WorkSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetWorkSchedule(workScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more contractsIds as a Contracts from a WorkSchedule
//----------------------------------------------------------------------------
func RemoveContractsFromWorkSchedule( workScheduleId uint64, contractsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the WorkSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkSchedule(workScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( contractsIds, ",")

		for _, contractsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EmploymentContract

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EmploymentContract
			// with a matching contractsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , contractsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EmploymentContractObj from the Contracts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Contracts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contracts", contractsId )
				return utils.RequestResult{false, msg, "removeContracts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetWorkSchedule(workScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more shiftsIds as a Shifts to a WorkSchedule
//----------------------------------------------------------------------------
func AddShiftsToWorkSchedule ( workScheduleId uint64, shiftsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkSchedule(workScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( shiftsIds, ",")

		for _, shiftsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkShift

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkShift
			// with a matching shiftsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , shiftsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Shifts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Shifts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Shifts", shiftsId )
				return utils.RequestResult{false, msg, "unassignShifts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetWorkSchedule(workScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more shiftsIds as a Shifts from a WorkSchedule
//----------------------------------------------------------------------------
func RemoveShiftsFromWorkSchedule( workScheduleId uint64, shiftsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the WorkSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkSchedule(workScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( shiftsIds, ",")

		for _, shiftsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkShift

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkShift
			// with a matching shiftsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , shiftsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove WorkShiftObj from the Shifts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Shifts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Shifts", shiftsId )
				return utils.RequestResult{false, msg, "removeShifts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetWorkSchedule(workScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more exceptionsIds as a Exceptions to a WorkSchedule
//----------------------------------------------------------------------------
func AddExceptionsToWorkSchedule ( workScheduleId uint64, exceptionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkSchedule(workScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( exceptionsIds, ",")

		for _, exceptionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ScheduleException

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ScheduleException
			// with a matching exceptionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exceptionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Exceptions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Exceptions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exceptions", exceptionsId )
				return utils.RequestResult{false, msg, "unassignExceptions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetWorkSchedule(workScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more exceptionsIds as a Exceptions from a WorkSchedule
//----------------------------------------------------------------------------
func RemoveExceptionsFromWorkSchedule( workScheduleId uint64, exceptionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the WorkSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkSchedule(workScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( exceptionsIds, ",")

		for _, exceptionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ScheduleException

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ScheduleException
			// with a matching exceptionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exceptionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ScheduleExceptionObj from the Exceptions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Exceptions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exceptions", exceptionsId )
				return utils.RequestResult{false, msg, "removeExceptions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetWorkSchedule(workScheduleId)

	} else {
		return parentRequestResult
	}
}

