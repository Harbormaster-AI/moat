package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TransferOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTransferOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateTransferOrder(obj model.TransferOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TransferOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TransferOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTransferOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTransferOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTransferOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TransferOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TransferOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TransferOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TransferOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTransferOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTransferOrder - returns all
//----------------------------------------------------------------------------
func GetAllTransferOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TransferOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all TransferOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TransferOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TransferOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTransferOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTransferOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTransferOrder(obj model.TransferOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TransferOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TransferOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTransferOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTransferOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTransferOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TransferOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTransferOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TransferOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TransferOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TransferOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTransferOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a OriginWarehouse on a TransferOrder
//----------------------------------------------------------------------------
func AssignOriginWarehouseToTransferOrder( transferOrderId uint64, originWarehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TransferOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrder(transferOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Warehouse

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Warehouse with a
		// matching originWarehouseId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, originWarehouseId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the OriginWarehouse	to the TransferOrder
			//----------------------------------------------------------------------------
			parentObj.OriginWarehouse = &childObj

			//----------------------------------------------------------------------------
			// save the TransferOrder
			//----------------------------------------------------------------------------
			return UpdateTransferOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OriginWarehouse", originWarehouseId )
			return utils.RequestResult{false, msg, "assignOriginWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a OriginWarehouse on a TransferOrder
//----------------------------------------------------------------------------
func UnassignOriginWarehouseFromTransferOrder(transferOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrder(transferOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrder)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the OriginWarehouse
		//----------------------------------------------------------------------------
		parentObj.OriginWarehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the OriginWarehouse
		//----------------------------------------------------------------------------
		parentObj.OriginWarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the TransferOrder
		//----------------------------------------------------------------------------
		return UpdateTransferOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a DestinationWarehouse on a TransferOrder
//----------------------------------------------------------------------------
func AssignDestinationWarehouseToTransferOrder( transferOrderId uint64, destinationWarehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TransferOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrder(transferOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Warehouse

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Warehouse with a
		// matching destinationWarehouseId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, destinationWarehouseId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the DestinationWarehouse	to the TransferOrder
			//----------------------------------------------------------------------------
			parentObj.DestinationWarehouse = &childObj

			//----------------------------------------------------------------------------
			// save the TransferOrder
			//----------------------------------------------------------------------------
			return UpdateTransferOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DestinationWarehouse", destinationWarehouseId )
			return utils.RequestResult{false, msg, "assignDestinationWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a DestinationWarehouse on a TransferOrder
//----------------------------------------------------------------------------
func UnassignDestinationWarehouseFromTransferOrder(transferOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrder(transferOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrder)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the DestinationWarehouse
		//----------------------------------------------------------------------------
		parentObj.DestinationWarehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the DestinationWarehouse
		//----------------------------------------------------------------------------
		parentObj.DestinationWarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the TransferOrder
		//----------------------------------------------------------------------------
		return UpdateTransferOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more linesIds as a Lines to a TransferOrder
//----------------------------------------------------------------------------
func AddLinesToTransferOrder ( transferOrderId uint64, linesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrder(transferOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TransferOrderLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TransferOrderLine
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
		// retrieve the modified TransferOrder from the gorm
		//----------------------------------------------------------------------------
		return GetTransferOrder(transferOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more linesIds as a Lines from a TransferOrder
//----------------------------------------------------------------------------
func RemoveLinesFromTransferOrder( transferOrderId uint64, linesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TransferOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrder(transferOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TransferOrderLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TransferOrderLine
			// with a matching linesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , linesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TransferOrderLineObj from the Lines array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lines").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lines", linesId )
				return utils.RequestResult{false, msg, "removeLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TransferOrder from the gorm
		//----------------------------------------------------------------------------
		return GetTransferOrder(transferOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a TransferOrder
//----------------------------------------------------------------------------
func AddTransactionsToTransferOrder ( transferOrderId uint64, transactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TransferOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrder(transferOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrder)

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
		// retrieve the modified TransferOrder from the gorm
		//----------------------------------------------------------------------------
		return GetTransferOrder(transferOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a TransferOrder
//----------------------------------------------------------------------------
func RemoveTransactionsFromTransferOrder( transferOrderId uint64, transactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TransferOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTransferOrder(transferOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TransferOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TransferOrder)

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
		// retrieve the modified TransferOrder from the gorm
		//----------------------------------------------------------------------------
		return GetTransferOrder(transferOrderId)

	} else {
		return parentRequestResult
	}
}

