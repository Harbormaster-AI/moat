import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {JobFamily} from '../models/JobFamily';
import {OrganizationService} from '../services/Organization.service';
import {JobProfileService} from '../services/JobProfile.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class JobFamilyService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	jobFamily : JobFamily;

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
	// add a JobFamily
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addJobFamily(name, description, Organization, JobProfiles) : Observable<any> {
		const uri_ = this.apiUrl + '/JobFamily/create';
		const obj = {
			      		name: name,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
			JobProfiles: JobProfiles != null && JobProfiles.length > 0 ? JobProfiles : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a JobFamily
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateJobFamily(name, description, Organization, JobProfiles, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/JobFamily/update/' + id;
		const obj = {
				      		name: name,
      		description: description,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
			JobProfiles: JobProfiles != null && JobProfiles.length > 0 ? JobProfiles : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a JobFamily
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteJobFamily(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/JobFamily/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a JobFamily
	// returns the results untouched as an Observable JobFamily
	// JobFamily model
	// delegates via URI
	//********************************************************************
	getJobFamily(id) : Observable<JobFamily> {
		const uri_ = this.apiUrl + '/JobFamily/load/' + id;

		return this.http.get<JobFamily>(uri_);
	}
	
	//********************************************************************
	// gets all JobFamily
	// returns the results untouched as JSON representation of an
	// Observable array of JobFamily models
	// delegates via URI
	//********************************************************************
	getJobFamilys() : Observable<JobFamily[]> {
		const uri_ = this.apiUrl + '/JobFamily/';

		return this
			.http.get<JobFamily[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a JobFamily
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( jobFamilyId, _organizationId ): Observable<any> {

		// get the JobFamily from storage
		this.loadHelper( jobFamilyId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.jobFamily.organization = tmp;

	// save the JobFamily
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a JobFamily
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( jobFamilyId ): Observable<any> {

		// get the JobFamily from storage
		this.loadHelper( jobFamilyId );

	// assign Organization to null
	this.jobFamily.organization = null;

	// save the JobFamily
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more jobProfilesIds as a JobProfiles
	// to a JobFamily
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addJobProfiles( jobFamilyId, jobProfilesIds ): Observable<any> {

		// get the JobFamily
		this.loadHelper( jobFamilyId );

	// split on a comma with no spaces
	var idList = jobProfilesIds.split(',')

	// iterate over array of jobProfiles ids
	idList.forEach(function (id) {
		// read the JobProfile
		var jobProfile = new JobProfileService(this.http).getJobProfile(id);
		// add the JobProfile if not already assigned
		if ( this.jobFamily.jobProfiles.indexOf(jobProfile) == -1 )
		this.jobFamily.jobProfiles.push(jobProfile);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more jobProfilesIds as a JobProfiles
	// from a JobFamily
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeJobProfiles( jobFamilyId, jobProfilesIds ): Observable<any> {

		// get the JobFamily
		this.loadHelper( jobFamilyId );


	// split on a comma with no spaces
	var idList 					= jobProfilesIds.split(',');
	var jobProfiles 	= this.jobFamily.jobProfiles;

	if ( jobProfiles != null && jobProfilesIds != null ) {

		// iterate over array of jobProfiles ids
		jobProfiles.forEach(function (obj) {
			if ( jobProfilesIds.indexOf(obj._id) > -1 ) {
				// remove the JobProfile
				this.jobFamily.jobProfiles.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a JobFamily
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/JobFamily/update/' + this.jobFamily;

	return  this.http.post(uri_, this.jobFamily );
}

	//********************************************************************
	// loadHelper - internal helper to load a JobFamily
	//********************************************************************	
	loadHelper( id ) {
		this.getJobFamily(id)
			.subscribe((res : JobFamily) => {
				this.jobFamily = res;
			});
	}
}