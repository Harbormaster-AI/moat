import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {FeatureSet} from '../models/FeatureSet';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {FeatureService} from '../services/Feature.service';
import {DataSetService} from '../services/DataSet.service';
import {Model_Service} from '../services/Model_.service';
import {ModelVersionService} from '../services/ModelVersion.service';
import {TagService} from '../services/Tag.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class FeatureSetService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	featureSet : FeatureSet;

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
	// add a FeatureSet
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addFeatureSet(name, refreshSchedule, Workspace, Features, Datasets, Models, ModelVersions, Tags, StoreType) : Observable<any> {
		const uri_ = this.apiUrl + '/FeatureSet/create';
		const obj = {
			      		name: name,
      		refreshSchedule: refreshSchedule,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Features: Features != null && Features.length > 0 ? Features : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		ModelVersions: ModelVersions != null && ModelVersions.length > 0 ? ModelVersions : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			StoreType: StoreType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a FeatureSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateFeatureSet(name, refreshSchedule, Workspace, Features, Datasets, Models, ModelVersions, Tags, StoreType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/FeatureSet/update/' + id;
		const obj = {
				      		name: name,
      		refreshSchedule: refreshSchedule,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Features: Features != null && Features.length > 0 ? Features : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Models: Models != null && Models.length > 0 ? Models : null,
      		ModelVersions: ModelVersions != null && ModelVersions.length > 0 ? ModelVersions : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			StoreType: StoreType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a FeatureSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteFeatureSet(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/FeatureSet/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a FeatureSet
	// returns the results untouched as an Observable FeatureSet
	// FeatureSet model
	// delegates via URI
	//********************************************************************
	getFeatureSet(id) : Observable<FeatureSet> {
		const uri_ = this.apiUrl + '/FeatureSet/load/' + id;

		return this.http.get<FeatureSet>(uri_);
	}
	
	//********************************************************************
	// gets all FeatureSet
	// returns the results untouched as JSON representation of an
	// Observable array of FeatureSet models
	// delegates via URI
	//********************************************************************
	getFeatureSets() : Observable<FeatureSet[]> {
		const uri_ = this.apiUrl + '/FeatureSet/';

		return this
			.http.get<FeatureSet[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a FeatureSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( featureSetId, _workspaceId ): Observable<any> {

		// get the FeatureSet from storage
		this.loadHelper( featureSetId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.featureSet.workspace = tmp;

	// save the FeatureSet
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a FeatureSet
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( featureSetId ): Observable<any> {

		// get the FeatureSet from storage
		this.loadHelper( featureSetId );

	// assign Workspace to null
	this.featureSet.workspace = null;

	// save the FeatureSet
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more featuresIds as a Features
	// to a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFeatures( featureSetId, featuresIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );

	// split on a comma with no spaces
	var idList = featuresIds.split(',')

	// iterate over array of features ids
	idList.forEach(function (id) {
		// read the Feature
		var feature = new FeatureService(this.http).getFeature(id);
		// add the Feature if not already assigned
		if ( this.featureSet.features.indexOf(feature) == -1 )
		this.featureSet.features.push(feature);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more featuresIds as a Features
	// from a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFeatures( featureSetId, featuresIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );


	// split on a comma with no spaces
	var idList 					= featuresIds.split(',');
	var features 	= this.featureSet.features;

	if ( features != null && featuresIds != null ) {

		// iterate over array of features ids
		features.forEach(function (obj) {
			if ( featuresIds.indexOf(obj._id) > -1 ) {
				// remove the Feature
				this.featureSet.features.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( featureSetId, datasetsIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.featureSet.datasets.indexOf(dataSet) == -1 )
		this.featureSet.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( featureSetId, datasetsIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.featureSet.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.featureSet.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more modelsIds as a Models
	// to a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModels( featureSetId, modelsIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );

	// split on a comma with no spaces
	var idList = modelsIds.split(',')

	// iterate over array of models ids
	idList.forEach(function (id) {
		// read the Model_
		var model_ = new Model_Service(this.http).getModel_(id);
		// add the Model_ if not already assigned
		if ( this.featureSet.models.indexOf(model_) == -1 )
		this.featureSet.models.push(model_);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelsIds as a Models
	// from a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModels( featureSetId, modelsIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );


	// split on a comma with no spaces
	var idList 					= modelsIds.split(',');
	var models 	= this.featureSet.models;

	if ( models != null && modelsIds != null ) {

		// iterate over array of models ids
		models.forEach(function (obj) {
			if ( modelsIds.indexOf(obj._id) > -1 ) {
				// remove the Model_
				this.featureSet.models.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more modelVersionsIds as a ModelVersions
	// to a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addModelVersions( featureSetId, modelVersionsIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );

	// split on a comma with no spaces
	var idList = modelVersionsIds.split(',')

	// iterate over array of modelVersions ids
	idList.forEach(function (id) {
		// read the ModelVersion
		var modelVersion = new ModelVersionService(this.http).getModelVersion(id);
		// add the ModelVersion if not already assigned
		if ( this.featureSet.modelVersions.indexOf(modelVersion) == -1 )
		this.featureSet.modelVersions.push(modelVersion);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more modelVersionsIds as a ModelVersions
	// from a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeModelVersions( featureSetId, modelVersionsIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );


	// split on a comma with no spaces
	var idList 					= modelVersionsIds.split(',');
	var modelVersions 	= this.featureSet.modelVersions;

	if ( modelVersions != null && modelVersionsIds != null ) {

		// iterate over array of modelVersions ids
		modelVersions.forEach(function (obj) {
			if ( modelVersionsIds.indexOf(obj._id) > -1 ) {
				// remove the ModelVersion
				this.featureSet.modelVersions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more tagsIds as a Tags
	// to a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTags( featureSetId, tagsIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );

	// split on a comma with no spaces
	var idList = tagsIds.split(',')

	// iterate over array of tags ids
	idList.forEach(function (id) {
		// read the Tag
		var tag = new TagService(this.http).getTag(id);
		// add the Tag if not already assigned
		if ( this.featureSet.tags.indexOf(tag) == -1 )
		this.featureSet.tags.push(tag);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tagsIds as a Tags
	// from a FeatureSet
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTags( featureSetId, tagsIds ): Observable<any> {

		// get the FeatureSet
		this.loadHelper( featureSetId );


	// split on a comma with no spaces
	var idList 					= tagsIds.split(',');
	var tags 	= this.featureSet.tags;

	if ( tags != null && tagsIds != null ) {

		// iterate over array of tags ids
		tags.forEach(function (obj) {
			if ( tagsIds.indexOf(obj._id) > -1 ) {
				// remove the Tag
				this.featureSet.tags.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a FeatureSet
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/FeatureSet/update/' + this.featureSet;

	return  this.http.post(uri_, this.featureSet );
}

	//********************************************************************
	// loadHelper - internal helper to load a FeatureSet
	//********************************************************************	
	loadHelper( id ) {
		this.getFeatureSet(id)
			.subscribe((res : FeatureSet) => {
				this.featureSet = res;
			});
	}
}