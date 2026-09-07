package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ConsentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateConsent - creates a new db entry
//----------------------------------------------------------------------------
func CreateConsent(obj model.Consent)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Consent with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Consent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateConsent", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetConsent - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetConsent(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Consent

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Consent with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Consent using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Consent using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetConsent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllConsent - returns all
//----------------------------------------------------------------------------
func GetAllConsent()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Consent

	//----------------------------------------------------------------------------
	// Request the ORM to find all Consent
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Consent" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Consent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllConsent", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateConsent - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateConsent(obj model.Consent)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Consent using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Consent using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateConsent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteConsent - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteConsent(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetConsent(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Consent using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Consent using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteConsent", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a Consent
//----------------------------------------------------------------------------
func AssignCustomerToConsent( consentId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

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
			// assign the Customer	to the Consent
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Consent
			//----------------------------------------------------------------------------
			return UpdateConsent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Consent
//----------------------------------------------------------------------------
func UnassignCustomerFromConsent(consentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Consent
		//----------------------------------------------------------------------------
		return UpdateConsent(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ApiClient on a Consent
//----------------------------------------------------------------------------
func AssignApiClientToConsent( consentId uint64, apiClientId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.APIClient

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a APIClient with a
		// matching apiClientId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, apiClientId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ApiClient	to the Consent
			//----------------------------------------------------------------------------
			parentObj.ApiClient = &childObj

			//----------------------------------------------------------------------------
			// save the Consent
			//----------------------------------------------------------------------------
			return UpdateConsent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ApiClient", apiClientId )
			return utils.RequestResult{false, msg, "assignApiClient", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ApiClient on a Consent
//----------------------------------------------------------------------------
func UnassignApiClientFromConsent(consentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Consent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConsent(consentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Consent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Consent)

		//----------------------------------------------------------------------------
		// assign an empty APIClient to the ApiClient
		//----------------------------------------------------------------------------
		parentObj.ApiClient = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ApiClient
		//----------------------------------------------------------------------------
		parentObj.ApiClientId = nil;

		//----------------------------------------------------------------------------
		// save the Consent
		//----------------------------------------------------------------------------
		return UpdateConsent(parentObj)

	} else {
		return parentRequestResult
	}

}


