import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Alert Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initMetricList();
	this.initDashboardList();
	this.initDataSetList();
	this.initQualityRuleList();
	this.initAnomalyList();
	this.initSubscriberList();
  }
}
