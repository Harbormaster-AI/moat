import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { FlightHealthEventService } from './FlightHealthEvent.service';

describe('FlightHealthEventService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [FlightHealthEventService] });
	});

  it('should be created', () => {
    const service: FlightHealthEventService = TestBed.get(FlightHealthEventService);
    expect(service).toBeTruthy();
  });
});
