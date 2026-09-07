import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InferenceEndpointService } from './InferenceEndpoint.service';

describe('InferenceEndpointService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InferenceEndpointService] });
	});

  it('should be created', () => {
    const service: InferenceEndpointService = TestBed.get(InferenceEndpointService);
    expect(service).toBeTruthy();
  });
});
