import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ClaimService } from './Claim.service';

describe('ClaimService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ClaimService] });
	});

  it('should be created', () => {
    const service: ClaimService = TestBed.get(ClaimService);
    expect(service).toBeTruthy();
  });
});
