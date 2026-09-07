import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Case_ Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initOrganizationList();
	this.initAccountList();
	this.initContactList();
	this.initUserList();
	this.initTeamList();
	this.initActivityList();
	this.initNoteList();
	this.initEmailMessageList();
	this.initOpportunityList();
  }
}
