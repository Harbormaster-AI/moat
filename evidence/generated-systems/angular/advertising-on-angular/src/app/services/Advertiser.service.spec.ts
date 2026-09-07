import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AdvertiserService } from './Advertiser.service';

describe('AdvertiserService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AdvertiserService] });
	});

  it('should be created', () => {
    const service: AdvertiserService = TestBed.get(AdvertiserService);
    expect(service).toBeTruthy();
  });
});
