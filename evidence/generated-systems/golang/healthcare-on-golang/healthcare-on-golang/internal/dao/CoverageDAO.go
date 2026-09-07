package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CoverageDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCoverage - creates a new db entry
//----------------------------------------------------------------------------
func CreateCoverage(obj model.Coverage)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Coverage with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Coverage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCoverage", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCoverage - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCoverage(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Coverage

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Coverage with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Coverage using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Coverage using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCoverage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCoverage - returns all
//----------------------------------------------------------------------------
func GetAllCoverage()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Coverage

	//----------------------------------------------------------------------------
	// Request the ORM to find all Coverage
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Coverage" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Coverage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCoverage", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCoverage - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCoverage(obj model.Coverage)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Coverage using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Coverage using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCoverage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCoverage - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCoverage(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Coverage with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCoverage(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Coverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Coverage)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Coverage using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Coverage using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCoverage", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Patient on a Coverage
//----------------------------------------------------------------------------
func AssignPatientToCoverage( coverageId uint64, patientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Coverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverage(coverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Coverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Coverage)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Patient

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Patient with a
		// matching patientId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, patientId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Patient	to the Coverage
			//----------------------------------------------------------------------------
			parentObj.Patient = &childObj

			//----------------------------------------------------------------------------
			// save the Coverage
			//----------------------------------------------------------------------------
			return UpdateCoverage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Patient", patientId )
			return utils.RequestResult{false, msg, "assignPatient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Patient on a Coverage
//----------------------------------------------------------------------------
func UnassignPatientFromCoverage(coverageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Coverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverage(coverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Coverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Coverage)

		//----------------------------------------------------------------------------
		// assign an empty Patient to the Patient
		//----------------------------------------------------------------------------
		parentObj.Patient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Patient
		//----------------------------------------------------------------------------
		parentObj.PatientId = nil;

		//----------------------------------------------------------------------------
		// save the Coverage
		//----------------------------------------------------------------------------
		return UpdateCoverage(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Plan on a Coverage
//----------------------------------------------------------------------------
func AssignPlanToCoverage( coverageId uint64, planId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Coverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverage(coverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Coverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Coverage)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InsurancePlan

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InsurancePlan with a
		// matching planId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, planId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Plan	to the Coverage
			//----------------------------------------------------------------------------
			parentObj.Plan = &childObj

			//----------------------------------------------------------------------------
			// save the Coverage
			//----------------------------------------------------------------------------
			return UpdateCoverage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plan", planId )
			return utils.RequestResult{false, msg, "assignPlan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plan on a Coverage
//----------------------------------------------------------------------------
func UnassignPlanFromCoverage(coverageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Coverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverage(coverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Coverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Coverage)

		//----------------------------------------------------------------------------
		// assign an empty InsurancePlan to the Plan
		//----------------------------------------------------------------------------
		parentObj.Plan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plan
		//----------------------------------------------------------------------------
		parentObj.PlanId = nil;

		//----------------------------------------------------------------------------
		// save the Coverage
		//----------------------------------------------------------------------------
		return UpdateCoverage(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more claimsIds as a Claims to a Coverage
//----------------------------------------------------------------------------
func AddClaimsToCoverage ( coverageId uint64, claimsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Coverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverage(coverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Coverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Coverage)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Claims using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "unassignClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Coverage from the gorm
		//----------------------------------------------------------------------------
		return GetCoverage(coverageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more claimsIds as a Claims from a Coverage
//----------------------------------------------------------------------------
func RemoveClaimsFromCoverage( coverageId uint64, claimsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Coverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverage(coverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Coverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Coverage)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimObj from the Claims array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "removeClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Coverage from the gorm
		//----------------------------------------------------------------------------
		return GetCoverage(coverageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more authorizationsIds as a Authorizations to a Coverage
//----------------------------------------------------------------------------
func AddAuthorizationsToCoverage ( coverageId uint64, authorizationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Coverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverage(coverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Coverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Coverage)

		// slice the ids on comma with no spaces
		ids := strings.Split( authorizationsIds, ",")

		for _, authorizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Authorization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Authorization
			// with a matching authorizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , authorizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Authorizations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Authorizations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Authorizations", authorizationsId )
				return utils.RequestResult{false, msg, "unassignAuthorizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Coverage from the gorm
		//----------------------------------------------------------------------------
		return GetCoverage(coverageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more authorizationsIds as a Authorizations from a Coverage
//----------------------------------------------------------------------------
func RemoveAuthorizationsFromCoverage( coverageId uint64, authorizationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Coverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCoverage(coverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Coverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Coverage)

		// slice the ids on comma with no spaces
		ids := strings.Split( authorizationsIds, ",")

		for _, authorizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Authorization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Authorization
			// with a matching authorizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , authorizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AuthorizationObj from the Authorizations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Authorizations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Authorizations", authorizationsId )
				return utils.RequestResult{false, msg, "removeAuthorizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Coverage from the gorm
		//----------------------------------------------------------------------------
		return GetCoverage(coverageId)

	} else {
		return parentRequestResult
	}
}

