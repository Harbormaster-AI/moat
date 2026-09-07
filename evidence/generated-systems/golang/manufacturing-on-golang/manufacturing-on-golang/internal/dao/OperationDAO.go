package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OperationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOperation - creates a new db entry
//----------------------------------------------------------------------------
func CreateOperation(obj model.Operation)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Operation with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Operation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOperation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOperation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOperation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Operation

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Operation with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Operation using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Operation using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOperation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOperation - returns all
//----------------------------------------------------------------------------
func GetAllOperation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Operation

	//----------------------------------------------------------------------------
	// Request the ORM to find all Operation
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Operation" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Operation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOperation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOperation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOperation(obj model.Operation)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Operation using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Operation using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOperation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOperation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOperation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Operation with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOperation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Operation)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Operation using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Operation using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOperation", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Routing on a Operation
//----------------------------------------------------------------------------
func AssignRoutingToOperation( operationId uint64, routingId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Operation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperation(operationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Routing

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Routing with a
		// matching routingId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, routingId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Routing	to the Operation
			//----------------------------------------------------------------------------
			parentObj.Routing = &childObj

			//----------------------------------------------------------------------------
			// save the Operation
			//----------------------------------------------------------------------------
			return UpdateOperation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Routing", routingId )
			return utils.RequestResult{false, msg, "assignRouting", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Routing on a Operation
//----------------------------------------------------------------------------
func UnassignRoutingFromOperation(operationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Operation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperation(operationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operation)

		//----------------------------------------------------------------------------
		// assign an empty Routing to the Routing
		//----------------------------------------------------------------------------
		parentObj.Routing = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Routing
		//----------------------------------------------------------------------------
		parentObj.RoutingId = nil;

		//----------------------------------------------------------------------------
		// save the Operation
		//----------------------------------------------------------------------------
		return UpdateOperation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a WorkCenter on a Operation
//----------------------------------------------------------------------------
func AssignWorkCenterToOperation( operationId uint64, workCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Operation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperation(operationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.WorkCenter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a WorkCenter with a
		// matching workCenterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workCenterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the WorkCenter	to the Operation
			//----------------------------------------------------------------------------
			parentObj.WorkCenter = &childObj

			//----------------------------------------------------------------------------
			// save the Operation
			//----------------------------------------------------------------------------
			return UpdateOperation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkCenter", workCenterId )
			return utils.RequestResult{false, msg, "assignWorkCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkCenter on a Operation
//----------------------------------------------------------------------------
func UnassignWorkCenterFromOperation(operationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Operation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperation(operationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operation)

		//----------------------------------------------------------------------------
		// assign an empty WorkCenter to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenterId = nil;

		//----------------------------------------------------------------------------
		// save the Operation
		//----------------------------------------------------------------------------
		return UpdateOperation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a InspectionPlan on a Operation
//----------------------------------------------------------------------------
func AssignInspectionPlanToOperation( operationId uint64, inspectionPlanId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Operation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperation(operationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InspectionPlan

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InspectionPlan with a
		// matching inspectionPlanId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, inspectionPlanId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InspectionPlan	to the Operation
			//----------------------------------------------------------------------------
			parentObj.InspectionPlan = &childObj

			//----------------------------------------------------------------------------
			// save the Operation
			//----------------------------------------------------------------------------
			return UpdateOperation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InspectionPlan", inspectionPlanId )
			return utils.RequestResult{false, msg, "assignInspectionPlan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InspectionPlan on a Operation
//----------------------------------------------------------------------------
func UnassignInspectionPlanFromOperation(operationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Operation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperation(operationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operation)

		//----------------------------------------------------------------------------
		// assign an empty InspectionPlan to the InspectionPlan
		//----------------------------------------------------------------------------
		parentObj.InspectionPlan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InspectionPlan
		//----------------------------------------------------------------------------
		parentObj.InspectionPlanId = nil;

		//----------------------------------------------------------------------------
		// save the Operation
		//----------------------------------------------------------------------------
		return UpdateOperation(parentObj)

	} else {
		return parentRequestResult
	}

}


