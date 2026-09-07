package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FXQuoteDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFXQuote - creates a new db entry
//----------------------------------------------------------------------------
func CreateFXQuote(obj model.FXQuote)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FXQuote with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FXQuote", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFXQuote", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFXQuote - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFXQuote(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FXQuote

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FXQuote with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FXQuote using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FXQuote using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFXQuote", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFXQuote - returns all
//----------------------------------------------------------------------------
func GetAllFXQuote()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FXQuote

	//----------------------------------------------------------------------------
	// Request the ORM to find all FXQuote
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FXQuote" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FXQuote", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFXQuote", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFXQuote - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFXQuote(obj model.FXQuote)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FXQuote using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FXQuote using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFXQuote", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFXQuote - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFXQuote(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FXQuote with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFXQuote(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXQuote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FXQuote)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FXQuote using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FXQuote using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFXQuote", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a RequestedBy on a FXQuote
//----------------------------------------------------------------------------
func AssignRequestedByToFXQuote( fXQuoteId uint64, requestedById uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FXQuote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXQuote(fXQuoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXQuote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXQuote)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching requestedById
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, requestedById).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the RequestedBy	to the FXQuote
			//----------------------------------------------------------------------------
			parentObj.RequestedBy = &childObj

			//----------------------------------------------------------------------------
			// save the FXQuote
			//----------------------------------------------------------------------------
			return UpdateFXQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RequestedBy", requestedById )
			return utils.RequestResult{false, msg, "assignRequestedBy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a RequestedBy on a FXQuote
//----------------------------------------------------------------------------
func UnassignRequestedByFromFXQuote(fXQuoteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FXQuote with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFXQuote(fXQuoteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FXQuote so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FXQuote)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the RequestedBy
		//----------------------------------------------------------------------------
		parentObj.RequestedBy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the RequestedBy
		//----------------------------------------------------------------------------
		parentObj.RequestedById = nil;

		//----------------------------------------------------------------------------
		// save the FXQuote
		//----------------------------------------------------------------------------
		return UpdateFXQuote(parentObj)

	} else {
		return parentRequestResult
	}

}


