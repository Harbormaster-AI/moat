import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ApprovalService } from '../../../services/Approval.service';
import { SubBaseComponent } from '../../Approval/sub.base.component';


@Component({
    selector: 'app-edit-approval',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditApprovalComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Approval';

    approvalForm: FormGroup;
    approval: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ApprovalService,
        private fb: FormBuilder
) {
        super(http);
        this.approvalForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  approverComment: ['', Validators.required],
      actionDate: ['', Validators.required],
      Approver: ['', ],
      Timesheet: ['', ],
      LeaveRequest: ['', ],
      Status: ['', ]
        });
    }

    
    updateApproval(approverComment, actionDate, Approver, Timesheet, LeaveRequest, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateApproval(approverComment, actionDate, Approver, Timesheet, LeaveRequest, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexApproval']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getApproval(params['id']).subscribe(res => {
                this.approval = res;
            });
        });
    }
}