import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TrainingRun} from '../models/TrainingRun';
import {ExperimentService} from '../services/Experiment.service';
import {ModelVersionService} from '../services/ModelVersion.service';
import {DataSetService} from '../services/DataSet.service';
import {FeatureService} from '../services/Feature.service';
import {RunMetricService} from '../services/RunMetric.service';
import {RunParameterService} from '../services/RunParameter.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TrainingRunService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	trainingRun : TrainingRun;

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
	// add a TrainingRun
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTrainingRun(runLabel, startedAt, completedAt, Experiment, ModelVersion, InputDatasets, Features, RunMetrics, RunParameters, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/TrainingRun/create';
		const obj = {
			      		runLabel: runLabel,
      		startedAt: startedAt,
      		completedAt: completedAt,
      		Experiment: Experiment != null && Experiment.length > 0 ? Experiment : null,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
      		InputDatasets: InputDatasets != null && InputDatasets.length > 0 ? InputDatasets : null,
      		Features: Features != null && Features.length > 0 ? Features : null,
      		RunMetrics: RunMetrics != null && RunMetrics.length > 0 ? RunMetrics : null,
      		RunParameters: RunParameters != null && RunParameters.length > 0 ? RunParameters : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TrainingRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTrainingRun(runLabel, startedAt, completedAt, Experiment, ModelVersion, InputDatasets, Features, RunMetrics, RunParameters, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TrainingRun/update/' + id;
		const obj = {
				      		runLabel: runLabel,
      		startedAt: startedAt,
      		completedAt: completedAt,
      		Experiment: Experiment != null && Experiment.length > 0 ? Experiment : null,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
      		InputDatasets: InputDatasets != null && InputDatasets.length > 0 ? InputDatasets : null,
      		Features: Features != null && Features.length > 0 ? Features : null,
      		RunMetrics: RunMetrics != null && RunMetrics.length > 0 ? RunMetrics : null,
      		RunParameters: RunParameters != null && RunParameters.length > 0 ? RunParameters : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TrainingRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTrainingRun(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TrainingRun/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TrainingRun
	// returns the results untouched as an Observable TrainingRun
	// TrainingRun model
	// delegates via URI
	//********************************************************************
	getTrainingRun(id) : Observable<TrainingRun> {
		const uri_ = this.apiUrl + '/TrainingRun/load/' + id;

		return this.http.get<TrainingRun>(uri_);
	}
	
	//********************************************************************
	// gets all TrainingRun
	// returns the results untouched as JSON representation of an
	// Observable array of TrainingRun models
	// delegates via URI
	//********************************************************************
	getTrainingRuns() : Observable<TrainingRun[]> {
		const uri_ = this.apiUrl + '/TrainingRun/';

		return this
			.http.get<TrainingRun[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Experiment on a TrainingRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignExperiment( trainingRunId, _experimentId ): Observable<any> {

		// get the TrainingRun from storage
		this.loadHelper( trainingRunId );

	// get the Experiment from storage
	var tmp 	= new ExperimentService(this.http).getExperiment(_experimentId);

	// assign the Experiment
	this.trainingRun.experiment = tmp;

	// save the TrainingRun
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Experiment on a TrainingRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignExperiment( trainingRunId ): Observable<any> {

		// get the TrainingRun from storage
		this.loadHelper( trainingRunId );

	// assign Experiment to null
	this.trainingRun.experiment = null;

	// save the TrainingRun
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ModelVersion on a TrainingRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignModelVersion( trainingRunId, _modelVersionId ): Observable<any> {

		// get the TrainingRun from storage
		this.loadHelper( trainingRunId );

	// get the ModelVersion from storage
	var tmp 	= new ModelVersionService(this.http).getModelVersion(_modelVersionId);

	// assign the ModelVersion
	this.trainingRun.modelVersion = tmp;

	// save the TrainingRun
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ModelVersion on a TrainingRun
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignModelVersion( trainingRunId ): Observable<any> {

		// get the TrainingRun from storage
		this.loadHelper( trainingRunId );

	// assign ModelVersion to null
	this.trainingRun.modelVersion = null;

	// save the TrainingRun
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more inputDatasetsIds as a InputDatasets
	// to a TrainingRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInputDatasets( trainingRunId, inputDatasetsIds ): Observable<any> {

		// get the TrainingRun
		this.loadHelper( trainingRunId );

	// split on a comma with no spaces
	var idList = inputDatasetsIds.split(',')

	// iterate over array of inputDatasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.trainingRun.inputDatasets.indexOf(dataSet) == -1 )
		this.trainingRun.inputDatasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more inputDatasetsIds as a InputDatasets
	// from a TrainingRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInputDatasets( trainingRunId, inputDatasetsIds ): Observable<any> {

		// get the TrainingRun
		this.loadHelper( trainingRunId );


	// split on a comma with no spaces
	var idList 					= inputDatasetsIds.split(',');
	var inputDatasets 	= this.trainingRun.inputDatasets;

	if ( inputDatasets != null && inputDatasetsIds != null ) {

		// iterate over array of inputDatasets ids
		inputDatasets.forEach(function (obj) {
			if ( inputDatasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.trainingRun.inputDatasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more featuresIds as a Features
	// to a TrainingRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFeatures( trainingRunId, featuresIds ): Observable<any> {

		// get the TrainingRun
		this.loadHelper( trainingRunId );

	// split on a comma with no spaces
	var idList = featuresIds.split(',')

	// iterate over array of features ids
	idList.forEach(function (id) {
		// read the Feature
		var feature = new FeatureService(this.http).getFeature(id);
		// add the Feature if not already assigned
		if ( this.trainingRun.features.indexOf(feature) == -1 )
		this.trainingRun.features.push(feature);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more featuresIds as a Features
	// from a TrainingRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFeatures( trainingRunId, featuresIds ): Observable<any> {

		// get the TrainingRun
		this.loadHelper( trainingRunId );


	// split on a comma with no spaces
	var idList 					= featuresIds.split(',');
	var features 	= this.trainingRun.features;

	if ( features != null && featuresIds != null ) {

		// iterate over array of features ids
		features.forEach(function (obj) {
			if ( featuresIds.indexOf(obj._id) > -1 ) {
				// remove the Feature
				this.trainingRun.features.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more runMetricsIds as a RunMetrics
	// to a TrainingRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRunMetrics( trainingRunId, runMetricsIds ): Observable<any> {

		// get the TrainingRun
		this.loadHelper( trainingRunId );

	// split on a comma with no spaces
	var idList = runMetricsIds.split(',')

	// iterate over array of runMetrics ids
	idList.forEach(function (id) {
		// read the RunMetric
		var runMetric = new RunMetricService(this.http).getRunMetric(id);
		// add the RunMetric if not already assigned
		if ( this.trainingRun.runMetrics.indexOf(runMetric) == -1 )
		this.trainingRun.runMetrics.push(runMetric);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more runMetricsIds as a RunMetrics
	// from a TrainingRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRunMetrics( trainingRunId, runMetricsIds ): Observable<any> {

		// get the TrainingRun
		this.loadHelper( trainingRunId );


	// split on a comma with no spaces
	var idList 					= runMetricsIds.split(',');
	var runMetrics 	= this.trainingRun.runMetrics;

	if ( runMetrics != null && runMetricsIds != null ) {

		// iterate over array of runMetrics ids
		runMetrics.forEach(function (obj) {
			if ( runMetricsIds.indexOf(obj._id) > -1 ) {
				// remove the RunMetric
				this.trainingRun.runMetrics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more runParametersIds as a RunParameters
	// to a TrainingRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRunParameters( trainingRunId, runParametersIds ): Observable<any> {

		// get the TrainingRun
		this.loadHelper( trainingRunId );

	// split on a comma with no spaces
	var idList = runParametersIds.split(',')

	// iterate over array of runParameters ids
	idList.forEach(function (id) {
		// read the RunParameter
		var runParameter = new RunParameterService(this.http).getRunParameter(id);
		// add the RunParameter if not already assigned
		if ( this.trainingRun.runParameters.indexOf(runParameter) == -1 )
		this.trainingRun.runParameters.push(runParameter);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more runParametersIds as a RunParameters
	// from a TrainingRun
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRunParameters( trainingRunId, runParametersIds ): Observable<any> {

		// get the TrainingRun
		this.loadHelper( trainingRunId );


	// split on a comma with no spaces
	var idList 					= runParametersIds.split(',');
	var runParameters 	= this.trainingRun.runParameters;

	if ( runParameters != null && runParametersIds != null ) {

		// iterate over array of runParameters ids
		runParameters.forEach(function (obj) {
			if ( runParametersIds.indexOf(obj._id) > -1 ) {
				// remove the RunParameter
				this.trainingRun.runParameters.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a TrainingRun
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TrainingRun/update/' + this.trainingRun;

	return  this.http.post(uri_, this.trainingRun );
}

	//********************************************************************
	// loadHelper - internal helper to load a TrainingRun
	//********************************************************************	
	loadHelper( id ) {
		this.getTrainingRun(id)
			.subscribe((res : TrainingRun) => {
				this.trainingRun = res;
			});
	}
}