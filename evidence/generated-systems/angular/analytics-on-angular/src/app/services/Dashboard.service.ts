import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Dashboard} from '../models/Dashboard';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {VisualizationService} from '../services/Visualization.service';
import {ReportService} from '../services/Report.service';
import {DataSetService} from '../services/DataSet.service';
import {AlertService} from '../services/Alert.service';
import {BIQueryService} from '../services/BIQuery.service';
import {TagService} from '../services/Tag.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DashboardService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dashboard : Dashboard;

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
	// add a Dashboard
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDashboard(title, theme, Workspace, Visualizations, Reports, Datasets, Alerts, Queries, Tags, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Dashboard/create';
		const obj = {
			      		title: title,
      		theme: theme,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Visualizations: Visualizations != null && Visualizations.length > 0 ? Visualizations : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		Queries: Queries != null && Queries.length > 0 ? Queries : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Dashboard
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDashboard(title, theme, Workspace, Visualizations, Reports, Datasets, Alerts, Queries, Tags, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Dashboard/update/' + id;
		const obj = {
				      		title: title,
      		theme: theme,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Visualizations: Visualizations != null && Visualizations.length > 0 ? Visualizations : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Alerts: Alerts != null && Alerts.length > 0 ? Alerts : null,
      		Queries: Queries != null && Queries.length > 0 ? Queries : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Dashboard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDashboard(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Dashboard/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Dashboard
	// returns the results untouched as an Observable Dashboard
	// Dashboard model
	// delegates via URI
	//********************************************************************
	getDashboard(id) : Observable<Dashboard> {
		const uri_ = this.apiUrl + '/Dashboard/load/' + id;

		return this.http.get<Dashboard>(uri_);
	}
	
	//********************************************************************
	// gets all Dashboard
	// returns the results untouched as JSON representation of an
	// Observable array of Dashboard models
	// delegates via URI
	//********************************************************************
	getDashboards() : Observable<Dashboard[]> {
		const uri_ = this.apiUrl + '/Dashboard/';

		return this
			.http.get<Dashboard[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a Dashboard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( dashboardId, _workspaceId ): Observable<any> {

		// get the Dashboard from storage
		this.loadHelper( dashboardId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.dashboard.workspace = tmp;

	// save the Dashboard
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a Dashboard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( dashboardId ): Observable<any> {

		// get the Dashboard from storage
		this.loadHelper( dashboardId );

	// assign Workspace to null
	this.dashboard.workspace = null;

	// save the Dashboard
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more visualizationsIds as a Visualizations
	// to a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVisualizations( dashboardId, visualizationsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );

	// split on a comma with no spaces
	var idList = visualizationsIds.split(',')

	// iterate over array of visualizations ids
	idList.forEach(function (id) {
		// read the Visualization
		var visualization = new VisualizationService(this.http).getVisualization(id);
		// add the Visualization if not already assigned
		if ( this.dashboard.visualizations.indexOf(visualization) == -1 )
		this.dashboard.visualizations.push(visualization);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more visualizationsIds as a Visualizations
	// from a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVisualizations( dashboardId, visualizationsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );


	// split on a comma with no spaces
	var idList 					= visualizationsIds.split(',');
	var visualizations 	= this.dashboard.visualizations;

	if ( visualizations != null && visualizationsIds != null ) {

		// iterate over array of visualizations ids
		visualizations.forEach(function (obj) {
			if ( visualizationsIds.indexOf(obj._id) > -1 ) {
				// remove the Visualization
				this.dashboard.visualizations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reportsIds as a Reports
	// to a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReports( dashboardId, reportsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );

	// split on a comma with no spaces
	var idList = reportsIds.split(',')

	// iterate over array of reports ids
	idList.forEach(function (id) {
		// read the Report
		var report = new ReportService(this.http).getReport(id);
		// add the Report if not already assigned
		if ( this.dashboard.reports.indexOf(report) == -1 )
		this.dashboard.reports.push(report);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reportsIds as a Reports
	// from a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReports( dashboardId, reportsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );


	// split on a comma with no spaces
	var idList 					= reportsIds.split(',');
	var reports 	= this.dashboard.reports;

	if ( reports != null && reportsIds != null ) {

		// iterate over array of reports ids
		reports.forEach(function (obj) {
			if ( reportsIds.indexOf(obj._id) > -1 ) {
				// remove the Report
				this.dashboard.reports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( dashboardId, datasetsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.dashboard.datasets.indexOf(dataSet) == -1 )
		this.dashboard.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( dashboardId, datasetsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.dashboard.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.dashboard.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more alertsIds as a Alerts
	// to a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAlerts( dashboardId, alertsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );

	// split on a comma with no spaces
	var idList = alertsIds.split(',')

	// iterate over array of alerts ids
	idList.forEach(function (id) {
		// read the Alert
		var alert = new AlertService(this.http).getAlert(id);
		// add the Alert if not already assigned
		if ( this.dashboard.alerts.indexOf(alert) == -1 )
		this.dashboard.alerts.push(alert);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more alertsIds as a Alerts
	// from a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAlerts( dashboardId, alertsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );


	// split on a comma with no spaces
	var idList 					= alertsIds.split(',');
	var alerts 	= this.dashboard.alerts;

	if ( alerts != null && alertsIds != null ) {

		// iterate over array of alerts ids
		alerts.forEach(function (obj) {
			if ( alertsIds.indexOf(obj._id) > -1 ) {
				// remove the Alert
				this.dashboard.alerts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more queriesIds as a Queries
	// to a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQueries( dashboardId, queriesIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );

	// split on a comma with no spaces
	var idList = queriesIds.split(',')

	// iterate over array of queries ids
	idList.forEach(function (id) {
		// read the BIQuery
		var bIQuery = new BIQueryService(this.http).getBIQuery(id);
		// add the BIQuery if not already assigned
		if ( this.dashboard.queries.indexOf(bIQuery) == -1 )
		this.dashboard.queries.push(bIQuery);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more queriesIds as a Queries
	// from a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQueries( dashboardId, queriesIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );


	// split on a comma with no spaces
	var idList 					= queriesIds.split(',');
	var queries 	= this.dashboard.queries;

	if ( queries != null && queriesIds != null ) {

		// iterate over array of queries ids
		queries.forEach(function (obj) {
			if ( queriesIds.indexOf(obj._id) > -1 ) {
				// remove the BIQuery
				this.dashboard.queries.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more tagsIds as a Tags
	// to a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTags( dashboardId, tagsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );

	// split on a comma with no spaces
	var idList = tagsIds.split(',')

	// iterate over array of tags ids
	idList.forEach(function (id) {
		// read the Tag
		var tag = new TagService(this.http).getTag(id);
		// add the Tag if not already assigned
		if ( this.dashboard.tags.indexOf(tag) == -1 )
		this.dashboard.tags.push(tag);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tagsIds as a Tags
	// from a Dashboard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTags( dashboardId, tagsIds ): Observable<any> {

		// get the Dashboard
		this.loadHelper( dashboardId );


	// split on a comma with no spaces
	var idList 					= tagsIds.split(',');
	var tags 	= this.dashboard.tags;

	if ( tags != null && tagsIds != null ) {

		// iterate over array of tags ids
		tags.forEach(function (obj) {
			if ( tagsIds.indexOf(obj._id) > -1 ) {
				// remove the Tag
				this.dashboard.tags.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Dashboard
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Dashboard/update/' + this.dashboard;

	return  this.http.post(uri_, this.dashboard );
}

	//********************************************************************
	// loadHelper - internal helper to load a Dashboard
	//********************************************************************	
	loadHelper( id ) {
		this.getDashboard(id)
			.subscribe((res : Dashboard) => {
				this.dashboard = res;
			});
	}
}