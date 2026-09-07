import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Competency} from '../models/Competency';
import {JobProfileService} from '../services/JobProfile.service';
import {CompetencyRatingService} from '../services/CompetencyRating.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CompetencyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	competency : Competency;

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
	// add a Competency
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCompetency(name, category, JobProfiles, CompetencyRatings) : Observable<any> {
		const uri_ = this.apiUrl + '/Competency/create';
		const obj = {
			      		name: name,
      		category: category,
      		JobProfiles: JobProfiles != null && JobProfiles.length > 0 ? JobProfiles : null,
			CompetencyRatings: CompetencyRatings != null && CompetencyRatings.length > 0 ? CompetencyRatings : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Competency
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCompetency(name, category, JobProfiles, CompetencyRatings, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Competency/update/' + id;
		const obj = {
				      		name: name,
      		category: category,
      		JobProfiles: JobProfiles != null && JobProfiles.length > 0 ? JobProfiles : null,
			CompetencyRatings: CompetencyRatings != null && CompetencyRatings.length > 0 ? CompetencyRatings : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Competency
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCompetency(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Competency/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Competency
	// returns the results untouched as an Observable Competency
	// Competency model
	// delegates via URI
	//********************************************************************
	getCompetency(id) : Observable<Competency> {
		const uri_ = this.apiUrl + '/Competency/load/' + id;

		return this.http.get<Competency>(uri_);
	}
	
	//********************************************************************
	// gets all Competency
	// returns the results untouched as JSON representation of an
	// Observable array of Competency models
	// delegates via URI
	//********************************************************************
	getCompetencys() : Observable<Competency[]> {
		const uri_ = this.apiUrl + '/Competency/';

		return this
			.http.get<Competency[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more jobProfilesIds as a JobProfiles
	// to a Competency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addJobProfiles( competencyId, jobProfilesIds ): Observable<any> {

		// get the Competency
		this.loadHelper( competencyId );

	// split on a comma with no spaces
	var idList = jobProfilesIds.split(',')

	// iterate over array of jobProfiles ids
	idList.forEach(function (id) {
		// read the JobProfile
		var jobProfile = new JobProfileService(this.http).getJobProfile(id);
		// add the JobProfile if not already assigned
		if ( this.competency.jobProfiles.indexOf(jobProfile) == -1 )
		this.competency.jobProfiles.push(jobProfile);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more jobProfilesIds as a JobProfiles
	// from a Competency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeJobProfiles( competencyId, jobProfilesIds ): Observable<any> {

		// get the Competency
		this.loadHelper( competencyId );


	// split on a comma with no spaces
	var idList 					= jobProfilesIds.split(',');
	var jobProfiles 	= this.competency.jobProfiles;

	if ( jobProfiles != null && jobProfilesIds != null ) {

		// iterate over array of jobProfiles ids
		jobProfiles.forEach(function (obj) {
			if ( jobProfilesIds.indexOf(obj._id) > -1 ) {
				// remove the JobProfile
				this.competency.jobProfiles.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more competencyRatingsIds as a CompetencyRatings
	// to a Competency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCompetencyRatings( competencyId, competencyRatingsIds ): Observable<any> {

		// get the Competency
		this.loadHelper( competencyId );

	// split on a comma with no spaces
	var idList = competencyRatingsIds.split(',')

	// iterate over array of competencyRatings ids
	idList.forEach(function (id) {
		// read the CompetencyRating
		var competencyRating = new CompetencyRatingService(this.http).getCompetencyRating(id);
		// add the CompetencyRating if not already assigned
		if ( this.competency.competencyRatings.indexOf(competencyRating) == -1 )
		this.competency.competencyRatings.push(competencyRating);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more competencyRatingsIds as a CompetencyRatings
	// from a Competency
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCompetencyRatings( competencyId, competencyRatingsIds ): Observable<any> {

		// get the Competency
		this.loadHelper( competencyId );


	// split on a comma with no spaces
	var idList 					= competencyRatingsIds.split(',');
	var competencyRatings 	= this.competency.competencyRatings;

	if ( competencyRatings != null && competencyRatingsIds != null ) {

		// iterate over array of competencyRatings ids
		competencyRatings.forEach(function (obj) {
			if ( competencyRatingsIds.indexOf(obj._id) > -1 ) {
				// remove the CompetencyRating
				this.competency.competencyRatings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Competency
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Competency/update/' + this.competency;

	return  this.http.post(uri_, this.competency );
}

	//********************************************************************
	// loadHelper - internal helper to load a Competency
	//********************************************************************	
	loadHelper( id ) {
		this.getCompetency(id)
			.subscribe((res : Competency) => {
				this.competency = res;
			});
	}
}