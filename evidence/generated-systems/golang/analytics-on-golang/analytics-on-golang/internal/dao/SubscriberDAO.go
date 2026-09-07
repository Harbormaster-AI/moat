package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing SubscriberDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateSubscriber - creates a new db entry
//----------------------------------------------------------------------------
func CreateSubscriber(obj model.Subscriber)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Subscriber with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Subscriber", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateSubscriber", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetSubscriber - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetSubscriber(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Subscriber

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Subscriber with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Subscriber using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Subscriber using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetSubscriber", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllSubscriber - returns all
//----------------------------------------------------------------------------
func GetAllSubscriber()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Subscriber

	//----------------------------------------------------------------------------
	// Request the ORM to find all Subscriber
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Subscriber" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Subscriber", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllSubscriber", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateSubscriber - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateSubscriber(obj model.Subscriber)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Subscriber using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Subscriber using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateSubscriber", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteSubscriber - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteSubscriber(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Subscriber with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetSubscriber(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Subscriber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Subscriber)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Subscriber using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Subscriber using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteSubscriber", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more alertsIds as a Alerts to a Subscriber
//----------------------------------------------------------------------------
func AddAlertsToSubscriber ( subscriberId uint64, alertsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Subscriber with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSubscriber(subscriberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Subscriber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Subscriber)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Alert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Alert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Alerts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "unassignAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Subscriber from the gorm
		//----------------------------------------------------------------------------
		return GetSubscriber(subscriberId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more alertsIds as a Alerts from a Subscriber
//----------------------------------------------------------------------------
func RemoveAlertsFromSubscriber( subscriberId uint64, alertsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Subscriber with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetSubscriber(subscriberId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Subscriber so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Subscriber)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Alert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Alert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AlertObj from the Alerts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "removeAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Subscriber from the gorm
		//----------------------------------------------------------------------------
		return GetSubscriber(subscriberId)

	} else {
		return parentRequestResult
	}
}

