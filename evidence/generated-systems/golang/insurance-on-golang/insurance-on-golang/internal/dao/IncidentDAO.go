package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing IncidentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateIncident - creates a new db entry
//----------------------------------------------------------------------------
func CreateIncident(obj model.Incident)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Incident with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Incident", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateIncident", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetIncident - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetIncident(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Incident

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Incident with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Incident using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Incident using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetIncident", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllIncident - returns all
//----------------------------------------------------------------------------
func GetAllIncident()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Incident

	//----------------------------------------------------------------------------
	// Request the ORM to find all Incident
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Incident" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Incident", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllIncident", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateIncident - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateIncident(obj model.Incident)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Incident using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Incident using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateIncident", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteIncident - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteIncident(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Incident with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetIncident(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Incident so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Incident)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Incident using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Incident using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteIncident", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Claim on a Incident
//----------------------------------------------------------------------------
func AssignClaimToIncident( incidentId uint64, claimId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Incident with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIncident(incidentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Incident so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Incident)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Claim

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Claim with a
		// matching claimId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, claimId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Claim	to the Incident
			//----------------------------------------------------------------------------
			parentObj.Claim = &childObj

			//----------------------------------------------------------------------------
			// save the Incident
			//----------------------------------------------------------------------------
			return UpdateIncident(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claim", claimId )
			return utils.RequestResult{false, msg, "assignClaim", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Claim on a Incident
//----------------------------------------------------------------------------
func UnassignClaimFromIncident(incidentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Incident with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIncident(incidentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Incident so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Incident)

		//----------------------------------------------------------------------------
		// assign an empty Claim to the Claim
		//----------------------------------------------------------------------------
		parentObj.Claim = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Claim
		//----------------------------------------------------------------------------
		parentObj.ClaimId = nil;

		//----------------------------------------------------------------------------
		// save the Incident
		//----------------------------------------------------------------------------
		return UpdateIncident(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more insuredObjectsIds as a InsuredObjects to a Incident
//----------------------------------------------------------------------------
func AddInsuredObjectsToIncident ( incidentId uint64, insuredObjectsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Incident with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIncident(incidentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Incident so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Incident)

		// slice the ids on comma with no spaces
		ids := strings.Split( insuredObjectsIds, ",")

		for _, insuredObjectsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsuredObject

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsuredObject
			// with a matching insuredObjectsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insuredObjectsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InsuredObjects using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsuredObjects").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsuredObjects", insuredObjectsId )
				return utils.RequestResult{false, msg, "unassignInsuredObjects", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Incident from the gorm
		//----------------------------------------------------------------------------
		return GetIncident(incidentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more insuredObjectsIds as a InsuredObjects from a Incident
//----------------------------------------------------------------------------
func RemoveInsuredObjectsFromIncident( incidentId uint64, insuredObjectsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Incident with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIncident(incidentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Incident so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Incident)

		// slice the ids on comma with no spaces
		ids := strings.Split( insuredObjectsIds, ",")

		for _, insuredObjectsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsuredObject

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsuredObject
			// with a matching insuredObjectsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , insuredObjectsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InsuredObjectObj from the InsuredObjects array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InsuredObjects").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsuredObjects", insuredObjectsId )
				return utils.RequestResult{false, msg, "removeInsuredObjects", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Incident from the gorm
		//----------------------------------------------------------------------------
		return GetIncident(incidentId)

	} else {
		return parentRequestResult
	}
}

