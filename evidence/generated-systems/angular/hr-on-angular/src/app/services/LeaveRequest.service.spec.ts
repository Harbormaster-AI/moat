import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LeaveRequestService } from './LeaveRequest.service';

describe('LeaveRequestService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LeaveRequestService] });
	});

  it('should be created', () => {
    const service: LeaveRequestService = TestBed.get(LeaveRequestService);
    expect(service).toBeTruthy();
  });
});
