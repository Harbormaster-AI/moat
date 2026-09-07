import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { WarehouseService } from './Warehouse.service';

describe('WarehouseService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [WarehouseService] });
	});

  it('should be created', () => {
    const service: WarehouseService = TestBed.get(WarehouseService);
    expect(service).toBeTruthy();
  });
});
