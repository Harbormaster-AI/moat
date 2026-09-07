package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InspectionLotDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInspectionLot - creates a new db entry
//----------------------------------------------------------------------------
func CreateInspectionLot(obj model.InspectionLot)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InspectionLot with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InspectionLot", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInspectionLot", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInspectionLot - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInspectionLot(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InspectionLot

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InspectionLot with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InspectionLot using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InspectionLot using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInspectionLot", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInspectionLot - returns all
//----------------------------------------------------------------------------
func GetAllInspectionLot()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InspectionLot

	//----------------------------------------------------------------------------
	// Request the ORM to find all InspectionLot
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InspectionLot" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InspectionLot", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInspectionLot", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInspectionLot - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInspectionLot(obj model.InspectionLot)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InspectionLot using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InspectionLot using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInspectionLot", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInspectionLot - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInspectionLot(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InspectionLot with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInspectionLot(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionLot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InspectionLot)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InspectionLot using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InspectionLot using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInspectionLot", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Item on a InspectionLot
//----------------------------------------------------------------------------
func AssignItemToInspectionLot( inspectionLotId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InspectionLot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionLot(inspectionLotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionLot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionLot)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching itemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, itemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Item	to the InspectionLot
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the InspectionLot
			//----------------------------------------------------------------------------
			return UpdateInspectionLot(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a InspectionLot
//----------------------------------------------------------------------------
func UnassignItemFromInspectionLot(inspectionLotId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InspectionLot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionLot(inspectionLotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionLot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionLot)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the InspectionLot
		//----------------------------------------------------------------------------
		return UpdateInspectionLot(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a WorkOrder on a InspectionLot
//----------------------------------------------------------------------------
func AssignWorkOrderToInspectionLot( inspectionLotId uint64, workOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InspectionLot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionLot(inspectionLotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionLot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionLot)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.WorkOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a WorkOrder with a
		// matching workOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the WorkOrder	to the InspectionLot
			//----------------------------------------------------------------------------
			parentObj.WorkOrder = &childObj

			//----------------------------------------------------------------------------
			// save the InspectionLot
			//----------------------------------------------------------------------------
			return UpdateInspectionLot(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrder", workOrderId )
			return utils.RequestResult{false, msg, "assignWorkOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkOrder on a InspectionLot
//----------------------------------------------------------------------------
func UnassignWorkOrderFromInspectionLot(inspectionLotId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InspectionLot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionLot(inspectionLotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionLot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionLot)

		//----------------------------------------------------------------------------
		// assign an empty WorkOrder to the WorkOrder
		//----------------------------------------------------------------------------
		parentObj.WorkOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkOrder
		//----------------------------------------------------------------------------
		parentObj.WorkOrderId = nil;

		//----------------------------------------------------------------------------
		// save the InspectionLot
		//----------------------------------------------------------------------------
		return UpdateInspectionLot(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a GoodsReceipt on a InspectionLot
//----------------------------------------------------------------------------
func AssignGoodsReceiptToInspectionLot( inspectionLotId uint64, goodsReceiptId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InspectionLot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionLot(inspectionLotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionLot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionLot)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.GoodsReceipt

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a GoodsReceipt with a
		// matching goodsReceiptId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, goodsReceiptId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the GoodsReceipt	to the InspectionLot
			//----------------------------------------------------------------------------
			parentObj.GoodsReceipt = &childObj

			//----------------------------------------------------------------------------
			// save the InspectionLot
			//----------------------------------------------------------------------------
			return UpdateInspectionLot(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GoodsReceipt", goodsReceiptId )
			return utils.RequestResult{false, msg, "assignGoodsReceipt", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a GoodsReceipt on a InspectionLot
//----------------------------------------------------------------------------
func UnassignGoodsReceiptFromInspectionLot(inspectionLotId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InspectionLot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionLot(inspectionLotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionLot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionLot)

		//----------------------------------------------------------------------------
		// assign an empty GoodsReceipt to the GoodsReceipt
		//----------------------------------------------------------------------------
		parentObj.GoodsReceipt = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the GoodsReceipt
		//----------------------------------------------------------------------------
		parentObj.GoodsReceiptId = nil;

		//----------------------------------------------------------------------------
		// save the InspectionLot
		//----------------------------------------------------------------------------
		return UpdateInspectionLot(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more resultsIds as a Results to a InspectionLot
//----------------------------------------------------------------------------
func AddResultsToInspectionLot ( inspectionLotId uint64, resultsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InspectionLot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionLot(inspectionLotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionLot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionLot)

		// slice the ids on comma with no spaces
		ids := strings.Split( resultsIds, ",")

		for _, resultsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InspectionResult

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InspectionResult
			// with a matching resultsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , resultsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Results using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Results").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Results", resultsId )
				return utils.RequestResult{false, msg, "unassignResults", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InspectionLot from the gorm
		//----------------------------------------------------------------------------
		return GetInspectionLot(inspectionLotId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more resultsIds as a Results from a InspectionLot
//----------------------------------------------------------------------------
func RemoveResultsFromInspectionLot( inspectionLotId uint64, resultsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InspectionLot with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionLot(inspectionLotId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionLot so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionLot)

		// slice the ids on comma with no spaces
		ids := strings.Split( resultsIds, ",")

		for _, resultsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InspectionResult

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InspectionResult
			// with a matching resultsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , resultsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InspectionResultObj from the Results array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Results").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Results", resultsId )
				return utils.RequestResult{false, msg, "removeResults", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InspectionLot from the gorm
		//----------------------------------------------------------------------------
		return GetInspectionLot(inspectionLotId)

	} else {
		return parentRequestResult
	}
}

