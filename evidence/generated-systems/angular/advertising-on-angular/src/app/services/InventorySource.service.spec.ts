import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InventorySourceService } from './InventorySource.service';

describe('InventorySourceService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InventorySourceService] });
	});

  it('should be created', () => {
    const service: InventorySourceService = TestBed.get(InventorySourceService);
    expect(service).toBeTruthy();
  });
});
