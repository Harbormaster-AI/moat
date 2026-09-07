import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CompliancePolicyService } from '../../../services/CompliancePolicy.service';
import { SubBaseComponent } from '../../CompliancePolicy/sub.base.component';


@Component({
    selector: 'app-edit-compliancePolicy',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCompliancePolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CompliancePolicy';

    compliancePolicyForm: FormGroup;
    compliancePolicy: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CompliancePolicyService,
        private fb: FormBuilder
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

    
    updateCompliancePolicy(name, policyCode, description, Institution, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCompliancePolicy(name, policyCode, description, Institution, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCompliancePolicy']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCompliancePolicy(params['id']).subscribe(res => {
                this.compliancePolicy = res;
            });
        });
    }
}