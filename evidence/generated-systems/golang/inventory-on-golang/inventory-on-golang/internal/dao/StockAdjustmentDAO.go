package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing StockAdjustmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateStockAdjustment - creates a new db entry
//----------------------------------------------------------------------------
func CreateStockAdjustment(obj model.StockAdjustment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a StockAdjustment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a StockAdjustment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateStockAdjustment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetStockAdjustment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetStockAdjustment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.StockAdjustment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a StockAdjustment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a StockAdjustment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a StockAdjustment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetStockAdjustment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllStockAdjustment - returns all
//----------------------------------------------------------------------------
func GetAllStockAdjustment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.StockAdjustment

	//----------------------------------------------------------------------------
	// Request the ORM to find all StockAdjustment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all StockAdjustment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all StockAdjustment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllStockAdjustment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateStockAdjustment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateStockAdjustment(obj model.StockAdjustment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a StockAdjustment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a StockAdjustment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateStockAdjustment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteStockAdjustment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteStockAdjustment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetStockAdjustment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.StockAdjustment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a StockAdjustment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a StockAdjustment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteStockAdjustment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Warehouse on a StockAdjustment
//----------------------------------------------------------------------------
func AssignWarehouseToStockAdjustment( stockAdjustmentId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustment(stockAdjustmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustment)

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
			// assign the Warehouse	to the StockAdjustment
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the StockAdjustment
			//----------------------------------------------------------------------------
			return UpdateStockAdjustment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a StockAdjustment
//----------------------------------------------------------------------------
func UnassignWarehouseFromStockAdjustment(stockAdjustmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustment(stockAdjustmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustment)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the StockAdjustment
		//----------------------------------------------------------------------------
		return UpdateStockAdjustment(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more linesIds as a Lines to a StockAdjustment
//----------------------------------------------------------------------------
func AddLinesToStockAdjustment ( stockAdjustmentId uint64, linesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustment(stockAdjustmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustment)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StockAdjustmentLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StockAdjustmentLine
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
		// retrieve the modified StockAdjustment from the gorm
		//----------------------------------------------------------------------------
		return GetStockAdjustment(stockAdjustmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more linesIds as a Lines from a StockAdjustment
//----------------------------------------------------------------------------
func RemoveLinesFromStockAdjustment( stockAdjustmentId uint64, linesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StockAdjustment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustment(stockAdjustmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustment)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StockAdjustmentLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StockAdjustmentLine
			// with a matching linesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , linesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove StockAdjustmentLineObj from the Lines array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lines").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lines", linesId )
				return utils.RequestResult{false, msg, "removeLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StockAdjustment from the gorm
		//----------------------------------------------------------------------------
		return GetStockAdjustment(stockAdjustmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a StockAdjustment
//----------------------------------------------------------------------------
func AddTransactionsToStockAdjustment ( stockAdjustmentId uint64, transactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockAdjustment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustment(stockAdjustmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustment)

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
		// retrieve the modified StockAdjustment from the gorm
		//----------------------------------------------------------------------------
		return GetStockAdjustment(stockAdjustmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a StockAdjustment
//----------------------------------------------------------------------------
func RemoveTransactionsFromStockAdjustment( stockAdjustmentId uint64, transactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StockAdjustment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockAdjustment(stockAdjustmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockAdjustment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockAdjustment)

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
		// retrieve the modified StockAdjustment from the gorm
		//----------------------------------------------------------------------------
		return GetStockAdjustment(stockAdjustmentId)

	} else {
		return parentRequestResult
	}
}

