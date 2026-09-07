import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PolicyAcknowledgementService } from './PolicyAcknowledgement.service';

describe('PolicyAcknowledgementService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PolicyAcknowledgementService] });
	});

  it('should be created', () => {
    const service: PolicyAcknowledgementService = TestBed.get(PolicyAcknowledgementService);
    expect(service).toBeTruthy();
  });
});
