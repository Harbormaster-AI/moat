import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InventoryItemService } from './InventoryItem.service';

describe('InventoryItemService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InventoryItemService] });
	});

  it('should be created', () => {
    const service: InventoryItemService = TestBed.get(InventoryItemService);
    expect(service).toBeTruthy();
  });
});
