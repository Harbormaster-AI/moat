package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PlantDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePlant - creates a new db entry
//----------------------------------------------------------------------------
func CreatePlant(obj model.Plant)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Plant with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Plant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePlant", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPlant - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPlant(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Plant

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Plant with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Plant using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Plant using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPlant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPlant - returns all
//----------------------------------------------------------------------------
func GetAllPlant()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Plant

	//----------------------------------------------------------------------------
	// Request the ORM to find all Plant
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Plant" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Plant", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPlant", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePlant - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePlant(obj model.Plant)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Plant using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Plant using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePlant", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePlant - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePlant(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPlant(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Plant)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Plant using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Plant using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePlant", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Enterprise on a Plant
//----------------------------------------------------------------------------
func AssignEnterpriseToPlant( plantId uint64, enterpriseId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Enterprise

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Enterprise with a
		// matching enterpriseId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, enterpriseId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Enterprise	to the Plant
			//----------------------------------------------------------------------------
			parentObj.Enterprise = &childObj

			//----------------------------------------------------------------------------
			// save the Plant
			//----------------------------------------------------------------------------
			return UpdatePlant(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Enterprise", enterpriseId )
			return utils.RequestResult{false, msg, "assignEnterprise", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Enterprise on a Plant
//----------------------------------------------------------------------------
func UnassignEnterpriseFromPlant(plantId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		//----------------------------------------------------------------------------
		// assign an empty Enterprise to the Enterprise
		//----------------------------------------------------------------------------
		parentObj.Enterprise = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Enterprise
		//----------------------------------------------------------------------------
		parentObj.EnterpriseId = nil;

		//----------------------------------------------------------------------------
		// save the Plant
		//----------------------------------------------------------------------------
		return UpdatePlant(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more productionLinesIds as a ProductionLines to a Plant
//----------------------------------------------------------------------------
func AddProductionLinesToPlant ( plantId uint64, productionLinesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		// slice the ids on comma with no spaces
		ids := strings.Split( productionLinesIds, ",")

		for _, productionLinesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductionLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductionLine
			// with a matching productionLinesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productionLinesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ProductionLines using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductionLines").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionLines", productionLinesId )
				return utils.RequestResult{false, msg, "unassignProductionLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more productionLinesIds as a ProductionLines from a Plant
//----------------------------------------------------------------------------
func RemoveProductionLinesFromPlant( plantId uint64, productionLinesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		// slice the ids on comma with no spaces
		ids := strings.Split( productionLinesIds, ",")

		for _, productionLinesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductionLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductionLine
			// with a matching productionLinesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productionLinesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ProductionLineObj from the ProductionLines array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductionLines").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionLines", productionLinesId )
				return utils.RequestResult{false, msg, "removeProductionLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more workCentersIds as a WorkCenters to a Plant
//----------------------------------------------------------------------------
func AddWorkCentersToPlant ( plantId uint64, workCentersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		// slice the ids on comma with no spaces
		ids := strings.Split( workCentersIds, ",")

		for _, workCentersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkCenter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkCenter
			// with a matching workCentersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workCentersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the WorkCenters using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkCenters").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkCenters", workCentersId )
				return utils.RequestResult{false, msg, "unassignWorkCenters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more workCentersIds as a WorkCenters from a Plant
//----------------------------------------------------------------------------
func RemoveWorkCentersFromPlant( plantId uint64, workCentersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		// slice the ids on comma with no spaces
		ids := strings.Split( workCentersIds, ",")

		for _, workCentersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkCenter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkCenter
			// with a matching workCentersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workCentersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove WorkCenterObj from the WorkCenters array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkCenters").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkCenters", workCentersId )
				return utils.RequestResult{false, msg, "removeWorkCenters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more warehousesIds as a Warehouses to a Plant
//----------------------------------------------------------------------------
func AddWarehousesToPlant ( plantId uint64, warehousesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		// slice the ids on comma with no spaces
		ids := strings.Split( warehousesIds, ",")

		for _, warehousesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Warehouse

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Warehouse
			// with a matching warehousesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , warehousesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Warehouses using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Warehouses").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouses", warehousesId )
				return utils.RequestResult{false, msg, "unassignWarehouses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more warehousesIds as a Warehouses from a Plant
//----------------------------------------------------------------------------
func RemoveWarehousesFromPlant( plantId uint64, warehousesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		// slice the ids on comma with no spaces
		ids := strings.Split( warehousesIds, ",")

		for _, warehousesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Warehouse

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Warehouse
			// with a matching warehousesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , warehousesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove WarehouseObj from the Warehouses array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Warehouses").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warehouses", warehousesId )
				return utils.RequestResult{false, msg, "removeWarehouses", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more assetsIds as a Assets to a Plant
//----------------------------------------------------------------------------
func AddAssetsToPlant ( plantId uint64, assetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

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
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more assetsIds as a Assets from a Plant
//----------------------------------------------------------------------------
func RemoveAssetsFromPlant( plantId uint64, assetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

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
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more productionSchedulesIds as a ProductionSchedules to a Plant
//----------------------------------------------------------------------------
func AddProductionSchedulesToPlant ( plantId uint64, productionSchedulesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		// slice the ids on comma with no spaces
		ids := strings.Split( productionSchedulesIds, ",")

		for _, productionSchedulesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductionSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductionSchedule
			// with a matching productionSchedulesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productionSchedulesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ProductionSchedules using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductionSchedules").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionSchedules", productionSchedulesId )
				return utils.RequestResult{false, msg, "unassignProductionSchedules", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more productionSchedulesIds as a ProductionSchedules from a Plant
//----------------------------------------------------------------------------
func RemoveProductionSchedulesFromPlant( plantId uint64, productionSchedulesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Plant with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlant(plantId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Plant so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Plant)

		// slice the ids on comma with no spaces
		ids := strings.Split( productionSchedulesIds, ",")

		for _, productionSchedulesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductionSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductionSchedule
			// with a matching productionSchedulesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productionSchedulesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ProductionScheduleObj from the ProductionSchedules array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductionSchedules").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionSchedules", productionSchedulesId )
				return utils.RequestResult{false, msg, "removeProductionSchedules", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Plant from the gorm
		//----------------------------------------------------------------------------
		return GetPlant(plantId)

	} else {
		return parentRequestResult
	}
}

