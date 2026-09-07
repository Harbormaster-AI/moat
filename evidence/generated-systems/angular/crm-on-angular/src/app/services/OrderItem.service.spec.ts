import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { OrderItemService } from './OrderItem.service';

describe('OrderItemService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [OrderItemService] });
	});

  it('should be created', () => {
    const service: OrderItemService = TestBed.get(OrderItemService);
    expect(service).toBeTruthy();
  });
});
