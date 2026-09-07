import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PayrollRunService } from './PayrollRun.service';

describe('PayrollRunService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PayrollRunService] });
	});

  it('should be created', () => {
    const service: PayrollRunService = TestBed.get(PayrollRunService);
    expect(service).toBeTruthy();
  });
});
