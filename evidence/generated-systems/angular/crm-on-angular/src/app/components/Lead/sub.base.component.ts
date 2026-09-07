import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Lead Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initOrganizationList();
	this.initUserList();
	this.initActivityList();
	this.initCampaignList();
	this.initAccountList();
	this.initContactList();
	this.initOpportunityList();
	this.initNoteList();
	this.initEmailMessageList();
  }
}
