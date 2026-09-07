import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Forecast} from '../models/Forecast';
import {ModelVersionService} from '../services/ModelVersion.service';
import {TimeSeriesService} from '../services/TimeSeries.service';
import {DataSetService} from '../services/DataSet.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ForecastService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	forecast : Forecast;

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
	// add a Forecast
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addForecast(name, horizon, ModelVersion, TimeSeries, Datasets, Granularity) : Observable<any> {
		const uri_ = this.apiUrl + '/Forecast/create';
		const obj = {
			      		name: name,
      		horizon: horizon,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
      		TimeSeries: TimeSeries != null && TimeSeries.length > 0 ? TimeSeries : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
			Granularity: Granularity
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Forecast
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateForecast(name, horizon, ModelVersion, TimeSeries, Datasets, Granularity, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Forecast/update/' + id;
		const obj = {
				      		name: name,
      		horizon: horizon,
      		ModelVersion: ModelVersion != null && ModelVersion.length > 0 ? ModelVersion : null,
      		TimeSeries: TimeSeries != null && TimeSeries.length > 0 ? TimeSeries : null,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
			Granularity: Granularity
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Forecast
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteForecast(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Forecast/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Forecast
	// returns the results untouched as an Observable Forecast
	// Forecast model
	// delegates via URI
	//********************************************************************
	getForecast(id) : Observable<Forecast> {
		const uri_ = this.apiUrl + '/Forecast/load/' + id;

		return this.http.get<Forecast>(uri_);
	}
	
	//********************************************************************
	// gets all Forecast
	// returns the results untouched as JSON representation of an
	// Observable array of Forecast models
	// delegates via URI
	//********************************************************************
	getForecasts() : Observable<Forecast[]> {
		const uri_ = this.apiUrl + '/Forecast/';

		return this
			.http.get<Forecast[]>(uri_);
	}
	
			//********************************************************************
	// assigns a ModelVersion on a Forecast
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignModelVersion( forecastId, _modelVersionId ): Observable<any> {

		// get the Forecast from storage
		this.loadHelper( forecastId );

	// get the ModelVersion from storage
	var tmp 	= new ModelVersionService(this.http).getModelVersion(_modelVersionId);

	// assign the ModelVersion
	this.forecast.modelVersion = tmp;

	// save the Forecast
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ModelVersion on a Forecast
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignModelVersion( forecastId ): Observable<any> {

		// get the Forecast from storage
		this.loadHelper( forecastId );

	// assign ModelVersion to null
	this.forecast.modelVersion = null;

	// save the Forecast
	return this.saveHelper();
}

		//********************************************************************
	// assigns a TimeSeries on a Forecast
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignTimeSeries( forecastId, _timeSeriesId ): Observable<any> {

		// get the Forecast from storage
		this.loadHelper( forecastId );

	// get the TimeSeries from storage
	var tmp 	= new TimeSeriesService(this.http).getTimeSeries(_timeSeriesId);

	// assign the TimeSeries
	this.forecast.timeSeries = tmp;

	// save the Forecast
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a TimeSeries on a Forecast
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignTimeSeries( forecastId ): Observable<any> {

		// get the Forecast from storage
		this.loadHelper( forecastId );

	// assign TimeSeries to null
	this.forecast.timeSeries = null;

	// save the Forecast
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a Forecast
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( forecastId, datasetsIds ): Observable<any> {

		// get the Forecast
		this.loadHelper( forecastId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.forecast.datasets.indexOf(dataSet) == -1 )
		this.forecast.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a Forecast
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( forecastId, datasetsIds ): Observable<any> {

		// get the Forecast
		this.loadHelper( forecastId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.forecast.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.forecast.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Forecast
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Forecast/update/' + this.forecast;

	return  this.http.post(uri_, this.forecast );
}

	//********************************************************************
	// loadHelper - internal helper to load a Forecast
	//********************************************************************	
	loadHelper( id ) {
		this.getForecast(id)
			.subscribe((res : Forecast) => {
				this.forecast = res;
			});
	}
}