import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CompliancePolicyService } from '../../../services/CompliancePolicy.service';
import { CompliancePolicy } from '../../../models/CompliancePolicy';
import { SubBaseComponent } from '../../CompliancePolicy/sub.base.component';

@Component({
    selector: 'app-create-compliancePolicy',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCompliancePolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Add CompliancePolicy';

    compliancePolicyForm: FormGroup;
    compliancePolicy: CompliancePolicy;

    constructor( http: HttpClient,
        private compliancePolicyService: CompliancePolicyService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.compliancePolicyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      policyCode: ['', Validators.required],
      description: ['', Validators.required],
      Institution: ['', ],
      Status: ['', ]
        });
    }

    
    addCompliancePolicy(name, policyCode, description, Institution, Status): void {
        this.compliancePolicyService
        .addCompliancePolicy(name, policyCode, description, Institution, Status)
            .subscribe(() => {
                this.router.navigate(['/indexCompliancePolicy']);
            });
    }

    ngOnInit(): void {
    }
}