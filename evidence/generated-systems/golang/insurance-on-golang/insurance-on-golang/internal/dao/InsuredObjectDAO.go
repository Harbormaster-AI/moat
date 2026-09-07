package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InsuredObjectDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInsuredObject - creates a new db entry
//----------------------------------------------------------------------------
func CreateInsuredObject(obj model.InsuredObject)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InsuredObject with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InsuredObject", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInsuredObject", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInsuredObject - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInsuredObject(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InsuredObject

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InsuredObject with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InsuredObject using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InsuredObject using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInsuredObject", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInsuredObject - returns all
//----------------------------------------------------------------------------
func GetAllInsuredObject()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InsuredObject

	//----------------------------------------------------------------------------
	// Request the ORM to find all InsuredObject
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InsuredObject" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InsuredObject", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInsuredObject", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInsuredObject - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInsuredObject(obj model.InsuredObject)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InsuredObject using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InsuredObject using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInsuredObject", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInsuredObject - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInsuredObject(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InsuredObject with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInsuredObject(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuredObject so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InsuredObject)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InsuredObject using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InsuredObject using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInsuredObject", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Policy on a InsuredObject
//----------------------------------------------------------------------------
func AssignPolicyToInsuredObject( insuredObjectId uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InsuredObject with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsuredObject(insuredObjectId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuredObject so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsuredObject)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Policy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Policy with a
		// matching policyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, policyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Policy	to the InsuredObject
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the InsuredObject
			//----------------------------------------------------------------------------
			return UpdateInsuredObject(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a InsuredObject
//----------------------------------------------------------------------------
func UnassignPolicyFromInsuredObject(insuredObjectId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsuredObject with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsuredObject(insuredObjectId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuredObject so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsuredObject)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the InsuredObject
		//----------------------------------------------------------------------------
		return UpdateInsuredObject(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more coveragesIds as a Coverages to a InsuredObject
//----------------------------------------------------------------------------
func AddCoveragesToInsuredObject ( insuredObjectId uint64, coveragesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsuredObject with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsuredObject(insuredObjectId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuredObject so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsuredObject)

		// slice the ids on comma with no spaces
		ids := strings.Split( coveragesIds, ",")

		for _, coveragesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PolicyCoverage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PolicyCoverage
			// with a matching coveragesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coveragesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Coverages using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Coverages").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverages", coveragesId )
				return utils.RequestResult{false, msg, "unassignCoverages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsuredObject from the gorm
		//----------------------------------------------------------------------------
		return GetInsuredObject(insuredObjectId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more coveragesIds as a Coverages from a InsuredObject
//----------------------------------------------------------------------------
func RemoveCoveragesFromInsuredObject( insuredObjectId uint64, coveragesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InsuredObject with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsuredObject(insuredObjectId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsuredObject so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsuredObject)

		// slice the ids on comma with no spaces
		ids := strings.Split( coveragesIds, ",")

		for _, coveragesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PolicyCoverage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PolicyCoverage
			// with a matching coveragesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coveragesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyCoverageObj from the Coverages array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Coverages").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverages", coveragesId )
				return utils.RequestResult{false, msg, "removeCoverages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsuredObject from the gorm
		//----------------------------------------------------------------------------
		return GetInsuredObject(insuredObjectId)

	} else {
		return parentRequestResult
	}
}

