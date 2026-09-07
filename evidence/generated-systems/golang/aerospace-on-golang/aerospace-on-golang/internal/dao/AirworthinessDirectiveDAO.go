package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AirworthinessDirectiveDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAirworthinessDirective - creates a new db entry
//----------------------------------------------------------------------------
func CreateAirworthinessDirective(obj model.AirworthinessDirective)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AirworthinessDirective with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AirworthinessDirective", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAirworthinessDirective", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAirworthinessDirective - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAirworthinessDirective(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AirworthinessDirective

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AirworthinessDirective with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AirworthinessDirective using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AirworthinessDirective using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAirworthinessDirective", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAirworthinessDirective - returns all
//----------------------------------------------------------------------------
func GetAllAirworthinessDirective()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AirworthinessDirective

	//----------------------------------------------------------------------------
	// Request the ORM to find all AirworthinessDirective
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AirworthinessDirective" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AirworthinessDirective", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAirworthinessDirective", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAirworthinessDirective - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAirworthinessDirective(obj model.AirworthinessDirective)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AirworthinessDirective using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AirworthinessDirective using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAirworthinessDirective", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAirworthinessDirective - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAirworthinessDirective(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AirworthinessDirective with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAirworthinessDirective(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AirworthinessDirective so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AirworthinessDirective)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AirworthinessDirective using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AirworthinessDirective using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAirworthinessDirective", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more workOrdersIds as a WorkOrders to a AirworthinessDirective
//----------------------------------------------------------------------------
func AddWorkOrdersToAirworthinessDirective ( airworthinessDirectiveId uint64, workOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AirworthinessDirective with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAirworthinessDirective(airworthinessDirectiveId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AirworthinessDirective so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AirworthinessDirective)

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
		// retrieve the modified AirworthinessDirective from the gorm
		//----------------------------------------------------------------------------
		return GetAirworthinessDirective(airworthinessDirectiveId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more workOrdersIds as a WorkOrders from a AirworthinessDirective
//----------------------------------------------------------------------------
func RemoveWorkOrdersFromAirworthinessDirective( airworthinessDirectiveId uint64, workOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AirworthinessDirective with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAirworthinessDirective(airworthinessDirectiveId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AirworthinessDirective so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AirworthinessDirective)

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
		// retrieve the modified AirworthinessDirective from the gorm
		//----------------------------------------------------------------------------
		return GetAirworthinessDirective(airworthinessDirectiveId)

	} else {
		return parentRequestResult
	}
}

