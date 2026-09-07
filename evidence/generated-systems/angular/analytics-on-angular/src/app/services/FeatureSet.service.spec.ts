import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { FeatureSetService } from './FeatureSet.service';

describe('FeatureSetService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [FeatureSetService] });
	});

  it('should be created', () => {
    const service: FeatureSetService = TestBed.get(FeatureSetService);
    expect(service).toBeTruthy();
  });
});
