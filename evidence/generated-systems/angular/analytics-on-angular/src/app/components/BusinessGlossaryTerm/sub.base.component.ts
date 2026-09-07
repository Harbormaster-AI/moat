import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all BusinessGlossaryTerm Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initBusinessGlossaryTermList();
	this.initMetricList();
	this.initDataSetList();
	this.initDimensionList();
	this.initMeasureList();
  }
}
