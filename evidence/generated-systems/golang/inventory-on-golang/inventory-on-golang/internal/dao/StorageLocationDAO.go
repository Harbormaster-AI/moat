package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing StorageLocationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateStorageLocation - creates a new db entry
//----------------------------------------------------------------------------
func CreateStorageLocation(obj model.StorageLocation)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a StorageLocation with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a StorageLocation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateStorageLocation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetStorageLocation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetStorageLocation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.StorageLocation

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a StorageLocation with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a StorageLocation using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a StorageLocation using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetStorageLocation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllStorageLocation - returns all
//----------------------------------------------------------------------------
func GetAllStorageLocation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.StorageLocation

	//----------------------------------------------------------------------------
	// Request the ORM to find all StorageLocation
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all StorageLocation" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all StorageLocation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllStorageLocation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateStorageLocation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateStorageLocation(obj model.StorageLocation)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a StorageLocation using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a StorageLocation using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateStorageLocation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteStorageLocation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteStorageLocation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the StorageLocation with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetStorageLocation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StorageLocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.StorageLocation)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a StorageLocation using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a StorageLocation using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteStorageLocation", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Warehouse on a StorageLocation
//----------------------------------------------------------------------------
func AssignWarehouseToStorageLocation( storageLocationId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StorageLocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStorageLocation(storageLocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StorageLocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StorageLocation)

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
			// assign the Warehouse	to the StorageLocation
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the StorageLocation
			//----------------------------------------------------------------------------
			return UpdateStorageLocation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a StorageLocation
//----------------------------------------------------------------------------
func UnassignWarehouseFromStorageLocation(storageLocationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StorageLocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStorageLocation(storageLocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StorageLocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StorageLocation)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the StorageLocation
		//----------------------------------------------------------------------------
		return UpdateStorageLocation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ParentLocation on a StorageLocation
//----------------------------------------------------------------------------
func AssignParentLocationToStorageLocation( storageLocationId uint64, parentLocationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StorageLocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStorageLocation(storageLocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StorageLocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StorageLocation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.StorageLocation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StorageLocation with a
		// matching parentLocationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, parentLocationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ParentLocation	to the StorageLocation
			//----------------------------------------------------------------------------
			parentObj.ParentLocation = &childObj

			//----------------------------------------------------------------------------
			// save the StorageLocation
			//----------------------------------------------------------------------------
			return UpdateStorageLocation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ParentLocation", parentLocationId )
			return utils.RequestResult{false, msg, "assignParentLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ParentLocation on a StorageLocation
//----------------------------------------------------------------------------
func UnassignParentLocationFromStorageLocation(storageLocationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StorageLocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStorageLocation(storageLocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StorageLocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StorageLocation)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the ParentLocation
		//----------------------------------------------------------------------------
		parentObj.ParentLocation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ParentLocation
		//----------------------------------------------------------------------------
		parentObj.ParentLocationId = nil;

		//----------------------------------------------------------------------------
		// save the StorageLocation
		//----------------------------------------------------------------------------
		return UpdateStorageLocation(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more childLocationsIds as a ChildLocations to a StorageLocation
//----------------------------------------------------------------------------
func AddChildLocationsToStorageLocation ( storageLocationId uint64, childLocationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StorageLocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStorageLocation(storageLocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StorageLocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StorageLocation)

		// slice the ids on comma with no spaces
		ids := strings.Split( childLocationsIds, ",")

		for _, childLocationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StorageLocation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StorageLocation
			// with a matching childLocationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childLocationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ChildLocations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ChildLocations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ChildLocations", childLocationsId )
				return utils.RequestResult{false, msg, "unassignChildLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StorageLocation from the gorm
		//----------------------------------------------------------------------------
		return GetStorageLocation(storageLocationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more childLocationsIds as a ChildLocations from a StorageLocation
//----------------------------------------------------------------------------
func RemoveChildLocationsFromStorageLocation( storageLocationId uint64, childLocationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StorageLocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStorageLocation(storageLocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StorageLocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StorageLocation)

		// slice the ids on comma with no spaces
		ids := strings.Split( childLocationsIds, ",")

		for _, childLocationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StorageLocation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StorageLocation
			// with a matching childLocationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , childLocationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove StorageLocationObj from the ChildLocations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ChildLocations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ChildLocations", childLocationsId )
				return utils.RequestResult{false, msg, "removeChildLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StorageLocation from the gorm
		//----------------------------------------------------------------------------
		return GetStorageLocation(storageLocationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more inventoryItemsIds as a InventoryItems to a StorageLocation
//----------------------------------------------------------------------------
func AddInventoryItemsToStorageLocation ( storageLocationId uint64, inventoryItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StorageLocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStorageLocation(storageLocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StorageLocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StorageLocation)

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
		// retrieve the modified StorageLocation from the gorm
		//----------------------------------------------------------------------------
		return GetStorageLocation(storageLocationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventoryItemsIds as a InventoryItems from a StorageLocation
//----------------------------------------------------------------------------
func RemoveInventoryItemsFromStorageLocation( storageLocationId uint64, inventoryItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StorageLocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStorageLocation(storageLocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StorageLocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StorageLocation)

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
		// retrieve the modified StorageLocation from the gorm
		//----------------------------------------------------------------------------
		return GetStorageLocation(storageLocationId)

	} else {
		return parentRequestResult
	}
}

