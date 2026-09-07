package dao

import (
    "advertising-on-golang/internal/model"
    "advertising-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RateCardDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRateCard - creates a new db entry
//----------------------------------------------------------------------------
func CreateRateCard(obj model.RateCard)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RateCard with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RateCard", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRateCard", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRateCard - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRateCard(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RateCard

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RateCard with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RateCard using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RateCard using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRateCard", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRateCard - returns all
//----------------------------------------------------------------------------
func GetAllRateCard()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RateCard

	//----------------------------------------------------------------------------
	// Request the ORM to find all RateCard
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all RateCard" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RateCard", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRateCard", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRateCard - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRateCard(obj model.RateCard)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RateCard using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RateCard using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRateCard", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRateCard - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRateCard(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RateCard with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRateCard(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RateCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.RateCard)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RateCard using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RateCard using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRateCard", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Publisher on a RateCard
//----------------------------------------------------------------------------
func AssignPublisherToRateCard( rateCardId uint64, publisherId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RateCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRateCard(rateCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RateCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RateCard)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Publisher

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Publisher with a
		// matching publisherId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, publisherId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Publisher	to the RateCard
			//----------------------------------------------------------------------------
			parentObj.Publisher = &childObj

			//----------------------------------------------------------------------------
			// save the RateCard
			//----------------------------------------------------------------------------
			return UpdateRateCard(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Publisher", publisherId )
			return utils.RequestResult{false, msg, "assignPublisher", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Publisher on a RateCard
//----------------------------------------------------------------------------
func UnassignPublisherFromRateCard(rateCardId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RateCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRateCard(rateCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RateCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RateCard)

		//----------------------------------------------------------------------------
		// assign an empty Publisher to the Publisher
		//----------------------------------------------------------------------------
		parentObj.Publisher = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Publisher
		//----------------------------------------------------------------------------
		parentObj.PublisherId = nil;

		//----------------------------------------------------------------------------
		// save the RateCard
		//----------------------------------------------------------------------------
		return UpdateRateCard(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more ratesIds as a Rates to a RateCard
//----------------------------------------------------------------------------
func AddRatesToRateCard ( rateCardId uint64, ratesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RateCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRateCard(rateCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RateCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RateCard)

		// slice the ids on comma with no spaces
		ids := strings.Split( ratesIds, ",")

		for _, ratesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Rate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Rate
			// with a matching ratesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ratesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Rates using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Rates").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Rates", ratesId )
				return utils.RequestResult{false, msg, "unassignRates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RateCard from the gorm
		//----------------------------------------------------------------------------
		return GetRateCard(rateCardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ratesIds as a Rates from a RateCard
//----------------------------------------------------------------------------
func RemoveRatesFromRateCard( rateCardId uint64, ratesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RateCard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRateCard(rateCardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RateCard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RateCard)

		// slice the ids on comma with no spaces
		ids := strings.Split( ratesIds, ",")

		for _, ratesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Rate

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Rate
			// with a matching ratesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ratesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RateObj from the Rates array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Rates").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Rates", ratesId )
				return utils.RequestResult{false, msg, "removeRates", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RateCard from the gorm
		//----------------------------------------------------------------------------
		return GetRateCard(rateCardId)

	} else {
		return parentRequestResult
	}
}

