package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing GoodsReceiptDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateGoodsReceipt - creates a new db entry
//----------------------------------------------------------------------------
func CreateGoodsReceipt(obj model.GoodsReceipt)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a GoodsReceipt with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a GoodsReceipt", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateGoodsReceipt", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetGoodsReceipt - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetGoodsReceipt(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.GoodsReceipt

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a GoodsReceipt with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a GoodsReceipt using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a GoodsReceipt using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetGoodsReceipt", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllGoodsReceipt - returns all
//----------------------------------------------------------------------------
func GetAllGoodsReceipt()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.GoodsReceipt

	//----------------------------------------------------------------------------
	// Request the ORM to find all GoodsReceipt
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all GoodsReceipt" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all GoodsReceipt", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllGoodsReceipt", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateGoodsReceipt - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateGoodsReceipt(obj model.GoodsReceipt)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a GoodsReceipt using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a GoodsReceipt using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateGoodsReceipt", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteGoodsReceipt - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteGoodsReceipt(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceipt with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetGoodsReceipt(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceipt so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.GoodsReceipt)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a GoodsReceipt using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a GoodsReceipt using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteGoodsReceipt", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a PurchaseOrder on a GoodsReceipt
//----------------------------------------------------------------------------
func AssignPurchaseOrderToGoodsReceipt( goodsReceiptId uint64, purchaseOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceipt with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceipt(goodsReceiptId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceipt so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceipt)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PurchaseOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PurchaseOrder with a
		// matching purchaseOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, purchaseOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PurchaseOrder	to the GoodsReceipt
			//----------------------------------------------------------------------------
			parentObj.PurchaseOrder = &childObj

			//----------------------------------------------------------------------------
			// save the GoodsReceipt
			//----------------------------------------------------------------------------
			return UpdateGoodsReceipt(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PurchaseOrder", purchaseOrderId )
			return utils.RequestResult{false, msg, "assignPurchaseOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PurchaseOrder on a GoodsReceipt
//----------------------------------------------------------------------------
func UnassignPurchaseOrderFromGoodsReceipt(goodsReceiptId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceipt with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceipt(goodsReceiptId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceipt so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceipt)

		//----------------------------------------------------------------------------
		// assign an empty PurchaseOrder to the PurchaseOrder
		//----------------------------------------------------------------------------
		parentObj.PurchaseOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PurchaseOrder
		//----------------------------------------------------------------------------
		parentObj.PurchaseOrderId = nil;

		//----------------------------------------------------------------------------
		// save the GoodsReceipt
		//----------------------------------------------------------------------------
		return UpdateGoodsReceipt(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Warehouse on a GoodsReceipt
//----------------------------------------------------------------------------
func AssignWarehouseToGoodsReceipt( goodsReceiptId uint64, warehouseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceipt with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceipt(goodsReceiptId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceipt so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceipt)

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
			// assign the Warehouse	to the GoodsReceipt
			//----------------------------------------------------------------------------
			parentObj.Warehouse = &childObj

			//----------------------------------------------------------------------------
			// save the GoodsReceipt
			//----------------------------------------------------------------------------
			return UpdateGoodsReceipt(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouse", warehouseId )
			return utils.RequestResult{false, msg, "assignWarehouse", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warehouse on a GoodsReceipt
//----------------------------------------------------------------------------
func UnassignWarehouseFromGoodsReceipt(goodsReceiptId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceipt with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceipt(goodsReceiptId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceipt so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceipt)

		//----------------------------------------------------------------------------
		// assign an empty Warehouse to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.Warehouse = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warehouse
		//----------------------------------------------------------------------------
		parentObj.WarehouseId = nil;

		//----------------------------------------------------------------------------
		// save the GoodsReceipt
		//----------------------------------------------------------------------------
		return UpdateGoodsReceipt(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more linesIds as a Lines to a GoodsReceipt
//----------------------------------------------------------------------------
func AddLinesToGoodsReceipt ( goodsReceiptId uint64, linesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the GoodsReceipt with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceipt(goodsReceiptId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceipt so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceipt)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GoodsReceiptLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GoodsReceiptLine
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
		// retrieve the modified GoodsReceipt from the gorm
		//----------------------------------------------------------------------------
		return GetGoodsReceipt(goodsReceiptId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more linesIds as a Lines from a GoodsReceipt
//----------------------------------------------------------------------------
func RemoveLinesFromGoodsReceipt( goodsReceiptId uint64, linesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the GoodsReceipt with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetGoodsReceipt(goodsReceiptId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.GoodsReceipt so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.GoodsReceipt)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GoodsReceiptLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GoodsReceiptLine
			// with a matching linesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , linesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove GoodsReceiptLineObj from the Lines array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lines").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lines", linesId )
				return utils.RequestResult{false, msg, "removeLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified GoodsReceipt from the gorm
		//----------------------------------------------------------------------------
		return GetGoodsReceipt(goodsReceiptId)

	} else {
		return parentRequestResult
	}
}

