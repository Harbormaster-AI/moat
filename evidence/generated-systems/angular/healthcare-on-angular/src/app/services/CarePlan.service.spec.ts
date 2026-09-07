import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CarePlanService } from './CarePlan.service';

describe('CarePlanService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CarePlanService] });
	});

  it('should be created', () => {
    const service: CarePlanService = TestBed.get(CarePlanService);
    expect(service).toBeTruthy();
  });
});
