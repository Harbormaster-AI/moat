import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ControlTest_Service } from './ControlTest_.service';

describe('ControlTest_Service', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ControlTest_Service] });
	});

  it('should be created', () => {
    const service: ControlTest_Service = TestBed.get(ControlTest_Service);
    expect(service).toBeTruthy();
  });
});
