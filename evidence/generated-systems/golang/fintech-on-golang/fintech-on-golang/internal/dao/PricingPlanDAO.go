package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PricingPlanDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePricingPlan - creates a new db entry
//----------------------------------------------------------------------------
func CreatePricingPlan(obj model.PricingPlan)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PricingPlan with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PricingPlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePricingPlan", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPricingPlan - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPricingPlan(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PricingPlan

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PricingPlan with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PricingPlan using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PricingPlan using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPricingPlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPricingPlan - returns all
//----------------------------------------------------------------------------
func GetAllPricingPlan()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PricingPlan

	//----------------------------------------------------------------------------
	// Request the ORM to find all PricingPlan
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PricingPlan" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PricingPlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPricingPlan", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePricingPlan - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePricingPlan(obj model.PricingPlan)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PricingPlan using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PricingPlan using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePricingPlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePricingPlan - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePricingPlan(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PricingPlan with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPricingPlan(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PricingPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PricingPlan)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PricingPlan using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PricingPlan using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePricingPlan", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ProductOffering on a PricingPlan
//----------------------------------------------------------------------------
func AssignProductOfferingToPricingPlan( pricingPlanId uint64, productOfferingId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PricingPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPricingPlan(pricingPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PricingPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PricingPlan)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ProductOffering

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ProductOffering with a
		// matching productOfferingId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productOfferingId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ProductOffering	to the PricingPlan
			//----------------------------------------------------------------------------
			parentObj.ProductOffering = &childObj

			//----------------------------------------------------------------------------
			// save the PricingPlan
			//----------------------------------------------------------------------------
			return UpdatePricingPlan(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductOffering", productOfferingId )
			return utils.RequestResult{false, msg, "assignProductOffering", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ProductOffering on a PricingPlan
//----------------------------------------------------------------------------
func UnassignProductOfferingFromPricingPlan(pricingPlanId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PricingPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPricingPlan(pricingPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PricingPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PricingPlan)

		//----------------------------------------------------------------------------
		// assign an empty ProductOffering to the ProductOffering
		//----------------------------------------------------------------------------
		parentObj.ProductOffering = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ProductOffering
		//----------------------------------------------------------------------------
		parentObj.ProductOfferingId = nil;

		//----------------------------------------------------------------------------
		// save the PricingPlan
		//----------------------------------------------------------------------------
		return UpdatePricingPlan(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more feeSchedulesIds as a FeeSchedules to a PricingPlan
//----------------------------------------------------------------------------
func AddFeeSchedulesToPricingPlan ( pricingPlanId uint64, feeSchedulesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PricingPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPricingPlan(pricingPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PricingPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PricingPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( feeSchedulesIds, ",")

		for _, feeSchedulesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeeSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeeSchedule
			// with a matching feeSchedulesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , feeSchedulesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FeeSchedules using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeeSchedules").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeeSchedules", feeSchedulesId )
				return utils.RequestResult{false, msg, "unassignFeeSchedules", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PricingPlan from the gorm
		//----------------------------------------------------------------------------
		return GetPricingPlan(pricingPlanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more feeSchedulesIds as a FeeSchedules from a PricingPlan
//----------------------------------------------------------------------------
func RemoveFeeSchedulesFromPricingPlan( pricingPlanId uint64, feeSchedulesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PricingPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPricingPlan(pricingPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PricingPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PricingPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( feeSchedulesIds, ",")

		for _, feeSchedulesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeeSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeeSchedule
			// with a matching feeSchedulesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , feeSchedulesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FeeScheduleObj from the FeeSchedules array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeeSchedules").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeeSchedules", feeSchedulesId )
				return utils.RequestResult{false, msg, "removeFeeSchedules", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PricingPlan from the gorm
		//----------------------------------------------------------------------------
		return GetPricingPlan(pricingPlanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more limitsIds as a Limits to a PricingPlan
//----------------------------------------------------------------------------
func AddLimitsToPricingPlan ( pricingPlanId uint64, limitsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PricingPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPricingPlan(pricingPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PricingPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PricingPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( limitsIds, ",")

		for _, limitsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.UsageLimit

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a UsageLimit
			// with a matching limitsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , limitsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Limits using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Limits").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Limits", limitsId )
				return utils.RequestResult{false, msg, "unassignLimits", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PricingPlan from the gorm
		//----------------------------------------------------------------------------
		return GetPricingPlan(pricingPlanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more limitsIds as a Limits from a PricingPlan
//----------------------------------------------------------------------------
func RemoveLimitsFromPricingPlan( pricingPlanId uint64, limitsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PricingPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPricingPlan(pricingPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PricingPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PricingPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( limitsIds, ",")

		for _, limitsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.UsageLimit

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a UsageLimit
			// with a matching limitsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , limitsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove UsageLimitObj from the Limits array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Limits").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Limits", limitsId )
				return utils.RequestResult{false, msg, "removeLimits", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PricingPlan from the gorm
		//----------------------------------------------------------------------------
		return GetPricingPlan(pricingPlanId)

	} else {
		return parentRequestResult
	}
}

