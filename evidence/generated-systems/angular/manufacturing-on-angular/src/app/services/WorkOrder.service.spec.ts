import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { WorkOrderService } from './WorkOrder.service';

describe('WorkOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [WorkOrderService] });
	});

  it('should be created', () => {
    const service: WorkOrderService = TestBed.get(WorkOrderService);
    expect(service).toBeTruthy();
  });
});
