package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LaboratoryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLaboratory - creates a new db entry
//----------------------------------------------------------------------------
func CreateLaboratory(obj model.Laboratory)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Laboratory with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Laboratory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLaboratory", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLaboratory - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLaboratory(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Laboratory

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Laboratory with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Laboratory using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Laboratory using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLaboratory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLaboratory - returns all
//----------------------------------------------------------------------------
func GetAllLaboratory()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Laboratory

	//----------------------------------------------------------------------------
	// Request the ORM to find all Laboratory
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Laboratory" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Laboratory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLaboratory", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLaboratory - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLaboratory(obj model.Laboratory)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Laboratory using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Laboratory using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLaboratory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLaboratory - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLaboratory(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Laboratory with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLaboratory(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Laboratory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Laboratory)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Laboratory using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Laboratory using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLaboratory", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Facility on a Laboratory
//----------------------------------------------------------------------------
func AssignFacilityToLaboratory( laboratoryId uint64, facilityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Laboratory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratory(laboratoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Laboratory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Laboratory)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Facility

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Facility with a
		// matching facilityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, facilityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Facility	to the Laboratory
			//----------------------------------------------------------------------------
			parentObj.Facility = &childObj

			//----------------------------------------------------------------------------
			// save the Laboratory
			//----------------------------------------------------------------------------
			return UpdateLaboratory(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facility", facilityId )
			return utils.RequestResult{false, msg, "assignFacility", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Facility on a Laboratory
//----------------------------------------------------------------------------
func UnassignFacilityFromLaboratory(laboratoryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Laboratory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratory(laboratoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Laboratory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Laboratory)

		//----------------------------------------------------------------------------
		// assign an empty Facility to the Facility
		//----------------------------------------------------------------------------
		parentObj.Facility = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Facility
		//----------------------------------------------------------------------------
		parentObj.FacilityId = nil;

		//----------------------------------------------------------------------------
		// save the Laboratory
		//----------------------------------------------------------------------------
		return UpdateLaboratory(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more laboratoryOrdersIds as a LaboratoryOrders to a Laboratory
//----------------------------------------------------------------------------
func AddLaboratoryOrdersToLaboratory ( laboratoryId uint64, laboratoryOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Laboratory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratory(laboratoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Laboratory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Laboratory)

		// slice the ids on comma with no spaces
		ids := strings.Split( laboratoryOrdersIds, ",")

		for _, laboratoryOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LaboratoryOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LaboratoryOrder
			// with a matching laboratoryOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , laboratoryOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LaboratoryOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LaboratoryOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LaboratoryOrders", laboratoryOrdersId )
				return utils.RequestResult{false, msg, "unassignLaboratoryOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Laboratory from the gorm
		//----------------------------------------------------------------------------
		return GetLaboratory(laboratoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more laboratoryOrdersIds as a LaboratoryOrders from a Laboratory
//----------------------------------------------------------------------------
func RemoveLaboratoryOrdersFromLaboratory( laboratoryId uint64, laboratoryOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Laboratory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratory(laboratoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Laboratory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Laboratory)

		// slice the ids on comma with no spaces
		ids := strings.Split( laboratoryOrdersIds, ",")

		for _, laboratoryOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LaboratoryOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LaboratoryOrder
			// with a matching laboratoryOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , laboratoryOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LaboratoryOrderObj from the LaboratoryOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LaboratoryOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LaboratoryOrders", laboratoryOrdersId )
				return utils.RequestResult{false, msg, "removeLaboratoryOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Laboratory from the gorm
		//----------------------------------------------------------------------------
		return GetLaboratory(laboratoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more labResultsIds as a LabResults to a Laboratory
//----------------------------------------------------------------------------
func AddLabResultsToLaboratory ( laboratoryId uint64, labResultsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Laboratory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratory(laboratoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Laboratory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Laboratory)

		// slice the ids on comma with no spaces
		ids := strings.Split( labResultsIds, ",")

		for _, labResultsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LabResult

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LabResult
			// with a matching labResultsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , labResultsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LabResults using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LabResults").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LabResults", labResultsId )
				return utils.RequestResult{false, msg, "unassignLabResults", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Laboratory from the gorm
		//----------------------------------------------------------------------------
		return GetLaboratory(laboratoryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more labResultsIds as a LabResults from a Laboratory
//----------------------------------------------------------------------------
func RemoveLabResultsFromLaboratory( laboratoryId uint64, labResultsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Laboratory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratory(laboratoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Laboratory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Laboratory)

		// slice the ids on comma with no spaces
		ids := strings.Split( labResultsIds, ",")

		for _, labResultsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LabResult

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LabResult
			// with a matching labResultsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , labResultsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LabResultObj from the LabResults array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LabResults").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LabResults", labResultsId )
				return utils.RequestResult{false, msg, "removeLabResults", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Laboratory from the gorm
		//----------------------------------------------------------------------------
		return GetLaboratory(laboratoryId)

	} else {
		return parentRequestResult
	}
}

