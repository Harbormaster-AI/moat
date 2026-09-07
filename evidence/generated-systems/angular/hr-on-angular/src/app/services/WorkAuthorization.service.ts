import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {WorkAuthorization} from '../models/WorkAuthorization';
import {EmployeeService} from '../services/Employee.service';
import {DocumentService} from '../services/Document.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class WorkAuthorizationService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	workAuthorization : WorkAuthorization;

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
	// add a WorkAuthorization
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addWorkAuthorization(country, expirationDate, Employee, Documents, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/WorkAuthorization/create';
		const obj = {
			      		country: country,
      		expirationDate: expirationDate,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Documents: Documents != null && Documents.length > 0 ? Documents : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a WorkAuthorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateWorkAuthorization(country, expirationDate, Employee, Documents, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/WorkAuthorization/update/' + id;
		const obj = {
				      		country: country,
      		expirationDate: expirationDate,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
      		Documents: Documents != null && Documents.length > 0 ? Documents : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a WorkAuthorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteWorkAuthorization(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/WorkAuthorization/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a WorkAuthorization
	// returns the results untouched as an Observable WorkAuthorization
	// WorkAuthorization model
	// delegates via URI
	//********************************************************************
	getWorkAuthorization(id) : Observable<WorkAuthorization> {
		const uri_ = this.apiUrl + '/WorkAuthorization/load/' + id;

		return this.http.get<WorkAuthorization>(uri_);
	}
	
	//********************************************************************
	// gets all WorkAuthorization
	// returns the results untouched as JSON representation of an
	// Observable array of WorkAuthorization models
	// delegates via URI
	//********************************************************************
	getWorkAuthorizations() : Observable<WorkAuthorization[]> {
		const uri_ = this.apiUrl + '/WorkAuthorization/';

		return this
			.http.get<WorkAuthorization[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Employee on a WorkAuthorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( workAuthorizationId, _employeeId ): Observable<any> {

		// get the WorkAuthorization from storage
		this.loadHelper( workAuthorizationId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.workAuthorization.employee = tmp;

	// save the WorkAuthorization
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a WorkAuthorization
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( workAuthorizationId ): Observable<any> {

		// get the WorkAuthorization from storage
		this.loadHelper( workAuthorizationId );

	// assign Employee to null
	this.workAuthorization.employee = null;

	// save the WorkAuthorization
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more documentsIds as a Documents
	// to a WorkAuthorization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDocuments( workAuthorizationId, documentsIds ): Observable<any> {

		// get the WorkAuthorization
		this.loadHelper( workAuthorizationId );

	// split on a comma with no spaces
	var idList = documentsIds.split(',')

	// iterate over array of documents ids
	idList.forEach(function (id) {
		// read the Document
		var document = new DocumentService(this.http).getDocument(id);
		// add the Document if not already assigned
		if ( this.workAuthorization.documents.indexOf(document) == -1 )
		this.workAuthorization.documents.push(document);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more documentsIds as a Documents
	// from a WorkAuthorization
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDocuments( workAuthorizationId, documentsIds ): Observable<any> {

		// get the WorkAuthorization
		this.loadHelper( workAuthorizationId );


	// split on a comma with no spaces
	var idList 					= documentsIds.split(',');
	var documents 	= this.workAuthorization.documents;

	if ( documents != null && documentsIds != null ) {

		// iterate over array of documents ids
		documents.forEach(function (obj) {
			if ( documentsIds.indexOf(obj._id) > -1 ) {
				// remove the Document
				this.workAuthorization.documents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a WorkAuthorization
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/WorkAuthorization/update/' + this.workAuthorization;

	return  this.http.post(uri_, this.workAuthorization );
}

	//********************************************************************
	// loadHelper - internal helper to load a WorkAuthorization
	//********************************************************************	
	loadHelper( id ) {
		this.getWorkAuthorization(id)
			.subscribe((res : WorkAuthorization) => {
				this.workAuthorization = res;
			});
	}
}