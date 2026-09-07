import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { QualityRuleService } from '../../../services/QualityRule.service';
import { SubBaseComponent } from '../../QualityRule/sub.base.component';


@Component({
    selector: 'app-edit-qualityRule',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditQualityRuleComponent extends SubBaseComponent implements OnInit {

    title = 'Edit QualityRule';

    qualityRuleForm: FormGroup;
    qualityRule: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: QualityRuleService,
        private fb: FormBuilder
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

    
    updateQualityRule(name, threshold, targetField, Dataset, Checks, Dimension, Operator): void {
        this.route.params.subscribe((params) => {

                        this.service.updateQualityRule(name, threshold, targetField, Dataset, Checks, Dimension, Operator, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexQualityRule']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getQualityRule(params['id']).subscribe(res => {
                this.qualityRule = res;
            });
        });
    }
}