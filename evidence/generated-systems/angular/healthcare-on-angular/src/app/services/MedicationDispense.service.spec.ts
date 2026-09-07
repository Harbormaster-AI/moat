import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MedicationDispenseService } from './MedicationDispense.service';

describe('MedicationDispenseService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MedicationDispenseService] });
	});

  it('should be created', () => {
    const service: MedicationDispenseService = TestBed.get(MedicationDispenseService);
    expect(service).toBeTruthy();
  });
});
