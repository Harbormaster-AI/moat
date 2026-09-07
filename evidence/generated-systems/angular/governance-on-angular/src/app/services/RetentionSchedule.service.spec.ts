import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RetentionScheduleService } from './RetentionSchedule.service';

describe('RetentionScheduleService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RetentionScheduleService] });
	});

  it('should be created', () => {
    const service: RetentionScheduleService = TestBed.get(RetentionScheduleService);
    expect(service).toBeTruthy();
  });
});
