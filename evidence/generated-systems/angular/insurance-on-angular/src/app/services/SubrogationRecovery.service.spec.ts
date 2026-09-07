import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SubrogationRecoveryService } from './SubrogationRecovery.service';

describe('SubrogationRecoveryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SubrogationRecoveryService] });
	});

  it('should be created', () => {
    const service: SubrogationRecoveryService = TestBed.get(SubrogationRecoveryService);
    expect(service).toBeTruthy();
  });
});
