import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ServiceBulletinService } from './ServiceBulletin.service';

describe('ServiceBulletinService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ServiceBulletinService] });
	});

  it('should be created', () => {
    const service: ServiceBulletinService = TestBed.get(ServiceBulletinService);
    expect(service).toBeTruthy();
  });
});
