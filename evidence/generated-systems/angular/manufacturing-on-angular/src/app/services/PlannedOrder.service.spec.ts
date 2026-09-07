import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PlannedOrderService } from './PlannedOrder.service';

describe('PlannedOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PlannedOrderService] });
	});

  it('should be created', () => {
    const service: PlannedOrderService = TestBed.get(PlannedOrderService);
    expect(service).toBeTruthy();
  });
});
