import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CreativeVariation} from '../models/CreativeVariation';
import {CreativeAssetService} from '../services/CreativeAsset.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CreativeVariationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	creativeVariation : CreativeVariation;

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
	// add a CreativeVariation
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCreativeVariation(name, language, headline, bodyText, callToAction, CreativeAsset) : Observable<any> {
		const uri_ = this.apiUrl + '/CreativeVariation/create';
		const obj = {
			      		name: name,
      		language: language,
      		headline: headline,
      		bodyText: bodyText,
      		callToAction: callToAction,
			CreativeAsset: CreativeAsset != null && CreativeAsset.length > 0 ? CreativeAsset : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CreativeVariation
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCreativeVariation(name, language, headline, bodyText, callToAction, CreativeAsset, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CreativeVariation/update/' + id;
		const obj = {
				      		name: name,
      		language: language,
      		headline: headline,
      		bodyText: bodyText,
      		callToAction: callToAction,
			CreativeAsset: CreativeAsset != null && CreativeAsset.length > 0 ? CreativeAsset : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CreativeVariation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCreativeVariation(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CreativeVariation/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CreativeVariation
	// returns the results untouched as an Observable CreativeVariation
	// CreativeVariation model
	// delegates via URI
	//********************************************************************
	getCreativeVariation(id) : Observable<CreativeVariation> {
		const uri_ = this.apiUrl + '/CreativeVariation/load/' + id;

		return this.http.get<CreativeVariation>(uri_);
	}
	
	//********************************************************************
	// gets all CreativeVariation
	// returns the results untouched as JSON representation of an
	// Observable array of CreativeVariation models
	// delegates via URI
	//********************************************************************
	getCreativeVariations() : Observable<CreativeVariation[]> {
		const uri_ = this.apiUrl + '/CreativeVariation/';

		return this
			.http.get<CreativeVariation[]>(uri_);
	}
	
			//********************************************************************
	// assigns a CreativeAsset on a CreativeVariation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCreativeAsset( creativeVariationId, _creativeAssetId ): Observable<any> {

		// get the CreativeVariation from storage
		this.loadHelper( creativeVariationId );

	// get the CreativeAsset from storage
	var tmp 	= new CreativeAssetService(this.http).getCreativeAsset(_creativeAssetId);

	// assign the CreativeAsset
	this.creativeVariation.creativeAsset = tmp;

	// save the CreativeVariation
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CreativeAsset on a CreativeVariation
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCreativeAsset( creativeVariationId ): Observable<any> {

		// get the CreativeVariation from storage
		this.loadHelper( creativeVariationId );

	// assign CreativeAsset to null
	this.creativeVariation.creativeAsset = null;

	// save the CreativeVariation
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CreativeVariation
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CreativeVariation/update/' + this.creativeVariation;

	return  this.http.post(uri_, this.creativeVariation );
}

	//********************************************************************
	// loadHelper - internal helper to load a CreativeVariation
	//********************************************************************	
	loadHelper( id ) {
		this.getCreativeVariation(id)
			.subscribe((res : CreativeVariation) => {
				this.creativeVariation = res;
			});
	}
}