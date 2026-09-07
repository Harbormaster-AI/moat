import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AgreementService } from './Agreement.service';

describe('AgreementService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AgreementService] });
	});

  it('should be created', () => {
    const service: AgreementService = TestBed.get(AgreementService);
    expect(service).toBeTruthy();
  });
});
