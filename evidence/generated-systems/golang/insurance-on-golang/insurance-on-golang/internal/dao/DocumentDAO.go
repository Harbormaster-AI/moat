package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DocumentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDocument - creates a new db entry
//----------------------------------------------------------------------------
func CreateDocument(obj model.Document)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Document with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Document", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDocument", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDocument - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDocument(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Document

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Document with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Document using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Document using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDocument", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDocument - returns all
//----------------------------------------------------------------------------
func GetAllDocument()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Document

	//----------------------------------------------------------------------------
	// Request the ORM to find all Document
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Document" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Document", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDocument", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDocument - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDocument(obj model.Document)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Document using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Document using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDocument", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDocument - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDocument(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDocument(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Document using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Document using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDocument", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Policy on a Document
//----------------------------------------------------------------------------
func AssignPolicyToDocument( documentId uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

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
			// assign the Policy	to the Document
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the Document
			//----------------------------------------------------------------------------
			return UpdateDocument(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a Document
//----------------------------------------------------------------------------
func UnassignPolicyFromDocument(documentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the Document
		//----------------------------------------------------------------------------
		return UpdateDocument(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Claim on a Document
//----------------------------------------------------------------------------
func AssignClaimToDocument( documentId uint64, claimId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

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
			// assign the Claim	to the Document
			//----------------------------------------------------------------------------
			parentObj.Claim = &childObj

			//----------------------------------------------------------------------------
			// save the Document
			//----------------------------------------------------------------------------
			return UpdateDocument(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claim", claimId )
			return utils.RequestResult{false, msg, "assignClaim", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Claim on a Document
//----------------------------------------------------------------------------
func UnassignClaimFromDocument(documentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// assign an empty Claim to the Claim
		//----------------------------------------------------------------------------
		parentObj.Claim = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Claim
		//----------------------------------------------------------------------------
		parentObj.ClaimId = nil;

		//----------------------------------------------------------------------------
		// save the Document
		//----------------------------------------------------------------------------
		return UpdateDocument(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Application on a Document
//----------------------------------------------------------------------------
func AssignApplicationToDocument( documentId uint64, applicationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Application

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Application with a
		// matching applicationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, applicationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Application	to the Document
			//----------------------------------------------------------------------------
			parentObj.Application = &childObj

			//----------------------------------------------------------------------------
			// save the Document
			//----------------------------------------------------------------------------
			return UpdateDocument(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Application", applicationId )
			return utils.RequestResult{false, msg, "assignApplication", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Application on a Document
//----------------------------------------------------------------------------
func UnassignApplicationFromDocument(documentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// assign an empty Application to the Application
		//----------------------------------------------------------------------------
		parentObj.Application = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Application
		//----------------------------------------------------------------------------
		parentObj.ApplicationId = nil;

		//----------------------------------------------------------------------------
		// save the Document
		//----------------------------------------------------------------------------
		return UpdateDocument(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Customer on a Document
//----------------------------------------------------------------------------
func AssignCustomerToDocument( documentId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

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
			// assign the Customer	to the Document
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Document
			//----------------------------------------------------------------------------
			return UpdateDocument(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Document
//----------------------------------------------------------------------------
func UnassignCustomerFromDocument(documentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Document with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDocument(documentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Document so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Document)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Document
		//----------------------------------------------------------------------------
		return UpdateDocument(parentObj)

	} else {
		return parentRequestResult
	}

}


