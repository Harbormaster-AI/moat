import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CreditorService } from './Creditor.service';

describe('CreditorService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CreditorService] });
	});

  it('should be created', () => {
    const service: CreditorService = TestBed.get(CreditorService);
    expect(service).toBeTruthy();
  });
});
