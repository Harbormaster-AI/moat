import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InboundShipmentService } from './InboundShipment.service';

describe('InboundShipmentService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InboundShipmentService] });
	});

  it('should be created', () => {
    const service: InboundShipmentService = TestBed.get(InboundShipmentService);
    expect(service).toBeTruthy();
  });
});
