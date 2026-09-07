package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LoanApplicationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLoanApplication - creates a new db entry
//----------------------------------------------------------------------------
func CreateLoanApplication(obj model.LoanApplication)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LoanApplication with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LoanApplication", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLoanApplication", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLoanApplication - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLoanApplication(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LoanApplication

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LoanApplication with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LoanApplication using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LoanApplication using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLoanApplication", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLoanApplication - returns all
//----------------------------------------------------------------------------
func GetAllLoanApplication()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LoanApplication

	//----------------------------------------------------------------------------
	// Request the ORM to find all LoanApplication
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LoanApplication" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LoanApplication", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLoanApplication", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLoanApplication - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLoanApplication(obj model.LoanApplication)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LoanApplication using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LoanApplication using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLoanApplication", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLoanApplication - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLoanApplication(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LoanApplication with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLoanApplication(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LoanApplication)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LoanApplication using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LoanApplication using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLoanApplication", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a LoanApplication
//----------------------------------------------------------------------------
func AssignCustomerToLoanApplication( loanApplicationId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanApplication(loanApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanApplication)

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
			// assign the Customer	to the LoanApplication
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the LoanApplication
			//----------------------------------------------------------------------------
			return UpdateLoanApplication(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a LoanApplication
//----------------------------------------------------------------------------
func UnassignCustomerFromLoanApplication(loanApplicationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanApplication(loanApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanApplication)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the LoanApplication
		//----------------------------------------------------------------------------
		return UpdateLoanApplication(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a RiskAssessment on a LoanApplication
//----------------------------------------------------------------------------
func AssignRiskAssessmentToLoanApplication( loanApplicationId uint64, riskAssessmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanApplication(loanApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanApplication)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.RiskAssessment

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a RiskAssessment with a
		// matching riskAssessmentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, riskAssessmentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the RiskAssessment	to the LoanApplication
			//----------------------------------------------------------------------------
			parentObj.RiskAssessment = &childObj

			//----------------------------------------------------------------------------
			// save the LoanApplication
			//----------------------------------------------------------------------------
			return UpdateLoanApplication(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RiskAssessment", riskAssessmentId )
			return utils.RequestResult{false, msg, "assignRiskAssessment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a RiskAssessment on a LoanApplication
//----------------------------------------------------------------------------
func UnassignRiskAssessmentFromLoanApplication(loanApplicationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanApplication(loanApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanApplication)

		//----------------------------------------------------------------------------
		// assign an empty RiskAssessment to the RiskAssessment
		//----------------------------------------------------------------------------
		parentObj.RiskAssessment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the RiskAssessment
		//----------------------------------------------------------------------------
		parentObj.RiskAssessmentId = nil;

		//----------------------------------------------------------------------------
		// save the LoanApplication
		//----------------------------------------------------------------------------
		return UpdateLoanApplication(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Loan on a LoanApplication
//----------------------------------------------------------------------------
func AssignLoanToLoanApplication( loanApplicationId uint64, loanId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanApplication(loanApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanApplication)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Loan

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Loan with a
		// matching loanId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, loanId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Loan	to the LoanApplication
			//----------------------------------------------------------------------------
			parentObj.Loan = &childObj

			//----------------------------------------------------------------------------
			// save the LoanApplication
			//----------------------------------------------------------------------------
			return UpdateLoanApplication(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Loan", loanId )
			return utils.RequestResult{false, msg, "assignLoan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Loan on a LoanApplication
//----------------------------------------------------------------------------
func UnassignLoanFromLoanApplication(loanApplicationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanApplication with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanApplication(loanApplicationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanApplication so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanApplication)

		//----------------------------------------------------------------------------
		// assign an empty Loan to the Loan
		//----------------------------------------------------------------------------
		parentObj.Loan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Loan
		//----------------------------------------------------------------------------
		parentObj.LoanId = nil;

		//----------------------------------------------------------------------------
		// save the LoanApplication
		//----------------------------------------------------------------------------
		return UpdateLoanApplication(parentObj)

	} else {
		return parentRequestResult
	}

}


