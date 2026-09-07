package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InvestmentPortfolioDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInvestmentPortfolio - creates a new db entry
//----------------------------------------------------------------------------
func CreateInvestmentPortfolio(obj model.InvestmentPortfolio)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InvestmentPortfolio with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InvestmentPortfolio", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInvestmentPortfolio", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInvestmentPortfolio - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInvestmentPortfolio(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InvestmentPortfolio

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InvestmentPortfolio with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InvestmentPortfolio using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InvestmentPortfolio using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInvestmentPortfolio", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInvestmentPortfolio - returns all
//----------------------------------------------------------------------------
func GetAllInvestmentPortfolio()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InvestmentPortfolio

	//----------------------------------------------------------------------------
	// Request the ORM to find all InvestmentPortfolio
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InvestmentPortfolio" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InvestmentPortfolio", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInvestmentPortfolio", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInvestmentPortfolio - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInvestmentPortfolio(obj model.InvestmentPortfolio)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InvestmentPortfolio using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InvestmentPortfolio using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInvestmentPortfolio", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInvestmentPortfolio - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInvestmentPortfolio(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InvestmentPortfolio with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInvestmentPortfolio(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentPortfolio so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InvestmentPortfolio)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InvestmentPortfolio using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InvestmentPortfolio using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInvestmentPortfolio", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a InvestmentPortfolio
//----------------------------------------------------------------------------
func AssignCustomerToInvestmentPortfolio( investmentPortfolioId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InvestmentPortfolio with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentPortfolio(investmentPortfolioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentPortfolio so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentPortfolio)

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
			// assign the Customer	to the InvestmentPortfolio
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the InvestmentPortfolio
			//----------------------------------------------------------------------------
			return UpdateInvestmentPortfolio(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a InvestmentPortfolio
//----------------------------------------------------------------------------
func UnassignCustomerFromInvestmentPortfolio(investmentPortfolioId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InvestmentPortfolio with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentPortfolio(investmentPortfolioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentPortfolio so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentPortfolio)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the InvestmentPortfolio
		//----------------------------------------------------------------------------
		return UpdateInvestmentPortfolio(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more accountsIds as a Accounts to a InvestmentPortfolio
//----------------------------------------------------------------------------
func AddAccountsToInvestmentPortfolio ( investmentPortfolioId uint64, accountsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InvestmentPortfolio with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentPortfolio(investmentPortfolioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentPortfolio so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentPortfolio)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InvestmentAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InvestmentAccount
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Accounts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "unassignAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InvestmentPortfolio from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentPortfolio(investmentPortfolioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more accountsIds as a Accounts from a InvestmentPortfolio
//----------------------------------------------------------------------------
func RemoveAccountsFromInvestmentPortfolio( investmentPortfolioId uint64, accountsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InvestmentPortfolio with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentPortfolio(investmentPortfolioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentPortfolio so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentPortfolio)

		// slice the ids on comma with no spaces
		ids := strings.Split( accountsIds, ",")

		for _, accountsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InvestmentAccount

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InvestmentAccount
			// with a matching accountsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , accountsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InvestmentAccountObj from the Accounts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Accounts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Accounts", accountsId )
				return utils.RequestResult{false, msg, "removeAccounts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InvestmentPortfolio from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentPortfolio(investmentPortfolioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more ordersIds as a Orders to a InvestmentPortfolio
//----------------------------------------------------------------------------
func AddOrdersToInvestmentPortfolio ( investmentPortfolioId uint64, ordersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InvestmentPortfolio with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentPortfolio(investmentPortfolioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentPortfolio so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentPortfolio)

		// slice the ids on comma with no spaces
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TradeOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TradeOrder
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Orders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "unassignOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InvestmentPortfolio from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentPortfolio(investmentPortfolioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ordersIds as a Orders from a InvestmentPortfolio
//----------------------------------------------------------------------------
func RemoveOrdersFromInvestmentPortfolio( investmentPortfolioId uint64, ordersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InvestmentPortfolio with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentPortfolio(investmentPortfolioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentPortfolio so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentPortfolio)

		// slice the ids on comma with no spaces
		ids := strings.Split( ordersIds, ",")

		for _, ordersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TradeOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TradeOrder
			// with a matching ordersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , ordersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TradeOrderObj from the Orders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Orders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Orders", ordersId )
				return utils.RequestResult{false, msg, "removeOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InvestmentPortfolio from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentPortfolio(investmentPortfolioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more holdingsIds as a Holdings to a InvestmentPortfolio
//----------------------------------------------------------------------------
func AddHoldingsToInvestmentPortfolio ( investmentPortfolioId uint64, holdingsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InvestmentPortfolio with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentPortfolio(investmentPortfolioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentPortfolio so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentPortfolio)

		// slice the ids on comma with no spaces
		ids := strings.Split( holdingsIds, ",")

		for _, holdingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Position

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Position
			// with a matching holdingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , holdingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Holdings using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Holdings").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Holdings", holdingsId )
				return utils.RequestResult{false, msg, "unassignHoldings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InvestmentPortfolio from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentPortfolio(investmentPortfolioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more holdingsIds as a Holdings from a InvestmentPortfolio
//----------------------------------------------------------------------------
func RemoveHoldingsFromInvestmentPortfolio( investmentPortfolioId uint64, holdingsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InvestmentPortfolio with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentPortfolio(investmentPortfolioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentPortfolio so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentPortfolio)

		// slice the ids on comma with no spaces
		ids := strings.Split( holdingsIds, ",")

		for _, holdingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Position

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Position
			// with a matching holdingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , holdingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PositionObj from the Holdings array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Holdings").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Holdings", holdingsId )
				return utils.RequestResult{false, msg, "removeHoldings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InvestmentPortfolio from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentPortfolio(investmentPortfolioId)

	} else {
		return parentRequestResult
	}
}

