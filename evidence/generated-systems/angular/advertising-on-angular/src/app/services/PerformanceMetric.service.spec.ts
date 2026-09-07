import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PerformanceMetricService } from './PerformanceMetric.service';

describe('PerformanceMetricService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PerformanceMetricService] });
	});

  it('should be created', () => {
    const service: PerformanceMetricService = TestBed.get(PerformanceMetricService);
    expect(service).toBeTruthy();
  });
});
