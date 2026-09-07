import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ObservationService } from './Observation.service';

describe('ObservationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ObservationService] });
	});

  it('should be created', () => {
    const service: ObservationService = TestBed.get(ObservationService);
    expect(service).toBeTruthy();
  });
});
