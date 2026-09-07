package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing HealthSystemDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateHealthSystem - creates a new db entry
//----------------------------------------------------------------------------
func CreateHealthSystem(obj model.HealthSystem)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a HealthSystem with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a HealthSystem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateHealthSystem", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetHealthSystem - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetHealthSystem(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.HealthSystem

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a HealthSystem with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a HealthSystem using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a HealthSystem using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetHealthSystem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllHealthSystem - returns all
//----------------------------------------------------------------------------
func GetAllHealthSystem()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.HealthSystem

	//----------------------------------------------------------------------------
	// Request the ORM to find all HealthSystem
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all HealthSystem" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all HealthSystem", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllHealthSystem", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateHealthSystem - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateHealthSystem(obj model.HealthSystem)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a HealthSystem using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a HealthSystem using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateHealthSystem", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteHealthSystem - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteHealthSystem(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the HealthSystem with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetHealthSystem(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.HealthSystem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.HealthSystem)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a HealthSystem using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a HealthSystem using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteHealthSystem", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more facilitiesIds as a Facilities to a HealthSystem
//----------------------------------------------------------------------------
func AddFacilitiesToHealthSystem ( healthSystemId uint64, facilitiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the HealthSystem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetHealthSystem(healthSystemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.HealthSystem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.HealthSystem)

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
		// retrieve the modified HealthSystem from the gorm
		//----------------------------------------------------------------------------
		return GetHealthSystem(healthSystemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more facilitiesIds as a Facilities from a HealthSystem
//----------------------------------------------------------------------------
func RemoveFacilitiesFromHealthSystem( healthSystemId uint64, facilitiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the HealthSystem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetHealthSystem(healthSystemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.HealthSystem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.HealthSystem)

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
		// retrieve the modified HealthSystem from the gorm
		//----------------------------------------------------------------------------
		return GetHealthSystem(healthSystemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more suppliersIds as a Suppliers to a HealthSystem
//----------------------------------------------------------------------------
func AddSuppliersToHealthSystem ( healthSystemId uint64, suppliersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the HealthSystem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetHealthSystem(healthSystemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.HealthSystem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.HealthSystem)

		// slice the ids on comma with no spaces
		ids := strings.Split( suppliersIds, ",")

		for _, suppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicalSupplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicalSupplier
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
		// retrieve the modified HealthSystem from the gorm
		//----------------------------------------------------------------------------
		return GetHealthSystem(healthSystemId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more suppliersIds as a Suppliers from a HealthSystem
//----------------------------------------------------------------------------
func RemoveSuppliersFromHealthSystem( healthSystemId uint64, suppliersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the HealthSystem with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetHealthSystem(healthSystemId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.HealthSystem so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.HealthSystem)

		// slice the ids on comma with no spaces
		ids := strings.Split( suppliersIds, ",")

		for _, suppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MedicalSupplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MedicalSupplier
			// with a matching suppliersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , suppliersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MedicalSupplierObj from the Suppliers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Suppliers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Suppliers", suppliersId )
				return utils.RequestResult{false, msg, "removeSuppliers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified HealthSystem from the gorm
		//----------------------------------------------------------------------------
		return GetHealthSystem(healthSystemId)

	} else {
		return parentRequestResult
	}
}

