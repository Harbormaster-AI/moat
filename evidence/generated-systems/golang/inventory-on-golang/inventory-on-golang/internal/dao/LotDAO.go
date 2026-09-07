package dao

import (
    "inventory-on-golang/internal/model"
    "inventory-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LotDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLot - creates a new db entry
//----------------------------------------------------------------------------
func CreateLot(obj model.Lot)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Lot with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Lot", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLot", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLot - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLot(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Lot

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Lot with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Lot using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Lot using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLot", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLot - returns all
//----------------------------------------------------------------------------
func GetAllLot()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Lot

	//----------------------------------------------------------------------------
	// Request the ORM to find all Lot
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Lot" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Lot", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLot", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLot - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLot(obj model.Lot)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Lot using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Lot using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLot", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLot - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLot(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Lot with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLot(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Lot)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Lot using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Lot using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLot", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Sku on a Lot
//----------------------------------------------------------------------------
func AssignSkuToLot( lotId uint64, skuId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Lot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLot(lotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lot)

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
			// assign the Sku	to the Lot
			//----------------------------------------------------------------------------
			parentObj.Sku = &childObj

			//----------------------------------------------------------------------------
			// save the Lot
			//----------------------------------------------------------------------------
			return UpdateLot(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sku", skuId )
			return utils.RequestResult{false, msg, "assignSku", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Sku on a Lot
//----------------------------------------------------------------------------
func UnassignSkuFromLot(lotId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLot(lotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lot)

		//----------------------------------------------------------------------------
		// assign an empty StockKeepingUnit to the Sku
		//----------------------------------------------------------------------------
		parentObj.Sku = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Sku
		//----------------------------------------------------------------------------
		parentObj.SkuId = nil;

		//----------------------------------------------------------------------------
		// save the Lot
		//----------------------------------------------------------------------------
		return UpdateLot(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more inventoryItemsIds as a InventoryItems to a Lot
//----------------------------------------------------------------------------
func AddInventoryItemsToLot ( lotId uint64, inventoryItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Lot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLot(lotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lot)

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
		// retrieve the modified Lot from the gorm
		//----------------------------------------------------------------------------
		return GetLot(lotId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventoryItemsIds as a InventoryItems from a Lot
//----------------------------------------------------------------------------
func RemoveInventoryItemsFromLot( lotId uint64, inventoryItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Lot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLot(lotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Lot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Lot)

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
		// retrieve the modified Lot from the gorm
		//----------------------------------------------------------------------------
		return GetLot(lotId)

	} else {
		return parentRequestResult
	}
}

