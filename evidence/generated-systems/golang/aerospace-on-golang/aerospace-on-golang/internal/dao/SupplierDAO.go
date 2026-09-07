package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
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
// adds one or more manufacturersIds as a Manufacturers to a Supplier
//----------------------------------------------------------------------------
func AddManufacturersToSupplier ( supplierId uint64, manufacturersIds string )(utils.RequestResult) {

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
		ids := strings.Split( manufacturersIds, ",")

		for _, manufacturersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AerospaceManufacturer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AerospaceManufacturer
			// with a matching manufacturersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , manufacturersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Manufacturers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Manufacturers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Manufacturers", manufacturersId )
				return utils.RequestResult{false, msg, "unassignManufacturers", childObj}
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
// removes one or more manufacturersIds as a Manufacturers from a Supplier
//----------------------------------------------------------------------------
func RemoveManufacturersFromSupplier( supplierId uint64, manufacturersIds string )(utils.RequestResult) {
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
		ids := strings.Split( manufacturersIds, ",")

		for _, manufacturersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AerospaceManufacturer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AerospaceManufacturer
			// with a matching manufacturersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , manufacturersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AerospaceManufacturerObj from the Manufacturers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Manufacturers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Manufacturers", manufacturersId )
				return utils.RequestResult{false, msg, "removeManufacturers", childObj}
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
// adds one or more componentsIds as a Components to a Supplier
//----------------------------------------------------------------------------
func AddComponentsToSupplier ( supplierId uint64, componentsIds string )(utils.RequestResult) {

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
		ids := strings.Split( componentsIds, ",")

		for _, componentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Component_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Component_
			// with a matching componentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , componentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Components using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Components").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Components", componentsId )
				return utils.RequestResult{false, msg, "unassignComponents", childObj}
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
// removes one or more componentsIds as a Components from a Supplier
//----------------------------------------------------------------------------
func RemoveComponentsFromSupplier( supplierId uint64, componentsIds string )(utils.RequestResult) {
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
		ids := strings.Split( componentsIds, ",")

		for _, componentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Component_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Component_
			// with a matching componentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , componentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Component_Obj from the Components array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Components").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Components", componentsId )
				return utils.RequestResult{false, msg, "removeComponents", childObj}
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
// adds one or more engineTypesIds as a EngineTypes to a Supplier
//----------------------------------------------------------------------------
func AddEngineTypesToSupplier ( supplierId uint64, engineTypesIds string )(utils.RequestResult) {

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
		ids := strings.Split( engineTypesIds, ",")

		for _, engineTypesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EngineType

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EngineType
			// with a matching engineTypesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , engineTypesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the EngineTypes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EngineTypes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EngineTypes", engineTypesId )
				return utils.RequestResult{false, msg, "unassignEngineTypes", childObj}
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
// removes one or more engineTypesIds as a EngineTypes from a Supplier
//----------------------------------------------------------------------------
func RemoveEngineTypesFromSupplier( supplierId uint64, engineTypesIds string )(utils.RequestResult) {
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
		ids := strings.Split( engineTypesIds, ",")

		for _, engineTypesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EngineType

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EngineType
			// with a matching engineTypesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , engineTypesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EngineTypeObj from the EngineTypes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EngineTypes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EngineTypes", engineTypesId )
				return utils.RequestResult{false, msg, "removeEngineTypes", childObj}
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
// adds one or more avionicsSuitesIds as a AvionicsSuites to a Supplier
//----------------------------------------------------------------------------
func AddAvionicsSuitesToSupplier ( supplierId uint64, avionicsSuitesIds string )(utils.RequestResult) {

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
		ids := strings.Split( avionicsSuitesIds, ",")

		for _, avionicsSuitesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AvionicsSuite

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AvionicsSuite
			// with a matching avionicsSuitesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , avionicsSuitesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AvionicsSuites using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AvionicsSuites").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AvionicsSuites", avionicsSuitesId )
				return utils.RequestResult{false, msg, "unassignAvionicsSuites", childObj}
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
// removes one or more avionicsSuitesIds as a AvionicsSuites from a Supplier
//----------------------------------------------------------------------------
func RemoveAvionicsSuitesFromSupplier( supplierId uint64, avionicsSuitesIds string )(utils.RequestResult) {
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
		ids := strings.Split( avionicsSuitesIds, ",")

		for _, avionicsSuitesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AvionicsSuite

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AvionicsSuite
			// with a matching avionicsSuitesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , avionicsSuitesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AvionicsSuiteObj from the AvionicsSuites array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AvionicsSuites").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AvionicsSuites", avionicsSuitesId )
				return utils.RequestResult{false, msg, "removeAvionicsSuites", childObj}
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
// adds one or more apusIds as a Apus to a Supplier
//----------------------------------------------------------------------------
func AddApusToSupplier ( supplierId uint64, apusIds string )(utils.RequestResult) {

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
		ids := strings.Split( apusIds, ",")

		for _, apusId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.APU

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a APU
			// with a matching apusId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , apusId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Apus using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Apus").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Apus", apusId )
				return utils.RequestResult{false, msg, "unassignApus", childObj}
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
// removes one or more apusIds as a Apus from a Supplier
//----------------------------------------------------------------------------
func RemoveApusFromSupplier( supplierId uint64, apusIds string )(utils.RequestResult) {
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
		ids := strings.Split( apusIds, ",")

		for _, apusId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.APU

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a APU
			// with a matching apusId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , apusId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove APUObj from the Apus array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Apus").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Apus", apusId )
				return utils.RequestResult{false, msg, "removeApus", childObj}
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
// adds one or more landingGearsIds as a LandingGears to a Supplier
//----------------------------------------------------------------------------
func AddLandingGearsToSupplier ( supplierId uint64, landingGearsIds string )(utils.RequestResult) {

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
		ids := strings.Split( landingGearsIds, ",")

		for _, landingGearsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LandingGear

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LandingGear
			// with a matching landingGearsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , landingGearsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LandingGears using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LandingGears").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LandingGears", landingGearsId )
				return utils.RequestResult{false, msg, "unassignLandingGears", childObj}
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
// removes one or more landingGearsIds as a LandingGears from a Supplier
//----------------------------------------------------------------------------
func RemoveLandingGearsFromSupplier( supplierId uint64, landingGearsIds string )(utils.RequestResult) {
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
		ids := strings.Split( landingGearsIds, ",")

		for _, landingGearsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LandingGear

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LandingGear
			// with a matching landingGearsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , landingGearsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LandingGearObj from the LandingGears array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LandingGears").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LandingGears", landingGearsId )
				return utils.RequestResult{false, msg, "removeLandingGears", childObj}
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

