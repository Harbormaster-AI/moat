import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PricingPlanService } from '../../../services/PricingPlan.service';
import { PricingPlan } from '../../../models/PricingPlan';
import { SubBaseComponent } from '../../PricingPlan/sub.base.component';

@Component({
    selector: 'app-create-pricingPlan',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePricingPlanComponent extends SubBaseComponent implements OnInit {

    title = 'Add PricingPlan';

    pricingPlanForm: FormGroup;
    pricingPlan: PricingPlan;

    constructor( http: HttpClient,
        private pricingPlanService: PricingPlanService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.pricingPlanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      planCode: ['', Validators.required],
      baseCurrency: ['', Validators.required],
      ProductOffering: ['', ],
      FeeSchedules: ['', ],
      Limits: ['', ],
      Status: ['', ]
        });
    }

    
    addPricingPlan(name, planCode, baseCurrency, ProductOffering, FeeSchedules, Limits, Status): void {
        this.pricingPlanService
        .addPricingPlan(name, planCode, baseCurrency, ProductOffering, FeeSchedules, Limits, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPricingPlan']);
            });
    }

    ngOnInit(): void {
    }
}