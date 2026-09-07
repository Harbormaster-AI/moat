import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all User Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initOrganizationList();
	this.initTeamList();
	this.initActivityList();
	this.initAccountList();
	this.initLeadList();
	this.initOpportunityList();
	this.initCase_List();
	this.initQuoteList();
	this.initOrderList();
	this.initContractList();
	this.initEmailMessageList();
  }
}
