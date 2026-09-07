package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LocationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLocation - creates a new db entry
//----------------------------------------------------------------------------
func CreateLocation(obj model.Location)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Location with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Location", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLocation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLocation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLocation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Location

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Location with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Location using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Location using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLocation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLocation - returns all
//----------------------------------------------------------------------------
func GetAllLocation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Location

	//----------------------------------------------------------------------------
	// Request the ORM to find all Location
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Location" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Location", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLocation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLocation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLocation(obj model.Location)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Location using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Location using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLocation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLocation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLocation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Location with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLocation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Location so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Location)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Location using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Location using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLocation", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Warehouse on a Location
//----------------------------------------------------------------------------
func AssignWarehouseToLocation( locationId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Location with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLocation(locationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Location so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Location)

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
			// assign the Warehouse	to the Location
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the Location
			//----------------------------------------------------------------------------
			return UpdateLocation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a Location
//----------------------------------------------------------------------------
func UnassignWarehouseFromLocation(locationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Location with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLocation(locationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Location so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Location)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the Location
		//----------------------------------------------------------------------------
		return UpdateLocation(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more inventoryItemsIds as a InventoryItems to a Location
//----------------------------------------------------------------------------
func AddInventoryItemsToLocation ( locationId uint64, inventoryItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Location with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLocation(locationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Location so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Location)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventoryItemsIds, ",")

		for _, inventoryItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching inventoryItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventoryItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InventoryItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventoryItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItems", inventoryItemsId )
				return utils.RequestResult{false, msg, "unassignInventoryItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Location from the gorm
		//----------------------------------------------------------------------------
		return GetLocation(locationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventoryItemsIds as a InventoryItems from a Location
//----------------------------------------------------------------------------
func RemoveInventoryItemsFromLocation( locationId uint64, inventoryItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Location with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLocation(locationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Location so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Location)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventoryItemsIds, ",")

		for _, inventoryItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching inventoryItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventoryItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InventoryItemObj from the InventoryItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventoryItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItems", inventoryItemsId )
				return utils.RequestResult{false, msg, "removeInventoryItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Location from the gorm
		//----------------------------------------------------------------------------
		return GetLocation(locationId)

	} else {
		return parentRequestResult
	}
}

