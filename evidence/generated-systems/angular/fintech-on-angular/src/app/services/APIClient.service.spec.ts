import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { APIClientService } from './APIClient.service';

describe('APIClientService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [APIClientService] });
	});

  it('should be created', () => {
    const service: APIClientService = TestBed.get(APIClientService);
    expect(service).toBeTruthy();
  });
});
