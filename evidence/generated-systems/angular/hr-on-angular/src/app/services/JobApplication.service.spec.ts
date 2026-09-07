import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { JobApplicationService } from './JobApplication.service';

describe('JobApplicationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [JobApplicationService] });
	});

  it('should be created', () => {
    const service: JobApplicationService = TestBed.get(JobApplicationService);
    expect(service).toBeTruthy();
  });
});
