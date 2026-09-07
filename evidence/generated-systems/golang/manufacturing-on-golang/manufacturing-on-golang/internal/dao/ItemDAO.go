package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ItemDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateItem - creates a new db entry
//----------------------------------------------------------------------------
func CreateItem(obj model.Item)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Item with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Item", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateItem", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetItem - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetItem(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Item

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Item with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Item using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Item using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllItem - returns all
//----------------------------------------------------------------------------
func GetAllItem()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Item

	//----------------------------------------------------------------------------
	// Request the ORM to find all Item
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Item" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Item", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllItem", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateItem - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateItem(obj model.Item)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Item using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Item using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateItem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteItem - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteItem(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetItem(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Item)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Item using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Item using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteItem", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a BusinessUnit on a Item
//----------------------------------------------------------------------------
func AssignBusinessUnitToItem( itemId uint64, businessUnitId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BusinessUnit

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BusinessUnit with a
		// matching businessUnitId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, businessUnitId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the BusinessUnit	to the Item
			//----------------------------------------------------------------------------
			parentObj.BusinessUnit = &childObj

			//----------------------------------------------------------------------------
			// save the Item
			//----------------------------------------------------------------------------
			return UpdateItem(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BusinessUnit", businessUnitId )
			return utils.RequestResult{false, msg, "assignBusinessUnit", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a BusinessUnit on a Item
//----------------------------------------------------------------------------
func UnassignBusinessUnitFromItem(itemId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		//----------------------------------------------------------------------------
		// assign an empty BusinessUnit to the BusinessUnit
		//----------------------------------------------------------------------------
		parentObj.BusinessUnit = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the BusinessUnit
		//----------------------------------------------------------------------------
		parentObj.BusinessUnitId = nil;

		//----------------------------------------------------------------------------
		// save the Item
		//----------------------------------------------------------------------------
		return UpdateItem(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more bomsIds as a Boms to a Item
//----------------------------------------------------------------------------
func AddBomsToItem ( itemId uint64, bomsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( bomsIds, ",")

		for _, bomsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BOM

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BOM
			// with a matching bomsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , bomsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Boms using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Boms").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Boms", bomsId )
				return utils.RequestResult{false, msg, "unassignBoms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more bomsIds as a Boms from a Item
//----------------------------------------------------------------------------
func RemoveBomsFromItem( itemId uint64, bomsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( bomsIds, ",")

		for _, bomsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BOM

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BOM
			// with a matching bomsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , bomsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BOMObj from the Boms array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Boms").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Boms", bomsId )
				return utils.RequestResult{false, msg, "removeBoms", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more routingsIds as a Routings to a Item
//----------------------------------------------------------------------------
func AddRoutingsToItem ( itemId uint64, routingsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( routingsIds, ",")

		for _, routingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Routing

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Routing
			// with a matching routingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , routingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Routings using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Routings").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Routings", routingsId )
				return utils.RequestResult{false, msg, "unassignRoutings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more routingsIds as a Routings from a Item
//----------------------------------------------------------------------------
func RemoveRoutingsFromItem( itemId uint64, routingsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( routingsIds, ",")

		for _, routingsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Routing

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Routing
			// with a matching routingsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , routingsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RoutingObj from the Routings array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Routings").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Routings", routingsId )
				return utils.RequestResult{false, msg, "removeRoutings", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more suppliersIds as a Suppliers to a Item
//----------------------------------------------------------------------------
func AddSuppliersToItem ( itemId uint64, suppliersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( suppliersIds, ",")

		for _, suppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Supplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Supplier
			// with a matching suppliersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , suppliersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Suppliers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Suppliers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Suppliers", suppliersId )
				return utils.RequestResult{false, msg, "unassignSuppliers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more suppliersIds as a Suppliers from a Item
//----------------------------------------------------------------------------
func RemoveSuppliersFromItem( itemId uint64, suppliersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( suppliersIds, ",")

		for _, suppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Supplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Supplier
			// with a matching suppliersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , suppliersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SupplierObj from the Suppliers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Suppliers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Suppliers", suppliersId )
				return utils.RequestResult{false, msg, "removeSuppliers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more qualitySpecificationsIds as a QualitySpecifications to a Item
//----------------------------------------------------------------------------
func AddQualitySpecificationsToItem ( itemId uint64, qualitySpecificationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( qualitySpecificationsIds, ",")

		for _, qualitySpecificationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QualitySpecification

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QualitySpecification
			// with a matching qualitySpecificationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , qualitySpecificationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the QualitySpecifications using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("QualitySpecifications").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "QualitySpecifications", qualitySpecificationsId )
				return utils.RequestResult{false, msg, "unassignQualitySpecifications", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more qualitySpecificationsIds as a QualitySpecifications from a Item
//----------------------------------------------------------------------------
func RemoveQualitySpecificationsFromItem( itemId uint64, qualitySpecificationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( qualitySpecificationsIds, ",")

		for _, qualitySpecificationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.QualitySpecification

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a QualitySpecification
			// with a matching qualitySpecificationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , qualitySpecificationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove QualitySpecificationObj from the QualitySpecifications array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("QualitySpecifications").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "QualitySpecifications", qualitySpecificationsId )
				return utils.RequestResult{false, msg, "removeQualitySpecifications", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more inventoryItemsIds as a InventoryItems to a Item
//----------------------------------------------------------------------------
func AddInventoryItemsToItem ( itemId uint64, inventoryItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventoryItemsIds, ",")

		for _, inventoryItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching inventoryItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventoryItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InventoryItems using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventoryItems").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItems", inventoryItemsId )
				return utils.RequestResult{false, msg, "unassignInventoryItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventoryItemsIds as a InventoryItems from a Item
//----------------------------------------------------------------------------
func RemoveInventoryItemsFromItem( itemId uint64, inventoryItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Item with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetItem(itemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Item so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Item)

		// slice the ids on comma with no spaces
		ids := strings.Split( inventoryItemsIds, ",")

		for _, inventoryItemsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InventoryItem

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InventoryItem
			// with a matching inventoryItemsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inventoryItemsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InventoryItemObj from the InventoryItems array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InventoryItems").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InventoryItems", inventoryItemsId )
				return utils.RequestResult{false, msg, "removeInventoryItems", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Item from the gorm
		//----------------------------------------------------------------------------
		return GetItem(itemId)

	} else {
		return parentRequestResult
	}
}

