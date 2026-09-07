import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PricingPlanService } from '../../../services/PricingPlan.service';
import { SubBaseComponent } from '../../PricingPlan/sub.base.component';


@Component({
    selector: 'app-edit-pricingPlan',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPricingPlanComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PricingPlan';

    pricingPlanForm: FormGroup;
    pricingPlan: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PricingPlanService,
        private fb: FormBuilder
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

    
    updatePricingPlan(name, planCode, baseCurrency, ProductOffering, FeeSchedules, Limits, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePricingPlan(name, planCode, baseCurrency, ProductOffering, FeeSchedules, Limits, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPricingPlan']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPricingPlan(params['id']).subscribe(res => {
                this.pricingPlan = res;
            });
        });
    }
}