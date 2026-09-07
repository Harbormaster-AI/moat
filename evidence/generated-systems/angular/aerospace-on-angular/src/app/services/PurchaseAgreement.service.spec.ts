import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PurchaseAgreementService } from './PurchaseAgreement.service';

describe('PurchaseAgreementService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PurchaseAgreementService] });
	});

  it('should be created', () => {
    const service: PurchaseAgreementService = TestBed.get(PurchaseAgreementService);
    expect(service).toBeTruthy();
  });
});
