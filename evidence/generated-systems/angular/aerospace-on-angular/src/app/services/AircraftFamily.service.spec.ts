import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AircraftFamilyService } from './AircraftFamily.service';

describe('AircraftFamilyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AircraftFamilyService] });
	});

  it('should be created', () => {
    const service: AircraftFamilyService = TestBed.get(AircraftFamilyService);
    expect(service).toBeTruthy();
  });
});
