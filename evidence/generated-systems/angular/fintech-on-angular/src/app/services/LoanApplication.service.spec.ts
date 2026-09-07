import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LoanApplicationService } from './LoanApplication.service';

describe('LoanApplicationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LoanApplicationService] });
	});

  it('should be created', () => {
    const service: LoanApplicationService = TestBed.get(LoanApplicationService);
    expect(service).toBeTruthy();
  });
});
