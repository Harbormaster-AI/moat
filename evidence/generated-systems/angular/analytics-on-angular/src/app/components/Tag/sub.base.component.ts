import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all Tag Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initDataSetList();
	this.initModel_List();
	this.initModelVersionList();
	this.initDashboardList();
	this.initReportList();
	this.initFeatureSetList();
	this.initMetricList();
  }
}
