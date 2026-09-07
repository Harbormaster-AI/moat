import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MetricService } from './Metric.service';

describe('MetricService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MetricService] });
	});

  it('should be created', () => {
    const service: MetricService = TestBed.get(MetricService);
    expect(service).toBeTruthy();
  });
});
