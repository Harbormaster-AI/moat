package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FacilityDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFacility - creates a new db entry
//----------------------------------------------------------------------------
func CreateFacility(obj model.Facility)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Facility with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Facility", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFacility", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFacility - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFacility(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Facility

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Facility with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Facility using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Facility using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFacility", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFacility - returns all
//----------------------------------------------------------------------------
func GetAllFacility()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Facility

	//----------------------------------------------------------------------------
	// Request the ORM to find all Facility
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Facility" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Facility", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFacility", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFacility - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFacility(obj model.Facility)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Facility using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Facility using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFacility", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFacility - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFacility(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFacility(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Facility)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Facility using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Facility using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFacility", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a HealthSystem on a Facility
//----------------------------------------------------------------------------
func AssignHealthSystemToFacility( facilityId uint64, healthSystemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.HealthSystem

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a HealthSystem with a
		// matching healthSystemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, healthSystemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the HealthSystem	to the Facility
			//----------------------------------------------------------------------------
			parentObj.HealthSystem = &childObj

			//----------------------------------------------------------------------------
			// save the Facility
			//----------------------------------------------------------------------------
			return UpdateFacility(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "HealthSystem", healthSystemId )
			return utils.RequestResult{false, msg, "assignHealthSystem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a HealthSystem on a Facility
//----------------------------------------------------------------------------
func UnassignHealthSystemFromFacility(facilityId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		//----------------------------------------------------------------------------
		// assign an empty HealthSystem to the HealthSystem
		//----------------------------------------------------------------------------
		parentObj.HealthSystem = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the HealthSystem
		//----------------------------------------------------------------------------
		parentObj.HealthSystemId = nil;

		//----------------------------------------------------------------------------
		// save the Facility
		//----------------------------------------------------------------------------
		return UpdateFacility(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more departmentsIds as a Departments to a Facility
//----------------------------------------------------------------------------
func AddDepartmentsToFacility ( facilityId uint64, departmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( departmentsIds, ",")

		for _, departmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Department

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Department
			// with a matching departmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , departmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Departments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Departments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Departments", departmentsId )
				return utils.RequestResult{false, msg, "unassignDepartments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more departmentsIds as a Departments from a Facility
//----------------------------------------------------------------------------
func RemoveDepartmentsFromFacility( facilityId uint64, departmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( departmentsIds, ",")

		for _, departmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Department

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Department
			// with a matching departmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , departmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DepartmentObj from the Departments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Departments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Departments", departmentsId )
				return utils.RequestResult{false, msg, "removeDepartments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more careTeamsIds as a CareTeams to a Facility
//----------------------------------------------------------------------------
func AddCareTeamsToFacility ( facilityId uint64, careTeamsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( careTeamsIds, ",")

		for _, careTeamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CareTeam

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CareTeam
			// with a matching careTeamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , careTeamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CareTeams using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CareTeams").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CareTeams", careTeamsId )
				return utils.RequestResult{false, msg, "unassignCareTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more careTeamsIds as a CareTeams from a Facility
//----------------------------------------------------------------------------
func RemoveCareTeamsFromFacility( facilityId uint64, careTeamsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( careTeamsIds, ",")

		for _, careTeamsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CareTeam

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CareTeam
			// with a matching careTeamsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , careTeamsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CareTeamObj from the CareTeams array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CareTeams").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CareTeams", careTeamsId )
				return utils.RequestResult{false, msg, "removeCareTeams", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more laboratoriesIds as a Laboratories to a Facility
//----------------------------------------------------------------------------
func AddLaboratoriesToFacility ( facilityId uint64, laboratoriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( laboratoriesIds, ",")

		for _, laboratoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Laboratory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Laboratory
			// with a matching laboratoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , laboratoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Laboratories using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Laboratories").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Laboratories", laboratoriesId )
				return utils.RequestResult{false, msg, "unassignLaboratories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more laboratoriesIds as a Laboratories from a Facility
//----------------------------------------------------------------------------
func RemoveLaboratoriesFromFacility( facilityId uint64, laboratoriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( laboratoriesIds, ",")

		for _, laboratoriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Laboratory

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Laboratory
			// with a matching laboratoriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , laboratoriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LaboratoryObj from the Laboratories array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Laboratories").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Laboratories", laboratoriesId )
				return utils.RequestResult{false, msg, "removeLaboratories", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more imagingCentersIds as a ImagingCenters to a Facility
//----------------------------------------------------------------------------
func AddImagingCentersToFacility ( facilityId uint64, imagingCentersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingCentersIds, ",")

		for _, imagingCentersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingCenter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingCenter
			// with a matching imagingCentersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingCentersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ImagingCenters using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingCenters").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingCenters", imagingCentersId )
				return utils.RequestResult{false, msg, "unassignImagingCenters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more imagingCentersIds as a ImagingCenters from a Facility
//----------------------------------------------------------------------------
func RemoveImagingCentersFromFacility( facilityId uint64, imagingCentersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( imagingCentersIds, ",")

		for _, imagingCentersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ImagingCenter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ImagingCenter
			// with a matching imagingCentersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , imagingCentersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ImagingCenterObj from the ImagingCenters array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ImagingCenters").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ImagingCenters", imagingCentersId )
				return utils.RequestResult{false, msg, "removeImagingCenters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more pharmaciesIds as a Pharmacies to a Facility
//----------------------------------------------------------------------------
func AddPharmaciesToFacility ( facilityId uint64, pharmaciesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( pharmaciesIds, ",")

		for _, pharmaciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Pharmacy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Pharmacy
			// with a matching pharmaciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , pharmaciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Pharmacies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Pharmacies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Pharmacies", pharmaciesId )
				return utils.RequestResult{false, msg, "unassignPharmacies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more pharmaciesIds as a Pharmacies from a Facility
//----------------------------------------------------------------------------
func RemovePharmaciesFromFacility( facilityId uint64, pharmaciesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

		// slice the ids on comma with no spaces
		ids := strings.Split( pharmaciesIds, ",")

		for _, pharmaciesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Pharmacy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Pharmacy
			// with a matching pharmaciesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , pharmaciesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PharmacyObj from the Pharmacies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Pharmacies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Pharmacies", pharmaciesId )
				return utils.RequestResult{false, msg, "removePharmacies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more inventoryItemsIds as a InventoryItems to a Facility
//----------------------------------------------------------------------------
func AddInventoryItemsToFacility ( facilityId uint64, inventoryItemsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

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
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inventoryItemsIds as a InventoryItems from a Facility
//----------------------------------------------------------------------------
func RemoveInventoryItemsFromFacility( facilityId uint64, inventoryItemsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Facility with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFacility(facilityId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Facility so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Facility)

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
		// retrieve the modified Facility from the gorm
		//----------------------------------------------------------------------------
		return GetFacility(facilityId)

	} else {
		return parentRequestResult
	}
}

