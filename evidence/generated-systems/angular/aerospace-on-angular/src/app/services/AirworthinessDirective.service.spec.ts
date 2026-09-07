import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AirworthinessDirectiveService } from './AirworthinessDirective.service';

describe('AirworthinessDirectiveService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AirworthinessDirectiveService] });
	});

  it('should be created', () => {
    const service: AirworthinessDirectiveService = TestBed.get(AirworthinessDirectiveService);
    expect(service).toBeTruthy();
  });
});
