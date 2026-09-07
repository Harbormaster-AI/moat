package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing StockKeepingUnitDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateStockKeepingUnit - creates a new db entry
//----------------------------------------------------------------------------
func CreateStockKeepingUnit(obj model.StockKeepingUnit)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a StockKeepingUnit with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a StockKeepingUnit", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateStockKeepingUnit", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetStockKeepingUnit - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetStockKeepingUnit(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.StockKeepingUnit

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a StockKeepingUnit with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a StockKeepingUnit using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a StockKeepingUnit using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetStockKeepingUnit", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllStockKeepingUnit - returns all
//----------------------------------------------------------------------------
func GetAllStockKeepingUnit()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.StockKeepingUnit

	//----------------------------------------------------------------------------
	// Request the ORM to find all StockKeepingUnit
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all StockKeepingUnit" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all StockKeepingUnit", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllStockKeepingUnit", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateStockKeepingUnit - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateStockKeepingUnit(obj model.StockKeepingUnit)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a StockKeepingUnit using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a StockKeepingUnit using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateStockKeepingUnit", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteStockKeepingUnit - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteStockKeepingUnit(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetStockKeepingUnit(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.StockKeepingUnit)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a StockKeepingUnit using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a StockKeepingUnit using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteStockKeepingUnit", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more inventoryItemsIds as a InventoryItems to a StockKeepingUnit
//----------------------------------------------------------------------------
func AddInventoryItemsToStockKeepingUnit ( stockKeepingUnitId uint64, inventoryItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventoryItemsIds, ",")

		for _, inventoryItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching inventoryItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventoryItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InventoryItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventoryItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItems", inventoryItemsId )
				return utils.RequestResult{false, msg, "unassignInventoryItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventoryItemsIds as a InventoryItems from a StockKeepingUnit
//----------------------------------------------------------------------------
func RemoveInventoryItemsFromStockKeepingUnit( stockKeepingUnitId uint64, inventoryItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventoryItemsIds, ",")

		for _, inventoryItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching inventoryItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventoryItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InventoryItemObj from the InventoryItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventoryItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItems", inventoryItemsId )
				return utils.RequestResult{false, msg, "removeInventoryItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more uomConversionsIds as a UomConversions to a StockKeepingUnit
//----------------------------------------------------------------------------
func AddUomConversionsToStockKeepingUnit ( stockKeepingUnitId uint64, uomConversionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( uomConversionsIds, ",")

		for _, uomConversionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.UoMConversion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a UoMConversion
			// with a matching uomConversionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , uomConversionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the UomConversions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("UomConversions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "UomConversions", uomConversionsId )
				return utils.RequestResult{false, msg, "unassignUomConversions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more uomConversionsIds as a UomConversions from a StockKeepingUnit
//----------------------------------------------------------------------------
func RemoveUomConversionsFromStockKeepingUnit( stockKeepingUnitId uint64, uomConversionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( uomConversionsIds, ",")

		for _, uomConversionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.UoMConversion

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a UoMConversion
			// with a matching uomConversionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , uomConversionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove UoMConversionObj from the UomConversions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("UomConversions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "UomConversions", uomConversionsId )
				return utils.RequestResult{false, msg, "removeUomConversions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more replenishmentPoliciesIds as a ReplenishmentPolicies to a StockKeepingUnit
//----------------------------------------------------------------------------
func AddReplenishmentPoliciesToStockKeepingUnit ( stockKeepingUnitId uint64, replenishmentPoliciesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( replenishmentPoliciesIds, ",")

		for _, replenishmentPoliciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ReplenishmentPolicy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ReplenishmentPolicy
			// with a matching replenishmentPoliciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , replenishmentPoliciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ReplenishmentPolicies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ReplenishmentPolicies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ReplenishmentPolicies", replenishmentPoliciesId )
				return utils.RequestResult{false, msg, "unassignReplenishmentPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more replenishmentPoliciesIds as a ReplenishmentPolicies from a StockKeepingUnit
//----------------------------------------------------------------------------
func RemoveReplenishmentPoliciesFromStockKeepingUnit( stockKeepingUnitId uint64, replenishmentPoliciesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( replenishmentPoliciesIds, ",")

		for _, replenishmentPoliciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ReplenishmentPolicy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ReplenishmentPolicy
			// with a matching replenishmentPoliciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , replenishmentPoliciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ReplenishmentPolicyObj from the ReplenishmentPolicies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ReplenishmentPolicies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ReplenishmentPolicies", replenishmentPoliciesId )
				return utils.RequestResult{false, msg, "removeReplenishmentPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more lotsIds as a Lots to a StockKeepingUnit
//----------------------------------------------------------------------------
func AddLotsToStockKeepingUnit ( stockKeepingUnitId uint64, lotsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( lotsIds, ",")

		for _, lotsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Lot

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Lot
			// with a matching lotsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lotsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Lots using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lots").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lots", lotsId )
				return utils.RequestResult{false, msg, "unassignLots", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more lotsIds as a Lots from a StockKeepingUnit
//----------------------------------------------------------------------------
func RemoveLotsFromStockKeepingUnit( stockKeepingUnitId uint64, lotsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

		// slice the ids on comma with no spaces
		ids := strings.Split( lotsIds, ",")

		for _, lotsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Lot

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Lot
			// with a matching lotsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lotsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LotObj from the Lots array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lots").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lots", lotsId )
				return utils.RequestResult{false, msg, "removeLots", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more serialNumbersIds as a SerialNumbers to a StockKeepingUnit
//----------------------------------------------------------------------------
func AddSerialNumbersToStockKeepingUnit ( stockKeepingUnitId uint64, serialNumbersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

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
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serialNumbersIds as a SerialNumbers from a StockKeepingUnit
//----------------------------------------------------------------------------
func RemoveSerialNumbersFromStockKeepingUnit( stockKeepingUnitId uint64, serialNumbersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the StockKeepingUnit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetStockKeepingUnit(stockKeepingUnitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.StockKeepingUnit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.StockKeepingUnit)

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
		// retrieve the modified StockKeepingUnit from the gorm
		//----------------------------------------------------------------------------
		return GetStockKeepingUnit(stockKeepingUnitId)

	} else {
		return parentRequestResult
	}
}

