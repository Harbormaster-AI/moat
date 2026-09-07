import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ComplianceAlertService } from '../../../services/ComplianceAlert.service';
import { SubBaseComponent } from '../../ComplianceAlert/sub.base.component';


@Component({
    selector: 'app-edit-complianceAlert',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditComplianceAlertComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ComplianceAlert';

    complianceAlertForm: FormGroup;
    complianceAlert: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ComplianceAlertService,
        private fb: FormBuilder
) {
        super(http);
        this.complianceAlertForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  alertCode: ['', Validators.required],
      raisedAt: ['', Validators.required],
      notes: ['', Validators.required],
      Screening: ['', ],
      Transaction: ['', ],
      Severity: ['', ],
      Status: ['', ]
        });
    }

    
    updateComplianceAlert(alertCode, raisedAt, notes, Screening, Transaction, Severity, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateComplianceAlert(alertCode, raisedAt, notes, Screening, Transaction, Severity, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexComplianceAlert']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getComplianceAlert(params['id']).subscribe(res => {
                this.complianceAlert = res;
            });
        });
    }
}