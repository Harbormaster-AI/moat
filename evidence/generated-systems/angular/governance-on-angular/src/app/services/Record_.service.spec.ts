import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { Record_Service } from './Record_.service';

describe('Record_Service', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [Record_Service] });
	});

  it('should be created', () => {
    const service: Record_Service = TestBed.get(Record_Service);
    expect(service).toBeTruthy();
  });
});
