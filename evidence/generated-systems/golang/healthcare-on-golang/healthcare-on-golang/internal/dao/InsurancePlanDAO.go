package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InsurancePlanDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInsurancePlan - creates a new db entry
//----------------------------------------------------------------------------
func CreateInsurancePlan(obj model.InsurancePlan)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InsurancePlan with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InsurancePlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInsurancePlan", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInsurancePlan - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInsurancePlan(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InsurancePlan

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InsurancePlan with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InsurancePlan using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InsurancePlan using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInsurancePlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInsurancePlan - returns all
//----------------------------------------------------------------------------
func GetAllInsurancePlan()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InsurancePlan

	//----------------------------------------------------------------------------
	// Request the ORM to find all InsurancePlan
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InsurancePlan" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InsurancePlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInsurancePlan", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInsurancePlan - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInsurancePlan(obj model.InsurancePlan)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InsurancePlan using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InsurancePlan using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInsurancePlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInsurancePlan - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInsurancePlan(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InsurancePlan with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInsurancePlan(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InsurancePlan)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InsurancePlan using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InsurancePlan using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInsurancePlan", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Payer on a InsurancePlan
//----------------------------------------------------------------------------
func AssignPayerToInsurancePlan( insurancePlanId uint64, payerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InsurancePlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurancePlan(insurancePlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsurancePlan)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InsurancePayer

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InsurancePayer with a
		// matching payerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, payerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Payer	to the InsurancePlan
			//----------------------------------------------------------------------------
			parentObj.Payer = &childObj

			//----------------------------------------------------------------------------
			// save the InsurancePlan
			//----------------------------------------------------------------------------
			return UpdateInsurancePlan(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Payer", payerId )
			return utils.RequestResult{false, msg, "assignPayer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Payer on a InsurancePlan
//----------------------------------------------------------------------------
func UnassignPayerFromInsurancePlan(insurancePlanId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsurancePlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurancePlan(insurancePlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsurancePlan)

		//----------------------------------------------------------------------------
		// assign an empty InsurancePayer to the Payer
		//----------------------------------------------------------------------------
		parentObj.Payer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Payer
		//----------------------------------------------------------------------------
		parentObj.PayerId = nil;

		//----------------------------------------------------------------------------
		// save the InsurancePlan
		//----------------------------------------------------------------------------
		return UpdateInsurancePlan(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more coveragesIds as a Coverages to a InsurancePlan
//----------------------------------------------------------------------------
func AddCoveragesToInsurancePlan ( insurancePlanId uint64, coveragesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsurancePlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurancePlan(insurancePlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsurancePlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( coveragesIds, ",")

		for _, coveragesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Coverage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Coverage
			// with a matching coveragesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coveragesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Coverages using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Coverages").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverages", coveragesId )
				return utils.RequestResult{false, msg, "unassignCoverages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsurancePlan from the gorm
		//----------------------------------------------------------------------------
		return GetInsurancePlan(insurancePlanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more coveragesIds as a Coverages from a InsurancePlan
//----------------------------------------------------------------------------
func RemoveCoveragesFromInsurancePlan( insurancePlanId uint64, coveragesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InsurancePlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurancePlan(insurancePlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsurancePlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( coveragesIds, ",")

		for _, coveragesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Coverage

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Coverage
			// with a matching coveragesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , coveragesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CoverageObj from the Coverages array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Coverages").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Coverages", coveragesId )
				return utils.RequestResult{false, msg, "removeCoverages", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsurancePlan from the gorm
		//----------------------------------------------------------------------------
		return GetInsurancePlan(insurancePlanId)

	} else {
		return parentRequestResult
	}
}

