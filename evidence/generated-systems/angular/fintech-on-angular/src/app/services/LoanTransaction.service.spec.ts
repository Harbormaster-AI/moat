import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LoanTransactionService } from './LoanTransaction.service';

describe('LoanTransactionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LoanTransactionService] });
	});

  it('should be created', () => {
    const service: LoanTransactionService = TestBed.get(LoanTransactionService);
    expect(service).toBeTruthy();
  });
});
