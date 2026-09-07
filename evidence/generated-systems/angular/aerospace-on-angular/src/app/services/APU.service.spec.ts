import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { APUService } from './APU.service';

describe('APUService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [APUService] });
	});

  it('should be created', () => {
    const service: APUService = TestBed.get(APUService);
    expect(service).toBeTruthy();
  });
});
