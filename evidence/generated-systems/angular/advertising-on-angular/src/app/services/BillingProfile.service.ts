import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {BillingProfile} from '../models/BillingProfile';
import {AdvertiserService} from '../services/Advertiser.service';
import {PaymentMethodService} from '../services/PaymentMethod.service';
import {AdAccountService} from '../services/AdAccount.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class BillingProfileService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	billingProfile : BillingProfile;

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
	// add a BillingProfile
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addBillingProfile(billingName, taxId, billingAddress, Advertiser, PaymentMethods, AdAccounts, PaymentTerms) : Observable<any> {
		const uri_ = this.apiUrl + '/BillingProfile/create';
		const obj = {
			      		billingName: billingName,
      		taxId: taxId,
      		billingAddress: billingAddress,
      		Advertiser: Advertiser != null && Advertiser.length > 0 ? Advertiser : null,
      		PaymentMethods: PaymentMethods != null && PaymentMethods.length > 0 ? PaymentMethods : null,
      		AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null,
			PaymentTerms: PaymentTerms
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a BillingProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateBillingProfile(billingName, taxId, billingAddress, Advertiser, PaymentMethods, AdAccounts, PaymentTerms, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/BillingProfile/update/' + id;
		const obj = {
				      		billingName: billingName,
      		taxId: taxId,
      		billingAddress: billingAddress,
      		Advertiser: Advertiser != null && Advertiser.length > 0 ? Advertiser : null,
      		PaymentMethods: PaymentMethods != null && PaymentMethods.length > 0 ? PaymentMethods : null,
      		AdAccounts: AdAccounts != null && AdAccounts.length > 0 ? AdAccounts : null,
			PaymentTerms: PaymentTerms
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a BillingProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteBillingProfile(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/BillingProfile/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a BillingProfile
	// returns the results untouched as an Observable BillingProfile
	// BillingProfile model
	// delegates via URI
	//********************************************************************
	getBillingProfile(id) : Observable<BillingProfile> {
		const uri_ = this.apiUrl + '/BillingProfile/load/' + id;

		return this.http.get<BillingProfile>(uri_);
	}
	
	//********************************************************************
	// gets all BillingProfile
	// returns the results untouched as JSON representation of an
	// Observable array of BillingProfile models
	// delegates via URI
	//********************************************************************
	getBillingProfiles() : Observable<BillingProfile[]> {
		const uri_ = this.apiUrl + '/BillingProfile/';

		return this
			.http.get<BillingProfile[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Advertiser on a BillingProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAdvertiser( billingProfileId, _advertiserId ): Observable<any> {

		// get the BillingProfile from storage
		this.loadHelper( billingProfileId );

	// get the Advertiser from storage
	var tmp 	= new AdvertiserService(this.http).getAdvertiser(_advertiserId);

	// assign the Advertiser
	this.billingProfile.advertiser = tmp;

	// save the BillingProfile
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Advertiser on a BillingProfile
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAdvertiser( billingProfileId ): Observable<any> {

		// get the BillingProfile from storage
		this.loadHelper( billingProfileId );

	// assign Advertiser to null
	this.billingProfile.advertiser = null;

	// save the BillingProfile
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more paymentMethodsIds as a PaymentMethods
	// to a BillingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addPaymentMethods( billingProfileId, paymentMethodsIds ): Observable<any> {

		// get the BillingProfile
		this.loadHelper( billingProfileId );

	// split on a comma with no spaces
	var idList = paymentMethodsIds.split(',')

	// iterate over array of paymentMethods ids
	idList.forEach(function (id) {
		// read the PaymentMethod
		var paymentMethod = new PaymentMethodService(this.http).getPaymentMethod(id);
		// add the PaymentMethod if not already assigned
		if ( this.billingProfile.paymentMethods.indexOf(paymentMethod) == -1 )
		this.billingProfile.paymentMethods.push(paymentMethod);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more paymentMethodsIds as a PaymentMethods
	// from a BillingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removePaymentMethods( billingProfileId, paymentMethodsIds ): Observable<any> {

		// get the BillingProfile
		this.loadHelper( billingProfileId );


	// split on a comma with no spaces
	var idList 					= paymentMethodsIds.split(',');
	var paymentMethods 	= this.billingProfile.paymentMethods;

	if ( paymentMethods != null && paymentMethodsIds != null ) {

		// iterate over array of paymentMethods ids
		paymentMethods.forEach(function (obj) {
			if ( paymentMethodsIds.indexOf(obj._id) > -1 ) {
				// remove the PaymentMethod
				this.billingProfile.paymentMethods.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

		//********************************************************************
	// adds one or more adAccountsIds as a AdAccounts
	// to a BillingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addAdAccounts( billingProfileId, adAccountsIds ): Observable<any> {

		// get the BillingProfile
		this.loadHelper( billingProfileId );

	// split on a comma with no spaces
	var idList = adAccountsIds.split(',')

	// iterate over array of adAccounts ids
	idList.forEach(function (id) {
		// read the AdAccount
		var adAccount = new AdAccountService(this.http).getAdAccount(id);
		// add the AdAccount if not already assigned
		if ( this.billingProfile.adAccounts.indexOf(adAccount) == -1 )
		this.billingProfile.adAccounts.push(adAccount);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more adAccountsIds as a AdAccounts
	// from a BillingProfile
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeAdAccounts( billingProfileId, adAccountsIds ): Observable<any> {

		// get the BillingProfile
		this.loadHelper( billingProfileId );


	// split on a comma with no spaces
	var idList 					= adAccountsIds.split(',');
	var adAccounts 	= this.billingProfile.adAccounts;

	if ( adAccounts != null && adAccountsIds != null ) {

		// iterate over array of adAccounts ids
		adAccounts.forEach(function (obj) {
			if ( adAccountsIds.indexOf(obj._id) > -1 ) {
				// remove the AdAccount
				this.billingProfile.adAccounts.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a BillingProfile
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/BillingProfile/update/' + this.billingProfile;

	return  this.http.post(uri_, this.billingProfile );
}

	//********************************************************************
	// loadHelper - internal helper to load a BillingProfile
	//********************************************************************	
	loadHelper( id ) {
		this.getBillingProfile(id)
			.subscribe((res : BillingProfile) => {
				this.billingProfile = res;
			});
	}
}