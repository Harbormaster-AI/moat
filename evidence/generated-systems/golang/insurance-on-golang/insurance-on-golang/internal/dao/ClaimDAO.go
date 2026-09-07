package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ClaimDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateClaim - creates a new db entry
//----------------------------------------------------------------------------
func CreateClaim(obj model.Claim)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Claim with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Claim", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateClaim", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetClaim - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetClaim(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Claim

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Claim with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Claim using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Claim using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetClaim", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllClaim - returns all
//----------------------------------------------------------------------------
func GetAllClaim()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Claim

	//----------------------------------------------------------------------------
	// Request the ORM to find all Claim
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Claim" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Claim", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllClaim", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateClaim - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateClaim(obj model.Claim)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Claim using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Claim using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateClaim", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteClaim - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteClaim(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetClaim(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Claim using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Claim using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteClaim", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Policy on a Claim
//----------------------------------------------------------------------------
func AssignPolicyToClaim( claimId uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

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
			// assign the Policy	to the Claim
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the Claim
			//----------------------------------------------------------------------------
			return UpdateClaim(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a Claim
//----------------------------------------------------------------------------
func UnassignPolicyFromClaim(claimId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the Claim
		//----------------------------------------------------------------------------
		return UpdateClaim(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Customer on a Claim
//----------------------------------------------------------------------------
func AssignCustomerToClaim( claimId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Customer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Customer with a
		// matching customerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, customerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Customer	to the Claim
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Claim
			//----------------------------------------------------------------------------
			return UpdateClaim(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Claim
//----------------------------------------------------------------------------
func UnassignCustomerFromClaim(claimId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Claim
		//----------------------------------------------------------------------------
		return UpdateClaim(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Adjuster on a Claim
//----------------------------------------------------------------------------
func AssignAdjusterToClaim( claimId uint64, adjusterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Adjuster

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Adjuster with a
		// matching adjusterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, adjusterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Adjuster	to the Claim
			//----------------------------------------------------------------------------
			parentObj.Adjuster = &childObj

			//----------------------------------------------------------------------------
			// save the Claim
			//----------------------------------------------------------------------------
			return UpdateClaim(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Adjuster", adjusterId )
			return utils.RequestResult{false, msg, "assignAdjuster", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Adjuster on a Claim
//----------------------------------------------------------------------------
func UnassignAdjusterFromClaim(claimId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// assign an empty Adjuster to the Adjuster
		//----------------------------------------------------------------------------
		parentObj.Adjuster = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Adjuster
		//----------------------------------------------------------------------------
		parentObj.AdjusterId = nil;

		//----------------------------------------------------------------------------
		// save the Claim
		//----------------------------------------------------------------------------
		return UpdateClaim(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Incident on a Claim
//----------------------------------------------------------------------------
func AssignIncidentToClaim( claimId uint64, incidentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Incident

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Incident with a
		// matching incidentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, incidentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Incident	to the Claim
			//----------------------------------------------------------------------------
			parentObj.Incident = &childObj

			//----------------------------------------------------------------------------
			// save the Claim
			//----------------------------------------------------------------------------
			return UpdateClaim(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Incident", incidentId )
			return utils.RequestResult{false, msg, "assignIncident", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Incident on a Claim
//----------------------------------------------------------------------------
func UnassignIncidentFromClaim(claimId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		//----------------------------------------------------------------------------
		// assign an empty Incident to the Incident
		//----------------------------------------------------------------------------
		parentObj.Incident = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Incident
		//----------------------------------------------------------------------------
		parentObj.IncidentId = nil;

		//----------------------------------------------------------------------------
		// save the Claim
		//----------------------------------------------------------------------------
		return UpdateClaim(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more exposuresIds as a Exposures to a Claim
//----------------------------------------------------------------------------
func AddExposuresToClaim ( claimId uint64, exposuresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( exposuresIds, ",")

		for _, exposuresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Exposure

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Exposure
			// with a matching exposuresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exposuresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Exposures using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Exposures").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exposures", exposuresId )
				return utils.RequestResult{false, msg, "unassignExposures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more exposuresIds as a Exposures from a Claim
//----------------------------------------------------------------------------
func RemoveExposuresFromClaim( claimId uint64, exposuresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( exposuresIds, ",")

		for _, exposuresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Exposure

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Exposure
			// with a matching exposuresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , exposuresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ExposureObj from the Exposures array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Exposures").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exposures", exposuresId )
				return utils.RequestResult{false, msg, "removeExposures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reservesIds as a Reserves to a Claim
//----------------------------------------------------------------------------
func AddReservesToClaim ( claimId uint64, reservesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( reservesIds, ",")

		for _, reservesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClaimReserve

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClaimReserve
			// with a matching reservesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reservesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reserves using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reserves").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reserves", reservesId )
				return utils.RequestResult{false, msg, "unassignReserves", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reservesIds as a Reserves from a Claim
//----------------------------------------------------------------------------
func RemoveReservesFromClaim( claimId uint64, reservesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( reservesIds, ",")

		for _, reservesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClaimReserve

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClaimReserve
			// with a matching reservesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reservesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimReserveObj from the Reserves array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reserves").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reserves", reservesId )
				return utils.RequestResult{false, msg, "removeReserves", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more claimPaymentsIds as a ClaimPayments to a Claim
//----------------------------------------------------------------------------
func AddClaimPaymentsToClaim ( claimId uint64, claimPaymentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimPaymentsIds, ",")

		for _, claimPaymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClaimPayment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClaimPayment
			// with a matching claimPaymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimPaymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ClaimPayments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ClaimPayments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ClaimPayments", claimPaymentsId )
				return utils.RequestResult{false, msg, "unassignClaimPayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more claimPaymentsIds as a ClaimPayments from a Claim
//----------------------------------------------------------------------------
func RemoveClaimPaymentsFromClaim( claimId uint64, claimPaymentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimPaymentsIds, ",")

		for _, claimPaymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClaimPayment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClaimPayment
			// with a matching claimPaymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimPaymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimPaymentObj from the ClaimPayments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ClaimPayments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ClaimPayments", claimPaymentsId )
				return utils.RequestResult{false, msg, "removeClaimPayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more serviceProvidersIds as a ServiceProviders to a Claim
//----------------------------------------------------------------------------
func AddServiceProvidersToClaim ( claimId uint64, serviceProvidersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( serviceProvidersIds, ",")

		for _, serviceProvidersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ServiceProvider

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ServiceProvider
			// with a matching serviceProvidersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , serviceProvidersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ServiceProviders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ServiceProviders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ServiceProviders", serviceProvidersId )
				return utils.RequestResult{false, msg, "unassignServiceProviders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more serviceProvidersIds as a ServiceProviders from a Claim
//----------------------------------------------------------------------------
func RemoveServiceProvidersFromClaim( claimId uint64, serviceProvidersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( serviceProvidersIds, ",")

		for _, serviceProvidersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ServiceProvider

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ServiceProvider
			// with a matching serviceProvidersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , serviceProvidersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ServiceProviderObj from the ServiceProviders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ServiceProviders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ServiceProviders", serviceProvidersId )
				return utils.RequestResult{false, msg, "removeServiceProviders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more subrogationsIds as a Subrogations to a Claim
//----------------------------------------------------------------------------
func AddSubrogationsToClaim ( claimId uint64, subrogationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( subrogationsIds, ",")

		for _, subrogationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SubrogationRecovery

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SubrogationRecovery
			// with a matching subrogationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , subrogationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Subrogations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Subrogations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Subrogations", subrogationsId )
				return utils.RequestResult{false, msg, "unassignSubrogations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more subrogationsIds as a Subrogations from a Claim
//----------------------------------------------------------------------------
func RemoveSubrogationsFromClaim( claimId uint64, subrogationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Claim with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaim(claimId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Claim so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Claim)

		// slice the ids on comma with no spaces
		ids := strings.Split( subrogationsIds, ",")

		for _, subrogationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SubrogationRecovery

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SubrogationRecovery
			// with a matching subrogationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , subrogationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SubrogationRecoveryObj from the Subrogations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Subrogations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Subrogations", subrogationsId )
				return utils.RequestResult{false, msg, "removeSubrogations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Claim from the gorm
		//----------------------------------------------------------------------------
		return GetClaim(claimId)

	} else {
		return parentRequestResult
	}
}

