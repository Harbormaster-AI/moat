import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PredictionService } from './Prediction.service';

describe('PredictionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PredictionService] });
	});

  it('should be created', () => {
    const service: PredictionService = TestBed.get(PredictionService);
    expect(service).toBeTruthy();
  });
});
