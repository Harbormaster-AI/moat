package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BuildScheduleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBuildSchedule - creates a new db entry
//----------------------------------------------------------------------------
func CreateBuildSchedule(obj model.BuildSchedule)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BuildSchedule with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BuildSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBuildSchedule", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBuildSchedule - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBuildSchedule(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BuildSchedule

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BuildSchedule with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BuildSchedule using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BuildSchedule using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBuildSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBuildSchedule - returns all
//----------------------------------------------------------------------------
func GetAllBuildSchedule()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BuildSchedule

	//----------------------------------------------------------------------------
	// Request the ORM to find all BuildSchedule
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BuildSchedule" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BuildSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBuildSchedule", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBuildSchedule - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBuildSchedule(obj model.BuildSchedule)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BuildSchedule using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BuildSchedule using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBuildSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBuildSchedule - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBuildSchedule(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BuildSchedule with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBuildSchedule(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BuildSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BuildSchedule)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BuildSchedule using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BuildSchedule using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBuildSchedule", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more productionOrdersIds as a ProductionOrders to a BuildSchedule
//----------------------------------------------------------------------------
func AddProductionOrdersToBuildSchedule ( buildScheduleId uint64, productionOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BuildSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBuildSchedule(buildScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BuildSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BuildSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( productionOrdersIds, ",")

		for _, productionOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductionOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductionOrder
			// with a matching productionOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productionOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ProductionOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductionOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionOrders", productionOrdersId )
				return utils.RequestResult{false, msg, "unassignProductionOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BuildSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetBuildSchedule(buildScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more productionOrdersIds as a ProductionOrders from a BuildSchedule
//----------------------------------------------------------------------------
func RemoveProductionOrdersFromBuildSchedule( buildScheduleId uint64, productionOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BuildSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBuildSchedule(buildScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BuildSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BuildSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( productionOrdersIds, ",")

		for _, productionOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ProductionOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ProductionOrder
			// with a matching productionOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , productionOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ProductionOrderObj from the ProductionOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProductionOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductionOrders", productionOrdersId )
				return utils.RequestResult{false, msg, "removeProductionOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BuildSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetBuildSchedule(buildScheduleId)

	} else {
		return parentRequestResult
	}
}

