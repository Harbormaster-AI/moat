import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MedicalSupplierService } from './MedicalSupplier.service';

describe('MedicalSupplierService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MedicalSupplierService] });
	});

  it('should be created', () => {
    const service: MedicalSupplierService = TestBed.get(MedicalSupplierService);
    expect(service).toBeTruthy();
  });
});
