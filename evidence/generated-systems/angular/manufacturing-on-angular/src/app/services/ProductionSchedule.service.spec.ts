import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ProductionScheduleService } from './ProductionSchedule.service';

describe('ProductionScheduleService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ProductionScheduleService] });
	});

  it('should be created', () => {
    const service: ProductionScheduleService = TestBed.get(ProductionScheduleService);
    expect(service).toBeTruthy();
  });
});
