package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AgreementDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAgreement - creates a new db entry
//----------------------------------------------------------------------------
func CreateAgreement(obj model.Agreement)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Agreement with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Agreement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAgreement", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAgreement - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAgreement(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Agreement

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Agreement with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Agreement using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Agreement using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAgreement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAgreement - returns all
//----------------------------------------------------------------------------
func GetAllAgreement()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Agreement

	//----------------------------------------------------------------------------
	// Request the ORM to find all Agreement
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Agreement" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Agreement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAgreement", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAgreement - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAgreement(obj model.Agreement)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Agreement using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Agreement using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAgreement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAgreement - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAgreement(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Agreement with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAgreement(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Agreement)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Agreement using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Agreement using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAgreement", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a Agreement
//----------------------------------------------------------------------------
func AssignCustomerToAgreement( agreementId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Agreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgreement(agreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agreement)

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
			// assign the Customer	to the Agreement
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Agreement
			//----------------------------------------------------------------------------
			return UpdateAgreement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Agreement
//----------------------------------------------------------------------------
func UnassignCustomerFromAgreement(agreementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Agreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgreement(agreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agreement)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Agreement
		//----------------------------------------------------------------------------
		return UpdateAgreement(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ProductOffering on a Agreement
//----------------------------------------------------------------------------
func AssignProductOfferingToAgreement( agreementId uint64, productOfferingId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Agreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgreement(agreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agreement)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ProductOffering

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ProductOffering with a
		// matching productOfferingId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productOfferingId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ProductOffering	to the Agreement
			//----------------------------------------------------------------------------
			parentObj.ProductOffering = &childObj

			//----------------------------------------------------------------------------
			// save the Agreement
			//----------------------------------------------------------------------------
			return UpdateAgreement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProductOffering", productOfferingId )
			return utils.RequestResult{false, msg, "assignProductOffering", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ProductOffering on a Agreement
//----------------------------------------------------------------------------
func UnassignProductOfferingFromAgreement(agreementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Agreement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAgreement(agreementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Agreement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Agreement)

		//----------------------------------------------------------------------------
		// assign an empty ProductOffering to the ProductOffering
		//----------------------------------------------------------------------------
		parentObj.ProductOffering = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ProductOffering
		//----------------------------------------------------------------------------
		parentObj.ProductOfferingId = nil;

		//----------------------------------------------------------------------------
		// save the Agreement
		//----------------------------------------------------------------------------
		return UpdateAgreement(parentObj)

	} else {
		return parentRequestResult
	}

}


