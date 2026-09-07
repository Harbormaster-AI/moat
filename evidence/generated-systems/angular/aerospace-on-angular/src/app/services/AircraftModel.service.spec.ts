import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AircraftModelService } from './AircraftModel.service';

describe('AircraftModelService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AircraftModelService] });
	});

  it('should be created', () => {
    const service: AircraftModelService = TestBed.get(AircraftModelService);
    expect(service).toBeTruthy();
  });
});
