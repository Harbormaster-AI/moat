import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {RateCard} from '../models/RateCard';
import {PublisherService} from '../services/Publisher.service';
import {RateService} from '../services/Rate.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RateCardService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	rateCard : RateCard;

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
	// add a RateCard
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRateCard(name, effectiveDate, currency, Publisher, Rates) : Observable<any> {
		const uri_ = this.apiUrl + '/RateCard/create';
		const obj = {
			      		name: name,
      		effectiveDate: effectiveDate,
      		currency: currency,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
			Rates: Rates != null && Rates.length > 0 ? Rates : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a RateCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRateCard(name, effectiveDate, currency, Publisher, Rates, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/RateCard/update/' + id;
		const obj = {
				      		name: name,
      		effectiveDate: effectiveDate,
      		currency: currency,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
			Rates: Rates != null && Rates.length > 0 ? Rates : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a RateCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRateCard(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/RateCard/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a RateCard
	// returns the results untouched as an Observable RateCard
	// RateCard model
	// delegates via URI
	//********************************************************************
	getRateCard(id) : Observable<RateCard> {
		const uri_ = this.apiUrl + '/RateCard/load/' + id;

		return this.http.get<RateCard>(uri_);
	}
	
	//********************************************************************
	// gets all RateCard
	// returns the results untouched as JSON representation of an
	// Observable array of RateCard models
	// delegates via URI
	//********************************************************************
	getRateCards() : Observable<RateCard[]> {
		const uri_ = this.apiUrl + '/RateCard/';

		return this
			.http.get<RateCard[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Publisher on a RateCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPublisher( rateCardId, _publisherId ): Observable<any> {

		// get the RateCard from storage
		this.loadHelper( rateCardId );

	// get the Publisher from storage
	var tmp 	= new PublisherService(this.http).getPublisher(_publisherId);

	// assign the Publisher
	this.rateCard.publisher = tmp;

	// save the RateCard
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Publisher on a RateCard
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPublisher( rateCardId ): Observable<any> {

		// get the RateCard from storage
		this.loadHelper( rateCardId );

	// assign Publisher to null
	this.rateCard.publisher = null;

	// save the RateCard
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more ratesIds as a Rates
	// to a RateCard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addRates( rateCardId, ratesIds ): Observable<any> {

		// get the RateCard
		this.loadHelper( rateCardId );

	// split on a comma with no spaces
	var idList = ratesIds.split(',')

	// iterate over array of rates ids
	idList.forEach(function (id) {
		// read the Rate
		var rate = new RateService(this.http).getRate(id);
		// add the Rate if not already assigned
		if ( this.rateCard.rates.indexOf(rate) == -1 )
		this.rateCard.rates.push(rate);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more ratesIds as a Rates
	// from a RateCard
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeRates( rateCardId, ratesIds ): Observable<any> {

		// get the RateCard
		this.loadHelper( rateCardId );


	// split on a comma with no spaces
	var idList 					= ratesIds.split(',');
	var rates 	= this.rateCard.rates;

	if ( rates != null && ratesIds != null ) {

		// iterate over array of rates ids
		rates.forEach(function (obj) {
			if ( ratesIds.indexOf(obj._id) > -1 ) {
				// remove the Rate
				this.rateCard.rates.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a RateCard
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/RateCard/update/' + this.rateCard;

	return  this.http.post(uri_, this.rateCard );
}

	//********************************************************************
	// loadHelper - internal helper to load a RateCard
	//********************************************************************	
	loadHelper( id ) {
		this.getRateCard(id)
			.subscribe((res : RateCard) => {
				this.rateCard = res;
			});
	}
}