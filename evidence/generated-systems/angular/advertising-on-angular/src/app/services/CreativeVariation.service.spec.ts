import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CreativeVariationService } from './CreativeVariation.service';

describe('CreativeVariationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CreativeVariationService] });
	});

  it('should be created', () => {
    const service: CreativeVariationService = TestBed.get(CreativeVariationService);
    expect(service).toBeTruthy();
  });
});
