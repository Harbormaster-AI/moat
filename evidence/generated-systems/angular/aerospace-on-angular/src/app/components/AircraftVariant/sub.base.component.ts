import { HttpClient } from '@angular/common/http';
import { BaseComponent } from '../base.component';

import { Directive } from '@angular/core';

/**
	Base class of all AircraftVariant Edit and Create Components.  
 **/
@Directive()
export class SubBaseComponent extends BaseComponent {

  constructor (http: HttpClient) { super(http); }
  
  ngOnInit() {
  	super.ngOnInit();
  	
	this.initAircraftModelList();
	this.initEngineTypeList();
	this.initAvionicsSuiteList();
	this.initAPUList();
	this.initLandingGearList();
	this.initCabinLayoutList();
	this.initAircraftOptionList();
	this.initAircraftPackageList();
  }
}
