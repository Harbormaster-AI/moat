package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing QuarantineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateQuarantine - creates a new db entry
//----------------------------------------------------------------------------
func CreateQuarantine(obj model.Quarantine)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Quarantine with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Quarantine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateQuarantine", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetQuarantine - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetQuarantine(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Quarantine

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Quarantine with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Quarantine using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Quarantine using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetQuarantine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllQuarantine - returns all
//----------------------------------------------------------------------------
func GetAllQuarantine()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Quarantine

	//----------------------------------------------------------------------------
	// Request the ORM to find all Quarantine
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Quarantine" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Quarantine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllQuarantine", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateQuarantine - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateQuarantine(obj model.Quarantine)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Quarantine using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Quarantine using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateQuarantine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteQuarantine - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteQuarantine(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Quarantine with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetQuarantine(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quarantine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Quarantine)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Quarantine using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Quarantine using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteQuarantine", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Warehouse on a Quarantine
//----------------------------------------------------------------------------
func AssignWarehouseToQuarantine( quarantineId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Quarantine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuarantine(quarantineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quarantine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quarantine)

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
			// assign the Warehouse	to the Quarantine
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the Quarantine
			//----------------------------------------------------------------------------
			return UpdateQuarantine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a Quarantine
//----------------------------------------------------------------------------
func UnassignWarehouseFromQuarantine(quarantineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quarantine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuarantine(quarantineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quarantine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quarantine)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the Quarantine
		//----------------------------------------------------------------------------
		return UpdateQuarantine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lot on a Quarantine
//----------------------------------------------------------------------------
func AssignLotToQuarantine( quarantineId uint64, lotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Quarantine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuarantine(quarantineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quarantine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quarantine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Lot

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Lot with a
		// matching lotId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, lotId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Lot	to the Quarantine
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the Quarantine
			//----------------------------------------------------------------------------
			return UpdateQuarantine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a Quarantine
//----------------------------------------------------------------------------
func UnassignLotFromQuarantine(quarantineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quarantine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuarantine(quarantineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quarantine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quarantine)

		//----------------------------------------------------------------------------
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the Quarantine
		//----------------------------------------------------------------------------
		return UpdateQuarantine(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more itemsIds as a Items to a Quarantine
//----------------------------------------------------------------------------
func AddItemsToQuarantine ( quarantineId uint64, itemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quarantine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuarantine(quarantineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quarantine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quarantine)

		// slice the ids on comma with no spaces
		ids := strings.Split( itemsIds, ",")

		for _, itemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching itemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , itemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Items using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Items").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Items", itemsId )
				return utils.RequestResult{false, msg, "unassignItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Quarantine from the gorm
		//----------------------------------------------------------------------------
		return GetQuarantine(quarantineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more itemsIds as a Items from a Quarantine
//----------------------------------------------------------------------------
func RemoveItemsFromQuarantine( quarantineId uint64, itemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Quarantine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuarantine(quarantineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quarantine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quarantine)

		// slice the ids on comma with no spaces
		ids := strings.Split( itemsIds, ",")

		for _, itemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching itemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , itemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InventoryItemObj from the Items array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Items").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Items", itemsId )
				return utils.RequestResult{false, msg, "removeItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Quarantine from the gorm
		//----------------------------------------------------------------------------
		return GetQuarantine(quarantineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a Quarantine
//----------------------------------------------------------------------------
func AddSerialNumbersToQuarantine ( quarantineId uint64, serialNumbersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quarantine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuarantine(quarantineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quarantine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quarantine)

		// slice the ids on comma with no spaces
		ids := strings.Split( serialNumbersIds, ",")

		for _, serialNumbersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SerialNumber

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SerialNumber
			// with a matching serialNumbersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , serialNumbersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SerialNumbers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SerialNumbers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SerialNumbers", serialNumbersId )
				return utils.RequestResult{false, msg, "unassignSerialNumbers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Quarantine from the gorm
		//----------------------------------------------------------------------------
		return GetQuarantine(quarantineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a Quarantine
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromQuarantine( quarantineId uint64, serialNumbersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Quarantine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuarantine(quarantineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quarantine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quarantine)

		// slice the ids on comma with no spaces
		ids := strings.Split( serialNumbersIds, ",")

		for _, serialNumbersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SerialNumber

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SerialNumber
			// with a matching serialNumbersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , serialNumbersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SerialNumberObj from the SerialNumbers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SerialNumbers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SerialNumbers", serialNumbersId )
				return utils.RequestResult{false, msg, "removeSerialNumbers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Quarantine from the gorm
		//----------------------------------------------------------------------------
		return GetQuarantine(quarantineId)

	} else {
		return parentRequestResult
	}
}

