import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ComplianceProgramService } from '../../../services/ComplianceProgram.service';
import { ComplianceProgram } from '../../../models/ComplianceProgram';
import { SubBaseComponent } from '../../ComplianceProgram/sub.base.component';

@Component({
    selector: 'app-create-complianceProgram',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateComplianceProgramComponent extends SubBaseComponent implements OnInit {

    title = 'Add ComplianceProgram';

    complianceProgramForm: FormGroup;
    complianceProgram: ComplianceProgram;

    constructor( http: HttpClient,
        private complianceProgramService: ComplianceProgramService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.complianceProgramForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      framework: ['', Validators.required],
      Organization: ['', ],
      Requirements: ['', ],
      Controls: ['', ],
      Attestations: ['', ],
      Regulations: ['', ],
      Status: ['', ]
        });
    }

    
    addComplianceProgram(name, framework, Organization, Requirements, Controls, Attestations, Regulations, Status): void {
        this.complianceProgramService
        .addComplianceProgram(name, framework, Organization, Requirements, Controls, Attestations, Regulations, Status)
            .subscribe(() => {
                this.router.navigate(['/indexComplianceProgram']);
            });
    }

    ngOnInit(): void {
    }
}