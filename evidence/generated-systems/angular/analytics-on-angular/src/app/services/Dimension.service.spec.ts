import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { DimensionService } from './Dimension.service';

describe('DimensionService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [DimensionService] });
	});

  it('should be created', () => {
    const service: DimensionService = TestBed.get(DimensionService);
    expect(service).toBeTruthy();
  });
});
