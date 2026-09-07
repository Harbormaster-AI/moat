package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing UoMConversionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateUoMConversion - creates a new db entry
//----------------------------------------------------------------------------
func CreateUoMConversion(obj model.UoMConversion)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a UoMConversion with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a UoMConversion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateUoMConversion", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetUoMConversion - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetUoMConversion(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.UoMConversion

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a UoMConversion with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a UoMConversion using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a UoMConversion using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetUoMConversion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllUoMConversion - returns all
//----------------------------------------------------------------------------
func GetAllUoMConversion()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.UoMConversion

	//----------------------------------------------------------------------------
	// Request the ORM to find all UoMConversion
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all UoMConversion" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all UoMConversion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllUoMConversion", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateUoMConversion - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateUoMConversion(obj model.UoMConversion)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a UoMConversion using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a UoMConversion using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateUoMConversion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteUoMConversion - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteUoMConversion(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the UoMConversion with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetUoMConversion(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UoMConversion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.UoMConversion)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a UoMConversion using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a UoMConversion using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteUoMConversion", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Sku on a UoMConversion
//----------------------------------------------------------------------------
func AssignSkuToUoMConversion( uoMConversionId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the UoMConversion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUoMConversion(uoMConversionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UoMConversion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.UoMConversion)

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
			// assign the Sku	to the UoMConversion
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the UoMConversion
			//----------------------------------------------------------------------------
			return UpdateUoMConversion(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a UoMConversion
//----------------------------------------------------------------------------
func UnassignSkuFromUoMConversion(uoMConversionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the UoMConversion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUoMConversion(uoMConversionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UoMConversion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.UoMConversion)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the UoMConversion
		//----------------------------------------------------------------------------
		return UpdateUoMConversion(parentObj)

	} else {
		return parentRequestResult
	}

}


