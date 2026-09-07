package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ServiceBulletinDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateServiceBulletin - creates a new db entry
//----------------------------------------------------------------------------
func CreateServiceBulletin(obj model.ServiceBulletin)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ServiceBulletin with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ServiceBulletin", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateServiceBulletin", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetServiceBulletin - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetServiceBulletin(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ServiceBulletin

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ServiceBulletin with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ServiceBulletin using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ServiceBulletin using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetServiceBulletin", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllServiceBulletin - returns all
//----------------------------------------------------------------------------
func GetAllServiceBulletin()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ServiceBulletin

	//----------------------------------------------------------------------------
	// Request the ORM to find all ServiceBulletin
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ServiceBulletin" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ServiceBulletin", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllServiceBulletin", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateServiceBulletin - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateServiceBulletin(obj model.ServiceBulletin)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ServiceBulletin using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ServiceBulletin using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateServiceBulletin", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteServiceBulletin - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteServiceBulletin(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ServiceBulletin with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetServiceBulletin(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ServiceBulletin so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ServiceBulletin)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ServiceBulletin using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ServiceBulletin using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteServiceBulletin", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more workOrdersIds as a WorkOrders to a ServiceBulletin
//----------------------------------------------------------------------------
func AddWorkOrdersToServiceBulletin ( serviceBulletinId uint64, workOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ServiceBulletin with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetServiceBulletin(serviceBulletinId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ServiceBulletin so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ServiceBulletin)

		// slice the ids on comma with no spaces
		ids := strings.Split( workOrdersIds, ",")

		for _, workOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceWorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceWorkOrder
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
		// retrieve the modified ServiceBulletin from the gorm
		//----------------------------------------------------------------------------
		return GetServiceBulletin(serviceBulletinId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more workOrdersIds as a WorkOrders from a ServiceBulletin
//----------------------------------------------------------------------------
func RemoveWorkOrdersFromServiceBulletin( serviceBulletinId uint64, workOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ServiceBulletin with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetServiceBulletin(serviceBulletinId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ServiceBulletin so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ServiceBulletin)

		// slice the ids on comma with no spaces
		ids := strings.Split( workOrdersIds, ",")

		for _, workOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceWorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceWorkOrder
			// with a matching workOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , workOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MaintenanceWorkOrderObj from the WorkOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("WorkOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrders", workOrdersId )
				return utils.RequestResult{false, msg, "removeWorkOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ServiceBulletin from the gorm
		//----------------------------------------------------------------------------
		return GetServiceBulletin(serviceBulletinId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more variantsIds as a Variants to a ServiceBulletin
//----------------------------------------------------------------------------
func AddVariantsToServiceBulletin ( serviceBulletinId uint64, variantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ServiceBulletin with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetServiceBulletin(serviceBulletinId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ServiceBulletin so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ServiceBulletin)

		// slice the ids on comma with no spaces
		ids := strings.Split( variantsIds, ",")

		for _, variantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftVariant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftVariant
			// with a matching variantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Variants using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variants").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variants", variantsId )
				return utils.RequestResult{false, msg, "unassignVariants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ServiceBulletin from the gorm
		//----------------------------------------------------------------------------
		return GetServiceBulletin(serviceBulletinId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more variantsIds as a Variants from a ServiceBulletin
//----------------------------------------------------------------------------
func RemoveVariantsFromServiceBulletin( serviceBulletinId uint64, variantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ServiceBulletin with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetServiceBulletin(serviceBulletinId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ServiceBulletin so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ServiceBulletin)

		// slice the ids on comma with no spaces
		ids := strings.Split( variantsIds, ",")

		for _, variantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftVariant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftVariant
			// with a matching variantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftVariantObj from the Variants array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variants").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variants", variantsId )
				return utils.RequestResult{false, msg, "removeVariants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ServiceBulletin from the gorm
		//----------------------------------------------------------------------------
		return GetServiceBulletin(serviceBulletinId)

	} else {
		return parentRequestResult
	}
}

