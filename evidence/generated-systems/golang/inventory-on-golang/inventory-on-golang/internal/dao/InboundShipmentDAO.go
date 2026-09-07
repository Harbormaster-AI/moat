package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InboundShipmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInboundShipment - creates a new db entry
//----------------------------------------------------------------------------
func CreateInboundShipment(obj model.InboundShipment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InboundShipment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InboundShipment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInboundShipment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInboundShipment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInboundShipment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InboundShipment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InboundShipment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InboundShipment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InboundShipment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInboundShipment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInboundShipment - returns all
//----------------------------------------------------------------------------
func GetAllInboundShipment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InboundShipment

	//----------------------------------------------------------------------------
	// Request the ORM to find all InboundShipment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InboundShipment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InboundShipment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInboundShipment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInboundShipment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInboundShipment(obj model.InboundShipment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InboundShipment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InboundShipment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInboundShipment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInboundShipment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInboundShipment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InboundShipment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInboundShipment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InboundShipment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InboundShipment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InboundShipment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInboundShipment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Warehouse on a InboundShipment
//----------------------------------------------------------------------------
func AssignWarehouseToInboundShipment( inboundShipmentId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InboundShipment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipment(inboundShipmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipment)

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
			// assign the Warehouse	to the InboundShipment
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the InboundShipment
			//----------------------------------------------------------------------------
			return UpdateInboundShipment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a InboundShipment
//----------------------------------------------------------------------------
func UnassignWarehouseFromInboundShipment(inboundShipmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InboundShipment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipment(inboundShipmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipment)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the InboundShipment
		//----------------------------------------------------------------------------
		return UpdateInboundShipment(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more linesIds as a Lines to a InboundShipment
//----------------------------------------------------------------------------
func AddLinesToInboundShipment ( inboundShipmentId uint64, linesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InboundShipment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipment(inboundShipmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipment)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InboundShipmentLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InboundShipmentLine
			// with a matching linesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , linesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Lines using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lines").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lines", linesId )
				return utils.RequestResult{false, msg, "unassignLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InboundShipment from the gorm
		//----------------------------------------------------------------------------
		return GetInboundShipment(inboundShipmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more linesIds as a Lines from a InboundShipment
//----------------------------------------------------------------------------
func RemoveLinesFromInboundShipment( inboundShipmentId uint64, linesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InboundShipment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipment(inboundShipmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipment)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InboundShipmentLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InboundShipmentLine
			// with a matching linesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , linesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InboundShipmentLineObj from the Lines array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lines").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lines", linesId )
				return utils.RequestResult{false, msg, "removeLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InboundShipment from the gorm
		//----------------------------------------------------------------------------
		return GetInboundShipment(inboundShipmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a InboundShipment
//----------------------------------------------------------------------------
func AddTransactionsToInboundShipment ( inboundShipmentId uint64, transactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InboundShipment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipment(inboundShipmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipment)

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
		// retrieve the modified InboundShipment from the gorm
		//----------------------------------------------------------------------------
		return GetInboundShipment(inboundShipmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a InboundShipment
//----------------------------------------------------------------------------
func RemoveTransactionsFromInboundShipment( inboundShipmentId uint64, transactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InboundShipment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInboundShipment(inboundShipmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InboundShipment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InboundShipment)

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
		// retrieve the modified InboundShipment from the gorm
		//----------------------------------------------------------------------------
		return GetInboundShipment(inboundShipmentId)

	} else {
		return parentRequestResult
	}
}

