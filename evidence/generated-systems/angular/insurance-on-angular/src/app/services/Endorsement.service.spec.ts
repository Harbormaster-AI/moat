import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { EndorsementService } from './Endorsement.service';

describe('EndorsementService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [EndorsementService] });
	});

  it('should be created', () => {
    const service: EndorsementService = TestBed.get(EndorsementService);
    expect(service).toBeTruthy();
  });
});
