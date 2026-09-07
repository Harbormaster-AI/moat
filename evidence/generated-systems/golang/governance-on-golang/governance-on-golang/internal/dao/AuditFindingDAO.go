package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AuditFindingDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAuditFinding - creates a new db entry
//----------------------------------------------------------------------------
func CreateAuditFinding(obj model.AuditFinding)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AuditFinding with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AuditFinding", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAuditFinding", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAuditFinding - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAuditFinding(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AuditFinding

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AuditFinding with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AuditFinding using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AuditFinding using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAuditFinding", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAuditFinding - returns all
//----------------------------------------------------------------------------
func GetAllAuditFinding()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AuditFinding

	//----------------------------------------------------------------------------
	// Request the ORM to find all AuditFinding
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AuditFinding" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AuditFinding", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAuditFinding", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAuditFinding - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAuditFinding(obj model.AuditFinding)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AuditFinding using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AuditFinding using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAuditFinding", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAuditFinding - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAuditFinding(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAuditFinding(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AuditFinding)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AuditFinding using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AuditFinding using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAuditFinding", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Engagement on a AuditFinding
//----------------------------------------------------------------------------
func AssignEngagementToAuditFinding( auditFindingId uint64, engagementId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AuditEngagement

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AuditEngagement with a
		// matching engagementId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, engagementId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Engagement	to the AuditFinding
			//----------------------------------------------------------------------------
			parentObj.Engagement = &childObj

			//----------------------------------------------------------------------------
			// save the AuditFinding
			//----------------------------------------------------------------------------
			return UpdateAuditFinding(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Engagement", engagementId )
			return utils.RequestResult{false, msg, "assignEngagement", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Engagement on a AuditFinding
//----------------------------------------------------------------------------
func UnassignEngagementFromAuditFinding(auditFindingId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		//----------------------------------------------------------------------------
		// assign an empty AuditEngagement to the Engagement
		//----------------------------------------------------------------------------
		parentObj.Engagement = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Engagement
		//----------------------------------------------------------------------------
		parentObj.EngagementId = nil;

		//----------------------------------------------------------------------------
		// save the AuditFinding
		//----------------------------------------------------------------------------
		return UpdateAuditFinding(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Workpaper on a AuditFinding
//----------------------------------------------------------------------------
func AssignWorkpaperToAuditFinding( auditFindingId uint64, workpaperId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AuditWorkpaper

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AuditWorkpaper with a
		// matching workpaperId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workpaperId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Workpaper	to the AuditFinding
			//----------------------------------------------------------------------------
			parentObj.Workpaper = &childObj

			//----------------------------------------------------------------------------
			// save the AuditFinding
			//----------------------------------------------------------------------------
			return UpdateAuditFinding(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workpaper", workpaperId )
			return utils.RequestResult{false, msg, "assignWorkpaper", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workpaper on a AuditFinding
//----------------------------------------------------------------------------
func UnassignWorkpaperFromAuditFinding(auditFindingId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		//----------------------------------------------------------------------------
		// assign an empty AuditWorkpaper to the Workpaper
		//----------------------------------------------------------------------------
		parentObj.Workpaper = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workpaper
		//----------------------------------------------------------------------------
		parentObj.WorkpaperId = nil;

		//----------------------------------------------------------------------------
		// save the AuditFinding
		//----------------------------------------------------------------------------
		return UpdateAuditFinding(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more correctiveActionsIds as a CorrectiveActions to a AuditFinding
//----------------------------------------------------------------------------
func AddCorrectiveActionsToAuditFinding ( auditFindingId uint64, correctiveActionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		// slice the ids on comma with no spaces
		ids := strings.Split( correctiveActionsIds, ",")

		for _, correctiveActionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CorrectiveAction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CorrectiveAction
			// with a matching correctiveActionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , correctiveActionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CorrectiveActions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CorrectiveActions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CorrectiveActions", correctiveActionsId )
				return utils.RequestResult{false, msg, "unassignCorrectiveActions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditFinding from the gorm
		//----------------------------------------------------------------------------
		return GetAuditFinding(auditFindingId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more correctiveActionsIds as a CorrectiveActions from a AuditFinding
//----------------------------------------------------------------------------
func RemoveCorrectiveActionsFromAuditFinding( auditFindingId uint64, correctiveActionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		// slice the ids on comma with no spaces
		ids := strings.Split( correctiveActionsIds, ",")

		for _, correctiveActionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CorrectiveAction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CorrectiveAction
			// with a matching correctiveActionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , correctiveActionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CorrectiveActionObj from the CorrectiveActions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CorrectiveActions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CorrectiveActions", correctiveActionsId )
				return utils.RequestResult{false, msg, "removeCorrectiveActions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditFinding from the gorm
		//----------------------------------------------------------------------------
		return GetAuditFinding(auditFindingId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more relatedRisksIds as a RelatedRisks to a AuditFinding
//----------------------------------------------------------------------------
func AddRelatedRisksToAuditFinding ( auditFindingId uint64, relatedRisksIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedRisksIds, ",")

		for _, relatedRisksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Risk

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Risk
			// with a matching relatedRisksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedRisksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RelatedRisks using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedRisks").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedRisks", relatedRisksId )
				return utils.RequestResult{false, msg, "unassignRelatedRisks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditFinding from the gorm
		//----------------------------------------------------------------------------
		return GetAuditFinding(auditFindingId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more relatedRisksIds as a RelatedRisks from a AuditFinding
//----------------------------------------------------------------------------
func RemoveRelatedRisksFromAuditFinding( auditFindingId uint64, relatedRisksIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedRisksIds, ",")

		for _, relatedRisksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Risk

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Risk
			// with a matching relatedRisksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedRisksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RiskObj from the RelatedRisks array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedRisks").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedRisks", relatedRisksId )
				return utils.RequestResult{false, msg, "removeRelatedRisks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditFinding from the gorm
		//----------------------------------------------------------------------------
		return GetAuditFinding(auditFindingId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more relatedControlsIds as a RelatedControls to a AuditFinding
//----------------------------------------------------------------------------
func AddRelatedControlsToAuditFinding ( auditFindingId uint64, relatedControlsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedControlsIds, ",")

		for _, relatedControlsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Control

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Control
			// with a matching relatedControlsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedControlsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RelatedControls using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedControls").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedControls", relatedControlsId )
				return utils.RequestResult{false, msg, "unassignRelatedControls", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditFinding from the gorm
		//----------------------------------------------------------------------------
		return GetAuditFinding(auditFindingId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more relatedControlsIds as a RelatedControls from a AuditFinding
//----------------------------------------------------------------------------
func RemoveRelatedControlsFromAuditFinding( auditFindingId uint64, relatedControlsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedControlsIds, ",")

		for _, relatedControlsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Control

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Control
			// with a matching relatedControlsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedControlsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ControlObj from the RelatedControls array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedControls").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedControls", relatedControlsId )
				return utils.RequestResult{false, msg, "removeRelatedControls", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditFinding from the gorm
		//----------------------------------------------------------------------------
		return GetAuditFinding(auditFindingId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more issuesIds as a Issues to a AuditFinding
//----------------------------------------------------------------------------
func AddIssuesToAuditFinding ( auditFindingId uint64, issuesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

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
		// retrieve the modified AuditFinding from the gorm
		//----------------------------------------------------------------------------
		return GetAuditFinding(auditFindingId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more issuesIds as a Issues from a AuditFinding
//----------------------------------------------------------------------------
func RemoveIssuesFromAuditFinding( auditFindingId uint64, issuesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditFinding with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditFinding(auditFindingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditFinding so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditFinding)

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
		// retrieve the modified AuditFinding from the gorm
		//----------------------------------------------------------------------------
		return GetAuditFinding(auditFindingId)

	} else {
		return parentRequestResult
	}
}

