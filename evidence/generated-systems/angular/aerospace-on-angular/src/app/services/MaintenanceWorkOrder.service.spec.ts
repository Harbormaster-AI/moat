import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { MaintenanceWorkOrderService } from './MaintenanceWorkOrder.service';

describe('MaintenanceWorkOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [MaintenanceWorkOrderService] });
	});

  it('should be created', () => {
    const service: MaintenanceWorkOrderService = TestBed.get(MaintenanceWorkOrderService);
    expect(service).toBeTruthy();
  });
});
