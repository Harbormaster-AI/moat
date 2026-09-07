import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SubscriberService } from './Subscriber.service';

describe('SubscriberService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SubscriberService] });
	});

  it('should be created', () => {
    const service: SubscriberService = TestBed.get(SubscriberService);
    expect(service).toBeTruthy();
  });
});
