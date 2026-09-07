import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TimeSeries} from '../models/TimeSeries';
import {DataSetService} from '../services/DataSet.service';
import {ForecastService} from '../services/Forecast.service';
import {AnomalyService} from '../services/Anomaly.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TimeSeriesService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	timeSeries : TimeSeries;

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
	// add a TimeSeries
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTimeSeries(name, timezone, Datasets, Forecasts, Anomalies, Granularity) : Observable<any> {
		const uri_ = this.apiUrl + '/TimeSeries/create';
		const obj = {
			      		name: name,
      		timezone: timezone,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Forecasts: Forecasts != null && Forecasts.length > 0 ? Forecasts : null,
      		Anomalies: Anomalies != null && Anomalies.length > 0 ? Anomalies : null,
			Granularity: Granularity
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TimeSeries
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTimeSeries(name, timezone, Datasets, Forecasts, Anomalies, Granularity, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TimeSeries/update/' + id;
		const obj = {
				      		name: name,
      		timezone: timezone,
      		Datasets: Datasets != null && Datasets.length > 0 ? Datasets : null,
      		Forecasts: Forecasts != null && Forecasts.length > 0 ? Forecasts : null,
      		Anomalies: Anomalies != null && Anomalies.length > 0 ? Anomalies : null,
			Granularity: Granularity
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TimeSeries
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTimeSeries(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TimeSeries/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TimeSeries
	// returns the results untouched as an Observable TimeSeries
	// TimeSeries model
	// delegates via URI
	//********************************************************************
	getTimeSeries(id) : Observable<TimeSeries> {
		const uri_ = this.apiUrl + '/TimeSeries/load/' + id;

		return this.http.get<TimeSeries>(uri_);
	}
	
	//********************************************************************
	// gets all TimeSeries
	// returns the results untouched as JSON representation of an
	// Observable array of TimeSeries models
	// delegates via URI
	//********************************************************************
	getTimeSeriess() : Observable<TimeSeries[]> {
		const uri_ = this.apiUrl + '/TimeSeries/';

		return this
			.http.get<TimeSeries[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more datasetsIds as a Datasets
	// to a TimeSeries
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDatasets( timeSeriesId, datasetsIds ): Observable<any> {

		// get the TimeSeries
		this.loadHelper( timeSeriesId );

	// split on a comma with no spaces
	var idList = datasetsIds.split(',')

	// iterate over array of datasets ids
	idList.forEach(function (id) {
		// read the DataSet
		var dataSet = new DataSetService(this.http).getDataSet(id);
		// add the DataSet if not already assigned
		if ( this.timeSeries.datasets.indexOf(dataSet) == -1 )
		this.timeSeries.datasets.push(dataSet);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more datasetsIds as a Datasets
	// from a TimeSeries
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDatasets( timeSeriesId, datasetsIds ): Observable<any> {

		// get the TimeSeries
		this.loadHelper( timeSeriesId );


	// split on a comma with no spaces
	var idList 					= datasetsIds.split(',');
	var datasets 	= this.timeSeries.datasets;

	if ( datasets != null && datasetsIds != null ) {

		// iterate over array of datasets ids
		datasets.forEach(function (obj) {
			if ( datasetsIds.indexOf(obj._id) > -1 ) {
				// remove the DataSet
				this.timeSeries.datasets.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more forecastsIds as a Forecasts
	// to a TimeSeries
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addForecasts( timeSeriesId, forecastsIds ): Observable<any> {

		// get the TimeSeries
		this.loadHelper( timeSeriesId );

	// split on a comma with no spaces
	var idList = forecastsIds.split(',')

	// iterate over array of forecasts ids
	idList.forEach(function (id) {
		// read the Forecast
		var forecast = new ForecastService(this.http).getForecast(id);
		// add the Forecast if not already assigned
		if ( this.timeSeries.forecasts.indexOf(forecast) == -1 )
		this.timeSeries.forecasts.push(forecast);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more forecastsIds as a Forecasts
	// from a TimeSeries
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeForecasts( timeSeriesId, forecastsIds ): Observable<any> {

		// get the TimeSeries
		this.loadHelper( timeSeriesId );


	// split on a comma with no spaces
	var idList 					= forecastsIds.split(',');
	var forecasts 	= this.timeSeries.forecasts;

	if ( forecasts != null && forecastsIds != null ) {

		// iterate over array of forecasts ids
		forecasts.forEach(function (obj) {
			if ( forecastsIds.indexOf(obj._id) > -1 ) {
				// remove the Forecast
				this.timeSeries.forecasts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more anomaliesIds as a Anomalies
	// to a TimeSeries
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAnomalies( timeSeriesId, anomaliesIds ): Observable<any> {

		// get the TimeSeries
		this.loadHelper( timeSeriesId );

	// split on a comma with no spaces
	var idList = anomaliesIds.split(',')

	// iterate over array of anomalies ids
	idList.forEach(function (id) {
		// read the Anomaly
		var anomaly = new AnomalyService(this.http).getAnomaly(id);
		// add the Anomaly if not already assigned
		if ( this.timeSeries.anomalies.indexOf(anomaly) == -1 )
		this.timeSeries.anomalies.push(anomaly);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more anomaliesIds as a Anomalies
	// from a TimeSeries
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAnomalies( timeSeriesId, anomaliesIds ): Observable<any> {

		// get the TimeSeries
		this.loadHelper( timeSeriesId );


	// split on a comma with no spaces
	var idList 					= anomaliesIds.split(',');
	var anomalies 	= this.timeSeries.anomalies;

	if ( anomalies != null && anomaliesIds != null ) {

		// iterate over array of anomalies ids
		anomalies.forEach(function (obj) {
			if ( anomaliesIds.indexOf(obj._id) > -1 ) {
				// remove the Anomaly
				this.timeSeries.anomalies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a TimeSeries
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TimeSeries/update/' + this.timeSeries;

	return  this.http.post(uri_, this.timeSeries );
}

	//********************************************************************
	// loadHelper - internal helper to load a TimeSeries
	//********************************************************************	
	loadHelper( id ) {
		this.getTimeSeries(id)
			.subscribe((res : TimeSeries) => {
				this.timeSeries = res;
			});
	}
}