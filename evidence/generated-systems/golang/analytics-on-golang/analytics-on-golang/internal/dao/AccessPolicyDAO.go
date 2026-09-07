package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AccessPolicyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAccessPolicy - creates a new db entry
//----------------------------------------------------------------------------
func CreateAccessPolicy(obj model.AccessPolicy)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AccessPolicy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AccessPolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAccessPolicy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAccessPolicy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAccessPolicy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AccessPolicy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AccessPolicy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AccessPolicy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AccessPolicy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAccessPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAccessPolicy - returns all
//----------------------------------------------------------------------------
func GetAllAccessPolicy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AccessPolicy

	//----------------------------------------------------------------------------
	// Request the ORM to find all AccessPolicy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AccessPolicy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AccessPolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAccessPolicy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAccessPolicy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAccessPolicy(obj model.AccessPolicy)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AccessPolicy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AccessPolicy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAccessPolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAccessPolicy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAccessPolicy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAccessPolicy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AccessPolicy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AccessPolicy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AccessPolicy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAccessPolicy", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a AccessPolicy
//----------------------------------------------------------------------------
func AssignWorkspaceToAccessPolicy( accessPolicyId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AnalyticsWorkspace

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AnalyticsWorkspace with a
		// matching workspaceId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workspaceId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Workspace	to the AccessPolicy
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the AccessPolicy
			//----------------------------------------------------------------------------
			return UpdateAccessPolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a AccessPolicy
//----------------------------------------------------------------------------
func UnassignWorkspaceFromAccessPolicy(accessPolicyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the AccessPolicy
		//----------------------------------------------------------------------------
		return UpdateAccessPolicy(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a AccessPolicy
//----------------------------------------------------------------------------
func AddDatasetsToAccessPolicy ( accessPolicyId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

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
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a AccessPolicy
//----------------------------------------------------------------------------
func RemoveDatasetsFromAccessPolicy( accessPolicyId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

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
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dashboardsIds as a Dashboards to a AccessPolicy
//----------------------------------------------------------------------------
func AddDashboardsToAccessPolicy ( accessPolicyId uint64, dashboardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( dashboardsIds, ",")

		for _, dashboardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dashboard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dashboard
			// with a matching dashboardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dashboardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Dashboards using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dashboards").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dashboards", dashboardsId )
				return utils.RequestResult{false, msg, "unassignDashboards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dashboardsIds as a Dashboards from a AccessPolicy
//----------------------------------------------------------------------------
func RemoveDashboardsFromAccessPolicy( accessPolicyId uint64, dashboardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( dashboardsIds, ",")

		for _, dashboardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dashboard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dashboard
			// with a matching dashboardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dashboardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DashboardObj from the Dashboards array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dashboards").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dashboards", dashboardsId )
				return utils.RequestResult{false, msg, "removeDashboards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reportsIds as a Reports to a AccessPolicy
//----------------------------------------------------------------------------
func AddReportsToAccessPolicy ( accessPolicyId uint64, reportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Report

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Report
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reports using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "unassignReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reportsIds as a Reports from a AccessPolicy
//----------------------------------------------------------------------------
func RemoveReportsFromAccessPolicy( accessPolicyId uint64, reportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Report

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Report
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ReportObj from the Reports array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "removeReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more modelsIds as a Models to a AccessPolicy
//----------------------------------------------------------------------------
func AddModelsToAccessPolicy ( accessPolicyId uint64, modelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

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
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelsIds as a Models from a AccessPolicy
//----------------------------------------------------------------------------
func RemoveModelsFromAccessPolicy( accessPolicyId uint64, modelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

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
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more featureSetsIds as a FeatureSets to a AccessPolicy
//----------------------------------------------------------------------------
func AddFeatureSetsToAccessPolicy ( accessPolicyId uint64, featureSetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( featureSetsIds, ",")

		for _, featureSetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeatureSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeatureSet
			// with a matching featureSetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featureSetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FeatureSets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeatureSets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeatureSets", featureSetsId )
				return utils.RequestResult{false, msg, "unassignFeatureSets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more featureSetsIds as a FeatureSets from a AccessPolicy
//----------------------------------------------------------------------------
func RemoveFeatureSetsFromAccessPolicy( accessPolicyId uint64, featureSetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AccessPolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAccessPolicy(accessPolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AccessPolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AccessPolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( featureSetsIds, ",")

		for _, featureSetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeatureSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeatureSet
			// with a matching featureSetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featureSetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FeatureSetObj from the FeatureSets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeatureSets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeatureSets", featureSetsId )
				return utils.RequestResult{false, msg, "removeFeatureSets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AccessPolicy from the gorm
		//----------------------------------------------------------------------------
		return GetAccessPolicy(accessPolicyId)

	} else {
		return parentRequestResult
	}
}

