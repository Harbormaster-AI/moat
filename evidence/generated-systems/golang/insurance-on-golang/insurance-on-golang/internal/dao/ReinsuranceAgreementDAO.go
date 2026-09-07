package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ReinsuranceAgreementDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateReinsuranceAgreement - creates a new db entry
//----------------------------------------------------------------------------
func CreateReinsuranceAgreement(obj model.ReinsuranceAgreement)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ReinsuranceAgreement with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ReinsuranceAgreement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateReinsuranceAgreement", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetReinsuranceAgreement - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetReinsuranceAgreement(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ReinsuranceAgreement

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ReinsuranceAgreement with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ReinsuranceAgreement using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ReinsuranceAgreement using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetReinsuranceAgreement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllReinsuranceAgreement - returns all
//----------------------------------------------------------------------------
func GetAllReinsuranceAgreement()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ReinsuranceAgreement

	//----------------------------------------------------------------------------
	// Request the ORM to find all ReinsuranceAgreement
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ReinsuranceAgreement" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ReinsuranceAgreement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllReinsuranceAgreement", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateReinsuranceAgreement - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateReinsuranceAgreement(obj model.ReinsuranceAgreement)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ReinsuranceAgreement using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ReinsuranceAgreement using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateReinsuranceAgreement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteReinsuranceAgreement - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteReinsuranceAgreement(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ReinsuranceAgreement with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetReinsuranceAgreement(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReinsuranceAgreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ReinsuranceAgreement)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ReinsuranceAgreement using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ReinsuranceAgreement using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteReinsuranceAgreement", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Insurer on a ReinsuranceAgreement
//----------------------------------------------------------------------------
func AssignInsurerToReinsuranceAgreement( reinsuranceAgreementId uint64, insurerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ReinsuranceAgreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReinsuranceAgreement(reinsuranceAgreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReinsuranceAgreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReinsuranceAgreement)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Insurer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Insurer with a
		// matching insurerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, insurerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Insurer	to the ReinsuranceAgreement
			//----------------------------------------------------------------------------
			parentObj.Insurer = &childObj

			//----------------------------------------------------------------------------
			// save the ReinsuranceAgreement
			//----------------------------------------------------------------------------
			return UpdateReinsuranceAgreement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Insurer", insurerId )
			return utils.RequestResult{false, msg, "assignInsurer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Insurer on a ReinsuranceAgreement
//----------------------------------------------------------------------------
func UnassignInsurerFromReinsuranceAgreement(reinsuranceAgreementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ReinsuranceAgreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReinsuranceAgreement(reinsuranceAgreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReinsuranceAgreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReinsuranceAgreement)

		//----------------------------------------------------------------------------
		// assign an empty Insurer to the Insurer
		//----------------------------------------------------------------------------
		parentObj.Insurer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Insurer
		//----------------------------------------------------------------------------
		parentObj.InsurerId = nil;

		//----------------------------------------------------------------------------
		// save the ReinsuranceAgreement
		//----------------------------------------------------------------------------
		return UpdateReinsuranceAgreement(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a ReinsuranceAgreement
//----------------------------------------------------------------------------
func AddPoliciesToReinsuranceAgreement ( reinsuranceAgreementId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ReinsuranceAgreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReinsuranceAgreement(reinsuranceAgreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReinsuranceAgreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReinsuranceAgreement)

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
		// retrieve the modified ReinsuranceAgreement from the gorm
		//----------------------------------------------------------------------------
		return GetReinsuranceAgreement(reinsuranceAgreementId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a ReinsuranceAgreement
//----------------------------------------------------------------------------
func RemovePoliciesFromReinsuranceAgreement( reinsuranceAgreementId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ReinsuranceAgreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReinsuranceAgreement(reinsuranceAgreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ReinsuranceAgreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ReinsuranceAgreement)

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
		// retrieve the modified ReinsuranceAgreement from the gorm
		//----------------------------------------------------------------------------
		return GetReinsuranceAgreement(reinsuranceAgreementId)

	} else {
		return parentRequestResult
	}
}

