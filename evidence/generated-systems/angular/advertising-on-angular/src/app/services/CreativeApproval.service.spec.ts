import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CreativeApprovalService } from './CreativeApproval.service';

describe('CreativeApprovalService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CreativeApprovalService] });
	});

  it('should be created', () => {
    const service: CreativeApprovalService = TestBed.get(CreativeApprovalService);
    expect(service).toBeTruthy();
  });
});
