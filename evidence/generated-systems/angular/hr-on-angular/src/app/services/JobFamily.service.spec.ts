import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { JobFamilyService } from './JobFamily.service';

describe('JobFamilyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [JobFamilyService] });
	});

  it('should be created', () => {
    const service: JobFamilyService = TestBed.get(JobFamilyService);
    expect(service).toBeTruthy();
  });
});
