import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FeeScheduleService } from '../../../services/FeeSchedule.service';
import { FeeSchedule } from '../../../models/FeeSchedule';
import { SubBaseComponent } from '../../FeeSchedule/sub.base.component';

@Component({
    selector: 'app-create-feeSchedule',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFeeScheduleComponent extends SubBaseComponent implements OnInit {

    title = 'Add FeeSchedule';

    feeScheduleForm: FormGroup;
    feeSchedule: FeeSchedule;

    constructor( http: HttpClient,
        private feeScheduleService: FeeScheduleService,
        private fb: FormBuilder,
        private router: Router
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

    
    addFeeSchedule(name, amount, percentage, minimum, maximum, PricingPlan, FeeType, CalculationMethod): void {
        this.feeScheduleService
        .addFeeSchedule(name, amount, percentage, minimum, maximum, PricingPlan, FeeType, CalculationMethod)
            .subscribe(() => {
                this.router.navigate(['/indexFeeSchedule']);
            });
    }

    ngOnInit(): void {
    }
}