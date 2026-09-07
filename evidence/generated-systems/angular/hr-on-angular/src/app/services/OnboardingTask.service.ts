import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {OnboardingTask} from '../models/OnboardingTask';
import {EmployeeService} from '../services/Employee.service';
import {OfferService} from '../services/Offer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OnboardingTaskService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	onboardingTask : OnboardingTask;

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
	// add a OnboardingTask
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOnboardingTask(taskNumber, name, dueDate, Employee, AssignedTo, Dependencies, RelatedOffer, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/OnboardingTask/create';
		const obj = {
			      		taskNumber: taskNumber,
      		name: name,
      		dueDate: dueDate,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		AssignedTo: AssignedTo != null && AssignedTo.length > 0 ? AssignedTo : null,
      		Dependencies: Dependencies != null && Dependencies.length > 0 ? Dependencies : null,
      		RelatedOffer: RelatedOffer != null && RelatedOffer.length > 0 ? RelatedOffer : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a OnboardingTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOnboardingTask(taskNumber, name, dueDate, Employee, AssignedTo, Dependencies, RelatedOffer, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/OnboardingTask/update/' + id;
		const obj = {
				      		taskNumber: taskNumber,
      		name: name,
      		dueDate: dueDate,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		AssignedTo: AssignedTo != null && AssignedTo.length > 0 ? AssignedTo : null,
      		Dependencies: Dependencies != null && Dependencies.length > 0 ? Dependencies : null,
      		RelatedOffer: RelatedOffer != null && RelatedOffer.length > 0 ? RelatedOffer : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a OnboardingTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOnboardingTask(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/OnboardingTask/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a OnboardingTask
	// returns the results untouched as an Observable OnboardingTask
	// OnboardingTask model
	// delegates via URI
	//********************************************************************
	getOnboardingTask(id) : Observable<OnboardingTask> {
		const uri_ = this.apiUrl + '/OnboardingTask/load/' + id;

		return this.http.get<OnboardingTask>(uri_);
	}
	
	//********************************************************************
	// gets all OnboardingTask
	// returns the results untouched as JSON representation of an
	// Observable array of OnboardingTask models
	// delegates via URI
	//********************************************************************
	getOnboardingTasks() : Observable<OnboardingTask[]> {
		const uri_ = this.apiUrl + '/OnboardingTask/';

		return this
			.http.get<OnboardingTask[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a OnboardingTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( onboardingTaskId, _employeeId ): Observable<any> {

		// get the OnboardingTask from storage
		this.loadHelper( onboardingTaskId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.onboardingTask.employee = tmp;

	// save the OnboardingTask
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a OnboardingTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( onboardingTaskId ): Observable<any> {

		// get the OnboardingTask from storage
		this.loadHelper( onboardingTaskId );

	// assign Employee to null
	this.onboardingTask.employee = null;

	// save the OnboardingTask
	return this.saveHelper();
}

		//********************************************************************
	// assigns a AssignedTo on a OnboardingTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAssignedTo( onboardingTaskId, _assignedToId ): Observable<any> {

		// get the OnboardingTask from storage
		this.loadHelper( onboardingTaskId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_assignedToId);

	// assign the AssignedTo
	this.onboardingTask.assignedTo = tmp;

	// save the OnboardingTask
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a AssignedTo on a OnboardingTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAssignedTo( onboardingTaskId ): Observable<any> {

		// get the OnboardingTask from storage
		this.loadHelper( onboardingTaskId );

	// assign AssignedTo to null
	this.onboardingTask.assignedTo = null;

	// save the OnboardingTask
	return this.saveHelper();
}

		//********************************************************************
	// assigns a RelatedOffer on a OnboardingTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRelatedOffer( onboardingTaskId, _relatedOfferId ): Observable<any> {

		// get the OnboardingTask from storage
		this.loadHelper( onboardingTaskId );

	// get the Offer from storage
	var tmp 	= new OfferService(this.http).getOffer(_relatedOfferId);

	// assign the RelatedOffer
	this.onboardingTask.relatedOffer = tmp;

	// save the OnboardingTask
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a RelatedOffer on a OnboardingTask
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRelatedOffer( onboardingTaskId ): Observable<any> {

		// get the OnboardingTask from storage
		this.loadHelper( onboardingTaskId );

	// assign RelatedOffer to null
	this.onboardingTask.relatedOffer = null;

	// save the OnboardingTask
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more dependenciesIds as a Dependencies
	// to a OnboardingTask
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDependencies( onboardingTaskId, dependenciesIds ): Observable<any> {

		// get the OnboardingTask
		this.loadHelper( onboardingTaskId );

	// split on a comma with no spaces
	var idList = dependenciesIds.split(',')

	// iterate over array of dependencies ids
	idList.forEach(function (id) {
		// read the OnboardingTask
		var onboardingTask = new OnboardingTaskService(this.http).getOnboardingTask(id);
		// add the OnboardingTask if not already assigned
		if ( this.onboardingTask.dependencies.indexOf(onboardingTask) == -1 )
		this.onboardingTask.dependencies.push(onboardingTask);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more dependenciesIds as a Dependencies
	// from a OnboardingTask
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDependencies( onboardingTaskId, dependenciesIds ): Observable<any> {

		// get the OnboardingTask
		this.loadHelper( onboardingTaskId );


	// split on a comma with no spaces
	var idList 					= dependenciesIds.split(',');
	var dependencies 	= this.onboardingTask.dependencies;

	if ( dependencies != null && dependenciesIds != null ) {

		// iterate over array of dependencies ids
		dependencies.forEach(function (obj) {
			if ( dependenciesIds.indexOf(obj._id) > -1 ) {
				// remove the OnboardingTask
				this.onboardingTask.dependencies.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a OnboardingTask
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/OnboardingTask/update/' + this.onboardingTask;

	return  this.http.post(uri_, this.onboardingTask );
}

	//********************************************************************
	// loadHelper - internal helper to load a OnboardingTask
	//********************************************************************	
	loadHelper( id ) {
		this.getOnboardingTask(id)
			.subscribe((res : OnboardingTask) => {
				this.onboardingTask = res;
			});
	}
}