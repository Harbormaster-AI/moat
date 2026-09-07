package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SalesOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSalesOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreateSalesOrder(obj model.SalesOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a SalesOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a SalesOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSalesOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSalesOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSalesOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.SalesOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a SalesOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a SalesOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a SalesOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSalesOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSalesOrder - returns all
//----------------------------------------------------------------------------
func GetAllSalesOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.SalesOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all SalesOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all SalesOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all SalesOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSalesOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSalesOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSalesOrder(obj model.SalesOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a SalesOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a SalesOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSalesOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSalesOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSalesOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the SalesOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSalesOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.SalesOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a SalesOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a SalesOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSalesOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Customer on a SalesOrder
//----------------------------------------------------------------------------
func AssignCustomerToSalesOrder( salesOrderId uint64, customerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SalesOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrder(salesOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrder)

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
			// assign the Customer	to the SalesOrder
			//----------------------------------------------------------------------------
			parentObj.Customer = &childObj

			//----------------------------------------------------------------------------
			// save the SalesOrder
			//----------------------------------------------------------------------------
			return UpdateSalesOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customer", customerId )
			return utils.RequestResult{false, msg, "assignCustomer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Customer on a SalesOrder
//----------------------------------------------------------------------------
func UnassignCustomerFromSalesOrder(salesOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrder(salesOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrder)

		//----------------------------------------------------------------------------
		// assign an empty Customer to the Customer
		//----------------------------------------------------------------------------
		parentObj.Customer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Customer
		//----------------------------------------------------------------------------
		parentObj.CustomerId = nil;

		//----------------------------------------------------------------------------
		// save the SalesOrder
		//----------------------------------------------------------------------------
		return UpdateSalesOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Plant on a SalesOrder
//----------------------------------------------------------------------------
func AssignPlantToSalesOrder( salesOrderId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the SalesOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrder(salesOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Plant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Plant with a
		// matching plantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, plantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Plant	to the SalesOrder
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the SalesOrder
			//----------------------------------------------------------------------------
			return UpdateSalesOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a SalesOrder
//----------------------------------------------------------------------------
func UnassignPlantFromSalesOrder(salesOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrder(salesOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrder)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the SalesOrder
		//----------------------------------------------------------------------------
		return UpdateSalesOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more linesIds as a Lines to a SalesOrder
//----------------------------------------------------------------------------
func AddLinesToSalesOrder ( salesOrderId uint64, linesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrder(salesOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SalesOrderLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SalesOrderLine
			// with a matching linesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , linesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Lines using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lines").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lines", linesId )
				return utils.RequestResult{false, msg, "unassignLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesOrder from the gorm
		//----------------------------------------------------------------------------
		return GetSalesOrder(salesOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more linesIds as a Lines from a SalesOrder
//----------------------------------------------------------------------------
func RemoveLinesFromSalesOrder( salesOrderId uint64, linesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SalesOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrder(salesOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( linesIds, ",")

		for _, linesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SalesOrderLine

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SalesOrderLine
			// with a matching linesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , linesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SalesOrderLineObj from the Lines array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Lines").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Lines", linesId )
				return utils.RequestResult{false, msg, "removeLines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesOrder from the gorm
		//----------------------------------------------------------------------------
		return GetSalesOrder(salesOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more workOrdersIds as a WorkOrders to a SalesOrder
//----------------------------------------------------------------------------
func AddWorkOrdersToSalesOrder ( salesOrderId uint64, workOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the SalesOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrder(salesOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( workOrdersIds, ",")

		for _, workOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkOrder
			// with a matching workOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the WorkOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrders", workOrdersId )
				return utils.RequestResult{false, msg, "unassignWorkOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesOrder from the gorm
		//----------------------------------------------------------------------------
		return GetSalesOrder(salesOrderId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more workOrdersIds as a WorkOrders from a SalesOrder
//----------------------------------------------------------------------------
func RemoveWorkOrdersFromSalesOrder( salesOrderId uint64, workOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the SalesOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSalesOrder(salesOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.SalesOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.SalesOrder)

		// slice the ids on comma with no spaces
		ids := strings.Split( workOrdersIds, ",")

		for _, workOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.WorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a WorkOrder
			// with a matching workOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove WorkOrderObj from the WorkOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrders", workOrdersId )
				return utils.RequestResult{false, msg, "removeWorkOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified SalesOrder from the gorm
		//----------------------------------------------------------------------------
		return GetSalesOrder(salesOrderId)

	} else {
		return parentRequestResult
	}
}

