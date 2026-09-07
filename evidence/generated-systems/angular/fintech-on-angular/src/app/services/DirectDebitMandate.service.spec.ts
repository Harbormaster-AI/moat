import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DirectDebitMandateService } from './DirectDebitMandate.service';

describe('DirectDebitMandateService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DirectDebitMandateService] });
	});

  it('should be created', () => {
    const service: DirectDebitMandateService = TestBed.get(DirectDebitMandateService);
    expect(service).toBeTruthy();
  });
});
