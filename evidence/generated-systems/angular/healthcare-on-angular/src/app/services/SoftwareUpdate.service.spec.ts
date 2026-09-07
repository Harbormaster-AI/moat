import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SoftwareUpdateService } from './SoftwareUpdate.service';

describe('SoftwareUpdateService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SoftwareUpdateService] });
	});

  it('should be created', () => {
    const service: SoftwareUpdateService = TestBed.get(SoftwareUpdateService);
    expect(service).toBeTruthy();
  });
});
