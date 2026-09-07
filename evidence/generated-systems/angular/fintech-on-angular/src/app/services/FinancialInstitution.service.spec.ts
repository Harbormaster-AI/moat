import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { FinancialInstitutionService } from './FinancialInstitution.service';

describe('FinancialInstitutionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [FinancialInstitutionService] });
	});

  it('should be created', () => {
    const service: FinancialInstitutionService = TestBed.get(FinancialInstitutionService);
    expect(service).toBeTruthy();
  });
});
