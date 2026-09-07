package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SupplierDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSupplier - creates a new db entry
//----------------------------------------------------------------------------
func CreateSupplier(obj model.Supplier)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Supplier with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Supplier", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSupplier", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSupplier - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSupplier(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Supplier

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Supplier with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Supplier using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Supplier using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSupplier", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSupplier - returns all
//----------------------------------------------------------------------------
func GetAllSupplier()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Supplier

	//----------------------------------------------------------------------------
	// Request the ORM to find all Supplier
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Supplier" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Supplier", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSupplier", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSupplier - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSupplier(obj model.Supplier)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Supplier using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Supplier using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSupplier", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSupplier - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSupplier(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Supplier with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSupplier(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Supplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Supplier)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Supplier using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Supplier using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSupplier", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more enterprisesIds as a Enterprises to a Supplier
//----------------------------------------------------------------------------
func AddEnterprisesToSupplier ( supplierId uint64, enterprisesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Supplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSupplier(supplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Supplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Supplier)

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
		// retrieve the modified Supplier from the gorm
		//----------------------------------------------------------------------------
		return GetSupplier(supplierId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more enterprisesIds as a Enterprises from a Supplier
//----------------------------------------------------------------------------
func RemoveEnterprisesFromSupplier( supplierId uint64, enterprisesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Supplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSupplier(supplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Supplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Supplier)

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
		// retrieve the modified Supplier from the gorm
		//----------------------------------------------------------------------------
		return GetSupplier(supplierId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more itemsIds as a Items to a Supplier
//----------------------------------------------------------------------------
func AddItemsToSupplier ( supplierId uint64, itemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Supplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSupplier(supplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Supplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Supplier)

		// slice the ids on comma with no spaces
		ids := strings.Split( itemsIds, ",")

		for _, itemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Item

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Item
			// with a matching itemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , itemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Items using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Items").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Items", itemsId )
				return utils.RequestResult{false, msg, "unassignItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Supplier from the gorm
		//----------------------------------------------------------------------------
		return GetSupplier(supplierId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more itemsIds as a Items from a Supplier
//----------------------------------------------------------------------------
func RemoveItemsFromSupplier( supplierId uint64, itemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Supplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSupplier(supplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Supplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Supplier)

		// slice the ids on comma with no spaces
		ids := strings.Split( itemsIds, ",")

		for _, itemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Item

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Item
			// with a matching itemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , itemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ItemObj from the Items array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Items").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Items", itemsId )
				return utils.RequestResult{false, msg, "removeItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Supplier from the gorm
		//----------------------------------------------------------------------------
		return GetSupplier(supplierId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more purchaseOrdersIds as a PurchaseOrders to a Supplier
//----------------------------------------------------------------------------
func AddPurchaseOrdersToSupplier ( supplierId uint64, purchaseOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Supplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSupplier(supplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Supplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Supplier)

		// slice the ids on comma with no spaces
		ids := strings.Split( purchaseOrdersIds, ",")

		for _, purchaseOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PurchaseOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PurchaseOrder
			// with a matching purchaseOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , purchaseOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PurchaseOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PurchaseOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PurchaseOrders", purchaseOrdersId )
				return utils.RequestResult{false, msg, "unassignPurchaseOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Supplier from the gorm
		//----------------------------------------------------------------------------
		return GetSupplier(supplierId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more purchaseOrdersIds as a PurchaseOrders from a Supplier
//----------------------------------------------------------------------------
func RemovePurchaseOrdersFromSupplier( supplierId uint64, purchaseOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Supplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSupplier(supplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Supplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Supplier)

		// slice the ids on comma with no spaces
		ids := strings.Split( purchaseOrdersIds, ",")

		for _, purchaseOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PurchaseOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PurchaseOrder
			// with a matching purchaseOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , purchaseOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PurchaseOrderObj from the PurchaseOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PurchaseOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PurchaseOrders", purchaseOrdersId )
				return utils.RequestResult{false, msg, "removePurchaseOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Supplier from the gorm
		//----------------------------------------------------------------------------
		return GetSupplier(supplierId)

	} else {
		return parentRequestResult
	}
}

