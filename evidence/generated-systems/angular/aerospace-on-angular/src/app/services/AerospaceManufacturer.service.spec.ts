import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AerospaceManufacturerService } from './AerospaceManufacturer.service';

describe('AerospaceManufacturerService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AerospaceManufacturerService] });
	});

  it('should be created', () => {
    const service: AerospaceManufacturerService = TestBed.get(AerospaceManufacturerService);
    expect(service).toBeTruthy();
  });
});
