import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AttestationService } from './Attestation.service';

describe('AttestationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AttestationService] });
	});

  it('should be created', () => {
    const service: AttestationService = TestBed.get(AttestationService);
    expect(service).toBeTruthy();
  });
});
