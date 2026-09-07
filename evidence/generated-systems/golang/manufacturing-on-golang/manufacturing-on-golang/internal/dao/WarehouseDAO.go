package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing WarehouseDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateWarehouse - creates a new db entry
//----------------------------------------------------------------------------
func CreateWarehouse(obj model.Warehouse)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Warehouse with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Warehouse", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateWarehouse", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetWarehouse - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetWarehouse(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Warehouse

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Warehouse with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Warehouse using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Warehouse using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetWarehouse", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllWarehouse - returns all
//----------------------------------------------------------------------------
func GetAllWarehouse()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Warehouse

	//----------------------------------------------------------------------------
	// Request the ORM to find all Warehouse
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Warehouse" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Warehouse", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllWarehouse", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateWarehouse - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateWarehouse(obj model.Warehouse)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Warehouse using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Warehouse using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateWarehouse", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteWarehouse - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteWarehouse(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetWarehouse(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Warehouse)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Warehouse using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Warehouse using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteWarehouse", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Plant on a Warehouse
//----------------------------------------------------------------------------
func AssignPlantToWarehouse( warehouseId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

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
			// assign the Plant	to the Warehouse
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the Warehouse
			//----------------------------------------------------------------------------
			return UpdateWarehouse(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a Warehouse
//----------------------------------------------------------------------------
func UnassignPlantFromWarehouse(warehouseId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the Warehouse
		//----------------------------------------------------------------------------
		return UpdateWarehouse(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more locationsIds as a Locations to a Warehouse
//----------------------------------------------------------------------------
func AddLocationsToWarehouse ( warehouseId uint64, locationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( locationsIds, ",")

		for _, locationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Location

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Location
			// with a matching locationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , locationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Locations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Locations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Locations", locationsId )
				return utils.RequestResult{false, msg, "unassignLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more locationsIds as a Locations from a Warehouse
//----------------------------------------------------------------------------
func RemoveLocationsFromWarehouse( warehouseId uint64, locationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

		// slice the ids on comma with no spaces
		ids := strings.Split( locationsIds, ",")

		for _, locationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Location

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Location
			// with a matching locationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , locationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LocationObj from the Locations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Locations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Locations", locationsId )
				return utils.RequestResult{false, msg, "removeLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more inventoryItemsIds as a InventoryItems to a Warehouse
//----------------------------------------------------------------------------
func AddInventoryItemsToWarehouse ( warehouseId uint64, inventoryItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

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
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventoryItemsIds as a InventoryItems from a Warehouse
//----------------------------------------------------------------------------
func RemoveInventoryItemsFromWarehouse( warehouseId uint64, inventoryItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Warehouse with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWarehouse(warehouseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Warehouse so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Warehouse)

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
		// retrieve the modified Warehouse from the gorm
		//----------------------------------------------------------------------------
		return GetWarehouse(warehouseId)

	} else {
		return parentRequestResult
	}
}

