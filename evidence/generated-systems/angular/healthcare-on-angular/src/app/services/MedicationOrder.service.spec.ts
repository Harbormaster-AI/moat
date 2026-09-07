import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MedicationOrderService } from './MedicationOrder.service';

describe('MedicationOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MedicationOrderService] });
	});

  it('should be created', () => {
    const service: MedicationOrderService = TestBed.get(MedicationOrderService);
    expect(service).toBeTruthy();
  });
});
