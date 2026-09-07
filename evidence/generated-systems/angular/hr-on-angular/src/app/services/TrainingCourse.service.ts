import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TrainingCourse} from '../models/TrainingCourse';
import {TrainingEnrollmentService} from '../services/TrainingEnrollment.service';
import {JobProfileService} from '../services/JobProfile.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TrainingCourseService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	trainingCourse : TrainingCourse;

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
	// add a TrainingCourse
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTrainingCourse(code, title, durationHours, Prerequisites, Enrollments, JobProfiles, DeliveryMethod) : Observable<any> {
		const uri_ = this.apiUrl + '/TrainingCourse/create';
		const obj = {
			      		code: code,
      		title: title,
      		durationHours: durationHours,
      		Prerequisites: Prerequisites != null && Prerequisites.length > 0 ? Prerequisites : null,
      		Enrollments: Enrollments != null && Enrollments.length > 0 ? Enrollments : null,
      		JobProfiles: JobProfiles != null && JobProfiles.length > 0 ? JobProfiles : null,
			DeliveryMethod: DeliveryMethod
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TrainingCourse
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTrainingCourse(code, title, durationHours, Prerequisites, Enrollments, JobProfiles, DeliveryMethod, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TrainingCourse/update/' + id;
		const obj = {
				      		code: code,
      		title: title,
      		durationHours: durationHours,
      		Prerequisites: Prerequisites != null && Prerequisites.length > 0 ? Prerequisites : null,
      		Enrollments: Enrollments != null && Enrollments.length > 0 ? Enrollments : null,
      		JobProfiles: JobProfiles != null && JobProfiles.length > 0 ? JobProfiles : null,
			DeliveryMethod: DeliveryMethod
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TrainingCourse
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTrainingCourse(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TrainingCourse/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TrainingCourse
	// returns the results untouched as an Observable TrainingCourse
	// TrainingCourse model
	// delegates via URI
	//********************************************************************
	getTrainingCourse(id) : Observable<TrainingCourse> {
		const uri_ = this.apiUrl + '/TrainingCourse/load/' + id;

		return this.http.get<TrainingCourse>(uri_);
	}
	
	//********************************************************************
	// gets all TrainingCourse
	// returns the results untouched as JSON representation of an
	// Observable array of TrainingCourse models
	// delegates via URI
	//********************************************************************
	getTrainingCourses() : Observable<TrainingCourse[]> {
		const uri_ = this.apiUrl + '/TrainingCourse/';

		return this
			.http.get<TrainingCourse[]>(uri_);
	}
	
		
		//********************************************************************
	// adds one or more prerequisitesIds as a Prerequisites
	// to a TrainingCourse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPrerequisites( trainingCourseId, prerequisitesIds ): Observable<any> {

		// get the TrainingCourse
		this.loadHelper( trainingCourseId );

	// split on a comma with no spaces
	var idList = prerequisitesIds.split(',')

	// iterate over array of prerequisites ids
	idList.forEach(function (id) {
		// read the TrainingCourse
		var trainingCourse = new TrainingCourseService(this.http).getTrainingCourse(id);
		// add the TrainingCourse if not already assigned
		if ( this.trainingCourse.prerequisites.indexOf(trainingCourse) == -1 )
		this.trainingCourse.prerequisites.push(trainingCourse);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more prerequisitesIds as a Prerequisites
	// from a TrainingCourse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePrerequisites( trainingCourseId, prerequisitesIds ): Observable<any> {

		// get the TrainingCourse
		this.loadHelper( trainingCourseId );


	// split on a comma with no spaces
	var idList 					= prerequisitesIds.split(',');
	var prerequisites 	= this.trainingCourse.prerequisites;

	if ( prerequisites != null && prerequisitesIds != null ) {

		// iterate over array of prerequisites ids
		prerequisites.forEach(function (obj) {
			if ( prerequisitesIds.indexOf(obj._id) > -1 ) {
				// remove the TrainingCourse
				this.trainingCourse.prerequisites.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more enrollmentsIds as a Enrollments
	// to a TrainingCourse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addEnrollments( trainingCourseId, enrollmentsIds ): Observable<any> {

		// get the TrainingCourse
		this.loadHelper( trainingCourseId );

	// split on a comma with no spaces
	var idList = enrollmentsIds.split(',')

	// iterate over array of enrollments ids
	idList.forEach(function (id) {
		// read the TrainingEnrollment
		var trainingEnrollment = new TrainingEnrollmentService(this.http).getTrainingEnrollment(id);
		// add the TrainingEnrollment if not already assigned
		if ( this.trainingCourse.enrollments.indexOf(trainingEnrollment) == -1 )
		this.trainingCourse.enrollments.push(trainingEnrollment);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more enrollmentsIds as a Enrollments
	// from a TrainingCourse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeEnrollments( trainingCourseId, enrollmentsIds ): Observable<any> {

		// get the TrainingCourse
		this.loadHelper( trainingCourseId );


	// split on a comma with no spaces
	var idList 					= enrollmentsIds.split(',');
	var enrollments 	= this.trainingCourse.enrollments;

	if ( enrollments != null && enrollmentsIds != null ) {

		// iterate over array of enrollments ids
		enrollments.forEach(function (obj) {
			if ( enrollmentsIds.indexOf(obj._id) > -1 ) {
				// remove the TrainingEnrollment
				this.trainingCourse.enrollments.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more jobProfilesIds as a JobProfiles
	// to a TrainingCourse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addJobProfiles( trainingCourseId, jobProfilesIds ): Observable<any> {

		// get the TrainingCourse
		this.loadHelper( trainingCourseId );

	// split on a comma with no spaces
	var idList = jobProfilesIds.split(',')

	// iterate over array of jobProfiles ids
	idList.forEach(function (id) {
		// read the JobProfile
		var jobProfile = new JobProfileService(this.http).getJobProfile(id);
		// add the JobProfile if not already assigned
		if ( this.trainingCourse.jobProfiles.indexOf(jobProfile) == -1 )
		this.trainingCourse.jobProfiles.push(jobProfile);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more jobProfilesIds as a JobProfiles
	// from a TrainingCourse
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeJobProfiles( trainingCourseId, jobProfilesIds ): Observable<any> {

		// get the TrainingCourse
		this.loadHelper( trainingCourseId );


	// split on a comma with no spaces
	var idList 					= jobProfilesIds.split(',');
	var jobProfiles 	= this.trainingCourse.jobProfiles;

	if ( jobProfiles != null && jobProfilesIds != null ) {

		// iterate over array of jobProfiles ids
		jobProfiles.forEach(function (obj) {
			if ( jobProfilesIds.indexOf(obj._id) > -1 ) {
				// remove the JobProfile
				this.trainingCourse.jobProfiles.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a TrainingCourse
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TrainingCourse/update/' + this.trainingCourse;

	return  this.http.post(uri_, this.trainingCourse );
}

	//********************************************************************
	// loadHelper - internal helper to load a TrainingCourse
	//********************************************************************	
	loadHelper( id ) {
		this.getTrainingCourse(id)
			.subscribe((res : TrainingCourse) => {
				this.trainingCourse = res;
			});
	}
}