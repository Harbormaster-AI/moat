package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AuditWorkpaperDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAuditWorkpaper - creates a new db entry
//----------------------------------------------------------------------------
func CreateAuditWorkpaper(obj model.AuditWorkpaper)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AuditWorkpaper with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AuditWorkpaper", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAuditWorkpaper", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAuditWorkpaper - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAuditWorkpaper(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AuditWorkpaper

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AuditWorkpaper with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AuditWorkpaper using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AuditWorkpaper using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAuditWorkpaper", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAuditWorkpaper - returns all
//----------------------------------------------------------------------------
func GetAllAuditWorkpaper()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AuditWorkpaper

	//----------------------------------------------------------------------------
	// Request the ORM to find all AuditWorkpaper
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AuditWorkpaper" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AuditWorkpaper", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAuditWorkpaper", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAuditWorkpaper - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAuditWorkpaper(obj model.AuditWorkpaper)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AuditWorkpaper using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AuditWorkpaper using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAuditWorkpaper", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAuditWorkpaper - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAuditWorkpaper(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AuditWorkpaper with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAuditWorkpaper(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditWorkpaper so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AuditWorkpaper)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AuditWorkpaper using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AuditWorkpaper using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAuditWorkpaper", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Engagement on a AuditWorkpaper
//----------------------------------------------------------------------------
func AssignEngagementToAuditWorkpaper( auditWorkpaperId uint64, engagementId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AuditWorkpaper with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditWorkpaper(auditWorkpaperId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditWorkpaper so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditWorkpaper)

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
			// assign the Engagement	to the AuditWorkpaper
			//----------------------------------------------------------------------------
			parentObj.Engagement = &childObj

			//----------------------------------------------------------------------------
			// save the AuditWorkpaper
			//----------------------------------------------------------------------------
			return UpdateAuditWorkpaper(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Engagement", engagementId )
			return utils.RequestResult{false, msg, "assignEngagement", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Engagement on a AuditWorkpaper
//----------------------------------------------------------------------------
func UnassignEngagementFromAuditWorkpaper(auditWorkpaperId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditWorkpaper with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditWorkpaper(auditWorkpaperId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditWorkpaper so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditWorkpaper)

		//----------------------------------------------------------------------------
		// assign an empty AuditEngagement to the Engagement
		//----------------------------------------------------------------------------
		parentObj.Engagement = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Engagement
		//----------------------------------------------------------------------------
		parentObj.EngagementId = nil;

		//----------------------------------------------------------------------------
		// save the AuditWorkpaper
		//----------------------------------------------------------------------------
		return UpdateAuditWorkpaper(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more evidenceIds as a Evidence to a AuditWorkpaper
//----------------------------------------------------------------------------
func AddEvidenceToAuditWorkpaper ( auditWorkpaperId uint64, evidenceIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditWorkpaper with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditWorkpaper(auditWorkpaperId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditWorkpaper so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditWorkpaper)

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
		// retrieve the modified AuditWorkpaper from the gorm
		//----------------------------------------------------------------------------
		return GetAuditWorkpaper(auditWorkpaperId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more evidenceIds as a Evidence from a AuditWorkpaper
//----------------------------------------------------------------------------
func RemoveEvidenceFromAuditWorkpaper( auditWorkpaperId uint64, evidenceIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditWorkpaper with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditWorkpaper(auditWorkpaperId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditWorkpaper so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditWorkpaper)

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
		// retrieve the modified AuditWorkpaper from the gorm
		//----------------------------------------------------------------------------
		return GetAuditWorkpaper(auditWorkpaperId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more findingsIds as a Findings to a AuditWorkpaper
//----------------------------------------------------------------------------
func AddFindingsToAuditWorkpaper ( auditWorkpaperId uint64, findingsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditWorkpaper with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditWorkpaper(auditWorkpaperId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditWorkpaper so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditWorkpaper)

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
		// retrieve the modified AuditWorkpaper from the gorm
		//----------------------------------------------------------------------------
		return GetAuditWorkpaper(auditWorkpaperId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more findingsIds as a Findings from a AuditWorkpaper
//----------------------------------------------------------------------------
func RemoveFindingsFromAuditWorkpaper( auditWorkpaperId uint64, findingsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditWorkpaper with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditWorkpaper(auditWorkpaperId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditWorkpaper so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditWorkpaper)

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
		// retrieve the modified AuditWorkpaper from the gorm
		//----------------------------------------------------------------------------
		return GetAuditWorkpaper(auditWorkpaperId)

	} else {
		return parentRequestResult
	}
}

