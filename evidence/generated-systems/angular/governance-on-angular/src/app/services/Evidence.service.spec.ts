import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EvidenceService } from './Evidence.service';

describe('EvidenceService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EvidenceService] });
	});

  it('should be created', () => {
    const service: EvidenceService = TestBed.get(EvidenceService);
    expect(service).toBeTruthy();
  });
});
