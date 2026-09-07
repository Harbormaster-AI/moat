import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LeaveRequestService } from '../../../services/LeaveRequest.service';
import { SubBaseComponent } from '../../LeaveRequest/sub.base.component';


@Component({
    selector: 'app-edit-leaveRequest',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLeaveRequestComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LeaveRequest';

    leaveRequestForm: FormGroup;
    leaveRequest: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LeaveRequestService,
        private fb: FormBuilder
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

    
    updateLeaveRequest(requestNumber, startDate, endDate, reason, hours, Employee, LeavePolicy, Approvals, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLeaveRequest(requestNumber, startDate, endDate, reason, hours, Employee, LeavePolicy, Approvals, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLeaveRequest']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLeaveRequest(params['id']).subscribe(res => {
                this.leaveRequest = res;
            });
        });
    }
}