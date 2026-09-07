import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AircraftOrderService } from './AircraftOrder.service';

describe('AircraftOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AircraftOrderService] });
	});

  it('should be created', () => {
    const service: AircraftOrderService = TestBed.get(AircraftOrderService);
    expect(service).toBeTruthy();
  });
});
