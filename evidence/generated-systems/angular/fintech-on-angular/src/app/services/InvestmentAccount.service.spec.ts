import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InvestmentAccountService } from './InvestmentAccount.service';

describe('InvestmentAccountService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InvestmentAccountService] });
	});

  it('should be created', () => {
    const service: InvestmentAccountService = TestBed.get(InvestmentAccountService);
    expect(service).toBeTruthy();
  });
});
