import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {OpportunityStageHistory} from '../models/OpportunityStageHistory';
import {OpportunityService} from '../services/Opportunity.service';
import {UserService} from '../services/User.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class OpportunityStageHistoryService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	opportunityStageHistory : OpportunityStageHistory;

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
	// add a OpportunityStageHistory
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addOpportunityStageHistory(changedAt, comment, Opportunity, ChangedBy, FromStage, ToStage) : Observable<any> {
		const uri_ = this.apiUrl + '/OpportunityStageHistory/create';
		const obj = {
			      		changedAt: changedAt,
      		comment: comment,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		ChangedBy: ChangedBy != null && ChangedBy.length > 0 ? ChangedBy : null,
      		FromStage: FromStage,
			ToStage: ToStage
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a OpportunityStageHistory
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateOpportunityStageHistory(changedAt, comment, Opportunity, ChangedBy, FromStage, ToStage, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/OpportunityStageHistory/update/' + id;
		const obj = {
				      		changedAt: changedAt,
      		comment: comment,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		ChangedBy: ChangedBy != null && ChangedBy.length > 0 ? ChangedBy : null,
      		FromStage: FromStage,
			ToStage: ToStage
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a OpportunityStageHistory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteOpportunityStageHistory(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/OpportunityStageHistory/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a OpportunityStageHistory
	// returns the results untouched as an Observable OpportunityStageHistory
	// OpportunityStageHistory model
	// delegates via URI
	//********************************************************************
	getOpportunityStageHistory(id) : Observable<OpportunityStageHistory> {
		const uri_ = this.apiUrl + '/OpportunityStageHistory/load/' + id;

		return this.http.get<OpportunityStageHistory>(uri_);
	}
	
	//********************************************************************
	// gets all OpportunityStageHistory
	// returns the results untouched as JSON representation of an
	// Observable array of OpportunityStageHistory models
	// delegates via URI
	//********************************************************************
	getOpportunityStageHistorys() : Observable<OpportunityStageHistory[]> {
		const uri_ = this.apiUrl + '/OpportunityStageHistory/';

		return this
			.http.get<OpportunityStageHistory[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Opportunity on a OpportunityStageHistory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOpportunity( opportunityStageHistoryId, _opportunityId ): Observable<any> {

		// get the OpportunityStageHistory from storage
		this.loadHelper( opportunityStageHistoryId );

	// get the Opportunity from storage
	var tmp 	= new OpportunityService(this.http).getOpportunity(_opportunityId);

	// assign the Opportunity
	this.opportunityStageHistory.opportunity = tmp;

	// save the OpportunityStageHistory
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Opportunity on a OpportunityStageHistory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOpportunity( opportunityStageHistoryId ): Observable<any> {

		// get the OpportunityStageHistory from storage
		this.loadHelper( opportunityStageHistoryId );

	// assign Opportunity to null
	this.opportunityStageHistory.opportunity = null;

	// save the OpportunityStageHistory
	return this.saveHelper();
}

		//********************************************************************
	// assigns a ChangedBy on a OpportunityStageHistory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignChangedBy( opportunityStageHistoryId, _changedById ): Observable<any> {

		// get the OpportunityStageHistory from storage
		this.loadHelper( opportunityStageHistoryId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_changedById);

	// assign the ChangedBy
	this.opportunityStageHistory.changedBy = tmp;

	// save the OpportunityStageHistory
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a ChangedBy on a OpportunityStageHistory
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignChangedBy( opportunityStageHistoryId ): Observable<any> {

		// get the OpportunityStageHistory from storage
		this.loadHelper( opportunityStageHistoryId );

	// assign ChangedBy to null
	this.opportunityStageHistory.changedBy = null;

	// save the OpportunityStageHistory
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a OpportunityStageHistory
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/OpportunityStageHistory/update/' + this.opportunityStageHistory;

	return  this.http.post(uri_, this.opportunityStageHistory );
}

	//********************************************************************
	// loadHelper - internal helper to load a OpportunityStageHistory
	//********************************************************************	
	loadHelper( id ) {
		this.getOpportunityStageHistory(id)
			.subscribe((res : OpportunityStageHistory) => {
				this.opportunityStageHistory = res;
			});
	}
}