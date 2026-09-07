import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {AccessPolicy} from '../models/AccessPolicy';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {DataSetService} from '../services/DataSet.service';
import {DashboardService} from '../services/Dashboard.service';
import {ReportService} from '../services/Report.service';
import {Model_Service} from '../services/Model_.service';
import {FeatureSetService} from '../services/FeatureSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class AccessPolicyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	accessPolicy : AccessPolicy;

	//********************************************************************
	// Catch all for the return value of a service call
	//********************************************************************
	result: any;

	//********************************************************************
	// sole constructor, injected with the HttpClient
	//********************************************************************
	constructor(private http: HttpClient) {
		super();
	}

		//********************************************************************
	// add a AccessPolicy
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addAccessPolicy(name, subjectName, Workspace, Datasets, Dashboards, Reports, Models, FeatureSets, AccessLevel, SubjectType) : Observable<any> {
		const uri_ = this.apiUrl + '/AccessPolicy/create';
		const obj = {
			      		name: name,
      		subjectName: subjectName,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		AccessLevel: AccessLevel,
			SubjectType: SubjectType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a AccessPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateAccessPolicy(name, subjectName, Workspace, Datasets, Dashboards, Reports, Models, FeatureSets, AccessLevel, SubjectType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/AccessPolicy/update/' + id;
		const obj = {
				      		name: name,
      		subjectName: subjectName,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		AccessLevel: AccessLevel,
			SubjectType: SubjectType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a AccessPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteAccessPolicy(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/AccessPolicy/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a AccessPolicy
	// returns the results untouched as an Observable AccessPolicy
	// AccessPolicy model
	// delegates via URI
	//********************************************************************
	getAccessPolicy(id) : Observable<AccessPolicy> {
		const uri_ = this.apiUrl + '/AccessPolicy/load/' + id;

		return this.http.get<AccessPolicy>(uri_);
	}
	
	//********************************************************************
	// gets all AccessPolicy
	// returns the results untouched as JSON representation of an
	// Observable array of AccessPolicy models
	// delegates via URI
	//********************************************************************
	getAccessPolicys() : Observable<AccessPolicy[]> {
		const uri_ = this.apiUrl + '/AccessPolicy/';

		return this
			.http.get<AccessPolicy[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a AccessPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( accessPolicyId, _workspaceId ): Observable<any> {

		// get the AccessPolicy from storage
		this.loadHelper( accessPolicyId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.accessPolicy.workspace = tmp;

	// save the AccessPolicy
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a AccessPolicy
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( accessPolicyId ): Observable<any> {

		// get the AccessPolicy from storage
		this.loadHelper( accessPolicyId );

	// assign Workspace to null
	this.accessPolicy.workspace = null;

	// save the AccessPolicy
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( accessPolicyId, datasetsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.accessPolicy.datasets.indexOf(dataSet) == -1 )
		this.accessPolicy.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( accessPolicyId, datasetsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.accessPolicy.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.accessPolicy.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dashboardsIds as a Dashboards
	// to a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDashboards( accessPolicyId, dashboardsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );

	// split on a comma with no spaces
	var idList = dashboardsIds.split(',')

	// iterate over array of dashboards ids
	idList.forEach(function (id) {
		// read the Dashboard
		var dashboard = new DashboardService(this.http).getDashboard(id);
		// add the Dashboard if not already assigned
		if ( this.accessPolicy.dashboards.indexOf(dashboard) == -1 )
		this.accessPolicy.dashboards.push(dashboard);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dashboardsIds as a Dashboards
	// from a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDashboards( accessPolicyId, dashboardsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );


	// split on a comma with no spaces
	var idList 					= dashboardsIds.split(',');
	var dashboards 	= this.accessPolicy.dashboards;

	if ( dashboards != null && dashboardsIds != null ) {

		// iterate over array of dashboards ids
		dashboards.forEach(function (obj) {
			if ( dashboardsIds.indexOf(obj._id) > -1 ) {
				// remove the Dashboard
				this.accessPolicy.dashboards.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reportsIds as a Reports
	// to a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReports( accessPolicyId, reportsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );

	// split on a comma with no spaces
	var idList = reportsIds.split(',')

	// iterate over array of reports ids
	idList.forEach(function (id) {
		// read the Report
		var report = new ReportService(this.http).getReport(id);
		// add the Report if not already assigned
		if ( this.accessPolicy.reports.indexOf(report) == -1 )
		this.accessPolicy.reports.push(report);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reportsIds as a Reports
	// from a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReports( accessPolicyId, reportsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );


	// split on a comma with no spaces
	var idList 					= reportsIds.split(',');
	var reports 	= this.accessPolicy.reports;

	if ( reports != null && reportsIds != null ) {

		// iterate over array of reports ids
		reports.forEach(function (obj) {
			if ( reportsIds.indexOf(obj._id) > -1 ) {
				// remove the Report
				this.accessPolicy.reports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more modelsIds as a Models
	// to a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModels( accessPolicyId, modelsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );

	// split on a comma with no spaces
	var idList = modelsIds.split(',')

	// iterate over array of models ids
	idList.forEach(function (id) {
		// read the Model_
		var model_ = new Model_Service(this.http).getModel_(id);
		// add the Model_ if not already assigned
		if ( this.accessPolicy.models.indexOf(model_) == -1 )
		this.accessPolicy.models.push(model_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelsIds as a Models
	// from a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModels( accessPolicyId, modelsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );


	// split on a comma with no spaces
	var idList 					= modelsIds.split(',');
	var models 	= this.accessPolicy.models;

	if ( models != null && modelsIds != null ) {

		// iterate over array of models ids
		models.forEach(function (obj) {
			if ( modelsIds.indexOf(obj._id) > -1 ) {
				// remove the Model_
				this.accessPolicy.models.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more featureSetsIds as a FeatureSets
	// to a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFeatureSets( accessPolicyId, featureSetsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );

	// split on a comma with no spaces
	var idList = featureSetsIds.split(',')

	// iterate over array of featureSets ids
	idList.forEach(function (id) {
		// read the FeatureSet
		var featureSet = new FeatureSetService(this.http).getFeatureSet(id);
		// add the FeatureSet if not already assigned
		if ( this.accessPolicy.featureSets.indexOf(featureSet) == -1 )
		this.accessPolicy.featureSets.push(featureSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more featureSetsIds as a FeatureSets
	// from a AccessPolicy
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFeatureSets( accessPolicyId, featureSetsIds ): Observable<any> {

		// get the AccessPolicy
		this.loadHelper( accessPolicyId );


	// split on a comma with no spaces
	var idList 					= featureSetsIds.split(',');
	var featureSets 	= this.accessPolicy.featureSets;

	if ( featureSets != null && featureSetsIds != null ) {

		// iterate over array of featureSets ids
		featureSets.forEach(function (obj) {
			if ( featureSetsIds.indexOf(obj._id) > -1 ) {
				// remove the FeatureSet
				this.accessPolicy.featureSets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a AccessPolicy
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/AccessPolicy/update/' + this.accessPolicy;

	return  this.http.post(uri_, this.accessPolicy );
}

	//********************************************************************
	// loadHelper - internal helper to load a AccessPolicy
	//********************************************************************	
	loadHelper( id ) {
		this.getAccessPolicy(id)
			.subscribe((res : AccessPolicy) => {
				this.accessPolicy = res;
			});
	}
}