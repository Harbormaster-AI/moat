import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TaxWithholdingService } from './TaxWithholding.service';

describe('TaxWithholdingService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TaxWithholdingService] });
	});

  it('should be created', () => {
    const service: TaxWithholdingService = TestBed.get(TaxWithholdingService);
    expect(service).toBeTruthy();
  });
});
