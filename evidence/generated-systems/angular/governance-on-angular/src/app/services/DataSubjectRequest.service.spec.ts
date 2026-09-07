import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DataSubjectRequestService } from './DataSubjectRequest.service';

describe('DataSubjectRequestService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DataSubjectRequestService] });
	});

  it('should be created', () => {
    const service: DataSubjectRequestService = TestBed.get(DataSubjectRequestService);
    expect(service).toBeTruthy();
  });
});
