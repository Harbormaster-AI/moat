import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ComplianceAlertService } from '../../../services/ComplianceAlert.service';
import { ComplianceAlert } from '../../../models/ComplianceAlert';
import { SubBaseComponent } from '../../ComplianceAlert/sub.base.component';

@Component({
    selector: 'app-create-complianceAlert',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateComplianceAlertComponent extends SubBaseComponent implements OnInit {

    title = 'Add ComplianceAlert';

    complianceAlertForm: FormGroup;
    complianceAlert: ComplianceAlert;

    constructor( http: HttpClient,
        private complianceAlertService: ComplianceAlertService,
        private fb: FormBuilder,
        private router: Router
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

    
    addComplianceAlert(alertCode, raisedAt, notes, Screening, Transaction, Severity, Status): void {
        this.complianceAlertService
        .addComplianceAlert(alertCode, raisedAt, notes, Screening, Transaction, Severity, Status)
            .subscribe(() => {
                this.router.navigate(['/indexComplianceAlert']);
            });
    }

    ngOnInit(): void {
    }
}