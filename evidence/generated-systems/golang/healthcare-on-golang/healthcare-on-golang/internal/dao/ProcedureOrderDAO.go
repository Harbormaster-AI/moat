package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ProcedureOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateProcedureOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateProcedureOrder(obj model.ProcedureOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ProcedureOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ProcedureOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateProcedureOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetProcedureOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetProcedureOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ProcedureOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ProcedureOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ProcedureOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ProcedureOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetProcedureOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllProcedureOrder - returns all
//----------------------------------------------------------------------------
func GetAllProcedureOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ProcedureOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all ProcedureOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ProcedureOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ProcedureOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllProcedureOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateProcedureOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateProcedureOrder(obj model.ProcedureOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ProcedureOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ProcedureOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateProcedureOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteProcedureOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteProcedureOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ProcedureOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetProcedureOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProcedureOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ProcedureOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ProcedureOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ProcedureOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteProcedureOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Order on a ProcedureOrder
//----------------------------------------------------------------------------
func AssignOrderToProcedureOrder( procedureOrderId uint64, orderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProcedureOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedureOrder(procedureOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProcedureOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProcedureOrder)

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
			// assign the Order	to the ProcedureOrder
			//----------------------------------------------------------------------------
			parentObj.Order = &childObj

			//----------------------------------------------------------------------------
			// save the ProcedureOrder
			//----------------------------------------------------------------------------
			return UpdateProcedureOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Order", orderId )
			return utils.RequestResult{false, msg, "assignOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Order on a ProcedureOrder
//----------------------------------------------------------------------------
func UnassignOrderFromProcedureOrder(procedureOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProcedureOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedureOrder(procedureOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProcedureOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProcedureOrder)

		//----------------------------------------------------------------------------
		// assign an empty ClinicalOrder to the Order
		//----------------------------------------------------------------------------
		parentObj.Order = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Order
		//----------------------------------------------------------------------------
		parentObj.OrderId = nil;

		//----------------------------------------------------------------------------
		// save the ProcedureOrder
		//----------------------------------------------------------------------------
		return UpdateProcedureOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Facility on a ProcedureOrder
//----------------------------------------------------------------------------
func AssignFacilityToProcedureOrder( procedureOrderId uint64, facilityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProcedureOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedureOrder(procedureOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProcedureOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProcedureOrder)

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
			// assign the Facility	to the ProcedureOrder
			//----------------------------------------------------------------------------
			parentObj.Facility = &childObj

			//----------------------------------------------------------------------------
			// save the ProcedureOrder
			//----------------------------------------------------------------------------
			return UpdateProcedureOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facility", facilityId )
			return utils.RequestResult{false, msg, "assignFacility", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Facility on a ProcedureOrder
//----------------------------------------------------------------------------
func UnassignFacilityFromProcedureOrder(procedureOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProcedureOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedureOrder(procedureOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProcedureOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProcedureOrder)

		//----------------------------------------------------------------------------
		// assign an empty Facility to the Facility
		//----------------------------------------------------------------------------
		parentObj.Facility = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Facility
		//----------------------------------------------------------------------------
		parentObj.FacilityId = nil;

		//----------------------------------------------------------------------------
		// save the ProcedureOrder
		//----------------------------------------------------------------------------
		return UpdateProcedureOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Procedure on a ProcedureOrder
//----------------------------------------------------------------------------
func AssignProcedureToProcedureOrder( procedureOrderId uint64, procedureId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProcedureOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedureOrder(procedureOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProcedureOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProcedureOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Procedure

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Procedure with a
		// matching procedureId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, procedureId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Procedure	to the ProcedureOrder
			//----------------------------------------------------------------------------
			parentObj.Procedure = &childObj

			//----------------------------------------------------------------------------
			// save the ProcedureOrder
			//----------------------------------------------------------------------------
			return UpdateProcedureOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Procedure", procedureId )
			return utils.RequestResult{false, msg, "assignProcedure", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Procedure on a ProcedureOrder
//----------------------------------------------------------------------------
func UnassignProcedureFromProcedureOrder(procedureOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProcedureOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedureOrder(procedureOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProcedureOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProcedureOrder)

		//----------------------------------------------------------------------------
		// assign an empty Procedure to the Procedure
		//----------------------------------------------------------------------------
		parentObj.Procedure = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Procedure
		//----------------------------------------------------------------------------
		parentObj.ProcedureId = nil;

		//----------------------------------------------------------------------------
		// save the ProcedureOrder
		//----------------------------------------------------------------------------
		return UpdateProcedureOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


