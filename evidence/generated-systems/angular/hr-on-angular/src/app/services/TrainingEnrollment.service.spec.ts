import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TrainingEnrollmentService } from './TrainingEnrollment.service';

describe('TrainingEnrollmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TrainingEnrollmentService] });
	});

  it('should be created', () => {
    const service: TrainingEnrollmentService = TestBed.get(TrainingEnrollmentService);
    expect(service).toBeTruthy();
  });
});
