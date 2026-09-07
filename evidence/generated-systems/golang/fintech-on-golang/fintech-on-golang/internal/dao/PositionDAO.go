package dao

import (
    "fintech-on-golang/internal/model"
    "fintech-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PositionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePosition - creates a new db entry
//----------------------------------------------------------------------------
func CreatePosition(obj model.Position)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Position with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Position", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePosition", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPosition - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPosition(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Position

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Position with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Position using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Position using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPosition", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPosition - returns all
//----------------------------------------------------------------------------
func GetAllPosition()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Position

	//----------------------------------------------------------------------------
	// Request the ORM to find all Position
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Position" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Position", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPosition", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePosition - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePosition(obj model.Position)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Position using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Position using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePosition", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePosition - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePosition(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPosition(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Position using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Position using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePosition", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Portfolio on a Position
//----------------------------------------------------------------------------
func AssignPortfolioToPosition( positionId uint64, portfolioId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

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
			// assign the Portfolio	to the Position
			//----------------------------------------------------------------------------
			parentObj.Portfolio = &childObj

			//----------------------------------------------------------------------------
			// save the Position
			//----------------------------------------------------------------------------
			return UpdatePosition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Portfolio", portfolioId )
			return utils.RequestResult{false, msg, "assignPortfolio", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Portfolio on a Position
//----------------------------------------------------------------------------
func UnassignPortfolioFromPosition(positionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// assign an empty InvestmentPortfolio to the Portfolio
		//----------------------------------------------------------------------------
		parentObj.Portfolio = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Portfolio
		//----------------------------------------------------------------------------
		parentObj.PortfolioId = nil;

		//----------------------------------------------------------------------------
		// save the Position
		//----------------------------------------------------------------------------
		return UpdatePosition(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Security on a Position
//----------------------------------------------------------------------------
func AssignSecurityToPosition( positionId uint64, securityId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

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
			// assign the Security	to the Position
			//----------------------------------------------------------------------------
			parentObj.Security = &childObj

			//----------------------------------------------------------------------------
			// save the Position
			//----------------------------------------------------------------------------
			return UpdatePosition(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Security", securityId )
			return utils.RequestResult{false, msg, "assignSecurity", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Security on a Position
//----------------------------------------------------------------------------
func UnassignSecurityFromPosition(positionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Position with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPosition(positionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Position so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Position)

		//----------------------------------------------------------------------------
		// assign an empty Security to the Security
		//----------------------------------------------------------------------------
		parentObj.Security = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Security
		//----------------------------------------------------------------------------
		parentObj.SecurityId = nil;

		//----------------------------------------------------------------------------
		// save the Position
		//----------------------------------------------------------------------------
		return UpdatePosition(parentObj)

	} else {
		return parentRequestResult
	}

}


