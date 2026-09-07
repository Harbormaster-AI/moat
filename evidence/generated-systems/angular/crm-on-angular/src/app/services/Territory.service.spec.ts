import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { TerritoryService } from './Territory.service';

describe('TerritoryService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [TerritoryService] });
	});

  it('should be created', () => {
    const service: TerritoryService = TestBed.get(TerritoryService);
    expect(service).toBeTruthy();
  });
});
