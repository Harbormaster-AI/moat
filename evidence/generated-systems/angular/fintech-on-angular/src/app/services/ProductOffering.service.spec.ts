import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ProductOfferingService } from './ProductOffering.service';

describe('ProductOfferingService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ProductOfferingService] });
	});

  it('should be created', () => {
    const service: ProductOfferingService = TestBed.get(ProductOfferingService);
    expect(service).toBeTruthy();
  });
});
