import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MaintenancePlanService } from '../../../services/MaintenancePlan.service';
import { MaintenancePlan } from '../../../models/MaintenancePlan';
import { SubBaseComponent } from '../../MaintenancePlan/sub.base.component';

@Component({
    selector: 'app-create-maintenancePlan',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateMaintenancePlanComponent extends SubBaseComponent implements OnInit {

    title = 'Add MaintenancePlan';

    maintenancePlanForm: FormGroup;
    maintenancePlan: MaintenancePlan;

    constructor( http: HttpClient,
        private maintenancePlanService: MaintenancePlanService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.maintenancePlanForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  planNumber: ['', Validators.required],
      interval: ['', Validators.required],
      lastServiceDate: ['', Validators.required],
      Asset: ['', ],
      MaintenanceOrders: ['', ],
      Strategy: ['', ]
        });
    }

    
    addMaintenancePlan(planNumber, interval, lastServiceDate, Asset, MaintenanceOrders, Strategy): void {
        this.maintenancePlanService
        .addMaintenancePlan(planNumber, interval, lastServiceDate, Asset, MaintenanceOrders, Strategy)
            .subscribe(() => {
                this.router.navigate(['/indexMaintenancePlan']);
            });
    }

    ngOnInit(): void {
    }
}