import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PerformanceCycle} from '../models/PerformanceCycle';
import {OrganizationService} from '../services/Organization.service';
import {PerformanceReviewService} from '../services/PerformanceReview.service';
import {GoalService} from '../services/Goal.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PerformanceCycleService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	performanceCycle : PerformanceCycle;

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
	// add a PerformanceCycle
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPerformanceCycle(name, startDate, endDate, Organization, Reviews, Goals, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PerformanceCycle/create';
		const obj = {
			      		name: name,
      		startDate: startDate,
      		endDate: endDate,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Reviews: Reviews != null && Reviews.length > 0 ? Reviews : null,
      		Goals: Goals != null && Goals.length > 0 ? Goals : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PerformanceCycle
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePerformanceCycle(name, startDate, endDate, Organization, Reviews, Goals, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PerformanceCycle/update/' + id;
		const obj = {
				      		name: name,
      		startDate: startDate,
      		endDate: endDate,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Reviews: Reviews != null && Reviews.length > 0 ? Reviews : null,
      		Goals: Goals != null && Goals.length > 0 ? Goals : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PerformanceCycle
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePerformanceCycle(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PerformanceCycle/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PerformanceCycle
	// returns the results untouched as an Observable PerformanceCycle
	// PerformanceCycle model
	// delegates via URI
	//********************************************************************
	getPerformanceCycle(id) : Observable<PerformanceCycle> {
		const uri_ = this.apiUrl + '/PerformanceCycle/load/' + id;

		return this.http.get<PerformanceCycle>(uri_);
	}
	
	//********************************************************************
	// gets all PerformanceCycle
	// returns the results untouched as JSON representation of an
	// Observable array of PerformanceCycle models
	// delegates via URI
	//********************************************************************
	getPerformanceCycles() : Observable<PerformanceCycle[]> {
		const uri_ = this.apiUrl + '/PerformanceCycle/';

		return this
			.http.get<PerformanceCycle[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a PerformanceCycle
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( performanceCycleId, _organizationId ): Observable<any> {

		// get the PerformanceCycle from storage
		this.loadHelper( performanceCycleId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.performanceCycle.organization = tmp;

	// save the PerformanceCycle
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a PerformanceCycle
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( performanceCycleId ): Observable<any> {

		// get the PerformanceCycle from storage
		this.loadHelper( performanceCycleId );

	// assign Organization to null
	this.performanceCycle.organization = null;

	// save the PerformanceCycle
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more reviewsIds as a Reviews
	// to a PerformanceCycle
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addReviews( performanceCycleId, reviewsIds ): Observable<any> {

		// get the PerformanceCycle
		this.loadHelper( performanceCycleId );

	// split on a comma with no spaces
	var idList = reviewsIds.split(',')

	// iterate over array of reviews ids
	idList.forEach(function (id) {
		// read the PerformanceReview
		var performanceReview = new PerformanceReviewService(this.http).getPerformanceReview(id);
		// add the PerformanceReview if not already assigned
		if ( this.performanceCycle.reviews.indexOf(performanceReview) == -1 )
		this.performanceCycle.reviews.push(performanceReview);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more reviewsIds as a Reviews
	// from a PerformanceCycle
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeReviews( performanceCycleId, reviewsIds ): Observable<any> {

		// get the PerformanceCycle
		this.loadHelper( performanceCycleId );


	// split on a comma with no spaces
	var idList 					= reviewsIds.split(',');
	var reviews 	= this.performanceCycle.reviews;

	if ( reviews != null && reviewsIds != null ) {

		// iterate over array of reviews ids
		reviews.forEach(function (obj) {
			if ( reviewsIds.indexOf(obj._id) > -1 ) {
				// remove the PerformanceReview
				this.performanceCycle.reviews.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more goalsIds as a Goals
	// to a PerformanceCycle
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addGoals( performanceCycleId, goalsIds ): Observable<any> {

		// get the PerformanceCycle
		this.loadHelper( performanceCycleId );

	// split on a comma with no spaces
	var idList = goalsIds.split(',')

	// iterate over array of goals ids
	idList.forEach(function (id) {
		// read the Goal
		var goal = new GoalService(this.http).getGoal(id);
		// add the Goal if not already assigned
		if ( this.performanceCycle.goals.indexOf(goal) == -1 )
		this.performanceCycle.goals.push(goal);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more goalsIds as a Goals
	// from a PerformanceCycle
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeGoals( performanceCycleId, goalsIds ): Observable<any> {

		// get the PerformanceCycle
		this.loadHelper( performanceCycleId );


	// split on a comma with no spaces
	var idList 					= goalsIds.split(',');
	var goals 	= this.performanceCycle.goals;

	if ( goals != null && goalsIds != null ) {

		// iterate over array of goals ids
		goals.forEach(function (obj) {
			if ( goalsIds.indexOf(obj._id) > -1 ) {
				// remove the Goal
				this.performanceCycle.goals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PerformanceCycle
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PerformanceCycle/update/' + this.performanceCycle;

	return  this.http.post(uri_, this.performanceCycle );
}

	//********************************************************************
	// loadHelper - internal helper to load a PerformanceCycle
	//********************************************************************	
	loadHelper( id ) {
		this.getPerformanceCycle(id)
			.subscribe((res : PerformanceCycle) => {
				this.performanceCycle = res;
			});
	}
}