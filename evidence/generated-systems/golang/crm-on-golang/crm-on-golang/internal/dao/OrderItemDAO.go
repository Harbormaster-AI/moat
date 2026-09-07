package dao

import (
    "crm-on-golang/internal/model"
    "crm-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OrderItemDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOrderItem - creates a new db entry
//----------------------------------------------------------------------------
func CreateOrderItem(obj model.OrderItem)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a OrderItem with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a OrderItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOrderItem", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOrderItem - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOrderItem(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.OrderItem

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a OrderItem with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a OrderItem using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a OrderItem using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOrderItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOrderItem - returns all
//----------------------------------------------------------------------------
func GetAllOrderItem()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.OrderItem

	//----------------------------------------------------------------------------
	// Request the ORM to find all OrderItem
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all OrderItem" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all OrderItem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOrderItem", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOrderItem - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOrderItem(obj model.OrderItem)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a OrderItem using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a OrderItem using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOrderItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOrderItem - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOrderItem(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the OrderItem with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOrderItem(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OrderItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.OrderItem)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a OrderItem using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a OrderItem using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOrderItem", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Order on a OrderItem
//----------------------------------------------------------------------------
func AssignOrderToOrderItem( orderItemId uint64, orderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OrderItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrderItem(orderItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OrderItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OrderItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Order

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Order with a
		// matching orderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, orderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Order	to the OrderItem
			//----------------------------------------------------------------------------
			parentObj.Order = &childObj

			//----------------------------------------------------------------------------
			// save the OrderItem
			//----------------------------------------------------------------------------
			return UpdateOrderItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Order", orderId )
			return utils.RequestResult{false, msg, "assignOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Order on a OrderItem
//----------------------------------------------------------------------------
func UnassignOrderFromOrderItem(orderItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OrderItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrderItem(orderItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OrderItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OrderItem)

		//----------------------------------------------------------------------------
		// assign an empty Order to the Order
		//----------------------------------------------------------------------------
		parentObj.Order = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Order
		//----------------------------------------------------------------------------
		parentObj.OrderId = nil;

		//----------------------------------------------------------------------------
		// save the OrderItem
		//----------------------------------------------------------------------------
		return UpdateOrderItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Product on a OrderItem
//----------------------------------------------------------------------------
func AssignProductToOrderItem( orderItemId uint64, productId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OrderItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrderItem(orderItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OrderItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OrderItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Product

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Product with a
		// matching productId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, productId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Product	to the OrderItem
			//----------------------------------------------------------------------------
			parentObj.Product = &childObj

			//----------------------------------------------------------------------------
			// save the OrderItem
			//----------------------------------------------------------------------------
			return UpdateOrderItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Product", productId )
			return utils.RequestResult{false, msg, "assignProduct", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Product on a OrderItem
//----------------------------------------------------------------------------
func UnassignProductFromOrderItem(orderItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OrderItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrderItem(orderItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OrderItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OrderItem)

		//----------------------------------------------------------------------------
		// assign an empty Product to the Product
		//----------------------------------------------------------------------------
		parentObj.Product = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Product
		//----------------------------------------------------------------------------
		parentObj.ProductId = nil;

		//----------------------------------------------------------------------------
		// save the OrderItem
		//----------------------------------------------------------------------------
		return UpdateOrderItem(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PriceBookEntry on a OrderItem
//----------------------------------------------------------------------------
func AssignPriceBookEntryToOrderItem( orderItemId uint64, priceBookEntryId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the OrderItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrderItem(orderItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OrderItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OrderItem)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PriceBookEntry

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PriceBookEntry with a
		// matching priceBookEntryId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, priceBookEntryId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PriceBookEntry	to the OrderItem
			//----------------------------------------------------------------------------
			parentObj.PriceBookEntry = &childObj

			//----------------------------------------------------------------------------
			// save the OrderItem
			//----------------------------------------------------------------------------
			return UpdateOrderItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PriceBookEntry", priceBookEntryId )
			return utils.RequestResult{false, msg, "assignPriceBookEntry", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PriceBookEntry on a OrderItem
//----------------------------------------------------------------------------
func UnassignPriceBookEntryFromOrderItem(orderItemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the OrderItem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrderItem(orderItemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.OrderItem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.OrderItem)

		//----------------------------------------------------------------------------
		// assign an empty PriceBookEntry to the PriceBookEntry
		//----------------------------------------------------------------------------
		parentObj.PriceBookEntry = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PriceBookEntry
		//----------------------------------------------------------------------------
		parentObj.PriceBookEntryId = nil;

		//----------------------------------------------------------------------------
		// save the OrderItem
		//----------------------------------------------------------------------------
		return UpdateOrderItem(parentObj)

	} else {
		return parentRequestResult
	}

}


