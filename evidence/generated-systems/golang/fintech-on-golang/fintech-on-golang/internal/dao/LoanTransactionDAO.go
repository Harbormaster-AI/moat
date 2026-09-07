package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LoanTransactionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLoanTransaction - creates a new db entry
//----------------------------------------------------------------------------
func CreateLoanTransaction(obj model.LoanTransaction)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LoanTransaction with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LoanTransaction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLoanTransaction", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLoanTransaction - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLoanTransaction(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LoanTransaction

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LoanTransaction with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LoanTransaction using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LoanTransaction using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLoanTransaction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLoanTransaction - returns all
//----------------------------------------------------------------------------
func GetAllLoanTransaction()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LoanTransaction

	//----------------------------------------------------------------------------
	// Request the ORM to find all LoanTransaction
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LoanTransaction" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LoanTransaction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLoanTransaction", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLoanTransaction - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLoanTransaction(obj model.LoanTransaction)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LoanTransaction using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LoanTransaction using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLoanTransaction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLoanTransaction - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLoanTransaction(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LoanTransaction with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLoanTransaction(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LoanTransaction)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LoanTransaction using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LoanTransaction using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLoanTransaction", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Loan on a LoanTransaction
//----------------------------------------------------------------------------
func AssignLoanToLoanTransaction( loanTransactionId uint64, loanId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LoanTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanTransaction(loanTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanTransaction)

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
			// assign the Loan	to the LoanTransaction
			//----------------------------------------------------------------------------
			parentObj.Loan = &childObj

			//----------------------------------------------------------------------------
			// save the LoanTransaction
			//----------------------------------------------------------------------------
			return UpdateLoanTransaction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Loan", loanId )
			return utils.RequestResult{false, msg, "assignLoan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Loan on a LoanTransaction
//----------------------------------------------------------------------------
func UnassignLoanFromLoanTransaction(loanTransactionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LoanTransaction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoanTransaction(loanTransactionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LoanTransaction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LoanTransaction)

		//----------------------------------------------------------------------------
		// assign an empty Loan to the Loan
		//----------------------------------------------------------------------------
		parentObj.Loan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Loan
		//----------------------------------------------------------------------------
		parentObj.LoanId = nil;

		//----------------------------------------------------------------------------
		// save the LoanTransaction
		//----------------------------------------------------------------------------
		return UpdateLoanTransaction(parentObj)

	} else {
		return parentRequestResult
	}

}


