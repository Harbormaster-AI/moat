import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EvaluationMetricService } from './EvaluationMetric.service';

describe('EvaluationMetricService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EvaluationMetricService] });
	});

  it('should be created', () => {
    const service: EvaluationMetricService = TestBed.get(EvaluationMetricService);
    expect(service).toBeTruthy();
  });
});
