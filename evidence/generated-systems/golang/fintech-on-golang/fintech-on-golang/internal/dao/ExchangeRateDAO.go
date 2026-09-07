package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ExchangeRateDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateExchangeRate - creates a new db entry
//----------------------------------------------------------------------------
func CreateExchangeRate(obj model.ExchangeRate)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ExchangeRate with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ExchangeRate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateExchangeRate", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetExchangeRate - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetExchangeRate(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ExchangeRate

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ExchangeRate with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ExchangeRate using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ExchangeRate using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetExchangeRate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllExchangeRate - returns all
//----------------------------------------------------------------------------
func GetAllExchangeRate()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ExchangeRate

	//----------------------------------------------------------------------------
	// Request the ORM to find all ExchangeRate
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ExchangeRate" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ExchangeRate", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllExchangeRate", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateExchangeRate - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateExchangeRate(obj model.ExchangeRate)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ExchangeRate using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ExchangeRate using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateExchangeRate", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteExchangeRate - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteExchangeRate(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetExchangeRate(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ExchangeRate)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ExchangeRate using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ExchangeRate using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteExchangeRate", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more usedByQuotesIds as a UsedByQuotes to a ExchangeRate
//----------------------------------------------------------------------------
func AddUsedByQuotesToExchangeRate ( exchangeRateId uint64, usedByQuotesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExchangeRate)

		// slice the ids on comma with no spaces
		ids := strings.Split( usedByQuotesIds, ",")

		for _, usedByQuotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FXQuote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FXQuote
			// with a matching usedByQuotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , usedByQuotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the UsedByQuotes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("UsedByQuotes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "UsedByQuotes", usedByQuotesId )
				return utils.RequestResult{false, msg, "unassignUsedByQuotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ExchangeRate from the gorm
		//----------------------------------------------------------------------------
		return GetExchangeRate(exchangeRateId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more usedByQuotesIds as a UsedByQuotes from a ExchangeRate
//----------------------------------------------------------------------------
func RemoveUsedByQuotesFromExchangeRate( exchangeRateId uint64, usedByQuotesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ExchangeRate with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExchangeRate(exchangeRateId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ExchangeRate so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ExchangeRate)

		// slice the ids on comma with no spaces
		ids := strings.Split( usedByQuotesIds, ",")

		for _, usedByQuotesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FXQuote

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FXQuote
			// with a matching usedByQuotesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , usedByQuotesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FXQuoteObj from the UsedByQuotes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("UsedByQuotes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "UsedByQuotes", usedByQuotesId )
				return utils.RequestResult{false, msg, "removeUsedByQuotes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ExchangeRate from the gorm
		//----------------------------------------------------------------------------
		return GetExchangeRate(exchangeRateId)

	} else {
		return parentRequestResult
	}
}

