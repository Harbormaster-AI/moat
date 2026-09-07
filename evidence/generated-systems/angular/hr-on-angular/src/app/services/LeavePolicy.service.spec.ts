import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LeavePolicyService } from './LeavePolicy.service';

describe('LeavePolicyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LeavePolicyService] });
	});

  it('should be created', () => {
    const service: LeavePolicyService = TestBed.get(LeavePolicyService);
    expect(service).toBeTruthy();
  });
});
