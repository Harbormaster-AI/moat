import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ReplenishmentPolicyService } from '../../../services/ReplenishmentPolicy.service';
import { SubBaseComponent } from '../../ReplenishmentPolicy/sub.base.component';


@Component({
    selector: 'app-edit-replenishmentPolicy',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditReplenishmentPolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ReplenishmentPolicy';

    replenishmentPolicyForm: FormGroup;
    replenishmentPolicy: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ReplenishmentPolicyService,
        private fb: FormBuilder
) {
        super(http);
        this.replenishmentPolicyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  minLevel: ['', Validators.required],
      maxLevel: ['', Validators.required],
      reorderPoint: ['', Validators.required],
      reorderQuantity: ['', Validators.required],
      leadTimeDays: ['', Validators.required],
      reviewPeriodDays: ['', Validators.required],
      Sku: ['', ],
      Warehouse: ['', ],
      Location: ['', ],
      PolicyType: ['', ]
        });
    }

    
    updateReplenishmentPolicy(minLevel, maxLevel, reorderPoint, reorderQuantity, leadTimeDays, reviewPeriodDays, Sku, Warehouse, Location, PolicyType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateReplenishmentPolicy(minLevel, maxLevel, reorderPoint, reorderQuantity, leadTimeDays, reviewPeriodDays, Sku, Warehouse, Location, PolicyType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexReplenishmentPolicy']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getReplenishmentPolicy(params['id']).subscribe(res => {
                this.replenishmentPolicy = res;
            });
        });
    }
}