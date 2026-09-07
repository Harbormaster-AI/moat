import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Customer Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initFinancialInstitutionList();
	this.initAccountList();
	this.initWalletList();
	this.initPaymentCardList();
	this.initKYCProfileList();
	this.initConsentList();
	this.initAgreementList();
	this.initLoanApplicationList();
	this.initLoanList();
	this.initInvestmentPortfolioList();
	this.initDisputeList();
  }
}
