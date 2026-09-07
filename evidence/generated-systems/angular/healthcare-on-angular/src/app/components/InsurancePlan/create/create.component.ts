import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InsurancePlanService } from '../../../services/InsurancePlan.service';
import { InsurancePlan } from '../../../models/InsurancePlan';
import { SubBaseComponent } from '../../InsurancePlan/sub.base.component';

@Component({
    selector: 'app-create-insurancePlan',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInsurancePlanComponent extends SubBaseComponent implements OnInit {

    title = 'Add InsurancePlan';

    insurancePlanForm: FormGroup;
    insurancePlan: InsurancePlan;

    constructor( http: HttpClient,
        private insurancePlanService: InsurancePlanService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.insurancePlanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      planCode: ['', Validators.required],
      Payer: ['', ],
      Coverages: ['', ],
      PlanType: ['', ]
        });
    }

    
    addInsurancePlan(name, planCode, Payer, Coverages, PlanType): void {
        this.insurancePlanService
        .addInsurancePlan(name, planCode, Payer, Coverages, PlanType)
            .subscribe(() => {
                this.router.navigate(['/indexInsurancePlan']);
            });
    }

    ngOnInit(): void {
    }
}