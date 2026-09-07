import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { JobRequisitionService } from './JobRequisition.service';

describe('JobRequisitionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [JobRequisitionService] });
	});

  it('should be created', () => {
    const service: JobRequisitionService = TestBed.get(JobRequisitionService);
    expect(service).toBeTruthy();
  });
});
