import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BIQuery} from '../models/BIQuery';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {DataSetService} from '../services/DataSet.service';
import {ReportService} from '../services/Report.service';
import {DashboardService} from '../services/Dashboard.service';
import {NotebookService} from '../services/Notebook.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BIQueryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	bIQuery : BIQuery;

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
	// add a BIQuery
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBIQuery(name, text, Workspace, Datasets, Reports, Dashboards, Notebooks, Dialect) : Observable<any> {
		const uri_ = this.apiUrl + '/BIQuery/create';
		const obj = {
			      		name: name,
      		text: text,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Notebooks: Notebooks != null && Notebooks.length > 0 ? Notebooks : null,
			Dialect: Dialect
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BIQuery
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBIQuery(name, text, Workspace, Datasets, Reports, Dashboards, Notebooks, Dialect, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BIQuery/update/' + id;
		const obj = {
				      		name: name,
      		text: text,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Reports: Reports != null && Reports.length > 0 ? Reports : null,
      		Dashboards: Dashboards != null && Dashboards.length > 0 ? Dashboards : null,
      		Notebooks: Notebooks != null && Notebooks.length > 0 ? Notebooks : null,
			Dialect: Dialect
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BIQuery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBIQuery(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BIQuery/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BIQuery
	// returns the results untouched as an Observable BIQuery
	// BIQuery model
	// delegates via URI
	//********************************************************************
	getBIQuery(id) : Observable<BIQuery> {
		const uri_ = this.apiUrl + '/BIQuery/load/' + id;

		return this.http.get<BIQuery>(uri_);
	}
	
	//********************************************************************
	// gets all BIQuery
	// returns the results untouched as JSON representation of an
	// Observable array of BIQuery models
	// delegates via URI
	//********************************************************************
	getBIQuerys() : Observable<BIQuery[]> {
		const uri_ = this.apiUrl + '/BIQuery/';

		return this
			.http.get<BIQuery[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a BIQuery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( bIQueryId, _workspaceId ): Observable<any> {

		// get the BIQuery from storage
		this.loadHelper( bIQueryId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.bIQuery.workspace = tmp;

	// save the BIQuery
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a BIQuery
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( bIQueryId ): Observable<any> {

		// get the BIQuery from storage
		this.loadHelper( bIQueryId );

	// assign Workspace to null
	this.bIQuery.workspace = null;

	// save the BIQuery
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a BIQuery
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( bIQueryId, datasetsIds ): Observable<any> {

		// get the BIQuery
		this.loadHelper( bIQueryId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.bIQuery.datasets.indexOf(dataSet) == -1 )
		this.bIQuery.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a BIQuery
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( bIQueryId, datasetsIds ): Observable<any> {

		// get the BIQuery
		this.loadHelper( bIQueryId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.bIQuery.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.bIQuery.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more reportsIds as a Reports
	// to a BIQuery
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReports( bIQueryId, reportsIds ): Observable<any> {

		// get the BIQuery
		this.loadHelper( bIQueryId );

	// split on a comma with no spaces
	var idList = reportsIds.split(',')

	// iterate over array of reports ids
	idList.forEach(function (id) {
		// read the Report
		var report = new ReportService(this.http).getReport(id);
		// add the Report if not already assigned
		if ( this.bIQuery.reports.indexOf(report) == -1 )
		this.bIQuery.reports.push(report);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reportsIds as a Reports
	// from a BIQuery
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReports( bIQueryId, reportsIds ): Observable<any> {

		// get the BIQuery
		this.loadHelper( bIQueryId );


	// split on a comma with no spaces
	var idList 					= reportsIds.split(',');
	var reports 	= this.bIQuery.reports;

	if ( reports != null && reportsIds != null ) {

		// iterate over array of reports ids
		reports.forEach(function (obj) {
			if ( reportsIds.indexOf(obj._id) > -1 ) {
				// remove the Report
				this.bIQuery.reports.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more dashboardsIds as a Dashboards
	// to a BIQuery
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDashboards( bIQueryId, dashboardsIds ): Observable<any> {

		// get the BIQuery
		this.loadHelper( bIQueryId );

	// split on a comma with no spaces
	var idList = dashboardsIds.split(',')

	// iterate over array of dashboards ids
	idList.forEach(function (id) {
		// read the Dashboard
		var dashboard = new DashboardService(this.http).getDashboard(id);
		// add the Dashboard if not already assigned
		if ( this.bIQuery.dashboards.indexOf(dashboard) == -1 )
		this.bIQuery.dashboards.push(dashboard);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dashboardsIds as a Dashboards
	// from a BIQuery
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDashboards( bIQueryId, dashboardsIds ): Observable<any> {

		// get the BIQuery
		this.loadHelper( bIQueryId );


	// split on a comma with no spaces
	var idList 					= dashboardsIds.split(',');
	var dashboards 	= this.bIQuery.dashboards;

	if ( dashboards != null && dashboardsIds != null ) {

		// iterate over array of dashboards ids
		dashboards.forEach(function (obj) {
			if ( dashboardsIds.indexOf(obj._id) > -1 ) {
				// remove the Dashboard
				this.bIQuery.dashboards.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more notebooksIds as a Notebooks
	// to a BIQuery
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addNotebooks( bIQueryId, notebooksIds ): Observable<any> {

		// get the BIQuery
		this.loadHelper( bIQueryId );

	// split on a comma with no spaces
	var idList = notebooksIds.split(',')

	// iterate over array of notebooks ids
	idList.forEach(function (id) {
		// read the Notebook
		var notebook = new NotebookService(this.http).getNotebook(id);
		// add the Notebook if not already assigned
		if ( this.bIQuery.notebooks.indexOf(notebook) == -1 )
		this.bIQuery.notebooks.push(notebook);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more notebooksIds as a Notebooks
	// from a BIQuery
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeNotebooks( bIQueryId, notebooksIds ): Observable<any> {

		// get the BIQuery
		this.loadHelper( bIQueryId );


	// split on a comma with no spaces
	var idList 					= notebooksIds.split(',');
	var notebooks 	= this.bIQuery.notebooks;

	if ( notebooks != null && notebooksIds != null ) {

		// iterate over array of notebooks ids
		notebooks.forEach(function (obj) {
			if ( notebooksIds.indexOf(obj._id) > -1 ) {
				// remove the Notebook
				this.bIQuery.notebooks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BIQuery
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BIQuery/update/' + this.bIQuery;

	return  this.http.post(uri_, this.bIQuery );
}

	//********************************************************************
	// loadHelper - internal helper to load a BIQuery
	//********************************************************************	
	loadHelper( id ) {
		this.getBIQuery(id)
			.subscribe((res : BIQuery) => {
				this.bIQuery = res;
			});
	}
}