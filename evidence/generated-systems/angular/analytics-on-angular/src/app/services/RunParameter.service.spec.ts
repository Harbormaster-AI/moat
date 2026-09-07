import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RunParameterService } from './RunParameter.service';

describe('RunParameterService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RunParameterService] });
	});

  it('should be created', () => {
    const service: RunParameterService = TestBed.get(RunParameterService);
    expect(service).toBeTruthy();
  });
});
