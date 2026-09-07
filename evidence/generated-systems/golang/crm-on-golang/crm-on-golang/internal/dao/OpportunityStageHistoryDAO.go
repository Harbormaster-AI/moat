package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OpportunityStageHistoryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOpportunityStageHistory - creates a new db entry
//----------------------------------------------------------------------------
func CreateOpportunityStageHistory(obj model.OpportunityStageHistory)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a OpportunityStageHistory with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a OpportunityStageHistory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOpportunityStageHistory", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOpportunityStageHistory - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOpportunityStageHistory(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.OpportunityStageHistory

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a OpportunityStageHistory with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a OpportunityStageHistory using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a OpportunityStageHistory using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOpportunityStageHistory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOpportunityStageHistory - returns all
//----------------------------------------------------------------------------
func GetAllOpportunityStageHistory()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.OpportunityStageHistory

	//----------------------------------------------------------------------------
	// Request the ORM to find all OpportunityStageHistory
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all OpportunityStageHistory" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all OpportunityStageHistory", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOpportunityStageHistory", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOpportunityStageHistory - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOpportunityStageHistory(obj model.OpportunityStageHistory)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a OpportunityStageHistory using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a OpportunityStageHistory using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOpportunityStageHistory", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOpportunityStageHistory - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOpportunityStageHistory(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the OpportunityStageHistory with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOpportunityStageHistory(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityStageHistory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.OpportunityStageHistory)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a OpportunityStageHistory using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a OpportunityStageHistory using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOpportunityStageHistory", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Opportunity on a OpportunityStageHistory
//----------------------------------------------------------------------------
func AssignOpportunityToOpportunityStageHistory( opportunityStageHistoryId uint64, opportunityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OpportunityStageHistory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityStageHistory(opportunityStageHistoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityStageHistory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityStageHistory)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Opportunity

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Opportunity with a
		// matching opportunityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, opportunityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Opportunity	to the OpportunityStageHistory
			//----------------------------------------------------------------------------
			parentObj.Opportunity = &childObj

			//----------------------------------------------------------------------------
			// save the OpportunityStageHistory
			//----------------------------------------------------------------------------
			return UpdateOpportunityStageHistory(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Opportunity", opportunityId )
			return utils.RequestResult{false, msg, "assignOpportunity", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Opportunity on a OpportunityStageHistory
//----------------------------------------------------------------------------
func UnassignOpportunityFromOpportunityStageHistory(opportunityStageHistoryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OpportunityStageHistory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityStageHistory(opportunityStageHistoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityStageHistory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityStageHistory)

		//----------------------------------------------------------------------------
		// assign an empty Opportunity to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.Opportunity = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Opportunity
		//----------------------------------------------------------------------------
		parentObj.OpportunityId = nil;

		//----------------------------------------------------------------------------
		// save the OpportunityStageHistory
		//----------------------------------------------------------------------------
		return UpdateOpportunityStageHistory(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ChangedBy on a OpportunityStageHistory
//----------------------------------------------------------------------------
func AssignChangedByToOpportunityStageHistory( opportunityStageHistoryId uint64, changedById uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OpportunityStageHistory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityStageHistory(opportunityStageHistoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityStageHistory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityStageHistory)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.User

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a User with a
		// matching changedById
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, changedById).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ChangedBy	to the OpportunityStageHistory
			//----------------------------------------------------------------------------
			parentObj.ChangedBy = &childObj

			//----------------------------------------------------------------------------
			// save the OpportunityStageHistory
			//----------------------------------------------------------------------------
			return UpdateOpportunityStageHistory(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ChangedBy", changedById )
			return utils.RequestResult{false, msg, "assignChangedBy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ChangedBy on a OpportunityStageHistory
//----------------------------------------------------------------------------
func UnassignChangedByFromOpportunityStageHistory(opportunityStageHistoryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OpportunityStageHistory with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOpportunityStageHistory(opportunityStageHistoryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OpportunityStageHistory so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OpportunityStageHistory)

		//----------------------------------------------------------------------------
		// assign an empty User to the ChangedBy
		//----------------------------------------------------------------------------
		parentObj.ChangedBy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ChangedBy
		//----------------------------------------------------------------------------
		parentObj.ChangedById = nil;

		//----------------------------------------------------------------------------
		// save the OpportunityStageHistory
		//----------------------------------------------------------------------------
		return UpdateOpportunityStageHistory(parentObj)

	} else {
		return parentRequestResult
	}

}


