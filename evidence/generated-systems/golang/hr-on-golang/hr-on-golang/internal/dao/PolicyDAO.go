package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
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
// adds one or more acknowledgementsIds as a Acknowledgements to a Policy
//----------------------------------------------------------------------------
func AddAcknowledgementsToPolicy ( policyId uint64, acknowledgementsIds string )(utils.RequestResult) {

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
		ids := strings.Split( acknowledgementsIds, ",")

		for _, acknowledgementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PolicyAcknowledgement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PolicyAcknowledgement
			// with a matching acknowledgementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , acknowledgementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Acknowledgements using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Acknowledgements").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Acknowledgements", acknowledgementsId )
				return utils.RequestResult{false, msg, "unassignAcknowledgements", childObj}
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
// removes one or more acknowledgementsIds as a Acknowledgements from a Policy
//----------------------------------------------------------------------------
func RemoveAcknowledgementsFromPolicy( policyId uint64, acknowledgementsIds string )(utils.RequestResult) {
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
		ids := strings.Split( acknowledgementsIds, ",")

		for _, acknowledgementsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PolicyAcknowledgement

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PolicyAcknowledgement
			// with a matching acknowledgementsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , acknowledgementsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PolicyAcknowledgementObj from the Acknowledgements array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Acknowledgements").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Acknowledgements", acknowledgementsId )
				return utils.RequestResult{false, msg, "removeAcknowledgements", childObj}
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

