import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LeavePolicyService } from '../../../services/LeavePolicy.service';
import { LeavePolicy } from '../../../models/LeavePolicy';
import { SubBaseComponent } from '../../LeavePolicy/sub.base.component';

@Component({
    selector: 'app-create-leavePolicy',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLeavePolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Add LeavePolicy';

    leavePolicyForm: FormGroup;
    leavePolicy: LeavePolicy;

    constructor( http: HttpClient,
        private leavePolicyService: LeavePolicyService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.leavePolicyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      accrualRate: ['', Validators.required],
      carryoverAllowed: ['', Validators.required],
      maxBalance: ['', Validators.required],
      Organization: ['', ],
      LeaveRequests: ['', ],
      LeaveCategory: ['', ],
      AccrualUnit: ['', ]
        });
    }

    
    addLeavePolicy(name, accrualRate, carryoverAllowed, maxBalance, Organization, LeaveRequests, LeaveCategory, AccrualUnit): void {
        this.leavePolicyService
        .addLeavePolicy(name, accrualRate, carryoverAllowed, maxBalance, Organization, LeaveRequests, LeaveCategory, AccrualUnit)
            .subscribe(() => {
                this.router.navigate(['/indexLeavePolicy']);
            });
    }

    ngOnInit(): void {
    }
}