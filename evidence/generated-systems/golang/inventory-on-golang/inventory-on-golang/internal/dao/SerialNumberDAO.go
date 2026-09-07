package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SerialNumberDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSerialNumber - creates a new db entry
//----------------------------------------------------------------------------
func CreateSerialNumber(obj model.SerialNumber)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SerialNumber with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SerialNumber", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSerialNumber", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSerialNumber - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSerialNumber(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SerialNumber

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SerialNumber with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SerialNumber using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SerialNumber using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSerialNumber", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSerialNumber - returns all
//----------------------------------------------------------------------------
func GetAllSerialNumber()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SerialNumber

	//----------------------------------------------------------------------------
	// Request the ORM to find all SerialNumber
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SerialNumber" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SerialNumber", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSerialNumber", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSerialNumber - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSerialNumber(obj model.SerialNumber)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SerialNumber using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SerialNumber using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSerialNumber", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSerialNumber - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSerialNumber(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SerialNumber with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSerialNumber(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SerialNumber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SerialNumber)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SerialNumber using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SerialNumber using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSerialNumber", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Sku on a SerialNumber
//----------------------------------------------------------------------------
func AssignSkuToSerialNumber( serialNumberId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SerialNumber with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSerialNumber(serialNumberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SerialNumber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SerialNumber)

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
			// assign the Sku	to the SerialNumber
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the SerialNumber
			//----------------------------------------------------------------------------
			return UpdateSerialNumber(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a SerialNumber
//----------------------------------------------------------------------------
func UnassignSkuFromSerialNumber(serialNumberId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SerialNumber with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSerialNumber(serialNumberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SerialNumber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SerialNumber)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the SerialNumber
		//----------------------------------------------------------------------------
		return UpdateSerialNumber(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CurrentInventoryItem on a SerialNumber
//----------------------------------------------------------------------------
func AssignCurrentInventoryItemToSerialNumber( serialNumberId uint64, currentInventoryItemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SerialNumber with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSerialNumber(serialNumberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SerialNumber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SerialNumber)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InventoryItem

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InventoryItem with a
		// matching currentInventoryItemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, currentInventoryItemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CurrentInventoryItem	to the SerialNumber
			//----------------------------------------------------------------------------
			parentObj.CurrentInventoryItem = &childObj

			//----------------------------------------------------------------------------
			// save the SerialNumber
			//----------------------------------------------------------------------------
			return UpdateSerialNumber(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CurrentInventoryItem", currentInventoryItemId )
			return utils.RequestResult{false, msg, "assignCurrentInventoryItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CurrentInventoryItem on a SerialNumber
//----------------------------------------------------------------------------
func UnassignCurrentInventoryItemFromSerialNumber(serialNumberId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SerialNumber with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSerialNumber(serialNumberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SerialNumber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SerialNumber)

		//----------------------------------------------------------------------------
		// assign an empty InventoryItem to the CurrentInventoryItem
		//----------------------------------------------------------------------------
		parentObj.CurrentInventoryItem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CurrentInventoryItem
		//----------------------------------------------------------------------------
		parentObj.CurrentInventoryItemId = nil;

		//----------------------------------------------------------------------------
		// save the SerialNumber
		//----------------------------------------------------------------------------
		return UpdateSerialNumber(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lot on a SerialNumber
//----------------------------------------------------------------------------
func AssignLotToSerialNumber( serialNumberId uint64, lotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SerialNumber with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSerialNumber(serialNumberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SerialNumber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SerialNumber)

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
			// assign the Lot	to the SerialNumber
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the SerialNumber
			//----------------------------------------------------------------------------
			return UpdateSerialNumber(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a SerialNumber
//----------------------------------------------------------------------------
func UnassignLotFromSerialNumber(serialNumberId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SerialNumber with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSerialNumber(serialNumberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SerialNumber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SerialNumber)

		//----------------------------------------------------------------------------
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the SerialNumber
		//----------------------------------------------------------------------------
		return UpdateSerialNumber(parentObj)

	} else {
		return parentRequestResult
	}

}


