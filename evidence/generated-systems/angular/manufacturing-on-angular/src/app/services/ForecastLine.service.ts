import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {ForecastLine} from '../models/ForecastLine';
import {ForecastService} from '../services/Forecast.service';
import {ItemService} from '../services/Item.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ForecastLineService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	forecastLine : ForecastLine;

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
	// add a ForecastLine
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addForecastLine(period, quantity, confidence, Forecast, Item) : Observable<any> {
		const uri_ = this.apiUrl + '/ForecastLine/create';
		const obj = {
			      		period: period,
      		quantity: quantity,
      		confidence: confidence,
      		Forecast: Forecast != null && Forecast.length > 0 ? Forecast : null,
			Item: Item != null && Item.length > 0 ? Item : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a ForecastLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateForecastLine(period, quantity, confidence, Forecast, Item, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/ForecastLine/update/' + id;
		const obj = {
				      		period: period,
      		quantity: quantity,
      		confidence: confidence,
      		Forecast: Forecast != null && Forecast.length > 0 ? Forecast : null,
			Item: Item != null && Item.length > 0 ? Item : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a ForecastLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteForecastLine(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/ForecastLine/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a ForecastLine
	// returns the results untouched as an Observable ForecastLine
	// ForecastLine model
	// delegates via URI
	//********************************************************************
	getForecastLine(id) : Observable<ForecastLine> {
		const uri_ = this.apiUrl + '/ForecastLine/load/' + id;

		return this.http.get<ForecastLine>(uri_);
	}
	
	//********************************************************************
	// gets all ForecastLine
	// returns the results untouched as JSON representation of an
	// Observable array of ForecastLine models
	// delegates via URI
	//********************************************************************
	getForecastLines() : Observable<ForecastLine[]> {
		const uri_ = this.apiUrl + '/ForecastLine/';

		return this
			.http.get<ForecastLine[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Forecast on a ForecastLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignForecast( forecastLineId, _forecastId ): Observable<any> {

		// get the ForecastLine from storage
		this.loadHelper( forecastLineId );

	// get the Forecast from storage
	var tmp 	= new ForecastService(this.http).getForecast(_forecastId);

	// assign the Forecast
	this.forecastLine.forecast = tmp;

	// save the ForecastLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Forecast on a ForecastLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignForecast( forecastLineId ): Observable<any> {

		// get the ForecastLine from storage
		this.loadHelper( forecastLineId );

	// assign Forecast to null
	this.forecastLine.forecast = null;

	// save the ForecastLine
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Item on a ForecastLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignItem( forecastLineId, _itemId ): Observable<any> {

		// get the ForecastLine from storage
		this.loadHelper( forecastLineId );

	// get the Item from storage
	var tmp 	= new ItemService(this.http).getItem(_itemId);

	// assign the Item
	this.forecastLine.item = tmp;

	// save the ForecastLine
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Item on a ForecastLine
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignItem( forecastLineId ): Observable<any> {

		// get the ForecastLine from storage
		this.loadHelper( forecastLineId );

	// assign Item to null
	this.forecastLine.item = null;

	// save the ForecastLine
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a ForecastLine
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/ForecastLine/update/' + this.forecastLine;

	return  this.http.post(uri_, this.forecastLine );
}

	//********************************************************************
	// loadHelper - internal helper to load a ForecastLine
	//********************************************************************	
	loadHelper( id ) {
		this.getForecastLine(id)
			.subscribe((res : ForecastLine) => {
				this.forecastLine = res;
			});
	}
}