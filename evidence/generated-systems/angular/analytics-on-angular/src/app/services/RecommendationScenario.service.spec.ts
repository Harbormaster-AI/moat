import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { RecommendationScenarioService } from './RecommendationScenario.service';

describe('RecommendationScenarioService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [RecommendationScenarioService] });
	});

  it('should be created', () => {
    const service: RecommendationScenarioService = TestBed.get(RecommendationScenarioService);
    expect(service).toBeTruthy();
  });
});
