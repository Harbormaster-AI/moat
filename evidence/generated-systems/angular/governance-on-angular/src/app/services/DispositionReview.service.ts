import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {DispositionReview} from '../models/DispositionReview';
import {Record_Service} from '../services/Record_.service';
import {RetentionScheduleService} from '../services/RetentionSchedule.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DispositionReviewService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	dispositionReview : DispositionReview;

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
	// add a DispositionReview
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDispositionReview(reviewDate, reviewer, notes, Record, RetentionSchedule, Outcome) : Observable<any> {
		const uri_ = this.apiUrl + '/DispositionReview/create';
		const obj = {
			      		reviewDate: reviewDate,
      		reviewer: reviewer,
      		notes: notes,
      		Record: Record != null && Record.length > 0 ? Record : null,
      		RetentionSchedule: RetentionSchedule != null && RetentionSchedule.length > 0 ? RetentionSchedule : null,
			Outcome: Outcome
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a DispositionReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDispositionReview(reviewDate, reviewer, notes, Record, RetentionSchedule, Outcome, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/DispositionReview/update/' + id;
		const obj = {
				      		reviewDate: reviewDate,
      		reviewer: reviewer,
      		notes: notes,
      		Record: Record != null && Record.length > 0 ? Record : null,
      		RetentionSchedule: RetentionSchedule != null && RetentionSchedule.length > 0 ? RetentionSchedule : null,
			Outcome: Outcome
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a DispositionReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDispositionReview(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/DispositionReview/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a DispositionReview
	// returns the results untouched as an Observable DispositionReview
	// DispositionReview model
	// delegates via URI
	//********************************************************************
	getDispositionReview(id) : Observable<DispositionReview> {
		const uri_ = this.apiUrl + '/DispositionReview/load/' + id;

		return this.http.get<DispositionReview>(uri_);
	}
	
	//********************************************************************
	// gets all DispositionReview
	// returns the results untouched as JSON representation of an
	// Observable array of DispositionReview models
	// delegates via URI
	//********************************************************************
	getDispositionReviews() : Observable<DispositionReview[]> {
		const uri_ = this.apiUrl + '/DispositionReview/';

		return this
			.http.get<DispositionReview[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Record on a DispositionReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRecord( dispositionReviewId, _recordId ): Observable<any> {

		// get the DispositionReview from storage
		this.loadHelper( dispositionReviewId );

	// get the Record_ from storage
	var tmp 	= new Record_Service(this.http).getRecord_(_recordId);

	// assign the Record
	this.dispositionReview.record = tmp;

	// save the DispositionReview
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Record on a DispositionReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRecord( dispositionReviewId ): Observable<any> {

		// get the DispositionReview from storage
		this.loadHelper( dispositionReviewId );

	// assign Record to null
	this.dispositionReview.record = null;

	// save the DispositionReview
	return this.saveHelper();
}

		//********************************************************************
	// assigns a RetentionSchedule on a DispositionReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRetentionSchedule( dispositionReviewId, _retentionScheduleId ): Observable<any> {

		// get the DispositionReview from storage
		this.loadHelper( dispositionReviewId );

	// get the RetentionSchedule from storage
	var tmp 	= new RetentionScheduleService(this.http).getRetentionSchedule(_retentionScheduleId);

	// assign the RetentionSchedule
	this.dispositionReview.retentionSchedule = tmp;

	// save the DispositionReview
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a RetentionSchedule on a DispositionReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRetentionSchedule( dispositionReviewId ): Observable<any> {

		// get the DispositionReview from storage
		this.loadHelper( dispositionReviewId );

	// assign RetentionSchedule to null
	this.dispositionReview.retentionSchedule = null;

	// save the DispositionReview
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a DispositionReview
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/DispositionReview/update/' + this.dispositionReview;

	return  this.http.post(uri_, this.dispositionReview );
}

	//********************************************************************
	// loadHelper - internal helper to load a DispositionReview
	//********************************************************************	
	loadHelper( id ) {
		this.getDispositionReview(id)
			.subscribe((res : DispositionReview) => {
				this.dispositionReview = res;
			});
	}
}