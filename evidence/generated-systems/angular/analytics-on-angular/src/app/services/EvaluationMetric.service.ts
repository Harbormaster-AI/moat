import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {EvaluationMetric} from '../models/EvaluationMetric';
import {ModelVersionService} from '../services/ModelVersion.service';
import {MetricService} from '../services/Metric.service';
import {DataSetService} from '../services/DataSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class EvaluationMetricService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	evaluationMetric : EvaluationMetric;

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
	// add a EvaluationMetric
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addEvaluationMetric(name, value, ModelVersion, Metric, Dataset) : Observable<any> {
		const uri_ = this.apiUrl + '/EvaluationMetric/create';
		const obj = {
			      		name: name,
      		value: value,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
      		Metric: Metric != null && Metric.length > 0 ? Metric : null,
			Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a EvaluationMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateEvaluationMetric(name, value, ModelVersion, Metric, Dataset, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/EvaluationMetric/update/' + id;
		const obj = {
				      		name: name,
      		value: value,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
      		Metric: Metric != null && Metric.length > 0 ? Metric : null,
			Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a EvaluationMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteEvaluationMetric(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/EvaluationMetric/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a EvaluationMetric
	// returns the results untouched as an Observable EvaluationMetric
	// EvaluationMetric model
	// delegates via URI
	//********************************************************************
	getEvaluationMetric(id) : Observable<EvaluationMetric> {
		const uri_ = this.apiUrl + '/EvaluationMetric/load/' + id;

		return this.http.get<EvaluationMetric>(uri_);
	}
	
	//********************************************************************
	// gets all EvaluationMetric
	// returns the results untouched as JSON representation of an
	// Observable array of EvaluationMetric models
	// delegates via URI
	//********************************************************************
	getEvaluationMetrics() : Observable<EvaluationMetric[]> {
		const uri_ = this.apiUrl + '/EvaluationMetric/';

		return this
			.http.get<EvaluationMetric[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ModelVersion on a EvaluationMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignModelVersion( evaluationMetricId, _modelVersionId ): Observable<any> {

		// get the EvaluationMetric from storage
		this.loadHelper( evaluationMetricId );

	// get the ModelVersion from storage
	var tmp 	= new ModelVersionService(this.http).getModelVersion(_modelVersionId);

	// assign the ModelVersion
	this.evaluationMetric.modelVersion = tmp;

	// save the EvaluationMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ModelVersion on a EvaluationMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignModelVersion( evaluationMetricId ): Observable<any> {

		// get the EvaluationMetric from storage
		this.loadHelper( evaluationMetricId );

	// assign ModelVersion to null
	this.evaluationMetric.modelVersion = null;

	// save the EvaluationMetric
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Metric on a EvaluationMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMetric( evaluationMetricId, _metricId ): Observable<any> {

		// get the EvaluationMetric from storage
		this.loadHelper( evaluationMetricId );

	// get the Metric from storage
	var tmp 	= new MetricService(this.http).getMetric(_metricId);

	// assign the Metric
	this.evaluationMetric.metric = tmp;

	// save the EvaluationMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Metric on a EvaluationMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMetric( evaluationMetricId ): Observable<any> {

		// get the EvaluationMetric from storage
		this.loadHelper( evaluationMetricId );

	// assign Metric to null
	this.evaluationMetric.metric = null;

	// save the EvaluationMetric
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Dataset on a EvaluationMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDataset( evaluationMetricId, _datasetId ): Observable<any> {

		// get the EvaluationMetric from storage
		this.loadHelper( evaluationMetricId );

	// get the DataSet from storage
	var tmp 	= new DataSetService(this.http).getDataSet(_datasetId);

	// assign the Dataset
	this.evaluationMetric.dataset = tmp;

	// save the EvaluationMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dataset on a EvaluationMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDataset( evaluationMetricId ): Observable<any> {

		// get the EvaluationMetric from storage
		this.loadHelper( evaluationMetricId );

	// assign Dataset to null
	this.evaluationMetric.dataset = null;

	// save the EvaluationMetric
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a EvaluationMetric
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/EvaluationMetric/update/' + this.evaluationMetric;

	return  this.http.post(uri_, this.evaluationMetric );
}

	//********************************************************************
	// loadHelper - internal helper to load a EvaluationMetric
	//********************************************************************	
	loadHelper( id ) {
		this.getEvaluationMetric(id)
			.subscribe((res : EvaluationMetric) => {
				this.evaluationMetric = res;
			});
	}
}