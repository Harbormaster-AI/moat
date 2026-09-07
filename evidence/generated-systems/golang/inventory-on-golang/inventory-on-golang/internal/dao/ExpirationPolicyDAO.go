package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ExpirationPolicyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateExpirationPolicy - creates a new db entry
//----------------------------------------------------------------------------
func CreateExpirationPolicy(obj model.ExpirationPolicy)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ExpirationPolicy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ExpirationPolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateExpirationPolicy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetExpirationPolicy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetExpirationPolicy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ExpirationPolicy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ExpirationPolicy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ExpirationPolicy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ExpirationPolicy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetExpirationPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllExpirationPolicy - returns all
//----------------------------------------------------------------------------
func GetAllExpirationPolicy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ExpirationPolicy

	//----------------------------------------------------------------------------
	// Request the ORM to find all ExpirationPolicy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ExpirationPolicy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ExpirationPolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllExpirationPolicy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateExpirationPolicy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateExpirationPolicy(obj model.ExpirationPolicy)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ExpirationPolicy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ExpirationPolicy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateExpirationPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteExpirationPolicy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteExpirationPolicy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ExpirationPolicy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetExpirationPolicy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExpirationPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ExpirationPolicy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ExpirationPolicy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ExpirationPolicy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteExpirationPolicy", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Sku on a ExpirationPolicy
//----------------------------------------------------------------------------
func AssignSkuToExpirationPolicy( expirationPolicyId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ExpirationPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExpirationPolicy(expirationPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExpirationPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExpirationPolicy)

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
			// assign the Sku	to the ExpirationPolicy
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the ExpirationPolicy
			//----------------------------------------------------------------------------
			return UpdateExpirationPolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a ExpirationPolicy
//----------------------------------------------------------------------------
func UnassignSkuFromExpirationPolicy(expirationPolicyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExpirationPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExpirationPolicy(expirationPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExpirationPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExpirationPolicy)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the ExpirationPolicy
		//----------------------------------------------------------------------------
		return UpdateExpirationPolicy(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Warehouse on a ExpirationPolicy
//----------------------------------------------------------------------------
func AssignWarehouseToExpirationPolicy( expirationPolicyId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ExpirationPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExpirationPolicy(expirationPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExpirationPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExpirationPolicy)

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
			// assign the Warehouse	to the ExpirationPolicy
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the ExpirationPolicy
			//----------------------------------------------------------------------------
			return UpdateExpirationPolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a ExpirationPolicy
//----------------------------------------------------------------------------
func UnassignWarehouseFromExpirationPolicy(expirationPolicyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExpirationPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExpirationPolicy(expirationPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExpirationPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExpirationPolicy)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the ExpirationPolicy
		//----------------------------------------------------------------------------
		return UpdateExpirationPolicy(parentObj)

	} else {
		return parentRequestResult
	}

}


