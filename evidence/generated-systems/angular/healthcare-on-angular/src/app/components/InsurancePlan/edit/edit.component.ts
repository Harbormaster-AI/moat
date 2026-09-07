import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InsurancePlanService } from '../../../services/InsurancePlan.service';
import { SubBaseComponent } from '../../InsurancePlan/sub.base.component';


@Component({
    selector: 'app-edit-insurancePlan',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInsurancePlanComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InsurancePlan';

    insurancePlanForm: FormGroup;
    insurancePlan: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InsurancePlanService,
        private fb: FormBuilder
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

    
    updateInsurancePlan(name, planCode, Payer, Coverages, PlanType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInsurancePlan(name, planCode, Payer, Coverages, PlanType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInsurancePlan']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInsurancePlan(params['id']).subscribe(res => {
                this.insurancePlan = res;
            });
        });
    }
}