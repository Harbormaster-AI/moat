import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ApprovalService } from './Approval.service';

describe('ApprovalService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ApprovalService] });
	});

  it('should be created', () => {
    const service: ApprovalService = TestBed.get(ApprovalService);
    expect(service).toBeTruthy();
  });
});
