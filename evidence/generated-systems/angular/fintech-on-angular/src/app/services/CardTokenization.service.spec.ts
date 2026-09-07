import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { CardTokenizationService } from './CardTokenization.service';

describe('CardTokenizationService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [CardTokenizationService] });
	});

  it('should be created', () => {
    const service: CardTokenizationService = TestBed.get(CardTokenizationService);
    expect(service).toBeTruthy();
  });
});
