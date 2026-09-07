package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MRPRunDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMRPRun - creates a new db entry
//----------------------------------------------------------------------------
func CreateMRPRun(obj model.MRPRun)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a MRPRun with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MRPRun", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMRPRun", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMRPRun - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMRPRun(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MRPRun

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MRPRun with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MRPRun using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MRPRun using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMRPRun", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMRPRun - returns all
//----------------------------------------------------------------------------
func GetAllMRPRun()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MRPRun

	//----------------------------------------------------------------------------
	// Request the ORM to find all MRPRun
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MRPRun" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MRPRun", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMRPRun", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMRPRun - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMRPRun(obj model.MRPRun)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a MRPRun using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MRPRun using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMRPRun", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMRPRun - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMRPRun(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MRPRun with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMRPRun(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MRPRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MRPRun)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MRPRun using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MRPRun using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMRPRun", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Plant on a MRPRun
//----------------------------------------------------------------------------
func AssignPlantToMRPRun( mRPRunId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the MRPRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMRPRun(mRPRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MRPRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MRPRun)

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
			// assign the Plant	to the MRPRun
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the MRPRun
			//----------------------------------------------------------------------------
			return UpdateMRPRun(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a MRPRun
//----------------------------------------------------------------------------
func UnassignPlantFromMRPRun(mRPRunId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MRPRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMRPRun(mRPRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MRPRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MRPRun)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the MRPRun
		//----------------------------------------------------------------------------
		return UpdateMRPRun(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more plannedOrdersIds as a PlannedOrders to a MRPRun
//----------------------------------------------------------------------------
func AddPlannedOrdersToMRPRun ( mRPRunId uint64, plannedOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MRPRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMRPRun(mRPRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MRPRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MRPRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( plannedOrdersIds, ",")

		for _, plannedOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PlannedOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PlannedOrder
			// with a matching plannedOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , plannedOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PlannedOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PlannedOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PlannedOrders", plannedOrdersId )
				return utils.RequestResult{false, msg, "unassignPlannedOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MRPRun from the gorm
		//----------------------------------------------------------------------------
		return GetMRPRun(mRPRunId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more plannedOrdersIds as a PlannedOrders from a MRPRun
//----------------------------------------------------------------------------
func RemovePlannedOrdersFromMRPRun( mRPRunId uint64, plannedOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the MRPRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMRPRun(mRPRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MRPRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MRPRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( plannedOrdersIds, ",")

		for _, plannedOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PlannedOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PlannedOrder
			// with a matching plannedOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , plannedOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PlannedOrderObj from the PlannedOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PlannedOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PlannedOrders", plannedOrdersId )
				return utils.RequestResult{false, msg, "removePlannedOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MRPRun from the gorm
		//----------------------------------------------------------------------------
		return GetMRPRun(mRPRunId)

	} else {
		return parentRequestResult
	}
}

