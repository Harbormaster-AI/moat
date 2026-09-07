import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Department} from '../models/Department';
import {FacilityService} from '../services/Facility.service';
import {CareTeamService} from '../services/CareTeam.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DepartmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	department : Department;

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
	// add a Department
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDepartment(name, Facility, CareTeams, DepartmentType) : Observable<any> {
		const uri_ = this.apiUrl + '/Department/create';
		const obj = {
			      		name: name,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		CareTeams: CareTeams != null && CareTeams.length > 0 ? CareTeams : null,
			DepartmentType: DepartmentType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDepartment(name, Facility, CareTeams, DepartmentType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Department/update/' + id;
		const obj = {
				      		name: name,
      		Facility: Facility != null && Facility.length > 0 ? Facility : null,
      		CareTeams: CareTeams != null && CareTeams.length > 0 ? CareTeams : null,
			DepartmentType: DepartmentType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDepartment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Department/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Department
	// returns the results untouched as an Observable Department
	// Department model
	// delegates via URI
	//********************************************************************
	getDepartment(id) : Observable<Department> {
		const uri_ = this.apiUrl + '/Department/load/' + id;

		return this.http.get<Department>(uri_);
	}
	
	//********************************************************************
	// gets all Department
	// returns the results untouched as JSON representation of an
	// Observable array of Department models
	// delegates via URI
	//********************************************************************
	getDepartments() : Observable<Department[]> {
		const uri_ = this.apiUrl + '/Department/';

		return this
			.http.get<Department[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Facility on a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignFacility( departmentId, _facilityId ): Observable<any> {

		// get the Department from storage
		this.loadHelper( departmentId );

	// get the Facility from storage
	var tmp 	= new FacilityService(this.http).getFacility(_facilityId);

	// assign the Facility
	this.department.facility = tmp;

	// save the Department
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Facility on a Department
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignFacility( departmentId ): Observable<any> {

		// get the Department from storage
		this.loadHelper( departmentId );

	// assign Facility to null
	this.department.facility = null;

	// save the Department
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more careTeamsIds as a CareTeams
	// to a Department
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addCareTeams( departmentId, careTeamsIds ): Observable<any> {

		// get the Department
		this.loadHelper( departmentId );

	// split on a comma with no spaces
	var idList = careTeamsIds.split(',')

	// iterate over array of careTeams ids
	idList.forEach(function (id) {
		// read the CareTeam
		var careTeam = new CareTeamService(this.http).getCareTeam(id);
		// add the CareTeam if not already assigned
		if ( this.department.careTeams.indexOf(careTeam) == -1 )
		this.department.careTeams.push(careTeam);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more careTeamsIds as a CareTeams
	// from a Department
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeCareTeams( departmentId, careTeamsIds ): Observable<any> {

		// get the Department
		this.loadHelper( departmentId );


	// split on a comma with no spaces
	var idList 					= careTeamsIds.split(',');
	var careTeams 	= this.department.careTeams;

	if ( careTeams != null && careTeamsIds != null ) {

		// iterate over array of careTeams ids
		careTeams.forEach(function (obj) {
			if ( careTeamsIds.indexOf(obj._id) > -1 ) {
				// remove the CareTeam
				this.department.careTeams.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Department
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Department/update/' + this.department;

	return  this.http.post(uri_, this.department );
}

	//********************************************************************
	// loadHelper - internal helper to load a Department
	//********************************************************************	
	loadHelper( id ) {
		this.getDepartment(id)
			.subscribe((res : Department) => {
				this.department = res;
			});
	}
}