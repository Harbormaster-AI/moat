import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Campaign Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initOrganizationList();
	this.initCampaignList();
	this.initCampaignList();
	this.initCampaignMemberList();
	this.initOpportunityList();
	this.initAccountList();
	this.initLeadList();
	this.initContactList();
	this.initTeamList();
	this.initActivityList();
  }
}
