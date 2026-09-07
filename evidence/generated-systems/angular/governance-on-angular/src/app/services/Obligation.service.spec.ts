import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ObligationService } from './Obligation.service';

describe('ObligationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ObligationService] });
	});

  it('should be created', () => {
    const service: ObligationService = TestBed.get(ObligationService);
    expect(service).toBeTruthy();
  });
});
