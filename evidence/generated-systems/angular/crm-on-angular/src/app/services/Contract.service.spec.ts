import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ContractService } from './Contract.service';

describe('ContractService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ContractService] });
	});

  it('should be created', () => {
    const service: ContractService = TestBed.get(ContractService);
    expect(service).toBeTruthy();
  });
});
