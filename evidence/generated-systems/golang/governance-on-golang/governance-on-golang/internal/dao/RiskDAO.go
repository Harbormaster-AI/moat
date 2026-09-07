package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RiskDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRisk - creates a new db entry
//----------------------------------------------------------------------------
func CreateRisk(obj model.Risk)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Risk with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Risk", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRisk", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRisk - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRisk(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Risk

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Risk with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Risk using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Risk using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRisk", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRisk - returns all
//----------------------------------------------------------------------------
func GetAllRisk()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Risk

	//----------------------------------------------------------------------------
	// Request the ORM to find all Risk
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Risk" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Risk", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRisk", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRisk - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRisk(obj model.Risk)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Risk using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Risk using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRisk", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRisk - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRisk(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRisk(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Risk)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Risk using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Risk using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRisk", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Risk
//----------------------------------------------------------------------------
func AssignOrganizationToRisk( riskId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the Risk
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Risk
			//----------------------------------------------------------------------------
			return UpdateRisk(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Risk
//----------------------------------------------------------------------------
func UnassignOrganizationFromRisk(riskId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Risk
		//----------------------------------------------------------------------------
		return UpdateRisk(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more controlsIds as a Controls to a Risk
//----------------------------------------------------------------------------
func AddControlsToRisk ( riskId uint64, controlsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

		// slice the ids on comma with no spaces
		ids := strings.Split( controlsIds, ",")

		for _, controlsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Control

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Control
			// with a matching controlsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , controlsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Controls using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Controls").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Controls", controlsId )
				return utils.RequestResult{false, msg, "unassignControls", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Risk from the gorm
		//----------------------------------------------------------------------------
		return GetRisk(riskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more controlsIds as a Controls from a Risk
//----------------------------------------------------------------------------
func RemoveControlsFromRisk( riskId uint64, controlsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

		// slice the ids on comma with no spaces
		ids := strings.Split( controlsIds, ",")

		for _, controlsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Control

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Control
			// with a matching controlsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , controlsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ControlObj from the Controls array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Controls").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Controls", controlsId )
				return utils.RequestResult{false, msg, "removeControls", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Risk from the gorm
		//----------------------------------------------------------------------------
		return GetRisk(riskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more assessmentsIds as a Assessments to a Risk
//----------------------------------------------------------------------------
func AddAssessmentsToRisk ( riskId uint64, assessmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

		// slice the ids on comma with no spaces
		ids := strings.Split( assessmentsIds, ",")

		for _, assessmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RiskAssessment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RiskAssessment
			// with a matching assessmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assessmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Assessments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assessments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assessments", assessmentsId )
				return utils.RequestResult{false, msg, "unassignAssessments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Risk from the gorm
		//----------------------------------------------------------------------------
		return GetRisk(riskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more assessmentsIds as a Assessments from a Risk
//----------------------------------------------------------------------------
func RemoveAssessmentsFromRisk( riskId uint64, assessmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

		// slice the ids on comma with no spaces
		ids := strings.Split( assessmentsIds, ",")

		for _, assessmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RiskAssessment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RiskAssessment
			// with a matching assessmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , assessmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RiskAssessmentObj from the Assessments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Assessments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assessments", assessmentsId )
				return utils.RequestResult{false, msg, "removeAssessments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Risk from the gorm
		//----------------------------------------------------------------------------
		return GetRisk(riskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more issuesIds as a Issues to a Risk
//----------------------------------------------------------------------------
func AddIssuesToRisk ( riskId uint64, issuesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

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
		// retrieve the modified Risk from the gorm
		//----------------------------------------------------------------------------
		return GetRisk(riskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more issuesIds as a Issues from a Risk
//----------------------------------------------------------------------------
func RemoveIssuesFromRisk( riskId uint64, issuesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

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
		// retrieve the modified Risk from the gorm
		//----------------------------------------------------------------------------
		return GetRisk(riskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more findingsIds as a Findings to a Risk
//----------------------------------------------------------------------------
func AddFindingsToRisk ( riskId uint64, findingsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

		// slice the ids on comma with no spaces
		ids := strings.Split( findingsIds, ",")

		for _, findingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AuditFinding

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AuditFinding
			// with a matching findingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , findingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Findings using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Findings").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Findings", findingsId )
				return utils.RequestResult{false, msg, "unassignFindings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Risk from the gorm
		//----------------------------------------------------------------------------
		return GetRisk(riskId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more findingsIds as a Findings from a Risk
//----------------------------------------------------------------------------
func RemoveFindingsFromRisk( riskId uint64, findingsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Risk with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRisk(riskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Risk so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Risk)

		// slice the ids on comma with no spaces
		ids := strings.Split( findingsIds, ",")

		for _, findingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AuditFinding

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AuditFinding
			// with a matching findingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , findingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AuditFindingObj from the Findings array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Findings").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Findings", findingsId )
				return utils.RequestResult{false, msg, "removeFindings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Risk from the gorm
		//----------------------------------------------------------------------------
		return GetRisk(riskId)

	} else {
		return parentRequestResult
	}
}

