package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
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
// assigns a Sku on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignSkuToInventoryTransaction( inventoryTransactionId uint64, skuId uint64 )(utils.RequestResult){

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
		var childObj model.StockKeepingUnit

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StockKeepingUnit with a
		// matching skuId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, skuId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Sku	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignSkuFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

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
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Warehouse on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignWarehouseToInventoryTransaction( inventoryTransactionId uint64, warehouseId uint64 )(utils.RequestResult){

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
		var childObj model.Warehouse

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Warehouse with a
		// matching warehouseId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, warehouseId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Warehouse	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignWarehouseFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

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
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

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
		var childObj model.StorageLocation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StorageLocation with a
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
		// assign an empty StorageLocation to the Location
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
// assigns a Lot on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignLotToInventoryTransaction( inventoryTransactionId uint64, lotId uint64 )(utils.RequestResult){

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
		var childObj model.Lot

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Lot with a
		// matching lotId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, lotId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Lot	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignLotFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

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
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a RelatedReservation on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignRelatedReservationToInventoryTransaction( inventoryTransactionId uint64, relatedReservationId uint64 )(utils.RequestResult){

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
		var childObj model.Reservation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Reservation with a
		// matching relatedReservationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, relatedReservationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the RelatedReservation	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.RelatedReservation = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedReservation", relatedReservationId )
			return utils.RequestResult{false, msg, "assignRelatedReservation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a RelatedReservation on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignRelatedReservationFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

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
		// assign an empty Reservation to the RelatedReservation
		//----------------------------------------------------------------------------
		parentObj.RelatedReservation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the RelatedReservation
		//----------------------------------------------------------------------------
		parentObj.RelatedReservationId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a TransferOrder on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignTransferOrderToInventoryTransaction( inventoryTransactionId uint64, transferOrderId uint64 )(utils.RequestResult){

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
		var childObj model.TransferOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a TransferOrder with a
		// matching transferOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, transferOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the TransferOrder	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.TransferOrder = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TransferOrder", transferOrderId )
			return utils.RequestResult{false, msg, "assignTransferOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TransferOrder on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignTransferOrderFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

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
		// assign an empty TransferOrder to the TransferOrder
		//----------------------------------------------------------------------------
		parentObj.TransferOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TransferOrder
		//----------------------------------------------------------------------------
		parentObj.TransferOrderId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Adjustment on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignAdjustmentToInventoryTransaction( inventoryTransactionId uint64, adjustmentId uint64 )(utils.RequestResult){

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
		var childObj model.StockAdjustment

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StockAdjustment with a
		// matching adjustmentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, adjustmentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Adjustment	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.Adjustment = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Adjustment", adjustmentId )
			return utils.RequestResult{false, msg, "assignAdjustment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Adjustment on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignAdjustmentFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

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
		// assign an empty StockAdjustment to the Adjustment
		//----------------------------------------------------------------------------
		parentObj.Adjustment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Adjustment
		//----------------------------------------------------------------------------
		parentObj.AdjustmentId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CycleCount on a InventoryTransaction
//----------------------------------------------------------------------------
func AssignCycleCountToInventoryTransaction( inventoryTransactionId uint64, cycleCountId uint64 )(utils.RequestResult){

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
		var childObj model.CycleCount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CycleCount with a
		// matching cycleCountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, cycleCountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CycleCount	to the InventoryTransaction
			//----------------------------------------------------------------------------
			parentObj.CycleCount = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryTransaction
			//----------------------------------------------------------------------------
			return UpdateInventoryTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CycleCount", cycleCountId )
			return utils.RequestResult{false, msg, "assignCycleCount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CycleCount on a InventoryTransaction
//----------------------------------------------------------------------------
func UnassignCycleCountFromInventoryTransaction(inventoryTransactionId uint64)(utils.RequestResult) {

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
		// assign an empty CycleCount to the CycleCount
		//----------------------------------------------------------------------------
		parentObj.CycleCount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CycleCount
		//----------------------------------------------------------------------------
		parentObj.CycleCountId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryTransaction
		//----------------------------------------------------------------------------
		return UpdateInventoryTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a InventoryTransaction
//----------------------------------------------------------------------------
func AddSerialNumbersToInventoryTransaction ( inventoryTransactionId uint64, serialNumbersIds string )(utils.RequestResult) {

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
		// retrieve the modified InventoryTransaction from the gorm
		//----------------------------------------------------------------------------
		return GetInventoryTransaction(inventoryTransactionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a InventoryTransaction
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromInventoryTransaction( inventoryTransactionId uint64, serialNumbersIds string )(utils.RequestResult) {
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
		// retrieve the modified InventoryTransaction from the gorm
		//----------------------------------------------------------------------------
		return GetInventoryTransaction(inventoryTransactionId)

	} else {
		return parentRequestResult
	}
}

