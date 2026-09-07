package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ProductionScheduleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateProductionSchedule - creates a new db entry
//----------------------------------------------------------------------------
func CreateProductionSchedule(obj model.ProductionSchedule)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ProductionSchedule with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ProductionSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateProductionSchedule", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetProductionSchedule - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetProductionSchedule(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ProductionSchedule

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ProductionSchedule with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ProductionSchedule using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ProductionSchedule using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetProductionSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllProductionSchedule - returns all
//----------------------------------------------------------------------------
func GetAllProductionSchedule()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ProductionSchedule

	//----------------------------------------------------------------------------
	// Request the ORM to find all ProductionSchedule
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ProductionSchedule" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ProductionSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllProductionSchedule", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateProductionSchedule - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateProductionSchedule(obj model.ProductionSchedule)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ProductionSchedule using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ProductionSchedule using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateProductionSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteProductionSchedule - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteProductionSchedule(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ProductionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetProductionSchedule(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ProductionSchedule)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ProductionSchedule using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ProductionSchedule using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteProductionSchedule", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Plant on a ProductionSchedule
//----------------------------------------------------------------------------
func AssignPlantToProductionSchedule( productionScheduleId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ProductionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionSchedule(productionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionSchedule)

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
			// assign the Plant	to the ProductionSchedule
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the ProductionSchedule
			//----------------------------------------------------------------------------
			return UpdateProductionSchedule(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a ProductionSchedule
//----------------------------------------------------------------------------
func UnassignPlantFromProductionSchedule(productionScheduleId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionSchedule(productionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionSchedule)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the ProductionSchedule
		//----------------------------------------------------------------------------
		return UpdateProductionSchedule(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more workOrdersIds as a WorkOrders to a ProductionSchedule
//----------------------------------------------------------------------------
func AddWorkOrdersToProductionSchedule ( productionScheduleId uint64, workOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ProductionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionSchedule(productionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( workOrdersIds, ",")

		for _, workOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkOrder
			// with a matching workOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the WorkOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrders", workOrdersId )
				return utils.RequestResult{false, msg, "unassignWorkOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ProductionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetProductionSchedule(productionScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more workOrdersIds as a WorkOrders from a ProductionSchedule
//----------------------------------------------------------------------------
func RemoveWorkOrdersFromProductionSchedule( productionScheduleId uint64, workOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ProductionSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProductionSchedule(productionScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ProductionSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ProductionSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( workOrdersIds, ",")

		for _, workOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkOrder
			// with a matching workOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove WorkOrderObj from the WorkOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrders", workOrdersId )
				return utils.RequestResult{false, msg, "removeWorkOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ProductionSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetProductionSchedule(productionScheduleId)

	} else {
		return parentRequestResult
	}
}

