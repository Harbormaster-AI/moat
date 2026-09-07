import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EmploymentContractService } from './EmploymentContract.service';

describe('EmploymentContractService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EmploymentContractService] });
	});

  it('should be created', () => {
    const service: EmploymentContractService = TestBed.get(EmploymentContractService);
    expect(service).toBeTruthy();
  });
});
