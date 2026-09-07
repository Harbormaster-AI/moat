import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BOMService } from './BOM.service';

describe('BOMService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BOMService] });
	});

  it('should be created', () => {
    const service: BOMService = TestBed.get(BOMService);
    expect(service).toBeTruthy();
  });
});
