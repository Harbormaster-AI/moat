import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Certification} from '../models/Certification';
import {EmployeeService} from '../services/Employee.service';
import {TrainingCourseService} from '../services/TrainingCourse.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CertificationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	certification : Certification;

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
	// add a Certification
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCertification(name, issuer, validFrom, validTo, credentialId, Employee, Course) : Observable<any> {
		const uri_ = this.apiUrl + '/Certification/create';
		const obj = {
			      		name: name,
      		issuer: issuer,
      		validFrom: validFrom,
      		validTo: validTo,
      		credentialId: credentialId,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			Course: Course != null && Course.length > 0 ? Course : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Certification
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCertification(name, issuer, validFrom, validTo, credentialId, Employee, Course, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Certification/update/' + id;
		const obj = {
				      		name: name,
      		issuer: issuer,
      		validFrom: validFrom,
      		validTo: validTo,
      		credentialId: credentialId,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			Course: Course != null && Course.length > 0 ? Course : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Certification
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCertification(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Certification/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Certification
	// returns the results untouched as an Observable Certification
	// Certification model
	// delegates via URI
	//********************************************************************
	getCertification(id) : Observable<Certification> {
		const uri_ = this.apiUrl + '/Certification/load/' + id;

		return this.http.get<Certification>(uri_);
	}
	
	//********************************************************************
	// gets all Certification
	// returns the results untouched as JSON representation of an
	// Observable array of Certification models
	// delegates via URI
	//********************************************************************
	getCertifications() : Observable<Certification[]> {
		const uri_ = this.apiUrl + '/Certification/';

		return this
			.http.get<Certification[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a Certification
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( certificationId, _employeeId ): Observable<any> {

		// get the Certification from storage
		this.loadHelper( certificationId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.certification.employee = tmp;

	// save the Certification
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a Certification
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( certificationId ): Observable<any> {

		// get the Certification from storage
		this.loadHelper( certificationId );

	// assign Employee to null
	this.certification.employee = null;

	// save the Certification
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Course on a Certification
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCourse( certificationId, _courseId ): Observable<any> {

		// get the Certification from storage
		this.loadHelper( certificationId );

	// get the TrainingCourse from storage
	var tmp 	= new TrainingCourseService(this.http).getTrainingCourse(_courseId);

	// assign the Course
	this.certification.course = tmp;

	// save the Certification
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Course on a Certification
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCourse( certificationId ): Observable<any> {

		// get the Certification from storage
		this.loadHelper( certificationId );

	// assign Course to null
	this.certification.course = null;

	// save the Certification
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Certification
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Certification/update/' + this.certification;

	return  this.http.post(uri_, this.certification );
}

	//********************************************************************
	// loadHelper - internal helper to load a Certification
	//********************************************************************	
	loadHelper( id ) {
		this.getCertification(id)
			.subscribe((res : Certification) => {
				this.certification = res;
			});
	}
}