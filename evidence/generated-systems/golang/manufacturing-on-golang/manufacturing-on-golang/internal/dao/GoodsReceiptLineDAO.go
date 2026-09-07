package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing GoodsReceiptLineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateGoodsReceiptLine - creates a new db entry
//----------------------------------------------------------------------------
func CreateGoodsReceiptLine(obj model.GoodsReceiptLine)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a GoodsReceiptLine with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a GoodsReceiptLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateGoodsReceiptLine", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetGoodsReceiptLine - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetGoodsReceiptLine(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.GoodsReceiptLine

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a GoodsReceiptLine with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a GoodsReceiptLine using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a GoodsReceiptLine using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetGoodsReceiptLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllGoodsReceiptLine - returns all
//----------------------------------------------------------------------------
func GetAllGoodsReceiptLine()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.GoodsReceiptLine

	//----------------------------------------------------------------------------
	// Request the ORM to find all GoodsReceiptLine
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all GoodsReceiptLine" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all GoodsReceiptLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllGoodsReceiptLine", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateGoodsReceiptLine - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateGoodsReceiptLine(obj model.GoodsReceiptLine)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a GoodsReceiptLine using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a GoodsReceiptLine using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateGoodsReceiptLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteGoodsReceiptLine - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteGoodsReceiptLine(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceiptLine with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetGoodsReceiptLine(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceiptLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.GoodsReceiptLine)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a GoodsReceiptLine using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a GoodsReceiptLine using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteGoodsReceiptLine", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a GoodsReceipt on a GoodsReceiptLine
//----------------------------------------------------------------------------
func AssignGoodsReceiptToGoodsReceiptLine( goodsReceiptLineId uint64, goodsReceiptId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceiptLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceiptLine(goodsReceiptLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceiptLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceiptLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.GoodsReceipt

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a GoodsReceipt with a
		// matching goodsReceiptId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, goodsReceiptId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the GoodsReceipt	to the GoodsReceiptLine
			//----------------------------------------------------------------------------
			parentObj.GoodsReceipt = &childObj

			//----------------------------------------------------------------------------
			// save the GoodsReceiptLine
			//----------------------------------------------------------------------------
			return UpdateGoodsReceiptLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GoodsReceipt", goodsReceiptId )
			return utils.RequestResult{false, msg, "assignGoodsReceipt", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a GoodsReceipt on a GoodsReceiptLine
//----------------------------------------------------------------------------
func UnassignGoodsReceiptFromGoodsReceiptLine(goodsReceiptLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceiptLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceiptLine(goodsReceiptLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceiptLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceiptLine)

		//----------------------------------------------------------------------------
		// assign an empty GoodsReceipt to the GoodsReceipt
		//----------------------------------------------------------------------------
		parentObj.GoodsReceipt = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the GoodsReceipt
		//----------------------------------------------------------------------------
		parentObj.GoodsReceiptId = nil;

		//----------------------------------------------------------------------------
		// save the GoodsReceiptLine
		//----------------------------------------------------------------------------
		return UpdateGoodsReceiptLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Item on a GoodsReceiptLine
//----------------------------------------------------------------------------
func AssignItemToGoodsReceiptLine( goodsReceiptLineId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceiptLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceiptLine(goodsReceiptLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceiptLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceiptLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching itemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, itemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Item	to the GoodsReceiptLine
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the GoodsReceiptLine
			//----------------------------------------------------------------------------
			return UpdateGoodsReceiptLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a GoodsReceiptLine
//----------------------------------------------------------------------------
func UnassignItemFromGoodsReceiptLine(goodsReceiptLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceiptLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceiptLine(goodsReceiptLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceiptLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceiptLine)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the GoodsReceiptLine
		//----------------------------------------------------------------------------
		return UpdateGoodsReceiptLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a InventoryTransaction on a GoodsReceiptLine
//----------------------------------------------------------------------------
func AssignInventoryTransactionToGoodsReceiptLine( goodsReceiptLineId uint64, inventoryTransactionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceiptLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceiptLine(goodsReceiptLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceiptLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceiptLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InventoryTransaction

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InventoryTransaction with a
		// matching inventoryTransactionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, inventoryTransactionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InventoryTransaction	to the GoodsReceiptLine
			//----------------------------------------------------------------------------
			parentObj.InventoryTransaction = &childObj

			//----------------------------------------------------------------------------
			// save the GoodsReceiptLine
			//----------------------------------------------------------------------------
			return UpdateGoodsReceiptLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryTransaction", inventoryTransactionId )
			return utils.RequestResult{false, msg, "assignInventoryTransaction", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InventoryTransaction on a GoodsReceiptLine
//----------------------------------------------------------------------------
func UnassignInventoryTransactionFromGoodsReceiptLine(goodsReceiptLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceiptLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceiptLine(goodsReceiptLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceiptLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceiptLine)

		//----------------------------------------------------------------------------
		// assign an empty InventoryTransaction to the InventoryTransaction
		//----------------------------------------------------------------------------
		parentObj.InventoryTransaction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InventoryTransaction
		//----------------------------------------------------------------------------
		parentObj.InventoryTransactionId = nil;

		//----------------------------------------------------------------------------
		// save the GoodsReceiptLine
		//----------------------------------------------------------------------------
		return UpdateGoodsReceiptLine(parentObj)

	} else {
		return parentRequestResult
	}

}


