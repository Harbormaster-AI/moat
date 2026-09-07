package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ComplianceProgramDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateComplianceProgram - creates a new db entry
//----------------------------------------------------------------------------
func CreateComplianceProgram(obj model.ComplianceProgram)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ComplianceProgram with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ComplianceProgram", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateComplianceProgram", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetComplianceProgram - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetComplianceProgram(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ComplianceProgram

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ComplianceProgram with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ComplianceProgram using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ComplianceProgram using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetComplianceProgram", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllComplianceProgram - returns all
//----------------------------------------------------------------------------
func GetAllComplianceProgram()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ComplianceProgram

	//----------------------------------------------------------------------------
	// Request the ORM to find all ComplianceProgram
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ComplianceProgram" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ComplianceProgram", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllComplianceProgram", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateComplianceProgram - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateComplianceProgram(obj model.ComplianceProgram)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ComplianceProgram using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ComplianceProgram using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateComplianceProgram", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteComplianceProgram - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteComplianceProgram(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetComplianceProgram(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ComplianceProgram)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ComplianceProgram using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ComplianceProgram using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteComplianceProgram", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a ComplianceProgram
//----------------------------------------------------------------------------
func AssignOrganizationToComplianceProgram( complianceProgramId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

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
			// assign the Organization	to the ComplianceProgram
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the ComplianceProgram
			//----------------------------------------------------------------------------
			return UpdateComplianceProgram(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a ComplianceProgram
//----------------------------------------------------------------------------
func UnassignOrganizationFromComplianceProgram(complianceProgramId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the ComplianceProgram
		//----------------------------------------------------------------------------
		return UpdateComplianceProgram(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more requirementsIds as a Requirements to a ComplianceProgram
//----------------------------------------------------------------------------
func AddRequirementsToComplianceProgram ( complianceProgramId uint64, requirementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( requirementsIds, ",")

		for _, requirementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceRequirement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceRequirement
			// with a matching requirementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , requirementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Requirements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Requirements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Requirements", requirementsId )
				return utils.RequestResult{false, msg, "unassignRequirements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ComplianceProgram from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceProgram(complianceProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more requirementsIds as a Requirements from a ComplianceProgram
//----------------------------------------------------------------------------
func RemoveRequirementsFromComplianceProgram( complianceProgramId uint64, requirementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( requirementsIds, ",")

		for _, requirementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceRequirement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceRequirement
			// with a matching requirementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , requirementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ComplianceRequirementObj from the Requirements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Requirements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Requirements", requirementsId )
				return utils.RequestResult{false, msg, "removeRequirements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ComplianceProgram from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceProgram(complianceProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more controlsIds as a Controls to a ComplianceProgram
//----------------------------------------------------------------------------
func AddControlsToComplianceProgram ( complianceProgramId uint64, controlsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

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
		// retrieve the modified ComplianceProgram from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceProgram(complianceProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more controlsIds as a Controls from a ComplianceProgram
//----------------------------------------------------------------------------
func RemoveControlsFromComplianceProgram( complianceProgramId uint64, controlsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

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
		// retrieve the modified ComplianceProgram from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceProgram(complianceProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more attestationsIds as a Attestations to a ComplianceProgram
//----------------------------------------------------------------------------
func AddAttestationsToComplianceProgram ( complianceProgramId uint64, attestationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( attestationsIds, ",")

		for _, attestationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Attestation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Attestation
			// with a matching attestationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , attestationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Attestations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Attestations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Attestations", attestationsId )
				return utils.RequestResult{false, msg, "unassignAttestations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ComplianceProgram from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceProgram(complianceProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more attestationsIds as a Attestations from a ComplianceProgram
//----------------------------------------------------------------------------
func RemoveAttestationsFromComplianceProgram( complianceProgramId uint64, attestationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( attestationsIds, ",")

		for _, attestationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Attestation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Attestation
			// with a matching attestationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , attestationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AttestationObj from the Attestations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Attestations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Attestations", attestationsId )
				return utils.RequestResult{false, msg, "removeAttestations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ComplianceProgram from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceProgram(complianceProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more regulationsIds as a Regulations to a ComplianceProgram
//----------------------------------------------------------------------------
func AddRegulationsToComplianceProgram ( complianceProgramId uint64, regulationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( regulationsIds, ",")

		for _, regulationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Regulation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Regulation
			// with a matching regulationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , regulationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Regulations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Regulations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Regulations", regulationsId )
				return utils.RequestResult{false, msg, "unassignRegulations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ComplianceProgram from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceProgram(complianceProgramId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more regulationsIds as a Regulations from a ComplianceProgram
//----------------------------------------------------------------------------
func RemoveRegulationsFromComplianceProgram( complianceProgramId uint64, regulationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ComplianceProgram with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceProgram(complianceProgramId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceProgram so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceProgram)

		// slice the ids on comma with no spaces
		ids := strings.Split( regulationsIds, ",")

		for _, regulationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Regulation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Regulation
			// with a matching regulationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , regulationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RegulationObj from the Regulations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Regulations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Regulations", regulationsId )
				return utils.RequestResult{false, msg, "removeRegulations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ComplianceProgram from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceProgram(complianceProgramId)

	} else {
		return parentRequestResult
	}
}

