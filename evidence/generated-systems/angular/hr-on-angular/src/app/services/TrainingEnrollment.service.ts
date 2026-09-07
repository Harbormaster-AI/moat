import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {TrainingEnrollment} from '../models/TrainingEnrollment';
import {TrainingCourseService} from '../services/TrainingCourse.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class TrainingEnrollmentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	trainingEnrollment : TrainingEnrollment;

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
	// add a TrainingEnrollment
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addTrainingEnrollment(enrollmentNumber, completionDate, score, Course, Employee, Instructor, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/TrainingEnrollment/create';
		const obj = {
			      		enrollmentNumber: enrollmentNumber,
      		completionDate: completionDate,
      		score: score,
      		Course: Course != null && Course.length > 0 ? Course : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Instructor: Instructor != null && Instructor.length > 0 ? Instructor : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a TrainingEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateTrainingEnrollment(enrollmentNumber, completionDate, score, Course, Employee, Instructor, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/TrainingEnrollment/update/' + id;
		const obj = {
				      		enrollmentNumber: enrollmentNumber,
      		completionDate: completionDate,
      		score: score,
      		Course: Course != null && Course.length > 0 ? Course : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Instructor: Instructor != null && Instructor.length > 0 ? Instructor : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a TrainingEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteTrainingEnrollment(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/TrainingEnrollment/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a TrainingEnrollment
	// returns the results untouched as an Observable TrainingEnrollment
	// TrainingEnrollment model
	// delegates via URI
	//********************************************************************
	getTrainingEnrollment(id) : Observable<TrainingEnrollment> {
		const uri_ = this.apiUrl + '/TrainingEnrollment/load/' + id;

		return this.http.get<TrainingEnrollment>(uri_);
	}
	
	//********************************************************************
	// gets all TrainingEnrollment
	// returns the results untouched as JSON representation of an
	// Observable array of TrainingEnrollment models
	// delegates via URI
	//********************************************************************
	getTrainingEnrollments() : Observable<TrainingEnrollment[]> {
		const uri_ = this.apiUrl + '/TrainingEnrollment/';

		return this
			.http.get<TrainingEnrollment[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Course on a TrainingEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCourse( trainingEnrollmentId, _courseId ): Observable<any> {

		// get the TrainingEnrollment from storage
		this.loadHelper( trainingEnrollmentId );

	// get the TrainingCourse from storage
	var tmp 	= new TrainingCourseService(this.http).getTrainingCourse(_courseId);

	// assign the Course
	this.trainingEnrollment.course = tmp;

	// save the TrainingEnrollment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Course on a TrainingEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCourse( trainingEnrollmentId ): Observable<any> {

		// get the TrainingEnrollment from storage
		this.loadHelper( trainingEnrollmentId );

	// assign Course to null
	this.trainingEnrollment.course = null;

	// save the TrainingEnrollment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Employee on a TrainingEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( trainingEnrollmentId, _employeeId ): Observable<any> {

		// get the TrainingEnrollment from storage
		this.loadHelper( trainingEnrollmentId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.trainingEnrollment.employee = tmp;

	// save the TrainingEnrollment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a TrainingEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( trainingEnrollmentId ): Observable<any> {

		// get the TrainingEnrollment from storage
		this.loadHelper( trainingEnrollmentId );

	// assign Employee to null
	this.trainingEnrollment.employee = null;

	// save the TrainingEnrollment
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Instructor on a TrainingEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignInstructor( trainingEnrollmentId, _instructorId ): Observable<any> {

		// get the TrainingEnrollment from storage
		this.loadHelper( trainingEnrollmentId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_instructorId);

	// assign the Instructor
	this.trainingEnrollment.instructor = tmp;

	// save the TrainingEnrollment
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Instructor on a TrainingEnrollment
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignInstructor( trainingEnrollmentId ): Observable<any> {

		// get the TrainingEnrollment from storage
		this.loadHelper( trainingEnrollmentId );

	// assign Instructor to null
	this.trainingEnrollment.instructor = null;

	// save the TrainingEnrollment
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a TrainingEnrollment
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/TrainingEnrollment/update/' + this.trainingEnrollment;

	return  this.http.post(uri_, this.trainingEnrollment );
}

	//********************************************************************
	// loadHelper - internal helper to load a TrainingEnrollment
	//********************************************************************	
	loadHelper( id ) {
		this.getTrainingEnrollment(id)
			.subscribe((res : TrainingEnrollment) => {
				this.trainingEnrollment = res;
			});
	}
}