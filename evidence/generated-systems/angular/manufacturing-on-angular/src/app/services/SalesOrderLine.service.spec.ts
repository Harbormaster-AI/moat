import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SalesOrderLineService } from './SalesOrderLine.service';

describe('SalesOrderLineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SalesOrderLineService] });
	});

  it('should be created', () => {
    const service: SalesOrderLineService = TestBed.get(SalesOrderLineService);
    expect(service).toBeTruthy();
  });
});
