package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MaintenanceOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMaintenanceOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateMaintenanceOrder(obj model.MaintenanceOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a MaintenanceOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MaintenanceOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMaintenanceOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMaintenanceOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMaintenanceOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MaintenanceOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MaintenanceOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MaintenanceOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MaintenanceOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMaintenanceOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMaintenanceOrder - returns all
//----------------------------------------------------------------------------
func GetAllMaintenanceOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MaintenanceOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all MaintenanceOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MaintenanceOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MaintenanceOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMaintenanceOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMaintenanceOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMaintenanceOrder(obj model.MaintenanceOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a MaintenanceOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MaintenanceOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMaintenanceOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMaintenanceOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMaintenanceOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMaintenanceOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MaintenanceOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MaintenanceOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MaintenanceOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMaintenanceOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Asset on a MaintenanceOrder
//----------------------------------------------------------------------------
func AssignAssetToMaintenanceOrder( maintenanceOrderId uint64, assetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceOrder(maintenanceOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Asset

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Asset with a
		// matching assetId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, assetId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Asset	to the MaintenanceOrder
			//----------------------------------------------------------------------------
			parentObj.Asset = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenanceOrder
			//----------------------------------------------------------------------------
			return UpdateMaintenanceOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Asset", assetId )
			return utils.RequestResult{false, msg, "assignAsset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Asset on a MaintenanceOrder
//----------------------------------------------------------------------------
func UnassignAssetFromMaintenanceOrder(maintenanceOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceOrder(maintenanceOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceOrder)

		//----------------------------------------------------------------------------
		// assign an empty Asset to the Asset
		//----------------------------------------------------------------------------
		parentObj.Asset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Asset
		//----------------------------------------------------------------------------
		parentObj.AssetId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenanceOrder
		//----------------------------------------------------------------------------
		return UpdateMaintenanceOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Plan on a MaintenanceOrder
//----------------------------------------------------------------------------
func AssignPlanToMaintenanceOrder( maintenanceOrderId uint64, planId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceOrder(maintenanceOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.MaintenancePlan

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a MaintenancePlan with a
		// matching planId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, planId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Plan	to the MaintenanceOrder
			//----------------------------------------------------------------------------
			parentObj.Plan = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenanceOrder
			//----------------------------------------------------------------------------
			return UpdateMaintenanceOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plan", planId )
			return utils.RequestResult{false, msg, "assignPlan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plan on a MaintenanceOrder
//----------------------------------------------------------------------------
func UnassignPlanFromMaintenanceOrder(maintenanceOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceOrder(maintenanceOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceOrder)

		//----------------------------------------------------------------------------
		// assign an empty MaintenancePlan to the Plan
		//----------------------------------------------------------------------------
		parentObj.Plan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plan
		//----------------------------------------------------------------------------
		parentObj.PlanId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenanceOrder
		//----------------------------------------------------------------------------
		return UpdateMaintenanceOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a WorkCenter on a MaintenanceOrder
//----------------------------------------------------------------------------
func AssignWorkCenterToMaintenanceOrder( maintenanceOrderId uint64, workCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceOrder(maintenanceOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.WorkCenter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a WorkCenter with a
		// matching workCenterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workCenterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the WorkCenter	to the MaintenanceOrder
			//----------------------------------------------------------------------------
			parentObj.WorkCenter = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenanceOrder
			//----------------------------------------------------------------------------
			return UpdateMaintenanceOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkCenter", workCenterId )
			return utils.RequestResult{false, msg, "assignWorkCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkCenter on a MaintenanceOrder
//----------------------------------------------------------------------------
func UnassignWorkCenterFromMaintenanceOrder(maintenanceOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenanceOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenanceOrder(maintenanceOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenanceOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenanceOrder)

		//----------------------------------------------------------------------------
		// assign an empty WorkCenter to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenterId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenanceOrder
		//----------------------------------------------------------------------------
		return UpdateMaintenanceOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


