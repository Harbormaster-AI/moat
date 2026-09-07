package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing APUDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAPU - creates a new db entry
//----------------------------------------------------------------------------
func CreateAPU(obj model.APU)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a APU with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a APU", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAPU", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAPU - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAPU(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.APU

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a APU with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a APU using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a APU using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAPU", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAPU - returns all
//----------------------------------------------------------------------------
func GetAllAPU()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.APU

	//----------------------------------------------------------------------------
	// Request the ORM to find all APU
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all APU" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all APU", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAPU", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAPU - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAPU(obj model.APU)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a APU using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a APU using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAPU", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAPU - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAPU(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the APU with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAPU(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.APU so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.APU)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a APU using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a APU using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAPU", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Supplier on a APU
//----------------------------------------------------------------------------
func AssignSupplierToAPU( aPUId uint64, supplierId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the APU with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAPU(aPUId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.APU so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.APU)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Supplier

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Supplier with a
		// matching supplierId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, supplierId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Supplier	to the APU
			//----------------------------------------------------------------------------
			parentObj.Supplier = &childObj

			//----------------------------------------------------------------------------
			// save the APU
			//----------------------------------------------------------------------------
			return UpdateAPU(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Supplier", supplierId )
			return utils.RequestResult{false, msg, "assignSupplier", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Supplier on a APU
//----------------------------------------------------------------------------
func UnassignSupplierFromAPU(aPUId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the APU with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAPU(aPUId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.APU so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.APU)

		//----------------------------------------------------------------------------
		// assign an empty Supplier to the Supplier
		//----------------------------------------------------------------------------
		parentObj.Supplier = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Supplier
		//----------------------------------------------------------------------------
		parentObj.SupplierId = nil;

		//----------------------------------------------------------------------------
		// save the APU
		//----------------------------------------------------------------------------
		return UpdateAPU(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more variantsIds as a Variants to a APU
//----------------------------------------------------------------------------
func AddVariantsToAPU ( aPUId uint64, variantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the APU with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAPU(aPUId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.APU so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.APU)

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
		// retrieve the modified APU from the gorm
		//----------------------------------------------------------------------------
		return GetAPU(aPUId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more variantsIds as a Variants from a APU
//----------------------------------------------------------------------------
func RemoveVariantsFromAPU( aPUId uint64, variantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the APU with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAPU(aPUId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.APU so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.APU)

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
		// retrieve the modified APU from the gorm
		//----------------------------------------------------------------------------
		return GetAPU(aPUId)

	} else {
		return parentRequestResult
	}
}

