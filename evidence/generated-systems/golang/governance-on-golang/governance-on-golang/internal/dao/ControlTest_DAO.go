package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ControlTest_DAO..." ) )
}

//----------------------------------------------------------------------------
// CreateControlTest_ - creates a new db entry
//----------------------------------------------------------------------------
func CreateControlTest_(obj model.ControlTest_)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ControlTest_ with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ControlTest_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateControlTest_", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetControlTest_ - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetControlTest_(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ControlTest_

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ControlTest_ with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ControlTest_ using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ControlTest_ using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetControlTest_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllControlTest_ - returns all
//----------------------------------------------------------------------------
func GetAllControlTest_()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ControlTest_

	//----------------------------------------------------------------------------
	// Request the ORM to find all ControlTest_
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ControlTest_" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ControlTest_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllControlTest_", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateControlTest_ - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateControlTest_(obj model.ControlTest_)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ControlTest_ using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ControlTest_ using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateControlTest_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteControlTest_ - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteControlTest_(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ControlTest_ with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetControlTest_(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ControlTest_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ControlTest_)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ControlTest_ using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ControlTest_ using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteControlTest_", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Control on a ControlTest_
//----------------------------------------------------------------------------
func AssignControlToControlTest_( controlTest_Id uint64, controlId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ControlTest_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControlTest_(controlTest_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ControlTest_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ControlTest_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Control

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Control with a
		// matching controlId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, controlId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Control	to the ControlTest_
			//----------------------------------------------------------------------------
			parentObj.Control = &childObj

			//----------------------------------------------------------------------------
			// save the ControlTest_
			//----------------------------------------------------------------------------
			return UpdateControlTest_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Control", controlId )
			return utils.RequestResult{false, msg, "assignControl", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Control on a ControlTest_
//----------------------------------------------------------------------------
func UnassignControlFromControlTest_(controlTest_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ControlTest_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControlTest_(controlTest_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ControlTest_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ControlTest_)

		//----------------------------------------------------------------------------
		// assign an empty Control to the Control
		//----------------------------------------------------------------------------
		parentObj.Control = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Control
		//----------------------------------------------------------------------------
		parentObj.ControlId = nil;

		//----------------------------------------------------------------------------
		// save the ControlTest_
		//----------------------------------------------------------------------------
		return UpdateControlTest_(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Engagement on a ControlTest_
//----------------------------------------------------------------------------
func AssignEngagementToControlTest_( controlTest_Id uint64, engagementId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ControlTest_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControlTest_(controlTest_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ControlTest_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ControlTest_)

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
			// assign the Engagement	to the ControlTest_
			//----------------------------------------------------------------------------
			parentObj.Engagement = &childObj

			//----------------------------------------------------------------------------
			// save the ControlTest_
			//----------------------------------------------------------------------------
			return UpdateControlTest_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Engagement", engagementId )
			return utils.RequestResult{false, msg, "assignEngagement", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Engagement on a ControlTest_
//----------------------------------------------------------------------------
func UnassignEngagementFromControlTest_(controlTest_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ControlTest_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControlTest_(controlTest_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ControlTest_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ControlTest_)

		//----------------------------------------------------------------------------
		// assign an empty AuditEngagement to the Engagement
		//----------------------------------------------------------------------------
		parentObj.Engagement = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Engagement
		//----------------------------------------------------------------------------
		parentObj.EngagementId = nil;

		//----------------------------------------------------------------------------
		// save the ControlTest_
		//----------------------------------------------------------------------------
		return UpdateControlTest_(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more evidenceIds as a Evidence to a ControlTest_
//----------------------------------------------------------------------------
func AddEvidenceToControlTest_ ( controlTest_Id uint64, evidenceIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ControlTest_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControlTest_(controlTest_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ControlTest_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ControlTest_)

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
		// retrieve the modified ControlTest_ from the gorm
		//----------------------------------------------------------------------------
		return GetControlTest_(controlTest_Id)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more evidenceIds as a Evidence from a ControlTest_
//----------------------------------------------------------------------------
func RemoveEvidenceFromControlTest_( controlTest_Id uint64, evidenceIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ControlTest_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetControlTest_(controlTest_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ControlTest_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ControlTest_)

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
		// retrieve the modified ControlTest_ from the gorm
		//----------------------------------------------------------------------------
		return GetControlTest_(controlTest_Id)

	} else {
		return parentRequestResult
	}
}

