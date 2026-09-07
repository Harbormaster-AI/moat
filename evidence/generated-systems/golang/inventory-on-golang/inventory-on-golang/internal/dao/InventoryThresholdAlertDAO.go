package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InventoryThresholdAlertDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInventoryThresholdAlert - creates a new db entry
//----------------------------------------------------------------------------
func CreateInventoryThresholdAlert(obj model.InventoryThresholdAlert)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InventoryThresholdAlert with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InventoryThresholdAlert", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInventoryThresholdAlert", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInventoryThresholdAlert - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInventoryThresholdAlert(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InventoryThresholdAlert

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InventoryThresholdAlert with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InventoryThresholdAlert using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InventoryThresholdAlert using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInventoryThresholdAlert", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInventoryThresholdAlert - returns all
//----------------------------------------------------------------------------
func GetAllInventoryThresholdAlert()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InventoryThresholdAlert

	//----------------------------------------------------------------------------
	// Request the ORM to find all InventoryThresholdAlert
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InventoryThresholdAlert" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InventoryThresholdAlert", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInventoryThresholdAlert", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInventoryThresholdAlert - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInventoryThresholdAlert(obj model.InventoryThresholdAlert)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InventoryThresholdAlert using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InventoryThresholdAlert using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInventoryThresholdAlert", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInventoryThresholdAlert - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInventoryThresholdAlert(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InventoryThresholdAlert with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInventoryThresholdAlert(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryThresholdAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InventoryThresholdAlert)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InventoryThresholdAlert using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InventoryThresholdAlert using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInventoryThresholdAlert", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Sku on a InventoryThresholdAlert
//----------------------------------------------------------------------------
func AssignSkuToInventoryThresholdAlert( inventoryThresholdAlertId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryThresholdAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryThresholdAlert(inventoryThresholdAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryThresholdAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryThresholdAlert)

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
			// assign the Sku	to the InventoryThresholdAlert
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryThresholdAlert
			//----------------------------------------------------------------------------
			return UpdateInventoryThresholdAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a InventoryThresholdAlert
//----------------------------------------------------------------------------
func UnassignSkuFromInventoryThresholdAlert(inventoryThresholdAlertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryThresholdAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryThresholdAlert(inventoryThresholdAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryThresholdAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryThresholdAlert)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryThresholdAlert
		//----------------------------------------------------------------------------
		return UpdateInventoryThresholdAlert(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Warehouse on a InventoryThresholdAlert
//----------------------------------------------------------------------------
func AssignWarehouseToInventoryThresholdAlert( inventoryThresholdAlertId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryThresholdAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryThresholdAlert(inventoryThresholdAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryThresholdAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryThresholdAlert)

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
			// assign the Warehouse	to the InventoryThresholdAlert
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryThresholdAlert
			//----------------------------------------------------------------------------
			return UpdateInventoryThresholdAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a InventoryThresholdAlert
//----------------------------------------------------------------------------
func UnassignWarehouseFromInventoryThresholdAlert(inventoryThresholdAlertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryThresholdAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryThresholdAlert(inventoryThresholdAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryThresholdAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryThresholdAlert)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryThresholdAlert
		//----------------------------------------------------------------------------
		return UpdateInventoryThresholdAlert(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Location on a InventoryThresholdAlert
//----------------------------------------------------------------------------
func AssignLocationToInventoryThresholdAlert( inventoryThresholdAlertId uint64, locationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryThresholdAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryThresholdAlert(inventoryThresholdAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryThresholdAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryThresholdAlert)

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
			// assign the Location	to the InventoryThresholdAlert
			//----------------------------------------------------------------------------
			parentObj.Location = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryThresholdAlert
			//----------------------------------------------------------------------------
			return UpdateInventoryThresholdAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Location", locationId )
			return utils.RequestResult{false, msg, "assignLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Location on a InventoryThresholdAlert
//----------------------------------------------------------------------------
func UnassignLocationFromInventoryThresholdAlert(inventoryThresholdAlertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryThresholdAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryThresholdAlert(inventoryThresholdAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryThresholdAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryThresholdAlert)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the Location
		//----------------------------------------------------------------------------
		parentObj.Location = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Location
		//----------------------------------------------------------------------------
		parentObj.LocationId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryThresholdAlert
		//----------------------------------------------------------------------------
		return UpdateInventoryThresholdAlert(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a RelatedPolicy on a InventoryThresholdAlert
//----------------------------------------------------------------------------
func AssignRelatedPolicyToInventoryThresholdAlert( inventoryThresholdAlertId uint64, relatedPolicyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InventoryThresholdAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryThresholdAlert(inventoryThresholdAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryThresholdAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryThresholdAlert)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ReplenishmentPolicy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ReplenishmentPolicy with a
		// matching relatedPolicyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, relatedPolicyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the RelatedPolicy	to the InventoryThresholdAlert
			//----------------------------------------------------------------------------
			parentObj.RelatedPolicy = &childObj

			//----------------------------------------------------------------------------
			// save the InventoryThresholdAlert
			//----------------------------------------------------------------------------
			return UpdateInventoryThresholdAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedPolicy", relatedPolicyId )
			return utils.RequestResult{false, msg, "assignRelatedPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a RelatedPolicy on a InventoryThresholdAlert
//----------------------------------------------------------------------------
func UnassignRelatedPolicyFromInventoryThresholdAlert(inventoryThresholdAlertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InventoryThresholdAlert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInventoryThresholdAlert(inventoryThresholdAlertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InventoryThresholdAlert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InventoryThresholdAlert)

		//----------------------------------------------------------------------------
		// assign an empty ReplenishmentPolicy to the RelatedPolicy
		//----------------------------------------------------------------------------
		parentObj.RelatedPolicy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the RelatedPolicy
		//----------------------------------------------------------------------------
		parentObj.RelatedPolicyId = nil;

		//----------------------------------------------------------------------------
		// save the InventoryThresholdAlert
		//----------------------------------------------------------------------------
		return UpdateInventoryThresholdAlert(parentObj)

	} else {
		return parentRequestResult
	}

}


