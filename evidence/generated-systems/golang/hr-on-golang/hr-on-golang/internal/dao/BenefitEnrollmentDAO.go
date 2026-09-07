package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BenefitEnrollmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBenefitEnrollment - creates a new db entry
//----------------------------------------------------------------------------
func CreateBenefitEnrollment(obj model.BenefitEnrollment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BenefitEnrollment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BenefitEnrollment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBenefitEnrollment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBenefitEnrollment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBenefitEnrollment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BenefitEnrollment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BenefitEnrollment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BenefitEnrollment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BenefitEnrollment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBenefitEnrollment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBenefitEnrollment - returns all
//----------------------------------------------------------------------------
func GetAllBenefitEnrollment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BenefitEnrollment

	//----------------------------------------------------------------------------
	// Request the ORM to find all BenefitEnrollment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BenefitEnrollment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BenefitEnrollment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBenefitEnrollment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBenefitEnrollment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBenefitEnrollment(obj model.BenefitEnrollment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BenefitEnrollment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BenefitEnrollment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBenefitEnrollment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBenefitEnrollment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBenefitEnrollment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BenefitEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBenefitEnrollment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BenefitEnrollment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BenefitEnrollment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BenefitEnrollment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBenefitEnrollment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a BenefitPlan on a BenefitEnrollment
//----------------------------------------------------------------------------
func AssignBenefitPlanToBenefitEnrollment( benefitEnrollmentId uint64, benefitPlanId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BenefitEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitEnrollment(benefitEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitEnrollment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BenefitPlan

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BenefitPlan with a
		// matching benefitPlanId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, benefitPlanId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the BenefitPlan	to the BenefitEnrollment
			//----------------------------------------------------------------------------
			parentObj.BenefitPlan = &childObj

			//----------------------------------------------------------------------------
			// save the BenefitEnrollment
			//----------------------------------------------------------------------------
			return UpdateBenefitEnrollment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BenefitPlan", benefitPlanId )
			return utils.RequestResult{false, msg, "assignBenefitPlan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a BenefitPlan on a BenefitEnrollment
//----------------------------------------------------------------------------
func UnassignBenefitPlanFromBenefitEnrollment(benefitEnrollmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BenefitEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitEnrollment(benefitEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitEnrollment)

		//----------------------------------------------------------------------------
		// assign an empty BenefitPlan to the BenefitPlan
		//----------------------------------------------------------------------------
		parentObj.BenefitPlan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the BenefitPlan
		//----------------------------------------------------------------------------
		parentObj.BenefitPlanId = nil;

		//----------------------------------------------------------------------------
		// save the BenefitEnrollment
		//----------------------------------------------------------------------------
		return UpdateBenefitEnrollment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Employee on a BenefitEnrollment
//----------------------------------------------------------------------------
func AssignEmployeeToBenefitEnrollment( benefitEnrollmentId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BenefitEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitEnrollment(benefitEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitEnrollment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching employeeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, employeeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Employee	to the BenefitEnrollment
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the BenefitEnrollment
			//----------------------------------------------------------------------------
			return UpdateBenefitEnrollment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a BenefitEnrollment
//----------------------------------------------------------------------------
func UnassignEmployeeFromBenefitEnrollment(benefitEnrollmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BenefitEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitEnrollment(benefitEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitEnrollment)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the BenefitEnrollment
		//----------------------------------------------------------------------------
		return UpdateBenefitEnrollment(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more dependentsIds as a Dependents to a BenefitEnrollment
//----------------------------------------------------------------------------
func AddDependentsToBenefitEnrollment ( benefitEnrollmentId uint64, dependentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BenefitEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitEnrollment(benefitEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitEnrollment)

		// slice the ids on comma with no spaces
		ids := strings.Split( dependentsIds, ",")

		for _, dependentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dependent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dependent
			// with a matching dependentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dependentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Dependents using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dependents").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dependents", dependentsId )
				return utils.RequestResult{false, msg, "unassignDependents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BenefitEnrollment from the gorm
		//----------------------------------------------------------------------------
		return GetBenefitEnrollment(benefitEnrollmentId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dependentsIds as a Dependents from a BenefitEnrollment
//----------------------------------------------------------------------------
func RemoveDependentsFromBenefitEnrollment( benefitEnrollmentId uint64, dependentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BenefitEnrollment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitEnrollment(benefitEnrollmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitEnrollment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitEnrollment)

		// slice the ids on comma with no spaces
		ids := strings.Split( dependentsIds, ",")

		for _, dependentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dependent

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dependent
			// with a matching dependentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dependentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DependentObj from the Dependents array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dependents").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dependents", dependentsId )
				return utils.RequestResult{false, msg, "removeDependents", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BenefitEnrollment from the gorm
		//----------------------------------------------------------------------------
		return GetBenefitEnrollment(benefitEnrollmentId)

	} else {
		return parentRequestResult
	}
}

