import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LeaveRequestService } from '../../../services/LeaveRequest.service';
import { LeaveRequest } from '../../../models/LeaveRequest';
import { SubBaseComponent } from '../../LeaveRequest/sub.base.component';

@Component({
    selector: 'app-create-leaveRequest',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLeaveRequestComponent extends SubBaseComponent implements OnInit {

    title = 'Add LeaveRequest';

    leaveRequestForm: FormGroup;
    leaveRequest: LeaveRequest;

    constructor( http: HttpClient,
        private leaveRequestService: LeaveRequestService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.leaveRequestForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  requestNumber: ['', Validators.required],
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      reason: ['', Validators.required],
      hours: ['', Validators.required],
      Employee: ['', ],
      LeavePolicy: ['', ],
      Approvals: ['', ],
      Status: ['', ]
        });
    }

    
    addLeaveRequest(requestNumber, startDate, endDate, reason, hours, Employee, LeavePolicy, Approvals, Status): void {
        this.leaveRequestService
        .addLeaveRequest(requestNumber, startDate, endDate, reason, hours, Employee, LeavePolicy, Approvals, Status)
            .subscribe(() => {
                this.router.navigate(['/indexLeaveRequest']);
            });
    }

    ngOnInit(): void {
    }
}