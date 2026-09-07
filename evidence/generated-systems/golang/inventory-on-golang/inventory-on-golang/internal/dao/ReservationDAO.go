package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ReservationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateReservation - creates a new db entry
//----------------------------------------------------------------------------
func CreateReservation(obj model.Reservation)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Reservation with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Reservation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateReservation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetReservation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetReservation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Reservation

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Reservation with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Reservation using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Reservation using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetReservation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllReservation - returns all
//----------------------------------------------------------------------------
func GetAllReservation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Reservation

	//----------------------------------------------------------------------------
	// Request the ORM to find all Reservation
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Reservation" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Reservation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllReservation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateReservation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateReservation(obj model.Reservation)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Reservation using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Reservation using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateReservation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteReservation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteReservation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetReservation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Reservation)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Reservation using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Reservation using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteReservation", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Sku on a Reservation
//----------------------------------------------------------------------------
func AssignSkuToReservation( reservationId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

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
			// assign the Sku	to the Reservation
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the Reservation
			//----------------------------------------------------------------------------
			return UpdateReservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a Reservation
//----------------------------------------------------------------------------
func UnassignSkuFromReservation(reservationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the Reservation
		//----------------------------------------------------------------------------
		return UpdateReservation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Warehouse on a Reservation
//----------------------------------------------------------------------------
func AssignWarehouseToReservation( reservationId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

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
			// assign the Warehouse	to the Reservation
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the Reservation
			//----------------------------------------------------------------------------
			return UpdateReservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a Reservation
//----------------------------------------------------------------------------
func UnassignWarehouseFromReservation(reservationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the Reservation
		//----------------------------------------------------------------------------
		return UpdateReservation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Location on a Reservation
//----------------------------------------------------------------------------
func AssignLocationToReservation( reservationId uint64, locationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

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
			// assign the Location	to the Reservation
			//----------------------------------------------------------------------------
			parentObj.Location = &childObj

			//----------------------------------------------------------------------------
			// save the Reservation
			//----------------------------------------------------------------------------
			return UpdateReservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Location", locationId )
			return utils.RequestResult{false, msg, "assignLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Location on a Reservation
//----------------------------------------------------------------------------
func UnassignLocationFromReservation(reservationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the Location
		//----------------------------------------------------------------------------
		parentObj.Location = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Location
		//----------------------------------------------------------------------------
		parentObj.LocationId = nil;

		//----------------------------------------------------------------------------
		// save the Reservation
		//----------------------------------------------------------------------------
		return UpdateReservation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a InventoryItem on a Reservation
//----------------------------------------------------------------------------
func AssignInventoryItemToReservation( reservationId uint64, inventoryItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

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
			// assign the InventoryItem	to the Reservation
			//----------------------------------------------------------------------------
			parentObj.InventoryItem = &childObj

			//----------------------------------------------------------------------------
			// save the Reservation
			//----------------------------------------------------------------------------
			return UpdateReservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItem", inventoryItemId )
			return utils.RequestResult{false, msg, "assignInventoryItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InventoryItem on a Reservation
//----------------------------------------------------------------------------
func UnassignInventoryItemFromReservation(reservationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

		//----------------------------------------------------------------------------
		// assign an empty InventoryItem to the InventoryItem
		//----------------------------------------------------------------------------
		parentObj.InventoryItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InventoryItem
		//----------------------------------------------------------------------------
		parentObj.InventoryItemId = nil;

		//----------------------------------------------------------------------------
		// save the Reservation
		//----------------------------------------------------------------------------
		return UpdateReservation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lot on a Reservation
//----------------------------------------------------------------------------
func AssignLotToReservation( reservationId uint64, lotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

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
			// assign the Lot	to the Reservation
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the Reservation
			//----------------------------------------------------------------------------
			return UpdateReservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a Reservation
//----------------------------------------------------------------------------
func UnassignLotFromReservation(reservationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

		//----------------------------------------------------------------------------
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the Reservation
		//----------------------------------------------------------------------------
		return UpdateReservation(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a DemandSignal on a Reservation
//----------------------------------------------------------------------------
func AssignDemandSignalToReservation( reservationId uint64, demandSignalId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.DemandSignal

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a DemandSignal with a
		// matching demandSignalId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, demandSignalId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the DemandSignal	to the Reservation
			//----------------------------------------------------------------------------
			parentObj.DemandSignal = &childObj

			//----------------------------------------------------------------------------
			// save the Reservation
			//----------------------------------------------------------------------------
			return UpdateReservation(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DemandSignal", demandSignalId )
			return utils.RequestResult{false, msg, "assignDemandSignal", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a DemandSignal on a Reservation
//----------------------------------------------------------------------------
func UnassignDemandSignalFromReservation(reservationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

		//----------------------------------------------------------------------------
		// assign an empty DemandSignal to the DemandSignal
		//----------------------------------------------------------------------------
		parentObj.DemandSignal = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the DemandSignal
		//----------------------------------------------------------------------------
		parentObj.DemandSignalId = nil;

		//----------------------------------------------------------------------------
		// save the Reservation
		//----------------------------------------------------------------------------
		return UpdateReservation(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a Reservation
//----------------------------------------------------------------------------
func AddSerialNumbersToReservation ( reservationId uint64, serialNumbersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

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
		// retrieve the modified Reservation from the gorm
		//----------------------------------------------------------------------------
		return GetReservation(reservationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a Reservation
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromReservation( reservationId uint64, serialNumbersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Reservation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReservation(reservationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Reservation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Reservation)

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
		// retrieve the modified Reservation from the gorm
		//----------------------------------------------------------------------------
		return GetReservation(reservationId)

	} else {
		return parentRequestResult
	}
}

