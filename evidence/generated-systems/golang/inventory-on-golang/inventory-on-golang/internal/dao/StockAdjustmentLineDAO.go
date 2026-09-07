package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing StockAdjustmentLineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateStockAdjustmentLine - creates a new db entry
//----------------------------------------------------------------------------
func CreateStockAdjustmentLine(obj model.StockAdjustmentLine)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a StockAdjustmentLine with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a StockAdjustmentLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateStockAdjustmentLine", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetStockAdjustmentLine - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetStockAdjustmentLine(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.StockAdjustmentLine

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a StockAdjustmentLine with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a StockAdjustmentLine using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a StockAdjustmentLine using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetStockAdjustmentLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllStockAdjustmentLine - returns all
//----------------------------------------------------------------------------
func GetAllStockAdjustmentLine()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.StockAdjustmentLine

	//----------------------------------------------------------------------------
	// Request the ORM to find all StockAdjustmentLine
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all StockAdjustmentLine" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all StockAdjustmentLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllStockAdjustmentLine", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateStockAdjustmentLine - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateStockAdjustmentLine(obj model.StockAdjustmentLine)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a StockAdjustmentLine using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a StockAdjustmentLine using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateStockAdjustmentLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteStockAdjustmentLine - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteStockAdjustmentLine(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetStockAdjustmentLine(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.StockAdjustmentLine)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a StockAdjustmentLine using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a StockAdjustmentLine using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteStockAdjustmentLine", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Adjustment on a StockAdjustmentLine
//----------------------------------------------------------------------------
func AssignAdjustmentToStockAdjustmentLine( stockAdjustmentLineId uint64, adjustmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

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
			// assign the Adjustment	to the StockAdjustmentLine
			//----------------------------------------------------------------------------
			parentObj.Adjustment = &childObj

			//----------------------------------------------------------------------------
			// save the StockAdjustmentLine
			//----------------------------------------------------------------------------
			return UpdateStockAdjustmentLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Adjustment", adjustmentId )
			return utils.RequestResult{false, msg, "assignAdjustment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Adjustment on a StockAdjustmentLine
//----------------------------------------------------------------------------
func UnassignAdjustmentFromStockAdjustmentLine(stockAdjustmentLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

		//----------------------------------------------------------------------------
		// assign an empty StockAdjustment to the Adjustment
		//----------------------------------------------------------------------------
		parentObj.Adjustment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Adjustment
		//----------------------------------------------------------------------------
		parentObj.AdjustmentId = nil;

		//----------------------------------------------------------------------------
		// save the StockAdjustmentLine
		//----------------------------------------------------------------------------
		return UpdateStockAdjustmentLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Sku on a StockAdjustmentLine
//----------------------------------------------------------------------------
func AssignSkuToStockAdjustmentLine( stockAdjustmentLineId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

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
			// assign the Sku	to the StockAdjustmentLine
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the StockAdjustmentLine
			//----------------------------------------------------------------------------
			return UpdateStockAdjustmentLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a StockAdjustmentLine
//----------------------------------------------------------------------------
func UnassignSkuFromStockAdjustmentLine(stockAdjustmentLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the StockAdjustmentLine
		//----------------------------------------------------------------------------
		return UpdateStockAdjustmentLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lot on a StockAdjustmentLine
//----------------------------------------------------------------------------
func AssignLotToStockAdjustmentLine( stockAdjustmentLineId uint64, lotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

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
			// assign the Lot	to the StockAdjustmentLine
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the StockAdjustmentLine
			//----------------------------------------------------------------------------
			return UpdateStockAdjustmentLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a StockAdjustmentLine
//----------------------------------------------------------------------------
func UnassignLotFromStockAdjustmentLine(stockAdjustmentLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

		//----------------------------------------------------------------------------
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the StockAdjustmentLine
		//----------------------------------------------------------------------------
		return UpdateStockAdjustmentLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Location on a StockAdjustmentLine
//----------------------------------------------------------------------------
func AssignLocationToStockAdjustmentLine( stockAdjustmentLineId uint64, locationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

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
			// assign the Location	to the StockAdjustmentLine
			//----------------------------------------------------------------------------
			parentObj.Location = &childObj

			//----------------------------------------------------------------------------
			// save the StockAdjustmentLine
			//----------------------------------------------------------------------------
			return UpdateStockAdjustmentLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Location", locationId )
			return utils.RequestResult{false, msg, "assignLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Location on a StockAdjustmentLine
//----------------------------------------------------------------------------
func UnassignLocationFromStockAdjustmentLine(stockAdjustmentLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the Location
		//----------------------------------------------------------------------------
		parentObj.Location = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Location
		//----------------------------------------------------------------------------
		parentObj.LocationId = nil;

		//----------------------------------------------------------------------------
		// save the StockAdjustmentLine
		//----------------------------------------------------------------------------
		return UpdateStockAdjustmentLine(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a StockAdjustmentLine
//----------------------------------------------------------------------------
func AddSerialNumbersToStockAdjustmentLine ( stockAdjustmentLineId uint64, serialNumbersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

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
		// retrieve the modified StockAdjustmentLine from the gorm
		//----------------------------------------------------------------------------
		return GetStockAdjustmentLine(stockAdjustmentLineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a StockAdjustmentLine
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromStockAdjustmentLine( stockAdjustmentLineId uint64, serialNumbersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StockAdjustmentLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustmentLine(stockAdjustmentLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustmentLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustmentLine)

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
		// retrieve the modified StockAdjustmentLine from the gorm
		//----------------------------------------------------------------------------
		return GetStockAdjustmentLine(stockAdjustmentLineId)

	} else {
		return parentRequestResult
	}
}

