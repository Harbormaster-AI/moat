package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing MedicalSupplierDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateMedicalSupplier - creates a new db entry
//----------------------------------------------------------------------------
func CreateMedicalSupplier(obj model.MedicalSupplier)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a MedicalSupplier with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a MedicalSupplier", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateMedicalSupplier", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetMedicalSupplier - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetMedicalSupplier(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.MedicalSupplier

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a MedicalSupplier with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a MedicalSupplier using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a MedicalSupplier using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetMedicalSupplier", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllMedicalSupplier - returns all
//----------------------------------------------------------------------------
func GetAllMedicalSupplier()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.MedicalSupplier

	//----------------------------------------------------------------------------
	// Request the ORM to find all MedicalSupplier
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all MedicalSupplier" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all MedicalSupplier", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllMedicalSupplier", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateMedicalSupplier - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateMedicalSupplier(obj model.MedicalSupplier)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a MedicalSupplier using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a MedicalSupplier using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateMedicalSupplier", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteMedicalSupplier - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteMedicalSupplier(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the MedicalSupplier with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetMedicalSupplier(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalSupplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.MedicalSupplier)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a MedicalSupplier using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a MedicalSupplier using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteMedicalSupplier", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more facilitiesIds as a Facilities to a MedicalSupplier
//----------------------------------------------------------------------------
func AddFacilitiesToMedicalSupplier ( medicalSupplierId uint64, facilitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicalSupplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalSupplier(medicalSupplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalSupplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalSupplier)

		// slice the ids on comma with no spaces
		ids := strings.Split( facilitiesIds, ",")

		for _, facilitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Facility

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Facility
			// with a matching facilitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , facilitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Facilities using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Facilities").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facilities", facilitiesId )
				return utils.RequestResult{false, msg, "unassignFacilities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MedicalSupplier from the gorm
		//----------------------------------------------------------------------------
		return GetMedicalSupplier(medicalSupplierId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more facilitiesIds as a Facilities from a MedicalSupplier
//----------------------------------------------------------------------------
func RemoveFacilitiesFromMedicalSupplier( medicalSupplierId uint64, facilitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the MedicalSupplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalSupplier(medicalSupplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalSupplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalSupplier)

		// slice the ids on comma with no spaces
		ids := strings.Split( facilitiesIds, ",")

		for _, facilitiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Facility

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Facility
			// with a matching facilitiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , facilitiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FacilityObj from the Facilities array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Facilities").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Facilities", facilitiesId )
				return utils.RequestResult{false, msg, "removeFacilities", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified MedicalSupplier from the gorm
		//----------------------------------------------------------------------------
		return GetMedicalSupplier(medicalSupplierId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more inventoryItemsIds as a InventoryItems to a MedicalSupplier
//----------------------------------------------------------------------------
func AddInventoryItemsToMedicalSupplier ( medicalSupplierId uint64, inventoryItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the MedicalSupplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalSupplier(medicalSupplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalSupplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalSupplier)

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
		// retrieve the modified MedicalSupplier from the gorm
		//----------------------------------------------------------------------------
		return GetMedicalSupplier(medicalSupplierId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventoryItemsIds as a InventoryItems from a MedicalSupplier
//----------------------------------------------------------------------------
func RemoveInventoryItemsFromMedicalSupplier( medicalSupplierId uint64, inventoryItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the MedicalSupplier with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetMedicalSupplier(medicalSupplierId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.MedicalSupplier so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.MedicalSupplier)

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
		// retrieve the modified MedicalSupplier from the gorm
		//----------------------------------------------------------------------------
		return GetMedicalSupplier(medicalSupplierId)

	} else {
		return parentRequestResult
	}
}

