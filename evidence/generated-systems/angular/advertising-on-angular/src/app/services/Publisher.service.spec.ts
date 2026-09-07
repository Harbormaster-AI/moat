import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PublisherService } from './Publisher.service';

describe('PublisherService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PublisherService] });
	});

  it('should be created', () => {
    const service: PublisherService = TestBed.get(PublisherService);
    expect(service).toBeTruthy();
  });
});
