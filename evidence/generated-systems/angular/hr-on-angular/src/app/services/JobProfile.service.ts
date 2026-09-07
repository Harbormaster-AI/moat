import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {JobProfile} from '../models/JobProfile';
import {JobFamilyService} from '../services/JobFamily.service';
import {CompetencyService} from '../services/Competency.service';
import {TrainingCourseService} from '../services/TrainingCourse.service';
import {PositionService} from '../services/Position.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class JobProfileService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	jobProfile : JobProfile;

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
	// add a JobProfile
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addJobProfile(title, jobCode, JobFamily, Competencies, TrainingRecommendations, Positions, JobLevel, ExemptStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/JobProfile/create';
		const obj = {
			      		title: title,
      		jobCode: jobCode,
      		JobFamily: JobFamily != null && JobFamily.length > 0 ? JobFamily : null,
      		Competencies: Competencies != null && Competencies.length > 0 ? Competencies : null,
      		TrainingRecommendations: TrainingRecommendations != null && TrainingRecommendations.length > 0 ? TrainingRecommendations : null,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
      		JobLevel: JobLevel,
			ExemptStatus: ExemptStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a JobProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateJobProfile(title, jobCode, JobFamily, Competencies, TrainingRecommendations, Positions, JobLevel, ExemptStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/JobProfile/update/' + id;
		const obj = {
				      		title: title,
      		jobCode: jobCode,
      		JobFamily: JobFamily != null && JobFamily.length > 0 ? JobFamily : null,
      		Competencies: Competencies != null && Competencies.length > 0 ? Competencies : null,
      		TrainingRecommendations: TrainingRecommendations != null && TrainingRecommendations.length > 0 ? TrainingRecommendations : null,
      		Positions: Positions != null && Positions.length > 0 ? Positions : null,
      		JobLevel: JobLevel,
			ExemptStatus: ExemptStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a JobProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteJobProfile(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/JobProfile/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a JobProfile
	// returns the results untouched as an Observable JobProfile
	// JobProfile model
	// delegates via URI
	//********************************************************************
	getJobProfile(id) : Observable<JobProfile> {
		const uri_ = this.apiUrl + '/JobProfile/load/' + id;

		return this.http.get<JobProfile>(uri_);
	}
	
	//********************************************************************
	// gets all JobProfile
	// returns the results untouched as JSON representation of an
	// Observable array of JobProfile models
	// delegates via URI
	//********************************************************************
	getJobProfiles() : Observable<JobProfile[]> {
		const uri_ = this.apiUrl + '/JobProfile/';

		return this
			.http.get<JobProfile[]>(uri_);
	}
	
			//********************************************************************
	// assigns a JobFamily on a JobProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignJobFamily( jobProfileId, _jobFamilyId ): Observable<any> {

		// get the JobProfile from storage
		this.loadHelper( jobProfileId );

	// get the JobFamily from storage
	var tmp 	= new JobFamilyService(this.http).getJobFamily(_jobFamilyId);

	// assign the JobFamily
	this.jobProfile.jobFamily = tmp;

	// save the JobProfile
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a JobFamily on a JobProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignJobFamily( jobProfileId ): Observable<any> {

		// get the JobProfile from storage
		this.loadHelper( jobProfileId );

	// assign JobFamily to null
	this.jobProfile.jobFamily = null;

	// save the JobProfile
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more competenciesIds as a Competencies
	// to a JobProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCompetencies( jobProfileId, competenciesIds ): Observable<any> {

		// get the JobProfile
		this.loadHelper( jobProfileId );

	// split on a comma with no spaces
	var idList = competenciesIds.split(',')

	// iterate over array of competencies ids
	idList.forEach(function (id) {
		// read the Competency
		var competency = new CompetencyService(this.http).getCompetency(id);
		// add the Competency if not already assigned
		if ( this.jobProfile.competencies.indexOf(competency) == -1 )
		this.jobProfile.competencies.push(competency);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more competenciesIds as a Competencies
	// from a JobProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCompetencies( jobProfileId, competenciesIds ): Observable<any> {

		// get the JobProfile
		this.loadHelper( jobProfileId );


	// split on a comma with no spaces
	var idList 					= competenciesIds.split(',');
	var competencies 	= this.jobProfile.competencies;

	if ( competencies != null && competenciesIds != null ) {

		// iterate over array of competencies ids
		competencies.forEach(function (obj) {
			if ( competenciesIds.indexOf(obj._id) > -1 ) {
				// remove the Competency
				this.jobProfile.competencies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more trainingRecommendationsIds as a TrainingRecommendations
	// to a JobProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addTrainingRecommendations( jobProfileId, trainingRecommendationsIds ): Observable<any> {

		// get the JobProfile
		this.loadHelper( jobProfileId );

	// split on a comma with no spaces
	var idList = trainingRecommendationsIds.split(',')

	// iterate over array of trainingRecommendations ids
	idList.forEach(function (id) {
		// read the TrainingCourse
		var trainingCourse = new TrainingCourseService(this.http).getTrainingCourse(id);
		// add the TrainingCourse if not already assigned
		if ( this.jobProfile.trainingRecommendations.indexOf(trainingCourse) == -1 )
		this.jobProfile.trainingRecommendations.push(trainingCourse);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more trainingRecommendationsIds as a TrainingRecommendations
	// from a JobProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeTrainingRecommendations( jobProfileId, trainingRecommendationsIds ): Observable<any> {

		// get the JobProfile
		this.loadHelper( jobProfileId );


	// split on a comma with no spaces
	var idList 					= trainingRecommendationsIds.split(',');
	var trainingRecommendations 	= this.jobProfile.trainingRecommendations;

	if ( trainingRecommendations != null && trainingRecommendationsIds != null ) {

		// iterate over array of trainingRecommendations ids
		trainingRecommendations.forEach(function (obj) {
			if ( trainingRecommendationsIds.indexOf(obj._id) > -1 ) {
				// remove the TrainingCourse
				this.jobProfile.trainingRecommendations.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more positionsIds as a Positions
	// to a JobProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPositions( jobProfileId, positionsIds ): Observable<any> {

		// get the JobProfile
		this.loadHelper( jobProfileId );

	// split on a comma with no spaces
	var idList = positionsIds.split(',')

	// iterate over array of positions ids
	idList.forEach(function (id) {
		// read the Position
		var position = new PositionService(this.http).getPosition(id);
		// add the Position if not already assigned
		if ( this.jobProfile.positions.indexOf(position) == -1 )
		this.jobProfile.positions.push(position);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more positionsIds as a Positions
	// from a JobProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePositions( jobProfileId, positionsIds ): Observable<any> {

		// get the JobProfile
		this.loadHelper( jobProfileId );


	// split on a comma with no spaces
	var idList 					= positionsIds.split(',');
	var positions 	= this.jobProfile.positions;

	if ( positions != null && positionsIds != null ) {

		// iterate over array of positions ids
		positions.forEach(function (obj) {
			if ( positionsIds.indexOf(obj._id) > -1 ) {
				// remove the Position
				this.jobProfile.positions.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a JobProfile
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/JobProfile/update/' + this.jobProfile;

	return  this.http.post(uri_, this.jobProfile );
}

	//********************************************************************
	// loadHelper - internal helper to load a JobProfile
	//********************************************************************	
	loadHelper( id ) {
		this.getJobProfile(id)
			.subscribe((res : JobProfile) => {
				this.jobProfile = res;
			});
	}
}