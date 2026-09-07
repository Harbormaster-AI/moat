package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LaboratoryOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLaboratoryOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateLaboratoryOrder(obj model.LaboratoryOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LaboratoryOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LaboratoryOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLaboratoryOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLaboratoryOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLaboratoryOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LaboratoryOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LaboratoryOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LaboratoryOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LaboratoryOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLaboratoryOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLaboratoryOrder - returns all
//----------------------------------------------------------------------------
func GetAllLaboratoryOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LaboratoryOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all LaboratoryOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LaboratoryOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LaboratoryOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLaboratoryOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLaboratoryOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLaboratoryOrder(obj model.LaboratoryOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LaboratoryOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LaboratoryOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLaboratoryOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLaboratoryOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLaboratoryOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LaboratoryOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLaboratoryOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LaboratoryOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LaboratoryOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LaboratoryOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LaboratoryOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLaboratoryOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Order on a LaboratoryOrder
//----------------------------------------------------------------------------
func AssignOrderToLaboratoryOrder( laboratoryOrderId uint64, orderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LaboratoryOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratoryOrder(laboratoryOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LaboratoryOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LaboratoryOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ClinicalOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ClinicalOrder with a
		// matching orderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, orderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Order	to the LaboratoryOrder
			//----------------------------------------------------------------------------
			parentObj.Order = &childObj

			//----------------------------------------------------------------------------
			// save the LaboratoryOrder
			//----------------------------------------------------------------------------
			return UpdateLaboratoryOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Order", orderId )
			return utils.RequestResult{false, msg, "assignOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Order on a LaboratoryOrder
//----------------------------------------------------------------------------
func UnassignOrderFromLaboratoryOrder(laboratoryOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LaboratoryOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratoryOrder(laboratoryOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LaboratoryOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LaboratoryOrder)

		//----------------------------------------------------------------------------
		// assign an empty ClinicalOrder to the Order
		//----------------------------------------------------------------------------
		parentObj.Order = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Order
		//----------------------------------------------------------------------------
		parentObj.OrderId = nil;

		//----------------------------------------------------------------------------
		// save the LaboratoryOrder
		//----------------------------------------------------------------------------
		return UpdateLaboratoryOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Laboratory on a LaboratoryOrder
//----------------------------------------------------------------------------
func AssignLaboratoryToLaboratoryOrder( laboratoryOrderId uint64, laboratoryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LaboratoryOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratoryOrder(laboratoryOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LaboratoryOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LaboratoryOrder)

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
			// assign the Laboratory	to the LaboratoryOrder
			//----------------------------------------------------------------------------
			parentObj.Laboratory = &childObj

			//----------------------------------------------------------------------------
			// save the LaboratoryOrder
			//----------------------------------------------------------------------------
			return UpdateLaboratoryOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Laboratory", laboratoryId )
			return utils.RequestResult{false, msg, "assignLaboratory", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Laboratory on a LaboratoryOrder
//----------------------------------------------------------------------------
func UnassignLaboratoryFromLaboratoryOrder(laboratoryOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LaboratoryOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratoryOrder(laboratoryOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LaboratoryOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LaboratoryOrder)

		//----------------------------------------------------------------------------
		// assign an empty Laboratory to the Laboratory
		//----------------------------------------------------------------------------
		parentObj.Laboratory = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Laboratory
		//----------------------------------------------------------------------------
		parentObj.LaboratoryId = nil;

		//----------------------------------------------------------------------------
		// save the LaboratoryOrder
		//----------------------------------------------------------------------------
		return UpdateLaboratoryOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more resultsIds as a Results to a LaboratoryOrder
//----------------------------------------------------------------------------
func AddResultsToLaboratoryOrder ( laboratoryOrderId uint64, resultsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LaboratoryOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratoryOrder(laboratoryOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LaboratoryOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LaboratoryOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( resultsIds, ",")

		for _, resultsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LabResult

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LabResult
			// with a matching resultsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , resultsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Results using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Results").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Results", resultsId )
				return utils.RequestResult{false, msg, "unassignResults", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LaboratoryOrder from the gorm
		//----------------------------------------------------------------------------
		return GetLaboratoryOrder(laboratoryOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more resultsIds as a Results from a LaboratoryOrder
//----------------------------------------------------------------------------
func RemoveResultsFromLaboratoryOrder( laboratoryOrderId uint64, resultsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LaboratoryOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLaboratoryOrder(laboratoryOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LaboratoryOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LaboratoryOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( resultsIds, ",")

		for _, resultsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LabResult

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LabResult
			// with a matching resultsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , resultsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LabResultObj from the Results array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Results").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Results", resultsId )
				return utils.RequestResult{false, msg, "removeResults", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LaboratoryOrder from the gorm
		//----------------------------------------------------------------------------
		return GetLaboratoryOrder(laboratoryOrderId)

	} else {
		return parentRequestResult
	}
}

