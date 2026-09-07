import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all DataSet Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initAnalyticsWorkspaceList();
	this.initDataSourceList();
	this.initDataPipelineList();
	this.initSemanticModelList();
	this.initDimensionList();
	this.initMeasureList();
	this.initMetricList();
	this.initQualityRuleList();
	this.initLineageNodeList();
	this.initTagList();
  }
}
