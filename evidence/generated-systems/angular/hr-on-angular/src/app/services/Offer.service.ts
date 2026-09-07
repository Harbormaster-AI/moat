import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Offer} from '../models/Offer';
import {JobRequisitionService} from '../services/JobRequisition.service';
import {CandidateService} from '../services/Candidate.service';
import {EmployeeService} from '../services/Employee.service';
import {EmploymentContractService} from '../services/EmploymentContract.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OfferService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	offer : Offer;

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
	// add a Offer
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOffer(offerNumber, proposedStartDate, baseSalary, signOnBonus, Requisition, Candidate, ApprovedBy, Contract, Status) : Observable<any> {
		const uri_ = this.apiUrl + '/Offer/create';
		const obj = {
			      		offerNumber: offerNumber,
      		proposedStartDate: proposedStartDate,
      		baseSalary: baseSalary,
      		signOnBonus: signOnBonus,
      		Requisition: Requisition != null && Requisition.length > 0 ? Requisition : null,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		ApprovedBy: ApprovedBy != null && ApprovedBy.length > 0 ? ApprovedBy : null,
      		Contract: Contract != null && Contract.length > 0 ? Contract : null,
			Status: Status
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOffer(offerNumber, proposedStartDate, baseSalary, signOnBonus, Requisition, Candidate, ApprovedBy, Contract, Status, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Offer/update/' + id;
		const obj = {
				      		offerNumber: offerNumber,
      		proposedStartDate: proposedStartDate,
      		baseSalary: baseSalary,
      		signOnBonus: signOnBonus,
      		Requisition: Requisition != null && Requisition.length > 0 ? Requisition : null,
      		Candidate: Candidate != null && Candidate.length > 0 ? Candidate : null,
      		ApprovedBy: ApprovedBy != null && ApprovedBy.length > 0 ? ApprovedBy : null,
      		Contract: Contract != null && Contract.length > 0 ? Contract : null,
			Status: Status
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOffer(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Offer/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Offer
	// returns the results untouched as an Observable Offer
	// Offer model
	// delegates via URI
	//********************************************************************
	getOffer(id) : Observable<Offer> {
		const uri_ = this.apiUrl + '/Offer/load/' + id;

		return this.http.get<Offer>(uri_);
	}
	
	//********************************************************************
	// gets all Offer
	// returns the results untouched as JSON representation of an
	// Observable array of Offer models
	// delegates via URI
	//********************************************************************
	getOffers() : Observable<Offer[]> {
		const uri_ = this.apiUrl + '/Offer/';

		return this
			.http.get<Offer[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Requisition on a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignRequisition( offerId, _requisitionId ): Observable<any> {

		// get the Offer from storage
		this.loadHelper( offerId );

	// get the JobRequisition from storage
	var tmp 	= new JobRequisitionService(this.http).getJobRequisition(_requisitionId);

	// assign the Requisition
	this.offer.requisition = tmp;

	// save the Offer
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Requisition on a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignRequisition( offerId ): Observable<any> {

		// get the Offer from storage
		this.loadHelper( offerId );

	// assign Requisition to null
	this.offer.requisition = null;

	// save the Offer
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Candidate on a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCandidate( offerId, _candidateId ): Observable<any> {

		// get the Offer from storage
		this.loadHelper( offerId );

	// get the Candidate from storage
	var tmp 	= new CandidateService(this.http).getCandidate(_candidateId);

	// assign the Candidate
	this.offer.candidate = tmp;

	// save the Offer
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Candidate on a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCandidate( offerId ): Observable<any> {

		// get the Offer from storage
		this.loadHelper( offerId );

	// assign Candidate to null
	this.offer.candidate = null;

	// save the Offer
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ApprovedBy on a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignApprovedBy( offerId, _approvedById ): Observable<any> {

		// get the Offer from storage
		this.loadHelper( offerId );

	// get the Employee from storage
	var tmp 	= new EmployeeService(this.http).getEmployee(_approvedById);

	// assign the ApprovedBy
	this.offer.approvedBy = tmp;

	// save the Offer
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ApprovedBy on a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignApprovedBy( offerId ): Observable<any> {

		// get the Offer from storage
		this.loadHelper( offerId );

	// assign ApprovedBy to null
	this.offer.approvedBy = null;

	// save the Offer
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Contract on a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignContract( offerId, _contractId ): Observable<any> {

		// get the Offer from storage
		this.loadHelper( offerId );

	// get the EmploymentContract from storage
	var tmp 	= new EmploymentContractService(this.http).getEmploymentContract(_contractId);

	// assign the Contract
	this.offer.contract = tmp;

	// save the Offer
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Contract on a Offer
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignContract( offerId ): Observable<any> {

		// get the Offer from storage
		this.loadHelper( offerId );

	// assign Contract to null
	this.offer.contract = null;

	// save the Offer
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Offer
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Offer/update/' + this.offer;

	return  this.http.post(uri_, this.offer );
}

	//********************************************************************
	// loadHelper - internal helper to load a Offer
	//********************************************************************	
	loadHelper( id ) {
		this.getOffer(id)
			.subscribe((res : Offer) => {
				this.offer = res;
			});
	}
}