import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ClinicalOrderService } from './ClinicalOrder.service';

describe('ClinicalOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ClinicalOrderService] });
	});

  it('should be created', () => {
    const service: ClinicalOrderService = TestBed.get(ClinicalOrderService);
    expect(service).toBeTruthy();
  });
});
