import { TestBed } from '@angular/core/testing';
import { FormGroup,  FormBuilder,  Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

import { QualityRuleService } from './QualityRule.service';

describe('QualityRuleService', () => {
  	beforeEach(() => {
	  TestBed.configureTestingModule({ imports: [HttpClient, FormGroup, FormBuilder, Validators], providers: [QualityRuleService] });
	});

  it('should be created', () => {
    const service: QualityRuleService = TestBed.get(QualityRuleService);
    expect(service).toBeTruthy();
  });
});
