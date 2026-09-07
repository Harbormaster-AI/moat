import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { WorkShiftService } from './WorkShift.service';

describe('WorkShiftService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [WorkShiftService] });
	});

  it('should be created', () => {
    const service: WorkShiftService = TestBed.get(WorkShiftService);
    expect(service).toBeTruthy();
  });
});
