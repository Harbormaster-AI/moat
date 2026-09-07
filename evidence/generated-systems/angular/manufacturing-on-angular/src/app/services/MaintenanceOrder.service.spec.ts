import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MaintenanceOrderService } from './MaintenanceOrder.service';

describe('MaintenanceOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MaintenanceOrderService] });
	});

  it('should be created', () => {
    const service: MaintenanceOrderService = TestBed.get(MaintenanceOrderService);
    expect(service).toBeTruthy();
  });
});
