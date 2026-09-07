package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RateDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRate - creates a new db entry
//----------------------------------------------------------------------------
func CreateRate(obj model.Rate)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Rate with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Rate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRate", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRate - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRate(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Rate

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Rate with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Rate using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Rate using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRate - returns all
//----------------------------------------------------------------------------
func GetAllRate()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Rate

	//----------------------------------------------------------------------------
	// Request the ORM to find all Rate
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Rate" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Rate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRate", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRate - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRate(obj model.Rate)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Rate using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Rate using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRate - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRate(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Rate with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRate(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Rate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Rate)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Rate using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Rate using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRate", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a RateCard on a Rate
//----------------------------------------------------------------------------
func AssignRateCardToRate( rateId uint64, rateCardId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Rate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRate(rateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Rate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Rate)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.RateCard

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a RateCard with a
		// matching rateCardId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, rateCardId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the RateCard	to the Rate
			//----------------------------------------------------------------------------
			parentObj.RateCard = &childObj

			//----------------------------------------------------------------------------
			// save the Rate
			//----------------------------------------------------------------------------
			return UpdateRate(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RateCard", rateCardId )
			return utils.RequestResult{false, msg, "assignRateCard", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a RateCard on a Rate
//----------------------------------------------------------------------------
func UnassignRateCardFromRate(rateId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Rate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRate(rateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Rate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Rate)

		//----------------------------------------------------------------------------
		// assign an empty RateCard to the RateCard
		//----------------------------------------------------------------------------
		parentObj.RateCard = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the RateCard
		//----------------------------------------------------------------------------
		parentObj.RateCardId = nil;

		//----------------------------------------------------------------------------
		// save the Rate
		//----------------------------------------------------------------------------
		return UpdateRate(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a AdSlot on a Rate
//----------------------------------------------------------------------------
func AssignAdSlotToRate( rateId uint64, adSlotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Rate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRate(rateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Rate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Rate)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AdSlot

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AdSlot with a
		// matching adSlotId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, adSlotId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AdSlot	to the Rate
			//----------------------------------------------------------------------------
			parentObj.AdSlot = &childObj

			//----------------------------------------------------------------------------
			// save the Rate
			//----------------------------------------------------------------------------
			return UpdateRate(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AdSlot", adSlotId )
			return utils.RequestResult{false, msg, "assignAdSlot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AdSlot on a Rate
//----------------------------------------------------------------------------
func UnassignAdSlotFromRate(rateId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Rate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRate(rateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Rate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Rate)

		//----------------------------------------------------------------------------
		// assign an empty AdSlot to the AdSlot
		//----------------------------------------------------------------------------
		parentObj.AdSlot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AdSlot
		//----------------------------------------------------------------------------
		parentObj.AdSlotId = nil;

		//----------------------------------------------------------------------------
		// save the Rate
		//----------------------------------------------------------------------------
		return UpdateRate(parentObj)

	} else {
		return parentRequestResult
	}

}


