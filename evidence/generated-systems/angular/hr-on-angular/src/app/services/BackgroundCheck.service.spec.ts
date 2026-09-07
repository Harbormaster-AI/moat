import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BackgroundCheckService } from './BackgroundCheck.service';

describe('BackgroundCheckService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BackgroundCheckService] });
	});

  it('should be created', () => {
    const service: BackgroundCheckService = TestBed.get(BackgroundCheckService);
    expect(service).toBeTruthy();
  });
});
