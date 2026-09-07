import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EmploymentAssignmentService } from './EmploymentAssignment.service';

describe('EmploymentAssignmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EmploymentAssignmentService] });
	});

  it('should be created', () => {
    const service: EmploymentAssignmentService = TestBed.get(EmploymentAssignmentService);
    expect(service).toBeTruthy();
  });
});
