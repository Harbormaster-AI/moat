import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ComplianceRequirementService } from '../../../services/ComplianceRequirement.service';
import { SubBaseComponent } from '../../ComplianceRequirement/sub.base.component';


@Component({
    selector: 'app-edit-complianceRequirement',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditComplianceRequirementComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ComplianceRequirement';

    complianceRequirementForm: FormGroup;
    complianceRequirement: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ComplianceRequirementService,
        private fb: FormBuilder
) {
        super(http);
        this.complianceRequirementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      source: ['', Validators.required],
      citation: ['', Validators.required],
      ComplianceProgram: ['', ],
      Policies: ['', ],
      Controls: ['', ],
      Obligations: ['', ],
      Applicability: ['', ],
      Status: ['', ]
        });
    }

    
    updateComplianceRequirement(name, source, citation, ComplianceProgram, Policies, Controls, Obligations, Applicability, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateComplianceRequirement(name, source, citation, ComplianceProgram, Policies, Controls, Obligations, Applicability, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexComplianceRequirement']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getComplianceRequirement(params['id']).subscribe(res => {
                this.complianceRequirement = res;
            });
        });
    }
}