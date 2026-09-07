package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OrganizationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOrganization - creates a new db entry
//----------------------------------------------------------------------------
func CreateOrganization(obj model.Organization)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Organization with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Organization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOrganization", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOrganization - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOrganization(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Organization

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Organization with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Organization using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Organization using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOrganization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOrganization - returns all
//----------------------------------------------------------------------------
func GetAllOrganization()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Organization

	//----------------------------------------------------------------------------
	// Request the ORM to find all Organization
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Organization" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Organization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOrganization", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOrganization - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOrganization(obj model.Organization)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Organization using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Organization using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOrganization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOrganization - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOrganization(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOrganization(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Organization)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Organization using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Organization using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOrganization", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more governanceBodiesIds as a GovernanceBodies to a Organization
//----------------------------------------------------------------------------
func AddGovernanceBodiesToOrganization ( organizationId uint64, governanceBodiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( governanceBodiesIds, ",")

		for _, governanceBodiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GovernanceBody

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GovernanceBody
			// with a matching governanceBodiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , governanceBodiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the GovernanceBodies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("GovernanceBodies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GovernanceBodies", governanceBodiesId )
				return utils.RequestResult{false, msg, "unassignGovernanceBodies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more governanceBodiesIds as a GovernanceBodies from a Organization
//----------------------------------------------------------------------------
func RemoveGovernanceBodiesFromOrganization( organizationId uint64, governanceBodiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( governanceBodiesIds, ",")

		for _, governanceBodiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.GovernanceBody

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a GovernanceBody
			// with a matching governanceBodiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , governanceBodiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove GovernanceBodyObj from the GovernanceBodies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("GovernanceBodies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "GovernanceBodies", governanceBodiesId )
				return utils.RequestResult{false, msg, "removeGovernanceBodies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a Organization
//----------------------------------------------------------------------------
func AddPoliciesToOrganization ( organizationId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Policies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "unassignPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a Organization
//----------------------------------------------------------------------------
func RemovePoliciesFromOrganization( organizationId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Policy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Policy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyObj from the Policies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "removePolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more risksIds as a Risks to a Organization
//----------------------------------------------------------------------------
func AddRisksToOrganization ( organizationId uint64, risksIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

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
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more risksIds as a Risks from a Organization
//----------------------------------------------------------------------------
func RemoveRisksFromOrganization( organizationId uint64, risksIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

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
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more thirdPartiesIds as a ThirdParties to a Organization
//----------------------------------------------------------------------------
func AddThirdPartiesToOrganization ( organizationId uint64, thirdPartiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( thirdPartiesIds, ",")

		for _, thirdPartiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ThirdParty

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdParty
			// with a matching thirdPartiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , thirdPartiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ThirdParties using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ThirdParties").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdParties", thirdPartiesId )
				return utils.RequestResult{false, msg, "unassignThirdParties", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more thirdPartiesIds as a ThirdParties from a Organization
//----------------------------------------------------------------------------
func RemoveThirdPartiesFromOrganization( organizationId uint64, thirdPartiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( thirdPartiesIds, ",")

		for _, thirdPartiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ThirdParty

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ThirdParty
			// with a matching thirdPartiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , thirdPartiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ThirdPartyObj from the ThirdParties array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ThirdParties").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ThirdParties", thirdPartiesId )
				return utils.RequestResult{false, msg, "removeThirdParties", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more recordsRepositoriesIds as a RecordsRepositories to a Organization
//----------------------------------------------------------------------------
func AddRecordsRepositoriesToOrganization ( organizationId uint64, recordsRepositoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( recordsRepositoriesIds, ",")

		for _, recordsRepositoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RecordsRepository

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RecordsRepository
			// with a matching recordsRepositoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , recordsRepositoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RecordsRepositories using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RecordsRepositories").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RecordsRepositories", recordsRepositoriesId )
				return utils.RequestResult{false, msg, "unassignRecordsRepositories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more recordsRepositoriesIds as a RecordsRepositories from a Organization
//----------------------------------------------------------------------------
func RemoveRecordsRepositoriesFromOrganization( organizationId uint64, recordsRepositoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( recordsRepositoriesIds, ",")

		for _, recordsRepositoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RecordsRepository

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RecordsRepository
			// with a matching recordsRepositoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , recordsRepositoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RecordsRepositoryObj from the RecordsRepositories array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RecordsRepositories").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RecordsRepositories", recordsRepositoriesId )
				return utils.RequestResult{false, msg, "removeRecordsRepositories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataProcessingActivitiesIds as a DataProcessingActivities to a Organization
//----------------------------------------------------------------------------
func AddDataProcessingActivitiesToOrganization ( organizationId uint64, dataProcessingActivitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataProcessingActivitiesIds, ",")

		for _, dataProcessingActivitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataProcessingActivity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataProcessingActivity
			// with a matching dataProcessingActivitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataProcessingActivitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DataProcessingActivities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataProcessingActivities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataProcessingActivities", dataProcessingActivitiesId )
				return utils.RequestResult{false, msg, "unassignDataProcessingActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataProcessingActivitiesIds as a DataProcessingActivities from a Organization
//----------------------------------------------------------------------------
func RemoveDataProcessingActivitiesFromOrganization( organizationId uint64, dataProcessingActivitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataProcessingActivitiesIds, ",")

		for _, dataProcessingActivitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataProcessingActivity

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataProcessingActivity
			// with a matching dataProcessingActivitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataProcessingActivitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataProcessingActivityObj from the DataProcessingActivities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataProcessingActivities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataProcessingActivities", dataProcessingActivitiesId )
				return utils.RequestResult{false, msg, "removeDataProcessingActivities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more complianceProgramsIds as a CompliancePrograms to a Organization
//----------------------------------------------------------------------------
func AddComplianceProgramsToOrganization ( organizationId uint64, complianceProgramsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( complianceProgramsIds, ",")

		for _, complianceProgramsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceProgram

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceProgram
			// with a matching complianceProgramsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , complianceProgramsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CompliancePrograms using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompliancePrograms").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompliancePrograms", complianceProgramsId )
				return utils.RequestResult{false, msg, "unassignCompliancePrograms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more complianceProgramsIds as a CompliancePrograms from a Organization
//----------------------------------------------------------------------------
func RemoveComplianceProgramsFromOrganization( organizationId uint64, complianceProgramsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( complianceProgramsIds, ",")

		for _, complianceProgramsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceProgram

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceProgram
			// with a matching complianceProgramsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , complianceProgramsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ComplianceProgramObj from the CompliancePrograms array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CompliancePrograms").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CompliancePrograms", complianceProgramsId )
				return utils.RequestResult{false, msg, "removeCompliancePrograms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more auditProgramsIds as a AuditPrograms to a Organization
//----------------------------------------------------------------------------
func AddAuditProgramsToOrganization ( organizationId uint64, auditProgramsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( auditProgramsIds, ",")

		for _, auditProgramsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AuditProgram

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AuditProgram
			// with a matching auditProgramsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , auditProgramsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AuditPrograms using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AuditPrograms").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AuditPrograms", auditProgramsId )
				return utils.RequestResult{false, msg, "unassignAuditPrograms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more auditProgramsIds as a AuditPrograms from a Organization
//----------------------------------------------------------------------------
func RemoveAuditProgramsFromOrganization( organizationId uint64, auditProgramsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( auditProgramsIds, ",")

		for _, auditProgramsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AuditProgram

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AuditProgram
			// with a matching auditProgramsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , auditProgramsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AuditProgramObj from the AuditPrograms array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AuditPrograms").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AuditPrograms", auditProgramsId )
				return utils.RequestResult{false, msg, "removeAuditPrograms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more businessUnitsIds as a BusinessUnits to a Organization
//----------------------------------------------------------------------------
func AddBusinessUnitsToOrganization ( organizationId uint64, businessUnitsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

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
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more businessUnitsIds as a BusinessUnits from a Organization
//----------------------------------------------------------------------------
func RemoveBusinessUnitsFromOrganization( organizationId uint64, businessUnitsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

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
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more mattersIds as a Matters to a Organization
//----------------------------------------------------------------------------
func AddMattersToOrganization ( organizationId uint64, mattersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( mattersIds, ",")

		for _, mattersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Matter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Matter
			// with a matching mattersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , mattersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Matters using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Matters").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Matters", mattersId )
				return utils.RequestResult{false, msg, "unassignMatters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more mattersIds as a Matters from a Organization
//----------------------------------------------------------------------------
func RemoveMattersFromOrganization( organizationId uint64, mattersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( mattersIds, ",")

		for _, mattersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Matter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Matter
			// with a matching mattersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , mattersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MatterObj from the Matters array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Matters").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Matters", mattersId )
				return utils.RequestResult{false, msg, "removeMatters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataBreachesIds as a DataBreaches to a Organization
//----------------------------------------------------------------------------
func AddDataBreachesToOrganization ( organizationId uint64, dataBreachesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataBreachesIds, ",")

		for _, dataBreachesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataBreach

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataBreach
			// with a matching dataBreachesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataBreachesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DataBreaches using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataBreaches").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataBreaches", dataBreachesId )
				return utils.RequestResult{false, msg, "unassignDataBreaches", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataBreachesIds as a DataBreaches from a Organization
//----------------------------------------------------------------------------
func RemoveDataBreachesFromOrganization( organizationId uint64, dataBreachesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataBreachesIds, ",")

		for _, dataBreachesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataBreach

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataBreach
			// with a matching dataBreachesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataBreachesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataBreachObj from the DataBreaches array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataBreaches").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataBreaches", dataBreachesId )
				return utils.RequestResult{false, msg, "removeDataBreaches", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

