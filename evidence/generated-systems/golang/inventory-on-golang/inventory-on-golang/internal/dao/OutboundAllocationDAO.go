package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OutboundAllocationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOutboundAllocation - creates a new db entry
//----------------------------------------------------------------------------
func CreateOutboundAllocation(obj model.OutboundAllocation)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a OutboundAllocation with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a OutboundAllocation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOutboundAllocation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOutboundAllocation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOutboundAllocation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.OutboundAllocation

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a OutboundAllocation with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a OutboundAllocation using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a OutboundAllocation using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOutboundAllocation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOutboundAllocation - returns all
//----------------------------------------------------------------------------
func GetAllOutboundAllocation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.OutboundAllocation

	//----------------------------------------------------------------------------
	// Request the ORM to find all OutboundAllocation
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all OutboundAllocation" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all OutboundAllocation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOutboundAllocation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOutboundAllocation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOutboundAllocation(obj model.OutboundAllocation)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a OutboundAllocation using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a OutboundAllocation using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOutboundAllocation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOutboundAllocation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOutboundAllocation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOutboundAllocation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a OutboundAllocation using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a OutboundAllocation using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOutboundAllocation", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Warehouse on a OutboundAllocation
//----------------------------------------------------------------------------
func AssignWarehouseToOutboundAllocation( outboundAllocationId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

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
			// assign the Warehouse	to the OutboundAllocation
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the OutboundAllocation
			//----------------------------------------------------------------------------
			return UpdateOutboundAllocation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a OutboundAllocation
//----------------------------------------------------------------------------
func UnassignWarehouseFromOutboundAllocation(outboundAllocationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the OutboundAllocation
		//----------------------------------------------------------------------------
		return UpdateOutboundAllocation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Sku on a OutboundAllocation
//----------------------------------------------------------------------------
func AssignSkuToOutboundAllocation( outboundAllocationId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

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
			// assign the Sku	to the OutboundAllocation
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the OutboundAllocation
			//----------------------------------------------------------------------------
			return UpdateOutboundAllocation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a OutboundAllocation
//----------------------------------------------------------------------------
func UnassignSkuFromOutboundAllocation(outboundAllocationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the OutboundAllocation
		//----------------------------------------------------------------------------
		return UpdateOutboundAllocation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a InventoryItem on a OutboundAllocation
//----------------------------------------------------------------------------
func AssignInventoryItemToOutboundAllocation( outboundAllocationId uint64, inventoryItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InventoryItem

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InventoryItem with a
		// matching inventoryItemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, inventoryItemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InventoryItem	to the OutboundAllocation
			//----------------------------------------------------------------------------
			parentObj.InventoryItem = &childObj

			//----------------------------------------------------------------------------
			// save the OutboundAllocation
			//----------------------------------------------------------------------------
			return UpdateOutboundAllocation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItem", inventoryItemId )
			return utils.RequestResult{false, msg, "assignInventoryItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InventoryItem on a OutboundAllocation
//----------------------------------------------------------------------------
func UnassignInventoryItemFromOutboundAllocation(outboundAllocationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// assign an empty InventoryItem to the InventoryItem
		//----------------------------------------------------------------------------
		parentObj.InventoryItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InventoryItem
		//----------------------------------------------------------------------------
		parentObj.InventoryItemId = nil;

		//----------------------------------------------------------------------------
		// save the OutboundAllocation
		//----------------------------------------------------------------------------
		return UpdateOutboundAllocation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Reservation on a OutboundAllocation
//----------------------------------------------------------------------------
func AssignReservationToOutboundAllocation( outboundAllocationId uint64, reservationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Reservation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Reservation with a
		// matching reservationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, reservationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Reservation	to the OutboundAllocation
			//----------------------------------------------------------------------------
			parentObj.Reservation = &childObj

			//----------------------------------------------------------------------------
			// save the OutboundAllocation
			//----------------------------------------------------------------------------
			return UpdateOutboundAllocation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reservation", reservationId )
			return utils.RequestResult{false, msg, "assignReservation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Reservation on a OutboundAllocation
//----------------------------------------------------------------------------
func UnassignReservationFromOutboundAllocation(outboundAllocationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// assign an empty Reservation to the Reservation
		//----------------------------------------------------------------------------
		parentObj.Reservation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Reservation
		//----------------------------------------------------------------------------
		parentObj.ReservationId = nil;

		//----------------------------------------------------------------------------
		// save the OutboundAllocation
		//----------------------------------------------------------------------------
		return UpdateOutboundAllocation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lot on a OutboundAllocation
//----------------------------------------------------------------------------
func AssignLotToOutboundAllocation( outboundAllocationId uint64, lotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

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
			// assign the Lot	to the OutboundAllocation
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the OutboundAllocation
			//----------------------------------------------------------------------------
			return UpdateOutboundAllocation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a OutboundAllocation
//----------------------------------------------------------------------------
func UnassignLotFromOutboundAllocation(outboundAllocationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the OutboundAllocation
		//----------------------------------------------------------------------------
		return UpdateOutboundAllocation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a SourceLocation on a OutboundAllocation
//----------------------------------------------------------------------------
func AssignSourceLocationToOutboundAllocation( outboundAllocationId uint64, sourceLocationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.StorageLocation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StorageLocation with a
		// matching sourceLocationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, sourceLocationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SourceLocation	to the OutboundAllocation
			//----------------------------------------------------------------------------
			parentObj.SourceLocation = &childObj

			//----------------------------------------------------------------------------
			// save the OutboundAllocation
			//----------------------------------------------------------------------------
			return UpdateOutboundAllocation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SourceLocation", sourceLocationId )
			return utils.RequestResult{false, msg, "assignSourceLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SourceLocation on a OutboundAllocation
//----------------------------------------------------------------------------
func UnassignSourceLocationFromOutboundAllocation(outboundAllocationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the SourceLocation
		//----------------------------------------------------------------------------
		parentObj.SourceLocation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SourceLocation
		//----------------------------------------------------------------------------
		parentObj.SourceLocationId = nil;

		//----------------------------------------------------------------------------
		// save the OutboundAllocation
		//----------------------------------------------------------------------------
		return UpdateOutboundAllocation(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a OutboundAllocation
//----------------------------------------------------------------------------
func AddSerialNumbersToOutboundAllocation ( outboundAllocationId uint64, serialNumbersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

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
		// retrieve the modified OutboundAllocation from the gorm
		//----------------------------------------------------------------------------
		return GetOutboundAllocation(outboundAllocationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a OutboundAllocation
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromOutboundAllocation( outboundAllocationId uint64, serialNumbersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the OutboundAllocation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOutboundAllocation(outboundAllocationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OutboundAllocation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OutboundAllocation)

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
		// retrieve the modified OutboundAllocation from the gorm
		//----------------------------------------------------------------------------
		return GetOutboundAllocation(outboundAllocationId)

	} else {
		return parentRequestResult
	}
}

