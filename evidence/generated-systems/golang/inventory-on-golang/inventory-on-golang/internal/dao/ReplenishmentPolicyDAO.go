package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ReplenishmentPolicyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateReplenishmentPolicy - creates a new db entry
//----------------------------------------------------------------------------
func CreateReplenishmentPolicy(obj model.ReplenishmentPolicy)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ReplenishmentPolicy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ReplenishmentPolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateReplenishmentPolicy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetReplenishmentPolicy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetReplenishmentPolicy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ReplenishmentPolicy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ReplenishmentPolicy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ReplenishmentPolicy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ReplenishmentPolicy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetReplenishmentPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllReplenishmentPolicy - returns all
//----------------------------------------------------------------------------
func GetAllReplenishmentPolicy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ReplenishmentPolicy

	//----------------------------------------------------------------------------
	// Request the ORM to find all ReplenishmentPolicy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ReplenishmentPolicy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ReplenishmentPolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllReplenishmentPolicy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateReplenishmentPolicy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateReplenishmentPolicy(obj model.ReplenishmentPolicy)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ReplenishmentPolicy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ReplenishmentPolicy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateReplenishmentPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteReplenishmentPolicy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteReplenishmentPolicy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ReplenishmentPolicy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetReplenishmentPolicy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReplenishmentPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ReplenishmentPolicy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ReplenishmentPolicy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ReplenishmentPolicy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteReplenishmentPolicy", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Sku on a ReplenishmentPolicy
//----------------------------------------------------------------------------
func AssignSkuToReplenishmentPolicy( replenishmentPolicyId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ReplenishmentPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReplenishmentPolicy(replenishmentPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReplenishmentPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReplenishmentPolicy)

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
			// assign the Sku	to the ReplenishmentPolicy
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the ReplenishmentPolicy
			//----------------------------------------------------------------------------
			return UpdateReplenishmentPolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a ReplenishmentPolicy
//----------------------------------------------------------------------------
func UnassignSkuFromReplenishmentPolicy(replenishmentPolicyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ReplenishmentPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReplenishmentPolicy(replenishmentPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReplenishmentPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReplenishmentPolicy)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the ReplenishmentPolicy
		//----------------------------------------------------------------------------
		return UpdateReplenishmentPolicy(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Warehouse on a ReplenishmentPolicy
//----------------------------------------------------------------------------
func AssignWarehouseToReplenishmentPolicy( replenishmentPolicyId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ReplenishmentPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReplenishmentPolicy(replenishmentPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReplenishmentPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReplenishmentPolicy)

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
			// assign the Warehouse	to the ReplenishmentPolicy
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the ReplenishmentPolicy
			//----------------------------------------------------------------------------
			return UpdateReplenishmentPolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a ReplenishmentPolicy
//----------------------------------------------------------------------------
func UnassignWarehouseFromReplenishmentPolicy(replenishmentPolicyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ReplenishmentPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReplenishmentPolicy(replenishmentPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReplenishmentPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReplenishmentPolicy)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the ReplenishmentPolicy
		//----------------------------------------------------------------------------
		return UpdateReplenishmentPolicy(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Location on a ReplenishmentPolicy
//----------------------------------------------------------------------------
func AssignLocationToReplenishmentPolicy( replenishmentPolicyId uint64, locationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ReplenishmentPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReplenishmentPolicy(replenishmentPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReplenishmentPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReplenishmentPolicy)

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
			// assign the Location	to the ReplenishmentPolicy
			//----------------------------------------------------------------------------
			parentObj.Location = &childObj

			//----------------------------------------------------------------------------
			// save the ReplenishmentPolicy
			//----------------------------------------------------------------------------
			return UpdateReplenishmentPolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Location", locationId )
			return utils.RequestResult{false, msg, "assignLocation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Location on a ReplenishmentPolicy
//----------------------------------------------------------------------------
func UnassignLocationFromReplenishmentPolicy(replenishmentPolicyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ReplenishmentPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReplenishmentPolicy(replenishmentPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReplenishmentPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReplenishmentPolicy)

		//----------------------------------------------------------------------------
		// assign an empty StorageLocation to the Location
		//----------------------------------------------------------------------------
		parentObj.Location = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Location
		//----------------------------------------------------------------------------
		parentObj.LocationId = nil;

		//----------------------------------------------------------------------------
		// save the ReplenishmentPolicy
		//----------------------------------------------------------------------------
		return UpdateReplenishmentPolicy(parentObj)

	} else {
		return parentRequestResult
	}

}


