import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ForecastLineService } from './ForecastLine.service';

describe('ForecastLineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ForecastLineService] });
	});

  it('should be created', () => {
    const service: ForecastLineService = TestBed.get(ForecastLineService);
    expect(service).toBeTruthy();
  });
});
