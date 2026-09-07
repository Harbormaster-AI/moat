package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InsurancePayerDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInsurancePayer - creates a new db entry
//----------------------------------------------------------------------------
func CreateInsurancePayer(obj model.InsurancePayer)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InsurancePayer with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InsurancePayer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInsurancePayer", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInsurancePayer - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInsurancePayer(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InsurancePayer

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InsurancePayer with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InsurancePayer using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InsurancePayer using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInsurancePayer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInsurancePayer - returns all
//----------------------------------------------------------------------------
func GetAllInsurancePayer()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InsurancePayer

	//----------------------------------------------------------------------------
	// Request the ORM to find all InsurancePayer
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InsurancePayer" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InsurancePayer", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInsurancePayer", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInsurancePayer - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInsurancePayer(obj model.InsurancePayer)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InsurancePayer using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InsurancePayer using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInsurancePayer", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInsurancePayer - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInsurancePayer(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InsurancePayer with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInsurancePayer(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePayer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InsurancePayer)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InsurancePayer using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InsurancePayer using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInsurancePayer", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more plansIds as a Plans to a InsurancePayer
//----------------------------------------------------------------------------
func AddPlansToInsurancePayer ( insurancePayerId uint64, plansIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsurancePayer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurancePayer(insurancePayerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePayer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsurancePayer)

		// slice the ids on comma with no spaces
		ids := strings.Split( plansIds, ",")

		for _, plansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsurancePlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsurancePlan
			// with a matching plansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , plansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Plans using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Plans").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plans", plansId )
				return utils.RequestResult{false, msg, "unassignPlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsurancePayer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurancePayer(insurancePayerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more plansIds as a Plans from a InsurancePayer
//----------------------------------------------------------------------------
func RemovePlansFromInsurancePayer( insurancePayerId uint64, plansIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InsurancePayer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurancePayer(insurancePayerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePayer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsurancePayer)

		// slice the ids on comma with no spaces
		ids := strings.Split( plansIds, ",")

		for _, plansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InsurancePlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InsurancePlan
			// with a matching plansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , plansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InsurancePlanObj from the Plans array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Plans").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plans", plansId )
				return utils.RequestResult{false, msg, "removePlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsurancePayer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurancePayer(insurancePayerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more claimsIds as a Claims to a InsurancePayer
//----------------------------------------------------------------------------
func AddClaimsToInsurancePayer ( insurancePayerId uint64, claimsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InsurancePayer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurancePayer(insurancePayerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePayer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsurancePayer)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Claims using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "unassignClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsurancePayer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurancePayer(insurancePayerId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more claimsIds as a Claims from a InsurancePayer
//----------------------------------------------------------------------------
func RemoveClaimsFromInsurancePayer( insurancePayerId uint64, claimsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InsurancePayer with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInsurancePayer(insurancePayerId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InsurancePayer so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InsurancePayer)

		// slice the ids on comma with no spaces
		ids := strings.Split( claimsIds, ",")

		for _, claimsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Claim

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Claim
			// with a matching claimsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , claimsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ClaimObj from the Claims array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Claims").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Claims", claimsId )
				return utils.RequestResult{false, msg, "removeClaims", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InsurancePayer from the gorm
		//----------------------------------------------------------------------------
		return GetInsurancePayer(insurancePayerId)

	} else {
		return parentRequestResult
	}
}

