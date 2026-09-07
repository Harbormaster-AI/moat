import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UsageLimitService } from '../../../services/UsageLimit.service';
import { UsageLimit } from '../../../models/UsageLimit';
import { SubBaseComponent } from '../../UsageLimit/sub.base.component';

@Component({
    selector: 'app-create-usageLimit',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateUsageLimitComponent extends SubBaseComponent implements OnInit {

    title = 'Add UsageLimit';

    usageLimitForm: FormGroup;
    usageLimit: UsageLimit;

    constructor( http: HttpClient,
        private usageLimitService: UsageLimitService,
        private fb: FormBuilder,
        private router: Router
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

    
    addUsageLimit(name, amount, count, PricingPlan, Scope, Period): void {
        this.usageLimitService
        .addUsageLimit(name, amount, count, PricingPlan, Scope, Period)
            .subscribe(() => {
                this.router.navigate(['/indexUsageLimit']);
            });
    }

    ngOnInit(): void {
    }
}