import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { FeeScheduleService } from './FeeSchedule.service';

describe('FeeScheduleService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [FeeScheduleService] });
	});

  it('should be created', () => {
    const service: FeeScheduleService = TestBed.get(FeeScheduleService);
    expect(service).toBeTruthy();
  });
});
