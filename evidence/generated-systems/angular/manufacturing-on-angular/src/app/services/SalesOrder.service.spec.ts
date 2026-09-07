import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SalesOrderService } from './SalesOrder.service';

describe('SalesOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SalesOrderService] });
	});

  it('should be created', () => {
    const service: SalesOrderService = TestBed.get(SalesOrderService);
    expect(service).toBeTruthy();
  });
});
