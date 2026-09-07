import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {RunMetric} from '../models/RunMetric';
import {TrainingRunService} from '../services/TrainingRun.service';
import {MetricService} from '../services/Metric.service';
import {DataSetService} from '../services/DataSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RunMetricService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	runMetric : RunMetric;

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
	// add a RunMetric
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRunMetric(name, value, TrainingRun, Metric, Dataset) : Observable<any> {
		const uri_ = this.apiUrl + '/RunMetric/create';
		const obj = {
			      		name: name,
      		value: value,
      		TrainingRun: TrainingRun != null && TrainingRun.length > 0 ? TrainingRun : null,
      		Metric: Metric != null && Metric.length > 0 ? Metric : null,
			Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a RunMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRunMetric(name, value, TrainingRun, Metric, Dataset, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/RunMetric/update/' + id;
		const obj = {
				      		name: name,
      		value: value,
      		TrainingRun: TrainingRun != null && TrainingRun.length > 0 ? TrainingRun : null,
      		Metric: Metric != null && Metric.length > 0 ? Metric : null,
			Dataset: Dataset != null && Dataset.length > 0 ? Dataset : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a RunMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRunMetric(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/RunMetric/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a RunMetric
	// returns the results untouched as an Observable RunMetric
	// RunMetric model
	// delegates via URI
	//********************************************************************
	getRunMetric(id) : Observable<RunMetric> {
		const uri_ = this.apiUrl + '/RunMetric/load/' + id;

		return this.http.get<RunMetric>(uri_);
	}
	
	//********************************************************************
	// gets all RunMetric
	// returns the results untouched as JSON representation of an
	// Observable array of RunMetric models
	// delegates via URI
	//********************************************************************
	getRunMetrics() : Observable<RunMetric[]> {
		const uri_ = this.apiUrl + '/RunMetric/';

		return this
			.http.get<RunMetric[]>(uri_);
	}
	
			//********************************************************************
	// assigns a TrainingRun on a RunMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTrainingRun( runMetricId, _trainingRunId ): Observable<any> {

		// get the RunMetric from storage
		this.loadHelper( runMetricId );

	// get the TrainingRun from storage
	var tmp 	= new TrainingRunService(this.http).getTrainingRun(_trainingRunId);

	// assign the TrainingRun
	this.runMetric.trainingRun = tmp;

	// save the RunMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TrainingRun on a RunMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTrainingRun( runMetricId ): Observable<any> {

		// get the RunMetric from storage
		this.loadHelper( runMetricId );

	// assign TrainingRun to null
	this.runMetric.trainingRun = null;

	// save the RunMetric
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Metric on a RunMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignMetric( runMetricId, _metricId ): Observable<any> {

		// get the RunMetric from storage
		this.loadHelper( runMetricId );

	// get the Metric from storage
	var tmp 	= new MetricService(this.http).getMetric(_metricId);

	// assign the Metric
	this.runMetric.metric = tmp;

	// save the RunMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Metric on a RunMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignMetric( runMetricId ): Observable<any> {

		// get the RunMetric from storage
		this.loadHelper( runMetricId );

	// assign Metric to null
	this.runMetric.metric = null;

	// save the RunMetric
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Dataset on a RunMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignDataset( runMetricId, _datasetId ): Observable<any> {

		// get the RunMetric from storage
		this.loadHelper( runMetricId );

	// get the DataSet from storage
	var tmp 	= new DataSetService(this.http).getDataSet(_datasetId);

	// assign the Dataset
	this.runMetric.dataset = tmp;

	// save the RunMetric
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Dataset on a RunMetric
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignDataset( runMetricId ): Observable<any> {

		// get the RunMetric from storage
		this.loadHelper( runMetricId );

	// assign Dataset to null
	this.runMetric.dataset = null;

	// save the RunMetric
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a RunMetric
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/RunMetric/update/' + this.runMetric;

	return  this.http.post(uri_, this.runMetric );
}

	//********************************************************************
	// loadHelper - internal helper to load a RunMetric
	//********************************************************************	
	loadHelper( id ) {
		this.getRunMetric(id)
			.subscribe((res : RunMetric) => {
				this.runMetric = res;
			});
	}
}