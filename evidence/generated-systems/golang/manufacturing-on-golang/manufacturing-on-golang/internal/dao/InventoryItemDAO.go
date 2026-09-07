package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
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
// assigns a Item on a InventoryItem
//----------------------------------------------------------------------------
func AssignItemToInventoryItem( inventoryItemId uint64, itemId uint64 )(utils.RequestResult){

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
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching itemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, itemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Item	to the InventoryItem
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryItem
			//----------------------------------------------------------------------------
			return UpdateInventoryItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a InventoryItem
//----------------------------------------------------------------------------
func UnassignItemFromInventoryItem(inventoryItemId uint64)(utils.RequestResult) {

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
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryItem
		//----------------------------------------------------------------------------
		return UpdateInventoryItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Location on a InventoryItem
//----------------------------------------------------------------------------
func AssignLocationToInventoryItem( inventoryItemId uint64, locationId uint64 )(utils.RequestResult){

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
		var childObj model.Location

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Location with a
		// matching locationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, locationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Location	to the InventoryItem
			//----------------------------------------------------------------------------
			parentObj.Location = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryItem
			//----------------------------------------------------------------------------
			return UpdateInventoryItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Location", locationId )
			return utils.RequestResult{false, msg, "assignLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Location on a InventoryItem
//----------------------------------------------------------------------------
func UnassignLocationFromInventoryItem(inventoryItemId uint64)(utils.RequestResult) {

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
		// assign an empty Location to the Location
		//----------------------------------------------------------------------------
		parentObj.Location = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Location
		//----------------------------------------------------------------------------
		parentObj.LocationId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryItem
		//----------------------------------------------------------------------------
		return UpdateInventoryItem(parentObj)

	} else {
		return parentRequestResult
	}

}


