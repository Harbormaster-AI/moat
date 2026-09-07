package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RegulationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRegulation - creates a new db entry
//----------------------------------------------------------------------------
func CreateRegulation(obj model.Regulation)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Regulation with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Regulation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRegulation", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRegulation - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRegulation(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Regulation

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Regulation with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Regulation using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Regulation using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRegulation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRegulation - returns all
//----------------------------------------------------------------------------
func GetAllRegulation()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Regulation

	//----------------------------------------------------------------------------
	// Request the ORM to find all Regulation
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Regulation" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Regulation", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRegulation", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRegulation - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRegulation(obj model.Regulation)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Regulation using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Regulation using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRegulation", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRegulation - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRegulation(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Regulation with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRegulation(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Regulation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Regulation)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Regulation using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Regulation using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRegulation", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more obligationsIds as a Obligations to a Regulation
//----------------------------------------------------------------------------
func AddObligationsToRegulation ( regulationId uint64, obligationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Regulation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRegulation(regulationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Regulation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Regulation)

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
		// retrieve the modified Regulation from the gorm
		//----------------------------------------------------------------------------
		return GetRegulation(regulationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more obligationsIds as a Obligations from a Regulation
//----------------------------------------------------------------------------
func RemoveObligationsFromRegulation( regulationId uint64, obligationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Regulation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRegulation(regulationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Regulation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Regulation)

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
		// retrieve the modified Regulation from the gorm
		//----------------------------------------------------------------------------
		return GetRegulation(regulationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more complianceProgramsIds as a CompliancePrograms to a Regulation
//----------------------------------------------------------------------------
func AddComplianceProgramsToRegulation ( regulationId uint64, complianceProgramsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Regulation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRegulation(regulationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Regulation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Regulation)

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
		// retrieve the modified Regulation from the gorm
		//----------------------------------------------------------------------------
		return GetRegulation(regulationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more complianceProgramsIds as a CompliancePrograms from a Regulation
//----------------------------------------------------------------------------
func RemoveComplianceProgramsFromRegulation( regulationId uint64, complianceProgramsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Regulation with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRegulation(regulationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Regulation so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Regulation)

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
		// retrieve the modified Regulation from the gorm
		//----------------------------------------------------------------------------
		return GetRegulation(regulationId)

	} else {
		return parentRequestResult
	}
}

