import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ShiftAssignmentService } from './ShiftAssignment.service';

describe('ShiftAssignmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ShiftAssignmentService] });
	});

  it('should be created', () => {
    const service: ShiftAssignmentService = TestBed.get(ShiftAssignmentService);
    expect(service).toBeTruthy();
  });
});
