import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { ProductionLineService } from './ProductionLine.service';

describe('ProductionLineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [ProductionLineService] });
	});

  it('should be created', () => {
    const service: ProductionLineService = TestBed.get(ProductionLineService);
    expect(service).toBeTruthy();
  });
});
