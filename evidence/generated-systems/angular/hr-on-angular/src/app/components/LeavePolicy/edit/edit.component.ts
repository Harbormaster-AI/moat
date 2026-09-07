import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LeavePolicyService } from '../../../services/LeavePolicy.service';
import { SubBaseComponent } from '../../LeavePolicy/sub.base.component';


@Component({
    selector: 'app-edit-leavePolicy',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLeavePolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LeavePolicy';

    leavePolicyForm: FormGroup;
    leavePolicy: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LeavePolicyService,
        private fb: FormBuilder
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

    
    updateLeavePolicy(name, accrualRate, carryoverAllowed, maxBalance, Organization, LeaveRequests, LeaveCategory, AccrualUnit): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLeavePolicy(name, accrualRate, carryoverAllowed, maxBalance, Organization, LeaveRequests, LeaveCategory, AccrualUnit, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLeavePolicy']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLeavePolicy(params['id']).subscribe(res => {
                this.leavePolicy = res;
            });
        });
    }
}