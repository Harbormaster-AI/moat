package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TradeOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTradeOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateTradeOrder(obj model.TradeOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TradeOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TradeOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTradeOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTradeOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTradeOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TradeOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TradeOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TradeOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TradeOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTradeOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTradeOrder - returns all
//----------------------------------------------------------------------------
func GetAllTradeOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TradeOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all TradeOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TradeOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TradeOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTradeOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTradeOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTradeOrder(obj model.TradeOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TradeOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TradeOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTradeOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTradeOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTradeOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TradeOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTradeOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TradeOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TradeOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TradeOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TradeOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTradeOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Portfolio on a TradeOrder
//----------------------------------------------------------------------------
func AssignPortfolioToTradeOrder( tradeOrderId uint64, portfolioId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TradeOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTradeOrder(tradeOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TradeOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TradeOrder)

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
			// assign the Portfolio	to the TradeOrder
			//----------------------------------------------------------------------------
			parentObj.Portfolio = &childObj

			//----------------------------------------------------------------------------
			// save the TradeOrder
			//----------------------------------------------------------------------------
			return UpdateTradeOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Portfolio", portfolioId )
			return utils.RequestResult{false, msg, "assignPortfolio", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Portfolio on a TradeOrder
//----------------------------------------------------------------------------
func UnassignPortfolioFromTradeOrder(tradeOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TradeOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTradeOrder(tradeOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TradeOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TradeOrder)

		//----------------------------------------------------------------------------
		// assign an empty InvestmentPortfolio to the Portfolio
		//----------------------------------------------------------------------------
		parentObj.Portfolio = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Portfolio
		//----------------------------------------------------------------------------
		parentObj.PortfolioId = nil;

		//----------------------------------------------------------------------------
		// save the TradeOrder
		//----------------------------------------------------------------------------
		return UpdateTradeOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Security on a TradeOrder
//----------------------------------------------------------------------------
func AssignSecurityToTradeOrder( tradeOrderId uint64, securityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TradeOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTradeOrder(tradeOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TradeOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TradeOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Security

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Security with a
		// matching securityId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, securityId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Security	to the TradeOrder
			//----------------------------------------------------------------------------
			parentObj.Security = &childObj

			//----------------------------------------------------------------------------
			// save the TradeOrder
			//----------------------------------------------------------------------------
			return UpdateTradeOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Security", securityId )
			return utils.RequestResult{false, msg, "assignSecurity", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Security on a TradeOrder
//----------------------------------------------------------------------------
func UnassignSecurityFromTradeOrder(tradeOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TradeOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTradeOrder(tradeOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TradeOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TradeOrder)

		//----------------------------------------------------------------------------
		// assign an empty Security to the Security
		//----------------------------------------------------------------------------
		parentObj.Security = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Security
		//----------------------------------------------------------------------------
		parentObj.SecurityId = nil;

		//----------------------------------------------------------------------------
		// save the TradeOrder
		//----------------------------------------------------------------------------
		return UpdateTradeOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more tradesIds as a Trades to a TradeOrder
//----------------------------------------------------------------------------
func AddTradesToTradeOrder ( tradeOrderId uint64, tradesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TradeOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTradeOrder(tradeOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TradeOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TradeOrder)

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
		// retrieve the modified TradeOrder from the gorm
		//----------------------------------------------------------------------------
		return GetTradeOrder(tradeOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more tradesIds as a Trades from a TradeOrder
//----------------------------------------------------------------------------
func RemoveTradesFromTradeOrder( tradeOrderId uint64, tradesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TradeOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTradeOrder(tradeOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TradeOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TradeOrder)

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
		// retrieve the modified TradeOrder from the gorm
		//----------------------------------------------------------------------------
		return GetTradeOrder(tradeOrderId)

	} else {
		return parentRequestResult
	}
}

