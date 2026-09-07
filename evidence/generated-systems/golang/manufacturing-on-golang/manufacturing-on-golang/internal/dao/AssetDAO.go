package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AssetDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAsset - creates a new db entry
//----------------------------------------------------------------------------
func CreateAsset(obj model.Asset)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Asset with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Asset", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAsset", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAsset - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAsset(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Asset

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Asset with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Asset using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Asset using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAsset", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAsset - returns all
//----------------------------------------------------------------------------
func GetAllAsset()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Asset

	//----------------------------------------------------------------------------
	// Request the ORM to find all Asset
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Asset" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Asset", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAsset", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAsset - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAsset(obj model.Asset)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Asset using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Asset using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAsset", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAsset - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAsset(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Asset with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAsset(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Asset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Asset)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Asset using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Asset using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAsset", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Plant on a Asset
//----------------------------------------------------------------------------
func AssignPlantToAsset( assetId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Asset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAsset(assetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Asset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Asset)

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
			// assign the Plant	to the Asset
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the Asset
			//----------------------------------------------------------------------------
			return UpdateAsset(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a Asset
//----------------------------------------------------------------------------
func UnassignPlantFromAsset(assetId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Asset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAsset(assetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Asset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Asset)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the Asset
		//----------------------------------------------------------------------------
		return UpdateAsset(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a WorkCenter on a Asset
//----------------------------------------------------------------------------
func AssignWorkCenterToAsset( assetId uint64, workCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Asset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAsset(assetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Asset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Asset)

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
			// assign the WorkCenter	to the Asset
			//----------------------------------------------------------------------------
			parentObj.WorkCenter = &childObj

			//----------------------------------------------------------------------------
			// save the Asset
			//----------------------------------------------------------------------------
			return UpdateAsset(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkCenter", workCenterId )
			return utils.RequestResult{false, msg, "assignWorkCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkCenter on a Asset
//----------------------------------------------------------------------------
func UnassignWorkCenterFromAsset(assetId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Asset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAsset(assetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Asset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Asset)

		//----------------------------------------------------------------------------
		// assign an empty WorkCenter to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenterId = nil;

		//----------------------------------------------------------------------------
		// save the Asset
		//----------------------------------------------------------------------------
		return UpdateAsset(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more maintenanceOrdersIds as a MaintenanceOrders to a Asset
//----------------------------------------------------------------------------
func AddMaintenanceOrdersToAsset ( assetId uint64, maintenanceOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Asset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAsset(assetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Asset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Asset)

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
		// retrieve the modified Asset from the gorm
		//----------------------------------------------------------------------------
		return GetAsset(assetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more maintenanceOrdersIds as a MaintenanceOrders from a Asset
//----------------------------------------------------------------------------
func RemoveMaintenanceOrdersFromAsset( assetId uint64, maintenanceOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Asset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAsset(assetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Asset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Asset)

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
		// retrieve the modified Asset from the gorm
		//----------------------------------------------------------------------------
		return GetAsset(assetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more maintenancePlansIds as a MaintenancePlans to a Asset
//----------------------------------------------------------------------------
func AddMaintenancePlansToAsset ( assetId uint64, maintenancePlansIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Asset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAsset(assetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Asset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Asset)

		// slice the ids on comma with no spaces
		ids := strings.Split( maintenancePlansIds, ",")

		for _, maintenancePlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenancePlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenancePlan
			// with a matching maintenancePlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , maintenancePlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the MaintenancePlans using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MaintenancePlans").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MaintenancePlans", maintenancePlansId )
				return utils.RequestResult{false, msg, "unassignMaintenancePlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Asset from the gorm
		//----------------------------------------------------------------------------
		return GetAsset(assetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more maintenancePlansIds as a MaintenancePlans from a Asset
//----------------------------------------------------------------------------
func RemoveMaintenancePlansFromAsset( assetId uint64, maintenancePlansIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Asset with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAsset(assetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Asset so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Asset)

		// slice the ids on comma with no spaces
		ids := strings.Split( maintenancePlansIds, ",")

		for _, maintenancePlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenancePlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenancePlan
			// with a matching maintenancePlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , maintenancePlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MaintenancePlanObj from the MaintenancePlans array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MaintenancePlans").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MaintenancePlans", maintenancePlansId )
				return utils.RequestResult{false, msg, "removeMaintenancePlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Asset from the gorm
		//----------------------------------------------------------------------------
		return GetAsset(assetId)

	} else {
		return parentRequestResult
	}
}

