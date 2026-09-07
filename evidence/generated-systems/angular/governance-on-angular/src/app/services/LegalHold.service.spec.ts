import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LegalHoldService } from './LegalHold.service';

describe('LegalHoldService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LegalHoldService] });
	});

  it('should be created', () => {
    const service: LegalHoldService = TestBed.get(LegalHoldService);
    expect(service).toBeTruthy();
  });
});
