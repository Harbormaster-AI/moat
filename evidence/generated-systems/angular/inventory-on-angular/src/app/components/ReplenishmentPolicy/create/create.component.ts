import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ReplenishmentPolicyService } from '../../../services/ReplenishmentPolicy.service';
import { ReplenishmentPolicy } from '../../../models/ReplenishmentPolicy';
import { SubBaseComponent } from '../../ReplenishmentPolicy/sub.base.component';

@Component({
    selector: 'app-create-replenishmentPolicy',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateReplenishmentPolicyComponent extends SubBaseComponent implements OnInit {

    title = 'Add ReplenishmentPolicy';

    replenishmentPolicyForm: FormGroup;
    replenishmentPolicy: ReplenishmentPolicy;

    constructor( http: HttpClient,
        private replenishmentPolicyService: ReplenishmentPolicyService,
        private fb: FormBuilder,
        private router: Router
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

    
    addReplenishmentPolicy(minLevel, maxLevel, reorderPoint, reorderQuantity, leadTimeDays, reviewPeriodDays, Sku, Warehouse, Location, PolicyType): void {
        this.replenishmentPolicyService
        .addReplenishmentPolicy(minLevel, maxLevel, reorderPoint, reorderQuantity, leadTimeDays, reviewPeriodDays, Sku, Warehouse, Location, PolicyType)
            .subscribe(() => {
                this.router.navigate(['/indexReplenishmentPolicy']);
            });
    }

    ngOnInit(): void {
    }
}