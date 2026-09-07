package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InventoryTransactionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInventoryTransaction - creates a new db entry
//----------------------------------------------------------------------------
func CreateInventoryTransaction(obj model.InventoryTransaction)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InventoryTransaction with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InventoryTransaction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInventoryTransaction", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInventoryTransaction - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInventoryTransaction(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InventoryTransaction

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InventoryTransaction with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InventoryTransaction using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InventoryTransaction using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInventoryTransaction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInventoryTransaction - returns all
//----------------------------------------------------------------------------
func GetAllInventoryTransaction()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InventoryTransaction

	//----------------------------------------------------------------------------
	// Request the ORM to find all InventoryTransaction
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InventoryTransaction" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InventoryTransaction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInventoryTransaction", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInventoryTransaction - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInventoryTransaction(obj model.InventoryTransaction)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InventoryTransaction using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InventoryTransaction using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInventoryTransaction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInventoryTransaction - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInventoryTransaction(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInventoryTransaction(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InventoryTransaction)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InventoryTransaction using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InventoryTransaction using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInventoryTransaction", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Item on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignItemToInventoryTransaction( inventoryTransactionId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

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
			// assign the Item	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignItemFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Location on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignLocationToInventoryTransaction( inventoryTransactionId uint64, locationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Location

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Location with a
		// matching locationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, locationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Location	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.Location = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Location", locationId )
			return utils.RequestResult{false, msg, "assignLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Location on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignLocationFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

		//----------------------------------------------------------------------------
		// assign an empty Location to the Location
		//----------------------------------------------------------------------------
		parentObj.Location = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Location
		//----------------------------------------------------------------------------
		parentObj.LocationId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a WorkOrder on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignWorkOrderToInventoryTransaction( inventoryTransactionId uint64, workOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.WorkOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a WorkOrder with a
		// matching workOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the WorkOrder	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.WorkOrder = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrder", workOrderId )
			return utils.RequestResult{false, msg, "assignWorkOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkOrder on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignWorkOrderFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

		//----------------------------------------------------------------------------
		// assign an empty WorkOrder to the WorkOrder
		//----------------------------------------------------------------------------
		parentObj.WorkOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkOrder
		//----------------------------------------------------------------------------
		parentObj.WorkOrderId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PurchaseOrder on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignPurchaseOrderToInventoryTransaction( inventoryTransactionId uint64, purchaseOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PurchaseOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PurchaseOrder with a
		// matching purchaseOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, purchaseOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PurchaseOrder	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.PurchaseOrder = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PurchaseOrder", purchaseOrderId )
			return utils.RequestResult{false, msg, "assignPurchaseOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PurchaseOrder on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignPurchaseOrderFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

		//----------------------------------------------------------------------------
		// assign an empty PurchaseOrder to the PurchaseOrder
		//----------------------------------------------------------------------------
		parentObj.PurchaseOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PurchaseOrder
		//----------------------------------------------------------------------------
		parentObj.PurchaseOrderId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a SalesOrder on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignSalesOrderToInventoryTransaction( inventoryTransactionId uint64, salesOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

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
			// assign the SalesOrder	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.SalesOrder = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalesOrder", salesOrderId )
			return utils.RequestResult{false, msg, "assignSalesOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SalesOrder on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignSalesOrderFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryTransaction(inventoryTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryTransaction)

		//----------------------------------------------------------------------------
		// assign an empty SalesOrder to the SalesOrder
		//----------------------------------------------------------------------------
		parentObj.SalesOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SalesOrder
		//----------------------------------------------------------------------------
		parentObj.SalesOrderId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}


