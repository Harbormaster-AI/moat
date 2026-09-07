import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AircraftOptionService } from './AircraftOption.service';

describe('AircraftOptionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AircraftOptionService] });
	});

  it('should be created', () => {
    const service: AircraftOptionService = TestBed.get(AircraftOptionService);
    expect(service).toBeTruthy();
  });
});
