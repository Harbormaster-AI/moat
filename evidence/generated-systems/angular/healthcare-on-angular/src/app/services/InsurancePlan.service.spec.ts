import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InsurancePlanService } from './InsurancePlan.service';

describe('InsurancePlanService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InsurancePlanService] });
	});

  it('should be created', () => {
    const service: InsurancePlanService = TestBed.get(InsurancePlanService);
    expect(service).toBeTruthy();
  });
});
