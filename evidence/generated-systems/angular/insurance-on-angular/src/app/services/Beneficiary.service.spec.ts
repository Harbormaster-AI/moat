import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BeneficiaryService } from './Beneficiary.service';

describe('BeneficiaryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BeneficiaryService] });
	});

  it('should be created', () => {
    const service: BeneficiaryService = TestBed.get(BeneficiaryService);
    expect(service).toBeTruthy();
  });
});
