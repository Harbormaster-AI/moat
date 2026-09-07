import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Feature} from '../models/Feature';
import {FeatureSetService} from '../services/FeatureSet.service';
import {DataSetService} from '../services/DataSet.service';
import {Model_Service} from '../services/Model_.service';
import {TrainingRunService} from '../services/TrainingRun.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FeatureService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	feature : Feature;

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
	// add a Feature
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFeature(name, description, FeatureSet, SourceDatasets, Models, TrainingRuns, DataType) : Observable<any> {
		const uri_ = this.apiUrl + '/Feature/create';
		const obj = {
			      		name: name,
      		description: description,
      		FeatureSet: FeatureSet != null && FeatureSet.length > 0 ? FeatureSet : null,
      		SourceDatasets: SourceDatasets != null && SourceDatasets.length > 0 ? SourceDatasets : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		TrainingRuns: TrainingRuns != null && TrainingRuns.length > 0 ? TrainingRuns : null,
			DataType: DataType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Feature
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFeature(name, description, FeatureSet, SourceDatasets, Models, TrainingRuns, DataType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Feature/update/' + id;
		const obj = {
				      		name: name,
      		description: description,
      		FeatureSet: FeatureSet != null && FeatureSet.length > 0 ? FeatureSet : null,
      		SourceDatasets: SourceDatasets != null && SourceDatasets.length > 0 ? SourceDatasets : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		TrainingRuns: TrainingRuns != null && TrainingRuns.length > 0 ? TrainingRuns : null,
			DataType: DataType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Feature
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFeature(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Feature/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Feature
	// returns the results untouched as an Observable Feature
	// Feature model
	// delegates via URI
	//********************************************************************
	getFeature(id) : Observable<Feature> {
		const uri_ = this.apiUrl + '/Feature/load/' + id;

		return this.http.get<Feature>(uri_);
	}
	
	//********************************************************************
	// gets all Feature
	// returns the results untouched as JSON representation of an
	// Observable array of Feature models
	// delegates via URI
	//********************************************************************
	getFeatures() : Observable<Feature[]> {
		const uri_ = this.apiUrl + '/Feature/';

		return this
			.http.get<Feature[]>(uri_);
	}
	
			//********************************************************************
	// assigns a FeatureSet on a Feature
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFeatureSet( featureId, _featureSetId ): Observable<any> {

		// get the Feature from storage
		this.loadHelper( featureId );

	// get the FeatureSet from storage
	var tmp 	= new FeatureSetService(this.http).getFeatureSet(_featureSetId);

	// assign the FeatureSet
	this.feature.featureSet = tmp;

	// save the Feature
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a FeatureSet on a Feature
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFeatureSet( featureId ): Observable<any> {

		// get the Feature from storage
		this.loadHelper( featureId );

	// assign FeatureSet to null
	this.feature.featureSet = null;

	// save the Feature
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more sourceDatasetsIds as a SourceDatasets
	// to a Feature
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addSourceDatasets( featureId, sourceDatasetsIds ): Observable<any> {

		// get the Feature
		this.loadHelper( featureId );

	// split on a comma with no spaces
	var idList = sourceDatasetsIds.split(',')

	// iterate over array of sourceDatasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.feature.sourceDatasets.indexOf(dataSet) == -1 )
		this.feature.sourceDatasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more sourceDatasetsIds as a SourceDatasets
	// from a Feature
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeSourceDatasets( featureId, sourceDatasetsIds ): Observable<any> {

		// get the Feature
		this.loadHelper( featureId );


	// split on a comma with no spaces
	var idList 					= sourceDatasetsIds.split(',');
	var sourceDatasets 	= this.feature.sourceDatasets;

	if ( sourceDatasets != null && sourceDatasetsIds != null ) {

		// iterate over array of sourceDatasets ids
		sourceDatasets.forEach(function (obj) {
			if ( sourceDatasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.feature.sourceDatasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more modelsIds as a Models
	// to a Feature
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModels( featureId, modelsIds ): Observable<any> {

		// get the Feature
		this.loadHelper( featureId );

	// split on a comma with no spaces
	var idList = modelsIds.split(',')

	// iterate over array of models ids
	idList.forEach(function (id) {
		// read the Model_
		var model_ = new Model_Service(this.http).getModel_(id);
		// add the Model_ if not already assigned
		if ( this.feature.models.indexOf(model_) == -1 )
		this.feature.models.push(model_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelsIds as a Models
	// from a Feature
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModels( featureId, modelsIds ): Observable<any> {

		// get the Feature
		this.loadHelper( featureId );


	// split on a comma with no spaces
	var idList 					= modelsIds.split(',');
	var models 	= this.feature.models;

	if ( models != null && modelsIds != null ) {

		// iterate over array of models ids
		models.forEach(function (obj) {
			if ( modelsIds.indexOf(obj._id) > -1 ) {
				// remove the Model_
				this.feature.models.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more trainingRunsIds as a TrainingRuns
	// to a Feature
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTrainingRuns( featureId, trainingRunsIds ): Observable<any> {

		// get the Feature
		this.loadHelper( featureId );

	// split on a comma with no spaces
	var idList = trainingRunsIds.split(',')

	// iterate over array of trainingRuns ids
	idList.forEach(function (id) {
		// read the TrainingRun
		var trainingRun = new TrainingRunService(this.http).getTrainingRun(id);
		// add the TrainingRun if not already assigned
		if ( this.feature.trainingRuns.indexOf(trainingRun) == -1 )
		this.feature.trainingRuns.push(trainingRun);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more trainingRunsIds as a TrainingRuns
	// from a Feature
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTrainingRuns( featureId, trainingRunsIds ): Observable<any> {

		// get the Feature
		this.loadHelper( featureId );


	// split on a comma with no spaces
	var idList 					= trainingRunsIds.split(',');
	var trainingRuns 	= this.feature.trainingRuns;

	if ( trainingRuns != null && trainingRunsIds != null ) {

		// iterate over array of trainingRuns ids
		trainingRuns.forEach(function (obj) {
			if ( trainingRunsIds.indexOf(obj._id) > -1 ) {
				// remove the TrainingRun
				this.feature.trainingRuns.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Feature
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Feature/update/' + this.feature;

	return  this.http.post(uri_, this.feature );
}

	//********************************************************************
	// loadHelper - internal helper to load a Feature
	//********************************************************************	
	loadHelper( id ) {
		this.getFeature(id)
			.subscribe((res : Feature) => {
				this.feature = res;
			});
	}
}