import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RoutingService } from './Routing.service';

describe('RoutingService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RoutingService] });
	});

  it('should be created', () => {
    const service: RoutingService = TestBed.get(RoutingService);
    expect(service).toBeTruthy();
  });
});
