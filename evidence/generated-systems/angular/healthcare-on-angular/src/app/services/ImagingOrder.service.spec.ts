import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ImagingOrderService } from './ImagingOrder.service';

describe('ImagingOrderService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ImagingOrderService] });
	});

  it('should be created', () => {
    const service: ImagingOrderService = TestBed.get(ImagingOrderService);
    expect(service).toBeTruthy();
  });
});
