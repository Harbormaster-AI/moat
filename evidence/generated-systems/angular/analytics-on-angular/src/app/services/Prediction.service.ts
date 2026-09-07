import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Prediction} from '../models/Prediction';
import {InferenceEndpointService} from '../services/InferenceEndpoint.service';
import {ModelVersionService} from '../services/ModelVersion.service';
import {DataSetService} from '../services/DataSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PredictionService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	prediction : Prediction;

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
	// add a Prediction
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPrediction(referenceKey, predictedAt, score, Endpoint, ModelVersion, Dataset) : Observable<any> {
		const uri_ = this.apiUrl + '/Prediction/create';
		const obj = {
			      		referenceKey: referenceKey,
      		predictedAt: predictedAt,
      		score: score,
      		Endpoint: Endpoint != null && Endpoint.length > 0 ? Endpoint : null,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
			Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Prediction
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePrediction(referenceKey, predictedAt, score, Endpoint, ModelVersion, Dataset, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Prediction/update/' + id;
		const obj = {
				      		referenceKey: referenceKey,
      		predictedAt: predictedAt,
      		score: score,
      		Endpoint: Endpoint != null && Endpoint.length > 0 ? Endpoint : null,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
			Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Prediction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePrediction(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Prediction/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Prediction
	// returns the results untouched as an Observable Prediction
	// Prediction model
	// delegates via URI
	//********************************************************************
	getPrediction(id) : Observable<Prediction> {
		const uri_ = this.apiUrl + '/Prediction/load/' + id;

		return this.http.get<Prediction>(uri_);
	}
	
	//********************************************************************
	// gets all Prediction
	// returns the results untouched as JSON representation of an
	// Observable array of Prediction models
	// delegates via URI
	//********************************************************************
	getPredictions() : Observable<Prediction[]> {
		const uri_ = this.apiUrl + '/Prediction/';

		return this
			.http.get<Prediction[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Endpoint on a Prediction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEndpoint( predictionId, _endpointId ): Observable<any> {

		// get the Prediction from storage
		this.loadHelper( predictionId );

	// get the InferenceEndpoint from storage
	var tmp 	= new InferenceEndpointService(this.http).getInferenceEndpoint(_endpointId);

	// assign the Endpoint
	this.prediction.endpoint = tmp;

	// save the Prediction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Endpoint on a Prediction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEndpoint( predictionId ): Observable<any> {

		// get the Prediction from storage
		this.loadHelper( predictionId );

	// assign Endpoint to null
	this.prediction.endpoint = null;

	// save the Prediction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ModelVersion on a Prediction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignModelVersion( predictionId, _modelVersionId ): Observable<any> {

		// get the Prediction from storage
		this.loadHelper( predictionId );

	// get the ModelVersion from storage
	var tmp 	= new ModelVersionService(this.http).getModelVersion(_modelVersionId);

	// assign the ModelVersion
	this.prediction.modelVersion = tmp;

	// save the Prediction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ModelVersion on a Prediction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignModelVersion( predictionId ): Observable<any> {

		// get the Prediction from storage
		this.loadHelper( predictionId );

	// assign ModelVersion to null
	this.prediction.modelVersion = null;

	// save the Prediction
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Dataset on a Prediction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDataset( predictionId, _datasetId ): Observable<any> {

		// get the Prediction from storage
		this.loadHelper( predictionId );

	// get the DataSet from storage
	var tmp 	= new DataSetService(this.http).getDataSet(_datasetId);

	// assign the Dataset
	this.prediction.dataset = tmp;

	// save the Prediction
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dataset on a Prediction
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDataset( predictionId ): Observable<any> {

		// get the Prediction from storage
		this.loadHelper( predictionId );

	// assign Dataset to null
	this.prediction.dataset = null;

	// save the Prediction
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Prediction
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Prediction/update/' + this.prediction;

	return  this.http.post(uri_, this.prediction );
}

	//********************************************************************
	// loadHelper - internal helper to load a Prediction
	//********************************************************************	
	loadHelper( id ) {
		this.getPrediction(id)
			.subscribe((res : Prediction) => {
				this.prediction = res;
			});
	}
}