import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all EmailMessage Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initOrganizationList();
	this.initUserList();
	this.initAccountList();
	this.initContactList();
	this.initLeadList();
	this.initCase_List();
	this.initOpportunityList();
	this.initCampaignList();
  }
}
