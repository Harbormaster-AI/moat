package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LabResultDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLabResult - creates a new db entry
//----------------------------------------------------------------------------
func CreateLabResult(obj model.LabResult)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LabResult with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LabResult", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLabResult", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLabResult - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLabResult(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LabResult

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LabResult with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LabResult using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LabResult using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLabResult", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLabResult - returns all
//----------------------------------------------------------------------------
func GetAllLabResult()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LabResult

	//----------------------------------------------------------------------------
	// Request the ORM to find all LabResult
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LabResult" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LabResult", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLabResult", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLabResult - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLabResult(obj model.LabResult)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LabResult using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LabResult using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLabResult", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLabResult - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLabResult(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LabResult with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLabResult(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LabResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LabResult)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LabResult using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LabResult using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLabResult", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a LaboratoryOrder on a LabResult
//----------------------------------------------------------------------------
func AssignLaboratoryOrderToLabResult( labResultId uint64, laboratoryOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LabResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLabResult(labResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LabResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LabResult)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LaboratoryOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LaboratoryOrder with a
		// matching laboratoryOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, laboratoryOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LaboratoryOrder	to the LabResult
			//----------------------------------------------------------------------------
			parentObj.LaboratoryOrder = &childObj

			//----------------------------------------------------------------------------
			// save the LabResult
			//----------------------------------------------------------------------------
			return UpdateLabResult(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LaboratoryOrder", laboratoryOrderId )
			return utils.RequestResult{false, msg, "assignLaboratoryOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LaboratoryOrder on a LabResult
//----------------------------------------------------------------------------
func UnassignLaboratoryOrderFromLabResult(labResultId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LabResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLabResult(labResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LabResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LabResult)

		//----------------------------------------------------------------------------
		// assign an empty LaboratoryOrder to the LaboratoryOrder
		//----------------------------------------------------------------------------
		parentObj.LaboratoryOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LaboratoryOrder
		//----------------------------------------------------------------------------
		parentObj.LaboratoryOrderId = nil;

		//----------------------------------------------------------------------------
		// save the LabResult
		//----------------------------------------------------------------------------
		return UpdateLabResult(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Laboratory on a LabResult
//----------------------------------------------------------------------------
func AssignLaboratoryToLabResult( labResultId uint64, laboratoryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LabResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLabResult(labResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LabResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LabResult)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Laboratory

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Laboratory with a
		// matching laboratoryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, laboratoryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Laboratory	to the LabResult
			//----------------------------------------------------------------------------
			parentObj.Laboratory = &childObj

			//----------------------------------------------------------------------------
			// save the LabResult
			//----------------------------------------------------------------------------
			return UpdateLabResult(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Laboratory", laboratoryId )
			return utils.RequestResult{false, msg, "assignLaboratory", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Laboratory on a LabResult
//----------------------------------------------------------------------------
func UnassignLaboratoryFromLabResult(labResultId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LabResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLabResult(labResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LabResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LabResult)

		//----------------------------------------------------------------------------
		// assign an empty Laboratory to the Laboratory
		//----------------------------------------------------------------------------
		parentObj.Laboratory = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Laboratory
		//----------------------------------------------------------------------------
		parentObj.LaboratoryId = nil;

		//----------------------------------------------------------------------------
		// save the LabResult
		//----------------------------------------------------------------------------
		return UpdateLabResult(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more observationsIds as a Observations to a LabResult
//----------------------------------------------------------------------------
func AddObservationsToLabResult ( labResultId uint64, observationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LabResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLabResult(labResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LabResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LabResult)

		// slice the ids on comma with no spaces
		ids := strings.Split( observationsIds, ",")

		for _, observationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Observation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Observation
			// with a matching observationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , observationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Observations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Observations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Observations", observationsId )
				return utils.RequestResult{false, msg, "unassignObservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LabResult from the gorm
		//----------------------------------------------------------------------------
		return GetLabResult(labResultId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more observationsIds as a Observations from a LabResult
//----------------------------------------------------------------------------
func RemoveObservationsFromLabResult( labResultId uint64, observationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LabResult with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLabResult(labResultId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LabResult so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LabResult)

		// slice the ids on comma with no spaces
		ids := strings.Split( observationsIds, ",")

		for _, observationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Observation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Observation
			// with a matching observationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , observationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ObservationObj from the Observations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Observations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Observations", observationsId )
				return utils.RequestResult{false, msg, "removeObservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LabResult from the gorm
		//----------------------------------------------------------------------------
		return GetLabResult(labResultId)

	} else {
		return parentRequestResult
	}
}

