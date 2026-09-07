import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuditEngagementService } from '../../../services/AuditEngagement.service';
import { AuditEngagement } from '../../../models/AuditEngagement';
import { SubBaseComponent } from '../../AuditEngagement/sub.base.component';

@Component({
    selector: 'app-create-auditEngagement',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAuditEngagementComponent extends SubBaseComponent implements OnInit {

    title = 'Add AuditEngagement';

    auditEngagementForm: FormGroup;
    auditEngagement: AuditEngagement;

    constructor( http: HttpClient,
        private auditEngagementService: AuditEngagementService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.auditEngagementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      AuditProgram: ['', ],
      BusinessUnits: ['', ],
      ControlTests: ['', ],
      Workpapers: ['', ],
      Findings: ['', ],
      Status: ['', ]
        });
    }

    
    addAuditEngagement(title, startDate, endDate, AuditProgram, BusinessUnits, ControlTests, Workpapers, Findings, Status): void {
        this.auditEngagementService
        .addAuditEngagement(title, startDate, endDate, AuditProgram, BusinessUnits, ControlTests, Workpapers, Findings, Status)
            .subscribe(() => {
                this.router.navigate(['/indexAuditEngagement']);
            });
    }

    ngOnInit(): void {
    }
}