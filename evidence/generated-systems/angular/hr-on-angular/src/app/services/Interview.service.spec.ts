import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InterviewService } from './Interview.service';

describe('InterviewService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InterviewService] });
	});

  it('should be created', () => {
    const service: InterviewService = TestBed.get(InterviewService);
    expect(service).toBeTruthy();
  });
});
