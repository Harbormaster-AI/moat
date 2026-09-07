import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InsertionOrderService } from './InsertionOrder.service';

describe('InsertionOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InsertionOrderService] });
	});

  it('should be created', () => {
    const service: InsertionOrderService = TestBed.get(InsertionOrderService);
    expect(service).toBeTruthy();
  });
});
