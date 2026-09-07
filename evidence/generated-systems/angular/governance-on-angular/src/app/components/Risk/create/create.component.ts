import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { RiskService } from '../../../services/Risk.service';
import { Risk } from '../../../models/Risk';
import { SubBaseComponent } from '../../Risk/sub.base.component';

@Component({
    selector: 'app-create-risk',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateRiskComponent extends SubBaseComponent implements OnInit {

    title = 'Add Risk';

    riskForm: FormGroup;
    risk: Risk;

    constructor( http: HttpClient,
        private riskService: RiskService,
        private fb: FormBuilder,
        private router: Router
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

    
    addRisk(name, description, inherentRiskScore, residualRiskScore, Organization, Controls, Assessments, Issues, Findings, Category, Impact, Likelihood, Status): void {
        this.riskService
        .addRisk(name, description, inherentRiskScore, residualRiskScore, Organization, Controls, Assessments, Issues, Findings, Category, Impact, Likelihood, Status)
            .subscribe(() => {
                this.router.navigate(['/indexRisk']);
            });
    }

    ngOnInit(): void {
    }
}