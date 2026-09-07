import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {PerformanceReview} from '../models/PerformanceReview';
import {EmployeeService} from '../services/Employee.service';
import {PerformanceCycleService} from '../services/PerformanceCycle.service';
import {CompetencyRatingService} from '../services/CompetencyRating.service';
import {GoalService} from '../services/Goal.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class PerformanceReviewService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	performanceReview : PerformanceReview;

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
	// add a PerformanceReview
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addPerformanceReview(reviewNumber, reviewDate, reviewerComments, Employee, Reviewer, Cycle, CompetencyRatings, Goals, Rating, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/PerformanceReview/create';
		const obj = {
			      		reviewNumber: reviewNumber,
      		reviewDate: reviewDate,
      		reviewerComments: reviewerComments,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Reviewer: Reviewer != null && Reviewer.length > 0 ? Reviewer : null,
      		Cycle: Cycle != null && Cycle.length > 0 ? Cycle : null,
      		CompetencyRatings: CompetencyRatings != null && CompetencyRatings.length > 0 ? CompetencyRatings : null,
      		Goals: Goals != null && Goals.length > 0 ? Goals : null,
      		Rating: Rating,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a PerformanceReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updatePerformanceReview(reviewNumber, reviewDate, reviewerComments, Employee, Reviewer, Cycle, CompetencyRatings, Goals, Rating, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/PerformanceReview/update/' + id;
		const obj = {
				      		reviewNumber: reviewNumber,
      		reviewDate: reviewDate,
      		reviewerComments: reviewerComments,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Reviewer: Reviewer != null && Reviewer.length > 0 ? Reviewer : null,
      		Cycle: Cycle != null && Cycle.length > 0 ? Cycle : null,
      		CompetencyRatings: CompetencyRatings != null && CompetencyRatings.length > 0 ? CompetencyRatings : null,
      		Goals: Goals != null && Goals.length > 0 ? Goals : null,
      		Rating: Rating,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a PerformanceReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deletePerformanceReview(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/PerformanceReview/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a PerformanceReview
	// returns the results untouched as an Observable PerformanceReview
	// PerformanceReview model
	// delegates via URI
	//********************************************************************
	getPerformanceReview(id) : Observable<PerformanceReview> {
		const uri_ = this.apiUrl + '/PerformanceReview/load/' + id;

		return this.http.get<PerformanceReview>(uri_);
	}
	
	//********************************************************************
	// gets all PerformanceReview
	// returns the results untouched as JSON representation of an
	// Observable array of PerformanceReview models
	// delegates via URI
	//********************************************************************
	getPerformanceReviews() : Observable<PerformanceReview[]> {
		const uri_ = this.apiUrl + '/PerformanceReview/';

		return this
			.http.get<PerformanceReview[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a PerformanceReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( performanceReviewId, _employeeId ): Observable<any> {

		// get the PerformanceReview from storage
		this.loadHelper( performanceReviewId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.performanceReview.employee = tmp;

	// save the PerformanceReview
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a PerformanceReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( performanceReviewId ): Observable<any> {

		// get the PerformanceReview from storage
		this.loadHelper( performanceReviewId );

	// assign Employee to null
	this.performanceReview.employee = null;

	// save the PerformanceReview
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Reviewer on a PerformanceReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignReviewer( performanceReviewId, _reviewerId ): Observable<any> {

		// get the PerformanceReview from storage
		this.loadHelper( performanceReviewId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_reviewerId);

	// assign the Reviewer
	this.performanceReview.reviewer = tmp;

	// save the PerformanceReview
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Reviewer on a PerformanceReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignReviewer( performanceReviewId ): Observable<any> {

		// get the PerformanceReview from storage
		this.loadHelper( performanceReviewId );

	// assign Reviewer to null
	this.performanceReview.reviewer = null;

	// save the PerformanceReview
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Cycle on a PerformanceReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCycle( performanceReviewId, _cycleId ): Observable<any> {

		// get the PerformanceReview from storage
		this.loadHelper( performanceReviewId );

	// get the PerformanceCycle from storage
	var tmp 	= new PerformanceCycleService(this.http).getPerformanceCycle(_cycleId);

	// assign the Cycle
	this.performanceReview.cycle = tmp;

	// save the PerformanceReview
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Cycle on a PerformanceReview
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCycle( performanceReviewId ): Observable<any> {

		// get the PerformanceReview from storage
		this.loadHelper( performanceReviewId );

	// assign Cycle to null
	this.performanceReview.cycle = null;

	// save the PerformanceReview
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more competencyRatingsIds as a CompetencyRatings
	// to a PerformanceReview
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCompetencyRatings( performanceReviewId, competencyRatingsIds ): Observable<any> {

		// get the PerformanceReview
		this.loadHelper( performanceReviewId );

	// split on a comma with no spaces
	var idList = competencyRatingsIds.split(',')

	// iterate over array of competencyRatings ids
	idList.forEach(function (id) {
		// read the CompetencyRating
		var competencyRating = new CompetencyRatingService(this.http).getCompetencyRating(id);
		// add the CompetencyRating if not already assigned
		if ( this.performanceReview.competencyRatings.indexOf(competencyRating) == -1 )
		this.performanceReview.competencyRatings.push(competencyRating);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more competencyRatingsIds as a CompetencyRatings
	// from a PerformanceReview
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCompetencyRatings( performanceReviewId, competencyRatingsIds ): Observable<any> {

		// get the PerformanceReview
		this.loadHelper( performanceReviewId );


	// split on a comma with no spaces
	var idList 					= competencyRatingsIds.split(',');
	var competencyRatings 	= this.performanceReview.competencyRatings;

	if ( competencyRatings != null && competencyRatingsIds != null ) {

		// iterate over array of competencyRatings ids
		competencyRatings.forEach(function (obj) {
			if ( competencyRatingsIds.indexOf(obj._id) > -1 ) {
				// remove the CompetencyRating
				this.performanceReview.competencyRatings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more goalsIds as a Goals
	// to a PerformanceReview
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addGoals( performanceReviewId, goalsIds ): Observable<any> {

		// get the PerformanceReview
		this.loadHelper( performanceReviewId );

	// split on a comma with no spaces
	var idList = goalsIds.split(',')

	// iterate over array of goals ids
	idList.forEach(function (id) {
		// read the Goal
		var goal = new GoalService(this.http).getGoal(id);
		// add the Goal if not already assigned
		if ( this.performanceReview.goals.indexOf(goal) == -1 )
		this.performanceReview.goals.push(goal);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more goalsIds as a Goals
	// from a PerformanceReview
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeGoals( performanceReviewId, goalsIds ): Observable<any> {

		// get the PerformanceReview
		this.loadHelper( performanceReviewId );


	// split on a comma with no spaces
	var idList 					= goalsIds.split(',');
	var goals 	= this.performanceReview.goals;

	if ( goals != null && goalsIds != null ) {

		// iterate over array of goals ids
		goals.forEach(function (obj) {
			if ( goalsIds.indexOf(obj._id) > -1 ) {
				// remove the Goal
				this.performanceReview.goals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a PerformanceReview
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/PerformanceReview/update/' + this.performanceReview;

	return  this.http.post(uri_, this.performanceReview );
}

	//********************************************************************
	// loadHelper - internal helper to load a PerformanceReview
	//********************************************************************	
	loadHelper( id ) {
		this.getPerformanceReview(id)
			.subscribe((res : PerformanceReview) => {
				this.performanceReview = res;
			});
	}
}