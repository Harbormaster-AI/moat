package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing UnderwriterDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateUnderwriter - creates a new db entry
//----------------------------------------------------------------------------
func CreateUnderwriter(obj model.Underwriter)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Underwriter with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Underwriter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateUnderwriter", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetUnderwriter - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetUnderwriter(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Underwriter

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Underwriter with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Underwriter using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Underwriter using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetUnderwriter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllUnderwriter - returns all
//----------------------------------------------------------------------------
func GetAllUnderwriter()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Underwriter

	//----------------------------------------------------------------------------
	// Request the ORM to find all Underwriter
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Underwriter" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Underwriter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllUnderwriter", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateUnderwriter - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateUnderwriter(obj model.Underwriter)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Underwriter using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Underwriter using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateUnderwriter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteUnderwriter - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteUnderwriter(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Underwriter with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetUnderwriter(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Underwriter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Underwriter)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Underwriter using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Underwriter using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteUnderwriter", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Insurer on a Underwriter
//----------------------------------------------------------------------------
func AssignInsurerToUnderwriter( underwriterId uint64, insurerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Underwriter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUnderwriter(underwriterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Underwriter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Underwriter)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Insurer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Insurer with a
		// matching insurerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, insurerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Insurer	to the Underwriter
			//----------------------------------------------------------------------------
			parentObj.Insurer = &childObj

			//----------------------------------------------------------------------------
			// save the Underwriter
			//----------------------------------------------------------------------------
			return UpdateUnderwriter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Insurer", insurerId )
			return utils.RequestResult{false, msg, "assignInsurer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Insurer on a Underwriter
//----------------------------------------------------------------------------
func UnassignInsurerFromUnderwriter(underwriterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Underwriter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUnderwriter(underwriterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Underwriter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Underwriter)

		//----------------------------------------------------------------------------
		// assign an empty Insurer to the Insurer
		//----------------------------------------------------------------------------
		parentObj.Insurer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Insurer
		//----------------------------------------------------------------------------
		parentObj.InsurerId = nil;

		//----------------------------------------------------------------------------
		// save the Underwriter
		//----------------------------------------------------------------------------
		return UpdateUnderwriter(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more decisionsIds as a Decisions to a Underwriter
//----------------------------------------------------------------------------
func AddDecisionsToUnderwriter ( underwriterId uint64, decisionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Underwriter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUnderwriter(underwriterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Underwriter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Underwriter)

		// slice the ids on comma with no spaces
		ids := strings.Split( decisionsIds, ",")

		for _, decisionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.UnderwritingDecision

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a UnderwritingDecision
			// with a matching decisionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , decisionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Decisions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Decisions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Decisions", decisionsId )
				return utils.RequestResult{false, msg, "unassignDecisions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Underwriter from the gorm
		//----------------------------------------------------------------------------
		return GetUnderwriter(underwriterId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more decisionsIds as a Decisions from a Underwriter
//----------------------------------------------------------------------------
func RemoveDecisionsFromUnderwriter( underwriterId uint64, decisionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Underwriter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetUnderwriter(underwriterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Underwriter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Underwriter)

		// slice the ids on comma with no spaces
		ids := strings.Split( decisionsIds, ",")

		for _, decisionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.UnderwritingDecision

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a UnderwritingDecision
			// with a matching decisionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , decisionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove UnderwritingDecisionObj from the Decisions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Decisions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Decisions", decisionsId )
				return utils.RequestResult{false, msg, "removeDecisions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Underwriter from the gorm
		//----------------------------------------------------------------------------
		return GetUnderwriter(underwriterId)

	} else {
		return parentRequestResult
	}
}

