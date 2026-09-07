package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CycleCountEntryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCycleCountEntry - creates a new db entry
//----------------------------------------------------------------------------
func CreateCycleCountEntry(obj model.CycleCountEntry)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CycleCountEntry with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CycleCountEntry", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCycleCountEntry", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCycleCountEntry - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCycleCountEntry(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CycleCountEntry

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CycleCountEntry with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CycleCountEntry using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CycleCountEntry using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCycleCountEntry", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCycleCountEntry - returns all
//----------------------------------------------------------------------------
func GetAllCycleCountEntry()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CycleCountEntry

	//----------------------------------------------------------------------------
	// Request the ORM to find all CycleCountEntry
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CycleCountEntry" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CycleCountEntry", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCycleCountEntry", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCycleCountEntry - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCycleCountEntry(obj model.CycleCountEntry)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CycleCountEntry using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CycleCountEntry using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCycleCountEntry", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCycleCountEntry - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCycleCountEntry(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCycleCountEntry(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CycleCountEntry)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CycleCountEntry using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CycleCountEntry using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCycleCountEntry", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a CycleCount on a CycleCountEntry
//----------------------------------------------------------------------------
func AssignCycleCountToCycleCountEntry( cycleCountEntryId uint64, cycleCountId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CycleCount

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CycleCount with a
		// matching cycleCountId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, cycleCountId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CycleCount	to the CycleCountEntry
			//----------------------------------------------------------------------------
			parentObj.CycleCount = &childObj

			//----------------------------------------------------------------------------
			// save the CycleCountEntry
			//----------------------------------------------------------------------------
			return UpdateCycleCountEntry(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CycleCount", cycleCountId )
			return utils.RequestResult{false, msg, "assignCycleCount", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CycleCount on a CycleCountEntry
//----------------------------------------------------------------------------
func UnassignCycleCountFromCycleCountEntry(cycleCountEntryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

		//----------------------------------------------------------------------------
		// assign an empty CycleCount to the CycleCount
		//----------------------------------------------------------------------------
		parentObj.CycleCount = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CycleCount
		//----------------------------------------------------------------------------
		parentObj.CycleCountId = nil;

		//----------------------------------------------------------------------------
		// save the CycleCountEntry
		//----------------------------------------------------------------------------
		return UpdateCycleCountEntry(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Sku on a CycleCountEntry
//----------------------------------------------------------------------------
func AssignSkuToCycleCountEntry( cycleCountEntryId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

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
			// assign the Sku	to the CycleCountEntry
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the CycleCountEntry
			//----------------------------------------------------------------------------
			return UpdateCycleCountEntry(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a CycleCountEntry
//----------------------------------------------------------------------------
func UnassignSkuFromCycleCountEntry(cycleCountEntryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the CycleCountEntry
		//----------------------------------------------------------------------------
		return UpdateCycleCountEntry(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Lot on a CycleCountEntry
//----------------------------------------------------------------------------
func AssignLotToCycleCountEntry( cycleCountEntryId uint64, lotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

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
			// assign the Lot	to the CycleCountEntry
			//----------------------------------------------------------------------------
			parentObj.Lot = &childObj

			//----------------------------------------------------------------------------
			// save the CycleCountEntry
			//----------------------------------------------------------------------------
			return UpdateCycleCountEntry(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lot", lotId )
			return utils.RequestResult{false, msg, "assignLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Lot on a CycleCountEntry
//----------------------------------------------------------------------------
func UnassignLotFromCycleCountEntry(cycleCountEntryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

		//----------------------------------------------------------------------------
		// assign an empty Lot to the Lot
		//----------------------------------------------------------------------------
		parentObj.Lot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Lot
		//----------------------------------------------------------------------------
		parentObj.LotId = nil;

		//----------------------------------------------------------------------------
		// save the CycleCountEntry
		//----------------------------------------------------------------------------
		return UpdateCycleCountEntry(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Location on a CycleCountEntry
//----------------------------------------------------------------------------
func AssignLocationToCycleCountEntry( cycleCountEntryId uint64, locationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

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
			// assign the Location	to the CycleCountEntry
			//----------------------------------------------------------------------------
			parentObj.Location = &childObj

			//----------------------------------------------------------------------------
			// save the CycleCountEntry
			//----------------------------------------------------------------------------
			return UpdateCycleCountEntry(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Location", locationId )
			return utils.RequestResult{false, msg, "assignLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Location on a CycleCountEntry
//----------------------------------------------------------------------------
func UnassignLocationFromCycleCountEntry(cycleCountEntryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the Location
		//----------------------------------------------------------------------------
		parentObj.Location = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Location
		//----------------------------------------------------------------------------
		parentObj.LocationId = nil;

		//----------------------------------------------------------------------------
		// save the CycleCountEntry
		//----------------------------------------------------------------------------
		return UpdateCycleCountEntry(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a CycleCountEntry
//----------------------------------------------------------------------------
func AddSerialNumbersToCycleCountEntry ( cycleCountEntryId uint64, serialNumbersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

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
		// retrieve the modified CycleCountEntry from the gorm
		//----------------------------------------------------------------------------
		return GetCycleCountEntry(cycleCountEntryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a CycleCountEntry
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromCycleCountEntry( cycleCountEntryId uint64, serialNumbersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the CycleCountEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCycleCountEntry(cycleCountEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CycleCountEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CycleCountEntry)

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
		// retrieve the modified CycleCountEntry from the gorm
		//----------------------------------------------------------------------------
		return GetCycleCountEntry(cycleCountEntryId)

	} else {
		return parentRequestResult
	}
}

