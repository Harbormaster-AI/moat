import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { InboundShipmentLineService } from './InboundShipmentLine.service';

describe('InboundShipmentLineService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [InboundShipmentLineService] });
	});

  it('should be created', () => {
    const service: InboundShipmentLineService = TestBed.get(InboundShipmentLineService);
    expect(service).toBeTruthy();
  });
});
