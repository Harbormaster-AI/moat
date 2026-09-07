package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing WorkCenterDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateWorkCenter - creates a new db entry
//----------------------------------------------------------------------------
func CreateWorkCenter(obj model.WorkCenter)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a WorkCenter with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a WorkCenter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateWorkCenter", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetWorkCenter - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetWorkCenter(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.WorkCenter

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a WorkCenter with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a WorkCenter using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a WorkCenter using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetWorkCenter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllWorkCenter - returns all
//----------------------------------------------------------------------------
func GetAllWorkCenter()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.WorkCenter

	//----------------------------------------------------------------------------
	// Request the ORM to find all WorkCenter
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all WorkCenter" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all WorkCenter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllWorkCenter", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateWorkCenter - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateWorkCenter(obj model.WorkCenter)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a WorkCenter using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a WorkCenter using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateWorkCenter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteWorkCenter - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteWorkCenter(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the WorkCenter with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetWorkCenter(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.WorkCenter)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a WorkCenter using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a WorkCenter using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteWorkCenter", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ProductionLine on a WorkCenter
//----------------------------------------------------------------------------
func AssignProductionLineToWorkCenter( workCenterId uint64, productionLineId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the WorkCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkCenter(workCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkCenter)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ProductionLine

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ProductionLine with a
		// matching productionLineId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productionLineId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ProductionLine	to the WorkCenter
			//----------------------------------------------------------------------------
			parentObj.ProductionLine = &childObj

			//----------------------------------------------------------------------------
			// save the WorkCenter
			//----------------------------------------------------------------------------
			return UpdateWorkCenter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionLine", productionLineId )
			return utils.RequestResult{false, msg, "assignProductionLine", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ProductionLine on a WorkCenter
//----------------------------------------------------------------------------
func UnassignProductionLineFromWorkCenter(workCenterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkCenter(workCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkCenter)

		//----------------------------------------------------------------------------
		// assign an empty ProductionLine to the ProductionLine
		//----------------------------------------------------------------------------
		parentObj.ProductionLine = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ProductionLine
		//----------------------------------------------------------------------------
		parentObj.ProductionLineId = nil;

		//----------------------------------------------------------------------------
		// save the WorkCenter
		//----------------------------------------------------------------------------
		return UpdateWorkCenter(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more assetsIds as a Assets to a WorkCenter
//----------------------------------------------------------------------------
func AddAssetsToWorkCenter ( workCenterId uint64, assetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkCenter(workCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( assetsIds, ",")

		for _, assetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Asset

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Asset
			// with a matching assetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Assets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assets", assetsId )
				return utils.RequestResult{false, msg, "unassignAssets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkCenter from the gorm
		//----------------------------------------------------------------------------
		return GetWorkCenter(workCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more assetsIds as a Assets from a WorkCenter
//----------------------------------------------------------------------------
func RemoveAssetsFromWorkCenter( workCenterId uint64, assetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the WorkCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkCenter(workCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( assetsIds, ",")

		for _, assetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Asset

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Asset
			// with a matching assetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AssetObj from the Assets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assets", assetsId )
				return utils.RequestResult{false, msg, "removeAssets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkCenter from the gorm
		//----------------------------------------------------------------------------
		return GetWorkCenter(workCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more maintenanceOrdersIds as a MaintenanceOrders to a WorkCenter
//----------------------------------------------------------------------------
func AddMaintenanceOrdersToWorkCenter ( workCenterId uint64, maintenanceOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkCenter(workCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( maintenanceOrdersIds, ",")

		for _, maintenanceOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceOrder
			// with a matching maintenanceOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , maintenanceOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the MaintenanceOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MaintenanceOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MaintenanceOrders", maintenanceOrdersId )
				return utils.RequestResult{false, msg, "unassignMaintenanceOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkCenter from the gorm
		//----------------------------------------------------------------------------
		return GetWorkCenter(workCenterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more maintenanceOrdersIds as a MaintenanceOrders from a WorkCenter
//----------------------------------------------------------------------------
func RemoveMaintenanceOrdersFromWorkCenter( workCenterId uint64, maintenanceOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the WorkCenter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkCenter(workCenterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkCenter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkCenter)

		// slice the ids on comma with no spaces
		ids := strings.Split( maintenanceOrdersIds, ",")

		for _, maintenanceOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceOrder
			// with a matching maintenanceOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , maintenanceOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MaintenanceOrderObj from the MaintenanceOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MaintenanceOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MaintenanceOrders", maintenanceOrdersId )
				return utils.RequestResult{false, msg, "removeMaintenanceOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified WorkCenter from the gorm
		//----------------------------------------------------------------------------
		return GetWorkCenter(workCenterId)

	} else {
		return parentRequestResult
	}
}

