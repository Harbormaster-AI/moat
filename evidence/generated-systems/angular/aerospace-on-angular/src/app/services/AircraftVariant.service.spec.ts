import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AircraftVariantService } from './AircraftVariant.service';

describe('AircraftVariantService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AircraftVariantService] });
	});

  it('should be created', () => {
    const service: AircraftVariantService = TestBed.get(AircraftVariantService);
    expect(service).toBeTruthy();
  });
});
