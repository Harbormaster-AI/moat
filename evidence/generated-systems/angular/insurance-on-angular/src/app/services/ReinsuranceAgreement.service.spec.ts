import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ReinsuranceAgreementService } from './ReinsuranceAgreement.service';

describe('ReinsuranceAgreementService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ReinsuranceAgreementService] });
	});

  it('should be created', () => {
    const service: ReinsuranceAgreementService = TestBed.get(ReinsuranceAgreementService);
    expect(service).toBeTruthy();
  });
});
