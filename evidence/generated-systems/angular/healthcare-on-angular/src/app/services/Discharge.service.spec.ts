import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DischargeService } from './Discharge.service';

describe('DischargeService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DischargeService] });
	});

  it('should be created', () => {
    const service: DischargeService = TestBed.get(DischargeService);
    expect(service).toBeTruthy();
  });
});
