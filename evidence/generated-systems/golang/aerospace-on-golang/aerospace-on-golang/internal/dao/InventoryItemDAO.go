package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InventoryItemDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInventoryItem - creates a new db entry
//----------------------------------------------------------------------------
func CreateInventoryItem(obj model.InventoryItem)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InventoryItem with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InventoryItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInventoryItem", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInventoryItem - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInventoryItem(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InventoryItem

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InventoryItem with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InventoryItem using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InventoryItem using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInventoryItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInventoryItem - returns all
//----------------------------------------------------------------------------
func GetAllInventoryItem()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InventoryItem

	//----------------------------------------------------------------------------
	// Request the ORM to find all InventoryItem
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InventoryItem" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InventoryItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInventoryItem", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInventoryItem - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInventoryItem(obj model.InventoryItem)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InventoryItem using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InventoryItem using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInventoryItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInventoryItem - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInventoryItem(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InventoryItem with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInventoryItem(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InventoryItem)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InventoryItem using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InventoryItem using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInventoryItem", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Component on a InventoryItem
//----------------------------------------------------------------------------
func AssignComponentToInventoryItem( inventoryItemId uint64, componentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryItem(inventoryItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Component_

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Component_ with a
		// matching componentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, componentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Component	to the InventoryItem
			//----------------------------------------------------------------------------
			parentObj.Component = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryItem
			//----------------------------------------------------------------------------
			return UpdateInventoryItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Component", componentId )
			return utils.RequestResult{false, msg, "assignComponent", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Component on a InventoryItem
//----------------------------------------------------------------------------
func UnassignComponentFromInventoryItem(inventoryItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryItem(inventoryItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryItem)

		//----------------------------------------------------------------------------
		// assign an empty Component_ to the Component
		//----------------------------------------------------------------------------
		parentObj.Component = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Component
		//----------------------------------------------------------------------------
		parentObj.ComponentId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryItem
		//----------------------------------------------------------------------------
		return UpdateInventoryItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Warehouse on a InventoryItem
//----------------------------------------------------------------------------
func AssignWarehouseToInventoryItem( inventoryItemId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryItem(inventoryItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Warehouse

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Warehouse with a
		// matching warehouseId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, warehouseId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Warehouse	to the InventoryItem
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryItem
			//----------------------------------------------------------------------------
			return UpdateInventoryItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a InventoryItem
//----------------------------------------------------------------------------
func UnassignWarehouseFromInventoryItem(inventoryItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryItem(inventoryItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryItem)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryItem
		//----------------------------------------------------------------------------
		return UpdateInventoryItem(parentObj)

	} else {
		return parentRequestResult
	}

}


