import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TimeEntryService } from './TimeEntry.service';

describe('TimeEntryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TimeEntryService] });
	});

  it('should be created', () => {
    const service: TimeEntryService = TestBed.get(TimeEntryService);
    expect(service).toBeTruthy();
  });
});
