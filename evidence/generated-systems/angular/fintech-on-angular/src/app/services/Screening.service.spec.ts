import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ScreeningService } from './Screening.service';

describe('ScreeningService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ScreeningService] });
	});

  it('should be created', () => {
    const service: ScreeningService = TestBed.get(ScreeningService);
    expect(service).toBeTruthy();
  });
});
