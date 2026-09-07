package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PolicyCoverageDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePolicyCoverage - creates a new db entry
//----------------------------------------------------------------------------
func CreatePolicyCoverage(obj model.PolicyCoverage)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PolicyCoverage with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PolicyCoverage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePolicyCoverage", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPolicyCoverage - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPolicyCoverage(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PolicyCoverage

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PolicyCoverage with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PolicyCoverage using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PolicyCoverage using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPolicyCoverage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPolicyCoverage - returns all
//----------------------------------------------------------------------------
func GetAllPolicyCoverage()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PolicyCoverage

	//----------------------------------------------------------------------------
	// Request the ORM to find all PolicyCoverage
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PolicyCoverage" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PolicyCoverage", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPolicyCoverage", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePolicyCoverage - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePolicyCoverage(obj model.PolicyCoverage)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PolicyCoverage using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PolicyCoverage using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePolicyCoverage", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePolicyCoverage - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePolicyCoverage(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PolicyCoverage with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPolicyCoverage(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyCoverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PolicyCoverage)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PolicyCoverage using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PolicyCoverage using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePolicyCoverage", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Policy on a PolicyCoverage
//----------------------------------------------------------------------------
func AssignPolicyToPolicyCoverage( policyCoverageId uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PolicyCoverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicyCoverage(policyCoverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyCoverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PolicyCoverage)

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
			// assign the Policy	to the PolicyCoverage
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the PolicyCoverage
			//----------------------------------------------------------------------------
			return UpdatePolicyCoverage(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a PolicyCoverage
//----------------------------------------------------------------------------
func UnassignPolicyFromPolicyCoverage(policyCoverageId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PolicyCoverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicyCoverage(policyCoverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyCoverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PolicyCoverage)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the PolicyCoverage
		//----------------------------------------------------------------------------
		return UpdatePolicyCoverage(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more insuredObjectsIds as a InsuredObjects to a PolicyCoverage
//----------------------------------------------------------------------------
func AddInsuredObjectsToPolicyCoverage ( policyCoverageId uint64, insuredObjectsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PolicyCoverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicyCoverage(policyCoverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyCoverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PolicyCoverage)

		// slice the ids on comma with no spaces
		ids := strings.Split( insuredObjectsIds, ",")

		for _, insuredObjectsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsuredObject

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsuredObject
			// with a matching insuredObjectsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insuredObjectsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InsuredObjects using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsuredObjects").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsuredObjects", insuredObjectsId )
				return utils.RequestResult{false, msg, "unassignInsuredObjects", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PolicyCoverage from the gorm
		//----------------------------------------------------------------------------
		return GetPolicyCoverage(policyCoverageId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more insuredObjectsIds as a InsuredObjects from a PolicyCoverage
//----------------------------------------------------------------------------
func RemoveInsuredObjectsFromPolicyCoverage( policyCoverageId uint64, insuredObjectsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the PolicyCoverage with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicyCoverage(policyCoverageId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyCoverage so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PolicyCoverage)

		// slice the ids on comma with no spaces
		ids := strings.Split( insuredObjectsIds, ",")

		for _, insuredObjectsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsuredObject

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsuredObject
			// with a matching insuredObjectsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insuredObjectsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InsuredObjectObj from the InsuredObjects array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsuredObjects").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsuredObjects", insuredObjectsId )
				return utils.RequestResult{false, msg, "removeInsuredObjects", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified PolicyCoverage from the gorm
		//----------------------------------------------------------------------------
		return GetPolicyCoverage(policyCoverageId)

	} else {
		return parentRequestResult
	}
}

