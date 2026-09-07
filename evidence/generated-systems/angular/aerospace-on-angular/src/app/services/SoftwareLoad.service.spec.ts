import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { SoftwareLoadService } from './SoftwareLoad.service';

describe('SoftwareLoadService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [SoftwareLoadService] });
	});

  it('should be created', () => {
    const service: SoftwareLoadService = TestBed.get(SoftwareLoadService);
    expect(service).toBeTruthy();
  });
});
