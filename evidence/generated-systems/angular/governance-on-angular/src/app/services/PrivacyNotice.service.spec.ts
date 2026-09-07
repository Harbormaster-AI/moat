import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { PrivacyNoticeService } from './PrivacyNotice.service';

describe('PrivacyNoticeService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [PrivacyNoticeService] });
	});

  it('should be created', () => {
    const service: PrivacyNoticeService = TestBed.get(PrivacyNoticeService);
    expect(service).toBeTruthy();
  });
});
