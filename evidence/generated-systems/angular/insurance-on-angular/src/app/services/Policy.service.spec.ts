import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PolicyService } from './Policy.service';

describe('PolicyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PolicyService] });
	});

  it('should be created', () => {
    const service: PolicyService = TestBed.get(PolicyService);
    expect(service).toBeTruthy();
  });
});
