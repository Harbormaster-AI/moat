package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LoanDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLoan - creates a new db entry
//----------------------------------------------------------------------------
func CreateLoan(obj model.Loan)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Loan with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Loan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLoan", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLoan - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLoan(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Loan

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Loan with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Loan using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Loan using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLoan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLoan - returns all
//----------------------------------------------------------------------------
func GetAllLoan()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Loan

	//----------------------------------------------------------------------------
	// Request the ORM to find all Loan
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Loan" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Loan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLoan", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLoan - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLoan(obj model.Loan)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Loan using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Loan using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLoan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLoan - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLoan(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Loan with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLoan(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Loan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Loan)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Loan using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Loan using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLoan", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a Loan
//----------------------------------------------------------------------------
func AssignCustomerToLoan( loanId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Loan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoan(loanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Loan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Loan)

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
			// assign the Customer	to the Loan
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the Loan
			//----------------------------------------------------------------------------
			return UpdateLoan(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a Loan
//----------------------------------------------------------------------------
func UnassignCustomerFromLoan(loanId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Loan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoan(loanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Loan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Loan)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the Loan
		//----------------------------------------------------------------------------
		return UpdateLoan(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more scheduleIds as a Schedule to a Loan
//----------------------------------------------------------------------------
func AddScheduleToLoan ( loanId uint64, scheduleIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Loan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoan(loanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Loan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Loan)

		// slice the ids on comma with no spaces
		ids := strings.Split( scheduleIds, ",")

		for _, scheduleId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RepaymentSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RepaymentSchedule
			// with a matching scheduleId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , scheduleId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Schedule using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Schedule").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Schedule", scheduleId )
				return utils.RequestResult{false, msg, "unassignSchedule", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Loan from the gorm
		//----------------------------------------------------------------------------
		return GetLoan(loanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more scheduleIds as a Schedule from a Loan
//----------------------------------------------------------------------------
func RemoveScheduleFromLoan( loanId uint64, scheduleIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Loan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoan(loanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Loan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Loan)

		// slice the ids on comma with no spaces
		ids := strings.Split( scheduleIds, ",")

		for _, scheduleId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RepaymentSchedule

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RepaymentSchedule
			// with a matching scheduleId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , scheduleId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RepaymentScheduleObj from the Schedule array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Schedule").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Schedule", scheduleId )
				return utils.RequestResult{false, msg, "removeSchedule", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Loan from the gorm
		//----------------------------------------------------------------------------
		return GetLoan(loanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more collateralIds as a Collateral to a Loan
//----------------------------------------------------------------------------
func AddCollateralToLoan ( loanId uint64, collateralIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Loan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoan(loanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Loan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Loan)

		// slice the ids on comma with no spaces
		ids := strings.Split( collateralIds, ",")

		for _, collateralId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Collateral

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Collateral
			// with a matching collateralId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , collateralId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Collateral using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Collateral").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Collateral", collateralId )
				return utils.RequestResult{false, msg, "unassignCollateral", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Loan from the gorm
		//----------------------------------------------------------------------------
		return GetLoan(loanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more collateralIds as a Collateral from a Loan
//----------------------------------------------------------------------------
func RemoveCollateralFromLoan( loanId uint64, collateralIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Loan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoan(loanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Loan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Loan)

		// slice the ids on comma with no spaces
		ids := strings.Split( collateralIds, ",")

		for _, collateralId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Collateral

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Collateral
			// with a matching collateralId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , collateralId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CollateralObj from the Collateral array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Collateral").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Collateral", collateralId )
				return utils.RequestResult{false, msg, "removeCollateral", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Loan from the gorm
		//----------------------------------------------------------------------------
		return GetLoan(loanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more transactionsIds as a Transactions to a Loan
//----------------------------------------------------------------------------
func AddTransactionsToLoan ( loanId uint64, transactionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Loan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoan(loanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Loan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Loan)

		// slice the ids on comma with no spaces
		ids := strings.Split( transactionsIds, ",")

		for _, transactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LoanTransaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanTransaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Transactions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "unassignTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Loan from the gorm
		//----------------------------------------------------------------------------
		return GetLoan(loanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more transactionsIds as a Transactions from a Loan
//----------------------------------------------------------------------------
func RemoveTransactionsFromLoan( loanId uint64, transactionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Loan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLoan(loanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Loan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Loan)

		// slice the ids on comma with no spaces
		ids := strings.Split( transactionsIds, ",")

		for _, transactionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LoanTransaction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LoanTransaction
			// with a matching transactionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , transactionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LoanTransactionObj from the Transactions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Transactions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Transactions", transactionsId )
				return utils.RequestResult{false, msg, "removeTransactions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Loan from the gorm
		//----------------------------------------------------------------------------
		return GetLoan(loanId)

	} else {
		return parentRequestResult
	}
}

