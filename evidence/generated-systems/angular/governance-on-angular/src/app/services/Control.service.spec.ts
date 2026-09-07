import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ControlService } from './Control.service';

describe('ControlService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ControlService] });
	});

  it('should be created', () => {
    const service: ControlService = TestBed.get(ControlService);
    expect(service).toBeTruthy();
  });
});
