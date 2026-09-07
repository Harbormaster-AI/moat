package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FraudScenarioDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFraudScenario - creates a new db entry
//----------------------------------------------------------------------------
func CreateFraudScenario(obj model.FraudScenario)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FraudScenario with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FraudScenario", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFraudScenario", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFraudScenario - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFraudScenario(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FraudScenario

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FraudScenario with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FraudScenario using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FraudScenario using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFraudScenario", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFraudScenario - returns all
//----------------------------------------------------------------------------
func GetAllFraudScenario()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FraudScenario

	//----------------------------------------------------------------------------
	// Request the ORM to find all FraudScenario
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FraudScenario" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FraudScenario", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFraudScenario", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFraudScenario - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFraudScenario(obj model.FraudScenario)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FraudScenario using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FraudScenario using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFraudScenario", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFraudScenario - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFraudScenario(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FraudScenario with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFraudScenario(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FraudScenario)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FraudScenario using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FraudScenario using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFraudScenario", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more modelsIds as a Models to a FraudScenario
//----------------------------------------------------------------------------
func AddModelsToFraudScenario ( fraudScenarioId uint64, modelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FraudScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudScenario(fraudScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( modelsIds, ",")

		for _, modelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Model_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Model_
			// with a matching modelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , modelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Models using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Models").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Models", modelsId )
				return utils.RequestResult{false, msg, "unassignModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FraudScenario from the gorm
		//----------------------------------------------------------------------------
		return GetFraudScenario(fraudScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelsIds as a Models from a FraudScenario
//----------------------------------------------------------------------------
func RemoveModelsFromFraudScenario( fraudScenarioId uint64, modelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FraudScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudScenario(fraudScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( modelsIds, ",")

		for _, modelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Model_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Model_
			// with a matching modelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , modelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Model_Obj from the Models array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Models").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Models", modelsId )
				return utils.RequestResult{false, msg, "removeModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FraudScenario from the gorm
		//----------------------------------------------------------------------------
		return GetFraudScenario(fraudScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a FraudScenario
//----------------------------------------------------------------------------
func AddDatasetsToFraudScenario ( fraudScenarioId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FraudScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudScenario(fraudScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( datasetsIds, ",")

		for _, datasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching datasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , datasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Datasets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Datasets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Datasets", datasetsId )
				return utils.RequestResult{false, msg, "unassignDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FraudScenario from the gorm
		//----------------------------------------------------------------------------
		return GetFraudScenario(fraudScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a FraudScenario
//----------------------------------------------------------------------------
func RemoveDatasetsFromFraudScenario( fraudScenarioId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FraudScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudScenario(fraudScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( datasetsIds, ",")

		for _, datasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching datasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , datasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSetObj from the Datasets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Datasets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Datasets", datasetsId )
				return utils.RequestResult{false, msg, "removeDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FraudScenario from the gorm
		//----------------------------------------------------------------------------
		return GetFraudScenario(fraudScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more alertsIds as a Alerts to a FraudScenario
//----------------------------------------------------------------------------
func AddAlertsToFraudScenario ( fraudScenarioId uint64, alertsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FraudScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudScenario(fraudScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudScenario)

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
		// retrieve the modified FraudScenario from the gorm
		//----------------------------------------------------------------------------
		return GetFraudScenario(fraudScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more alertsIds as a Alerts from a FraudScenario
//----------------------------------------------------------------------------
func RemoveAlertsFromFraudScenario( fraudScenarioId uint64, alertsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FraudScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudScenario(fraudScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudScenario)

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
		// retrieve the modified FraudScenario from the gorm
		//----------------------------------------------------------------------------
		return GetFraudScenario(fraudScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more signalsIds as a Signals to a FraudScenario
//----------------------------------------------------------------------------
func AddSignalsToFraudScenario ( fraudScenarioId uint64, signalsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FraudScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudScenario(fraudScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( signalsIds, ",")

		for _, signalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FraudSignal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FraudSignal
			// with a matching signalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , signalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Signals using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Signals").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Signals", signalsId )
				return utils.RequestResult{false, msg, "unassignSignals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FraudScenario from the gorm
		//----------------------------------------------------------------------------
		return GetFraudScenario(fraudScenarioId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more signalsIds as a Signals from a FraudScenario
//----------------------------------------------------------------------------
func RemoveSignalsFromFraudScenario( fraudScenarioId uint64, signalsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the FraudScenario with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudScenario(fraudScenarioId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudScenario so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudScenario)

		// slice the ids on comma with no spaces
		ids := strings.Split( signalsIds, ",")

		for _, signalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FraudSignal

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FraudSignal
			// with a matching signalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , signalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FraudSignalObj from the Signals array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Signals").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Signals", signalsId )
				return utils.RequestResult{false, msg, "removeSignals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified FraudScenario from the gorm
		//----------------------------------------------------------------------------
		return GetFraudScenario(fraudScenarioId)

	} else {
		return parentRequestResult
	}
}

