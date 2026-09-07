package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
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
// assigns a Application on a Quote
//----------------------------------------------------------------------------
func AssignApplicationToQuote( quoteId uint64, applicationId uint64 )(utils.RequestResult){

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
		var childObj model.Application

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Application with a
		// matching applicationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, applicationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Application	to the Quote
			//----------------------------------------------------------------------------
			parentObj.Application = &childObj

			//----------------------------------------------------------------------------
			// save the Quote
			//----------------------------------------------------------------------------
			return UpdateQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Application", applicationId )
			return utils.RequestResult{false, msg, "assignApplication", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Application on a Quote
//----------------------------------------------------------------------------
func UnassignApplicationFromQuote(quoteId uint64)(utils.RequestResult) {

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
		// assign an empty Application to the Application
		//----------------------------------------------------------------------------
		parentObj.Application = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Application
		//----------------------------------------------------------------------------
		parentObj.ApplicationId = nil;

		//----------------------------------------------------------------------------
		// save the Quote
		//----------------------------------------------------------------------------
		return UpdateQuote(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Policy on a Quote
//----------------------------------------------------------------------------
func AssignPolicyToQuote( quoteId uint64, policyId uint64 )(utils.RequestResult){

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
		var childObj model.Policy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Policy with a
		// matching policyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, policyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Policy	to the Quote
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the Quote
			//----------------------------------------------------------------------------
			return UpdateQuote(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a Quote
//----------------------------------------------------------------------------
func UnassignPolicyFromQuote(quoteId uint64)(utils.RequestResult) {

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
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the Quote
		//----------------------------------------------------------------------------
		return UpdateQuote(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more underwritingDecisionsIds as a UnderwritingDecisions to a Quote
//----------------------------------------------------------------------------
func AddUnderwritingDecisionsToQuote ( quoteId uint64, underwritingDecisionsIds string )(utils.RequestResult) {

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

		// slice the ids on comma with no spaces
		ids := strings.Split( underwritingDecisionsIds, ",")

		for _, underwritingDecisionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.UnderwritingDecision

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a UnderwritingDecision
			// with a matching underwritingDecisionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , underwritingDecisionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the UnderwritingDecisions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("UnderwritingDecisions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "UnderwritingDecisions", underwritingDecisionsId )
				return utils.RequestResult{false, msg, "unassignUnderwritingDecisions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Quote from the gorm
		//----------------------------------------------------------------------------
		return GetQuote(quoteId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more underwritingDecisionsIds as a UnderwritingDecisions from a Quote
//----------------------------------------------------------------------------
func RemoveUnderwritingDecisionsFromQuote( quoteId uint64, underwritingDecisionsIds string )(utils.RequestResult) {
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

		// slice the ids on comma with no spaces
		ids := strings.Split( underwritingDecisionsIds, ",")

		for _, underwritingDecisionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.UnderwritingDecision

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a UnderwritingDecision
			// with a matching underwritingDecisionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , underwritingDecisionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove UnderwritingDecisionObj from the UnderwritingDecisions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("UnderwritingDecisions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "UnderwritingDecisions", underwritingDecisionsId )
				return utils.RequestResult{false, msg, "removeUnderwritingDecisions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Quote from the gorm
		//----------------------------------------------------------------------------
		return GetQuote(quoteId)

	} else {
		return parentRequestResult
	}
}

