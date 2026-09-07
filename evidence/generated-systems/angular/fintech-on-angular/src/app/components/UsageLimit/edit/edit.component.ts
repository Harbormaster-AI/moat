import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { UsageLimitService } from '../../../services/UsageLimit.service';
import { SubBaseComponent } from '../../UsageLimit/sub.base.component';


@Component({
    selector: 'app-edit-usageLimit',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditUsageLimitComponent extends SubBaseComponent implements OnInit {

    title = 'Edit UsageLimit';

    usageLimitForm: FormGroup;
    usageLimit: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: UsageLimitService,
        private fb: FormBuilder
) {
        super(http);
        this.usageLimitForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      amount: ['', Validators.required],
      count: ['', Validators.required],
      PricingPlan: ['', ],
      Scope: ['', ],
      Period: ['', ]
        });
    }

    
    updateUsageLimit(name, amount, count, PricingPlan, Scope, Period): void {
        this.route.params.subscribe((params) => {

                        this.service.updateUsageLimit(name, amount, count, PricingPlan, Scope, Period, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexUsageLimit']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getUsageLimit(params['id']).subscribe(res => {
                this.usageLimit = res;
            });
        });
    }
}