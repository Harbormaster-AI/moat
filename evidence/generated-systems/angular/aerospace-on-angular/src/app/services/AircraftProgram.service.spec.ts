import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AircraftProgramService } from './AircraftProgram.service';

describe('AircraftProgramService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AircraftProgramService] });
	});

  it('should be created', () => {
    const service: AircraftProgramService = TestBed.get(AircraftProgramService);
    expect(service).toBeTruthy();
  });
});
