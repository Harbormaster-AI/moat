import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RunMetricService } from './RunMetric.service';

describe('RunMetricService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RunMetricService] });
	});

  it('should be created', () => {
    const service: RunMetricService = TestBed.get(RunMetricService);
    expect(service).toBeTruthy();
  });
});
