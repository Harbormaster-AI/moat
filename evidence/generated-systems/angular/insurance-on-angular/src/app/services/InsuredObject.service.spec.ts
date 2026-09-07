import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InsuredObjectService } from './InsuredObject.service';

describe('InsuredObjectService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InsuredObjectService] });
	});

  it('should be created', () => {
    const service: InsuredObjectService = TestBed.get(InsuredObjectService);
    expect(service).toBeTruthy();
  });
});
