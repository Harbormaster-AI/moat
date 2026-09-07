package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PurchaseOrderLineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePurchaseOrderLine - creates a new db entry
//----------------------------------------------------------------------------
func CreatePurchaseOrderLine(obj model.PurchaseOrderLine)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PurchaseOrderLine with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PurchaseOrderLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePurchaseOrderLine", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPurchaseOrderLine - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPurchaseOrderLine(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PurchaseOrderLine

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PurchaseOrderLine with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PurchaseOrderLine using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PurchaseOrderLine using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPurchaseOrderLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPurchaseOrderLine - returns all
//----------------------------------------------------------------------------
func GetAllPurchaseOrderLine()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PurchaseOrderLine

	//----------------------------------------------------------------------------
	// Request the ORM to find all PurchaseOrderLine
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PurchaseOrderLine" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PurchaseOrderLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPurchaseOrderLine", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePurchaseOrderLine - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePurchaseOrderLine(obj model.PurchaseOrderLine)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PurchaseOrderLine using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PurchaseOrderLine using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePurchaseOrderLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePurchaseOrderLine - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePurchaseOrderLine(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPurchaseOrderLine(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PurchaseOrderLine)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PurchaseOrderLine using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PurchaseOrderLine using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePurchaseOrderLine", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a PurchaseOrder on a PurchaseOrderLine
//----------------------------------------------------------------------------
func AssignPurchaseOrderToPurchaseOrderLine( purchaseOrderLineId uint64, purchaseOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrderLine(purchaseOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrderLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PurchaseOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PurchaseOrder with a
		// matching purchaseOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, purchaseOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PurchaseOrder	to the PurchaseOrderLine
			//----------------------------------------------------------------------------
			parentObj.PurchaseOrder = &childObj

			//----------------------------------------------------------------------------
			// save the PurchaseOrderLine
			//----------------------------------------------------------------------------
			return UpdatePurchaseOrderLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PurchaseOrder", purchaseOrderId )
			return utils.RequestResult{false, msg, "assignPurchaseOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PurchaseOrder on a PurchaseOrderLine
//----------------------------------------------------------------------------
func UnassignPurchaseOrderFromPurchaseOrderLine(purchaseOrderLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrderLine(purchaseOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrderLine)

		//----------------------------------------------------------------------------
		// assign an empty PurchaseOrder to the PurchaseOrder
		//----------------------------------------------------------------------------
		parentObj.PurchaseOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PurchaseOrder
		//----------------------------------------------------------------------------
		parentObj.PurchaseOrderId = nil;

		//----------------------------------------------------------------------------
		// save the PurchaseOrderLine
		//----------------------------------------------------------------------------
		return UpdatePurchaseOrderLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Item on a PurchaseOrderLine
//----------------------------------------------------------------------------
func AssignItemToPurchaseOrderLine( purchaseOrderLineId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrderLine(purchaseOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrderLine)

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
			// assign the Item	to the PurchaseOrderLine
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the PurchaseOrderLine
			//----------------------------------------------------------------------------
			return UpdatePurchaseOrderLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a PurchaseOrderLine
//----------------------------------------------------------------------------
func UnassignItemFromPurchaseOrderLine(purchaseOrderLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrderLine(purchaseOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrderLine)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the PurchaseOrderLine
		//----------------------------------------------------------------------------
		return UpdatePurchaseOrderLine(parentObj)

	} else {
		return parentRequestResult
	}

}


