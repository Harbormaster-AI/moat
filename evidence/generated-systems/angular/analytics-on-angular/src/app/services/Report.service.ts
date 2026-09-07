import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Report} from '../models/Report';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {VisualizationService} from '../services/Visualization.service';
import {DataSetService} from '../services/DataSet.service';
import {SemanticModelService} from '../services/SemanticModel.service';
import {BIQueryService} from '../services/BIQuery.service';
import {TagService} from '../services/Tag.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ReportService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	report : Report;

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
	// add a Report
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addReport(title, audience, Workspace, Visualizations, Datasets, SemanticModels, Queries, Tags, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Report/create';
		const obj = {
			      		title: title,
      		audience: audience,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Visualizations: Visualizations != null && Visualizations.length > 0 ? Visualizations : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		SemanticModels: SemanticModels != null && SemanticModels.length > 0 ? SemanticModels : null,
      		Queries: Queries != null && Queries.length > 0 ? Queries : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Report
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateReport(title, audience, Workspace, Visualizations, Datasets, SemanticModels, Queries, Tags, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Report/update/' + id;
		const obj = {
				      		title: title,
      		audience: audience,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Visualizations: Visualizations != null && Visualizations.length > 0 ? Visualizations : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		SemanticModels: SemanticModels != null && SemanticModels.length > 0 ? SemanticModels : null,
      		Queries: Queries != null && Queries.length > 0 ? Queries : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Report
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteReport(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Report/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Report
	// returns the results untouched as an Observable Report
	// Report model
	// delegates via URI
	//********************************************************************
	getReport(id) : Observable<Report> {
		const uri_ = this.apiUrl + '/Report/load/' + id;

		return this.http.get<Report>(uri_);
	}
	
	//********************************************************************
	// gets all Report
	// returns the results untouched as JSON representation of an
	// Observable array of Report models
	// delegates via URI
	//********************************************************************
	getReports() : Observable<Report[]> {
		const uri_ = this.apiUrl + '/Report/';

		return this
			.http.get<Report[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a Report
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( reportId, _workspaceId ): Observable<any> {

		// get the Report from storage
		this.loadHelper( reportId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.report.workspace = tmp;

	// save the Report
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a Report
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( reportId ): Observable<any> {

		// get the Report from storage
		this.loadHelper( reportId );

	// assign Workspace to null
	this.report.workspace = null;

	// save the Report
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more visualizationsIds as a Visualizations
	// to a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVisualizations( reportId, visualizationsIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );

	// split on a comma with no spaces
	var idList = visualizationsIds.split(',')

	// iterate over array of visualizations ids
	idList.forEach(function (id) {
		// read the Visualization
		var visualization = new VisualizationService(this.http).getVisualization(id);
		// add the Visualization if not already assigned
		if ( this.report.visualizations.indexOf(visualization) == -1 )
		this.report.visualizations.push(visualization);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more visualizationsIds as a Visualizations
	// from a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVisualizations( reportId, visualizationsIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );


	// split on a comma with no spaces
	var idList 					= visualizationsIds.split(',');
	var visualizations 	= this.report.visualizations;

	if ( visualizations != null && visualizationsIds != null ) {

		// iterate over array of visualizations ids
		visualizations.forEach(function (obj) {
			if ( visualizationsIds.indexOf(obj._id) > -1 ) {
				// remove the Visualization
				this.report.visualizations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( reportId, datasetsIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.report.datasets.indexOf(dataSet) == -1 )
		this.report.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( reportId, datasetsIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.report.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.report.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more semanticModelsIds as a SemanticModels
	// to a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSemanticModels( reportId, semanticModelsIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );

	// split on a comma with no spaces
	var idList = semanticModelsIds.split(',')

	// iterate over array of semanticModels ids
	idList.forEach(function (id) {
		// read the SemanticModel
		var semanticModel = new SemanticModelService(this.http).getSemanticModel(id);
		// add the SemanticModel if not already assigned
		if ( this.report.semanticModels.indexOf(semanticModel) == -1 )
		this.report.semanticModels.push(semanticModel);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more semanticModelsIds as a SemanticModels
	// from a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSemanticModels( reportId, semanticModelsIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );


	// split on a comma with no spaces
	var idList 					= semanticModelsIds.split(',');
	var semanticModels 	= this.report.semanticModels;

	if ( semanticModels != null && semanticModelsIds != null ) {

		// iterate over array of semanticModels ids
		semanticModels.forEach(function (obj) {
			if ( semanticModelsIds.indexOf(obj._id) > -1 ) {
				// remove the SemanticModel
				this.report.semanticModels.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more queriesIds as a Queries
	// to a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQueries( reportId, queriesIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );

	// split on a comma with no spaces
	var idList = queriesIds.split(',')

	// iterate over array of queries ids
	idList.forEach(function (id) {
		// read the BIQuery
		var bIQuery = new BIQueryService(this.http).getBIQuery(id);
		// add the BIQuery if not already assigned
		if ( this.report.queries.indexOf(bIQuery) == -1 )
		this.report.queries.push(bIQuery);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more queriesIds as a Queries
	// from a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQueries( reportId, queriesIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );


	// split on a comma with no spaces
	var idList 					= queriesIds.split(',');
	var queries 	= this.report.queries;

	if ( queries != null && queriesIds != null ) {

		// iterate over array of queries ids
		queries.forEach(function (obj) {
			if ( queriesIds.indexOf(obj._id) > -1 ) {
				// remove the BIQuery
				this.report.queries.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more tagsIds as a Tags
	// to a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTags( reportId, tagsIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );

	// split on a comma with no spaces
	var idList = tagsIds.split(',')

	// iterate over array of tags ids
	idList.forEach(function (id) {
		// read the Tag
		var tag = new TagService(this.http).getTag(id);
		// add the Tag if not already assigned
		if ( this.report.tags.indexOf(tag) == -1 )
		this.report.tags.push(tag);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tagsIds as a Tags
	// from a Report
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTags( reportId, tagsIds ): Observable<any> {

		// get the Report
		this.loadHelper( reportId );


	// split on a comma with no spaces
	var idList 					= tagsIds.split(',');
	var tags 	= this.report.tags;

	if ( tags != null && tagsIds != null ) {

		// iterate over array of tags ids
		tags.forEach(function (obj) {
			if ( tagsIds.indexOf(obj._id) > -1 ) {
				// remove the Tag
				this.report.tags.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Report
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Report/update/' + this.report;

	return  this.http.post(uri_, this.report );
}

	//********************************************************************
	// loadHelper - internal helper to load a Report
	//********************************************************************	
	loadHelper( id ) {
		this.getReport(id)
			.subscribe((res : Report) => {
				this.report = res;
			});
	}
}