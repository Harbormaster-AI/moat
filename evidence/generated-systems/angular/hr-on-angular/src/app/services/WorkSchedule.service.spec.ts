import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { WorkScheduleService } from './WorkSchedule.service';

describe('WorkScheduleService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [WorkScheduleService] });
	});

  it('should be created', () => {
    const service: WorkScheduleService = TestBed.get(WorkScheduleService);
    expect(service).toBeTruthy();
  });
});
