package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EvidenceDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEvidence - creates a new db entry
//----------------------------------------------------------------------------
func CreateEvidence(obj model.Evidence)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Evidence with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Evidence", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEvidence", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEvidence - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEvidence(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Evidence

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Evidence with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Evidence using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Evidence using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEvidence", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEvidence - returns all
//----------------------------------------------------------------------------
func GetAllEvidence()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Evidence

	//----------------------------------------------------------------------------
	// Request the ORM to find all Evidence
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Evidence" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Evidence", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEvidence", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEvidence - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEvidence(obj model.Evidence)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Evidence using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Evidence using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEvidence", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEvidence - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEvidence(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Evidence with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEvidence(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Evidence so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Evidence)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Evidence using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Evidence using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEvidence", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ControlTest on a Evidence
//----------------------------------------------------------------------------
func AssignControlTestToEvidence( evidenceId uint64, controlTestId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Evidence with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvidence(evidenceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Evidence so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Evidence)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ControlTest_

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ControlTest_ with a
		// matching controlTestId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, controlTestId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ControlTest	to the Evidence
			//----------------------------------------------------------------------------
			parentObj.ControlTest = &childObj

			//----------------------------------------------------------------------------
			// save the Evidence
			//----------------------------------------------------------------------------
			return UpdateEvidence(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ControlTest", controlTestId )
			return utils.RequestResult{false, msg, "assignControlTest", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ControlTest on a Evidence
//----------------------------------------------------------------------------
func UnassignControlTestFromEvidence(evidenceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Evidence with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvidence(evidenceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Evidence so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Evidence)

		//----------------------------------------------------------------------------
		// assign an empty ControlTest_ to the ControlTest
		//----------------------------------------------------------------------------
		parentObj.ControlTest = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ControlTest
		//----------------------------------------------------------------------------
		parentObj.ControlTestId = nil;

		//----------------------------------------------------------------------------
		// save the Evidence
		//----------------------------------------------------------------------------
		return UpdateEvidence(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Control on a Evidence
//----------------------------------------------------------------------------
func AssignControlToEvidence( evidenceId uint64, controlId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Evidence with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvidence(evidenceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Evidence so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Evidence)

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
			// assign the Control	to the Evidence
			//----------------------------------------------------------------------------
			parentObj.Control = &childObj

			//----------------------------------------------------------------------------
			// save the Evidence
			//----------------------------------------------------------------------------
			return UpdateEvidence(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Control", controlId )
			return utils.RequestResult{false, msg, "assignControl", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Control on a Evidence
//----------------------------------------------------------------------------
func UnassignControlFromEvidence(evidenceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Evidence with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvidence(evidenceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Evidence so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Evidence)

		//----------------------------------------------------------------------------
		// assign an empty Control to the Control
		//----------------------------------------------------------------------------
		parentObj.Control = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Control
		//----------------------------------------------------------------------------
		parentObj.ControlId = nil;

		//----------------------------------------------------------------------------
		// save the Evidence
		//----------------------------------------------------------------------------
		return UpdateEvidence(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Obligation on a Evidence
//----------------------------------------------------------------------------
func AssignObligationToEvidence( evidenceId uint64, obligationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Evidence with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvidence(evidenceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Evidence so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Evidence)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Obligation

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Obligation with a
		// matching obligationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, obligationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Obligation	to the Evidence
			//----------------------------------------------------------------------------
			parentObj.Obligation = &childObj

			//----------------------------------------------------------------------------
			// save the Evidence
			//----------------------------------------------------------------------------
			return UpdateEvidence(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Obligation", obligationId )
			return utils.RequestResult{false, msg, "assignObligation", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Obligation on a Evidence
//----------------------------------------------------------------------------
func UnassignObligationFromEvidence(evidenceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Evidence with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvidence(evidenceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Evidence so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Evidence)

		//----------------------------------------------------------------------------
		// assign an empty Obligation to the Obligation
		//----------------------------------------------------------------------------
		parentObj.Obligation = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Obligation
		//----------------------------------------------------------------------------
		parentObj.ObligationId = nil;

		//----------------------------------------------------------------------------
		// save the Evidence
		//----------------------------------------------------------------------------
		return UpdateEvidence(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Workpaper on a Evidence
//----------------------------------------------------------------------------
func AssignWorkpaperToEvidence( evidenceId uint64, workpaperId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Evidence with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvidence(evidenceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Evidence so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Evidence)

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
			// assign the Workpaper	to the Evidence
			//----------------------------------------------------------------------------
			parentObj.Workpaper = &childObj

			//----------------------------------------------------------------------------
			// save the Evidence
			//----------------------------------------------------------------------------
			return UpdateEvidence(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workpaper", workpaperId )
			return utils.RequestResult{false, msg, "assignWorkpaper", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workpaper on a Evidence
//----------------------------------------------------------------------------
func UnassignWorkpaperFromEvidence(evidenceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Evidence with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvidence(evidenceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Evidence so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Evidence)

		//----------------------------------------------------------------------------
		// assign an empty AuditWorkpaper to the Workpaper
		//----------------------------------------------------------------------------
		parentObj.Workpaper = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workpaper
		//----------------------------------------------------------------------------
		parentObj.WorkpaperId = nil;

		//----------------------------------------------------------------------------
		// save the Evidence
		//----------------------------------------------------------------------------
		return UpdateEvidence(parentObj)

	} else {
		return parentRequestResult
	}

}


