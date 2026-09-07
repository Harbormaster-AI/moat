import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AuditWorkpaperService } from './AuditWorkpaper.service';

describe('AuditWorkpaperService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AuditWorkpaperService] });
	});

  it('should be created', () => {
    const service: AuditWorkpaperService = TestBed.get(AuditWorkpaperService);
    expect(service).toBeTruthy();
  });
});
