import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Model_} from '../models/Model_';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {ModelVersionService} from '../services/ModelVersion.service';
import {FeatureSetService} from '../services/FeatureSet.service';
import {ExperimentService} from '../services/Experiment.service';
import {TagService} from '../services/Tag.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class Model_Service extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	model_ : Model_;

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
	// add a Model_
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addModel_(name, taskDescription, Workspace, Versions, FeatureSets, Experiments, Tags, ModelType) : Observable<any> {
		const uri_ = this.apiUrl + '/Model_/create';
		const obj = {
			      		name: name,
      		taskDescription: taskDescription,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Versions: Versions != null && Versions.length > 0 ? Versions : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		Experiments: Experiments != null && Experiments.length > 0 ? Experiments : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			ModelType: ModelType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Model_
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateModel_(name, taskDescription, Workspace, Versions, FeatureSets, Experiments, Tags, ModelType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Model_/update/' + id;
		const obj = {
				      		name: name,
      		taskDescription: taskDescription,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Versions: Versions != null && Versions.length > 0 ? Versions : null,
      		FeatureSets: FeatureSets != null && FeatureSets.length > 0 ? FeatureSets : null,
      		Experiments: Experiments != null && Experiments.length > 0 ? Experiments : null,
      		Tags: Tags != null && Tags.length > 0 ? Tags : null,
			ModelType: ModelType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Model_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteModel_(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Model_/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Model_
	// returns the results untouched as an Observable Model_
	// Model_ model
	// delegates via URI
	//********************************************************************
	getModel_(id) : Observable<Model_> {
		const uri_ = this.apiUrl + '/Model_/load/' + id;

		return this.http.get<Model_>(uri_);
	}
	
	//********************************************************************
	// gets all Model_
	// returns the results untouched as JSON representation of an
	// Observable array of Model_ models
	// delegates via URI
	//********************************************************************
	getModel_s() : Observable<Model_[]> {
		const uri_ = this.apiUrl + '/Model_/';

		return this
			.http.get<Model_[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Workspace on a Model_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( model_Id, _workspaceId ): Observable<any> {

		// get the Model_ from storage
		this.loadHelper( model_Id );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.model_.workspace = tmp;

	// save the Model_
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a Model_
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( model_Id ): Observable<any> {

		// get the Model_ from storage
		this.loadHelper( model_Id );

	// assign Workspace to null
	this.model_.workspace = null;

	// save the Model_
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more versionsIds as a Versions
	// to a Model_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addVersions( model_Id, versionsIds ): Observable<any> {

		// get the Model_
		this.loadHelper( model_Id );

	// split on a comma with no spaces
	var idList = versionsIds.split(',')

	// iterate over array of versions ids
	idList.forEach(function (id) {
		// read the ModelVersion
		var modelVersion = new ModelVersionService(this.http).getModelVersion(id);
		// add the ModelVersion if not already assigned
		if ( this.model_.versions.indexOf(modelVersion) == -1 )
		this.model_.versions.push(modelVersion);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more versionsIds as a Versions
	// from a Model_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeVersions( model_Id, versionsIds ): Observable<any> {

		// get the Model_
		this.loadHelper( model_Id );


	// split on a comma with no spaces
	var idList 					= versionsIds.split(',');
	var versions 	= this.model_.versions;

	if ( versions != null && versionsIds != null ) {

		// iterate over array of versions ids
		versions.forEach(function (obj) {
			if ( versionsIds.indexOf(obj._id) > -1 ) {
				// remove the ModelVersion
				this.model_.versions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more featureSetsIds as a FeatureSets
	// to a Model_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addFeatureSets( model_Id, featureSetsIds ): Observable<any> {

		// get the Model_
		this.loadHelper( model_Id );

	// split on a comma with no spaces
	var idList = featureSetsIds.split(',')

	// iterate over array of featureSets ids
	idList.forEach(function (id) {
		// read the FeatureSet
		var featureSet = new FeatureSetService(this.http).getFeatureSet(id);
		// add the FeatureSet if not already assigned
		if ( this.model_.featureSets.indexOf(featureSet) == -1 )
		this.model_.featureSets.push(featureSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more featureSetsIds as a FeatureSets
	// from a Model_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeFeatureSets( model_Id, featureSetsIds ): Observable<any> {

		// get the Model_
		this.loadHelper( model_Id );


	// split on a comma with no spaces
	var idList 					= featureSetsIds.split(',');
	var featureSets 	= this.model_.featureSets;

	if ( featureSets != null && featureSetsIds != null ) {

		// iterate over array of featureSets ids
		featureSets.forEach(function (obj) {
			if ( featureSetsIds.indexOf(obj._id) > -1 ) {
				// remove the FeatureSet
				this.model_.featureSets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more experimentsIds as a Experiments
	// to a Model_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addExperiments( model_Id, experimentsIds ): Observable<any> {

		// get the Model_
		this.loadHelper( model_Id );

	// split on a comma with no spaces
	var idList = experimentsIds.split(',')

	// iterate over array of experiments ids
	idList.forEach(function (id) {
		// read the Experiment
		var experiment = new ExperimentService(this.http).getExperiment(id);
		// add the Experiment if not already assigned
		if ( this.model_.experiments.indexOf(experiment) == -1 )
		this.model_.experiments.push(experiment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more experimentsIds as a Experiments
	// from a Model_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeExperiments( model_Id, experimentsIds ): Observable<any> {

		// get the Model_
		this.loadHelper( model_Id );


	// split on a comma with no spaces
	var idList 					= experimentsIds.split(',');
	var experiments 	= this.model_.experiments;

	if ( experiments != null && experimentsIds != null ) {

		// iterate over array of experiments ids
		experiments.forEach(function (obj) {
			if ( experimentsIds.indexOf(obj._id) > -1 ) {
				// remove the Experiment
				this.model_.experiments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more tagsIds as a Tags
	// to a Model_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTags( model_Id, tagsIds ): Observable<any> {

		// get the Model_
		this.loadHelper( model_Id );

	// split on a comma with no spaces
	var idList = tagsIds.split(',')

	// iterate over array of tags ids
	idList.forEach(function (id) {
		// read the Tag
		var tag = new TagService(this.http).getTag(id);
		// add the Tag if not already assigned
		if ( this.model_.tags.indexOf(tag) == -1 )
		this.model_.tags.push(tag);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more tagsIds as a Tags
	// from a Model_
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTags( model_Id, tagsIds ): Observable<any> {

		// get the Model_
		this.loadHelper( model_Id );


	// split on a comma with no spaces
	var idList 					= tagsIds.split(',');
	var tags 	= this.model_.tags;

	if ( tags != null && tagsIds != null ) {

		// iterate over array of tags ids
		tags.forEach(function (obj) {
			if ( tagsIds.indexOf(obj._id) > -1 ) {
				// remove the Tag
				this.model_.tags.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Model_
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Model_/update/' + this.model_;

	return  this.http.post(uri_, this.model_ );
}

	//********************************************************************
	// loadHelper - internal helper to load a Model_
	//********************************************************************	
	loadHelper( id ) {
		this.getModel_(id)
			.subscribe((res : Model_) => {
				this.model_ = res;
			});
	}
}