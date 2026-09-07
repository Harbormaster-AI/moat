import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {VerifiedAddress} from '../models/VerifiedAddress';
import {KYCProfileService} from '../services/KYCProfile.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class VerifiedAddressService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	verifiedAddress : VerifiedAddress;

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
	// add a VerifiedAddress
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addVerifiedAddress(address, verifiedAt, KycProfile, VerificationStatus) : Observable<any> {
		const uri_ = this.apiUrl + '/VerifiedAddress/create';
		const obj = {
			      		address: address,
      		verifiedAt: verifiedAt,
      		KycProfile: KycProfile != null && KycProfile.length > 0 ? KycProfile : null,
			VerificationStatus: VerificationStatus
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a VerifiedAddress
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateVerifiedAddress(address, verifiedAt, KycProfile, VerificationStatus, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/VerifiedAddress/update/' + id;
		const obj = {
				      		address: address,
      		verifiedAt: verifiedAt,
      		KycProfile: KycProfile != null && KycProfile.length > 0 ? KycProfile : null,
			VerificationStatus: VerificationStatus
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a VerifiedAddress
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteVerifiedAddress(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/VerifiedAddress/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a VerifiedAddress
	// returns the results untouched as an Observable VerifiedAddress
	// VerifiedAddress model
	// delegates via URI
	//********************************************************************
	getVerifiedAddress(id) : Observable<VerifiedAddress> {
		const uri_ = this.apiUrl + '/VerifiedAddress/load/' + id;

		return this.http.get<VerifiedAddress>(uri_);
	}
	
	//********************************************************************
	// gets all VerifiedAddress
	// returns the results untouched as JSON representation of an
	// Observable array of VerifiedAddress models
	// delegates via URI
	//********************************************************************
	getVerifiedAddresss() : Observable<VerifiedAddress[]> {
		const uri_ = this.apiUrl + '/VerifiedAddress/';

		return this
			.http.get<VerifiedAddress[]>(uri_);
	}
	
			//********************************************************************
	// assigns a KycProfile on a VerifiedAddress
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignKycProfile( verifiedAddressId, _kycProfileId ): Observable<any> {

		// get the VerifiedAddress from storage
		this.loadHelper( verifiedAddressId );

	// get the KYCProfile from storage
	var tmp 	= new KYCProfileService(this.http).getKYCProfile(_kycProfileId);

	// assign the KycProfile
	this.verifiedAddress.kycProfile = tmp;

	// save the VerifiedAddress
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a KycProfile on a VerifiedAddress
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignKycProfile( verifiedAddressId ): Observable<any> {

		// get the VerifiedAddress from storage
		this.loadHelper( verifiedAddressId );

	// assign KycProfile to null
	this.verifiedAddress.kycProfile = null;

	// save the VerifiedAddress
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a VerifiedAddress
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/VerifiedAddress/update/' + this.verifiedAddress;

	return  this.http.post(uri_, this.verifiedAddress );
}

	//********************************************************************
	// loadHelper - internal helper to load a VerifiedAddress
	//********************************************************************	
	loadHelper( id ) {
		this.getVerifiedAddress(id)
			.subscribe((res : VerifiedAddress) => {
				this.verifiedAddress = res;
			});
	}
}