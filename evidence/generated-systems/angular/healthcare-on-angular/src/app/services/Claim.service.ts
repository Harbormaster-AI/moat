import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Claim} from '../models/Claim';
import {PatientService} from '../services/Patient.service';
import {CoverageService} from '../services/Coverage.service';
import {EncounterService} from '../services/Encounter.service';
import {InvoiceService} from '../services/Invoice.service';
import {InsurancePayerService} from '../services/InsurancePayer.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class ClaimService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	claim : Claim;

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
	// add a Claim
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addClaim(claimNumber, totalAmount, Patient, Coverage, Encounter, Invoices, Payer, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Claim/create';
		const obj = {
			      		claimNumber: claimNumber,
      		totalAmount: totalAmount,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Coverage: Coverage != null && Coverage.length > 0 ? Coverage : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Invoices: Invoices != null && Invoices.length > 0 ? Invoices : null,
      		Payer: Payer != null && Payer.length > 0 ? Payer : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateClaim(claimNumber, totalAmount, Patient, Coverage, Encounter, Invoices, Payer, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Claim/update/' + id;
		const obj = {
				      		claimNumber: claimNumber,
      		totalAmount: totalAmount,
      		Patient: Patient != null && Patient.length > 0 ? Patient : null,
      		Coverage: Coverage != null && Coverage.length > 0 ? Coverage : null,
      		Encounter: Encounter != null && Encounter.length > 0 ? Encounter : null,
      		Invoices: Invoices != null && Invoices.length > 0 ? Invoices : null,
      		Payer: Payer != null && Payer.length > 0 ? Payer : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteClaim(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Claim/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Claim
	// returns the results untouched as an Observable Claim
	// Claim model
	// delegates via URI
	//********************************************************************
	getClaim(id) : Observable<Claim> {
		const uri_ = this.apiUrl + '/Claim/load/' + id;

		return this.http.get<Claim>(uri_);
	}
	
	//********************************************************************
	// gets all Claim
	// returns the results untouched as JSON representation of an
	// Observable array of Claim models
	// delegates via URI
	//********************************************************************
	getClaims() : Observable<Claim[]> {
		const uri_ = this.apiUrl + '/Claim/';

		return this
			.http.get<Claim[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Patient on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPatient( claimId, _patientId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// get the Patient from storage
	var tmp 	= new PatientService(this.http).getPatient(_patientId);

	// assign the Patient
	this.claim.patient = tmp;

	// save the Claim
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Patient on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPatient( claimId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// assign Patient to null
	this.claim.patient = null;

	// save the Claim
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Coverage on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCoverage( claimId, _coverageId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// get the Coverage from storage
	var tmp 	= new CoverageService(this.http).getCoverage(_coverageId);

	// assign the Coverage
	this.claim.coverage = tmp;

	// save the Claim
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Coverage on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCoverage( claimId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// assign Coverage to null
	this.claim.coverage = null;

	// save the Claim
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Encounter on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignEncounter( claimId, _encounterId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// get the Encounter from storage
	var tmp 	= new EncounterService(this.http).getEncounter(_encounterId);

	// assign the Encounter
	this.claim.encounter = tmp;

	// save the Claim
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Encounter on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignEncounter( claimId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// assign Encounter to null
	this.claim.encounter = null;

	// save the Claim
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Payer on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignPayer( claimId, _payerId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// get the InsurancePayer from storage
	var tmp 	= new InsurancePayerService(this.http).getInsurancePayer(_payerId);

	// assign the Payer
	this.claim.payer = tmp;

	// save the Claim
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Payer on a Claim
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignPayer( claimId ): Observable<any> {

		// get the Claim from storage
		this.loadHelper( claimId );

	// assign Payer to null
	this.claim.payer = null;

	// save the Claim
	return this.saveHelper();
}

	
		//********************************************************************
	// adds one or more invoicesIds as a Invoices
	// to a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	addInvoices( claimId, invoicesIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );

	// split on a comma with no spaces
	var idList = invoicesIds.split(',')

	// iterate over array of invoices ids
	idList.forEach(function (id) {
		// read the Invoice
		var invoice = new InvoiceService(this.http).getInvoice(id);
		// add the Invoice if not already assigned
		if ( this.claim.invoices.indexOf(invoice) == -1 )
		this.claim.invoices.push(invoice);
	});

	// save it
	return this.saveHelper();
}

	//********************************************************************
	// removes one or more invoicesIds as a Invoices
	// from a Claim
	// returns a Promise
	// delegates via URI
	//********************************************************************
	removeInvoices( claimId, invoicesIds ): Observable<any> {

		// get the Claim
		this.loadHelper( claimId );


	// split on a comma with no spaces
	var idList 					= invoicesIds.split(',');
	var invoices 	= this.claim.invoices;

	if ( invoices != null && invoicesIds != null ) {

		// iterate over array of invoices ids
		invoices.forEach(function (obj) {
			if ( invoicesIds.indexOf(obj._id) > -1 ) {
				// remove the Invoice
				this.claim.invoices.pop(obj);
			}
		});

		// save it
		return this.saveHelper();
	}
}

	
	//********************************************************************
	// saveHelper - internal helper to save a Claim
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Claim/update/' + this.claim;

	return  this.http.post(uri_, this.claim );
}

	//********************************************************************
	// loadHelper - internal helper to load a Claim
	//********************************************************************	
	loadHelper( id ) {
		this.getClaim(id)
			.subscribe((res : Claim) => {
				this.claim = res;
			});
	}
}