import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MRPRunService } from './MRPRun.service';

describe('MRPRunService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MRPRunService] });
	});

  it('should be created', () => {
    const service: MRPRunService = TestBed.get(MRPRunService);
    expect(service).toBeTruthy();
  });
});
