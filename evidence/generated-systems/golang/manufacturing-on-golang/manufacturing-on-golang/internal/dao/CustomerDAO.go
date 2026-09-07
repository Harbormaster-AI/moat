package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CustomerDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCustomer - creates a new db entry
//----------------------------------------------------------------------------
func CreateCustomer(obj model.Customer)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Customer with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Customer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCustomer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCustomer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCustomer(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Customer

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Customer with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Customer using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Customer using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCustomer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCustomer - returns all
//----------------------------------------------------------------------------
func GetAllCustomer()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Customer

	//----------------------------------------------------------------------------
	// Request the ORM to find all Customer
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Customer" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Customer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCustomer", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCustomer - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCustomer(obj model.Customer)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Customer using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Customer using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCustomer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCustomer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCustomer(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCustomer(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Customer)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Customer using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Customer using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCustomer", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more enterprisesIds as a Enterprises to a Customer
//----------------------------------------------------------------------------
func AddEnterprisesToCustomer ( customerId uint64, enterprisesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( enterprisesIds, ",")

		for _, enterprisesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Enterprise

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Enterprise
			// with a matching enterprisesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , enterprisesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Enterprises using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Enterprises").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Enterprises", enterprisesId )
				return utils.RequestResult{false, msg, "unassignEnterprises", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more enterprisesIds as a Enterprises from a Customer
//----------------------------------------------------------------------------
func RemoveEnterprisesFromCustomer( customerId uint64, enterprisesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( enterprisesIds, ",")

		for _, enterprisesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Enterprise

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Enterprise
			// with a matching enterprisesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , enterprisesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EnterpriseObj from the Enterprises array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Enterprises").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Enterprises", enterprisesId )
				return utils.RequestResult{false, msg, "removeEnterprises", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more salesOrdersIds as a SalesOrders to a Customer
//----------------------------------------------------------------------------
func AddSalesOrdersToCustomer ( customerId uint64, salesOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( salesOrdersIds, ",")

		for _, salesOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SalesOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SalesOrder
			// with a matching salesOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , salesOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SalesOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SalesOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalesOrders", salesOrdersId )
				return utils.RequestResult{false, msg, "unassignSalesOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more salesOrdersIds as a SalesOrders from a Customer
//----------------------------------------------------------------------------
func RemoveSalesOrdersFromCustomer( customerId uint64, salesOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Customer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCustomer(customerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Customer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Customer)

		// slice the ids on comma with no spaces
		ids := strings.Split( salesOrdersIds, ",")

		for _, salesOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SalesOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SalesOrder
			// with a matching salesOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , salesOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SalesOrderObj from the SalesOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SalesOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalesOrders", salesOrdersId )
				return utils.RequestResult{false, msg, "removeSalesOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Customer from the gorm
		//----------------------------------------------------------------------------
		return GetCustomer(customerId)

	} else {
		return parentRequestResult
	}
}

