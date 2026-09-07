import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { RiskService } from '../../../services/Risk.service';
import { SubBaseComponent } from '../../Risk/sub.base.component';


@Component({
    selector: 'app-edit-risk',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditRiskComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Risk';

    riskForm: FormGroup;
    risk: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: RiskService,
        private fb: FormBuilder
) {
        super(http);
        this.riskForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      description: ['', Validators.required],
      inherentRiskScore: ['', Validators.required],
      residualRiskScore: ['', Validators.required],
      Organization: ['', ],
      Controls: ['', ],
      Assessments: ['', ],
      Issues: ['', ],
      Findings: ['', ],
      Category: ['', ],
      Impact: ['', ],
      Likelihood: ['', ],
      Status: ['', ]
        });
    }

    
    updateRisk(name, description, inherentRiskScore, residualRiskScore, Organization, Controls, Assessments, Issues, Findings, Category, Impact, Likelihood, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateRisk(name, description, inherentRiskScore, residualRiskScore, Organization, Controls, Assessments, Issues, Findings, Category, Impact, Likelihood, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexRisk']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getRisk(params['id']).subscribe(res => {
                this.risk = res;
            });
        });
    }
}