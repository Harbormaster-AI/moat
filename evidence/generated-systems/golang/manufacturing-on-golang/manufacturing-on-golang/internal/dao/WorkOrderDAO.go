package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing WorkOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateWorkOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateWorkOrder(obj model.WorkOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a WorkOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a WorkOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateWorkOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetWorkOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetWorkOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.WorkOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a WorkOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a WorkOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a WorkOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetWorkOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllWorkOrder - returns all
//----------------------------------------------------------------------------
func GetAllWorkOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.WorkOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all WorkOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all WorkOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all WorkOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllWorkOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateWorkOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateWorkOrder(obj model.WorkOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a WorkOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a WorkOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateWorkOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteWorkOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteWorkOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetWorkOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a WorkOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a WorkOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteWorkOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Item on a WorkOrder
//----------------------------------------------------------------------------
func AssignItemToWorkOrder( workOrderId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching itemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, itemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Item	to the WorkOrder
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the WorkOrder
			//----------------------------------------------------------------------------
			return UpdateWorkOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a WorkOrder
//----------------------------------------------------------------------------
func UnassignItemFromWorkOrder(workOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the WorkOrder
		//----------------------------------------------------------------------------
		return UpdateWorkOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Plant on a WorkOrder
//----------------------------------------------------------------------------
func AssignPlantToWorkOrder( workOrderId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Plant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Plant with a
		// matching plantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, plantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Plant	to the WorkOrder
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the WorkOrder
			//----------------------------------------------------------------------------
			return UpdateWorkOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a WorkOrder
//----------------------------------------------------------------------------
func UnassignPlantFromWorkOrder(workOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the WorkOrder
		//----------------------------------------------------------------------------
		return UpdateWorkOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Routing on a WorkOrder
//----------------------------------------------------------------------------
func AssignRoutingToWorkOrder( workOrderId uint64, routingId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

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
			// assign the Routing	to the WorkOrder
			//----------------------------------------------------------------------------
			parentObj.Routing = &childObj

			//----------------------------------------------------------------------------
			// save the WorkOrder
			//----------------------------------------------------------------------------
			return UpdateWorkOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Routing", routingId )
			return utils.RequestResult{false, msg, "assignRouting", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Routing on a WorkOrder
//----------------------------------------------------------------------------
func UnassignRoutingFromWorkOrder(workOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// assign an empty Routing to the Routing
		//----------------------------------------------------------------------------
		parentObj.Routing = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Routing
		//----------------------------------------------------------------------------
		parentObj.RoutingId = nil;

		//----------------------------------------------------------------------------
		// save the WorkOrder
		//----------------------------------------------------------------------------
		return UpdateWorkOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Bom on a WorkOrder
//----------------------------------------------------------------------------
func AssignBomToWorkOrder( workOrderId uint64, bomId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BOM

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BOM with a
		// matching bomId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, bomId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Bom	to the WorkOrder
			//----------------------------------------------------------------------------
			parentObj.Bom = &childObj

			//----------------------------------------------------------------------------
			// save the WorkOrder
			//----------------------------------------------------------------------------
			return UpdateWorkOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Bom", bomId )
			return utils.RequestResult{false, msg, "assignBom", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Bom on a WorkOrder
//----------------------------------------------------------------------------
func UnassignBomFromWorkOrder(workOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// assign an empty BOM to the Bom
		//----------------------------------------------------------------------------
		parentObj.Bom = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Bom
		//----------------------------------------------------------------------------
		parentObj.BomId = nil;

		//----------------------------------------------------------------------------
		// save the WorkOrder
		//----------------------------------------------------------------------------
		return UpdateWorkOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ProductionSchedule on a WorkOrder
//----------------------------------------------------------------------------
func AssignProductionScheduleToWorkOrder( workOrderId uint64, productionScheduleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ProductionSchedule

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ProductionSchedule with a
		// matching productionScheduleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productionScheduleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ProductionSchedule	to the WorkOrder
			//----------------------------------------------------------------------------
			parentObj.ProductionSchedule = &childObj

			//----------------------------------------------------------------------------
			// save the WorkOrder
			//----------------------------------------------------------------------------
			return UpdateWorkOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionSchedule", productionScheduleId )
			return utils.RequestResult{false, msg, "assignProductionSchedule", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ProductionSchedule on a WorkOrder
//----------------------------------------------------------------------------
func UnassignProductionScheduleFromWorkOrder(workOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// assign an empty ProductionSchedule to the ProductionSchedule
		//----------------------------------------------------------------------------
		parentObj.ProductionSchedule = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ProductionSchedule
		//----------------------------------------------------------------------------
		parentObj.ProductionScheduleId = nil;

		//----------------------------------------------------------------------------
		// save the WorkOrder
		//----------------------------------------------------------------------------
		return UpdateWorkOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a SalesOrder on a WorkOrder
//----------------------------------------------------------------------------
func AssignSalesOrderToWorkOrder( workOrderId uint64, salesOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.SalesOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a SalesOrder with a
		// matching salesOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, salesOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SalesOrder	to the WorkOrder
			//----------------------------------------------------------------------------
			parentObj.SalesOrder = &childObj

			//----------------------------------------------------------------------------
			// save the WorkOrder
			//----------------------------------------------------------------------------
			return UpdateWorkOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalesOrder", salesOrderId )
			return utils.RequestResult{false, msg, "assignSalesOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SalesOrder on a WorkOrder
//----------------------------------------------------------------------------
func UnassignSalesOrderFromWorkOrder(workOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkOrder(workOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkOrder)

		//----------------------------------------------------------------------------
		// assign an empty SalesOrder to the SalesOrder
		//----------------------------------------------------------------------------
		parentObj.SalesOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SalesOrder
		//----------------------------------------------------------------------------
		parentObj.SalesOrderId = nil;

		//----------------------------------------------------------------------------
		// save the WorkOrder
		//----------------------------------------------------------------------------
		return UpdateWorkOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


