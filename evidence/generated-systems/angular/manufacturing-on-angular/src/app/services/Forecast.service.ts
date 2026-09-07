import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Forecast} from '../models/Forecast';
import {ForecastLineService} from '../services/ForecastLine.service';
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
	addForecast(forecastNumber, forecastHorizonStart, forecastHorizonEnd, Lines, Method) : Observable<any> {
		const uri_ = this.apiUrl + '/Forecast/create';
		const obj = {
			      		forecastNumber: forecastNumber,
      		forecastHorizonStart: forecastHorizonStart,
      		forecastHorizonEnd: forecastHorizonEnd,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
			Method: Method
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Forecast
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateForecast(forecastNumber, forecastHorizonStart, forecastHorizonEnd, Lines, Method, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Forecast/update/' + id;
		const obj = {
				      		forecastNumber: forecastNumber,
      		forecastHorizonStart: forecastHorizonStart,
      		forecastHorizonEnd: forecastHorizonEnd,
      		Lines: Lines != null && Lines.length > 0 ? Lines : null,
			Method: Method
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
	// adds one or more linesIds as a Lines
	// to a Forecast
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addLines( forecastId, linesIds ): Observable<any> {

		// get the Forecast
		this.loadHelper( forecastId );

	// split on a comma with no spaces
	var idList = linesIds.split(',')

	// iterate over array of lines ids
	idList.forEach(function (id) {
		// read the ForecastLine
		var forecastLine = new ForecastLineService(this.http).getForecastLine(id);
		// add the ForecastLine if not already assigned
		if ( this.forecast.lines.indexOf(forecastLine) == -1 )
		this.forecast.lines.push(forecastLine);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more linesIds as a Lines
	// from a Forecast
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeLines( forecastId, linesIds ): Observable<any> {

		// get the Forecast
		this.loadHelper( forecastId );


	// split on a comma with no spaces
	var idList 					= linesIds.split(',');
	var lines 	= this.forecast.lines;

	if ( lines != null && linesIds != null ) {

		// iterate over array of lines ids
		lines.forEach(function (obj) {
			if ( linesIds.indexOf(obj._id) > -1 ) {
				// remove the ForecastLine
				this.forecast.lines.pop(obj);
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