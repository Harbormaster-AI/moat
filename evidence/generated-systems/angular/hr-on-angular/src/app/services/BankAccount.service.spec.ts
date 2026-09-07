import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BankAccountService } from './BankAccount.service';

describe('BankAccountService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BankAccountService] });
	});

  it('should be created', () => {
    const service: BankAccountService = TestBed.get(BankAccountService);
    expect(service).toBeTruthy();
  });
});
