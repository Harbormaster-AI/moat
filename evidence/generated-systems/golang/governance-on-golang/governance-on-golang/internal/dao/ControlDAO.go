package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ControlDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateControl - creates a new db entry
//----------------------------------------------------------------------------
func CreateControl(obj model.Control)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Control with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Control", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateControl", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetControl - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetControl(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Control

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Control with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Control using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Control using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetControl", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllControl - returns all
//----------------------------------------------------------------------------
func GetAllControl()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Control

	//----------------------------------------------------------------------------
	// Request the ORM to find all Control
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Control" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Control", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllControl", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateControl - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateControl(obj model.Control)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Control using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Control using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateControl", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteControl - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteControl(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetControl(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Control)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Control using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Control using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteControl", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Policy on a Control
//----------------------------------------------------------------------------
func AssignPolicyToControl( controlId uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

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
			// assign the Policy	to the Control
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the Control
			//----------------------------------------------------------------------------
			return UpdateControl(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a Control
//----------------------------------------------------------------------------
func UnassignPolicyFromControl(controlId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the Control
		//----------------------------------------------------------------------------
		return UpdateControl(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more controlTestsIds as a ControlTests to a Control
//----------------------------------------------------------------------------
func AddControlTestsToControl ( controlId uint64, controlTestsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( controlTestsIds, ",")

		for _, controlTestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ControlTest_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ControlTest_
			// with a matching controlTestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , controlTestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ControlTests using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ControlTests").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ControlTests", controlTestsId )
				return utils.RequestResult{false, msg, "unassignControlTests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more controlTestsIds as a ControlTests from a Control
//----------------------------------------------------------------------------
func RemoveControlTestsFromControl( controlId uint64, controlTestsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( controlTestsIds, ",")

		for _, controlTestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ControlTest_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ControlTest_
			// with a matching controlTestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , controlTestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ControlTest_Obj from the ControlTests array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ControlTests").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ControlTests", controlTestsId )
				return utils.RequestResult{false, msg, "removeControlTests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more evidenceIds as a Evidence to a Control
//----------------------------------------------------------------------------
func AddEvidenceToControl ( controlId uint64, evidenceIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( evidenceIds, ",")

		for _, evidenceId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Evidence

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Evidence
			// with a matching evidenceId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , evidenceId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Evidence using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Evidence").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Evidence", evidenceId )
				return utils.RequestResult{false, msg, "unassignEvidence", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more evidenceIds as a Evidence from a Control
//----------------------------------------------------------------------------
func RemoveEvidenceFromControl( controlId uint64, evidenceIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( evidenceIds, ",")

		for _, evidenceId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Evidence

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Evidence
			// with a matching evidenceId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , evidenceId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EvidenceObj from the Evidence array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Evidence").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Evidence", evidenceId )
				return utils.RequestResult{false, msg, "removeEvidence", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more risksIds as a Risks to a Control
//----------------------------------------------------------------------------
func AddRisksToControl ( controlId uint64, risksIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( risksIds, ",")

		for _, risksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Risk

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Risk
			// with a matching risksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , risksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Risks using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Risks").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Risks", risksId )
				return utils.RequestResult{false, msg, "unassignRisks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more risksIds as a Risks from a Control
//----------------------------------------------------------------------------
func RemoveRisksFromControl( controlId uint64, risksIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( risksIds, ",")

		for _, risksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Risk

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Risk
			// with a matching risksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , risksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RiskObj from the Risks array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Risks").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Risks", risksId )
				return utils.RequestResult{false, msg, "removeRisks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more obligationsIds as a Obligations to a Control
//----------------------------------------------------------------------------
func AddObligationsToControl ( controlId uint64, obligationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( obligationsIds, ",")

		for _, obligationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Obligation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Obligation
			// with a matching obligationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , obligationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Obligations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Obligations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Obligations", obligationsId )
				return utils.RequestResult{false, msg, "unassignObligations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more obligationsIds as a Obligations from a Control
//----------------------------------------------------------------------------
func RemoveObligationsFromControl( controlId uint64, obligationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( obligationsIds, ",")

		for _, obligationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Obligation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Obligation
			// with a matching obligationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , obligationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ObligationObj from the Obligations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Obligations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Obligations", obligationsId )
				return utils.RequestResult{false, msg, "removeObligations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more proceduresIds as a Procedures to a Control
//----------------------------------------------------------------------------
func AddProceduresToControl ( controlId uint64, proceduresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( proceduresIds, ",")

		for _, proceduresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Procedure

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Procedure
			// with a matching proceduresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , proceduresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Procedures using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Procedures").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Procedures", proceduresId )
				return utils.RequestResult{false, msg, "unassignProcedures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more proceduresIds as a Procedures from a Control
//----------------------------------------------------------------------------
func RemoveProceduresFromControl( controlId uint64, proceduresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

		// slice the ids on comma with no spaces
		ids := strings.Split( proceduresIds, ",")

		for _, proceduresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Procedure

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Procedure
			// with a matching proceduresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , proceduresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ProcedureObj from the Procedures array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Procedures").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Procedures", proceduresId )
				return utils.RequestResult{false, msg, "removeProcedures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more issuesIds as a Issues to a Control
//----------------------------------------------------------------------------
func AddIssuesToControl ( controlId uint64, issuesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

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
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more issuesIds as a Issues from a Control
//----------------------------------------------------------------------------
func RemoveIssuesFromControl( controlId uint64, issuesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Control with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControl(controlId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Control so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Control)

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
		// retrieve the modified Control from the gorm
		//----------------------------------------------------------------------------
		return GetControl(controlId)

	} else {
		return parentRequestResult
	}
}

