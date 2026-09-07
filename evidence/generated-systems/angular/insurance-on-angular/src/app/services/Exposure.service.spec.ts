import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ExposureService } from './Exposure.service';

describe('ExposureService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ExposureService] });
	});

  it('should be created', () => {
    const service: ExposureService = TestBed.get(ExposureService);
    expect(service).toBeTruthy();
  });
});
