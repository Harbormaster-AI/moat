package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InvestmentAccountDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInvestmentAccount - creates a new db entry
//----------------------------------------------------------------------------
func CreateInvestmentAccount(obj model.InvestmentAccount)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InvestmentAccount with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InvestmentAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInvestmentAccount", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInvestmentAccount - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInvestmentAccount(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InvestmentAccount

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InvestmentAccount with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InvestmentAccount using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InvestmentAccount using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInvestmentAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInvestmentAccount - returns all
//----------------------------------------------------------------------------
func GetAllInvestmentAccount()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InvestmentAccount

	//----------------------------------------------------------------------------
	// Request the ORM to find all InvestmentAccount
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InvestmentAccount" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InvestmentAccount", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInvestmentAccount", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInvestmentAccount - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInvestmentAccount(obj model.InvestmentAccount)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InvestmentAccount using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InvestmentAccount using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInvestmentAccount", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInvestmentAccount - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInvestmentAccount(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InvestmentAccount with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInvestmentAccount(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InvestmentAccount)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InvestmentAccount using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InvestmentAccount using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInvestmentAccount", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Portfolio on a InvestmentAccount
//----------------------------------------------------------------------------
func AssignPortfolioToInvestmentAccount( investmentAccountId uint64, portfolioId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InvestmentAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentAccount(investmentAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentAccount)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InvestmentPortfolio

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InvestmentPortfolio with a
		// matching portfolioId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, portfolioId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Portfolio	to the InvestmentAccount
			//----------------------------------------------------------------------------
			parentObj.Portfolio = &childObj

			//----------------------------------------------------------------------------
			// save the InvestmentAccount
			//----------------------------------------------------------------------------
			return UpdateInvestmentAccount(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Portfolio", portfolioId )
			return utils.RequestResult{false, msg, "assignPortfolio", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Portfolio on a InvestmentAccount
//----------------------------------------------------------------------------
func UnassignPortfolioFromInvestmentAccount(investmentAccountId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InvestmentAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentAccount(investmentAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentAccount)

		//----------------------------------------------------------------------------
		// assign an empty InvestmentPortfolio to the Portfolio
		//----------------------------------------------------------------------------
		parentObj.Portfolio = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Portfolio
		//----------------------------------------------------------------------------
		parentObj.PortfolioId = nil;

		//----------------------------------------------------------------------------
		// save the InvestmentAccount
		//----------------------------------------------------------------------------
		return UpdateInvestmentAccount(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more tradesIds as a Trades to a InvestmentAccount
//----------------------------------------------------------------------------
func AddTradesToInvestmentAccount ( investmentAccountId uint64, tradesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InvestmentAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentAccount(investmentAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( tradesIds, ",")

		for _, tradesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Trade

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Trade
			// with a matching tradesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , tradesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Trades using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Trades").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Trades", tradesId )
				return utils.RequestResult{false, msg, "unassignTrades", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InvestmentAccount from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentAccount(investmentAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more tradesIds as a Trades from a InvestmentAccount
//----------------------------------------------------------------------------
func RemoveTradesFromInvestmentAccount( investmentAccountId uint64, tradesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InvestmentAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentAccount(investmentAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentAccount)

		// slice the ids on comma with no spaces
		ids := strings.Split( tradesIds, ",")

		for _, tradesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Trade

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Trade
			// with a matching tradesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , tradesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TradeObj from the Trades array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Trades").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Trades", tradesId )
				return utils.RequestResult{false, msg, "removeTrades", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InvestmentAccount from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentAccount(investmentAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more ordersIds as a Orders to a InvestmentAccount
//----------------------------------------------------------------------------
func AddOrdersToInvestmentAccount ( investmentAccountId uint64, ordersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InvestmentAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentAccount(investmentAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentAccount)

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
		// retrieve the modified InvestmentAccount from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentAccount(investmentAccountId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more ordersIds as a Orders from a InvestmentAccount
//----------------------------------------------------------------------------
func RemoveOrdersFromInvestmentAccount( investmentAccountId uint64, ordersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InvestmentAccount with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInvestmentAccount(investmentAccountId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InvestmentAccount so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InvestmentAccount)

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
		// retrieve the modified InvestmentAccount from the gorm
		//----------------------------------------------------------------------------
		return GetInvestmentAccount(investmentAccountId)

	} else {
		return parentRequestResult
	}
}

