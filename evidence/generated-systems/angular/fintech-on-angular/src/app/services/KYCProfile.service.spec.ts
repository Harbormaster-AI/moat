import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { KYCProfileService } from './KYCProfile.service';

describe('KYCProfileService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [KYCProfileService] });
	});

  it('should be created', () => {
    const service: KYCProfileService = TestBed.get(KYCProfileService);
    expect(service).toBeTruthy();
  });
});
