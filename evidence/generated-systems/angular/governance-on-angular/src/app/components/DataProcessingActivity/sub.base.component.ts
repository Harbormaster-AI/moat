import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all DataProcessingActivity Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initOrganizationList();
	this.initDataCategoryList();
	this.initSystem_List();
	this.initRecord_List();
	this.initPrivacyNoticeList();
	this.initThirdPartyList();
	this.initConsentList();
	this.initDataBreachList();
	this.initDataSubjectRequestList();
  }
}
