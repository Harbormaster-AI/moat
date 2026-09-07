import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Goal} from '../models/Goal';
import {EmployeeService} from '../services/Employee.service';
import {PerformanceCycleService} from '../services/PerformanceCycle.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class GoalService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	goal : Goal;

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
	// add a Goal
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addGoal(title, description, targetDate, weight, Employee, Cycle, ParentGoal, ChildGoals, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Goal/create';
		const obj = {
			      		title: title,
      		description: description,
      		targetDate: targetDate,
      		weight: weight,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Cycle: Cycle != null && Cycle.length > 0 ? Cycle : null,
      		ParentGoal: ParentGoal != null && ParentGoal.length > 0 ? ParentGoal : null,
      		ChildGoals: ChildGoals != null && ChildGoals.length > 0 ? ChildGoals : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Goal
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateGoal(title, description, targetDate, weight, Employee, Cycle, ParentGoal, ChildGoals, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Goal/update/' + id;
		const obj = {
				      		title: title,
      		description: description,
      		targetDate: targetDate,
      		weight: weight,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Cycle: Cycle != null && Cycle.length > 0 ? Cycle : null,
      		ParentGoal: ParentGoal != null && ParentGoal.length > 0 ? ParentGoal : null,
      		ChildGoals: ChildGoals != null && ChildGoals.length > 0 ? ChildGoals : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Goal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteGoal(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Goal/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Goal
	// returns the results untouched as an Observable Goal
	// Goal model
	// delegates via URI
	//********************************************************************
	getGoal(id) : Observable<Goal> {
		const uri_ = this.apiUrl + '/Goal/load/' + id;

		return this.http.get<Goal>(uri_);
	}
	
	//********************************************************************
	// gets all Goal
	// returns the results untouched as JSON representation of an
	// Observable array of Goal models
	// delegates via URI
	//********************************************************************
	getGoals() : Observable<Goal[]> {
		const uri_ = this.apiUrl + '/Goal/';

		return this
			.http.get<Goal[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a Goal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( goalId, _employeeId ): Observable<any> {

		// get the Goal from storage
		this.loadHelper( goalId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.goal.employee = tmp;

	// save the Goal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a Goal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( goalId ): Observable<any> {

		// get the Goal from storage
		this.loadHelper( goalId );

	// assign Employee to null
	this.goal.employee = null;

	// save the Goal
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Cycle on a Goal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCycle( goalId, _cycleId ): Observable<any> {

		// get the Goal from storage
		this.loadHelper( goalId );

	// get the PerformanceCycle from storage
	var tmp 	= new PerformanceCycleService(this.http).getPerformanceCycle(_cycleId);

	// assign the Cycle
	this.goal.cycle = tmp;

	// save the Goal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Cycle on a Goal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCycle( goalId ): Observable<any> {

		// get the Goal from storage
		this.loadHelper( goalId );

	// assign Cycle to null
	this.goal.cycle = null;

	// save the Goal
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ParentGoal on a Goal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignParentGoal( goalId, _parentGoalId ): Observable<any> {

		// get the Goal from storage
		this.loadHelper( goalId );

	// get the Goal from storage
	var tmp 	= new GoalService(this.http).getGoal(_parentGoalId);

	// assign the ParentGoal
	this.goal.parentGoal = tmp;

	// save the Goal
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ParentGoal on a Goal
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignParentGoal( goalId ): Observable<any> {

		// get the Goal from storage
		this.loadHelper( goalId );

	// assign ParentGoal to null
	this.goal.parentGoal = null;

	// save the Goal
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more childGoalsIds as a ChildGoals
	// to a Goal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addChildGoals( goalId, childGoalsIds ): Observable<any> {

		// get the Goal
		this.loadHelper( goalId );

	// split on a comma with no spaces
	var idList = childGoalsIds.split(',')

	// iterate over array of childGoals ids
	idList.forEach(function (id) {
		// read the Goal
		var goal = new GoalService(this.http).getGoal(id);
		// add the Goal if not already assigned
		if ( this.goal.childGoals.indexOf(goal) == -1 )
		this.goal.childGoals.push(goal);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more childGoalsIds as a ChildGoals
	// from a Goal
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeChildGoals( goalId, childGoalsIds ): Observable<any> {

		// get the Goal
		this.loadHelper( goalId );


	// split on a comma with no spaces
	var idList 					= childGoalsIds.split(',');
	var childGoals 	= this.goal.childGoals;

	if ( childGoals != null && childGoalsIds != null ) {

		// iterate over array of childGoals ids
		childGoals.forEach(function (obj) {
			if ( childGoalsIds.indexOf(obj._id) > -1 ) {
				// remove the Goal
				this.goal.childGoals.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Goal
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Goal/update/' + this.goal;

	return  this.http.post(uri_, this.goal );
}

	//********************************************************************
	// loadHelper - internal helper to load a Goal
	//********************************************************************	
	loadHelper( id ) {
		this.getGoal(id)
			.subscribe((res : Goal) => {
				this.goal = res;
			});
	}
}