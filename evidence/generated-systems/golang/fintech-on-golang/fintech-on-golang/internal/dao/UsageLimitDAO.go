package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing UsageLimitDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateUsageLimit - creates a new db entry
//----------------------------------------------------------------------------
func CreateUsageLimit(obj model.UsageLimit)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a UsageLimit with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a UsageLimit", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateUsageLimit", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetUsageLimit - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetUsageLimit(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.UsageLimit

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a UsageLimit with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a UsageLimit using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a UsageLimit using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetUsageLimit", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllUsageLimit - returns all
//----------------------------------------------------------------------------
func GetAllUsageLimit()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.UsageLimit

	//----------------------------------------------------------------------------
	// Request the ORM to find all UsageLimit
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all UsageLimit" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all UsageLimit", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllUsageLimit", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateUsageLimit - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateUsageLimit(obj model.UsageLimit)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a UsageLimit using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a UsageLimit using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateUsageLimit", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteUsageLimit - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteUsageLimit(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the UsageLimit with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetUsageLimit(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UsageLimit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.UsageLimit)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a UsageLimit using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a UsageLimit using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteUsageLimit", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a PricingPlan on a UsageLimit
//----------------------------------------------------------------------------
func AssignPricingPlanToUsageLimit( usageLimitId uint64, pricingPlanId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the UsageLimit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUsageLimit(usageLimitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UsageLimit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.UsageLimit)

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
			// assign the PricingPlan	to the UsageLimit
			//----------------------------------------------------------------------------
			parentObj.PricingPlan = &childObj

			//----------------------------------------------------------------------------
			// save the UsageLimit
			//----------------------------------------------------------------------------
			return UpdateUsageLimit(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PricingPlan", pricingPlanId )
			return utils.RequestResult{false, msg, "assignPricingPlan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PricingPlan on a UsageLimit
//----------------------------------------------------------------------------
func UnassignPricingPlanFromUsageLimit(usageLimitId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the UsageLimit with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUsageLimit(usageLimitId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UsageLimit so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.UsageLimit)

		//----------------------------------------------------------------------------
		// assign an empty PricingPlan to the PricingPlan
		//----------------------------------------------------------------------------
		parentObj.PricingPlan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PricingPlan
		//----------------------------------------------------------------------------
		parentObj.PricingPlanId = nil;

		//----------------------------------------------------------------------------
		// save the UsageLimit
		//----------------------------------------------------------------------------
		return UpdateUsageLimit(parentObj)

	} else {
		return parentRequestResult
	}

}


