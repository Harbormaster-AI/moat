import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ComplianceProgramService } from '../../../services/ComplianceProgram.service';
import { SubBaseComponent } from '../../ComplianceProgram/sub.base.component';


@Component({
    selector: 'app-edit-complianceProgram',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditComplianceProgramComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ComplianceProgram';

    complianceProgramForm: FormGroup;
    complianceProgram: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ComplianceProgramService,
        private fb: FormBuilder
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

    
    updateComplianceProgram(name, framework, Organization, Requirements, Controls, Attestations, Regulations, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateComplianceProgram(name, framework, Organization, Requirements, Controls, Attestations, Regulations, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexComplianceProgram']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getComplianceProgram(params['id']).subscribe(res => {
                this.complianceProgram = res;
            });
        });
    }
}