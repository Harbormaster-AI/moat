import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Notebook} from '../models/Notebook';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {DataSetService} from '../services/DataSet.service';
import {ExperimentService} from '../services/Experiment.service';
import {BIQueryService} from '../services/BIQuery.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class NotebookService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	notebook : Notebook;

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
	// add a Notebook
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addNotebook(title, repository, Workspace, Datasets, Experiments, Queries, Language) : Observable<any> {
		const uri_ = this.apiUrl + '/Notebook/create';
		const obj = {
			      		title: title,
      		repository: repository,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Experiments: Experiments != null && Experiments.length > 0 ? Experiments : null,
      		Queries: Queries != null && Queries.length > 0 ? Queries : null,
			Language: Language
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Notebook
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateNotebook(title, repository, Workspace, Datasets, Experiments, Queries, Language, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Notebook/update/' + id;
		const obj = {
				      		title: title,
      		repository: repository,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Experiments: Experiments != null && Experiments.length > 0 ? Experiments : null,
      		Queries: Queries != null && Queries.length > 0 ? Queries : null,
			Language: Language
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Notebook
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteNotebook(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Notebook/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Notebook
	// returns the results untouched as an Observable Notebook
	// Notebook model
	// delegates via URI
	//********************************************************************
	getNotebook(id) : Observable<Notebook> {
		const uri_ = this.apiUrl + '/Notebook/load/' + id;

		return this.http.get<Notebook>(uri_);
	}
	
	//********************************************************************
	// gets all Notebook
	// returns the results untouched as JSON representation of an
	// Observable array of Notebook models
	// delegates via URI
	//********************************************************************
	getNotebooks() : Observable<Notebook[]> {
		const uri_ = this.apiUrl + '/Notebook/';

		return this
			.http.get<Notebook[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a Notebook
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( notebookId, _workspaceId ): Observable<any> {

		// get the Notebook from storage
		this.loadHelper( notebookId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.notebook.workspace = tmp;

	// save the Notebook
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a Notebook
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( notebookId ): Observable<any> {

		// get the Notebook from storage
		this.loadHelper( notebookId );

	// assign Workspace to null
	this.notebook.workspace = null;

	// save the Notebook
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a Notebook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( notebookId, datasetsIds ): Observable<any> {

		// get the Notebook
		this.loadHelper( notebookId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.notebook.datasets.indexOf(dataSet) == -1 )
		this.notebook.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a Notebook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( notebookId, datasetsIds ): Observable<any> {

		// get the Notebook
		this.loadHelper( notebookId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.notebook.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.notebook.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more experimentsIds as a Experiments
	// to a Notebook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addExperiments( notebookId, experimentsIds ): Observable<any> {

		// get the Notebook
		this.loadHelper( notebookId );

	// split on a comma with no spaces
	var idList = experimentsIds.split(',')

	// iterate over array of experiments ids
	idList.forEach(function (id) {
		// read the Experiment
		var experiment = new ExperimentService(this.http).getExperiment(id);
		// add the Experiment if not already assigned
		if ( this.notebook.experiments.indexOf(experiment) == -1 )
		this.notebook.experiments.push(experiment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more experimentsIds as a Experiments
	// from a Notebook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeExperiments( notebookId, experimentsIds ): Observable<any> {

		// get the Notebook
		this.loadHelper( notebookId );


	// split on a comma with no spaces
	var idList 					= experimentsIds.split(',');
	var experiments 	= this.notebook.experiments;

	if ( experiments != null && experimentsIds != null ) {

		// iterate over array of experiments ids
		experiments.forEach(function (obj) {
			if ( experimentsIds.indexOf(obj._id) > -1 ) {
				// remove the Experiment
				this.notebook.experiments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more queriesIds as a Queries
	// to a Notebook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addQueries( notebookId, queriesIds ): Observable<any> {

		// get the Notebook
		this.loadHelper( notebookId );

	// split on a comma with no spaces
	var idList = queriesIds.split(',')

	// iterate over array of queries ids
	idList.forEach(function (id) {
		// read the BIQuery
		var bIQuery = new BIQueryService(this.http).getBIQuery(id);
		// add the BIQuery if not already assigned
		if ( this.notebook.queries.indexOf(bIQuery) == -1 )
		this.notebook.queries.push(bIQuery);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more queriesIds as a Queries
	// from a Notebook
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeQueries( notebookId, queriesIds ): Observable<any> {

		// get the Notebook
		this.loadHelper( notebookId );


	// split on a comma with no spaces
	var idList 					= queriesIds.split(',');
	var queries 	= this.notebook.queries;

	if ( queries != null && queriesIds != null ) {

		// iterate over array of queries ids
		queries.forEach(function (obj) {
			if ( queriesIds.indexOf(obj._id) > -1 ) {
				// remove the BIQuery
				this.notebook.queries.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Notebook
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Notebook/update/' + this.notebook;

	return  this.http.post(uri_, this.notebook );
}

	//********************************************************************
	// loadHelper - internal helper to load a Notebook
	//********************************************************************	
	loadHelper( id ) {
		this.getNotebook(id)
			.subscribe((res : Notebook) => {
				this.notebook = res;
			});
	}
}