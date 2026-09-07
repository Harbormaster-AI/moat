import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {CampaignMember} from '../models/CampaignMember';
import {CampaignService} from '../services/Campaign.service';
import {LeadService} from '../services/Lead.service';
import {ContactService} from '../services/Contact.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class CampaignMemberService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	campaignMember : CampaignMember;

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
	// add a CampaignMember
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addCampaignMember(responded, Campaign, Lead, Contact, Status, MemberType) : Observable<any> {
		const uri_ = this.apiUrl + '/CampaignMember/create';
		const obj = {
			      		responded: responded,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Lead: Lead != null && Lead.length > 0 ? Lead : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Status: Status,
			MemberType: MemberType
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a CampaignMember
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateCampaignMember(responded, Campaign, Lead, Contact, Status, MemberType, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/CampaignMember/update/' + id;
		const obj = {
				      		responded: responded,
      		Campaign: Campaign != null && Campaign.length > 0 ? Campaign : null,
      		Lead: Lead != null && Lead.length > 0 ? Lead : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Status: Status,
			MemberType: MemberType
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a CampaignMember
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteCampaignMember(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/CampaignMember/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a CampaignMember
	// returns the results untouched as an Observable CampaignMember
	// CampaignMember model
	// delegates via URI
	//********************************************************************
	getCampaignMember(id) : Observable<CampaignMember> {
		const uri_ = this.apiUrl + '/CampaignMember/load/' + id;

		return this.http.get<CampaignMember>(uri_);
	}
	
	//********************************************************************
	// gets all CampaignMember
	// returns the results untouched as JSON representation of an
	// Observable array of CampaignMember models
	// delegates via URI
	//********************************************************************
	getCampaignMembers() : Observable<CampaignMember[]> {
		const uri_ = this.apiUrl + '/CampaignMember/';

		return this
			.http.get<CampaignMember[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Campaign on a CampaignMember
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCampaign( campaignMemberId, _campaignId ): Observable<any> {

		// get the CampaignMember from storage
		this.loadHelper( campaignMemberId );

	// get the Campaign from storage
	var tmp 	= new CampaignService(this.http).getCampaign(_campaignId);

	// assign the Campaign
	this.campaignMember.campaign = tmp;

	// save the CampaignMember
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Campaign on a CampaignMember
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCampaign( campaignMemberId ): Observable<any> {

		// get the CampaignMember from storage
		this.loadHelper( campaignMemberId );

	// assign Campaign to null
	this.campaignMember.campaign = null;

	// save the CampaignMember
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lead on a CampaignMember
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLead( campaignMemberId, _leadId ): Observable<any> {

		// get the CampaignMember from storage
		this.loadHelper( campaignMemberId );

	// get the Lead from storage
	var tmp 	= new LeadService(this.http).getLead(_leadId);

	// assign the Lead
	this.campaignMember.lead = tmp;

	// save the CampaignMember
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lead on a CampaignMember
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLead( campaignMemberId ): Observable<any> {

		// get the CampaignMember from storage
		this.loadHelper( campaignMemberId );

	// assign Lead to null
	this.campaignMember.lead = null;

	// save the CampaignMember
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Contact on a CampaignMember
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignContact( campaignMemberId, _contactId ): Observable<any> {

		// get the CampaignMember from storage
		this.loadHelper( campaignMemberId );

	// get the Contact from storage
	var tmp 	= new ContactService(this.http).getContact(_contactId);

	// assign the Contact
	this.campaignMember.contact = tmp;

	// save the CampaignMember
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Contact on a CampaignMember
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignContact( campaignMemberId ): Observable<any> {

		// get the CampaignMember from storage
		this.loadHelper( campaignMemberId );

	// assign Contact to null
	this.campaignMember.contact = null;

	// save the CampaignMember
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a CampaignMember
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/CampaignMember/update/' + this.campaignMember;

	return  this.http.post(uri_, this.campaignMember );
}

	//********************************************************************
	// loadHelper - internal helper to load a CampaignMember
	//********************************************************************	
	loadHelper( id ) {
		this.getCampaignMember(id)
			.subscribe((res : CampaignMember) => {
				this.campaignMember = res;
			});
	}
}