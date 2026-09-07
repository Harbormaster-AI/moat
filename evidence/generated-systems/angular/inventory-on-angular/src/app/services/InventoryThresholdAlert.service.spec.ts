import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InventoryThresholdAlertService } from './InventoryThresholdAlert.service';

describe('InventoryThresholdAlertService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InventoryThresholdAlertService] });
	});

  it('should be created', () => {
    const service: InventoryThresholdAlertService = TestBed.get(InventoryThresholdAlertService);
    expect(service).toBeTruthy();
  });
});
