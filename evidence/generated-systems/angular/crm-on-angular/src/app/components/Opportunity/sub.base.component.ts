import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Opportunity Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initOrganizationList();
	this.initAccountList();
	this.initUserList();
	this.initContactList();
	this.initOpportunityLineItemList();
	this.initOpportunityStageHistoryList();
	this.initQuoteList();
	this.initOrderList();
	this.initCampaignList();
	this.initActivityList();
	this.initTeamList();
  }
}
