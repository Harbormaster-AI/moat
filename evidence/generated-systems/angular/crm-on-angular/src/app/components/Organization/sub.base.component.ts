import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Organization Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initUserList();
	this.initAccountList();
	this.initTeamList();
	this.initTerritoryList();
	this.initProductList();
	this.initPriceBookList();
	this.initCampaignList();
  }
}
