package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing UnderwritingDecisionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateUnderwritingDecision - creates a new db entry
//----------------------------------------------------------------------------
func CreateUnderwritingDecision(obj model.UnderwritingDecision)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a UnderwritingDecision with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a UnderwritingDecision", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateUnderwritingDecision", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetUnderwritingDecision - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetUnderwritingDecision(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.UnderwritingDecision

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a UnderwritingDecision with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a UnderwritingDecision using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a UnderwritingDecision using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetUnderwritingDecision", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllUnderwritingDecision - returns all
//----------------------------------------------------------------------------
func GetAllUnderwritingDecision()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.UnderwritingDecision

	//----------------------------------------------------------------------------
	// Request the ORM to find all UnderwritingDecision
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all UnderwritingDecision" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all UnderwritingDecision", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllUnderwritingDecision", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateUnderwritingDecision - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateUnderwritingDecision(obj model.UnderwritingDecision)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a UnderwritingDecision using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a UnderwritingDecision using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateUnderwritingDecision", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteUnderwritingDecision - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteUnderwritingDecision(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the UnderwritingDecision with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetUnderwritingDecision(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UnderwritingDecision so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.UnderwritingDecision)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a UnderwritingDecision using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a UnderwritingDecision using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteUnderwritingDecision", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Quote on a UnderwritingDecision
//----------------------------------------------------------------------------
func AssignQuoteToUnderwritingDecision( underwritingDecisionId uint64, quoteId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the UnderwritingDecision with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUnderwritingDecision(underwritingDecisionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UnderwritingDecision so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.UnderwritingDecision)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Quote

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Quote with a
		// matching quoteId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, quoteId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Quote	to the UnderwritingDecision
			//----------------------------------------------------------------------------
			parentObj.Quote = &childObj

			//----------------------------------------------------------------------------
			// save the UnderwritingDecision
			//----------------------------------------------------------------------------
			return UpdateUnderwritingDecision(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Quote", quoteId )
			return utils.RequestResult{false, msg, "assignQuote", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Quote on a UnderwritingDecision
//----------------------------------------------------------------------------
func UnassignQuoteFromUnderwritingDecision(underwritingDecisionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the UnderwritingDecision with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUnderwritingDecision(underwritingDecisionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UnderwritingDecision so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.UnderwritingDecision)

		//----------------------------------------------------------------------------
		// assign an empty Quote to the Quote
		//----------------------------------------------------------------------------
		parentObj.Quote = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Quote
		//----------------------------------------------------------------------------
		parentObj.QuoteId = nil;

		//----------------------------------------------------------------------------
		// save the UnderwritingDecision
		//----------------------------------------------------------------------------
		return UpdateUnderwritingDecision(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Underwriter on a UnderwritingDecision
//----------------------------------------------------------------------------
func AssignUnderwriterToUnderwritingDecision( underwritingDecisionId uint64, underwriterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the UnderwritingDecision with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUnderwritingDecision(underwritingDecisionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UnderwritingDecision so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.UnderwritingDecision)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Underwriter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Underwriter with a
		// matching underwriterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, underwriterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Underwriter	to the UnderwritingDecision
			//----------------------------------------------------------------------------
			parentObj.Underwriter = &childObj

			//----------------------------------------------------------------------------
			// save the UnderwritingDecision
			//----------------------------------------------------------------------------
			return UpdateUnderwritingDecision(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Underwriter", underwriterId )
			return utils.RequestResult{false, msg, "assignUnderwriter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Underwriter on a UnderwritingDecision
//----------------------------------------------------------------------------
func UnassignUnderwriterFromUnderwritingDecision(underwritingDecisionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the UnderwritingDecision with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUnderwritingDecision(underwritingDecisionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.UnderwritingDecision so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.UnderwritingDecision)

		//----------------------------------------------------------------------------
		// assign an empty Underwriter to the Underwriter
		//----------------------------------------------------------------------------
		parentObj.Underwriter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Underwriter
		//----------------------------------------------------------------------------
		parentObj.UnderwriterId = nil;

		//----------------------------------------------------------------------------
		// save the UnderwritingDecision
		//----------------------------------------------------------------------------
		return UpdateUnderwritingDecision(parentObj)

	} else {
		return parentRequestResult
	}

}


