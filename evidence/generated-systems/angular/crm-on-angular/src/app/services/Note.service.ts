import { Injectable } from '@angular/core';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import {Note} from '../models/Note';
import {OrganizationService} from '../services/Organization.service';
import {UserService} from '../services/User.service';
import {AccountService} from '../services/Account.service';
import {ContactService} from '../services/Contact.service';
import {OpportunityService} from '../services/Opportunity.service';
import {Case_Service} from '../services/Case_.service';
import {LeadService} from '../services/Lead.service';
import { HelperBaseService } from './helperbase.service';

@Injectable({
	providedIn: 'root'
})

export class NoteService extends HelperBaseService {

	//********************************************************************
	// general holder 
	//********************************************************************
	note : Note;

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
	// add a Note
	// returns the results untouched as a JSON representation
	// delegates via URI
	//********************************************************************
	addNote(title, content, createdAt, updatedAt, Organization, Owner, Account, Contact, Opportunity, Case, Lead) : Observable<any> {
		const uri_ = this.apiUrl + '/Note/create';
		const obj = {
			      		title: title,
      		content: content,
      		createdAt: createdAt,
      		updatedAt: updatedAt,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Case: Case != null && Case.length > 0 ? Case : null,
			Lead: Lead != null && Lead.length > 0 ? Lead : null
		};

		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// update a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
		updateNote(title, content, createdAt, updatedAt, Organization, Owner, Account, Contact, Opportunity, Case, Lead, id)  :  Observable<any>  {
			const uri_ = this.apiUrl + '/Note/update/' + id;
		const obj = {
				      		title: title,
      		content: content,
      		createdAt: createdAt,
      		updatedAt: updatedAt,
      		Organization: Organization != null && Organization.length > 0 ? Organization : null,
      		Owner: Owner != null && Owner.length > 0 ? Owner : null,
      		Account: Account != null && Account.length > 0 ? Account : null,
      		Contact: Contact != null && Contact.length > 0 ? Contact : null,
      		Opportunity: Opportunity != null && Opportunity.length > 0 ? Opportunity : null,
      		Case: Case != null && Case.length > 0 ? Case : null,
			Lead: Lead != null && Lead.length > 0 ? Lead : null
		};
		return this.http.post(uri_, obj);
	}

	//********************************************************************
	// delete a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	deleteNote(id)  : Observable<any> {
		const uri_ = this.apiUrl + '/Note/delete/' + id;

		return this.http.get(uri_);
	}
	
	//********************************************************************
	// loads a Note
	// returns the results untouched as an Observable Note
	// Note model
	// delegates via URI
	//********************************************************************
	getNote(id) : Observable<Note> {
		const uri_ = this.apiUrl + '/Note/load/' + id;

		return this.http.get<Note>(uri_);
	}
	
	//********************************************************************
	// gets all Note
	// returns the results untouched as JSON representation of an
	// Observable array of Note models
	// delegates via URI
	//********************************************************************
	getNotes() : Observable<Note[]> {
		const uri_ = this.apiUrl + '/Note/';

		return this
			.http.get<Note[]>(uri_);
	}
	
			//********************************************************************
	// assigns a Organization on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOrganization( noteId, _organizationId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// get the Organization from storage
	var tmp 	= new OrganizationService(this.http).getOrganization(_organizationId);

	// assign the Organization
	this.note.organization = tmp;

	// save the Note
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Organization on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOrganization( noteId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// assign Organization to null
	this.note.organization = null;

	// save the Note
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Owner on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOwner( noteId, _ownerId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// get the User from storage
	var tmp 	= new UserService(this.http).getUser(_ownerId);

	// assign the Owner
	this.note.owner = tmp;

	// save the Note
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Owner on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOwner( noteId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// assign Owner to null
	this.note.owner = null;

	// save the Note
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Account on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignAccount( noteId, _accountId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// get the Account from storage
	var tmp 	= new AccountService(this.http).getAccount(_accountId);

	// assign the Account
	this.note.account = tmp;

	// save the Note
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Account on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignAccount( noteId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// assign Account to null
	this.note.account = null;

	// save the Note
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Contact on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignContact( noteId, _contactId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// get the Contact from storage
	var tmp 	= new ContactService(this.http).getContact(_contactId);

	// assign the Contact
	this.note.contact = tmp;

	// save the Note
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Contact on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignContact( noteId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// assign Contact to null
	this.note.contact = null;

	// save the Note
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Opportunity on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignOpportunity( noteId, _opportunityId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// get the Opportunity from storage
	var tmp 	= new OpportunityService(this.http).getOpportunity(_opportunityId);

	// assign the Opportunity
	this.note.opportunity = tmp;

	// save the Note
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Opportunity on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignOpportunity( noteId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// assign Opportunity to null
	this.note.opportunity = null;

	// save the Note
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Case on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignCase( noteId, _caseId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// get the Case_ from storage
	var tmp 	= new Case_Service(this.http).getCase_(_caseId);

	// assign the Case
	this.note.case = tmp;

	// save the Note
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Case on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignCase( noteId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// assign Case to null
	this.note.case = null;

	// save the Note
	return this.saveHelper();
}

		//********************************************************************
	// assigns a Lead on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	assignLead( noteId, _leadId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// get the Lead from storage
	var tmp 	= new LeadService(this.http).getLead(_leadId);

	// assign the Lead
	this.note.lead = tmp;

	// save the Note
	return this.saveHelper();
}

	//********************************************************************
	// unassigns a Lead on a Note
	// returns an Observable
	// delegates via URI
	//********************************************************************
	unassignLead( noteId ): Observable<any> {

		// get the Note from storage
		this.loadHelper( noteId );

	// assign Lead to null
	this.note.lead = null;

	// save the Note
	return this.saveHelper();
}

	
	
	//********************************************************************
	// saveHelper - internal helper to save a Note
	//********************************************************************
	saveHelper() : Observable<any> {

		const uri_ = this.apiUrl + '/Note/update/' + this.note;

	return  this.http.post(uri_, this.note );
}

	//********************************************************************
	// loadHelper - internal helper to load a Note
	//********************************************************************	
	loadHelper( id ) {
		this.getNote(id)
			.subscribe((res : Note) => {
				this.note = res;
			});
	}
}