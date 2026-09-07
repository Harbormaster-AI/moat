package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FlightHealthEventDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFlightHealthEvent - creates a new db entry
//----------------------------------------------------------------------------
func CreateFlightHealthEvent(obj model.FlightHealthEvent)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FlightHealthEvent with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FlightHealthEvent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFlightHealthEvent", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFlightHealthEvent - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFlightHealthEvent(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FlightHealthEvent

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FlightHealthEvent with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FlightHealthEvent using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FlightHealthEvent using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFlightHealthEvent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFlightHealthEvent - returns all
//----------------------------------------------------------------------------
func GetAllFlightHealthEvent()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FlightHealthEvent

	//----------------------------------------------------------------------------
	// Request the ORM to find all FlightHealthEvent
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FlightHealthEvent" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FlightHealthEvent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFlightHealthEvent", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFlightHealthEvent - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFlightHealthEvent(obj model.FlightHealthEvent)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FlightHealthEvent using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FlightHealthEvent using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFlightHealthEvent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFlightHealthEvent - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFlightHealthEvent(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FlightHealthEvent with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFlightHealthEvent(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FlightHealthEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FlightHealthEvent)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FlightHealthEvent using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FlightHealthEvent using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFlightHealthEvent", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ConnectedAircraft on a FlightHealthEvent
//----------------------------------------------------------------------------
func AssignConnectedAircraftToFlightHealthEvent( flightHealthEventId uint64, connectedAircraftId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FlightHealthEvent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFlightHealthEvent(flightHealthEventId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FlightHealthEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FlightHealthEvent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ConnectedAircraft

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ConnectedAircraft with a
		// matching connectedAircraftId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, connectedAircraftId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ConnectedAircraft	to the FlightHealthEvent
			//----------------------------------------------------------------------------
			parentObj.ConnectedAircraft = &childObj

			//----------------------------------------------------------------------------
			// save the FlightHealthEvent
			//----------------------------------------------------------------------------
			return UpdateFlightHealthEvent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ConnectedAircraft", connectedAircraftId )
			return utils.RequestResult{false, msg, "assignConnectedAircraft", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ConnectedAircraft on a FlightHealthEvent
//----------------------------------------------------------------------------
func UnassignConnectedAircraftFromFlightHealthEvent(flightHealthEventId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FlightHealthEvent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFlightHealthEvent(flightHealthEventId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FlightHealthEvent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FlightHealthEvent)

		//----------------------------------------------------------------------------
		// assign an empty ConnectedAircraft to the ConnectedAircraft
		//----------------------------------------------------------------------------
		parentObj.ConnectedAircraft = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ConnectedAircraft
		//----------------------------------------------------------------------------
		parentObj.ConnectedAircraftId = nil;

		//----------------------------------------------------------------------------
		// save the FlightHealthEvent
		//----------------------------------------------------------------------------
		return UpdateFlightHealthEvent(parentObj)

	} else {
		return parentRequestResult
	}

}


