import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FeeScheduleService } from '../../../services/FeeSchedule.service';
import { SubBaseComponent } from '../../FeeSchedule/sub.base.component';


@Component({
    selector: 'app-edit-feeSchedule',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFeeScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Edit FeeSchedule';

    feeScheduleForm: FormGroup;
    feeSchedule: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FeeScheduleService,
        private fb: FormBuilder
) {
        super(http);
        this.feeScheduleForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      amount: ['', Validators.required],
      percentage: ['', Validators.required],
      minimum: ['', Validators.required],
      maximum: ['', Validators.required],
      PricingPlan: ['', ],
      FeeType: ['', ],
      CalculationMethod: ['', ]
        });
    }

    
    updateFeeSchedule(name, amount, percentage, minimum, maximum, PricingPlan, FeeType, CalculationMethod): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFeeSchedule(name, amount, percentage, minimum, maximum, PricingPlan, FeeType, CalculationMethod, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFeeSchedule']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFeeSchedule(params['id']).subscribe(res => {
                this.feeSchedule = res;
            });
        });
    }
}