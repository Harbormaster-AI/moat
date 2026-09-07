package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ComplianceRequirementDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateComplianceRequirement - creates a new db entry
//----------------------------------------------------------------------------
func CreateComplianceRequirement(obj model.ComplianceRequirement)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ComplianceRequirement with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ComplianceRequirement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateComplianceRequirement", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetComplianceRequirement - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetComplianceRequirement(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ComplianceRequirement

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ComplianceRequirement with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ComplianceRequirement using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ComplianceRequirement using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetComplianceRequirement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllComplianceRequirement - returns all
//----------------------------------------------------------------------------
func GetAllComplianceRequirement()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ComplianceRequirement

	//----------------------------------------------------------------------------
	// Request the ORM to find all ComplianceRequirement
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ComplianceRequirement" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ComplianceRequirement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllComplianceRequirement", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateComplianceRequirement - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateComplianceRequirement(obj model.ComplianceRequirement)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ComplianceRequirement using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ComplianceRequirement using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateComplianceRequirement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteComplianceRequirement - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteComplianceRequirement(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ComplianceRequirement with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetComplianceRequirement(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceRequirement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ComplianceRequirement)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ComplianceRequirement using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ComplianceRequirement using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteComplianceRequirement", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ComplianceProgram on a ComplianceRequirement
//----------------------------------------------------------------------------
func AssignComplianceProgramToComplianceRequirement( complianceRequirementId uint64, complianceProgramId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ComplianceRequirement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceRequirement(complianceRequirementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceRequirement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceRequirement)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ComplianceProgram

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ComplianceProgram with a
		// matching complianceProgramId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, complianceProgramId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ComplianceProgram	to the ComplianceRequirement
			//----------------------------------------------------------------------------
			parentObj.ComplianceProgram = &childObj

			//----------------------------------------------------------------------------
			// save the ComplianceRequirement
			//----------------------------------------------------------------------------
			return UpdateComplianceRequirement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ComplianceProgram", complianceProgramId )
			return utils.RequestResult{false, msg, "assignComplianceProgram", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ComplianceProgram on a ComplianceRequirement
//----------------------------------------------------------------------------
func UnassignComplianceProgramFromComplianceRequirement(complianceRequirementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceRequirement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceRequirement(complianceRequirementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceRequirement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceRequirement)

		//----------------------------------------------------------------------------
		// assign an empty ComplianceProgram to the ComplianceProgram
		//----------------------------------------------------------------------------
		parentObj.ComplianceProgram = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ComplianceProgram
		//----------------------------------------------------------------------------
		parentObj.ComplianceProgramId = nil;

		//----------------------------------------------------------------------------
		// save the ComplianceRequirement
		//----------------------------------------------------------------------------
		return UpdateComplianceRequirement(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a ComplianceRequirement
//----------------------------------------------------------------------------
func AddPoliciesToComplianceRequirement ( complianceRequirementId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceRequirement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceRequirement(complianceRequirementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceRequirement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceRequirement)

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
		// retrieve the modified ComplianceRequirement from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceRequirement(complianceRequirementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a ComplianceRequirement
//----------------------------------------------------------------------------
func RemovePoliciesFromComplianceRequirement( complianceRequirementId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ComplianceRequirement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceRequirement(complianceRequirementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceRequirement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceRequirement)

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
		// retrieve the modified ComplianceRequirement from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceRequirement(complianceRequirementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more controlsIds as a Controls to a ComplianceRequirement
//----------------------------------------------------------------------------
func AddControlsToComplianceRequirement ( complianceRequirementId uint64, controlsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceRequirement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceRequirement(complianceRequirementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceRequirement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceRequirement)

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
		// retrieve the modified ComplianceRequirement from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceRequirement(complianceRequirementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more controlsIds as a Controls from a ComplianceRequirement
//----------------------------------------------------------------------------
func RemoveControlsFromComplianceRequirement( complianceRequirementId uint64, controlsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ComplianceRequirement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceRequirement(complianceRequirementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceRequirement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceRequirement)

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
		// retrieve the modified ComplianceRequirement from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceRequirement(complianceRequirementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more obligationsIds as a Obligations to a ComplianceRequirement
//----------------------------------------------------------------------------
func AddObligationsToComplianceRequirement ( complianceRequirementId uint64, obligationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ComplianceRequirement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceRequirement(complianceRequirementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceRequirement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceRequirement)

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
		// retrieve the modified ComplianceRequirement from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceRequirement(complianceRequirementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more obligationsIds as a Obligations from a ComplianceRequirement
//----------------------------------------------------------------------------
func RemoveObligationsFromComplianceRequirement( complianceRequirementId uint64, obligationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ComplianceRequirement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetComplianceRequirement(complianceRequirementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ComplianceRequirement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ComplianceRequirement)

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
		// retrieve the modified ComplianceRequirement from the gorm
		//----------------------------------------------------------------------------
		return GetComplianceRequirement(complianceRequirementId)

	} else {
		return parentRequestResult
	}
}

