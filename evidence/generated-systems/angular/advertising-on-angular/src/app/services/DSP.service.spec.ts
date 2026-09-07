import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DSPService } from './DSP.service';

describe('DSPService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DSPService] });
	});

  it('should be created', () => {
    const service: DSPService = TestBed.get(DSPService);
    expect(service).toBeTruthy();
  });
});
