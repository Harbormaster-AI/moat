package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ClaimPaymentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateClaimPayment - creates a new db entry
//----------------------------------------------------------------------------
func CreateClaimPayment(obj model.ClaimPayment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ClaimPayment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ClaimPayment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateClaimPayment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetClaimPayment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetClaimPayment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ClaimPayment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ClaimPayment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ClaimPayment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ClaimPayment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetClaimPayment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllClaimPayment - returns all
//----------------------------------------------------------------------------
func GetAllClaimPayment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ClaimPayment

	//----------------------------------------------------------------------------
	// Request the ORM to find all ClaimPayment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ClaimPayment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ClaimPayment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllClaimPayment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateClaimPayment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateClaimPayment(obj model.ClaimPayment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ClaimPayment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ClaimPayment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateClaimPayment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteClaimPayment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteClaimPayment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetClaimPayment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ClaimPayment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ClaimPayment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteClaimPayment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Claim on a ClaimPayment
//----------------------------------------------------------------------------
func AssignClaimToClaimPayment( claimPaymentId uint64, claimId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Claim

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Claim with a
		// matching claimId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, claimId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Claim	to the ClaimPayment
			//----------------------------------------------------------------------------
			parentObj.Claim = &childObj

			//----------------------------------------------------------------------------
			// save the ClaimPayment
			//----------------------------------------------------------------------------
			return UpdateClaimPayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claim", claimId )
			return utils.RequestResult{false, msg, "assignClaim", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Claim on a ClaimPayment
//----------------------------------------------------------------------------
func UnassignClaimFromClaimPayment(claimPaymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// assign an empty Claim to the Claim
		//----------------------------------------------------------------------------
		parentObj.Claim = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Claim
		//----------------------------------------------------------------------------
		parentObj.ClaimId = nil;

		//----------------------------------------------------------------------------
		// save the ClaimPayment
		//----------------------------------------------------------------------------
		return UpdateClaimPayment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Exposure on a ClaimPayment
//----------------------------------------------------------------------------
func AssignExposureToClaimPayment( claimPaymentId uint64, exposureId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Exposure

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Exposure with a
		// matching exposureId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, exposureId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Exposure	to the ClaimPayment
			//----------------------------------------------------------------------------
			parentObj.Exposure = &childObj

			//----------------------------------------------------------------------------
			// save the ClaimPayment
			//----------------------------------------------------------------------------
			return UpdateClaimPayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Exposure", exposureId )
			return utils.RequestResult{false, msg, "assignExposure", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Exposure on a ClaimPayment
//----------------------------------------------------------------------------
func UnassignExposureFromClaimPayment(claimPaymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// assign an empty Exposure to the Exposure
		//----------------------------------------------------------------------------
		parentObj.Exposure = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Exposure
		//----------------------------------------------------------------------------
		parentObj.ExposureId = nil;

		//----------------------------------------------------------------------------
		// save the ClaimPayment
		//----------------------------------------------------------------------------
		return UpdateClaimPayment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Beneficiary on a ClaimPayment
//----------------------------------------------------------------------------
func AssignBeneficiaryToClaimPayment( claimPaymentId uint64, beneficiaryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Beneficiary

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Beneficiary with a
		// matching beneficiaryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, beneficiaryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Beneficiary	to the ClaimPayment
			//----------------------------------------------------------------------------
			parentObj.Beneficiary = &childObj

			//----------------------------------------------------------------------------
			// save the ClaimPayment
			//----------------------------------------------------------------------------
			return UpdateClaimPayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Beneficiary", beneficiaryId )
			return utils.RequestResult{false, msg, "assignBeneficiary", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Beneficiary on a ClaimPayment
//----------------------------------------------------------------------------
func UnassignBeneficiaryFromClaimPayment(claimPaymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// assign an empty Beneficiary to the Beneficiary
		//----------------------------------------------------------------------------
		parentObj.Beneficiary = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Beneficiary
		//----------------------------------------------------------------------------
		parentObj.BeneficiaryId = nil;

		//----------------------------------------------------------------------------
		// save the ClaimPayment
		//----------------------------------------------------------------------------
		return UpdateClaimPayment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ServiceProvider on a ClaimPayment
//----------------------------------------------------------------------------
func AssignServiceProviderToClaimPayment( claimPaymentId uint64, serviceProviderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ServiceProvider

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ServiceProvider with a
		// matching serviceProviderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, serviceProviderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ServiceProvider	to the ClaimPayment
			//----------------------------------------------------------------------------
			parentObj.ServiceProvider = &childObj

			//----------------------------------------------------------------------------
			// save the ClaimPayment
			//----------------------------------------------------------------------------
			return UpdateClaimPayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ServiceProvider", serviceProviderId )
			return utils.RequestResult{false, msg, "assignServiceProvider", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ServiceProvider on a ClaimPayment
//----------------------------------------------------------------------------
func UnassignServiceProviderFromClaimPayment(claimPaymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// assign an empty ServiceProvider to the ServiceProvider
		//----------------------------------------------------------------------------
		parentObj.ServiceProvider = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ServiceProvider
		//----------------------------------------------------------------------------
		parentObj.ServiceProviderId = nil;

		//----------------------------------------------------------------------------
		// save the ClaimPayment
		//----------------------------------------------------------------------------
		return UpdateClaimPayment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Customer on a ClaimPayment
//----------------------------------------------------------------------------
func AssignCustomerToClaimPayment( claimPaymentId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

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
			// assign the Customer	to the ClaimPayment
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the ClaimPayment
			//----------------------------------------------------------------------------
			return UpdateClaimPayment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a ClaimPayment
//----------------------------------------------------------------------------
func UnassignCustomerFromClaimPayment(claimPaymentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ClaimPayment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetClaimPayment(claimPaymentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ClaimPayment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ClaimPayment)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the ClaimPayment
		//----------------------------------------------------------------------------
		return UpdateClaimPayment(parentObj)

	} else {
		return parentRequestResult
	}

}


