import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PolicyCoverageService } from './PolicyCoverage.service';

describe('PolicyCoverageService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PolicyCoverageService] });
	});

  it('should be created', () => {
    const service: PolicyCoverageService = TestBed.get(PolicyCoverageService);
    expect(service).toBeTruthy();
  });
});
