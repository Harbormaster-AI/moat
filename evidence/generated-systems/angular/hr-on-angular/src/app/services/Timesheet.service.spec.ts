import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TimesheetService } from './Timesheet.service';

describe('TimesheetService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TimesheetService] });
	});

  it('should be created', () => {
    const service: TimesheetService = TestBed.get(TimesheetService);
    expect(service).toBeTruthy();
  });
});
