import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { FraudScenarioService } from './FraudScenario.service';

describe('FraudScenarioService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [FraudScenarioService] });
	});

  it('should be created', () => {
    const service: FraudScenarioService = TestBed.get(FraudScenarioService);
    expect(service).toBeTruthy();
  });
});
