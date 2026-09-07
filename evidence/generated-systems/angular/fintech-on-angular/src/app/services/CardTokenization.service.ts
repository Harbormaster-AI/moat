import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CardTokenization} from '../models/CardTokenization';
import {PaymentCardService} from '../services/PaymentCard.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CardTokenizationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	cardTokenization : CardTokenization;

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
	// add a CardTokenization
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCardTokenization(tokenReference, createdAt, Card, WalletProvider, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/CardTokenization/create';
		const obj = {
			      		tokenReference: tokenReference,
      		createdAt: createdAt,
      		Card: Card != null && Card.length > 0 ? Card : null,
      		WalletProvider: WalletProvider,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CardTokenization
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCardTokenization(tokenReference, createdAt, Card, WalletProvider, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CardTokenization/update/' + id;
		const obj = {
				      		tokenReference: tokenReference,
      		createdAt: createdAt,
      		Card: Card != null && Card.length > 0 ? Card : null,
      		WalletProvider: WalletProvider,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CardTokenization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCardTokenization(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CardTokenization/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CardTokenization
	// returns the results untouched as an Observable CardTokenization
	// CardTokenization model
	// delegates via URI
	//********************************************************************
	getCardTokenization(id) : Observable<CardTokenization> {
		const uri_ = this.apiUrl + '/CardTokenization/load/' + id;

		return this.http.get<CardTokenization>(uri_);
	}
	
	//********************************************************************
	// gets all CardTokenization
	// returns the results untouched as JSON representation of an
	// Observable array of CardTokenization models
	// delegates via URI
	//********************************************************************
	getCardTokenizations() : Observable<CardTokenization[]> {
		const uri_ = this.apiUrl + '/CardTokenization/';

		return this
			.http.get<CardTokenization[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Card on a CardTokenization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCard( cardTokenizationId, _cardId ): Observable<any> {

		// get the CardTokenization from storage
		this.loadHelper( cardTokenizationId );

	// get the PaymentCard from storage
	var tmp 	= new PaymentCardService(this.http).getPaymentCard(_cardId);

	// assign the Card
	this.cardTokenization.card = tmp;

	// save the CardTokenization
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Card on a CardTokenization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCard( cardTokenizationId ): Observable<any> {

		// get the CardTokenization from storage
		this.loadHelper( cardTokenizationId );

	// assign Card to null
	this.cardTokenization.card = null;

	// save the CardTokenization
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CardTokenization
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CardTokenization/update/' + this.cardTokenization;

	return  this.http.post(uri_, this.cardTokenization );
}

	//********************************************************************
	// loadHelper - internal helper to load a CardTokenization
	//********************************************************************	
	loadHelper( id ) {
		this.getCardTokenization(id)
			.subscribe((res : CardTokenization) => {
				this.cardTokenization = res;
			});
	}
}