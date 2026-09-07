import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { NonconformanceService } from './Nonconformance.service';

describe('NonconformanceService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [NonconformanceService] });
	});

  it('should be created', () => {
    const service: NonconformanceService = TestBed.get(NonconformanceService);
    expect(service).toBeTruthy();
  });
});
