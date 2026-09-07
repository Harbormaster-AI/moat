import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EmailMessageService } from './EmailMessage.service';

describe('EmailMessageService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EmailMessageService] });
	});

  it('should be created', () => {
    const service: EmailMessageService = TestBed.get(EmailMessageService);
    expect(service).toBeTruthy();
  });
});
