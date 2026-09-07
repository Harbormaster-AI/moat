package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TransferOrderLineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTransferOrderLine - creates a new db entry
//----------------------------------------------------------------------------
func CreateTransferOrderLine(obj model.TransferOrderLine)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TransferOrderLine with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TransferOrderLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTransferOrderLine", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTransferOrderLine - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTransferOrderLine(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TransferOrderLine

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TransferOrderLine with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TransferOrderLine using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TransferOrderLine using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTransferOrderLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTransferOrderLine - returns all
//----------------------------------------------------------------------------
func GetAllTransferOrderLine()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TransferOrderLine

	//----------------------------------------------------------------------------
	// Request the ORM to find all TransferOrderLine
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TransferOrderLine" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TransferOrderLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTransferOrderLine", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTransferOrderLine - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTransferOrderLine(obj model.TransferOrderLine)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TransferOrderLine using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TransferOrderLine using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTransferOrderLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTransferOrderLine - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTransferOrderLine(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTransferOrderLine(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TransferOrderLine)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TransferOrderLine using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TransferOrderLine using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTransferOrderLine", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a TransferOrder on a TransferOrderLine
//----------------------------------------------------------------------------
func AssignTransferOrderToTransferOrderLine( transferOrderLineId uint64, transferOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

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
			// assign the TransferOrder	to the TransferOrderLine
			//----------------------------------------------------------------------------
			parentObj.TransferOrder = &childObj

			//----------------------------------------------------------------------------
			// save the TransferOrderLine
			//----------------------------------------------------------------------------
			return UpdateTransferOrderLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TransferOrder", transferOrderId )
			return utils.RequestResult{false, msg, "assignTransferOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TransferOrder on a TransferOrderLine
//----------------------------------------------------------------------------
func UnassignTransferOrderFromTransferOrderLine(transferOrderLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

		//----------------------------------------------------------------------------
		// assign an empty TransferOrder to the TransferOrder
		//----------------------------------------------------------------------------
		parentObj.TransferOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TransferOrder
		//----------------------------------------------------------------------------
		parentObj.TransferOrderId = nil;

		//----------------------------------------------------------------------------
		// save the TransferOrderLine
		//----------------------------------------------------------------------------
		return UpdateTransferOrderLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Sku on a TransferOrderLine
//----------------------------------------------------------------------------
func AssignSkuToTransferOrderLine( transferOrderLineId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

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
			// assign the Sku	to the TransferOrderLine
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the TransferOrderLine
			//----------------------------------------------------------------------------
			return UpdateTransferOrderLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a TransferOrderLine
//----------------------------------------------------------------------------
func UnassignSkuFromTransferOrderLine(transferOrderLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the TransferOrderLine
		//----------------------------------------------------------------------------
		return UpdateTransferOrderLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lot on a TransferOrderLine
//----------------------------------------------------------------------------
func AssignLotToTransferOrderLine( transferOrderLineId uint64, lotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

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
			// assign the Lot	to the TransferOrderLine
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the TransferOrderLine
			//----------------------------------------------------------------------------
			return UpdateTransferOrderLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a TransferOrderLine
//----------------------------------------------------------------------------
func UnassignLotFromTransferOrderLine(transferOrderLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

		//----------------------------------------------------------------------------
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the TransferOrderLine
		//----------------------------------------------------------------------------
		return UpdateTransferOrderLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a FromLocation on a TransferOrderLine
//----------------------------------------------------------------------------
func AssignFromLocationToTransferOrderLine( transferOrderLineId uint64, fromLocationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.StorageLocation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StorageLocation with a
		// matching fromLocationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, fromLocationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the FromLocation	to the TransferOrderLine
			//----------------------------------------------------------------------------
			parentObj.FromLocation = &childObj

			//----------------------------------------------------------------------------
			// save the TransferOrderLine
			//----------------------------------------------------------------------------
			return UpdateTransferOrderLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FromLocation", fromLocationId )
			return utils.RequestResult{false, msg, "assignFromLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a FromLocation on a TransferOrderLine
//----------------------------------------------------------------------------
func UnassignFromLocationFromTransferOrderLine(transferOrderLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the FromLocation
		//----------------------------------------------------------------------------
		parentObj.FromLocation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the FromLocation
		//----------------------------------------------------------------------------
		parentObj.FromLocationId = nil;

		//----------------------------------------------------------------------------
		// save the TransferOrderLine
		//----------------------------------------------------------------------------
		return UpdateTransferOrderLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ToLocation on a TransferOrderLine
//----------------------------------------------------------------------------
func AssignToLocationToTransferOrderLine( transferOrderLineId uint64, toLocationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.StorageLocation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a StorageLocation with a
		// matching toLocationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, toLocationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ToLocation	to the TransferOrderLine
			//----------------------------------------------------------------------------
			parentObj.ToLocation = &childObj

			//----------------------------------------------------------------------------
			// save the TransferOrderLine
			//----------------------------------------------------------------------------
			return UpdateTransferOrderLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ToLocation", toLocationId )
			return utils.RequestResult{false, msg, "assignToLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ToLocation on a TransferOrderLine
//----------------------------------------------------------------------------
func UnassignToLocationFromTransferOrderLine(transferOrderLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the ToLocation
		//----------------------------------------------------------------------------
		parentObj.ToLocation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ToLocation
		//----------------------------------------------------------------------------
		parentObj.ToLocationId = nil;

		//----------------------------------------------------------------------------
		// save the TransferOrderLine
		//----------------------------------------------------------------------------
		return UpdateTransferOrderLine(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a TransferOrderLine
//----------------------------------------------------------------------------
func AddSerialNumbersToTransferOrderLine ( transferOrderLineId uint64, serialNumbersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

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
		// retrieve the modified TransferOrderLine from the gorm
		//----------------------------------------------------------------------------
		return GetTransferOrderLine(transferOrderLineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a TransferOrderLine
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromTransferOrderLine( transferOrderLineId uint64, serialNumbersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TransferOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrderLine(transferOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrderLine)

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
		// retrieve the modified TransferOrderLine from the gorm
		//----------------------------------------------------------------------------
		return GetTransferOrderLine(transferOrderLineId)

	} else {
		return parentRequestResult
	}
}

