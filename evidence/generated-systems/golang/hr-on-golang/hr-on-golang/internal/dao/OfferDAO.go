package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OfferDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOffer - creates a new db entry
//----------------------------------------------------------------------------
func CreateOffer(obj model.Offer)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Offer with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Offer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOffer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOffer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOffer(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Offer

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Offer with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Offer using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Offer using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOffer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOffer - returns all
//----------------------------------------------------------------------------
func GetAllOffer()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Offer

	//----------------------------------------------------------------------------
	// Request the ORM to find all Offer
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Offer" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Offer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOffer", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOffer - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOffer(obj model.Offer)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Offer using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Offer using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOffer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOffer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOffer(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Offer with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOffer(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Offer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Offer)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Offer using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Offer using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOffer", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Requisition on a Offer
//----------------------------------------------------------------------------
func AssignRequisitionToOffer( offerId uint64, requisitionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Offer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOffer(offerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Offer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Offer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.JobRequisition

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a JobRequisition with a
		// matching requisitionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, requisitionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Requisition	to the Offer
			//----------------------------------------------------------------------------
			parentObj.Requisition = &childObj

			//----------------------------------------------------------------------------
			// save the Offer
			//----------------------------------------------------------------------------
			return UpdateOffer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Requisition", requisitionId )
			return utils.RequestResult{false, msg, "assignRequisition", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Requisition on a Offer
//----------------------------------------------------------------------------
func UnassignRequisitionFromOffer(offerId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Offer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOffer(offerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Offer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Offer)

		//----------------------------------------------------------------------------
		// assign an empty JobRequisition to the Requisition
		//----------------------------------------------------------------------------
		parentObj.Requisition = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Requisition
		//----------------------------------------------------------------------------
		parentObj.RequisitionId = nil;

		//----------------------------------------------------------------------------
		// save the Offer
		//----------------------------------------------------------------------------
		return UpdateOffer(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Candidate on a Offer
//----------------------------------------------------------------------------
func AssignCandidateToOffer( offerId uint64, candidateId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Offer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOffer(offerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Offer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Offer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Candidate

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Candidate with a
		// matching candidateId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, candidateId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Candidate	to the Offer
			//----------------------------------------------------------------------------
			parentObj.Candidate = &childObj

			//----------------------------------------------------------------------------
			// save the Offer
			//----------------------------------------------------------------------------
			return UpdateOffer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Candidate", candidateId )
			return utils.RequestResult{false, msg, "assignCandidate", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Candidate on a Offer
//----------------------------------------------------------------------------
func UnassignCandidateFromOffer(offerId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Offer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOffer(offerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Offer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Offer)

		//----------------------------------------------------------------------------
		// assign an empty Candidate to the Candidate
		//----------------------------------------------------------------------------
		parentObj.Candidate = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Candidate
		//----------------------------------------------------------------------------
		parentObj.CandidateId = nil;

		//----------------------------------------------------------------------------
		// save the Offer
		//----------------------------------------------------------------------------
		return UpdateOffer(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ApprovedBy on a Offer
//----------------------------------------------------------------------------
func AssignApprovedByToOffer( offerId uint64, approvedById uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Offer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOffer(offerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Offer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Offer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching approvedById
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, approvedById).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ApprovedBy	to the Offer
			//----------------------------------------------------------------------------
			parentObj.ApprovedBy = &childObj

			//----------------------------------------------------------------------------
			// save the Offer
			//----------------------------------------------------------------------------
			return UpdateOffer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ApprovedBy", approvedById )
			return utils.RequestResult{false, msg, "assignApprovedBy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ApprovedBy on a Offer
//----------------------------------------------------------------------------
func UnassignApprovedByFromOffer(offerId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Offer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOffer(offerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Offer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Offer)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the ApprovedBy
		//----------------------------------------------------------------------------
		parentObj.ApprovedBy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ApprovedBy
		//----------------------------------------------------------------------------
		parentObj.ApprovedById = nil;

		//----------------------------------------------------------------------------
		// save the Offer
		//----------------------------------------------------------------------------
		return UpdateOffer(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Contract on a Offer
//----------------------------------------------------------------------------
func AssignContractToOffer( offerId uint64, contractId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Offer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOffer(offerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Offer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Offer)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.EmploymentContract

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a EmploymentContract with a
		// matching contractId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, contractId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Contract	to the Offer
			//----------------------------------------------------------------------------
			parentObj.Contract = &childObj

			//----------------------------------------------------------------------------
			// save the Offer
			//----------------------------------------------------------------------------
			return UpdateOffer(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Contract", contractId )
			return utils.RequestResult{false, msg, "assignContract", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Contract on a Offer
//----------------------------------------------------------------------------
func UnassignContractFromOffer(offerId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Offer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOffer(offerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Offer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Offer)

		//----------------------------------------------------------------------------
		// assign an empty EmploymentContract to the Contract
		//----------------------------------------------------------------------------
		parentObj.Contract = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Contract
		//----------------------------------------------------------------------------
		parentObj.ContractId = nil;

		//----------------------------------------------------------------------------
		// save the Offer
		//----------------------------------------------------------------------------
		return UpdateOffer(parentObj)

	} else {
		return parentRequestResult
	}

}


