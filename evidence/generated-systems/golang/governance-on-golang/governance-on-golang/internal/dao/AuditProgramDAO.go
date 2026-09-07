package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AuditProgramDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAuditProgram - creates a new db entry
//----------------------------------------------------------------------------
func CreateAuditProgram(obj model.AuditProgram)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AuditProgram with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AuditProgram", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAuditProgram", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAuditProgram - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAuditProgram(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AuditProgram

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AuditProgram with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AuditProgram using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AuditProgram using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAuditProgram", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAuditProgram - returns all
//----------------------------------------------------------------------------
func GetAllAuditProgram()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AuditProgram

	//----------------------------------------------------------------------------
	// Request the ORM to find all AuditProgram
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AuditProgram" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AuditProgram", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAuditProgram", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAuditProgram - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAuditProgram(obj model.AuditProgram)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AuditProgram using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AuditProgram using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAuditProgram", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAuditProgram - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAuditProgram(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AuditProgram with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAuditProgram(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AuditProgram)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AuditProgram using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AuditProgram using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAuditProgram", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a AuditProgram
//----------------------------------------------------------------------------
func AssignOrganizationToAuditProgram( auditProgramId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AuditProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditProgram(auditProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditProgram)

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
			// assign the Organization	to the AuditProgram
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the AuditProgram
			//----------------------------------------------------------------------------
			return UpdateAuditProgram(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a AuditProgram
//----------------------------------------------------------------------------
func UnassignOrganizationFromAuditProgram(auditProgramId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditProgram(auditProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditProgram)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the AuditProgram
		//----------------------------------------------------------------------------
		return UpdateAuditProgram(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more engagementsIds as a Engagements to a AuditProgram
//----------------------------------------------------------------------------
func AddEngagementsToAuditProgram ( auditProgramId uint64, engagementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AuditProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditProgram(auditProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( engagementsIds, ",")

		for _, engagementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AuditEngagement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AuditEngagement
			// with a matching engagementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , engagementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Engagements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Engagements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Engagements", engagementsId )
				return utils.RequestResult{false, msg, "unassignEngagements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditProgram from the gorm
		//----------------------------------------------------------------------------
		return GetAuditProgram(auditProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more engagementsIds as a Engagements from a AuditProgram
//----------------------------------------------------------------------------
func RemoveEngagementsFromAuditProgram( auditProgramId uint64, engagementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AuditProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAuditProgram(auditProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AuditProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AuditProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( engagementsIds, ",")

		for _, engagementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AuditEngagement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AuditEngagement
			// with a matching engagementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , engagementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AuditEngagementObj from the Engagements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Engagements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Engagements", engagementsId )
				return utils.RequestResult{false, msg, "removeEngagements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AuditProgram from the gorm
		//----------------------------------------------------------------------------
		return GetAuditProgram(auditProgramId)

	} else {
		return parentRequestResult
	}
}

