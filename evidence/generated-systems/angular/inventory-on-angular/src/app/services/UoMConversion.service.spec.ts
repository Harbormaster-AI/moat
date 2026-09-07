import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { UoMConversionService } from './UoMConversion.service';

describe('UoMConversionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [UoMConversionService] });
	});

  it('should be created', () => {
    const service: UoMConversionService = TestBed.get(UoMConversionService);
    expect(service).toBeTruthy();
  });
});
