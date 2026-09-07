import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { JobProfileService } from './JobProfile.service';

describe('JobProfileService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [JobProfileService] });
	});

  it('should be created', () => {
    const service: JobProfileService = TestBed.get(JobProfileService);
    expect(service).toBeTruthy();
  });
});
