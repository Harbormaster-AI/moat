package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CycleCountDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCycleCount - creates a new db entry
//----------------------------------------------------------------------------
func CreateCycleCount(obj model.CycleCount)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CycleCount with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CycleCount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCycleCount", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCycleCount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCycleCount(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CycleCount

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CycleCount with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CycleCount using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CycleCount using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCycleCount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCycleCount - returns all
//----------------------------------------------------------------------------
func GetAllCycleCount()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CycleCount

	//----------------------------------------------------------------------------
	// Request the ORM to find all CycleCount
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CycleCount" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CycleCount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCycleCount", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCycleCount - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCycleCount(obj model.CycleCount)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CycleCount using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CycleCount using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCycleCount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCycleCount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCycleCount(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CycleCount with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCycleCount(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CycleCount)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CycleCount using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CycleCount using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCycleCount", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Warehouse on a CycleCount
//----------------------------------------------------------------------------
func AssignWarehouseToCycleCount( cycleCountId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CycleCount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCount(cycleCountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCount)

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
			// assign the Warehouse	to the CycleCount
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the CycleCount
			//----------------------------------------------------------------------------
			return UpdateCycleCount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a CycleCount
//----------------------------------------------------------------------------
func UnassignWarehouseFromCycleCount(cycleCountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CycleCount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCount(cycleCountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCount)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the CycleCount
		//----------------------------------------------------------------------------
		return UpdateCycleCount(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more locationsIds as a Locations to a CycleCount
//----------------------------------------------------------------------------
func AddLocationsToCycleCount ( cycleCountId uint64, locationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CycleCount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCount(cycleCountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCount)

		// slice the ids on comma with no spaces
		ids := strings.Split( locationsIds, ",")

		for _, locationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StorageLocation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StorageLocation
			// with a matching locationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , locationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Locations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Locations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Locations", locationsId )
				return utils.RequestResult{false, msg, "unassignLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CycleCount from the gorm
		//----------------------------------------------------------------------------
		return GetCycleCount(cycleCountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more locationsIds as a Locations from a CycleCount
//----------------------------------------------------------------------------
func RemoveLocationsFromCycleCount( cycleCountId uint64, locationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CycleCount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCount(cycleCountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCount)

		// slice the ids on comma with no spaces
		ids := strings.Split( locationsIds, ",")

		for _, locationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.StorageLocation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a StorageLocation
			// with a matching locationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , locationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove StorageLocationObj from the Locations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Locations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Locations", locationsId )
				return utils.RequestResult{false, msg, "removeLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CycleCount from the gorm
		//----------------------------------------------------------------------------
		return GetCycleCount(cycleCountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more entriesIds as a Entries to a CycleCount
//----------------------------------------------------------------------------
func AddEntriesToCycleCount ( cycleCountId uint64, entriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CycleCount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCount(cycleCountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCount)

		// slice the ids on comma with no spaces
		ids := strings.Split( entriesIds, ",")

		for _, entriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CycleCountEntry

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CycleCountEntry
			// with a matching entriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , entriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Entries using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Entries").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Entries", entriesId )
				return utils.RequestResult{false, msg, "unassignEntries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CycleCount from the gorm
		//----------------------------------------------------------------------------
		return GetCycleCount(cycleCountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more entriesIds as a Entries from a CycleCount
//----------------------------------------------------------------------------
func RemoveEntriesFromCycleCount( cycleCountId uint64, entriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CycleCount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCount(cycleCountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCount)

		// slice the ids on comma with no spaces
		ids := strings.Split( entriesIds, ",")

		for _, entriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CycleCountEntry

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CycleCountEntry
			// with a matching entriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , entriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CycleCountEntryObj from the Entries array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Entries").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Entries", entriesId )
				return utils.RequestResult{false, msg, "removeEntries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified CycleCount from the gorm
		//----------------------------------------------------------------------------
		return GetCycleCount(cycleCountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a CycleCount
//----------------------------------------------------------------------------
func AddTransactionsToCycleCount ( cycleCountId uint64, transactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CycleCount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCount(cycleCountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCount)

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
		// retrieve the modified CycleCount from the gorm
		//----------------------------------------------------------------------------
		return GetCycleCount(cycleCountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a CycleCount
//----------------------------------------------------------------------------
func RemoveTransactionsFromCycleCount( cycleCountId uint64, transactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CycleCount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCount(cycleCountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCount)

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
		// retrieve the modified CycleCount from the gorm
		//----------------------------------------------------------------------------
		return GetCycleCount(cycleCountId)

	} else {
		return parentRequestResult
	}
}

