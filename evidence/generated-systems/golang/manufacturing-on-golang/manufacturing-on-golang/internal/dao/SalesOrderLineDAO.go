package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SalesOrderLineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSalesOrderLine - creates a new db entry
//----------------------------------------------------------------------------
func CreateSalesOrderLine(obj model.SalesOrderLine)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SalesOrderLine with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SalesOrderLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSalesOrderLine", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSalesOrderLine - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSalesOrderLine(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SalesOrderLine

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SalesOrderLine with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SalesOrderLine using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SalesOrderLine using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSalesOrderLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSalesOrderLine - returns all
//----------------------------------------------------------------------------
func GetAllSalesOrderLine()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SalesOrderLine

	//----------------------------------------------------------------------------
	// Request the ORM to find all SalesOrderLine
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SalesOrderLine" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SalesOrderLine", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSalesOrderLine", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSalesOrderLine - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSalesOrderLine(obj model.SalesOrderLine)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SalesOrderLine using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SalesOrderLine using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSalesOrderLine", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSalesOrderLine - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSalesOrderLine(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SalesOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSalesOrderLine(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SalesOrderLine)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SalesOrderLine using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SalesOrderLine using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSalesOrderLine", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a SalesOrder on a SalesOrderLine
//----------------------------------------------------------------------------
func AssignSalesOrderToSalesOrderLine( salesOrderLineId uint64, salesOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SalesOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrderLine(salesOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrderLine)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.SalesOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a SalesOrder with a
		// matching salesOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, salesOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SalesOrder	to the SalesOrderLine
			//----------------------------------------------------------------------------
			parentObj.SalesOrder = &childObj

			//----------------------------------------------------------------------------
			// save the SalesOrderLine
			//----------------------------------------------------------------------------
			return UpdateSalesOrderLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalesOrder", salesOrderId )
			return utils.RequestResult{false, msg, "assignSalesOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SalesOrder on a SalesOrderLine
//----------------------------------------------------------------------------
func UnassignSalesOrderFromSalesOrderLine(salesOrderLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrderLine(salesOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrderLine)

		//----------------------------------------------------------------------------
		// assign an empty SalesOrder to the SalesOrder
		//----------------------------------------------------------------------------
		parentObj.SalesOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SalesOrder
		//----------------------------------------------------------------------------
		parentObj.SalesOrderId = nil;

		//----------------------------------------------------------------------------
		// save the SalesOrderLine
		//----------------------------------------------------------------------------
		return UpdateSalesOrderLine(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Item on a SalesOrderLine
//----------------------------------------------------------------------------
func AssignItemToSalesOrderLine( salesOrderLineId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SalesOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrderLine(salesOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrderLine)

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
			// assign the Item	to the SalesOrderLine
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the SalesOrderLine
			//----------------------------------------------------------------------------
			return UpdateSalesOrderLine(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a SalesOrderLine
//----------------------------------------------------------------------------
func UnassignItemFromSalesOrderLine(salesOrderLineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesOrderLine with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrderLine(salesOrderLineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrderLine so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrderLine)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the SalesOrderLine
		//----------------------------------------------------------------------------
		return UpdateSalesOrderLine(parentObj)

	} else {
		return parentRequestResult
	}

}


