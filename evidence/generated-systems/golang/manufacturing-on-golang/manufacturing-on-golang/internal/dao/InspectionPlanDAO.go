package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InspectionPlanDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInspectionPlan - creates a new db entry
//----------------------------------------------------------------------------
func CreateInspectionPlan(obj model.InspectionPlan)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InspectionPlan with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InspectionPlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInspectionPlan", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInspectionPlan - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInspectionPlan(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InspectionPlan

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InspectionPlan with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InspectionPlan using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InspectionPlan using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInspectionPlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInspectionPlan - returns all
//----------------------------------------------------------------------------
func GetAllInspectionPlan()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InspectionPlan

	//----------------------------------------------------------------------------
	// Request the ORM to find all InspectionPlan
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InspectionPlan" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InspectionPlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInspectionPlan", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInspectionPlan - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInspectionPlan(obj model.InspectionPlan)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InspectionPlan using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InspectionPlan using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInspectionPlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInspectionPlan - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInspectionPlan(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InspectionPlan with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInspectionPlan(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InspectionPlan)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InspectionPlan using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InspectionPlan using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInspectionPlan", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Item on a InspectionPlan
//----------------------------------------------------------------------------
func AssignItemToInspectionPlan( inspectionPlanId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InspectionPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionPlan(inspectionPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionPlan)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching itemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, itemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Item	to the InspectionPlan
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the InspectionPlan
			//----------------------------------------------------------------------------
			return UpdateInspectionPlan(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a InspectionPlan
//----------------------------------------------------------------------------
func UnassignItemFromInspectionPlan(inspectionPlanId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InspectionPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionPlan(inspectionPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionPlan)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the InspectionPlan
		//----------------------------------------------------------------------------
		return UpdateInspectionPlan(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more characteristicsIds as a Characteristics to a InspectionPlan
//----------------------------------------------------------------------------
func AddCharacteristicsToInspectionPlan ( inspectionPlanId uint64, characteristicsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InspectionPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionPlan(inspectionPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( characteristicsIds, ",")

		for _, characteristicsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InspectionCharacteristic

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InspectionCharacteristic
			// with a matching characteristicsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , characteristicsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Characteristics using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Characteristics").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Characteristics", characteristicsId )
				return utils.RequestResult{false, msg, "unassignCharacteristics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InspectionPlan from the gorm
		//----------------------------------------------------------------------------
		return GetInspectionPlan(inspectionPlanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more characteristicsIds as a Characteristics from a InspectionPlan
//----------------------------------------------------------------------------
func RemoveCharacteristicsFromInspectionPlan( inspectionPlanId uint64, characteristicsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InspectionPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInspectionPlan(inspectionPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InspectionPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InspectionPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( characteristicsIds, ",")

		for _, characteristicsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InspectionCharacteristic

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InspectionCharacteristic
			// with a matching characteristicsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , characteristicsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InspectionCharacteristicObj from the Characteristics array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Characteristics").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Characteristics", characteristicsId )
				return utils.RequestResult{false, msg, "removeCharacteristics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InspectionPlan from the gorm
		//----------------------------------------------------------------------------
		return GetInspectionPlan(inspectionPlanId)

	} else {
		return parentRequestResult
	}
}

