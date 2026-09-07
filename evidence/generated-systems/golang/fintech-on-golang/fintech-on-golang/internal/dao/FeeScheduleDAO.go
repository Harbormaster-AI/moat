package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FeeScheduleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFeeSchedule - creates a new db entry
//----------------------------------------------------------------------------
func CreateFeeSchedule(obj model.FeeSchedule)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FeeSchedule with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FeeSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFeeSchedule", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFeeSchedule - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFeeSchedule(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FeeSchedule

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FeeSchedule with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FeeSchedule using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FeeSchedule using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFeeSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFeeSchedule - returns all
//----------------------------------------------------------------------------
func GetAllFeeSchedule()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FeeSchedule

	//----------------------------------------------------------------------------
	// Request the ORM to find all FeeSchedule
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FeeSchedule" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FeeSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFeeSchedule", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFeeSchedule - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFeeSchedule(obj model.FeeSchedule)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FeeSchedule using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FeeSchedule using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFeeSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFeeSchedule - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFeeSchedule(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FeeSchedule with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFeeSchedule(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeeSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FeeSchedule)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FeeSchedule using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FeeSchedule using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFeeSchedule", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a PricingPlan on a FeeSchedule
//----------------------------------------------------------------------------
func AssignPricingPlanToFeeSchedule( feeScheduleId uint64, pricingPlanId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FeeSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeeSchedule(feeScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeeSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeeSchedule)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PricingPlan

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PricingPlan with a
		// matching pricingPlanId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, pricingPlanId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PricingPlan	to the FeeSchedule
			//----------------------------------------------------------------------------
			parentObj.PricingPlan = &childObj

			//----------------------------------------------------------------------------
			// save the FeeSchedule
			//----------------------------------------------------------------------------
			return UpdateFeeSchedule(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PricingPlan", pricingPlanId )
			return utils.RequestResult{false, msg, "assignPricingPlan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PricingPlan on a FeeSchedule
//----------------------------------------------------------------------------
func UnassignPricingPlanFromFeeSchedule(feeScheduleId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FeeSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFeeSchedule(feeScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FeeSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FeeSchedule)

		//----------------------------------------------------------------------------
		// assign an empty PricingPlan to the PricingPlan
		//----------------------------------------------------------------------------
		parentObj.PricingPlan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PricingPlan
		//----------------------------------------------------------------------------
		parentObj.PricingPlanId = nil;

		//----------------------------------------------------------------------------
		// save the FeeSchedule
		//----------------------------------------------------------------------------
		return UpdateFeeSchedule(parentObj)

	} else {
		return parentRequestResult
	}

}


