import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { AssetService } from './Asset.service';

describe('AssetService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [AssetService] });
	});

  it('should be created', () => {
    const service: AssetService = TestBed.get(AssetService);
    expect(service).toBeTruthy();
  });
});
