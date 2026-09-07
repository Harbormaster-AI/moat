import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Experiment} from '../models/Experiment';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {TrainingRunService} from '../services/TrainingRun.service';
import {Model_Service} from '../services/Model_.service';
import {NotebookService} from '../services/Notebook.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ExperimentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	experiment : Experiment;

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
	// add a Experiment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addExperiment(name, objective, Workspace, TrainingRuns, Models, Notebooks, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Experiment/create';
		const obj = {
			      		name: name,
      		objective: objective,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		TrainingRuns: TrainingRuns != null && TrainingRuns.length > 0 ? TrainingRuns : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		Notebooks: Notebooks != null && Notebooks.length > 0 ? Notebooks : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Experiment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateExperiment(name, objective, Workspace, TrainingRuns, Models, Notebooks, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Experiment/update/' + id;
		const obj = {
				      		name: name,
      		objective: objective,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		TrainingRuns: TrainingRuns != null && TrainingRuns.length > 0 ? TrainingRuns : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		Notebooks: Notebooks != null && Notebooks.length > 0 ? Notebooks : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Experiment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteExperiment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Experiment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Experiment
	// returns the results untouched as an Observable Experiment
	// Experiment model
	// delegates via URI
	//********************************************************************
	getExperiment(id) : Observable<Experiment> {
		const uri_ = this.apiUrl + '/Experiment/load/' + id;

		return this.http.get<Experiment>(uri_);
	}
	
	//********************************************************************
	// gets all Experiment
	// returns the results untouched as JSON representation of an
	// Observable array of Experiment models
	// delegates via URI
	//********************************************************************
	getExperiments() : Observable<Experiment[]> {
		const uri_ = this.apiUrl + '/Experiment/';

		return this
			.http.get<Experiment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a Experiment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( experimentId, _workspaceId ): Observable<any> {

		// get the Experiment from storage
		this.loadHelper( experimentId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.experiment.workspace = tmp;

	// save the Experiment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a Experiment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( experimentId ): Observable<any> {

		// get the Experiment from storage
		this.loadHelper( experimentId );

	// assign Workspace to null
	this.experiment.workspace = null;

	// save the Experiment
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more trainingRunsIds as a TrainingRuns
	// to a Experiment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTrainingRuns( experimentId, trainingRunsIds ): Observable<any> {

		// get the Experiment
		this.loadHelper( experimentId );

	// split on a comma with no spaces
	var idList = trainingRunsIds.split(',')

	// iterate over array of trainingRuns ids
	idList.forEach(function (id) {
		// read the TrainingRun
		var trainingRun = new TrainingRunService(this.http).getTrainingRun(id);
		// add the TrainingRun if not already assigned
		if ( this.experiment.trainingRuns.indexOf(trainingRun) == -1 )
		this.experiment.trainingRuns.push(trainingRun);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more trainingRunsIds as a TrainingRuns
	// from a Experiment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTrainingRuns( experimentId, trainingRunsIds ): Observable<any> {

		// get the Experiment
		this.loadHelper( experimentId );


	// split on a comma with no spaces
	var idList 					= trainingRunsIds.split(',');
	var trainingRuns 	= this.experiment.trainingRuns;

	if ( trainingRuns != null && trainingRunsIds != null ) {

		// iterate over array of trainingRuns ids
		trainingRuns.forEach(function (obj) {
			if ( trainingRunsIds.indexOf(obj._id) > -1 ) {
				// remove the TrainingRun
				this.experiment.trainingRuns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more modelsIds as a Models
	// to a Experiment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModels( experimentId, modelsIds ): Observable<any> {

		// get the Experiment
		this.loadHelper( experimentId );

	// split on a comma with no spaces
	var idList = modelsIds.split(',')

	// iterate over array of models ids
	idList.forEach(function (id) {
		// read the Model_
		var model_ = new Model_Service(this.http).getModel_(id);
		// add the Model_ if not already assigned
		if ( this.experiment.models.indexOf(model_) == -1 )
		this.experiment.models.push(model_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelsIds as a Models
	// from a Experiment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModels( experimentId, modelsIds ): Observable<any> {

		// get the Experiment
		this.loadHelper( experimentId );


	// split on a comma with no spaces
	var idList 					= modelsIds.split(',');
	var models 	= this.experiment.models;

	if ( models != null && modelsIds != null ) {

		// iterate over array of models ids
		models.forEach(function (obj) {
			if ( modelsIds.indexOf(obj._id) > -1 ) {
				// remove the Model_
				this.experiment.models.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more notebooksIds as a Notebooks
	// to a Experiment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addNotebooks( experimentId, notebooksIds ): Observable<any> {

		// get the Experiment
		this.loadHelper( experimentId );

	// split on a comma with no spaces
	var idList = notebooksIds.split(',')

	// iterate over array of notebooks ids
	idList.forEach(function (id) {
		// read the Notebook
		var notebook = new NotebookService(this.http).getNotebook(id);
		// add the Notebook if not already assigned
		if ( this.experiment.notebooks.indexOf(notebook) == -1 )
		this.experiment.notebooks.push(notebook);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more notebooksIds as a Notebooks
	// from a Experiment
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeNotebooks( experimentId, notebooksIds ): Observable<any> {

		// get the Experiment
		this.loadHelper( experimentId );


	// split on a comma with no spaces
	var idList 					= notebooksIds.split(',');
	var notebooks 	= this.experiment.notebooks;

	if ( notebooks != null && notebooksIds != null ) {

		// iterate over array of notebooks ids
		notebooks.forEach(function (obj) {
			if ( notebooksIds.indexOf(obj._id) > -1 ) {
				// remove the Notebook
				this.experiment.notebooks.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Experiment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Experiment/update/' + this.experiment;

	return  this.http.post(uri_, this.experiment );
}

	//********************************************************************
	// loadHelper - internal helper to load a Experiment
	//********************************************************************	
	loadHelper( id ) {
		this.getExperiment(id)
			.subscribe((res : Experiment) => {
				this.experiment = res;
			});
	}
}