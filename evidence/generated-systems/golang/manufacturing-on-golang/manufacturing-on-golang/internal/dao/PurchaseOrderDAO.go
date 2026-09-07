package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PurchaseOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePurchaseOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreatePurchaseOrder(obj model.PurchaseOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PurchaseOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PurchaseOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePurchaseOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPurchaseOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPurchaseOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PurchaseOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PurchaseOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PurchaseOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PurchaseOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPurchaseOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPurchaseOrder - returns all
//----------------------------------------------------------------------------
func GetAllPurchaseOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PurchaseOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all PurchaseOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PurchaseOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PurchaseOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPurchaseOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePurchaseOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePurchaseOrder(obj model.PurchaseOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PurchaseOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PurchaseOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePurchaseOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePurchaseOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePurchaseOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPurchaseOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PurchaseOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PurchaseOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PurchaseOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePurchaseOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Supplier on a PurchaseOrder
//----------------------------------------------------------------------------
func AssignSupplierToPurchaseOrder( purchaseOrderId uint64, supplierId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrder(purchaseOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Supplier

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Supplier with a
		// matching supplierId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, supplierId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Supplier	to the PurchaseOrder
			//----------------------------------------------------------------------------
			parentObj.Supplier = &childObj

			//----------------------------------------------------------------------------
			// save the PurchaseOrder
			//----------------------------------------------------------------------------
			return UpdatePurchaseOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Supplier", supplierId )
			return utils.RequestResult{false, msg, "assignSupplier", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Supplier on a PurchaseOrder
//----------------------------------------------------------------------------
func UnassignSupplierFromPurchaseOrder(purchaseOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrder(purchaseOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrder)

		//----------------------------------------------------------------------------
		// assign an empty Supplier to the Supplier
		//----------------------------------------------------------------------------
		parentObj.Supplier = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Supplier
		//----------------------------------------------------------------------------
		parentObj.SupplierId = nil;

		//----------------------------------------------------------------------------
		// save the PurchaseOrder
		//----------------------------------------------------------------------------
		return UpdatePurchaseOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Plant on a PurchaseOrder
//----------------------------------------------------------------------------
func AssignPlantToPurchaseOrder( purchaseOrderId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrder(purchaseOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Plant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Plant with a
		// matching plantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, plantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Plant	to the PurchaseOrder
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the PurchaseOrder
			//----------------------------------------------------------------------------
			return UpdatePurchaseOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a PurchaseOrder
//----------------------------------------------------------------------------
func UnassignPlantFromPurchaseOrder(purchaseOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrder(purchaseOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrder)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the PurchaseOrder
		//----------------------------------------------------------------------------
		return UpdatePurchaseOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more linesIds as a Lines to a PurchaseOrder
//----------------------------------------------------------------------------
func AddLinesToPurchaseOrder ( purchaseOrderId uint64, linesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrder(purchaseOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PurchaseOrderLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PurchaseOrderLine
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
		// retrieve the modified PurchaseOrder from the gorm
		//----------------------------------------------------------------------------
		return GetPurchaseOrder(purchaseOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more linesIds as a Lines from a PurchaseOrder
//----------------------------------------------------------------------------
func RemoveLinesFromPurchaseOrder( purchaseOrderId uint64, linesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrder(purchaseOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PurchaseOrderLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PurchaseOrderLine
			// with a matching linesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , linesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PurchaseOrderLineObj from the Lines array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lines").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lines", linesId )
				return utils.RequestResult{false, msg, "removeLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PurchaseOrder from the gorm
		//----------------------------------------------------------------------------
		return GetPurchaseOrder(purchaseOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more goodsReceiptsIds as a GoodsReceipts to a PurchaseOrder
//----------------------------------------------------------------------------
func AddGoodsReceiptsToPurchaseOrder ( purchaseOrderId uint64, goodsReceiptsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrder(purchaseOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( goodsReceiptsIds, ",")

		for _, goodsReceiptsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GoodsReceipt

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GoodsReceipt
			// with a matching goodsReceiptsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , goodsReceiptsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the GoodsReceipts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("GoodsReceipts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GoodsReceipts", goodsReceiptsId )
				return utils.RequestResult{false, msg, "unassignGoodsReceipts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PurchaseOrder from the gorm
		//----------------------------------------------------------------------------
		return GetPurchaseOrder(purchaseOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more goodsReceiptsIds as a GoodsReceipts from a PurchaseOrder
//----------------------------------------------------------------------------
func RemoveGoodsReceiptsFromPurchaseOrder( purchaseOrderId uint64, goodsReceiptsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PurchaseOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPurchaseOrder(purchaseOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PurchaseOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PurchaseOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( goodsReceiptsIds, ",")

		for _, goodsReceiptsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GoodsReceipt

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GoodsReceipt
			// with a matching goodsReceiptsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , goodsReceiptsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove GoodsReceiptObj from the GoodsReceipts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("GoodsReceipts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GoodsReceipts", goodsReceiptsId )
				return utils.RequestResult{false, msg, "removeGoodsReceipts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PurchaseOrder from the gorm
		//----------------------------------------------------------------------------
		return GetPurchaseOrder(purchaseOrderId)

	} else {
		return parentRequestResult
	}
}

