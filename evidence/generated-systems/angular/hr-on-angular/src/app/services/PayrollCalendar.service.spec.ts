import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PayrollCalendarService } from './PayrollCalendar.service';

describe('PayrollCalendarService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PayrollCalendarService] });
	});

  it('should be created', () => {
    const service: PayrollCalendarService = TestBed.get(PayrollCalendarService);
    expect(service).toBeTruthy();
  });
});
