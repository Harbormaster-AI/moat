import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TimeSeriesService } from './TimeSeries.service';

describe('TimeSeriesService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TimeSeriesService] });
	});

  it('should be created', () => {
    const service: TimeSeriesService = TestBed.get(TimeSeriesService);
    expect(service).toBeTruthy();
  });
});
