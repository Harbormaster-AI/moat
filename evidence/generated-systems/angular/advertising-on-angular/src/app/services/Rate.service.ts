import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Rate} from '../models/Rate';
import {RateCardService} from '../services/RateCard.service';
import {AdSlotService} from '../services/AdSlot.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class RateService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	rate : Rate;

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
	// add a Rate
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addRate(unitPrice, RateCard, AdSlot, AdFormat, PricingModel) : Observable<any> {
		const uri_ = this.apiUrl + '/Rate/create';
		const obj = {
			      		unitPrice: unitPrice,
      		RateCard: RateCard != null && RateCard.length > 0 ? RateCard : null,
      		AdSlot: AdSlot != null && AdSlot.length > 0 ? AdSlot : null,
      		AdFormat: AdFormat,
			PricingModel: PricingModel
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Rate
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateRate(unitPrice, RateCard, AdSlot, AdFormat, PricingModel, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Rate/update/' + id;
		const obj = {
				      		unitPrice: unitPrice,
      		RateCard: RateCard != null && RateCard.length > 0 ? RateCard : null,
      		AdSlot: AdSlot != null && AdSlot.length > 0 ? AdSlot : null,
      		AdFormat: AdFormat,
			PricingModel: PricingModel
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Rate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteRate(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Rate/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Rate
	// returns the results untouched as an Observable Rate
	// Rate model
	// delegates via URI
	//********************************************************************
	getRate(id) : Observable<Rate> {
		const uri_ = this.apiUrl + '/Rate/load/' + id;

		return this.http.get<Rate>(uri_);
	}
	
	//********************************************************************
	// gets all Rate
	// returns the results untouched as JSON representation of an
	// Observable array of Rate models
	// delegates via URI
	//********************************************************************
	getRates() : Observable<Rate[]> {
		const uri_ = this.apiUrl + '/Rate/';

		return this
			.http.get<Rate[]>(uri_);
	}
	
			//********************************************************************
	// assigns a RateCard on a Rate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRateCard( rateId, _rateCardId ): Observable<any> {

		// get the Rate from storage
		this.loadHelper( rateId );

	// get the RateCard from storage
	var tmp 	= new RateCardService(this.http).getRateCard(_rateCardId);

	// assign the RateCard
	this.rate.rateCard = tmp;

	// save the Rate
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a RateCard on a Rate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRateCard( rateId ): Observable<any> {

		// get the Rate from storage
		this.loadHelper( rateId );

	// assign RateCard to null
	this.rate.rateCard = null;

	// save the Rate
	return this.saveHelper();
}

		//********************************************************************
	// assigns a AdSlot on a Rate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdSlot( rateId, _adSlotId ): Observable<any> {

		// get the Rate from storage
		this.loadHelper( rateId );

	// get the AdSlot from storage
	var tmp 	= new AdSlotService(this.http).getAdSlot(_adSlotId);

	// assign the AdSlot
	this.rate.adSlot = tmp;

	// save the Rate
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AdSlot on a Rate
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdSlot( rateId ): Observable<any> {

		// get the Rate from storage
		this.loadHelper( rateId );

	// assign AdSlot to null
	this.rate.adSlot = null;

	// save the Rate
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Rate
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Rate/update/' + this.rate;

	return  this.http.post(uri_, this.rate );
}

	//********************************************************************
	// loadHelper - internal helper to load a Rate
	//********************************************************************	
	loadHelper( id ) {
		this.getRate(id)
			.subscribe((res : Rate) => {
				this.rate = res;
			});
	}
}