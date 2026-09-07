import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MaintenancePlanService } from './MaintenancePlan.service';

describe('MaintenancePlanService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MaintenancePlanService] });
	});

  it('should be created', () => {
    const service: MaintenancePlanService = TestBed.get(MaintenancePlanService);
    expect(service).toBeTruthy();
  });
});
