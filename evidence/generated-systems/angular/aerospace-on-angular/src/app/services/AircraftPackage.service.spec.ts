import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AircraftPackageService } from './AircraftPackage.service';

describe('AircraftPackageService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AircraftPackageService] });
	});

  it('should be created', () => {
    const service: AircraftPackageService = TestBed.get(AircraftPackageService);
    expect(service).toBeTruthy();
  });
});
