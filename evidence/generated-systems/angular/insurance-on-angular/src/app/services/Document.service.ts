import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Document} from '../models/Document';
import {PolicyService} from '../services/Policy.service';
import {ClaimService} from '../services/Claim.service';
import {ApplicationService} from '../services/Application.service';
import {CustomerService} from '../services/Customer.service';
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
	addDocument(fileName, uploadedDate, Policy, Claim, Application, Customer, DocumentType) : Observable<any> {
		const uri_ = this.apiUrl + '/Document/create';
		const obj = {
			      		fileName: fileName,
      		uploadedDate: uploadedDate,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		Application: Application != null && Application.length > 0 ? Application : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
			DocumentType: DocumentType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateDocument(fileName, uploadedDate, Policy, Claim, Application, Customer, DocumentType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Document/update/' + id;
		const obj = {
				      		fileName: fileName,
      		uploadedDate: uploadedDate,
      		Policy: Policy != null && Policy.length > 0 ? Policy : null,
      		Claim: Claim != null && Claim.length > 0 ? Claim : null,
      		Application: Application != null && Application.length > 0 ? Application : null,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
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
	// assigns a Policy on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPolicy( documentId, _policyId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// get the Policy from storage
	var tmp 	= new PolicyService(this.http).getPolicy(_policyId);

	// assign the Policy
	this.document.policy = tmp;

	// save the Document
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Policy on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPolicy( documentId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// assign Policy to null
	this.document.policy = null;

	// save the Document
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Claim on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignClaim( documentId, _claimId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// get the Claim from storage
	var tmp 	= new ClaimService(this.http).getClaim(_claimId);

	// assign the Claim
	this.document.claim = tmp;

	// save the Document
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Claim on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignClaim( documentId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// assign Claim to null
	this.document.claim = null;

	// save the Document
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Application on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignApplication( documentId, _applicationId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// get the Application from storage
	var tmp 	= new ApplicationService(this.http).getApplication(_applicationId);

	// assign the Application
	this.document.application = tmp;

	// save the Document
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Application on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignApplication( documentId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// assign Application to null
	this.document.application = null;

	// save the Document
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Customer on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( documentId, _customerId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.document.customer = tmp;

	// save the Document
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a Document
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( documentId ): Observable<any> {

		// get the Document from storage
		this.loadHelper( documentId );

	// assign Customer to null
	this.document.customer = null;

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