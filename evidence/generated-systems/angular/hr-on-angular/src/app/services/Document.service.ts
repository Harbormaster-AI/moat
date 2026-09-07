import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Document} from '../models/Document';
import {CandidateService} from '../services/Candidate.service';
import {EmployeeService} from '../services/Employee.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class DocumentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	document : Document;

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
	// add a Document
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addDocument(name, fileUrl, uploadedDate, Candidate, Employee, DocumentType) : Observable<any> {
		const uri_ = this.apiUrl + '/Document/create';
		const obj = {
			      		name: name,
      		fileUrl: fileUrl,
      		uploadedDate: uploadedDate,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			DocumentType: DocumentType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDocument(name, fileUrl, uploadedDate, Candidate, Employee, DocumentType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Document/update/' + id;
		const obj = {
				      		name: name,
      		fileUrl: fileUrl,
      		uploadedDate: uploadedDate,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		Employee: Employee != null && Employee.length > 0 ? Employee : null,
			DocumentType: DocumentType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteDocument(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Document/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Document
	// returns the results untouched as an Observable Document
	// Document model
	// delegates via URI
	//********************************************************************
	getDocument(id) : Observable<Document> {
		const uri_ = this.apiUrl + '/Document/load/' + id;

		return this.http.get<Document>(uri_);
	}
	
	//********************************************************************
	// gets all Document
	// returns the results untouched as JSON representation of an
	// Observable array of Document models
	// delegates via URI
	//********************************************************************
	getDocuments() : Observable<Document[]> {
		const uri_ = this.apiUrl + '/Document/';

		return this
			.http.get<Document[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Candidate on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCandidate( documentId, _candidateId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// get the Candidate from storage
	var tmp 	= new CandidateService(this.http).getCandidate(_candidateId);

	// assign the Candidate
	this.document.candidate = tmp;

	// save the Document
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Candidate on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCandidate( documentId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// assign Candidate to null
	this.document.candidate = null;

	// save the Document
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Employee on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEmployee( documentId, _employeeId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_employeeId);

	// assign the Employee
	this.document.employee = tmp;

	// save the Document
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Employee on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEmployee( documentId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// assign Employee to null
	this.document.employee = null;

	// save the Document
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Document
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Document/update/' + this.document;

	return  this.http.post(uri_, this.document );
}

	//********************************************************************
	// loadHelper - internal helper to load a Document
	//********************************************************************	
	loadHelper( id ) {
		this.getDocument(id)
			.subscribe((res : Document) => {
				this.document = res;
			});
	}
}