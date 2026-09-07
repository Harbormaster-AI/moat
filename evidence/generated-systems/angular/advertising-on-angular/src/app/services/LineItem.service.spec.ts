import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { LineItemService } from './LineItem.service';

describe('LineItemService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [LineItemService] });
	});

  it('should be created', () => {
    const service: LineItemService = TestBed.get(LineItemService);
    expect(service).toBeTruthy();
  });
});
