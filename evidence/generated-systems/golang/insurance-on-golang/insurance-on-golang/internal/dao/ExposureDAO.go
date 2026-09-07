package dao

import (
    "insurance-on-golang/internal/model"
    "insurance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ExposureDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateExposure - creates a new db entry
//----------------------------------------------------------------------------
func CreateExposure(obj model.Exposure)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Exposure with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Exposure", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateExposure", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetExposure - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetExposure(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Exposure

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Exposure with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Exposure using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Exposure using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetExposure", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllExposure - returns all
//----------------------------------------------------------------------------
func GetAllExposure()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Exposure

	//----------------------------------------------------------------------------
	// Request the ORM to find all Exposure
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Exposure" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Exposure", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllExposure", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateExposure - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateExposure(obj model.Exposure)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Exposure using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Exposure using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateExposure", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteExposure - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteExposure(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetExposure(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Exposure)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Exposure using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Exposure using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteExposure", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Claim on a Exposure
//----------------------------------------------------------------------------
func AssignClaimToExposure( exposureId uint64, claimId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

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
			// assign the Claim	to the Exposure
			//----------------------------------------------------------------------------
			parentObj.Claim = &childObj

			//----------------------------------------------------------------------------
			// save the Exposure
			//----------------------------------------------------------------------------
			return UpdateExposure(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claim", claimId )
			return utils.RequestResult{false, msg, "assignClaim", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Claim on a Exposure
//----------------------------------------------------------------------------
func UnassignClaimFromExposure(exposureId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

		//----------------------------------------------------------------------------
		// assign an empty Claim to the Claim
		//----------------------------------------------------------------------------
		parentObj.Claim = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Claim
		//----------------------------------------------------------------------------
		parentObj.ClaimId = nil;

		//----------------------------------------------------------------------------
		// save the Exposure
		//----------------------------------------------------------------------------
		return UpdateExposure(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a PolicyCoverage on a Exposure
//----------------------------------------------------------------------------
func AssignPolicyCoverageToExposure( exposureId uint64, policyCoverageId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.PolicyCoverage

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a PolicyCoverage with a
		// matching policyCoverageId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, policyCoverageId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the PolicyCoverage	to the Exposure
			//----------------------------------------------------------------------------
			parentObj.PolicyCoverage = &childObj

			//----------------------------------------------------------------------------
			// save the Exposure
			//----------------------------------------------------------------------------
			return UpdateExposure(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PolicyCoverage", policyCoverageId )
			return utils.RequestResult{false, msg, "assignPolicyCoverage", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a PolicyCoverage on a Exposure
//----------------------------------------------------------------------------
func UnassignPolicyCoverageFromExposure(exposureId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

		//----------------------------------------------------------------------------
		// assign an empty PolicyCoverage to the PolicyCoverage
		//----------------------------------------------------------------------------
		parentObj.PolicyCoverage = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the PolicyCoverage
		//----------------------------------------------------------------------------
		parentObj.PolicyCoverageId = nil;

		//----------------------------------------------------------------------------
		// save the Exposure
		//----------------------------------------------------------------------------
		return UpdateExposure(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a InsuredObject on a Exposure
//----------------------------------------------------------------------------
func AssignInsuredObjectToExposure( exposureId uint64, insuredObjectId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InsuredObject

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InsuredObject with a
		// matching insuredObjectId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, insuredObjectId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InsuredObject	to the Exposure
			//----------------------------------------------------------------------------
			parentObj.InsuredObject = &childObj

			//----------------------------------------------------------------------------
			// save the Exposure
			//----------------------------------------------------------------------------
			return UpdateExposure(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InsuredObject", insuredObjectId )
			return utils.RequestResult{false, msg, "assignInsuredObject", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InsuredObject on a Exposure
//----------------------------------------------------------------------------
func UnassignInsuredObjectFromExposure(exposureId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

		//----------------------------------------------------------------------------
		// assign an empty InsuredObject to the InsuredObject
		//----------------------------------------------------------------------------
		parentObj.InsuredObject = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InsuredObject
		//----------------------------------------------------------------------------
		parentObj.InsuredObjectId = nil;

		//----------------------------------------------------------------------------
		// save the Exposure
		//----------------------------------------------------------------------------
		return UpdateExposure(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more reservesIds as a Reserves to a Exposure
//----------------------------------------------------------------------------
func AddReservesToExposure ( exposureId uint64, reservesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

		// slice the ids on comma with no spaces
		ids := strings.Split( reservesIds, ",")

		for _, reservesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClaimReserve

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClaimReserve
			// with a matching reservesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reservesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reserves using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reserves").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reserves", reservesId )
				return utils.RequestResult{false, msg, "unassignReserves", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Exposure from the gorm
		//----------------------------------------------------------------------------
		return GetExposure(exposureId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reservesIds as a Reserves from a Exposure
//----------------------------------------------------------------------------
func RemoveReservesFromExposure( exposureId uint64, reservesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

		// slice the ids on comma with no spaces
		ids := strings.Split( reservesIds, ",")

		for _, reservesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClaimReserve

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClaimReserve
			// with a matching reservesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reservesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimReserveObj from the Reserves array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reserves").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reserves", reservesId )
				return utils.RequestResult{false, msg, "removeReserves", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Exposure from the gorm
		//----------------------------------------------------------------------------
		return GetExposure(exposureId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more paymentsIds as a Payments to a Exposure
//----------------------------------------------------------------------------
func AddPaymentsToExposure ( exposureId uint64, paymentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClaimPayment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClaimPayment
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Payments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "unassignPayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Exposure from the gorm
		//----------------------------------------------------------------------------
		return GetExposure(exposureId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more paymentsIds as a Payments from a Exposure
//----------------------------------------------------------------------------
func RemovePaymentsFromExposure( exposureId uint64, paymentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Exposure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetExposure(exposureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exposure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exposure)

		// slice the ids on comma with no spaces
		ids := strings.Split( paymentsIds, ",")

		for _, paymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ClaimPayment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ClaimPayment
			// with a matching paymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , paymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimPaymentObj from the Payments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Payments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payments", paymentsId )
				return utils.RequestResult{false, msg, "removePayments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Exposure from the gorm
		//----------------------------------------------------------------------------
		return GetExposure(exposureId)

	} else {
		return parentRequestResult
	}
}

