import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InsurancePayerService } from './InsurancePayer.service';

describe('InsurancePayerService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InsurancePayerService] });
	});

  it('should be created', () => {
    const service: InsurancePayerService = TestBed.get(InsurancePayerService);
    expect(service).toBeTruthy();
  });
});
