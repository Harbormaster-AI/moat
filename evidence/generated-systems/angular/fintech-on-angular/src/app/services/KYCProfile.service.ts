import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {KYCProfile} from '../models/KYCProfile';
import {CustomerService} from '../services/Customer.service';
import {KYCDocumentService} from '../services/KYCDocument.service';
import {ScreeningService} from '../services/Screening.service';
import {VerifiedAddressService} from '../services/VerifiedAddress.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class KYCProfileService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	kYCProfile : KYCProfile;

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
	// add a KYCProfile
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addKYCProfile(profileId, createdAt, Customer, Documents, Screenings, Addresses, Status, VerificationLevel) : Observable<any> {
		const uri_ = this.apiUrl + '/KYCProfile/create';
		const obj = {
			      		profileId: profileId,
      		createdAt: createdAt,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Documents: Documents != null && Documents.length > 0 ? Documents : null,
      		Screenings: Screenings != null && Screenings.length > 0 ? Screenings : null,
      		Addresses: Addresses != null && Addresses.length > 0 ? Addresses : null,
      		Status: Status,
			VerificationLevel: VerificationLevel
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a KYCProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateKYCProfile(profileId, createdAt, Customer, Documents, Screenings, Addresses, Status, VerificationLevel, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/KYCProfile/update/' + id;
		const obj = {
				      		profileId: profileId,
      		createdAt: createdAt,
      		Customer: Customer != null && Customer.length > 0 ? Customer : null,
      		Documents: Documents != null && Documents.length > 0 ? Documents : null,
      		Screenings: Screenings != null && Screenings.length > 0 ? Screenings : null,
      		Addresses: Addresses != null && Addresses.length > 0 ? Addresses : null,
      		Status: Status,
			VerificationLevel: VerificationLevel
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a KYCProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteKYCProfile(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/KYCProfile/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a KYCProfile
	// returns the results untouched as an Observable KYCProfile
	// KYCProfile model
	// delegates via URI
	//********************************************************************
	getKYCProfile(id) : Observable<KYCProfile> {
		const uri_ = this.apiUrl + '/KYCProfile/load/' + id;

		return this.http.get<KYCProfile>(uri_);
	}
	
	//********************************************************************
	// gets all KYCProfile
	// returns the results untouched as JSON representation of an
	// Observable array of KYCProfile models
	// delegates via URI
	//********************************************************************
	getKYCProfiles() : Observable<KYCProfile[]> {
		const uri_ = this.apiUrl + '/KYCProfile/';

		return this
			.http.get<KYCProfile[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Customer on a KYCProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCustomer( kYCProfileId, _customerId ): Observable<any> {

		// get the KYCProfile from storage
		this.loadHelper( kYCProfileId );

	// get the Customer from storage
	var tmp 	= new CustomerService(this.http).getCustomer(_customerId);

	// assign the Customer
	this.kYCProfile.customer = tmp;

	// save the KYCProfile
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Customer on a KYCProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCustomer( kYCProfileId ): Observable<any> {

		// get the KYCProfile from storage
		this.loadHelper( kYCProfileId );

	// assign Customer to null
	this.kYCProfile.customer = null;

	// save the KYCProfile
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more documentsIds as a Documents
	// to a KYCProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addDocuments( kYCProfileId, documentsIds ): Observable<any> {

		// get the KYCProfile
		this.loadHelper( kYCProfileId );

	// split on a comma with no spaces
	var idList = documentsIds.split(',')

	// iterate over array of documents ids
	idList.forEach(function (id) {
		// read the KYCDocument
		var kYCDocument = new KYCDocumentService(this.http).getKYCDocument(id);
		// add the KYCDocument if not already assigned
		if ( this.kYCProfile.documents.indexOf(kYCDocument) == -1 )
		this.kYCProfile.documents.push(kYCDocument);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more documentsIds as a Documents
	// from a KYCProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeDocuments( kYCProfileId, documentsIds ): Observable<any> {

		// get the KYCProfile
		this.loadHelper( kYCProfileId );


	// split on a comma with no spaces
	var idList 					= documentsIds.split(',');
	var documents 	= this.kYCProfile.documents;

	if ( documents != null && documentsIds != null ) {

		// iterate over array of documents ids
		documents.forEach(function (obj) {
			if ( documentsIds.indexOf(obj._id) > -1 ) {
				// remove the KYCDocument
				this.kYCProfile.documents.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more screeningsIds as a Screenings
	// to a KYCProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addScreenings( kYCProfileId, screeningsIds ): Observable<any> {

		// get the KYCProfile
		this.loadHelper( kYCProfileId );

	// split on a comma with no spaces
	var idList = screeningsIds.split(',')

	// iterate over array of screenings ids
	idList.forEach(function (id) {
		// read the Screening
		var screening = new ScreeningService(this.http).getScreening(id);
		// add the Screening if not already assigned
		if ( this.kYCProfile.screenings.indexOf(screening) == -1 )
		this.kYCProfile.screenings.push(screening);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more screeningsIds as a Screenings
	// from a KYCProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeScreenings( kYCProfileId, screeningsIds ): Observable<any> {

		// get the KYCProfile
		this.loadHelper( kYCProfileId );


	// split on a comma with no spaces
	var idList 					= screeningsIds.split(',');
	var screenings 	= this.kYCProfile.screenings;

	if ( screenings != null && screeningsIds != null ) {

		// iterate over array of screenings ids
		screenings.forEach(function (obj) {
			if ( screeningsIds.indexOf(obj._id) > -1 ) {
				// remove the Screening
				this.kYCProfile.screenings.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more addressesIds as a Addresses
	// to a KYCProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAddresses( kYCProfileId, addressesIds ): Observable<any> {

		// get the KYCProfile
		this.loadHelper( kYCProfileId );

	// split on a comma with no spaces
	var idList = addressesIds.split(',')

	// iterate over array of addresses ids
	idList.forEach(function (id) {
		// read the VerifiedAddress
		var verifiedAddress = new VerifiedAddressService(this.http).getVerifiedAddress(id);
		// add the VerifiedAddress if not already assigned
		if ( this.kYCProfile.addresses.indexOf(verifiedAddress) == -1 )
		this.kYCProfile.addresses.push(verifiedAddress);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more addressesIds as a Addresses
	// from a KYCProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAddresses( kYCProfileId, addressesIds ): Observable<any> {

		// get the KYCProfile
		this.loadHelper( kYCProfileId );


	// split on a comma with no spaces
	var idList 					= addressesIds.split(',');
	var addresses 	= this.kYCProfile.addresses;

	if ( addresses != null && addressesIds != null ) {

		// iterate over array of addresses ids
		addresses.forEach(function (obj) {
			if ( addressesIds.indexOf(obj._id) > -1 ) {
				// remove the VerifiedAddress
				this.kYCProfile.addresses.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a KYCProfile
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/KYCProfile/update/' + this.kYCProfile;

	return  this.http.post(uri_, this.kYCProfile );
}

	//********************************************************************
	// loadHelper - internal helper to load a KYCProfile
	//********************************************************************	
	loadHelper( id ) {
		this.getKYCProfile(id)
			.subscribe((res : KYCProfile) => {
				this.kYCProfile = res;
			});
	}
}