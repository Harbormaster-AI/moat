import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RoleAssignmentService } from './RoleAssignment.service';

describe('RoleAssignmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RoleAssignmentService] });
	});

  it('should be created', () => {
    const service: RoleAssignmentService = TestBed.get(RoleAssignmentService);
    expect(service).toBeTruthy();
  });
});
