package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DependentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDependent - creates a new db entry
//----------------------------------------------------------------------------
func CreateDependent(obj model.Dependent)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Dependent with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Dependent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDependent", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDependent - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDependent(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Dependent

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Dependent with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Dependent using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Dependent using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDependent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDependent - returns all
//----------------------------------------------------------------------------
func GetAllDependent()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Dependent

	//----------------------------------------------------------------------------
	// Request the ORM to find all Dependent
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Dependent" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Dependent", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDependent", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDependent - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDependent(obj model.Dependent)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Dependent using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Dependent using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDependent", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDependent - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDependent(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Dependent with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDependent(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dependent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Dependent)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Dependent using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Dependent using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDependent", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a BenefitEnrollment on a Dependent
//----------------------------------------------------------------------------
func AssignBenefitEnrollmentToDependent( dependentId uint64, benefitEnrollmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dependent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDependent(dependentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dependent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dependent)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.BenefitEnrollment

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a BenefitEnrollment with a
		// matching benefitEnrollmentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, benefitEnrollmentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the BenefitEnrollment	to the Dependent
			//----------------------------------------------------------------------------
			parentObj.BenefitEnrollment = &childObj

			//----------------------------------------------------------------------------
			// save the Dependent
			//----------------------------------------------------------------------------
			return UpdateDependent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BenefitEnrollment", benefitEnrollmentId )
			return utils.RequestResult{false, msg, "assignBenefitEnrollment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a BenefitEnrollment on a Dependent
//----------------------------------------------------------------------------
func UnassignBenefitEnrollmentFromDependent(dependentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dependent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDependent(dependentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dependent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dependent)

		//----------------------------------------------------------------------------
		// assign an empty BenefitEnrollment to the BenefitEnrollment
		//----------------------------------------------------------------------------
		parentObj.BenefitEnrollment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the BenefitEnrollment
		//----------------------------------------------------------------------------
		parentObj.BenefitEnrollmentId = nil;

		//----------------------------------------------------------------------------
		// save the Dependent
		//----------------------------------------------------------------------------
		return UpdateDependent(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Employee on a Dependent
//----------------------------------------------------------------------------
func AssignEmployeeToDependent( dependentId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dependent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDependent(dependentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dependent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dependent)

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
			// assign the Employee	to the Dependent
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the Dependent
			//----------------------------------------------------------------------------
			return UpdateDependent(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a Dependent
//----------------------------------------------------------------------------
func UnassignEmployeeFromDependent(dependentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dependent with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDependent(dependentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dependent so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dependent)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the Dependent
		//----------------------------------------------------------------------------
		return UpdateDependent(parentObj)

	} else {
		return parentRequestResult
	}

}


