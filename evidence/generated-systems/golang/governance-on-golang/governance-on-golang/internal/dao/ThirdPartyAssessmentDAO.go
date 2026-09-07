package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ThirdPartyAssessmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateThirdPartyAssessment - creates a new db entry
//----------------------------------------------------------------------------
func CreateThirdPartyAssessment(obj model.ThirdPartyAssessment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ThirdPartyAssessment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ThirdPartyAssessment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateThirdPartyAssessment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetThirdPartyAssessment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetThirdPartyAssessment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ThirdPartyAssessment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ThirdPartyAssessment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ThirdPartyAssessment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ThirdPartyAssessment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetThirdPartyAssessment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllThirdPartyAssessment - returns all
//----------------------------------------------------------------------------
func GetAllThirdPartyAssessment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ThirdPartyAssessment

	//----------------------------------------------------------------------------
	// Request the ORM to find all ThirdPartyAssessment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ThirdPartyAssessment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ThirdPartyAssessment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllThirdPartyAssessment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateThirdPartyAssessment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateThirdPartyAssessment(obj model.ThirdPartyAssessment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ThirdPartyAssessment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ThirdPartyAssessment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateThirdPartyAssessment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteThirdPartyAssessment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteThirdPartyAssessment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyAssessment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetThirdPartyAssessment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyAssessment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ThirdPartyAssessment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ThirdPartyAssessment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ThirdPartyAssessment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteThirdPartyAssessment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ThirdParty on a ThirdPartyAssessment
//----------------------------------------------------------------------------
func AssignThirdPartyToThirdPartyAssessment( thirdPartyAssessmentId uint64, thirdPartyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyAssessment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyAssessment(thirdPartyAssessmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyAssessment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdPartyAssessment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ThirdParty

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ThirdParty with a
		// matching thirdPartyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, thirdPartyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ThirdParty	to the ThirdPartyAssessment
			//----------------------------------------------------------------------------
			parentObj.ThirdParty = &childObj

			//----------------------------------------------------------------------------
			// save the ThirdPartyAssessment
			//----------------------------------------------------------------------------
			return UpdateThirdPartyAssessment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdParty", thirdPartyId )
			return utils.RequestResult{false, msg, "assignThirdParty", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ThirdParty on a ThirdPartyAssessment
//----------------------------------------------------------------------------
func UnassignThirdPartyFromThirdPartyAssessment(thirdPartyAssessmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyAssessment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyAssessment(thirdPartyAssessmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyAssessment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdPartyAssessment)

		//----------------------------------------------------------------------------
		// assign an empty ThirdParty to the ThirdParty
		//----------------------------------------------------------------------------
		parentObj.ThirdParty = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ThirdParty
		//----------------------------------------------------------------------------
		parentObj.ThirdPartyId = nil;

		//----------------------------------------------------------------------------
		// save the ThirdPartyAssessment
		//----------------------------------------------------------------------------
		return UpdateThirdPartyAssessment(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more issuesIds as a Issues to a ThirdPartyAssessment
//----------------------------------------------------------------------------
func AddIssuesToThirdPartyAssessment ( thirdPartyAssessmentId uint64, issuesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyAssessment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyAssessment(thirdPartyAssessmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyAssessment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdPartyAssessment)

		// slice the ids on comma with no spaces
		ids := strings.Split( issuesIds, ",")

		for _, issuesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Issue

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Issue
			// with a matching issuesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , issuesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Issues using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Issues").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Issues", issuesId )
				return utils.RequestResult{false, msg, "unassignIssues", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ThirdPartyAssessment from the gorm
		//----------------------------------------------------------------------------
		return GetThirdPartyAssessment(thirdPartyAssessmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more issuesIds as a Issues from a ThirdPartyAssessment
//----------------------------------------------------------------------------
func RemoveIssuesFromThirdPartyAssessment( thirdPartyAssessmentId uint64, issuesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ThirdPartyAssessment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetThirdPartyAssessment(thirdPartyAssessmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ThirdPartyAssessment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ThirdPartyAssessment)

		// slice the ids on comma with no spaces
		ids := strings.Split( issuesIds, ",")

		for _, issuesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Issue

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Issue
			// with a matching issuesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , issuesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove IssueObj from the Issues array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Issues").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Issues", issuesId )
				return utils.RequestResult{false, msg, "removeIssues", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ThirdPartyAssessment from the gorm
		//----------------------------------------------------------------------------
		return GetThirdPartyAssessment(thirdPartyAssessmentId)

	} else {
		return parentRequestResult
	}
}

