import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InsurerService } from './Insurer.service';

describe('InsurerService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InsurerService] });
	});

  it('should be created', () => {
    const service: InsurerService = TestBed.get(InsurerService);
    expect(service).toBeTruthy();
  });
});
