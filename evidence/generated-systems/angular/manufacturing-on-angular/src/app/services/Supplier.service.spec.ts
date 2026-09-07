import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SupplierService } from './Supplier.service';

describe('SupplierService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SupplierService] });
	});

  it('should be created', () => {
    const service: SupplierService = TestBed.get(SupplierService);
    expect(service).toBeTruthy();
  });
});
