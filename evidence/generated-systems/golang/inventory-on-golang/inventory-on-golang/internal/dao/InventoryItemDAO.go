package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
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
// assigns a Sku on a InventoryItem
//----------------------------------------------------------------------------
func AssignSkuToInventoryItem( inventoryItemId uint64, skuId uint64 )(utils.RequestResult){

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
		var childObj model.StockKeepingUnit

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StockKeepingUnit with a
		// matching skuId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, skuId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Sku	to the InventoryItem
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryItem
			//----------------------------------------------------------------------------
			return UpdateInventoryItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a InventoryItem
//----------------------------------------------------------------------------
func UnassignSkuFromInventoryItem(inventoryItemId uint64)(utils.RequestResult) {

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
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

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
		var childObj model.StorageLocation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StorageLocation with a
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
		// assign an empty StorageLocation to the Location
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

//----------------------------------------------------------------------------
// assigns a Lot on a InventoryItem
//----------------------------------------------------------------------------
func AssignLotToInventoryItem( inventoryItemId uint64, lotId uint64 )(utils.RequestResult){

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
		var childObj model.Lot

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Lot with a
		// matching lotId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, lotId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Lot	to the InventoryItem
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryItem
			//----------------------------------------------------------------------------
			return UpdateInventoryItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a InventoryItem
//----------------------------------------------------------------------------
func UnassignLotFromInventoryItem(inventoryItemId uint64)(utils.RequestResult) {

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
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryItem
		//----------------------------------------------------------------------------
		return UpdateInventoryItem(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a InventoryItem
//----------------------------------------------------------------------------
func AddSerialNumbersToInventoryItem ( inventoryItemId uint64, serialNumbersIds string )(utils.RequestResult) {

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
		// retrieve the modified InventoryItem from the gorm
		//----------------------------------------------------------------------------
		return GetInventoryItem(inventoryItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a InventoryItem
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromInventoryItem( inventoryItemId uint64, serialNumbersIds string )(utils.RequestResult) {
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
		// retrieve the modified InventoryItem from the gorm
		//----------------------------------------------------------------------------
		return GetInventoryItem(inventoryItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a InventoryItem
//----------------------------------------------------------------------------
func AddTransactionsToInventoryItem ( inventoryItemId uint64, transactionsIds string )(utils.RequestResult) {

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

		// slice the ids on comma with no spaces
		ids := strings.Split( transactionsIds, ",")

		for _, transactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryTransaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryTransaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Transactions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "unassignTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InventoryItem from the gorm
		//----------------------------------------------------------------------------
		return GetInventoryItem(inventoryItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a InventoryItem
//----------------------------------------------------------------------------
func RemoveTransactionsFromInventoryItem( inventoryItemId uint64, transactionsIds string )(utils.RequestResult) {
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

		// slice the ids on comma with no spaces
		ids := strings.Split( transactionsIds, ",")

		for _, transactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryTransaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryTransaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InventoryTransactionObj from the Transactions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "removeTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InventoryItem from the gorm
		//----------------------------------------------------------------------------
		return GetInventoryItem(inventoryItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reservationsIds as a Reservations to a InventoryItem
//----------------------------------------------------------------------------
func AddReservationsToInventoryItem ( inventoryItemId uint64, reservationsIds string )(utils.RequestResult) {

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

		// slice the ids on comma with no spaces
		ids := strings.Split( reservationsIds, ",")

		for _, reservationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Reservation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Reservation
			// with a matching reservationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reservationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reservations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reservations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reservations", reservationsId )
				return utils.RequestResult{false, msg, "unassignReservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InventoryItem from the gorm
		//----------------------------------------------------------------------------
		return GetInventoryItem(inventoryItemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reservationsIds as a Reservations from a InventoryItem
//----------------------------------------------------------------------------
func RemoveReservationsFromInventoryItem( inventoryItemId uint64, reservationsIds string )(utils.RequestResult) {
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

		// slice the ids on comma with no spaces
		ids := strings.Split( reservationsIds, ",")

		for _, reservationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Reservation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Reservation
			// with a matching reservationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reservationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ReservationObj from the Reservations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reservations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reservations", reservationsId )
				return utils.RequestResult{false, msg, "removeReservations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InventoryItem from the gorm
		//----------------------------------------------------------------------------
		return GetInventoryItem(inventoryItemId)

	} else {
		return parentRequestResult
	}
}

