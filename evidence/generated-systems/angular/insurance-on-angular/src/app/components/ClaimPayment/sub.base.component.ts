import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all ClaimPayment Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initClaimList();
	this.initExposureList();
	this.initBeneficiaryList();
	this.initServiceProviderList();
	this.initCustomerList();
  }
}
