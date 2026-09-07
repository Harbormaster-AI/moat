package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MaintenancePlanDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMaintenancePlan - creates a new db entry
//----------------------------------------------------------------------------
func CreateMaintenancePlan(obj model.MaintenancePlan)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a MaintenancePlan with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MaintenancePlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMaintenancePlan", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMaintenancePlan - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMaintenancePlan(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MaintenancePlan

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MaintenancePlan with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MaintenancePlan using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MaintenancePlan using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMaintenancePlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMaintenancePlan - returns all
//----------------------------------------------------------------------------
func GetAllMaintenancePlan()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MaintenancePlan

	//----------------------------------------------------------------------------
	// Request the ORM to find all MaintenancePlan
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MaintenancePlan" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MaintenancePlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMaintenancePlan", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMaintenancePlan - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMaintenancePlan(obj model.MaintenancePlan)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a MaintenancePlan using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MaintenancePlan using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMaintenancePlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMaintenancePlan - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMaintenancePlan(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MaintenancePlan with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMaintenancePlan(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MaintenancePlan)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MaintenancePlan using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MaintenancePlan using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMaintenancePlan", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Asset on a MaintenancePlan
//----------------------------------------------------------------------------
func AssignAssetToMaintenancePlan( maintenancePlanId uint64, assetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MaintenancePlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenancePlan(maintenancePlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenancePlan)

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
			// assign the Asset	to the MaintenancePlan
			//----------------------------------------------------------------------------
			parentObj.Asset = &childObj

			//----------------------------------------------------------------------------
			// save the MaintenancePlan
			//----------------------------------------------------------------------------
			return UpdateMaintenancePlan(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Asset", assetId )
			return utils.RequestResult{false, msg, "assignAsset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Asset on a MaintenancePlan
//----------------------------------------------------------------------------
func UnassignAssetFromMaintenancePlan(maintenancePlanId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenancePlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenancePlan(maintenancePlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenancePlan)

		//----------------------------------------------------------------------------
		// assign an empty Asset to the Asset
		//----------------------------------------------------------------------------
		parentObj.Asset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Asset
		//----------------------------------------------------------------------------
		parentObj.AssetId = nil;

		//----------------------------------------------------------------------------
		// save the MaintenancePlan
		//----------------------------------------------------------------------------
		return UpdateMaintenancePlan(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more maintenanceOrdersIds as a MaintenanceOrders to a MaintenancePlan
//----------------------------------------------------------------------------
func AddMaintenanceOrdersToMaintenancePlan ( maintenancePlanId uint64, maintenanceOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MaintenancePlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenancePlan(maintenancePlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenancePlan)

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
		// retrieve the modified MaintenancePlan from the gorm
		//----------------------------------------------------------------------------
		return GetMaintenancePlan(maintenancePlanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more maintenanceOrdersIds as a MaintenanceOrders from a MaintenancePlan
//----------------------------------------------------------------------------
func RemoveMaintenanceOrdersFromMaintenancePlan( maintenancePlanId uint64, maintenanceOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the MaintenancePlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMaintenancePlan(maintenancePlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MaintenancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MaintenancePlan)

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
		// retrieve the modified MaintenancePlan from the gorm
		//----------------------------------------------------------------------------
		return GetMaintenancePlan(maintenancePlanId)

	} else {
		return parentRequestResult
	}
}

