import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PerformanceCycleService } from './PerformanceCycle.service';

describe('PerformanceCycleService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PerformanceCycleService] });
	});

  it('should be created', () => {
    const service: PerformanceCycleService = TestBed.get(PerformanceCycleService);
    expect(service).toBeTruthy();
  });
});
