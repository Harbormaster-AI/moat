import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ApprovalService } from '../../../services/Approval.service';
import { Approval } from '../../../models/Approval';
import { SubBaseComponent } from '../../Approval/sub.base.component';

@Component({
    selector: 'app-create-approval',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateApprovalComponent extends SubBaseComponent implements OnInit {

    title = 'Add Approval';

    approvalForm: FormGroup;
    approval: Approval;

    constructor( http: HttpClient,
        private approvalService: ApprovalService,
        private fb: FormBuilder,
        private router: Router
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

    
    addApproval(approverComment, actionDate, Approver, Timesheet, LeaveRequest, Status): void {
        this.approvalService
        .addApproval(approverComment, actionDate, Approver, Timesheet, LeaveRequest, Status)
            .subscribe(() => {
                this.router.navigate(['/indexApproval']);
            });
    }

    ngOnInit(): void {
    }
}