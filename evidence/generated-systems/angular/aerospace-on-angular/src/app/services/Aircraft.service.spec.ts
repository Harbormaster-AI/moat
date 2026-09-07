import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AircraftService } from './Aircraft.service';

describe('AircraftService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AircraftService] });
	});

  it('should be created', () => {
    const service: AircraftService = TestBed.get(AircraftService);
    expect(service).toBeTruthy();
  });
});
