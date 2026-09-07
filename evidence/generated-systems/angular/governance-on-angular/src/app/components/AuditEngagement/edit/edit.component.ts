import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AuditEngagementService } from '../../../services/AuditEngagement.service';
import { SubBaseComponent } from '../../AuditEngagement/sub.base.component';


@Component({
    selector: 'app-edit-auditEngagement',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAuditEngagementComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AuditEngagement';

    auditEngagementForm: FormGroup;
    auditEngagement: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AuditEngagementService,
        private fb: FormBuilder
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

    
    updateAuditEngagement(title, startDate, endDate, AuditProgram, BusinessUnits, ControlTests, Workpapers, Findings, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAuditEngagement(title, startDate, endDate, AuditProgram, BusinessUnits, ControlTests, Workpapers, Findings, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAuditEngagement']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAuditEngagement(params['id']).subscribe(res => {
                this.auditEngagement = res;
            });
        });
    }
}