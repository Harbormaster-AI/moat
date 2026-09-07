package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BenefitPlanDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBenefitPlan - creates a new db entry
//----------------------------------------------------------------------------
func CreateBenefitPlan(obj model.BenefitPlan)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BenefitPlan with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BenefitPlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBenefitPlan", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBenefitPlan - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBenefitPlan(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BenefitPlan

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BenefitPlan with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BenefitPlan using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BenefitPlan using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBenefitPlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBenefitPlan - returns all
//----------------------------------------------------------------------------
func GetAllBenefitPlan()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BenefitPlan

	//----------------------------------------------------------------------------
	// Request the ORM to find all BenefitPlan
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BenefitPlan" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BenefitPlan", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBenefitPlan", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBenefitPlan - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBenefitPlan(obj model.BenefitPlan)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BenefitPlan using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BenefitPlan using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBenefitPlan", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBenefitPlan - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBenefitPlan(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BenefitPlan with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBenefitPlan(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BenefitPlan)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BenefitPlan using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BenefitPlan using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBenefitPlan", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a BenefitPlan
//----------------------------------------------------------------------------
func AssignOrganizationToBenefitPlan( benefitPlanId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BenefitPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitPlan(benefitPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitPlan)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the BenefitPlan
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the BenefitPlan
			//----------------------------------------------------------------------------
			return UpdateBenefitPlan(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a BenefitPlan
//----------------------------------------------------------------------------
func UnassignOrganizationFromBenefitPlan(benefitPlanId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BenefitPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitPlan(benefitPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitPlan)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the BenefitPlan
		//----------------------------------------------------------------------------
		return UpdateBenefitPlan(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more enrollmentsIds as a Enrollments to a BenefitPlan
//----------------------------------------------------------------------------
func AddEnrollmentsToBenefitPlan ( benefitPlanId uint64, enrollmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BenefitPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitPlan(benefitPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( enrollmentsIds, ",")

		for _, enrollmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BenefitEnrollment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BenefitEnrollment
			// with a matching enrollmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , enrollmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Enrollments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Enrollments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Enrollments", enrollmentsId )
				return utils.RequestResult{false, msg, "unassignEnrollments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BenefitPlan from the gorm
		//----------------------------------------------------------------------------
		return GetBenefitPlan(benefitPlanId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more enrollmentsIds as a Enrollments from a BenefitPlan
//----------------------------------------------------------------------------
func RemoveEnrollmentsFromBenefitPlan( benefitPlanId uint64, enrollmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BenefitPlan with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBenefitPlan(benefitPlanId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BenefitPlan so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BenefitPlan)

		// slice the ids on comma with no spaces
		ids := strings.Split( enrollmentsIds, ",")

		for _, enrollmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BenefitEnrollment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BenefitEnrollment
			// with a matching enrollmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , enrollmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BenefitEnrollmentObj from the Enrollments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Enrollments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Enrollments", enrollmentsId )
				return utils.RequestResult{false, msg, "removeEnrollments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BenefitPlan from the gorm
		//----------------------------------------------------------------------------
		return GetBenefitPlan(benefitPlanId)

	} else {
		return parentRequestResult
	}
}

