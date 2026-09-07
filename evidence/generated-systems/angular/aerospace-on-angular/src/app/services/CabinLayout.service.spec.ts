import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CabinLayoutService } from './CabinLayout.service';

describe('CabinLayoutService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CabinLayoutService] });
	});

  it('should be created', () => {
    const service: CabinLayoutService = TestBed.get(CabinLayoutService);
    expect(service).toBeTruthy();
  });
});
