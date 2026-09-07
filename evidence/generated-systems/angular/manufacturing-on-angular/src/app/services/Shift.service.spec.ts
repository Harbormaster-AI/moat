import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ShiftService } from './Shift.service';

describe('ShiftService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ShiftService] });
	});

  it('should be created', () => {
    const service: ShiftService = TestBed.get(ShiftService);
    expect(service).toBeTruthy();
  });
});
