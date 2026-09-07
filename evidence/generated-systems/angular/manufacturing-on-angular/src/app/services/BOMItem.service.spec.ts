import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { BOMItemService } from './BOMItem.service';

describe('BOMItemService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [BOMItemService] });
	});

  it('should be created', () => {
    const service: BOMItemService = TestBed.get(BOMItemService);
    expect(service).toBeTruthy();
  });
});
