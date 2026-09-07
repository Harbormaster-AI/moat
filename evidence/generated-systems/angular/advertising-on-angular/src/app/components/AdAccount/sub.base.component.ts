import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all AdAccount Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initAdvertiserList();
	this.initUserList();
	this.initCampaignList();
	this.initBillingProfileList();
	this.initDSPList();
	this.initPerformanceMetricList();
  }
}
