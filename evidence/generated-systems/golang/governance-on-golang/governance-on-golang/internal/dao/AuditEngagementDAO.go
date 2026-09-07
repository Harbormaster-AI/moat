package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AuditEngagementDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAuditEngagement - creates a new db entry
//----------------------------------------------------------------------------
func CreateAuditEngagement(obj model.AuditEngagement)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AuditEngagement with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AuditEngagement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAuditEngagement", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAuditEngagement - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAuditEngagement(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AuditEngagement

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AuditEngagement with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AuditEngagement using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AuditEngagement using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAuditEngagement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAuditEngagement - returns all
//----------------------------------------------------------------------------
func GetAllAuditEngagement()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AuditEngagement

	//----------------------------------------------------------------------------
	// Request the ORM to find all AuditEngagement
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AuditEngagement" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AuditEngagement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAuditEngagement", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAuditEngagement - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAuditEngagement(obj model.AuditEngagement)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AuditEngagement using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AuditEngagement using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAuditEngagement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAuditEngagement - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAuditEngagement(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAuditEngagement(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AuditEngagement)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AuditEngagement using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AuditEngagement using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAuditEngagement", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a AuditProgram on a AuditEngagement
//----------------------------------------------------------------------------
func AssignAuditProgramToAuditEngagement( auditEngagementId uint64, auditProgramId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AuditProgram

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AuditProgram with a
		// matching auditProgramId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, auditProgramId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AuditProgram	to the AuditEngagement
			//----------------------------------------------------------------------------
			parentObj.AuditProgram = &childObj

			//----------------------------------------------------------------------------
			// save the AuditEngagement
			//----------------------------------------------------------------------------
			return UpdateAuditEngagement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AuditProgram", auditProgramId )
			return utils.RequestResult{false, msg, "assignAuditProgram", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AuditProgram on a AuditEngagement
//----------------------------------------------------------------------------
func UnassignAuditProgramFromAuditEngagement(auditEngagementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

		//----------------------------------------------------------------------------
		// assign an empty AuditProgram to the AuditProgram
		//----------------------------------------------------------------------------
		parentObj.AuditProgram = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AuditProgram
		//----------------------------------------------------------------------------
		parentObj.AuditProgramId = nil;

		//----------------------------------------------------------------------------
		// save the AuditEngagement
		//----------------------------------------------------------------------------
		return UpdateAuditEngagement(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more businessUnitsIds as a BusinessUnits to a AuditEngagement
//----------------------------------------------------------------------------
func AddBusinessUnitsToAuditEngagement ( auditEngagementId uint64, businessUnitsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

		// slice the ids on comma with no spaces
		ids := strings.Split( businessUnitsIds, ",")

		for _, businessUnitsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BusinessUnit

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BusinessUnit
			// with a matching businessUnitsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , businessUnitsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the BusinessUnits using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BusinessUnits").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BusinessUnits", businessUnitsId )
				return utils.RequestResult{false, msg, "unassignBusinessUnits", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditEngagement from the gorm
		//----------------------------------------------------------------------------
		return GetAuditEngagement(auditEngagementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more businessUnitsIds as a BusinessUnits from a AuditEngagement
//----------------------------------------------------------------------------
func RemoveBusinessUnitsFromAuditEngagement( auditEngagementId uint64, businessUnitsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

		// slice the ids on comma with no spaces
		ids := strings.Split( businessUnitsIds, ",")

		for _, businessUnitsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BusinessUnit

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BusinessUnit
			// with a matching businessUnitsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , businessUnitsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BusinessUnitObj from the BusinessUnits array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BusinessUnits").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BusinessUnits", businessUnitsId )
				return utils.RequestResult{false, msg, "removeBusinessUnits", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditEngagement from the gorm
		//----------------------------------------------------------------------------
		return GetAuditEngagement(auditEngagementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more controlTestsIds as a ControlTests to a AuditEngagement
//----------------------------------------------------------------------------
func AddControlTestsToAuditEngagement ( auditEngagementId uint64, controlTestsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

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
		// retrieve the modified AuditEngagement from the gorm
		//----------------------------------------------------------------------------
		return GetAuditEngagement(auditEngagementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more controlTestsIds as a ControlTests from a AuditEngagement
//----------------------------------------------------------------------------
func RemoveControlTestsFromAuditEngagement( auditEngagementId uint64, controlTestsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

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
		// retrieve the modified AuditEngagement from the gorm
		//----------------------------------------------------------------------------
		return GetAuditEngagement(auditEngagementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more workpapersIds as a Workpapers to a AuditEngagement
//----------------------------------------------------------------------------
func AddWorkpapersToAuditEngagement ( auditEngagementId uint64, workpapersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

		// slice the ids on comma with no spaces
		ids := strings.Split( workpapersIds, ",")

		for _, workpapersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AuditWorkpaper

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AuditWorkpaper
			// with a matching workpapersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workpapersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Workpapers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Workpapers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workpapers", workpapersId )
				return utils.RequestResult{false, msg, "unassignWorkpapers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditEngagement from the gorm
		//----------------------------------------------------------------------------
		return GetAuditEngagement(auditEngagementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more workpapersIds as a Workpapers from a AuditEngagement
//----------------------------------------------------------------------------
func RemoveWorkpapersFromAuditEngagement( auditEngagementId uint64, workpapersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

		// slice the ids on comma with no spaces
		ids := strings.Split( workpapersIds, ",")

		for _, workpapersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AuditWorkpaper

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AuditWorkpaper
			// with a matching workpapersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workpapersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AuditWorkpaperObj from the Workpapers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Workpapers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workpapers", workpapersId )
				return utils.RequestResult{false, msg, "removeWorkpapers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditEngagement from the gorm
		//----------------------------------------------------------------------------
		return GetAuditEngagement(auditEngagementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more findingsIds as a Findings to a AuditEngagement
//----------------------------------------------------------------------------
func AddFindingsToAuditEngagement ( auditEngagementId uint64, findingsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

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
		// retrieve the modified AuditEngagement from the gorm
		//----------------------------------------------------------------------------
		return GetAuditEngagement(auditEngagementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more findingsIds as a Findings from a AuditEngagement
//----------------------------------------------------------------------------
func RemoveFindingsFromAuditEngagement( auditEngagementId uint64, findingsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditEngagement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditEngagement(auditEngagementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditEngagement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditEngagement)

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
		// retrieve the modified AuditEngagement from the gorm
		//----------------------------------------------------------------------------
		return GetAuditEngagement(auditEngagementId)

	} else {
		return parentRequestResult
	}
}

