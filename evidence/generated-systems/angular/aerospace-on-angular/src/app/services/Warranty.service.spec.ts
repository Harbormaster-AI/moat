import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { WarrantyService } from './Warranty.service';

describe('WarrantyService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [WarrantyService] });
	});

  it('should be created', () => {
    const service: WarrantyService = TestBed.get(WarrantyService);
    expect(service).toBeTruthy();
  });
});
