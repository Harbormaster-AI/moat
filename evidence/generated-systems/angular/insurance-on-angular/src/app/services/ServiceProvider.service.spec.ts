import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ServiceProviderService } from './ServiceProvider.service';

describe('ServiceProviderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ServiceProviderService] });
	});

  it('should be created', () => {
    const service: ServiceProviderService = TestBed.get(ServiceProviderService);
    expect(service).toBeTruthy();
  });
});
