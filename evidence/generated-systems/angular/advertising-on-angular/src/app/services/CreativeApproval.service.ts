import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CreativeApproval} from '../models/CreativeApproval';
import {CreativeAssetService} from '../services/CreativeAsset.service';
import {PublisherService} from '../services/Publisher.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CreativeApprovalService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	creativeApproval : CreativeApproval;

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
	// add a CreativeApproval
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCreativeApproval(reviewer, reviewedAt, CreativeAsset, Publisher, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/CreativeApproval/create';
		const obj = {
			      		reviewer: reviewer,
      		reviewedAt: reviewedAt,
      		CreativeAsset: CreativeAsset != null && CreativeAsset.length > 0 ? CreativeAsset : null,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CreativeApproval
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCreativeApproval(reviewer, reviewedAt, CreativeAsset, Publisher, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CreativeApproval/update/' + id;
		const obj = {
				      		reviewer: reviewer,
      		reviewedAt: reviewedAt,
      		CreativeAsset: CreativeAsset != null && CreativeAsset.length > 0 ? CreativeAsset : null,
      		Publisher: Publisher != null && Publisher.length > 0 ? Publisher : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CreativeApproval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCreativeApproval(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CreativeApproval/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CreativeApproval
	// returns the results untouched as an Observable CreativeApproval
	// CreativeApproval model
	// delegates via URI
	//********************************************************************
	getCreativeApproval(id) : Observable<CreativeApproval> {
		const uri_ = this.apiUrl + '/CreativeApproval/load/' + id;

		return this.http.get<CreativeApproval>(uri_);
	}
	
	//********************************************************************
	// gets all CreativeApproval
	// returns the results untouched as JSON representation of an
	// Observable array of CreativeApproval models
	// delegates via URI
	//********************************************************************
	getCreativeApprovals() : Observable<CreativeApproval[]> {
		const uri_ = this.apiUrl + '/CreativeApproval/';

		return this
			.http.get<CreativeApproval[]>(uri_);
	}
	
			//********************************************************************
	// assigns a CreativeAsset on a CreativeApproval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCreativeAsset( creativeApprovalId, _creativeAssetId ): Observable<any> {

		// get the CreativeApproval from storage
		this.loadHelper( creativeApprovalId );

	// get the CreativeAsset from storage
	var tmp 	= new CreativeAssetService(this.http).getCreativeAsset(_creativeAssetId);

	// assign the CreativeAsset
	this.creativeApproval.creativeAsset = tmp;

	// save the CreativeApproval
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a CreativeAsset on a CreativeApproval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCreativeAsset( creativeApprovalId ): Observable<any> {

		// get the CreativeApproval from storage
		this.loadHelper( creativeApprovalId );

	// assign CreativeAsset to null
	this.creativeApproval.creativeAsset = null;

	// save the CreativeApproval
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Publisher on a CreativeApproval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPublisher( creativeApprovalId, _publisherId ): Observable<any> {

		// get the CreativeApproval from storage
		this.loadHelper( creativeApprovalId );

	// get the Publisher from storage
	var tmp 	= new PublisherService(this.http).getPublisher(_publisherId);

	// assign the Publisher
	this.creativeApproval.publisher = tmp;

	// save the CreativeApproval
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Publisher on a CreativeApproval
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPublisher( creativeApprovalId ): Observable<any> {

		// get the CreativeApproval from storage
		this.loadHelper( creativeApprovalId );

	// assign Publisher to null
	this.creativeApproval.publisher = null;

	// save the CreativeApproval
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CreativeApproval
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CreativeApproval/update/' + this.creativeApproval;

	return  this.http.post(uri_, this.creativeApproval );
}

	//********************************************************************
	// loadHelper - internal helper to load a CreativeApproval
	//********************************************************************	
	loadHelper( id ) {
		this.getCreativeApproval(id)
			.subscribe((res : CreativeApproval) => {
				this.creativeApproval = res;
			});
	}
}