package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing QuoteDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateQuote - creates a new db entry
//----------------------------------------------------------------------------
func CreateQuote(obj model.Quote)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Quote with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Quote", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateQuote", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetQuote - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetQuote(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Quote

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Quote with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Quote using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Quote using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetQuote", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllQuote - returns all
//----------------------------------------------------------------------------
func GetAllQuote()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Quote

	//----------------------------------------------------------------------------
	// Request the ORM to find all Quote
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Quote" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Quote", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllQuote", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateQuote - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateQuote(obj model.Quote)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Quote using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Quote using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateQuote", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteQuote - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteQuote(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetQuote(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Quote using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Quote using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteQuote", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a AircraftOrder on a Quote
//----------------------------------------------------------------------------
func AssignAircraftOrderToQuote( quoteId uint64, aircraftOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftOrder with a
		// matching aircraftOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, aircraftOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AircraftOrder	to the Quote
			//----------------------------------------------------------------------------
			parentObj.AircraftOrder = &childObj

			//----------------------------------------------------------------------------
			// save the Quote
			//----------------------------------------------------------------------------
			return UpdateQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AircraftOrder", aircraftOrderId )
			return utils.RequestResult{false, msg, "assignAircraftOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AircraftOrder on a Quote
//----------------------------------------------------------------------------
func UnassignAircraftOrderFromQuote(quoteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Quote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetQuote(quoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Quote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Quote)

		//----------------------------------------------------------------------------
		// assign an empty AircraftOrder to the AircraftOrder
		//----------------------------------------------------------------------------
		parentObj.AircraftOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AircraftOrder
		//----------------------------------------------------------------------------
		parentObj.AircraftOrderId = nil;

		//----------------------------------------------------------------------------
		// save the Quote
		//----------------------------------------------------------------------------
		return UpdateQuote(parentObj)

	} else {
		return parentRequestResult
	}

}


