import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CreativeAssetService } from './CreativeAsset.service';

describe('CreativeAssetService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CreativeAssetService] });
	});

  it('should be created', () => {
    const service: CreativeAssetService = TestBed.get(CreativeAssetService);
    expect(service).toBeTruthy();
  });
});
