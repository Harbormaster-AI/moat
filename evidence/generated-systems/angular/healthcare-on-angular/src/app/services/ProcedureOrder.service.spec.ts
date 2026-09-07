import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ProcedureOrderService } from './ProcedureOrder.service';

describe('ProcedureOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ProcedureOrderService] });
	});

  it('should be created', () => {
    const service: ProcedureOrderService = TestBed.get(ProcedureOrderService);
    expect(service).toBeTruthy();
  });
});
