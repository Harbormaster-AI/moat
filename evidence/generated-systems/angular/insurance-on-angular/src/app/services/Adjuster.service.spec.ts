import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AdjusterService } from './Adjuster.service';

describe('AdjusterService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AdjusterService] });
	});

  it('should be created', () => {
    const service: AdjusterService = TestBed.get(AdjusterService);
    expect(service).toBeTruthy();
  });
});
