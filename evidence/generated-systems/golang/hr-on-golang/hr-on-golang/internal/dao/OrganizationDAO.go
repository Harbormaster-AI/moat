package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OrganizationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOrganization - creates a new db entry
//----------------------------------------------------------------------------
func CreateOrganization(obj model.Organization)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Organization with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Organization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOrganization", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOrganization - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOrganization(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Organization

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Organization with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Organization using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Organization using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOrganization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOrganization - returns all
//----------------------------------------------------------------------------
func GetAllOrganization()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Organization

	//----------------------------------------------------------------------------
	// Request the ORM to find all Organization
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Organization" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Organization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOrganization", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOrganization - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOrganization(obj model.Organization)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Organization using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Organization using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOrganization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOrganization - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOrganization(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOrganization(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Organization)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Organization using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Organization using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOrganization", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more departmentsIds as a Departments to a Organization
//----------------------------------------------------------------------------
func AddDepartmentsToOrganization ( organizationId uint64, departmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( departmentsIds, ",")

		for _, departmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Department

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Department
			// with a matching departmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , departmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Departments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Departments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Departments", departmentsId )
				return utils.RequestResult{false, msg, "unassignDepartments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more departmentsIds as a Departments from a Organization
//----------------------------------------------------------------------------
func RemoveDepartmentsFromOrganization( organizationId uint64, departmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( departmentsIds, ",")

		for _, departmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Department

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Department
			// with a matching departmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , departmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DepartmentObj from the Departments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Departments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Departments", departmentsId )
				return utils.RequestResult{false, msg, "removeDepartments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more locationsIds as a Locations to a Organization
//----------------------------------------------------------------------------
func AddLocationsToOrganization ( organizationId uint64, locationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( locationsIds, ",")

		for _, locationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Location

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Location
			// with a matching locationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , locationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Locations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Locations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Locations", locationsId )
				return utils.RequestResult{false, msg, "unassignLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more locationsIds as a Locations from a Organization
//----------------------------------------------------------------------------
func RemoveLocationsFromOrganization( organizationId uint64, locationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( locationsIds, ",")

		for _, locationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Location

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Location
			// with a matching locationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , locationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LocationObj from the Locations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Locations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Locations", locationsId )
				return utils.RequestResult{false, msg, "removeLocations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more jobFamiliesIds as a JobFamilies to a Organization
//----------------------------------------------------------------------------
func AddJobFamiliesToOrganization ( organizationId uint64, jobFamiliesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( jobFamiliesIds, ",")

		for _, jobFamiliesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.JobFamily

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a JobFamily
			// with a matching jobFamiliesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , jobFamiliesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the JobFamilies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("JobFamilies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "JobFamilies", jobFamiliesId )
				return utils.RequestResult{false, msg, "unassignJobFamilies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more jobFamiliesIds as a JobFamilies from a Organization
//----------------------------------------------------------------------------
func RemoveJobFamiliesFromOrganization( organizationId uint64, jobFamiliesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( jobFamiliesIds, ",")

		for _, jobFamiliesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.JobFamily

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a JobFamily
			// with a matching jobFamiliesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , jobFamiliesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove JobFamilyObj from the JobFamilies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("JobFamilies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "JobFamilies", jobFamiliesId )
				return utils.RequestResult{false, msg, "removeJobFamilies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more benefitPlansIds as a BenefitPlans to a Organization
//----------------------------------------------------------------------------
func AddBenefitPlansToOrganization ( organizationId uint64, benefitPlansIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( benefitPlansIds, ",")

		for _, benefitPlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BenefitPlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BenefitPlan
			// with a matching benefitPlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , benefitPlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the BenefitPlans using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BenefitPlans").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BenefitPlans", benefitPlansId )
				return utils.RequestResult{false, msg, "unassignBenefitPlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more benefitPlansIds as a BenefitPlans from a Organization
//----------------------------------------------------------------------------
func RemoveBenefitPlansFromOrganization( organizationId uint64, benefitPlansIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( benefitPlansIds, ",")

		for _, benefitPlansId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BenefitPlan

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BenefitPlan
			// with a matching benefitPlansId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , benefitPlansId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BenefitPlanObj from the BenefitPlans array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BenefitPlans").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BenefitPlans", benefitPlansId )
				return utils.RequestResult{false, msg, "removeBenefitPlans", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more costCentersIds as a CostCenters to a Organization
//----------------------------------------------------------------------------
func AddCostCentersToOrganization ( organizationId uint64, costCentersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( costCentersIds, ",")

		for _, costCentersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CostCenter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CostCenter
			// with a matching costCentersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , costCentersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CostCenters using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CostCenters").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CostCenters", costCentersId )
				return utils.RequestResult{false, msg, "unassignCostCenters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more costCentersIds as a CostCenters from a Organization
//----------------------------------------------------------------------------
func RemoveCostCentersFromOrganization( organizationId uint64, costCentersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( costCentersIds, ",")

		for _, costCentersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CostCenter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CostCenter
			// with a matching costCentersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , costCentersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CostCenterObj from the CostCenters array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CostCenters").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CostCenters", costCentersId )
				return utils.RequestResult{false, msg, "removeCostCenters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more payrollCalendarsIds as a PayrollCalendars to a Organization
//----------------------------------------------------------------------------
func AddPayrollCalendarsToOrganization ( organizationId uint64, payrollCalendarsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( payrollCalendarsIds, ",")

		for _, payrollCalendarsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PayrollCalendar

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PayrollCalendar
			// with a matching payrollCalendarsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , payrollCalendarsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the PayrollCalendars using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PayrollCalendars").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PayrollCalendars", payrollCalendarsId )
				return utils.RequestResult{false, msg, "unassignPayrollCalendars", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more payrollCalendarsIds as a PayrollCalendars from a Organization
//----------------------------------------------------------------------------
func RemovePayrollCalendarsFromOrganization( organizationId uint64, payrollCalendarsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Organization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOrganization(organizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Organization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Organization)

		// slice the ids on comma with no spaces
		ids := strings.Split( payrollCalendarsIds, ",")

		for _, payrollCalendarsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.PayrollCalendar

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a PayrollCalendar
			// with a matching payrollCalendarsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , payrollCalendarsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PayrollCalendarObj from the PayrollCalendars array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("PayrollCalendars").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "PayrollCalendars", payrollCalendarsId )
				return utils.RequestResult{false, msg, "removePayrollCalendars", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Organization from the gorm
		//----------------------------------------------------------------------------
		return GetOrganization(organizationId)

	} else {
		return parentRequestResult
	}
}

