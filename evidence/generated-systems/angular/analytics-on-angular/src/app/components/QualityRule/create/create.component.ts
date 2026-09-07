import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { QualityRuleService } from '../../../services/QualityRule.service';
import { QualityRule } from '../../../models/QualityRule';
import { SubBaseComponent } from '../../QualityRule/sub.base.component';

@Component({
    selector: 'app-create-qualityRule',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateQualityRuleComponent extends SubBaseComponent implements OnInit {

    title = 'Add QualityRule';

    qualityRuleForm: FormGroup;
    qualityRule: QualityRule;

    constructor( http: HttpClient,
        private qualityRuleService: QualityRuleService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.qualityRuleForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      threshold: ['', Validators.required],
      targetField: ['', Validators.required],
      Dataset: ['', ],
      Checks: ['', ],
      Dimension: ['', ],
      Operator: ['', ]
        });
    }

    
    addQualityRule(name, threshold, targetField, Dataset, Checks, Dimension, Operator): void {
        this.qualityRuleService
        .addQualityRule(name, threshold, targetField, Dataset, Checks, Dimension, Operator)
            .subscribe(() => {
                this.router.navigate(['/indexQualityRule']);
            });
    }

    ngOnInit(): void {
    }
}