import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {KYCDocument} from '../models/KYCDocument';
import {KYCProfileService} from '../services/KYCProfile.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class KYCDocumentService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	kYCDocument : KYCDocument;

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
	// add a KYCDocument
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addKYCDocument(reference, issuedCountry, expirationDate, KycProfile, DocumentType, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/KYCDocument/create';
		const obj = {
			      		reference: reference,
      		issuedCountry: issuedCountry,
      		expirationDate: expirationDate,
      		KycProfile: KycProfile != null && KycProfile.length > 0 ? KycProfile : null,
      		DocumentType: DocumentType,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a KYCDocument
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateKYCDocument(reference, issuedCountry, expirationDate, KycProfile, DocumentType, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/KYCDocument/update/' + id;
		const obj = {
				      		reference: reference,
      		issuedCountry: issuedCountry,
      		expirationDate: expirationDate,
      		KycProfile: KycProfile != null && KycProfile.length > 0 ? KycProfile : null,
      		DocumentType: DocumentType,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a KYCDocument
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteKYCDocument(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/KYCDocument/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a KYCDocument
	// returns the results untouched as an Observable KYCDocument
	// KYCDocument model
	// delegates via URI
	//********************************************************************
	getKYCDocument(id) : Observable<KYCDocument> {
		const uri_ = this.apiUrl + '/KYCDocument/load/' + id;

		return this.http.get<KYCDocument>(uri_);
	}
	
	//********************************************************************
	// gets all KYCDocument
	// returns the results untouched as JSON representation of an
	// Observable array of KYCDocument models
	// delegates via URI
	//********************************************************************
	getKYCDocuments() : Observable<KYCDocument[]> {
		const uri_ = this.apiUrl + '/KYCDocument/';

		return this
			.http.get<KYCDocument[]>(uri_);
	}
	
			//********************************************************************
	// assigns a KycProfile on a KYCDocument
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignKycProfile( kYCDocumentId, _kycProfileId ): Observable<any> {

		// get the KYCDocument from storage
		this.loadHelper( kYCDocumentId );

	// get the KYCProfile from storage
	var tmp 	= new KYCProfileService(this.http).getKYCProfile(_kycProfileId);

	// assign the KycProfile
	this.kYCDocument.kycProfile = tmp;

	// save the KYCDocument
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a KycProfile on a KYCDocument
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignKycProfile( kYCDocumentId ): Observable<any> {

		// get the KYCDocument from storage
		this.loadHelper( kYCDocumentId );

	// assign KycProfile to null
	this.kYCDocument.kycProfile = null;

	// save the KYCDocument
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a KYCDocument
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/KYCDocument/update/' + this.kYCDocument;

	return  this.http.post(uri_, this.kYCDocument );
}

	//********************************************************************
	// loadHelper - internal helper to load a KYCDocument
	//********************************************************************	
	loadHelper( id ) {
		this.getKYCDocument(id)
			.subscribe((res : KYCDocument) => {
				this.kYCDocument = res;
			});
	}
}