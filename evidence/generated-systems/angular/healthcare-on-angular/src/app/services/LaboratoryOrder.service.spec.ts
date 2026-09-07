import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LaboratoryOrderService } from './LaboratoryOrder.service';

describe('LaboratoryOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LaboratoryOrderService] });
	});

  it('should be created', () => {
    const service: LaboratoryOrderService = TestBed.get(LaboratoryOrderService);
    expect(service).toBeTruthy();
  });
});
