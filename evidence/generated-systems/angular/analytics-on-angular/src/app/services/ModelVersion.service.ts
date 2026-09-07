import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ModelVersion} from '../models/ModelVersion';
import {Model_Service} from '../services/Model_.service';
import {TrainingRunService} from '../services/TrainingRun.service';
import {EvaluationMetricService} from '../services/EvaluationMetric.service';
import {InferenceEndpointService} from '../services/InferenceEndpoint.service';
import {FeatureSetService} from '../services/FeatureSet.service';
import {DataSetService} from '../services/DataSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ModelVersionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	modelVersion : ModelVersion;

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
	// add a ModelVersion
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addModelVersion(version, Model_, TrainingRun, EvaluationMetrics, Deployments, FeatureSets, Datasets, Lifecycle, TrainingStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/ModelVersion/create';
		const obj = {
			      		version: version,
      		Model_: Model_ != null && Model_.length > 0 ? Model_ : null,
      		TrainingRun: TrainingRun != null && TrainingRun.length > 0 ? TrainingRun : null,
      		EvaluationMetrics: EvaluationMetrics != null && EvaluationMetrics.length > 0 ? EvaluationMetrics : null,
      		Deployments: Deployments != null && Deployments.length > 0 ? Deployments : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Lifecycle: Lifecycle,
			TrainingStatus: TrainingStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ModelVersion
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateModelVersion(version, Model_, TrainingRun, EvaluationMetrics, Deployments, FeatureSets, Datasets, Lifecycle, TrainingStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ModelVersion/update/' + id;
		const obj = {
				      		version: version,
      		Model_: Model_ != null && Model_.length > 0 ? Model_ : null,
      		TrainingRun: TrainingRun != null && TrainingRun.length > 0 ? TrainingRun : null,
      		EvaluationMetrics: EvaluationMetrics != null && EvaluationMetrics.length > 0 ? EvaluationMetrics : null,
      		Deployments: Deployments != null && Deployments.length > 0 ? Deployments : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Lifecycle: Lifecycle,
			TrainingStatus: TrainingStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ModelVersion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteModelVersion(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ModelVersion/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ModelVersion
	// returns the results untouched as an Observable ModelVersion
	// ModelVersion model
	// delegates via URI
	//********************************************************************
	getModelVersion(id) : Observable<ModelVersion> {
		const uri_ = this.apiUrl + '/ModelVersion/load/' + id;

		return this.http.get<ModelVersion>(uri_);
	}
	
	//********************************************************************
	// gets all ModelVersion
	// returns the results untouched as JSON representation of an
	// Observable array of ModelVersion models
	// delegates via URI
	//********************************************************************
	getModelVersions() : Observable<ModelVersion[]> {
		const uri_ = this.apiUrl + '/ModelVersion/';

		return this
			.http.get<ModelVersion[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Model_ on a ModelVersion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignModel_( modelVersionId, _model_Id ): Observable<any> {

		// get the ModelVersion from storage
		this.loadHelper( modelVersionId );

	// get the Model_ from storage
	var tmp 	= new Model_Service(this.http).getModel_(_model_Id);

	// assign the Model_
	this.modelVersion.model_ = tmp;

	// save the ModelVersion
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Model_ on a ModelVersion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignModel_( modelVersionId ): Observable<any> {

		// get the ModelVersion from storage
		this.loadHelper( modelVersionId );

	// assign Model_ to null
	this.modelVersion.model_ = null;

	// save the ModelVersion
	return this.saveHelper();
}

		//********************************************************************
	// assigns a TrainingRun on a ModelVersion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTrainingRun( modelVersionId, _trainingRunId ): Observable<any> {

		// get the ModelVersion from storage
		this.loadHelper( modelVersionId );

	// get the TrainingRun from storage
	var tmp 	= new TrainingRunService(this.http).getTrainingRun(_trainingRunId);

	// assign the TrainingRun
	this.modelVersion.trainingRun = tmp;

	// save the ModelVersion
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TrainingRun on a ModelVersion
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTrainingRun( modelVersionId ): Observable<any> {

		// get the ModelVersion from storage
		this.loadHelper( modelVersionId );

	// assign TrainingRun to null
	this.modelVersion.trainingRun = null;

	// save the ModelVersion
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more evaluationMetricsIds as a EvaluationMetrics
	// to a ModelVersion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEvaluationMetrics( modelVersionId, evaluationMetricsIds ): Observable<any> {

		// get the ModelVersion
		this.loadHelper( modelVersionId );

	// split on a comma with no spaces
	var idList = evaluationMetricsIds.split(',')

	// iterate over array of evaluationMetrics ids
	idList.forEach(function (id) {
		// read the EvaluationMetric
		var evaluationMetric = new EvaluationMetricService(this.http).getEvaluationMetric(id);
		// add the EvaluationMetric if not already assigned
		if ( this.modelVersion.evaluationMetrics.indexOf(evaluationMetric) == -1 )
		this.modelVersion.evaluationMetrics.push(evaluationMetric);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more evaluationMetricsIds as a EvaluationMetrics
	// from a ModelVersion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEvaluationMetrics( modelVersionId, evaluationMetricsIds ): Observable<any> {

		// get the ModelVersion
		this.loadHelper( modelVersionId );


	// split on a comma with no spaces
	var idList 					= evaluationMetricsIds.split(',');
	var evaluationMetrics 	= this.modelVersion.evaluationMetrics;

	if ( evaluationMetrics != null && evaluationMetricsIds != null ) {

		// iterate over array of evaluationMetrics ids
		evaluationMetrics.forEach(function (obj) {
			if ( evaluationMetricsIds.indexOf(obj._id) > -1 ) {
				// remove the EvaluationMetric
				this.modelVersion.evaluationMetrics.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more deploymentsIds as a Deployments
	// to a ModelVersion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDeployments( modelVersionId, deploymentsIds ): Observable<any> {

		// get the ModelVersion
		this.loadHelper( modelVersionId );

	// split on a comma with no spaces
	var idList = deploymentsIds.split(',')

	// iterate over array of deployments ids
	idList.forEach(function (id) {
		// read the InferenceEndpoint
		var inferenceEndpoint = new InferenceEndpointService(this.http).getInferenceEndpoint(id);
		// add the InferenceEndpoint if not already assigned
		if ( this.modelVersion.deployments.indexOf(inferenceEndpoint) == -1 )
		this.modelVersion.deployments.push(inferenceEndpoint);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more deploymentsIds as a Deployments
	// from a ModelVersion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDeployments( modelVersionId, deploymentsIds ): Observable<any> {

		// get the ModelVersion
		this.loadHelper( modelVersionId );


	// split on a comma with no spaces
	var idList 					= deploymentsIds.split(',');
	var deployments 	= this.modelVersion.deployments;

	if ( deployments != null && deploymentsIds != null ) {

		// iterate over array of deployments ids
		deployments.forEach(function (obj) {
			if ( deploymentsIds.indexOf(obj._id) > -1 ) {
				// remove the InferenceEndpoint
				this.modelVersion.deployments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more featureSetsIds as a FeatureSets
	// to a ModelVersion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFeatureSets( modelVersionId, featureSetsIds ): Observable<any> {

		// get the ModelVersion
		this.loadHelper( modelVersionId );

	// split on a comma with no spaces
	var idList = featureSetsIds.split(',')

	// iterate over array of featureSets ids
	idList.forEach(function (id) {
		// read the FeatureSet
		var featureSet = new FeatureSetService(this.http).getFeatureSet(id);
		// add the FeatureSet if not already assigned
		if ( this.modelVersion.featureSets.indexOf(featureSet) == -1 )
		this.modelVersion.featureSets.push(featureSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more featureSetsIds as a FeatureSets
	// from a ModelVersion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFeatureSets( modelVersionId, featureSetsIds ): Observable<any> {

		// get the ModelVersion
		this.loadHelper( modelVersionId );


	// split on a comma with no spaces
	var idList 					= featureSetsIds.split(',');
	var featureSets 	= this.modelVersion.featureSets;

	if ( featureSets != null && featureSetsIds != null ) {

		// iterate over array of featureSets ids
		featureSets.forEach(function (obj) {
			if ( featureSetsIds.indexOf(obj._id) > -1 ) {
				// remove the FeatureSet
				this.modelVersion.featureSets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a ModelVersion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( modelVersionId, datasetsIds ): Observable<any> {

		// get the ModelVersion
		this.loadHelper( modelVersionId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.modelVersion.datasets.indexOf(dataSet) == -1 )
		this.modelVersion.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a ModelVersion
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( modelVersionId, datasetsIds ): Observable<any> {

		// get the ModelVersion
		this.loadHelper( modelVersionId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.modelVersion.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.modelVersion.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a ModelVersion
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ModelVersion/update/' + this.modelVersion;

	return  this.http.post(uri_, this.modelVersion );
}

	//********************************************************************
	// loadHelper - internal helper to load a ModelVersion
	//********************************************************************	
	loadHelper( id ) {
		this.getModelVersion(id)
			.subscribe((res : ModelVersion) => {
				this.modelVersion = res;
			});
	}
}