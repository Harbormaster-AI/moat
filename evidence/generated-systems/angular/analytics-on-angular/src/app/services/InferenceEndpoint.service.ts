import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {InferenceEndpoint} from '../models/InferenceEndpoint';
import {ModelVersionService} from '../services/ModelVersion.service';
import {AnalyticsWorkspaceService} from '../services/AnalyticsWorkspace.service';
import {PredictionService} from '../services/Prediction.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class InferenceEndpointService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	inferenceEndpoint : InferenceEndpoint;

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
	// add a InferenceEndpoint
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addInferenceEndpoint(name, endpointUrl, trafficShare, ModelVersion, Workspace, Predictions, Mode) : Observable<any> {
		const uri_ = this.apiUrl + '/InferenceEndpoint/create';
		const obj = {
			      		name: name,
      		endpointUrl: endpointUrl,
      		trafficShare: trafficShare,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Predictions: Predictions != null && Predictions.length > 0 ? Predictions : null,
			Mode: Mode
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a InferenceEndpoint
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateInferenceEndpoint(name, endpointUrl, trafficShare, ModelVersion, Workspace, Predictions, Mode, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/InferenceEndpoint/update/' + id;
		const obj = {
				      		name: name,
      		endpointUrl: endpointUrl,
      		trafficShare: trafficShare,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
      		Workspace: Workspace != null && Workspace.length > 0 ? Workspace : null,
      		Predictions: Predictions != null && Predictions.length > 0 ? Predictions : null,
			Mode: Mode
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a InferenceEndpoint
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteInferenceEndpoint(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/InferenceEndpoint/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a InferenceEndpoint
	// returns the results untouched as an Observable InferenceEndpoint
	// InferenceEndpoint model
	// delegates via URI
	//********************************************************************
	getInferenceEndpoint(id) : Observable<InferenceEndpoint> {
		const uri_ = this.apiUrl + '/InferenceEndpoint/load/' + id;

		return this.http.get<InferenceEndpoint>(uri_);
	}
	
	//********************************************************************
	// gets all InferenceEndpoint
	// returns the results untouched as JSON representation of an
	// Observable array of InferenceEndpoint models
	// delegates via URI
	//********************************************************************
	getInferenceEndpoints() : Observable<InferenceEndpoint[]> {
		const uri_ = this.apiUrl + '/InferenceEndpoint/';

		return this
			.http.get<InferenceEndpoint[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ModelVersion on a InferenceEndpoint
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignModelVersion( inferenceEndpointId, _modelVersionId ): Observable<any> {

		// get the InferenceEndpoint from storage
		this.loadHelper( inferenceEndpointId );

	// get the ModelVersion from storage
	var tmp 	= new ModelVersionService(this.http).getModelVersion(_modelVersionId);

	// assign the ModelVersion
	this.inferenceEndpoint.modelVersion = tmp;

	// save the InferenceEndpoint
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ModelVersion on a InferenceEndpoint
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignModelVersion( inferenceEndpointId ): Observable<any> {

		// get the InferenceEndpoint from storage
		this.loadHelper( inferenceEndpointId );

	// assign ModelVersion to null
	this.inferenceEndpoint.modelVersion = null;

	// save the InferenceEndpoint
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Workspace on a InferenceEndpoint
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignWorkspace( inferenceEndpointId, _workspaceId ): Observable<any> {

		// get the InferenceEndpoint from storage
		this.loadHelper( inferenceEndpointId );

	// get the AnalyticsWorkspace from storage
	var tmp 	= new AnalyticsWorkspaceService(this.http).getAnalyticsWorkspace(_workspaceId);

	// assign the Workspace
	this.inferenceEndpoint.workspace = tmp;

	// save the InferenceEndpoint
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Workspace on a InferenceEndpoint
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignWorkspace( inferenceEndpointId ): Observable<any> {

		// get the InferenceEndpoint from storage
		this.loadHelper( inferenceEndpointId );

	// assign Workspace to null
	this.inferenceEndpoint.workspace = null;

	// save the InferenceEndpoint
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more predictionsIds as a Predictions
	// to a InferenceEndpoint
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPredictions( inferenceEndpointId, predictionsIds ): Observable<any> {

		// get the InferenceEndpoint
		this.loadHelper( inferenceEndpointId );

	// split on a comma with no spaces
	var idList = predictionsIds.split(',')

	// iterate over array of predictions ids
	idList.forEach(function (id) {
		// read the Prediction
		var prediction = new PredictionService(this.http).getPrediction(id);
		// add the Prediction if not already assigned
		if ( this.inferenceEndpoint.predictions.indexOf(prediction) == -1 )
		this.inferenceEndpoint.predictions.push(prediction);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more predictionsIds as a Predictions
	// from a InferenceEndpoint
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePredictions( inferenceEndpointId, predictionsIds ): Observable<any> {

		// get the InferenceEndpoint
		this.loadHelper( inferenceEndpointId );


	// split on a comma with no spaces
	var idList 					= predictionsIds.split(',');
	var predictions 	= this.inferenceEndpoint.predictions;

	if ( predictions != null && predictionsIds != null ) {

		// iterate over array of predictions ids
		predictions.forEach(function (obj) {
			if ( predictionsIds.indexOf(obj._id) > -1 ) {
				// remove the Prediction
				this.inferenceEndpoint.predictions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a InferenceEndpoint
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/InferenceEndpoint/update/' + this.inferenceEndpoint;

	return  this.http.post(uri_, this.inferenceEndpoint );
}

	//********************************************************************
	// loadHelper - internal helper to load a InferenceEndpoint
	//********************************************************************	
	loadHelper( id ) {
		this.getInferenceEndpoint(id)
			.subscribe((res : InferenceEndpoint) => {
				this.inferenceEndpoint = res;
			});
	}
}