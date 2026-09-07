package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PolicyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePolicy - creates a new db entry
//----------------------------------------------------------------------------
func CreatePolicy(obj model.Policy)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Policy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Policy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePolicy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPolicy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPolicy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Policy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Policy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Policy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Policy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPolicy - returns all
//----------------------------------------------------------------------------
func GetAllPolicy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Policy

	//----------------------------------------------------------------------------
	// Request the ORM to find all Policy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Policy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Policy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPolicy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePolicy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePolicy(obj model.Policy)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Policy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Policy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePolicy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePolicy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPolicy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Policy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Policy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePolicy", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a Policy
//----------------------------------------------------------------------------
func AssignOrganizationToPolicy( policyId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

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
			// assign the Organization	to the Policy
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the Policy
			//----------------------------------------------------------------------------
			return UpdatePolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a Policy
//----------------------------------------------------------------------------
func UnassignOrganizationFromPolicy(policyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the Policy
		//----------------------------------------------------------------------------
		return UpdatePolicy(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more ownersIds as a Owners to a Policy
//----------------------------------------------------------------------------
func AddOwnersToPolicy ( policyId uint64, ownersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( ownersIds, ",")

		for _, ownersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Person

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Person
			// with a matching ownersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Owners using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Owners").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owners", ownersId )
				return utils.RequestResult{false, msg, "unassignOwners", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ownersIds as a Owners from a Policy
//----------------------------------------------------------------------------
func RemoveOwnersFromPolicy( policyId uint64, ownersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( ownersIds, ",")

		for _, ownersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Person

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Person
			// with a matching ownersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ownersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PersonObj from the Owners array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Owners").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owners", ownersId )
				return utils.RequestResult{false, msg, "removeOwners", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more relatedRequirementsIds as a RelatedRequirements to a Policy
//----------------------------------------------------------------------------
func AddRelatedRequirementsToPolicy ( policyId uint64, relatedRequirementsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedRequirementsIds, ",")

		for _, relatedRequirementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceRequirement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceRequirement
			// with a matching relatedRequirementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedRequirementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RelatedRequirements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedRequirements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedRequirements", relatedRequirementsId )
				return utils.RequestResult{false, msg, "unassignRelatedRequirements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more relatedRequirementsIds as a RelatedRequirements from a Policy
//----------------------------------------------------------------------------
func RemoveRelatedRequirementsFromPolicy( policyId uint64, relatedRequirementsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( relatedRequirementsIds, ",")

		for _, relatedRequirementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ComplianceRequirement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ComplianceRequirement
			// with a matching relatedRequirementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , relatedRequirementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ComplianceRequirementObj from the RelatedRequirements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RelatedRequirements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RelatedRequirements", relatedRequirementsId )
				return utils.RequestResult{false, msg, "removeRelatedRequirements", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more controlsIds as a Controls to a Policy
//----------------------------------------------------------------------------
func AddControlsToPolicy ( policyId uint64, controlsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

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
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more controlsIds as a Controls from a Policy
//----------------------------------------------------------------------------
func RemoveControlsFromPolicy( policyId uint64, controlsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

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
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more proceduresIds as a Procedures to a Policy
//----------------------------------------------------------------------------
func AddProceduresToPolicy ( policyId uint64, proceduresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

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
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more proceduresIds as a Procedures from a Policy
//----------------------------------------------------------------------------
func RemoveProceduresFromPolicy( policyId uint64, proceduresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

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
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more exceptionsIds as a Exceptions to a Policy
//----------------------------------------------------------------------------
func AddExceptionsToPolicy ( policyId uint64, exceptionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( exceptionsIds, ",")

		for _, exceptionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Exception_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Exception_
			// with a matching exceptionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exceptionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Exceptions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Exceptions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exceptions", exceptionsId )
				return utils.RequestResult{false, msg, "unassignExceptions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more exceptionsIds as a Exceptions from a Policy
//----------------------------------------------------------------------------
func RemoveExceptionsFromPolicy( policyId uint64, exceptionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

		// slice the ids on comma with no spaces
		ids := strings.Split( exceptionsIds, ",")

		for _, exceptionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Exception_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Exception_
			// with a matching exceptionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exceptionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Exception_Obj from the Exceptions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Exceptions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exceptions", exceptionsId )
				return utils.RequestResult{false, msg, "removeExceptions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more attestationsIds as a Attestations to a Policy
//----------------------------------------------------------------------------
func AddAttestationsToPolicy ( policyId uint64, attestationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

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
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more attestationsIds as a Attestations from a Policy
//----------------------------------------------------------------------------
func RemoveAttestationsFromPolicy( policyId uint64, attestationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Policy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicy(policyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Policy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Policy)

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
		// retrieve the modified Policy from the gorm
		//----------------------------------------------------------------------------
		return GetPolicy(policyId)

	} else {
		return parentRequestResult
	}
}

