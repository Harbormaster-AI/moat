import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Screening} from '../models/Screening';
import {JobApplicationService} from '../services/JobApplication.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ScreeningService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	screening : Screening;

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
	// add a Screening
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addScreening(name, completedDate, Application, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Screening/create';
		const obj = {
			      		name: name,
      		completedDate: completedDate,
      		Application: Application != null && Application.length > 0 ? Application : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Screening
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateScreening(name, completedDate, Application, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Screening/update/' + id;
		const obj = {
				      		name: name,
      		completedDate: completedDate,
      		Application: Application != null && Application.length > 0 ? Application : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Screening
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteScreening(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Screening/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Screening
	// returns the results untouched as an Observable Screening
	// Screening model
	// delegates via URI
	//********************************************************************
	getScreening(id) : Observable<Screening> {
		const uri_ = this.apiUrl + '/Screening/load/' + id;

		return this.http.get<Screening>(uri_);
	}
	
	//********************************************************************
	// gets all Screening
	// returns the results untouched as JSON representation of an
	// Observable array of Screening models
	// delegates via URI
	//********************************************************************
	getScreenings() : Observable<Screening[]> {
		const uri_ = this.apiUrl + '/Screening/';

		return this
			.http.get<Screening[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Application on a Screening
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignApplication( screeningId, _applicationId ): Observable<any> {

		// get the Screening from storage
		this.loadHelper( screeningId );

	// get the JobApplication from storage
	var tmp 	= new JobApplicationService(this.http).getJobApplication(_applicationId);

	// assign the Application
	this.screening.application = tmp;

	// save the Screening
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Application on a Screening
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignApplication( screeningId ): Observable<any> {

		// get the Screening from storage
		this.loadHelper( screeningId );

	// assign Application to null
	this.screening.application = null;

	// save the Screening
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Screening
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Screening/update/' + this.screening;

	return  this.http.post(uri_, this.screening );
}

	//********************************************************************
	// loadHelper - internal helper to load a Screening
	//********************************************************************	
	loadHelper( id ) {
		this.getScreening(id)
			.subscribe((res : Screening) => {
				this.screening = res;
			});
	}
}