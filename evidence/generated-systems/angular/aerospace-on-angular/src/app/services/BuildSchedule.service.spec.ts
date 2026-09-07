import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BuildScheduleService } from './BuildSchedule.service';

describe('BuildScheduleService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BuildScheduleService] });
	});

  it('should be created', () => {
    const service: BuildScheduleService = TestBed.get(BuildScheduleService);
    expect(service).toBeTruthy();
  });
});
