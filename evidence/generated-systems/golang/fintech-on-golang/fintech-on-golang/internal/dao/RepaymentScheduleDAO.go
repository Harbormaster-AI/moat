package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RepaymentScheduleDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRepaymentSchedule - creates a new db entry
//----------------------------------------------------------------------------
func CreateRepaymentSchedule(obj model.RepaymentSchedule)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RepaymentSchedule with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RepaymentSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRepaymentSchedule", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRepaymentSchedule - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRepaymentSchedule(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RepaymentSchedule

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RepaymentSchedule with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RepaymentSchedule using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RepaymentSchedule using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRepaymentSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRepaymentSchedule - returns all
//----------------------------------------------------------------------------
func GetAllRepaymentSchedule()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RepaymentSchedule

	//----------------------------------------------------------------------------
	// Request the ORM to find all RepaymentSchedule
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all RepaymentSchedule" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RepaymentSchedule", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRepaymentSchedule", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRepaymentSchedule - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRepaymentSchedule(obj model.RepaymentSchedule)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RepaymentSchedule using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RepaymentSchedule using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRepaymentSchedule", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRepaymentSchedule - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRepaymentSchedule(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRepaymentSchedule(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.RepaymentSchedule)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RepaymentSchedule using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RepaymentSchedule using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRepaymentSchedule", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Loan on a RepaymentSchedule
//----------------------------------------------------------------------------
func AssignLoanToRepaymentSchedule( repaymentScheduleId uint64, loanId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RepaymentSchedule)

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
			// assign the Loan	to the RepaymentSchedule
			//----------------------------------------------------------------------------
			parentObj.Loan = &childObj

			//----------------------------------------------------------------------------
			// save the RepaymentSchedule
			//----------------------------------------------------------------------------
			return UpdateRepaymentSchedule(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Loan", loanId )
			return utils.RequestResult{false, msg, "assignLoan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Loan on a RepaymentSchedule
//----------------------------------------------------------------------------
func UnassignLoanFromRepaymentSchedule(repaymentScheduleId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RepaymentSchedule)

		//----------------------------------------------------------------------------
		// assign an empty Loan to the Loan
		//----------------------------------------------------------------------------
		parentObj.Loan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Loan
		//----------------------------------------------------------------------------
		parentObj.LoanId = nil;

		//----------------------------------------------------------------------------
		// save the RepaymentSchedule
		//----------------------------------------------------------------------------
		return UpdateRepaymentSchedule(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more paymentsIds as a Payments to a RepaymentSchedule
//----------------------------------------------------------------------------
func AddPaymentsToRepaymentSchedule ( repaymentScheduleId uint64, paymentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RepaymentSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Payments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "unassignPayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RepaymentSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRepaymentSchedule(repaymentScheduleId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentsIds as a Payments from a RepaymentSchedule
//----------------------------------------------------------------------------
func RemovePaymentsFromRepaymentSchedule( repaymentScheduleId uint64, paymentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the RepaymentSchedule with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRepaymentSchedule(repaymentScheduleId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RepaymentSchedule so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RepaymentSchedule)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Transaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Transaction
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TransactionObj from the Payments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "removePayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified RepaymentSchedule from the gorm
		//----------------------------------------------------------------------------
		return GetRepaymentSchedule(repaymentScheduleId)

	} else {
		return parentRequestResult
	}
}

