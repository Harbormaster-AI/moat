import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Claim Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initPolicyList();
	this.initCustomerList();
	this.initAdjusterList();
	this.initIncidentList();
	this.initExposureList();
	this.initClaimReserveList();
	this.initClaimPaymentList();
	this.initServiceProviderList();
	this.initSubrogationRecoveryList();
  }
}
