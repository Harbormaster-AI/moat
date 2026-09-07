import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CompetencyRating} from '../models/CompetencyRating';
import {PerformanceReviewService} from '../services/PerformanceReview.service';
import {CompetencyService} from '../services/Competency.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CompetencyRatingService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	competencyRating : CompetencyRating;

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
	// add a CompetencyRating
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCompetencyRating(comment, Review, Competency, Rating) : Observable<any> {
		const uri_ = this.apiUrl + '/CompetencyRating/create';
		const obj = {
			      		comment: comment,
      		Review: Review != null && Review.length > 0 ? Review : null,
      		Competency: Competency != null && Competency.length > 0 ? Competency : null,
			Rating: Rating
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CompetencyRating
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCompetencyRating(comment, Review, Competency, Rating, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CompetencyRating/update/' + id;
		const obj = {
				      		comment: comment,
      		Review: Review != null && Review.length > 0 ? Review : null,
      		Competency: Competency != null && Competency.length > 0 ? Competency : null,
			Rating: Rating
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CompetencyRating
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCompetencyRating(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CompetencyRating/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CompetencyRating
	// returns the results untouched as an Observable CompetencyRating
	// CompetencyRating model
	// delegates via URI
	//********************************************************************
	getCompetencyRating(id) : Observable<CompetencyRating> {
		const uri_ = this.apiUrl + '/CompetencyRating/load/' + id;

		return this.http.get<CompetencyRating>(uri_);
	}
	
	//********************************************************************
	// gets all CompetencyRating
	// returns the results untouched as JSON representation of an
	// Observable array of CompetencyRating models
	// delegates via URI
	//********************************************************************
	getCompetencyRatings() : Observable<CompetencyRating[]> {
		const uri_ = this.apiUrl + '/CompetencyRating/';

		return this
			.http.get<CompetencyRating[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Review on a CompetencyRating
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignReview( competencyRatingId, _reviewId ): Observable<any> {

		// get the CompetencyRating from storage
		this.loadHelper( competencyRatingId );

	// get the PerformanceReview from storage
	var tmp 	= new PerformanceReviewService(this.http).getPerformanceReview(_reviewId);

	// assign the Review
	this.competencyRating.review = tmp;

	// save the CompetencyRating
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Review on a CompetencyRating
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignReview( competencyRatingId ): Observable<any> {

		// get the CompetencyRating from storage
		this.loadHelper( competencyRatingId );

	// assign Review to null
	this.competencyRating.review = null;

	// save the CompetencyRating
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Competency on a CompetencyRating
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCompetency( competencyRatingId, _competencyId ): Observable<any> {

		// get the CompetencyRating from storage
		this.loadHelper( competencyRatingId );

	// get the Competency from storage
	var tmp 	= new CompetencyService(this.http).getCompetency(_competencyId);

	// assign the Competency
	this.competencyRating.competency = tmp;

	// save the CompetencyRating
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Competency on a CompetencyRating
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCompetency( competencyRatingId ): Observable<any> {

		// get the CompetencyRating from storage
		this.loadHelper( competencyRatingId );

	// assign Competency to null
	this.competencyRating.competency = null;

	// save the CompetencyRating
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CompetencyRating
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CompetencyRating/update/' + this.competencyRating;

	return  this.http.post(uri_, this.competencyRating );
}

	//********************************************************************
	// loadHelper - internal helper to load a CompetencyRating
	//********************************************************************	
	loadHelper( id ) {
		this.getCompetencyRating(id)
			.subscribe((res : CompetencyRating) => {
				this.competencyRating = res;
			});
	}
}