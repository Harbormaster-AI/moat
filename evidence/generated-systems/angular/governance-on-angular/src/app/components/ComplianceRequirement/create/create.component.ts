import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ComplianceRequirementService } from '../../../services/ComplianceRequirement.service';
import { ComplianceRequirement } from '../../../models/ComplianceRequirement';
import { SubBaseComponent } from '../../ComplianceRequirement/sub.base.component';

@Component({
    selector: 'app-create-complianceRequirement',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateComplianceRequirementComponent extends SubBaseComponent implements OnInit {

    title = 'Add ComplianceRequirement';

    complianceRequirementForm: FormGroup;
    complianceRequirement: ComplianceRequirement;

    constructor( http: HttpClient,
        private complianceRequirementService: ComplianceRequirementService,
        private fb: FormBuilder,
        private router: Router
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

    
    addComplianceRequirement(name, source, citation, ComplianceProgram, Policies, Controls, Obligations, Applicability, Status): void {
        this.complianceRequirementService
        .addComplianceRequirement(name, source, citation, ComplianceProgram, Policies, Controls, Obligations, Applicability, Status)
            .subscribe(() => {
                this.router.navigate(['/indexComplianceRequirement']);
            });
    }

    ngOnInit(): void {
    }
}