package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ConnectedAircraftDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateConnectedAircraft - creates a new db entry
//----------------------------------------------------------------------------
func CreateConnectedAircraft(obj model.ConnectedAircraft)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ConnectedAircraft with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ConnectedAircraft", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateConnectedAircraft", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetConnectedAircraft - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetConnectedAircraft(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ConnectedAircraft

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ConnectedAircraft with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ConnectedAircraft using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ConnectedAircraft using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetConnectedAircraft", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllConnectedAircraft - returns all
//----------------------------------------------------------------------------
func GetAllConnectedAircraft()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ConnectedAircraft

	//----------------------------------------------------------------------------
	// Request the ORM to find all ConnectedAircraft
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ConnectedAircraft" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ConnectedAircraft", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllConnectedAircraft", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateConnectedAircraft - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateConnectedAircraft(obj model.ConnectedAircraft)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ConnectedAircraft using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ConnectedAircraft using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateConnectedAircraft", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteConnectedAircraft - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteConnectedAircraft(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ConnectedAircraft with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetConnectedAircraft(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConnectedAircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ConnectedAircraft)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ConnectedAircraft using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ConnectedAircraft using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteConnectedAircraft", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Aircraft on a ConnectedAircraft
//----------------------------------------------------------------------------
func AssignAircraftToConnectedAircraft( connectedAircraftId uint64, aircraftId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ConnectedAircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConnectedAircraft(connectedAircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConnectedAircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConnectedAircraft)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Aircraft

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Aircraft with a
		// matching aircraftId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, aircraftId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Aircraft	to the ConnectedAircraft
			//----------------------------------------------------------------------------
			parentObj.Aircraft = &childObj

			//----------------------------------------------------------------------------
			// save the ConnectedAircraft
			//----------------------------------------------------------------------------
			return UpdateConnectedAircraft(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Aircraft", aircraftId )
			return utils.RequestResult{false, msg, "assignAircraft", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Aircraft on a ConnectedAircraft
//----------------------------------------------------------------------------
func UnassignAircraftFromConnectedAircraft(connectedAircraftId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ConnectedAircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConnectedAircraft(connectedAircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConnectedAircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConnectedAircraft)

		//----------------------------------------------------------------------------
		// assign an empty Aircraft to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.Aircraft = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Aircraft
		//----------------------------------------------------------------------------
		parentObj.AircraftId = nil;

		//----------------------------------------------------------------------------
		// save the ConnectedAircraft
		//----------------------------------------------------------------------------
		return UpdateConnectedAircraft(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more flightHealthEventsIds as a FlightHealthEvents to a ConnectedAircraft
//----------------------------------------------------------------------------
func AddFlightHealthEventsToConnectedAircraft ( connectedAircraftId uint64, flightHealthEventsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ConnectedAircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConnectedAircraft(connectedAircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConnectedAircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConnectedAircraft)

		// slice the ids on comma with no spaces
		ids := strings.Split( flightHealthEventsIds, ",")

		for _, flightHealthEventsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FlightHealthEvent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FlightHealthEvent
			// with a matching flightHealthEventsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , flightHealthEventsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FlightHealthEvents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FlightHealthEvents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FlightHealthEvents", flightHealthEventsId )
				return utils.RequestResult{false, msg, "unassignFlightHealthEvents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ConnectedAircraft from the gorm
		//----------------------------------------------------------------------------
		return GetConnectedAircraft(connectedAircraftId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more flightHealthEventsIds as a FlightHealthEvents from a ConnectedAircraft
//----------------------------------------------------------------------------
func RemoveFlightHealthEventsFromConnectedAircraft( connectedAircraftId uint64, flightHealthEventsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ConnectedAircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConnectedAircraft(connectedAircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConnectedAircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConnectedAircraft)

		// slice the ids on comma with no spaces
		ids := strings.Split( flightHealthEventsIds, ",")

		for _, flightHealthEventsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FlightHealthEvent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FlightHealthEvent
			// with a matching flightHealthEventsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , flightHealthEventsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FlightHealthEventObj from the FlightHealthEvents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FlightHealthEvents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FlightHealthEvents", flightHealthEventsId )
				return utils.RequestResult{false, msg, "removeFlightHealthEvents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ConnectedAircraft from the gorm
		//----------------------------------------------------------------------------
		return GetConnectedAircraft(connectedAircraftId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more softwareLoadsIds as a SoftwareLoads to a ConnectedAircraft
//----------------------------------------------------------------------------
func AddSoftwareLoadsToConnectedAircraft ( connectedAircraftId uint64, softwareLoadsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ConnectedAircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConnectedAircraft(connectedAircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConnectedAircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConnectedAircraft)

		// slice the ids on comma with no spaces
		ids := strings.Split( softwareLoadsIds, ",")

		for _, softwareLoadsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SoftwareLoad

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SoftwareLoad
			// with a matching softwareLoadsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , softwareLoadsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SoftwareLoads using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SoftwareLoads").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SoftwareLoads", softwareLoadsId )
				return utils.RequestResult{false, msg, "unassignSoftwareLoads", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ConnectedAircraft from the gorm
		//----------------------------------------------------------------------------
		return GetConnectedAircraft(connectedAircraftId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more softwareLoadsIds as a SoftwareLoads from a ConnectedAircraft
//----------------------------------------------------------------------------
func RemoveSoftwareLoadsFromConnectedAircraft( connectedAircraftId uint64, softwareLoadsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ConnectedAircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetConnectedAircraft(connectedAircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ConnectedAircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ConnectedAircraft)

		// slice the ids on comma with no spaces
		ids := strings.Split( softwareLoadsIds, ",")

		for _, softwareLoadsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SoftwareLoad

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SoftwareLoad
			// with a matching softwareLoadsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , softwareLoadsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SoftwareLoadObj from the SoftwareLoads array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SoftwareLoads").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SoftwareLoads", softwareLoadsId )
				return utils.RequestResult{false, msg, "removeSoftwareLoads", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ConnectedAircraft from the gorm
		//----------------------------------------------------------------------------
		return GetConnectedAircraft(connectedAircraftId)

	} else {
		return parentRequestResult
	}
}

